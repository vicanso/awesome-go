package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/vicanso/go-axios"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

const (
	apiRepoInfo = "https://api.github.com/repos/%s"
)

var repoReg = regexp.MustCompile(`\* \[[\s\S]+\]\(\S+\)`)
var githubRepoReg = regexp.MustCompile(`\(https://github.com/(\S+)\)`)
var ins = axios.NewInstance(&axios.InstanceConfig{
	Timeout: time.Minute,
})
var defaultLogger *zap.Logger

// 用于保存repo的star数量，用于在获取失败时使用
var originalRepoStars = map[string]int{}

// github api的认证参数
var clientId *string
var clientSecret *string

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
	url := fmt.Sprintf(apiRepoInfo, name)
	if *clientId != "" && *clientSecret != "" {
		url += fmt.Sprintf("?client_id=%s&client_secret=%s", *clientId, *clientSecret)
	}
	// 如果出错则认为0
	resp, err := ins.Get(url)
	if err != nil {
		return
	}
	info := &GithubRepoInfo{}
	err = resp.JSON(info)
	if err != nil {
		return
	}
	count = info.StargazersCount
	return
}

// 按star对同一分类的repo重排
func arrangeByStar(lines []string) (result []string, err error) {
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
				count, e := getStarCount(name)
				if e != nil {
					defaultLogger.Error("get repo info fail",
						zap.String("repo", name),
						zap.Error(e),
					)
					count = originalRepoStars[name]
				}

				repo.star = count
				if repo.star != 0 {
					stars := strconv.Itoa(repo.star)
					if repo.star >= 1000 {
						stars = fmt.Sprintf("%.1fK", float32(repo.star)/1000)
					}
					repo.content += fmt.Sprintf(" Stars:`%s`.", stars)
					originalRepoStars[name] = repo.star
				}
				defaultLogger.Info("get repo info success",
					zap.String("repo", name),
					zap.Int("star", count),
				)
			}
			repos = append(repos, repo)
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
	clientId = flag.String("clientId", "", "github oauth app client id")
	clientSecret = flag.String("clientSecret", "", "github oauth app client secret")
	flag.Parse()

	resp, err := ins.Get("https://raw.githubusercontent.com/avelino/awesome-go/master/README.md")
	if err != nil {
		panic(err)
	}
	lines := strings.SplitN(string(resp.Data), "\n", -1)

	result, err := arrangeByStar(lines)
	if err != nil {
		panic(err)
	}
	buf, _ := json.Marshal(originalRepoStars)
	if len(buf) != 0 {
		_ = ioutil.WriteFile(filepath.Join(os.TempDir()+"awesome-go.json"), buf, 0600)
	}
	err = ioutil.WriteFile("./README.md", []byte(strings.Join(result, "\n")), 0600)
	if err != nil {
		panic(err)
	}

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	auth, err := ssh.NewPublicKeysFromFile("git", user.HomeDir+"/.ssh/id_rsa", "")
	if err != nil {
		panic(err)
	}

	// 提交 commit
	r, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}
	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}
	status, err := w.Status()
	if err != nil {
		panic(err)
	}
	// 如果文件未改变
	if status.IsClean() {
		return
	}

	_, err = w.Add("README.md")
	if err != nil {
		panic(err)
	}

	commit, err := w.Commit("auto generate", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "vicanso",
			Email: "vicansocanbico@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		panic(err)
	}
	_, err = r.CommitObject(commit)
	if err != nil {
		panic(err)
	}
	err = r.Push(&git.PushOptions{
		Auth: auth,
	})
	if err != nil {
		panic(err)
	}
}
