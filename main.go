package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/vicanso/go-axios"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	apiRepoInfo = "https://api.github.com/repos/%s"
)

var repoReg = regexp.MustCompile(`\- \[[\s\S]+\]\(\S+\)`)
var githubRepoReg = regexp.MustCompile(`\(https://github.com/(\S+)\)`)
var ins = axios.NewInstance(&axios.InstanceConfig{
	Timeout: time.Minute,
})
var defaultLogger *zap.Logger

var startCounts = map[string]int{}

// github api的认证参数
var clientToken *string

type Repo struct {
	content string
	star    int
}
type Repos []*Repo

func (rs Repos) Len() int {
	return len(rs)
}
func (rs Repos) Less(i, j int) bool {
	return rs[i].star < rs[j].star
}
func (rs Repos) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

type GithubRepoInfo struct {
	StargazersCount int `json:"stargazers_count"`
}

func init() {
	c := zap.NewProductionConfig()
	c.DisableCaller = true
	c.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 只针对panic 以上的日志增加stack trace
	l, err := c.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		panic(err)
	}
	defaultLogger = l
}

// 获取star的数量
func getStarCount(name string) (count int, err error) {
	count, ok := startCounts[name]
	if ok {
		return count, nil
	}
	url := fmt.Sprintf(apiRepoInfo, name)
	// 如果出错则认为0
	headers := make(http.Header)
	headers.Add("Authorization", fmt.Sprintf("token %s", *clientToken))
	resp, err := ins.Request(&axios.Config{
		URL:     url,
		Headers: headers,
	})
	if err != nil {
		return
	}
	info := &GithubRepoInfo{}
	err = resp.JSON(info)
	if err != nil {
		return
	}
	count = info.StargazersCount
	startCounts[name] = count
	return
}

// 按star对同一分类的repo重排
func arrangeByStar(lines []string, min int) (result []string, err error) {
	result = make([]string, 0, len(lines))
	newCatStart := false

	var repos Repos
	for _, line := range lines {
		if repoReg.MatchString(line) {
			// 新的分类开始
			if !newCatStart {
				newCatStart = true
				repos = make(Repos, 0)
			}
			subs := githubRepoReg.FindAllStringSubmatch(line, -1)
			repo := &Repo{
				content: line,
			}
			if len(subs) > 0 && len(subs[0]) == 2 {
				name := subs[0][1]
				nameSplits := strings.Split(name, "/")
				if len(nameSplits) > 2 {
					name = strings.Join(nameSplits[0:2], "/")
				}
				count, e := getStarCount(name)
				if e != nil {
					defaultLogger.Error("get repo info fail",
						zap.String("repo", name),
						zap.Error(e),
					)
				}

				repo.star = count
				if repo.star != 0 {
					stars := strconv.Itoa(repo.star)
					if repo.star >= 1000 {
						stars = fmt.Sprintf("%.1fK", float32(repo.star)/1000)
					}
					repo.content += fmt.Sprintf(" Stars:`%s`.", stars)
				}
				defaultLogger.Info("get repo info success",
					zap.String("repo", name),
					zap.Int("star", count),
				)
			}
			if repo.star > min {
				repos = append(repos, repo)
			}
			continue
		}
		// 如果开始不匹配，则该分类结束
		if newCatStart {
			newCatStart = false
			// 对repo按star从高至低排序
			sort.Sort(sort.Reverse(repos))
			for _, item := range repos {
				result = append(result, item.content)
			}
		}
		result = append(result, line)
	}
	return
}

func main() {
	// 如果没有不匹配，调用github api的频率限制较低
	clientToken = flag.String("clientToken", "", "github personal access token")
	flag.Parse()

	resp, err := ins.Get("https://raw.githubusercontent.com/avelino/awesome-go/main/README.md")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitN(string(resp.Data), "\n", -1)

	result, err := arrangeByStar(lines, -1)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("./README.md", []byte(strings.Join(result, "\n")), 0600)
	if err != nil {
		panic(err)
	}

	result, err = arrangeByStar(lines, 1000)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("./README-1k.md", []byte(strings.Join(result, "\n")), 0600)
	if err != nil {
		panic(err)
	}
}
