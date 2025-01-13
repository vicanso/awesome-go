# Awesome Go

<a href="https://awesome-go.com/"><img align="right" src="https://github.com/avelino/awesome-go/raw/main/tmpl/assets/logo.png" alt="awesome-go" title="awesome-go" /></a>

[![Build Status](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml/badge.svg?branch=main)](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml?query=branch%3Amain)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)
[![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](https://gophers.slack.com/messages/awesome)
[![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)
[![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/avelino/awesome-go/)
[![Last Commit](https://img.shields.io/github/last-commit/avelino/awesome-go)](https://img.shields.io/github/last-commit/avelino/awesome-go)

We use the _[Golang Bridge](https://github.com/gobridge/about-us/blob/master/README.md)_ community Slack for instant communication, follow the [form here to join](https://invite.slack.golangbridge.org/).

<a href="https://www.producthunt.com/posts/awesome-go?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-awesome-go" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=291535&theme=light" alt="awesome-go - Curated list awesome Go frameworks, libraries and software | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

**Sponsorships:**

_Special thanks to_

<div align="center">
<table cellpadding="5">
<tbody align="center">
<tr>
<td colspan="2">
<a href="https://bit.ly/awesome-go-workos">
<img src="https://avelino.run/sponsors/workos-logo-white-bg.svg" width="200" alt="WorkOS"><br/>
<b>Your app, enterprise-ready.</b><br/>
<sub>Start selling to enterprise customers with just a few lines of code.</sub><br/>
<sup>Add Single Sign-On (and more) in minutes instead of months.</sup>
</a>
</td>
</tr>
<tr>
<td colspan="2">
<a href="https://bit.ly/awesome-go-digitalocean">
<img src="https://avelino.run/sponsors/do_logo_horizontal_blue-210.png" width="200" alt="Digital Ocean">
</a>
</td>
</tr>
</tbody>
</table>
</div>

**Awesome Go has no monthly fee**_, but we have employees who **work hard** to keep it running. With money raised, we can repay the effort of each person involved! You can see how we calculate our billing and distribution as it is open to the entire community. Want to be a supporter of the project click [here](mailto:avelinorun+oss@gmail.com?subject=awesome-go%3A%20project%20support)._

> A curated list of awesome Go frameworks, libraries, and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

**Contributing:**

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/main/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

> _If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!_

## Contents


**[⬆ back to top](#contents)**

## Actor Model

_Libraries for building actor-based programs._

- [ProtoActor](https://github.com/asynkron/protoactor-go) - Proto Actor - Ultra fast distributed actors for Go, C# and Java/Kotlin. Stars:`5.1K`.
- [Ergo](https://github.com/ergo-services/ergo) - An actor-based Framework with network transparency for creating event-driven architecture in Golang. Inspired by Erlang. Stars:`3.8K`.
- [Hollywood](https://github.com/anthdm/hollywood) - Blazingly fast and light-weight Actor engine written in Golang. Stars:`1.5K`.

## Artificial Intelligence

_Libraries for building programs that leverage AI._

- [Ollama](https://github.com/jmorganca/ollama) - Run large language models locally. Stars:`107.3K`.
- [LocalAI](https://github.com/mudler/LocalAI) - Open Source OpenAI alternative, self-host AI models. Stars:`28.3K`.
- [langchaingo](https://github.com/tmc/langchaingo) - LangChainGo is a framework for developing applications powered by language models. Stars:`5.3K`.

**[⬆ back to top](#contents)**

## Audio and Music

_Libraries for manipulating audio._

- [Oto](https://github.com/hajimehoshi/oto) - A low-level library to play sound on multiple platforms. Stars:`1.6K`.

**[⬆ back to top](#contents)**

## Authentication and OAuth

_Libraries for implementing authentication schemes._

- [casbin](https://github.com/hsluoyz/casbin) - Authorization library that supports access control models like ACL, RBAC, and ABAC. Stars:`18.0K`.
- [jwt-go](https://github.com/golang-jwt/jwt) - A full featured implementation of JSON Web Tokens (JWT). This library supports the parsing and verification as well as the generation and signing of JWTs. Stars:`7.5K`.
- [goth](https://github.com/markbates/goth) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. Stars:`5.7K`.
- [oauth2](https://github.com/golang/oauth2) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine, and App Engine support. Stars:`5.5K`.
- [keto](https://github.com/ory/keto) - Open Source (Go) implementation of "Zanzibar: Google's Consistent, Global Authorization System". Ships gRPC, REST APIs, newSQL, and an easy and granular permission language. Supports ACL, RBAC, and other access models. Stars:`4.9K`.
- [authboss](https://github.com/volatiletech/authboss) - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure it, and start building your app without having to build an authentication system each time. Stars:`3.9K`.
- [openfga](https://github.com/openfga/openfga) - Implementation of fine-grained authorization based on the "Zanzibar: Google's Consistent, Global Authorization System" paper. Backed by [CNCF](https://www.cncf.io/). Stars:`3.1K`.
- [scs](https://github.com/alexedwards/scs) - Session Manager for HTTP servers. Stars:`2.2K`.
- [jwx](https://github.com/lestrrat-go/jwx) - Go module implementing various JWx (JWA/JWE/JWK/JWS/JWT, otherwise known as JOSE) technologies Stars:`2.0K`.
- [loginsrv](https://github.com/tarent/loginsrv) - JWT login microservice with pluggable backends such as OAuth2 (Github), htpasswd, osiam. Stars:`1.9K`.
- [osin](https://github.com/openshift/osin) - Golang OAuth2 server library. Stars:`1.9K`.
- [gologin](https://github.com/dghubble/gologin) - chainable handlers for login with OAuth1 and OAuth2 authentication providers. Stars:`1.9K`.
- [gorbac](https://github.com/mikespook/gorbac) - provides a lightweight role-based access control (RBAC) implementation in Golang. Stars:`1.6K`.
- [oidc](https://github.com/zitadel/oidc) - Easy to use OpenID Connect client and server library written for Go and certified by the OpenID Foundation Stars:`1.5K`.

**[⬆ back to top](#contents)**

## Blockchain

_Tools for building blockchains._

- [go-ethereum](https://github.com/ethereum/go-ethereum) - Official Go implementation of the Ethereum protocol. Stars:`48.1K`.
- [kubo](https://github.com/ipfs/kubo) - A blockchain framework implemented in Go. It provides content-addressable storage which can be used for decentralized storage in DApps. It is based on the IPFS protocol. Stars:`16.3K`.
- [lnd](https://github.com/lightningnetwork/lnd) - A complete implementation of a Lightning Network node. Stars:`7.8K`.
- [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) - A Framework for Building Public Blockchains in the Cosmos Ecosystem. Stars:`6.4K`.
- [tendermint](https://github.com/tendermint/tendermint) - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. Stars:`5.7K`.
- [solana-go](https://github.com/gagliardetto/solana-go) - Go library to interface with Solana JSON RPC and WebSocket interfaces. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Bot Building

_Libraries for building and working with bots._

- [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) - Simple and clean Telegram bot client. Stars:`5.9K`.
- [telebot](https://github.com/tucnak/telebot) - Telegram bot framework is written in Go. Stars:`4.1K`.
- [olivia](https://github.com/olivia-ai/olivia) - A chatbot built with an artificial neural network. Stars:`3.7K`.
- [wayback](https://github.com/wabarc/wayback) - A bot for Telegram, Mastodon, Slack, and other messaging platforms archives webpages. Stars:`1.9K`.
- [Kelp](https://github.com/stellar/kelp) - official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. Stars:`1.1K`.
- [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) - A golang implementation of a console-based trading bot for cryptocurrency exchanges. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Build Automation

_Libraries and tools help with build automation._

- [air](https://github.com/cosmtrek/air) - Air - Live reload for Go apps. Stars:`18.9K`.
- [Task](https://github.com/go-task/task) - simple "Make" alternative. Stars:`11.9K`.
- [realize](https://github.com/tockins/realize) - Go build a system with file watchers and live to reload. Run, build and watch file changes with custom paths. Stars:`4.5K`.
- [mage](https://github.com/magefile/mage) - Mage is a make/rake-like build tool using Go. Stars:`4.2K`.
- [mmake](https://github.com/tj/mmake) - Modern Make. Stars:`1.7K`.
- [xc](https://github.com/joerdav/xc) - Task runner with README.md defined tasks, executable markdown. Stars:`1.2K`.

**[⬆ back to top](#contents)**

## Command Line

### Advanced Console UIs

_Libraries for building Console Applications and Console User Interfaces._

- [bubbletea](https://github.com/charmbracelet/bubbletea) - Go framework to build terminal apps, based on The Elm Architecture. Stars:`29.0K`.
- [fx](https://github.com/antonmedv/fx) - Terminal JSON viewer & processor. Stars:`19.2K`.
- [termui](https://github.com/gizak/termui) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). Stars:`13.3K`.
- [gocui](https://github.com/jroimartin/gocui) - Minimalist Go library aimed at creating Console User Interfaces. Stars:`10.0K`.
- [lipgloss](https://github.com/charmbracelet/lipgloss) - Declaratively define styles for color, format and layout in the terminal. Stars:`8.4K`.
- [bubbles](https://github.com/charmbracelet/bubbles) - TUI components for bubbletea. Stars:`5.8K`.
- [go-prompt](https://github.com/c-bata/go-prompt) - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). Stars:`5.3K`.
- [pterm](https://github.com/pterm/pterm) - A library to beautify console output on every platform with many combinable components. Stars:`4.9K`.
- [termbox-go](https://github.com/nsf/termbox-go) - Termbox is a library for creating cross-platform text-based interfaces. Stars:`4.7K`.
- [progressbar](https://github.com/schollz/progressbar) - Basic thread-safe progress bar that works in every OS. Stars:`4.2K`.
- [termdash](https://github.com/mum4k/termdash) - Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). Stars:`2.8K`.
- [asciigraph](https://github.com/guptarohit/asciigraph) - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. Stars:`2.7K`.
- [spinner](https://github.com/briandowns/spinner) - Go package to easily provide a terminal spinner with options. Stars:`2.4K`.
- [mpb](https://github.com/vbauerster/mpb) - Multi progress bar for terminal applications. Stars:`2.4K`.
- [uiprogress](https://github.com/gosuri/uiprogress) - Flexible library to render progress bars in terminal applications. Stars:`2.1K`.
- [termenv](https://github.com/muesli/termenv) - Advanced ANSI style & color support for your terminal applications. Stars:`1.8K`.
- [uilive](https://github.com/gosuri/uilive) - Library for updating terminal output in real time. Stars:`1.7K`.
- [gookit/color](https://github.com/gookit/color) - Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. Stars:`1.5K`.
- [aurora](https://github.com/logrusorgru/aurora) - ANSI terminal colors that support fmt.Printf/Sprintf. Stars:`1.4K`.

**[⬆ back to top](#contents)**

### Standard CLI

_Libraries for building standard or basic Command Line applications._

- [cobra](https://github.com/spf13/cobra) - Commander for modern Go CLI interactions. Stars:`38.8K`.
- [urfave/cli](https://github.com/urfave/cli) - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). Stars:`22.6K`.
- [elvish](https://github.com/elves/elvish) - An expressive programming language and a versatile interactive shell. Stars:`5.8K`.
- [survey](https://github.com/go-survey/survey) - Build interactive and accessible prompts with full support for windows and posix terminals. Stars:`4.1K`.
- [kingpin](https://github.com/alecthomas/kingpin) - Command line and flag parser supporting sub commands (superseded by `kong`; see below). Stars:`3.5K`.
- [Dnote](https://github.com/dnote/dnote) - A simple command line notebook with multi-device sync. Stars:`2.8K`.
- [go-flags](https://github.com/jessevdk/go-flags) - go command line option parser. Stars:`2.6K`.
- [pflag](https://github.com/spf13/pflag) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. Stars:`2.5K`.
- [go-arg](https://github.com/alexflint/go-arg) - Struct-based argument parsing in Go. Stars:`2.1K`.
- [ops](https://github.com/nanovms/ops) - Unikernel Builder/Orchestrator. Stars:`1.3K`.
- [liner](https://github.com/peterh/liner) - Go readline-like library for command-line interfaces. Stars:`1.1K`.
- [carapace-bin](https://github.com/rsteube/carapace-bin) - Multi-shell multi-command argument completer. Stars:`1.0K`.

**[⬆ back to top](#contents)**

## Configuration

_Libraries for configuration parsing._

- [viper](https://github.com/spf13/viper) - Go configuration with fangs. Stars:`27.6K`.
- [godotenv](https://github.com/joho/godotenv) - Go port of Ruby's dotenv library (Loads environment variables from `.env`). Stars:`8.6K`.
- [sonic](https://github.com/bytedance/sonic) - A blazingly fast JSON serializing & deserializing library. Stars:`7.2K`.
- [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) - Go library for managing configuration data from environment variables. Stars:`5.1K`.
- [env](https://github.com/caarlos0/env) - Parse environment variables to Go structs (with defaults). Stars:`5.1K`.
- [ini](https://github.com/go-ini/ini) - Go package to read and write INI files. Stars:`3.5K`.
- [koanf](https://github.com/knadh/koanf) - Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. Stars:`2.9K`.
- [kong](https://github.com/alecthomas/kong) - Command-line parser with support for arbitrarily complex command-line structures and additional sources of configuration such as YAML, JSON, TOML, etc (successor to `kingpin`). Stars:`2.3K`.
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Minimalistic configuration reader (from files, ENV, and wherever you want). Stars:`1.7K`.

**[⬆ back to top](#contents)**

## Continuous Integration

_Tools for help with continuous integration._

- [drone](https://github.com/drone/drone) - Drone is a Continuous Integration platform built on Docker, written in Go. Stars:`32.4K`.
- [CDS](https://github.com/ovh/cds) - Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. Stars:`4.6K`.
- [woodpecker](https://github.com/woodpecker-ci/woodpecker) - Woodpecker is a community fork of the Drone CI system. Stars:`4.5K`.
- [muffet](https://github.com/raviqqe/muffet) - Fast website link checker in Go, see [alternatives](https://github.com/lycheeverse/lychee#features). Stars:`2.5K`.

**[⬆ back to top](#contents)**

## CSS Preprocessors

_Libraries for preprocessing CSS files._


**[⬆ back to top](#contents)**

## Data Integration Frameworks

_Frameworks for performing ELT / ETL_

- [Benthos](https://github.com/benthosdev/benthos) - A message streaming bridge between a range of protocols. Stars:`8.2K`.

**[⬆ back to top](#contents)**

## Data Structures and Algorithms

### Bit-packing and Compression

- [roaring](https://github.com/RoaringBitmap/roaring) - Go package implementing compressed bitsets. Stars:`2.6K`.

### Bit Sets

- [bitset](https://github.com/bits-and-blooms/bitset) - Go package implementing bitsets. Stars:`1.4K`.

### Bloom and Cuckoo Filters

- [bloom](https://github.com/bits-and-blooms/bloom) - Go package implementing Bloom filters. Stars:`2.5K`.
- [boomfilters](https://github.com/tylertreat/BoomFilters) - Probabilistic data structures for processing continuous, unbounded streams. Stars:`1.6K`.
- [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. Stars:`1.2K`.

### Data Structure and Algorithm Collections

- [gods](https://github.com/emirpasic/gods) - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. Stars:`16.6K`.
- [go-datastructures](https://github.com/Workiva/go-datastructures) - Collection of useful, performant, and thread-safe data structures. Stars:`7.7K`.
- [gostl](https://github.com/liyue201/gostl) - Data structure and algorithm library for go, designed to provide functions similar to C++ STL. Stars:`1.1K`.

### Iterators


### Maps

See also [Database](#database) for more complex key-value stores, and [Trees](#trees) for
additional ordered map implementations.


### Miscellaneous Data Structures and Algorithms

- [gota](https://github.com/kniren/gota) - Implementation of dataframes, series, and data wrangling methods for Go. Stars:`3.1K`.

### Nullable Types


### Queues

- [hatchet](https://github.com/hatchet-dev/hatchet) - Distributed, Fault-tolerant task queue. Stars:`4.4K`.

### Sets

- [golang-set](https://github.com/deckarep/golang-set) - Thread-Safe and Non-Thread-Safe high-performance sets for Go. Stars:`4.3K`.

### Text Analysis

- [bleve](https://github.com/blevesearch/bleve) - Modern text indexing library for go. Stars:`10.2K`.

### Trees


### Pipes


**[⬆ back to top](#contents)**

## Database

### Caches

_Data stores with expiring records, in-memory distributed data stores, or in-memory subsets of file-based databases._

- [groupcache](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. Stars:`13.0K`.
- [BigCache](https://github.com/allegro/bigcache) - Efficient key/value cache for gigabytes of data. Stars:`7.7K`.
- [GCache](https://github.com/bluele/gcache) - Cache library with support for expirable Cache, LFU, LRU and ARC. Stars:`2.6K`.
- [gocache](https://github.com/eko/gocache) - A complete Go cache library with multiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. Stars:`2.5K`.
- [fastcache](https://github.com/VictoriaMetrics/fastcache) - fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. Stars:`2.2K`.
- [cache2go](https://github.com/muesli/cache2go) - In-memory key:value cache which supports automatic invalidation based on timeouts. Stars:`2.1K`.
- [otter](https://github.com/maypok86/otter) - A high performance lockless cache for Go. Many times faster than Ristretto and friends. Stars:`1.8K`.

### Databases Implemented in Go

- [prometheus](https://github.com/prometheus/prometheus) - Monitoring system and time series database. Stars:`56.8K`.
- [tidb](https://github.com/pingcap/tidb) - TiDB is a distributed SQL database. Inspired by the design of Google F1. Stars:`37.7K`.
- [Milvus](https://github.com/milvus-io/milvus) - Milvus is a vector database for embedding management, analytics and search. Stars:`31.8K`.
- [cockroach](https://github.com/cockroachdb/cockroach) - Scalable, Geo-Replicated, Transactional Datastore. Stars:`30.4K`.
- [influxdb](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics. Stars:`29.3K`.
- [dgraph](https://github.com/dgraph-io/dgraph) - Scalable, Distributed, Low Latency, High Throughput Graph Database. Stars:`20.6K`.
- [dolt](https://github.com/dolthub/dolt) - Dolt – It's Git for Data. Stars:`18.1K`.
- [rqlite](https://github.com/rqlite/rqlite) - The lightweight, distributed, relational database built on SQLite. Stars:`16.0K`.
- [badger](https://github.com/dgraph-io/badger) - Fast key-value store in Go. Stars:`14.2K`.
- [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) - fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. Stars:`13.0K`.
- [immudb](https://github.com/codenotary/immudb) - immudb is a lightweight, high-speed immutable database for systems and applications written in Go. Stars:`8.7K`.
- [bbolt](https://github.com/etcd-io/bbolt) - An embedded key/value database for Go. Stars:`8.4K`.
- [goleveldb](https://github.com/syndtr/goleveldb) - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. Stars:`6.2K`.
- [pebble](https://github.com/cockroachdb/pebble) - RocksDB/LevelDB inspired key-value database in Go. Stars:`5.1K`.
- [rosedb](https://github.com/roseduan/rosedb) - An embedded k-v database based on LSM+WAL, supports string, list, hash, set, zset. Stars:`4.7K`.
- [buntdb](https://github.com/tidwall/buntdb) - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. Stars:`4.6K`.
- [ledisdb](https://github.com/siddontang/ledisdb) - Ledisdb is a high performance NoSQL like Redis based on LevelDB. Stars:`4.1K`.
- [redka](https://github.com/nalgeon/redka) - Redis re-implemented with SQLite. Stars:`3.6K`.
- [godis](https://github.com/hdt3213/godis) - A Golang implemented high-performance Redis server and cluster. Stars:`3.6K`.
- [nutsdb](https://github.com/xujiajun/nutsdb) - Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as list, set, sorted set. Stars:`3.4K`.
- [LinDB](https://github.com/lindb/lindb) - LinDB is a scalable, high performance, high availability distributed time series database. Stars:`3.0K`.
- [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang. Stars:`2.7K`.
- [lotusdb](https://github.com/flower-corp/lotusdb) - Fast k/v database compatible with lsm and b+tree. Stars:`2.1K`.
- [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) - CovenantSQL is a SQL database on blockchain. Stars:`1.5K`.
- [column](https://github.com/kelindar/column) - High-performance, columnar, embeddable in-memory store with bitmap indexing and transactions. Stars:`1.5K`.
- [diskv](https://github.com/peterbourgon/diskv) - Home-grown disk-backed key-value store. Stars:`1.4K`.
- [pogreb](https://github.com/akrylysov/pogreb) - Embedded key-value store for read-heavy workloads. Stars:`1.3K`.
- [Databunker](https://github.com/paranoidguy/databunker) - Personally identifiable information (PII) storage service built to comply with GDPR and CCPA. Stars:`1.3K`.
- [objectbox-go](https://github.com/objectbox/objectbox-go) - High-performance embedded Object Database (NoSQL) with Go API. Stars:`1.1K`.
- [eliasdb](https://github.com/krotik/eliasdb) - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. Stars:`1.0K`.

### Database Schema Migration

- [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library. Stars:`15.9K`.
- [bytebase](https://github.com/bytebase/bytebase) - Safe database schema change and version control for DevOps teams. Stars:`11.8K`.
- [goose](https://github.com/pressly/goose) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. Stars:`7.4K`.
- [atlas](https://github.com/ariga/atlas) - A Database Toolkit. A CLI designed to help companies better work with their data. Stars:`6.3K`.
- [dbmate](https://github.com/amacneil/dbmate) - A lightweight, framework-agnostic database migration tool. Stars:`5.6K`.
- [sql-migrate](https://github.com/rubenv/sql-migrate) - Database migration tool. Allows embedding migrations into the application using go-bindata. Stars:`3.3K`.
- [soda](https://github.com/gobuffalo/pop/tree/master/soda) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. Stars:`1.5K`.
- [skeema](https://github.com/skeema/skeema) - Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. Stars:`1.3K`.
- [gormigrate](https://github.com/go-gormigrate/gormigrate) - Database schema migration helper for Gorm ORM. Stars:`1.1K`.

### Database Tools

- [vitess](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. Stars:`19.0K`.
- [pgweb](https://github.com/sosedoff/pgweb) - Web-based PostgreSQL database browser. Stars:`8.7K`.
- [go-mysql](https://github.com/siddontang/go-mysql) - Go toolset to handle MySQL protocol and replication. Stars:`4.6K`.
- [pREST](https://github.com/prest/prest) - Simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new. Stars:`4.3K`.
- [chproxy](https://github.com/Vertamedia/chproxy) - HTTP proxy for ClickHouse database. Stars:`1.3K`.
- [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable) - Advanced scheduling for PostgreSQL. Stars:`1.1K`.

### SQL Query Builders

_Libraries for building and using SQL._

- [sqlc](https://github.com/kyleconroy/sqlc) - Generate type-safe code from SQL. Stars:`14.0K`.
- [Squirrel](https://github.com/Masterminds/squirrel) - Go library that helps you build SQL queries. Stars:`7.1K`.
- [xo](https://github.com/knq/xo) - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. Stars:`3.8K`.
- [jet](https://github.com/go-jet/jet) - Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. Stars:`2.8K`.
- [goqu](https://github.com/doug-martin/goqu) - Idiomatic SQL builder and query library. Stars:`2.4K`.
- [gendry](https://github.com/didi/gendry) - Non-invasive SQL builder and powerful data binder. Stars:`1.6K`.

**[⬆ back to top](#contents)**

## Database Drivers

### Interfaces to Multiple Backends

- [cayley](https://github.com/google/cayley) - Graph database with support for multiple backends. Stars:`14.9K`.

### Relational Database Drivers

- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go. Stars:`14.6K`.
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver supporting features beyond those exposed by database/sql. Stars:`11.1K`.
- [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql. Stars:`9.2K`.
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that uses database/sql. Stars:`8.2K`.
- [go-mssqldb](https://github.com/denisenkom/go-mssqldb) - Microsoft MSSQL driver for Go. Stars:`1.8K`.

### NoSQL Database Drivers

- [redis](https://github.com/redis/go-redis) - Redis client for Golang. Stars:`20.4K`.
- [redigo](https://github.com/gomodule/redigo) - Redigo is a Go client for the Redis database. Stars:`9.8K`.
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver for the Go language. Stars:`8.2K`.
- [mgo](https://github.com/globalsign/mgo) - (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. Stars:`2.0K`.
- [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language. Stars:`1.8K`.
- [gorethink](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB. Stars:`1.7K`.
- [qmgo](https://github.com/qiniu/qmgo) - The MongoDB driver for Go. It‘s based on official MongoDB driver but easier to use like Mgo. Stars:`1.3K`.

### Search and Analytic Databases

- [elastic](https://github.com/olivere/elastic) - Elasticsearch client for Go. Stars:`7.4K`.
- [go-elasticsearch](https://github.com/elastic/go-elasticsearch) - Official Elasticsearch client for Go. Stars:`5.7K`.
- [clickhouse-go](https://github.com/ClickHouse/clickhouse-go/) - ClickHouse SQL client for Go with a `database/sql` compatibility. Stars:`3.0K`.
- [elasticsql](https://github.com/cch123/elasticsql) - Convert sql to elasticsearch dsl in Go. Stars:`1.2K`.

**[⬆ back to top](#contents)**

## Date and Time

_Libraries for working with dates and times._

- [now](https://github.com/jinzhu/now) - Now is a time toolkit for golang. Stars:`4.5K`.
- [dateparse](https://github.com/araddon/dateparse) - Parse date's without knowing format in advance. Stars:`2.1K`.

**[⬆ back to top](#contents)**

## Distributed Systems

_Packages that help with building Distributed Systems._

- [go-zero](https://github.com/tal-tech/go-zero) - A web and rpc framework. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity. Stars:`29.7K`.
- [go-kit](https://github.com/go-kit/kit) - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. Stars:`26.7K`.
- [Kratos](https://github.com/go-kratos/kratos) - A modular-designed and easy-to-use microservices framework in Go. Stars:`23.6K`.
- [go-micro](https://github.com/micro/go-micro) - A distributed systems development framework. Stars:`22.1K`.
- [grpc-go](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC. Stars:`21.3K`.
- [NATS](https://github.com/nats-io/nats-server) - NATS is a simple, secure, and Stars:`16.3K`.
- [micro](https://github.com/micro/micro) - A distributed systems runtime for the cloud and beyond. Stars:`12.2K`.
- [Kitex](https://github.com/cloudwego/kitex) - A high-performance and strong-extensibility Golang RPC framework that helps developers build microservices. If the performance and extensibility are the main concerns when you develop microservices, Kitex can be a good choice. Stars:`7.2K`.
- [lura](https://github.com/luraproject/lura) - Ultra performant API Gateway framework with middlewares. Stars:`6.4K`.
- [dragonboat](https://github.com/lni/dragonboat) - A feature complete and high performance multi-group Raft library in Go. Stars:`5.1K`.
- [evans](https://github.com/ktr0731/evans) - Evans: more expressive universal gRPC client. Stars:`4.3K`.
- [emitter-io](https://github.com/emitter-io/emitter) - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. Stars:`3.9K`.
- [gleam](https://github.com/chrislusf/gleam) - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. Stars:`3.5K`.
- [glow](https://github.com/chrislusf/glow) - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. Stars:`3.2K`.
- [liftbridge](https://github.com/liftbridge-io/liftbridge) - Lightweight, fault-tolerant message streams for NATS. Stars:`2.6K`.
- [Dragonfly](https://github.com/dragonflyoss/Dragonfly2) - Provide efficient, stable and secure file distribution and image acceleration based on p2p technology to be the best practice and standard solution in cloud native architectures. Stars:`2.4K`.
- [go-eagle](https://github.com/go-eagle/eagle) - A Go framework for the API or Microservice with handy scaffolding tools. Stars:`2.2K`.
- [go-doudou](https://github.com/unionj-cloud/go-doudou) - A gossip protocol and OpenAPI 3.0 spec based decentralized microservice framework. Built-in go-doudou cli focusing on low-code and rapid dev can power up your productivity. Stars:`1.4K`.
- [mochi mqtt](https://github.com/mochi-co/mqtt) - Fully spec compliant, embeddable high-performance MQTT v5/v3 broker for IoT, smarthome, and pubsub. Stars:`1.4K`.
- [hprose](https://github.com/hprose/hprose-golang) - Very newbility RPC Library, support 25+ languages now. Stars:`1.3K`.
  performant communications system for digital systems, services, and devices.
- [raft](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol, by HashiCorp. Stars:`8.4K`.
- [rpcx](https://github.com/smallnest/rpcx) - Distributed pluggable RPC service framework like alibaba Dubbo. Stars:`8.1K`.
- [torrent](https://github.com/anacrolix/torrent) - BitTorrent client package. Stars:`5.6K`.
- [sponge](https://github.com/zhufuyi/sponge) - A distributed development framework that integrates automatic code generation, gin and grpc frameworks, base development frameworks. Stars:`1.6K`.
- [redis-lock](https://github.com/bsm/redislock) - Simplified distributed locking implementation using Redis. Stars:`1.5K`.

**[⬆ back to top](#contents)**

## Dynamic DNS

_Tools for updating dynamic DNS records._

- [GoDNS](https://github.com/timothyye/godns) - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. Stars:`1.5K`.

**[⬆ back to top](#contents)**

## Email

_Libraries and tools that implement email creation and sending._

- [MailHog](https://github.com/mailhog/MailHog) - Email and SMTP testing with web and API interface. Stars:`14.4K`.
- [Mailpit](https://github.com/axllent/mailpit) - Email and SMTP testing tool for developers. Stars:`6.4K`.
- [Maddy](https://github.com/foxcpp/maddy) - All-in-one (SMTP, IMAP, DKIM, DMARC, MTA-STS, DANE) email server Stars:`5.2K`.
- [mox](https://github.com/mjl-/mox) - Modern full-featured secure mail server for low-maintenance, self-hosted email. Stars:`3.8K`.
- [hermes](https://github.com/matcornic/hermes) - Golang package that generates clean, responsive HTML e-mails. Stars:`2.8K`.
- [email](https://github.com/jordan-wright/email) - A robust and flexible email library for Go. Stars:`2.7K`.
- [go-imap](https://github.com/emersion/go-imap) - IMAP library for clients and servers. Stars:`2.1K`.
- [email-verifier](https://github.com/AfterShip/email-verifier) - A Go library for email verification without sending any emails. Stars:`1.3K`.

**[⬆ back to top](#contents)**

## Embeddable Scripting Languages

_Embedding other languages inside your go code._

- [expr](https://github.com/antonmedv/expr) - Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing. Stars:`6.5K`.
- [gopher-lua](https://github.com/yuin/gopher-lua) - Lua 5.1 VM and compiler written in Go. Stars:`6.4K`.
- [goja](https://github.com/dop251/goja) - ECMAScript 5.1(+) implementation in Go. Stars:`5.8K`.
- [tengo](https://github.com/d5/tengo) - Bytecode compiled script language for Go. Stars:`3.6K`.
- [go-lua](https://github.com/Shopify/go-lua) - Port of the Lua 5.2 VM to pure Go. Stars:`3.2K`.
- [starlark-go](https://github.com/google/starlark-go) - Go implementation of Starlark: Python-like language with deterministic evaluation and hermetic execution. Stars:`2.4K`.
- [cel-go](https://github.com/google/cel-go) - Fast, portable, non-Turing complete expression evaluation with gradual typing. Stars:`2.4K`.
- [metacall](https://github.com/metacall/core) - Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, WebAssembly, Java, Cobol and more. Stars:`1.6K`.
- [go-python](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API. Stars:`1.5K`.
- [anko](https://github.com/mattn/anko) - Scriptable interpreter written in Go. Stars:`1.5K`.
- [Wa/凹语言](https://github.com/wa-lang/wa) - The Wa Programming Language embedded in Go. Stars:`1.4K`.

**[⬆ back to top](#contents)**

## Error Handling

_Libraries for handling errors._

- [go-multierror](https://github.com/hashicorp/go-multierror) - Go (golang) package for representing a list of errors as a single error. Stars:`2.4K`.
- [errors](https://github.com/cockroachdb/errors) - Go error library with error portability over the network. Stars:`2.1K`.
- [eris](https://github.com/rotisserie/eris) - A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. Stars:`1.5K`.
- [errorx](https://github.com/joomcode/errorx) - A feature rich error package with stack traces, composition of errors and more. Stars:`1.2K`.
- [multierr](https://github.com/uber-go/multierr) - Package for representing a list of errors as a single error. Stars:`1.1K`.
- [tracerr](https://github.com/ztrue/tracerr) - Golang errors with stack trace and source fragments. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## File Handling

_Libraries for handling files and file systems._

- [pdfcpu](https://github.com/pdfcpu/pdfcpu) - PDF processor. Stars:`7.2K`.
- [afero](https://github.com/spf13/afero) - FileSystem Abstraction System for Go. Stars:`6.0K`.
- [gdu](https://github.com/dundee/gdu) - Disk usage analyzer with console interface. Stars:`4.1K`.
- [go-wkhtmltopdf](https://github.com/SebastiaanKlippert/go-wkhtmltopdf) - A package to convert an HTML template to a PDF file. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Financial

_Packages for accounting and finance._

- [decimal](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers. Stars:`6.5K`.
- [ticker](https://github.com/achannarasappa/ticker) - Terminal stock watcher and stock position tracker. Stars:`5.1K`.
- [go-money](https://github.com/rhymond/go-money) - Implementation of Fowler's Money pattern. Stars:`1.7K`.
- [bbgo](https://github.com/c9s/bbgo) - A crypto trading bot framework written in Go. Including common crypto exchange API, standard indicators, back-testing and many built-in strategies. Stars:`1.3K`.

**[⬆ back to top](#contents)**

## Forms

_Libraries for working with forms._

- [nosurf](https://github.com/justinas/nosurf) - CSRF protection middleware for Go. Stars:`1.6K`.
- [gorilla/csrf](https://github.com/gorilla/csrf) - CSRF protection for Go web applications & services. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Functional

_Packages to support functional programming in Go._

- [mo](https://github.com/samber/mo) - Monads and popular FP abstractions, based on Go 1.18+ Generics (Option, Result, Either...). Stars:`2.8K`.
- [go-underscore](https://github.com/tobyhede/go-underscore) - Useful collection of helpfully functional Go collection utilities. Stars:`1.3K`.

**[⬆ back to top](#contents)**

## Game Development

_Awesome game development libraries._

- [Ebitengine](https://github.com/hajimehoshi/ebiten) - dead simple 2D game engine in Go. Stars:`11.3K`.
- [Leaf](https://github.com/name5566/leaf) - Lightweight game server framework. Stars:`5.3K`.
- [nano](https://github.com/lonng/nano) - Lightweight, facility, high performance golang based game server framework. Stars:`2.9K`.
- [g3n](https://github.com/g3n/engine) - Go 3D Game Engine. Stars:`2.8K`.
- [goworld](https://github.com/xiaonanln/goworld) - Scalable game server engine, featuring space-entity framework and hot-swapping. Stars:`2.6K`.
- [Pitaya](https://github.com/topfreegames/pitaya) - Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. Stars:`2.4K`.
- [go-sdl2](https://github.com/veandco/go-sdl2) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). Stars:`2.2K`.
- [raylib-go](https://github.com/gen2brain/raylib-go) - Go bindings for [raylib](https://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. Stars:`1.8K`.
- [engo](https://github.com/EngoEngine/engo) - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. Stars:`1.8K`.
- [Oak](https://github.com/oakmound/oak) - Pure Go game engine. Stars:`1.6K`.
- [termloop](https://github.com/JoelOtter/termloop) - Terminal-based game engine for Go, built on top of Termbox. Stars:`1.4K`.
- [gonet](https://github.com/xtaci/gonet) - Game server skeleton implemented with golang. Stars:`1.3K`.

**[⬆ back to top](#contents)**

## Generators

_Tools that generate Go code._

- [oapi-codegen](https://github.com/deepmap/oapi-codegen) - This package contains a set of utilities for generating Go boilerplate code for services based on OpenAPI 3.0 API definitions. Stars:`6.5K`.
- [go-linq](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go. Stars:`3.5K`.
- [jennifer](https://github.com/dave/jennifer) - Generate arbitrary Go code without templates. Stars:`3.4K`.
- [goderive](https://github.com/awalterschulze/goderive) - Derives functions from input types Stars:`1.2K`.
- [GoWrap](https://github.com/hexdigest/gowrap) - Generate decorators for Go interfaces using simple templates. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Geographic

_Geographic tools and servers_

- [Tile38](https://github.com/tidwall/tile38) - Geolocation DB with spatial index and realtime geofencing. Stars:`9.2K`.
- [S2 geometry](https://github.com/golang/geo) - S2 geometry library in Go. Stars:`1.7K`.

**[⬆ back to top](#contents)**

## Go Compilers

_Tools for compiling Go to other languages._

- [gopherjs](https://github.com/gopherjs/gopherjs) - Compiler from Go to JavaScript. Stars:`12.8K`.

**[⬆ back to top](#contents)**

## Goroutines

_Tools for managing and working with Goroutines._

- [ants](https://github.com/panjf2000/ants) - A high-performance and low-cost goroutine pool in Go. Stars:`13.2K`.
- [conc](https://github.com/sourcegraph/conc) - `conc` is your toolbelt for structured concurrency in go, making common tasks easier and safer. Stars:`9.6K`.
- [tunny](https://github.com/Jeffail/tunny) - Goroutine pool for golang. Stars:`3.9K`.
- [goworker](https://github.com/benmanns/goworker) - goworker is a Go-based background worker. Stars:`2.8K`.
- [pond](https://github.com/alitto/pond) - Minimalistic and High-performance goroutine worker pool written in Go. Stars:`1.6K`.
- [rill](https://github.com/destel/rill) - Go toolkit for clean, composable, channel-based concurrency.  Stars:`1.5K`.
- [workerpool](https://github.com/gammazero/workerpool) - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. Stars:`1.4K`.

**[⬆ back to top](#contents)**

## GUI

_Libraries for building GUI Applications._

_Toolkits_

- [fyne](https://github.com/fyne-io/fyne) - Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android. Stars:`25.5K`.
- [webview](https://github.com/zserge/webview) - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). Stars:`12.8K`.
- [qt](https://github.com/therecipe/qt) - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). Stars:`10.5K`.
- [ui](https://github.com/andlabs/ui) - Platform-native GUI library for Go. Cross platform. Stars:`8.3K`.
- [app](https://github.com/murlokswarm/app) - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. Stars:`8.1K`.
- [walk](https://github.com/lxn/walk) - Windows application library kit for Go. Stars:`6.9K`.
- [DarwinKit](https://github.com/progrium/darwinkit) - Build native macOS applications using Go. Stars:`5.1K`.
- [go-sciter](https://github.com/sciter-sdk/go-sciter) - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. Stars:`2.6K`.
- [gotk3](https://github.com/gotk3/gotk3) - Go bindings for GTK3. Stars:`2.1K`.
- [Cogent Core](https://github.com/cogentcore/core) - A framework for building 2D and 3D apps that run on macOS, Windows, Linux, iOS, Android, and the web. Stars:`1.8K`.
- [Spot](https://github.com/roblillack/spot) - Reactive, cross-platform desktop GUI toolkit. Stars:`1.2K`.

_Interaction_

- [robotgo](https://github.com/go-vgo/robotgo) - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. Stars:`9.8K`.
- [systray](https://github.com/getlantern/systray) - Cross platform Go library to place an icon and menu in the notification area. Stars:`3.4K`.

**[⬆ back to top](#contents)**

## Hardware

_Libraries, tools, and tutorials for interacting with hardware._

- [arduino-cli](https://github.com/arduino/arduino-cli) - Official Arduino CLI and library. Can run standalone, or be incorporated into larger Go projects. Stars:`4.4K`.
- [go-rpio](https://github.com/stianeikeland/go-rpio) - GPIO for Go, doesn't require cgo. Stars:`2.2K`.
- [ghw](https://github.com/jaypipes/ghw) - Golang hardware discovery/inspection library. Stars:`1.7K`.
- [emgo](https://github.com/ziutek/emgo) - Go-like language for programming embedded systems (e.g. STM32 MCU). Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Images

_Libraries for manipulating images._

- [gocv](https://github.com/hybridgroup/gocv) - Go package for computer vision using OpenCV 3.3+. Stars:`6.8K`.
- [imaginary](https://github.com/h2non/imaginary) - Fast and simple HTTP microservice for image resizing. Stars:`5.7K`.
- [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package. Stars:`5.3K`.
- [gg](https://github.com/fogleman/gg) - 2D rendering in pure Go. Stars:`4.5K`.
- [bild](https://github.com/anthonynsimon/bild) - Collection of image processing algorithms in pure Go. Stars:`4.0K`.
- [imagor](https://github.com/cshum/imagor) - Fast, secure image processing server and Go library, using libvips. Stars:`3.5K`.
- [gowitness](https://github.com/sensepost/gowitness) - Screenshoting webpages using go and headless chrome on command line. Stars:`3.5K`.
- [ln](https://github.com/fogleman/ln) - 3D line art rendering in Go. Stars:`3.3K`.
- [bimg](https://github.com/h2non/bimg) - Small package for fast and efficient image processing using libvips. Stars:`2.8K`.
- [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation. Stars:`2.2K`.
- [picfit](https://github.com/thoas/picfit) - An image resizing server written in Go. Stars:`2.2K`.
- [pt](https://github.com/fogleman/pt) - Path tracing engine written in Go. Stars:`2.1K`.
- [smartcrop](https://github.com/muesli/smartcrop) - Finds good crops for arbitrary images and crop sizes. Stars:`1.8K`.
- [imagick](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API. Stars:`1.8K`.
- [gift](https://github.com/disintegration/gift) - Package of image processing filters. Stars:`1.8K`.
- [canvas](https://github.com/tdewolff/canvas) - Vector graphics to PDF, SVG or rasterized image. Stars:`1.5K`.
- [govips](https://github.com/davidbyttow/govips) - A lightning fast image processing and resizing library for Go. Stars:`1.3K`.
- [geopattern](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string. Stars:`1.3K`.
- [stegify](https://github.com/DimitarPetrov/stegify) - Go tool for LSB steganography, capable of hiding any file within an image. Stars:`1.2K`.

**[⬆ back to top](#contents)**

## IoT (Internet of Things)

_Libraries for programming devices of the IoT._

- [gobot](https://github.com/hybridgroup/gobot/) - Gobot is a framework for robotics, physical computing, and the Internet of Things. Stars:`9.0K`.
- [flogo](https://github.com/tibcosoftware/flogo) - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. Stars:`2.4K`.
- [ekuiper](https://github.com/lf-edge/ekuiper) - Lightweight data stream processing engine for IoT edge. Stars:`1.5K`.
- [shifu](https://github.com/Edgenesis/shifu) - Kubernetes native IoT development framework. Stars:`1.3K`.
- [gatt](https://github.com/paypal/gatt) - Gatt is a Go package for building Bluetooth Low Energy peripherals. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Job Scheduler

_Libraries for scheduling jobs._
- [gocron](https://github.com/go-co-op/gocron) - Easy and fluent Go job scheduling. This is an actively maintained fork of [jasonlvhit/gocron](https://github.com/jasonlvhit/gocron). Stars:`5.8K`.
- [go-quartz](https://github.com/reugn/go-quartz) - Simple, zero-dependency scheduling library for Go. Stars:`1.8K`.
- [JobRunner](https://github.com/bamzi/jobrunner) - Smart and featureful cron job scheduler with job queuing and live monitoring built in. Stars:`1.1K`.
- [gron](https://github.com/roylee0704/gron) - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. Stars:`1.0K`.

**[⬆ back to top](#contents)**

## JSON

_Libraries for working with JSON._

- [GJSON](https://github.com/tidwall/gjson) - Get a JSON value with one line of code. Stars:`14.6K`.
- [gabs](https://github.com/Jeffail/gabs) - For parsing, creating and editing unknown or dynamic JSON in Go. Stars:`3.5K`.
- [gojson](https://github.com/ChimeraCoder/gojson) - Automatically generate Go (golang) struct definitions from example JSON. Stars:`2.7K`.
- [fastjson](https://github.com/valyala/fastjson) - Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection. Stars:`2.3K`.

**[⬆ back to top](#contents)**

## Logging

_Libraries for generating and working with log files._

- [logrus](https://github.com/Sirupsen/logrus) - Structured logger for Go. Stars:`24.9K`.
- [zap](https://github.com/uber-go/zap) - Fast, structured, leveled logging in Go. Stars:`22.4K`.
- [zerolog](https://github.com/rs/zerolog) - Zero-allocation JSON logger. Stars:`10.8K`.
- [spew](https://github.com/davecgh/go-spew) - Implements a deep pretty printer for Go data structures to aid in debugging. Stars:`6.1K`.
- [lumberjack](https://github.com/natefinch/lumberjack) - Simple rolling logger, implements io.WriteCloser. Stars:`4.9K`.
- [glog](https://github.com/golang/glog) - Leveled execution logs for Go. Stars:`3.6K`.
- [tail](https://github.com/hpcloud/tail) - Go package striving to emulate the features of the BSD tail program. Stars:`2.7K`.
- [pp](https://github.com/k0kubun/pp) - Colored pretty printer for Go language. Stars:`1.9K`.
- [seelog](https://github.com/cihub/seelog) - Logging functionality with flexible dispatching, filtering, and formatting. Stars:`1.6K`.
- [log](https://github.com/apex/log) - Structured logging package for Go. Stars:`1.4K`.
- [log15](https://github.com/inconshreveable/log15) - Simple, powerful logging for Go. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Machine Learning

_Libraries for Machine Learning._

- [GoLearn](https://github.com/sjwhitworth/golearn) - General Machine Learning library for Go. Stars:`9.3K`.
- [gorse](https://github.com/zhenghaoz/gorse) - An offline recommender system backend based on collaborative filtering written in Go. Stars:`8.7K`.
- [gorgonia](https://github.com/gorgonia/gorgonia) - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. Stars:`5.6K`.
- [m2cgen](https://github.com/BayesWitnesses/m2cgen) - A CLI tool to transpile trained classic ML models into a native Go code with zero dependencies, written in Python with Go language support. Stars:`2.8K`.
- [gosseract](https://github.com/otiai10/gosseract) - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. Stars:`2.8K`.
- [tfgo](https://github.com/galeone/tfgo) - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. Stars:`2.4K`.
- [goml](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go. Stars:`1.6K`.

**[⬆ back to top](#contents)**

## Messaging

_Libraries that implement messaging systems._

- [Asynq](https://github.com/hibiken/asynq) - A simple, reliable, and efficient distributed task queue for Go built on top of Redis. Stars:`10.5K`.
- [Centrifugo](https://github.com/centrifugal/centrifugo) - Real-time messaging (Websockets or SockJS) server in Go. Stars:`8.6K`.
- [gorush](https://github.com/appleboy/gorush) - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). Stars:`8.1K`.
- [machinery](https://github.com/RichardKnop/machinery) - Asynchronous task queue/job queue based on distributed message passing. Stars:`7.6K`.
- [go-socket.io](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework. Stars:`5.8K`.
- [NATS Go Client](https://github.com/nats-io/nats.go) - Go client for the NATS Stars:`5.7K`.
- [Confluent Kafka Golang Client](https://github.com/confluentinc/confluent-kafka-go) - confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform. Stars:`4.7K`.
- [Mercure](https://github.com/dunglas/mercure) - Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). Stars:`4.7K`.
- [melody](https://github.com/olahol/melody) - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. Stars:`3.8K`.
- [APNs2](https://github.com/sideshow/apns2) - HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. Stars:`3.0K`.
- [go-nsq](https://github.com/nsqio/go-nsq) - the official Go package for NSQ. Stars:`2.6K`.
- [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) - gopush-cluster is a go push server cluster. Stars:`2.1K`.
- [EventBus](https://github.com/asaskevich/EventBus) - The lightweight event bus with async compatibility. Stars:`1.8K`.
- [amqp](https://github.com/rabbitmq/amqp091-go) - Go RabbitMQ Client Library. Stars:`1.6K`.
- [Beaver](https://github.com/Clivern/Beaver) - A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. Stars:`1.6K`.
- [Chanify](https://github.com/chanify/chanify) - A push notification server send message to your iOS devices. Stars:`1.3K`.
- [dbus](https://github.com/godbus/dbus) - Native Go bindings for D-Bus. Stars:`1.0K`.
  messaging system.
- [sarama](https://github.com/Shopify/sarama) - Go library for Apache Kafka. Stars:`11.7K`.
- [Watermill](https://github.com/ThreeDotsLabs/watermill) - Working efficiently with message streams. Building event driven applications, enabling event sourcing, RPC over messages, sagas. Can use conventional pub/sub implementations like Kafka or RabbitMQ, but also HTTP or MySQL binlog. Stars:`8.0K`.
- [Uniqush-Push](https://github.com/uniqush/uniqush-push) - Redis backed unified push service for server-side notifications to mobile devices. Stars:`1.5K`.
- [zmq4](https://github.com/pebbe/zmq4) - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). Stars:`1.2K`.

**[⬆ back to top](#contents)**

## Microsoft Office

- [unioffice](https://github.com/unidoc/unioffice) - Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. Stars:`4.4K`.

### Microsoft Excel

_Libraries for working with Microsoft Excel._

- [excelize](https://github.com/xuri/excelize) - Golang library for reading and writing Microsoft Excel&trade; (XLSX) files. Stars:`18.6K`.
- [xlsx](https://github.com/tealeg/xlsx) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. Stars:`5.9K`.

### Microsoft Word

_Libraries for working with Microsoft Word._


**[⬆ back to top](#contents)**

## Miscellaneous

### Dependency Injection

_Libraries for working with dependency injection._

- [google/wire](https://github.com/google/wire) - Automated Initialization in Go. Stars:`13.3K`.
- [fx](https://github.com/uber-go/fx) - A dependency injection based application framework for Go (built on top of dig). Stars:`6.1K`.
- [dig](https://github.com/uber-go/dig) - A reflection based dependency injection toolkit for Go. Stars:`4.0K`.
- [do](https://github.com/samber/do) - A dependency injection framework based on Generics. Stars:`1.9K`.

**[⬆ back to top](#contents)**

### Project Layout

_**Unofficial** set of patterns for structuring projects._

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout) - Set of common historical and emerging project layout patterns in the Go ecosystem. Note: despite the org-name they do not represent official golang standards, see [this issue](https://github.com/golang-standards/project-layout/issues/117) for more information. Nonetheless, some may find the layout useful. Stars:`50.4K`.
- [go-blueprint](https://github.com/Melkeydev/go-blueprint) - Allows users to spin up a quick Go project using a popular framework. Stars:`6.5K`.
- [ardanlabs/service](https://github.com/ardanlabs/service) - A [starter kit](https://github.com/ardanlabs/service/wiki) for building production grade scalable web service applications. Stars:`3.6K`.
- [goxygen](https://github.com/shpota/goxygen) - Generate a modern Web project with Go and Angular, React, or Vue in seconds. Stars:`3.6K`.
- [pagoda](https://github.com/mikestefanello/pagoda) - Rapid, easy full-stack web development starter kit built in Go. Stars:`2.3K`.
- [nunu](https://github.com/go-nunu/nunu) - Nunu is a scaffolding tool for building Go applications. Stars:`2.0K`.
- [modern-go-application](https://github.com/sagikazarmark/modern-go-application) - Go application boilerplate and example applying modern practices. Stars:`1.9K`.

**[⬆ back to top](#contents)**

### Strings

_Libraries for working with strings._

- [xstrings](https://github.com/huandu/xstrings) - Collection of useful string functions ported from other languages. Stars:`1.4K`.

**[⬆ back to top](#contents)**

### Uncategorized

_These libraries were placed here because none of the other categories seemed to fit._

- [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). Stars:`10.8K`.
- [gatus](https://github.com/TwinProduction/gatus) - Automated service health dashboard. Stars:`6.9K`.
- [gofakeit](https://github.com/brianvoe/gofakeit) - Random data generator written in go. Stars:`4.7K`.
- [archiver](https://github.com/mholt/archiver) - Library and command for making and extracting .zip and .tar.gz archives. Stars:`4.4K`.
- [go-resiliency](https://github.com/eapache/go-resiliency) - Resiliency patterns for golang. Stars:`2.2K`.
- [base64Captcha](https://github.com/mojocn/base64Captcha) - Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. Stars:`2.2K`.
- [gosms](https://github.com/haxpax/gosms) - Your own local SMS gateway in Go that can be used to send SMS. Stars:`1.5K`.
- [go-commons-pool](https://github.com/jolestar/go-commons-pool) - Generic object pool for Golang. Stars:`1.2K`.
- [llvm](https://github.com/llir/llvm) - Library for interacting with LLVM IR in pure Go. Stars:`1.2K`.
- [shoutrrr](https://github.com/containrrr/shoutrrr) - Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Natural Language Processing

_Libraries for working with human languages._

See also [Text Processing](#text-processing) and [Text Analysis](#text-analysis).

### Language Detection

- [lingua-go](https://github.com/pemistahl/lingua-go) - An accurate natural language detection library, suitable for long and short text alike. Supports detecting multiple languages in mixed-language text. Stars:`1.2K`.

### Morphological Analyzers

- [spaGO](https://github.com/nlpodyssey/spago) - Self-contained Machine Learning and Natural Language Processing library in Go. Stars:`1.8K`.

### Slugifiers

- [slug](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support. Stars:`1.2K`.

### Tokenizers

- [prose](https://github.com/jdkato/prose) - Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. Stars:`3.1K`.
- [gse](https://github.com/go-ego/gse) - Go efficient text segmentation; support english, chinese, japanese and other. Stars:`2.6K`.
- [gojieba](https://github.com/yanyiwu/gojieba) - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. Stars:`2.5K`.

### Translation

- [go-i18n](https://github.com/nicksnyder/go-i18n/) - Package and an accompanying tool to work with localized text. Stars:`3.1K`.
- [go-pinyin](https://github.com/mozillazg/go-pinyin) - CN Hanzi to Hanyu Pinyin converter. Stars:`1.6K`.

### Transliteration


**[⬆ back to top](#contents)**

## Networking

_Libraries for working with various layers of the network._

- [fasthttp](https://github.com/valyala/fasthttp) - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. Stars:`22.1K`.
- [webrtc](https://github.com/pions/webrtc) - A pure Go implementation of the WebRTC API. Stars:`14.1K`.
- [kcptun](https://github.com/xtaci/kcptun) - Extremely simple & fast udp tunnel based on KCP protocol. Stars:`14.0K`.
- [quic-go](https://github.com/lucas-clemente/quic-go) - An implementation of the QUIC protocol in pure Go. Stars:`10.3K`.
- [gnet](https://github.com/panjf2000/gnet) - `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. Stars:`9.9K`.
- [cloudflared](https://github.com/cloudflare/cloudflared) - Cloudflare Tunnel client (formerly Argo Tunnel).  Stars:`9.6K`.
- [dns](https://github.com/miekg/dns) - Go library for working with DNS. Stars:`8.1K`.
- [gopacket](https://github.com/google/gopacket) - Go library for packet processing with libpcap bindings. Stars:`6.4K`.
- [kcp-go](https://github.com/xtaci/kcp-go) - KCP - Fast and Reliable ARQ Protocol. Stars:`4.1K`.
- [netpoll](https://github.com/cloudwego/netpoll) - A high-performance non-blocking I/O networking framework, which focused on RPC scenarios, developed by ByteDance. Stars:`4.1K`.
- [HTTPLab](https://github.com/gchaincl/httplab) - HTTPLabs let you inspect HTTP requests and forge responses. Stars:`4.1K`.
- [ssh](https://github.com/gliderlabs/ssh) - Higher-level API for building SSH servers (wraps crypto/ssh). Stars:`3.8K`.
- [gobgp](https://github.com/osrg/gobgp) - BGP implemented in the Go Programming Language. Stars:`3.7K`.
- [tun2socks](https://github.com/xjasonlyu/tun2socks) - A pure go implementation of tun2socks powered by [gVisor](https://gvisor.dev/) TCP/IP stack. Stars:`3.6K`.
- [fortio](https://github.com/fortio/fortio) - Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. Stars:`3.4K`.
- [nbio](https://github.com/lesismal/nbio) - Pure Go 1000k+ connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use. Stars:`2.3K`.
- [water](https://github.com/songgao/water) - Simple TUN/TAP library. Stars:`2.0K`.
- [gev](https://github.com/Allenxuxu/gev) - gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. Stars:`1.7K`.
- [go-getter](https://github.com/hashicorp/go-getter) - Go library for downloading files or directories from various sources using a URL. Stars:`1.7K`.
- [sftp](https://github.com/pkg/sftp) - Package sftp implements the SSH File Transfer Protocol as described in <https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt>. Stars:`1.5K`.
- [gws](https://github.com/lxzan/gws) - High-Performance WebSocket Server & Client With AsyncIO Supporting . Stars:`1.5K`.
- [grab](https://github.com/cavaliercoder/grab) - Go package for managing file downloads. Stars:`1.4K`.
- [NFF-Go](https://github.com/intel-go/nff-go) - Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). Stars:`1.4K`.
- [ftp](https://github.com/jlaffaye/ftp) - Package ftp implements a FTP client as described in [RFC 959](https://tools.ietf.org/html/rfc959). Stars:`1.3K`.
- [mdns](https://github.com/hashicorp/mdns) - Simple mDNS (Multicast DNS) client/server library in Golang. Stars:`1.2K`.
- [gosnmp](https://github.com/soniah/gosnmp) - Native Go library for performing SNMP actions. Stars:`1.2K`.

**[⬆ back to top](#contents)**

### HTTP Clients

_Libraries for making HTTP requests._

- [resty](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client. Stars:`10.4K`.
- [req](https://github.com/imroc/req) - Simple Go HTTP client with Black Magic (Less code and More efficiency). Stars:`4.4K`.
- [heimdall](https://github.com/gojektech/heimdall) - An enhanced http client with retry and hystrix capabilities. Stars:`2.6K`.
- [grequests](https://github.com/levigross/grequests) - A Go "clone" of the great and famous Requests library. Stars:`2.1K`.
- [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) - Retryable HTTP client in Go. Stars:`2.0K`.
- [sling](https://github.com/dghubble/sling) - Sling is a Go HTTP client library for creating and sending API requests. Stars:`1.7K`.
- [requests](https://github.com/carlmjohnson/requests) - HTTP requests for Gophers. Uses context.Context and doesn't hide the underlying net/http.Client, making it compatible with standard Go APIs. Also includes testing tools. Stars:`1.5K`.
- [gentleman](https://github.com/h2non/gentleman) - Full-featured plugin-driven HTTP client library. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## OpenGL

_Libraries for using OpenGL in Go._

- [glfw](https://github.com/go-gl/glfw) - Go bindings for GLFW 3. Stars:`1.6K`.
- [gl](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow). Stars:`1.1K`.

**[⬆ back to top](#contents)**

## ORM

_Libraries that implement Object-Relational Mapping or datamapping techniques._

- [GORM](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang, aims to be developer friendly. Stars:`37.3K`.
- [ent](https://github.com/facebook/ent) - An entity framework for Go. Simple, yet powerful ORM for modeling and querying data. Stars:`15.8K`.
- [SQLBoiler](https://github.com/volatiletech/sqlboiler) - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. Stars:`6.8K`.
- [bun](https://github.com/uptrace/bun) - SQL-first Golang ORM. Successor of go-pg. Stars:`3.9K`.
- [gorp](https://github.com/go-gorp/gorp) - Go Relational Persistence, ORM-ish library for Go. Stars:`3.7K`.
- [upper.io/db](https://github.com/upper/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. Stars:`3.6K`.
- [gormt](https://github.com/xxjwxc/gormt) - Mysql database to golang gorm struct. Stars:`2.4K`.
- [Prisma](https://github.com/prisma/prisma-client-go) - Prisma Client Go, Typesafe database access for Go. Stars:`2.2K`.
- [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) - A flexible and powerful SQL string builder library plus a zero-config ORM. Stars:`1.5K`.
- [pop/soda](https://github.com/gobuffalo/pop) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. Stars:`1.5K`.
- [reform](https://github.com/go-reform/reform) - Better ORM for Go, based on non-empty interfaces and code generation. Stars:`1.4K`.

**[⬆ back to top](#contents)**

## Package Management

_Official tooling for dependency and package management_


_Unofficial libraries for package and dependency management._

- [glide](https://github.com/Masterminds/glide) - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. Stars:`8.1K`.
- [syft](https://github.com/anchore/syft) - A CLI tool and Go library for generating a Software Bill of Materials (SBOM) from container images and filesystems. Stars:`6.5K`.
- [godep](https://github.com/tools/godep) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. Stars:`5.5K`.
- [govendor](https://github.com/kardianos/govendor) - Go Package Manager. Go vendor tool that works with the standard vendor file. Stars:`4.9K`.
- [gopm](https://github.com/gpmgo/gopm) - Go Package Manager. Stars:`2.5K`.
- [gpm](https://github.com/pote/gpm) - Barebones dependency manager for Go. Stars:`1.2K`.

**[⬆ back to top](#contents)**

## Performance

- [jaeger](https://github.com/jaegertracing/jaeger) - A distributed tracing system. Stars:`20.8K`.
- [pixie](https://github.com/pixie-labs/pixie) - No instrumentation tracing for Golang applications via eBPF. Stars:`5.7K`.
- [statsviz](https://github.com/arl/statsviz) - Live visualization of your Go application runtime statistics. Stars:`3.2K`.
- [profile](https://github.com/pkg/profile) - Simple profiling support package for Go. Stars:`2.0K`.

**[⬆ back to top](#contents)**

## Query Language

- [gqlgen](https://github.com/99designs/gqlgen) - go generate based graphql server library. Stars:`10.1K`.
- [graphql-go](https://github.com/graphql-go/graphql) - Implementation of GraphQL for Go. Stars:`10.0K`.
- [dasel](https://github.com/tomwright/dasel) - Query and update data structures using selectors from the command line. Comparable to jq/yq but supports JSON, YAML, TOML and XML with zero runtime dependencies. Stars:`7.3K`.
- [graphql](https://github.com/neelance/graphql-go) - GraphQL server with a focus on ease of use. Stars:`4.7K`.
- [gojsonq](https://github.com/thedevsaddam/gojsonq) - A simple Go package to Query over JSON Data. Stars:`2.2K`.

**[⬆ back to top](#contents)**

## Reflection


**[⬆ back to top](#contents)**

## Resource Embedding

- [statik](https://github.com/rakyll/statik) - Embeds static files into a Go executable. Stars:`3.8K`.
- [packr](https://github.com/gobuffalo/packr) - The simple and easy way to embed static files into Go binaries. Stars:`3.4K`.
- [go.rice](https://github.com/GeertJohan/go.rice) - go.rice is a Go package that makes working with resources such as HTML, JS, CSS, images, and templates very easy. Stars:`2.4K`.

**[⬆ back to top](#contents)**

## Science and Data Analysis

_Libraries for scientific computing and data analyzing._

- [gonum](https://github.com/gonum/gonum) - Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. Stars:`7.7K`.
- [stats](https://github.com/montanaflynn/stats) - Statistics package with common functions missing from the Golang standard library. Stars:`2.9K`.
- [gonum/plot](https://github.com/gonum/plot) - gonum/plot provides an API for building and drawing plots in Go. Stars:`2.8K`.
- [gosl](https://github.com/cpmech/gosl) - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. Stars:`1.8K`.
- [streamtools](https://github.com/nytlabs/streamtools) - general purpose, graphical tool for dealing with streams of data. Stars:`1.3K`.
- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Dataframes for machine-learning and statistics (similar to pandas). Stars:`1.2K`.

**[⬆ back to top](#contents)**

## Security

_Libraries that are used to help make your application more secure._

- [age](https://github.com/FiloSottile/age) - A simple, modern and secure encryption tool (and Go library) with small explicit keys, no config options, and UNIX-style composability. Stars:`17.9K`.
- [lego](https://github.com/go-acme/lego) - Pure Go ACME client library and CLI tool (for use with Let's Encrypt). Stars:`8.2K`.
- [CertMagic](https://github.com/caddyserver/certmagic) - Mature, robust, and powerful ACME client integration for fully-managed TLS certificate issuance and renewal. Stars:`5.1K`.
- [Cameradar](https://github.com/Ullaakut/cameradar) - Tool and library to remotely hack RTSP streams from surveillance cameras. Stars:`4.2K`.
- [memguard](https://github.com/awnumar/memguard) - A pure Go library for handling sensitive values in memory. Stars:`2.6K`.
- [Coraza](https://github.com/corazawaf/coraza) - Enterprise-ready, modsecurity and OWASP CRS compatible WAF library. Stars:`2.4K`.
- [secure](https://github.com/unrolled/secure) - HTTP middleware for Go that facilitates some quick security wins. Stars:`2.3K`.
- [acmetool](https://github.com/hlandau/acme) - ACME (Let's Encrypt) client tool with automatic renewal. Stars:`2.1K`.
- [themis](https://github.com/cossacklabs/themis) - high-level cryptographic library for solving typical data security tasks (secure data storage, secure messaging, zero-knowledge proof authentication), available for 14 languages, best fit for multi-platform apps. Stars:`1.9K`.
- [acra](https://github.com/cossacklabs/acra) - Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. Stars:`1.4K`.

**[⬆ back to top](#contents)**

## Serialization

_Libraries and tools for binary serialization._

- [jsoniter](https://github.com/json-iterator/go) - High-performance 100% compatible drop-in replacement of "encoding/json". Stars:`13.5K`.
- [goprotobuf](https://github.com/golang/protobuf) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. Stars:`9.8K`.
- [mapstructure](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures. Stars:`7.9K`.
- [gogoprotobuf](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets. Stars:`5.7K`.
- [go-codec](https://github.com/ugorji/go) - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. Stars:`1.9K`.

**[⬆ back to top](#contents)**

## Server Applications

- [Caddy](https://github.com/caddyserver/caddy) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use. Stars:`60.4K`.
- [minio](https://github.com/minio/minio) - Minio is a distributed object storage server. Stars:`49.4K`.
- [etcd](https://github.com/etcd-io/etcd) - Highly-available key value store for shared configuration and service discovery. Stars:`48.2K`.
- [pocketbase](https://github.com/pocketbase/pocketbase) - PocketBase is a realtime backend in 1 file consisting of embedded database (SQLite) with realtime subscriptions, built-in auth management and much more. Stars:`42.8K`.
- [SFTPGo](https://github.com/drakkan/sftpgo) - Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support. It can serve local filesystem and Cloud Storage backends such as S3 and Google Cloud Storage. Stars:`9.8K`.
- [RoadRunner](https://github.com/spiral/roadrunner) - High-performance PHP application server, load-balancer and process manager. Stars:`8.0K`.
- [Easegress](https://github.com/megaease/easegress) - A cloud native high availability/performance traffic orchestration system with observability and extensibility. Stars:`5.8K`.
- [flipt](https://github.com/markphelps/flipt) - A self contained feature flag solution written in Go and Vue.js Stars:`4.0K`.
- [Wish](https://github.com/charmbracelet/wish) - Make SSH apps, just like that! Stars:`3.8K`.
- [devd](https://github.com/cortesi/devd) - Local webserver for developers. Stars:`3.4K`.
- [Fider](https://github.com/getfider/fider) - Fider is an open platform to collect and organize customer feedback. Stars:`3.1K`.
- [algernon](https://github.com/xyproto/algernon) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. Stars:`2.9K`.
- [Flagr](https://github.com/checkr/flagr) - Flagr is an open-source feature flagging and A/B testing service. Stars:`2.5K`.
- [Trickster](https://github.com/tricksterproxy/trickster) - HTTP reverse proxy cache and time series accelerator. Stars:`2.0K`.
- [discovery](https://github.com/Bilibili/discovery) - A registry for resilient mid-tier load balancing and failover. Stars:`1.8K`.
- [go-feature-flag](https://github.com/thomaspoignant/go-feature-flag) - A simple, complete and lightweight self-hosted feature flag solution 100% Open Source. Stars:`1.5K`.
- [jackal](https://github.com/ortuman/jackal) - An XMPP server written in Go. Stars:`1.4K`.

**[⬆ back to top](#contents)**

## Stream Processing

_Libraries and tools for stream processing and reactive programming._

- [go-streams](https://github.com/reugn/go-streams) - Go stream processing library. Stars:`2.0K`.

**[⬆ back to top](#contents)**

## Template Engines

_Libraries and tools for templating and lexing._

- [templ](https://github.com/a-h/templ) - A HTML templating language that has great developer tooling. Stars:`8.7K`.
- [quicktemplate](https://github.com/valyala/quicktemplate) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. Stars:`3.2K`.
- [pongo2](https://github.com/flosch/pongo2) - Django-like template-engine for Go. Stars:`2.9K`.
- [maroto](https://github.com/johnfercher/maroto) - A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. Stars:`2.3K`.
- [jet](https://github.com/CloudyKit/jet) - Jet template engine. Stars:`1.3K`.

**[⬆ back to top](#contents)**

## Testing

_Libraries for testing codebases and generating test data._

### Testing Frameworks

- [Testify](https://github.com/stretchr/testify) - Sacred extension to the standard go testing package. Stars:`23.8K`.
- [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload. Stars:`8.3K`.
- [keploy](https://github.com/keploy/keploy) - Generate Testcase and Data Mocks from API calls automatically. Stars:`5.2K`.
- [go-cmp](https://github.com/google/go-cmp) - Package for comparing Go values in tests. Stars:`4.3K`.
- [testcontainers-go](https://github.com/testcontainers/testcontainers-go) - A Go package that makes it simple to create and clean up container-based dependencies for automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers that should be run as part of a test and clean up those resources when the test is done. Stars:`3.8K`.
- [httpexpect](https://github.com/gavv/httpexpect) - Concise, declarative, and easy to use end-to-end HTTP and REST API testing. Stars:`2.6K`.
- [godog](https://github.com/cucumber/godog) - Cucumber BDD framework for Go. Stars:`2.4K`.
- [is](https://github.com/matryer/is) - Professional lightweight testing mini-framework for Go. Stars:`1.8K`.
- [gnomock](https://github.com/orlangure/gnomock) - integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks. Stars:`1.4K`.
- [go-vcr](https://github.com/dnaeon/go-vcr) - Record and replay your HTTP interactions for fast, deterministic and accurate tests. Stars:`1.3K`.
- [testfixtures](https://github.com/go-testfixtures/testfixtures) - A helper for Rails' like test fixtures to test database applications. Stars:`1.1K`.

### Mock

- [gomock](https://github.com/golang/mock) - Mocking framework for the Go programming language. Stars:`9.3K`.
- [mockery](https://github.com/vektra/mockery) - Tool to generate Go interfaces. Stars:`6.3K`.
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database interactions. Stars:`6.2K`.
- [hoverfly](https://github.com/SpectoLabs/hoverfly) - HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. Stars:`2.4K`.
- [gock](https://github.com/h2non/gock) - Versatile HTTP mocking made easy. Stars:`2.1K`.
- [moq](https://github.com/matryer/moq) - Utility that generates a struct from any interface. The struct can be used in test code as a mock of the interface. Stars:`2.0K`.
- [httpmock](https://github.com/jarcoal/httpmock) - Easy mocking of HTTP responses from external resources. Stars:`2.0K`.
- [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects. Stars:`1.0K`.

### Fuzzing and delta-debugging/reducing/shrinking

- [go-fuzz](https://github.com/dvyukov/go-fuzz) - Randomized testing system. Stars:`4.8K`.
- [gofuzz](https://github.com/google/gofuzz) - Library for populating go objects with random values. Stars:`1.5K`.

### Selenium and browser control tools

- [chromedp](https://github.com/knq/chromedp) - a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. Stars:`11.3K`.
- [rod](https://github.com/go-rod/rod) - A Devtools driver to make web automation and scraping easy. Stars:`5.6K`.
- [selenoid](https://github.com/aerokube/selenoid) - alternative Selenium hub server that launches browsers within containers. Stars:`2.6K`.
- [playwright-go](https://github.com/mxschmitt/playwright-go) - browser automation library to control Chromium, Firefox and WebKit with a single API. Stars:`2.3K`.

### Fail injection

**[⬆ back to top](#contents)**

## Text Processing

_Libraries for parsing and manipulating texts._

See also [Natural Language Processing](#natural-language-processing) and [Text Analysis](#text-analysis).

### Formatters

- [go-humanize](https://github.com/dustin/go-humanize) - Formatters for time, numbers, and memory size to human readable format. Stars:`4.4K`.
- [sq](https://github.com/neilotoole/sq) - Convert data from SQL databases or document formats like CSV or Excel into formats such as JSON, Excel, CSV, HTML, Markdown, XML, and YAML. Stars:`2.2K`.

### Markup Languages

- [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go. Stars:`5.5K`.
- [toml](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection). Stars:`4.6K`.
- [goldmark](https://github.com/yuin/goldmark) - A Markdown parser written in Go. Easy to extend, standard (CommonMark) compliant, well structured. Stars:`3.8K`.
- [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) - Convert HTML to Markdown. Even works with entire websites and can be extended through rules. Stars:`2.5K`.
- [go-toml](https://github.com/pelletier/go-toml) - Go library for the TOML format with query support and handy cli tools. Stars:`1.8K`.

### Parsers/Encoders/Decoders

- [sh](https://github.com/mvdan/sh) - Shell parser and formatter. Stars:`7.4K`.
- [gofeed](https://github.com/mmcdole/gofeed) - Parse RSS and Atom feeds in Go. Stars:`2.6K`.
- [go-querystring](https://github.com/google/go-querystring) - Go library for encoding structs into URL query parameters. Stars:`2.0K`.
- [when](https://github.com/olebedev/when) - Natural EN and RU language date/time parser with pluggable rules. Stars:`1.4K`.

### Regular Expressions


### Sanitation

- [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer. Stars:`3.3K`.

### Scrapers

- [colly](https://github.com/asciimoo/colly) - Fast and Elegant Scraping Framework for Gophers. Stars:`23.6K`.
- [GoQuery](https://github.com/PuerkitoBio/goquery) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language. Stars:`14.2K`.
- [xurls](https://github.com/mvdan/xurls) - Extract urls from text. Stars:`1.2K`.

### RSS


### Utility/Miscellaneous

- [w2vgrep](https://github.com/arunsupe/semantic-grep) - A semantic grep tool using word embeddings to find semantically similar matches. For example, searching for "death" will find "dead", "killing", "murder". Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Third-party APIs

_Libraries for accessing third party APIs._

- [github](https://github.com/google/go-github) - Go library for accessing the GitHub REST API v3. Stars:`10.6K`.
- [go-openai](https://github.com/sashabaranov/go-openai) - OpenAI ChatGPT, DALL·E, Whisper API library for Go. Stars:`9.4K`.
- [discordgo](https://github.com/bwmarrin/discordgo) - Go bindings for the Discord Chat API. Stars:`5.2K`.
- [slack](https://github.com/slack-go/slack) - Slack API in Go. Stars:`4.7K`.
- [google](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go. Stars:`4.1K`.
- [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library. Stars:`3.8K`.
- [aws-sdk-go](https://github.com/aws/aws-sdk-go-v2) - The official AWS SDK for the Go programming language. Stars:`2.8K`.
- [minio-go](https://github.com/minio/minio-go) - Minio Go Library for Amazon S3 compatible cloud storage. Stars:`2.5K`.
- [stripe](https://github.com/stripe/stripe-go) - Go client for the Stripe API. Stars:`2.2K`.
- [go-twitter](https://github.com/dghubble/go-twitter) - Go client library for the Twitter v1.1 APIs. Stars:`1.6K`.
- [go-jira](https://github.com/andygrunwald/go-jira) - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira) Stars:`1.5K`.
- [facebook](https://github.com/huandu/facebook) - Go Library that supports the Facebook Graph API. Stars:`1.3K`.
- [anaconda](https://github.com/ChimeraCoder/anaconda) - Go client library for the Twitter 1.1 API. Stars:`1.1K`.
- [githubql](https://github.com/shurcooL/githubql) - Go library for accessing the GitHub GraphQL API v4. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Utilities

_General utilities and tools to make your life easier._

- [fzf](https://github.com/junegunn/fzf) - Command-line fuzzy finder written in Go. Stars:`67.1K`.
- [dive](https://github.com/wagoodman/dive) - A tool for exploring each layer in a Docker image. Stars:`49.0K`.
- [hub](https://github.com/github/hub) - wrap git commands with additional functionality to interact with github from the terminal. Stars:`22.9K`.
- [lo](https://github.com/samber/lo) - A Lodash like Go library based on Go 1.18+ Generics (map, filter, contains, find...) Stars:`18.4K`.
- [sqlx](https://github.com/jmoiron/sqlx) - provides a set of extensions on top of the excellent built-in database/sql package. Stars:`16.5K`.
- [ctop](https://github.com/bcicen/ctop) - [Top-like](https://ctop.sh) interface (e.g. htop) for container metrics. Stars:`15.7K`.
- [goreleaser](https://github.com/goreleaser/goreleaser) - Deliver Go binaries as fast and easily as possible. Stars:`14.1K`.
- [wuzz](https://github.com/asciimoo/wuzz) - Interactive cli tool for HTTP inspection. Stars:`10.6K`.
- [usql](https://github.com/knq/usql) - usql is a universal command-line interface for SQL databases. Stars:`9.2K`.
- [peco](https://github.com/peco/peco) - Simplistic interactive filtering tool. Stars:`7.7K`.
- [lancet](https://github.com/duke-git/lancet) - A comprehensive, efficient, and reusable util function library of go. Stars:`4.9K`.
- [go-funk](https://github.com/thoas/go-funk) - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). Stars:`4.8K`.
- [godropbox](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications from Dropbox. Stars:`4.2K`.
- [minify](https://github.com/tdewolff/minify) - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. Stars:`3.8K`.
- [panicparse](https://github.com/maruel/panicparse) - Groups similar goroutines and colorizes stack dump. Stars:`3.5K`.
- [goreporter](https://github.com/wgliang/goreporter) - Golang tool that does static analysis, unit testing, code review and generate code quality report. Stars:`3.1K`.
- [mc](https://github.com/minio/mc) - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. Stars:`2.9K`.
- [mergo](https://github.com/imdario/mergo) - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. Stars:`2.9K`.
- [create-go-app](https://github.com/create-go-app/cli) - A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command. Stars:`2.6K`.
- [retry-go](https://github.com/avast/retry-go) - Simple library for retry mechanism. Stars:`2.5K`.
- [EaseProbe](https://github.com/megaease/easeprobe) - A simple, standalone, and lightWeight tool that can do health/status checking daemon, support HTTP/TCP/SSH/Shell/Client/... probes, and Slack/Discord/Telegram/SMS... notification. Stars:`2.2K`.
- [filetype](https://github.com/h2non/filetype) - Small package to infer the file type checking the magic numbers signature. Stars:`2.1K`.
- [Storm](https://github.com/asdine/storm) - Simple and powerful toolkit for BoltDB. Stars:`2.1K`.
- [jump](https://github.com/gsamokovarov/jump) - Jump helps you navigate faster by learning your habits. Stars:`1.8K`.
- [boilr](https://github.com/tmrts/boilr) - Blazingly fast CLI tool for creating projects from boilerplate templates. Stars:`1.7K`.
- [mimetype](https://github.com/gabriel-vasile/mimetype) - Package for MIME type detection based on magic numbers. Stars:`1.7K`.
- [Failsafe-go](https://github.com/failsafe-go/failsafe-go) - Fault tolerance and resilience patterns for Go. Stars:`1.7K`.
- [mole](https://github.com/davrodpin/mole) - cli app to easily create ssh tunnels. Stars:`1.7K`.
- [gitbatch](https://github.com/isacikgoz/gitbatch) - manage your git repositories in one place. Stars:`1.5K`.
- [scany](https://github.com/georgysavva/scany) - Library for scanning data from a database into Go structs and more. Stars:`1.3K`.
- [bed](https://github.com/itchyny/bed) - A Vim-like binary editor written in Go. Stars:`1.3K`.
- [circuitbreaker](https://github.com/rubyist/circuitbreaker) - Circuit Breakers in Go. Stars:`1.1K`.
- [hostctl](https://github.com/guumaster/hostctl) - A CLI tool to manage /etc/hosts with easy commands. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## UUID

_Libraries for working with UUIDs._

- [uuid](https://github.com/google/uuid) - Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. Stars:`5.4K`.
- [ulid](https://github.com/oklog/ulid) - Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). Stars:`4.6K`.
- [xid](https://github.com/rs/xid) - Xid is a globally unique id generator library, ready to be safely used directly in your server code. Stars:`4.0K`.
- [uuid](https://github.com/gofrs/uuid) - Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. Stars:`1.6K`.

**[⬆ back to top](#contents)**

## Validation

_Libraries for validation._

- [validator](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. Stars:`17.3K`.
- [govalidator](https://github.com/asaskevich/govalidator) - Validators and sanitizers for strings, numerics, slices and structs. Stars:`6.1K`.
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. Stars:`3.8K`.
- [govalidator](https://github.com/thedevsaddam/govalidator) - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. Stars:`1.3K`.
- [validate](https://github.com/gookit/validate) - Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. Stars:`1.1K`.
**[⬆ back to top](#contents)**

## Version Control

_Libraries for version control._

- [go-git](https://github.com/go-git/go-git) - highly extensible Git implementation in pure Go. Stars:`6.2K`.
- [hercules](https://github.com/src-d/hercules) - gaining advanced insights from Git repository history. Stars:`2.7K`.
- [git2go](https://github.com/libgit2/git2go) - Go bindings for libgit2. Stars:`1.9K`.

**[⬆ back to top](#contents)**

## Video

_Libraries for manipulating video._

- [goav](https://github.com/giorgisio/goav) - Comprehensive Go bindings for FFmpeg. Stars:`2.1K`.
- [m3u8](https://github.com/grafov/m3u8) - Parser and generator library of M3U8 playlists for Apple HLS. Stars:`1.2K`.

**[⬆ back to top](#contents)**

## Web Frameworks

_Full stack web frameworks._

- [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. Stars:`79.8K`.
- [Fiber](https://github.com/gofiber/fiber) - An Express.js inspired web framework build on Fasthttp. Stars:`34.6K`.
- [Beego](https://github.com/beego/beego) - beego is an open-source, high-performance web framework for the Go programming language. Stars:`31.8K`.
- [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework. Stars:`30.3K`.
- [Revel](https://github.com/revel/revel) - High-productivity web framework for the Go language. Stars:`13.2K`.
- [GoFrame](https://github.com/gogf/gf) - GoFrame is a modular, powerful, high-performance and enterprise-class application development framework of Golang. Stars:`11.9K`.
- [Hertz](https://github.com/cloudwego/hertz) - A high-performance and strong-extensibility Go HTTP framework that helps developers build microservices. Stars:`5.7K`.
- [Goa](https://github.com/goadesign/goa) - Goa provides a holistic approach for developing remote APIs and microservices in Go. Stars:`5.7K`.
- [GoFr](https://github.com/gofr-dev/gofr) - Gofr is an opinionated microservice development framework. Stars:`3.9K`.
- [goravel](https://github.com/goravel/goravel) - A Laravel-inspired web framework with ORM, authentication, queue, task scheduling, and more built-in features. Stars:`3.1K`.
- [Huma](https://github.com/danielgtaylor/huma/) - Framework for modern REST/GraphQL APIs with built-in OpenAPI 3, generated documentation, and a CLI. Stars:`2.5K`.
- [Goyave](https://github.com/go-goyave/goyave) - Feature-complete REST API framework aimed at clean code and fast development, with powerful built-in functionalities. Stars:`1.7K`.
- [Atreugo](https://github.com/savsgio/atreugo) - High performance and extensible micro web framework with zero memory allocations in hot paths. Stars:`1.3K`.
- [Fuego](https://github.com/go-fuego/fuego) - The framework for busy Go developers! Web framework generating OpenAPI 3 spec from source code. Stars:`1.1K`.

**[⬆ back to top](#contents)**

### Middlewares

#### Actual middlewares

- [Tollbooth](https://github.com/didip/tollbooth) - Rate limit HTTP request handler. Stars:`2.7K`.
- [CORS](https://github.com/rs/cors) - Easily add CORS capabilities to your API. Stars:`2.7K`.
- [Limiter](https://github.com/ulule/limiter) - Dead simple rate limit middleware for Go. Stars:`2.1K`.

#### Libraries for creating HTTP middlewares

- [negroni](https://github.com/urfave/negroni) - Idiomatic HTTP middleware for Golang. Stars:`7.5K`.
- [alice](https://github.com/justinas/alice) - Painless middleware chaining for Go. Stars:`3.2K`.
- [render](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, and HTML template responses. Stars:`2.0K`.

**[⬆ back to top](#contents)**

### Routers

- [mux](https://github.com/gorilla/mux) - Powerful URL router and dispatcher for golang. Stars:`21.0K`.
- [chi](https://github.com/go-chi/chi) - Small, fast and expressive HTTP router built on net/context. Stars:`18.9K`.
- [httprouter](https://github.com/julienschmidt/httprouter) - High performance router. Use this and the standard http handlers to form a very high performance web framework. Stars:`16.7K`.
- [gocraft/web](https://github.com/gocraft/web) - Mux and middleware package in Go. Stars:`1.5K`.
- [Bone](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer. Stars:`1.3K`.

**[⬆ back to top](#contents)**

## WebAssembly

- [tinygo](https://github.com/tinygo-org/tinygo) - Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. Stars:`15.7K`.

**[⬆ back to top](#contents)**

## Webhooks Server

- [webhook](https://github.com/adnanh/webhook) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. Stars:`10.6K`.

**[⬆ back to top](#contents)**

## Windows

- [go-ole](https://github.com/go-ole/go-ole) - Win32 OLE implementation for golang. Stars:`1.2K`.

**[⬆ back to top](#contents)**

## Workflow Frameworks

_Libraries for creating Workflows._
- [Dagu](https://github.com/dagu-go/dagu) - No-code workflow executor. it executes DAGs defined in a simple YAML format. Stars:`1.8K`.

**[⬆ back to top](#contents)**

## XML

_Libraries and tools for manipulating XML._


## Zero Trust

_Libraries and tools to implement Zero Trust architectures._

- [Cosign](https://github.com/sigstore/cosign) - Container Signing, Verification and Storage in an OCI registry. Stars:`4.6K`.
- [OpenZiti](https://github.com/openziti/ziti) - A full, open source zero trust overlay network. Including numerous SDKs for numerous languages such as [golang](https://github.com/openziti/sdk-golang) allowing you to embed zero trust principles directly into your applications. The [OpenZiti Test Kitchen](https://github.com/openziti-test-kitchen) has numerous examples to draw inspiration from including a [zero trust ssh client - zssh](https://github.com/openziti-test-kitchen/zssh) Stars:`3.0K`.
- [Spire](https://github.com/spiffe/spire) - SPIRE (the SPIFFE Runtime Environment) is a toolchain of APIs for establishing trust between software systems across a wide variety of hosting platforms. Stars:`1.8K`.

## Code Analysis

_Source code analysis tools, also known as Static Application Security Testing (SAST) Tools._

- [golangci-lint](https://github.com/golangci/golangci-lint) – A fast Go linters runner. It runs linters in parallel, uses caching, supports `yaml` config, has integrations with all major IDE and has dozens of linters included. Stars:`16.0K`.
- [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#. Stars:`6.3K`.
- [revive](https://github.com/mgechev/revive) – ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for `golint`. Stars:`4.9K`.
- [errcheck](https://github.com/kisielk/errcheck) - Errcheck is a program for checking for unchecked errors in Go programs. Stars:`2.4K`.
- [GoPlantUML](https://github.com/jfeliu007/goplantuml) - Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. Stars:`1.9K`.
- [go-critic](https://github.com/go-critic/go-critic) - source code linter that brings checks that are currently not implemented in other linters. Stars:`1.9K`.
- [gcvis](https://github.com/davecheney/gcvis) - Visualise Go program GC trace data in real time. Stars:`1.1K`.

**[⬆ back to top](#contents)**

## Editor Plugins

_Plugin for text editors and IDEs._

- [vim-go](https://github.com/fatih/vim-go) - Go development plugin for Vim. Stars:`16.1K`.
- [gocode](https://github.com/nsf/gocode) - Autocompletion daemon for the Go programming language. Stars:`5.0K`.
- [vscode-go](https://github.com/golang/vscode-go) - Extension for Visual Studio Code (VS Code) which provides support for the Go language. Stars:`3.9K`.
- [GoSublime](https://github.com/DisposaBoy/GoSublime) - Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. Stars:`3.4K`.
- [go-plus](https://github.com/joefitzgerald/go-plus) - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. Stars:`1.5K`.
- [go-mode](https://github.com/dominikh/go-mode.el) - Go mode for GNU/Emacs. Stars:`1.4K`.

**[⬆ back to top](#contents)**

## Go Generate Tools

- [gotests](https://github.com/cweill/gotests) - Generate Go tests from your source code. Stars:`5.0K`.
- [genny](https://github.com/cheekybits/genny) - Elegant generics for Go. Stars:`1.7K`.

**[⬆ back to top](#contents)**

## Go Tools

- [go-swagger](https://github.com/go-swagger/go-swagger) - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. Stars:`9.6K`.
- [go-callvis](https://github.com/TrueFurby/go-callvis) - Visualize call graph of your Go program using dot format. Stars:`6.0K`.
- [OctoLinker](https://github.com/OctoLinker/browser-extension) - Navigate through go files efficiently with the OctoLinker browser extension for GitHub. Stars:`5.3K`.
- [go-size-analyzer](https://github.com/Zxilly/go-size-analyzer) - Analyze and visualize the size of dependencies in compiled Golang binaries, providing insight into their impact on the final build. Stars:`1.4K`.


**[⬆ back to top](#contents)**

## Software Packages

_Software written in Go._

**[⬆ back to top](#contents)**

### DevOps Tools

- [kubernetes](https://github.com/kubernetes/kubernetes) - Container Cluster Manager from Google. Stars:`112.3K`.
- [Moby](https://github.com/moby/moby) - Collaborative project for the container ecosystem to assemble container-based systems. Stars:`69.0K`.
- [traefik](https://github.com/containous/traefik) - Reverse proxy and load balancer with support for multiple backends. Stars:`52.7K`.
- [Gitea](https://github.com/go-gitea/gitea) - Fork of Gogs, entirely community driven. Stars:`46.6K`.
- [minikube](https://github.com/kubernetes/minikube) - Run Kubernetes locally. Stars:`29.8K`.
- [k3s](https://github.com/k3s-io/k3s) - Lightweight Kubernetes. Stars:`28.5K`.
- [k6](https://github.com/grafana/k6) - A modern load testing tool, using Go and JavaScript. Stars:`26.5K`.
- [Vegeta](https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000! Stars:`23.8K`.
- [Hey](https://github.com/rakyll/hey) - Hey is a tiny program that sends some load to a web application. Stars:`18.4K`.
- [Packer](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. Stars:`15.2K`.
- [kind](https://github.com/kubernetes-sigs/kind) - Kubernetes IN Docker - local clusters for testing Kubernetes. Stars:`13.7K`.
- [kubeshark](https://github.com/kubeshark/kubeshark) - API traffic analyzer for Kubernetes, inspired by Wireshark, purposely built for Kubernetes. Stars:`11.1K`.
- [GVM](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions. Stars:`10.5K`.
- [Flannel](https://github.com/flannel-io/flannel) - Flannel is a network fabric for containers, designed for Kubernetes. Stars:`8.9K`.
- [Ddosify](https://github.com/ddosify/ddosify) - High-performance load testing tool, written in Golang. Stars:`8.4K`.
- [ko](https://github.com/google/ko) - Command line tool for building and deploying Go applications on Kubernetes Stars:`7.8K`.
- [KubeVela](https://github.com/kubevela/kubevela) - Cloud native application delivery. Stars:`6.5K`.
- [bombardier](https://github.com/codesenberg/bombardier) - Fast cross-platform HTTP benchmarking tool. Stars:`6.1K`.
- [script](https://github.com/bitfield/script) - Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. Stars:`5.8K`.
- [k3d](https://github.com/k3d-io/k3d) - Little helper to run CNCF's k3s in Docker. Stars:`5.6K`.
- [podinfo](https://github.com/stefanprodan/podinfo) - Podinfo is a tiny web application made with Go that showcases best practices of running microservices in Kubernetes. Podinfo is used by CNCF projects like Flux and Flagger for end-to-end testing and workshops. Stars:`5.5K`.
- [gaia](https://github.com/gaia-pipeline/gaia) - Build powerful pipelines in any programming language. Stars:`5.2K`.
- [gox](https://github.com/mitchellh/gox) - Dead simple, no frills Go cross compile tool. Stars:`4.6K`.
- [Pomerium](https://github.com/pomerium/pomerium) - Pomerium is an identity-aware access proxy. Stars:`4.1K`.
- [tau](https://github.com/taubyte/tau) - Easily build Cloud Computing Platforms with features like Serverless WebAssembly Functions, Frontend Hosting, CI/CD, Object Storage, K/V Database, and Pub-Sub Messaging. Stars:`3.7K`.
- [bosun](https://github.com/bosun-monitor/bosun) - Time Series Alerting Framework. Stars:`3.4K`.
- [Fleet device management](https://github.com/fleetdm/fleet) - Lightweight, programmable telemetry for servers and workstations. Stars:`3.3K`.
- [s5cmd](https://github.com/peak/s5cmd) - Blazing fast S3 and local filesystem execution tool. Stars:`2.8K`.
- [kubeblocks](https://github.com/apecloud/kubeblocks) - KubeBlocks is an open-source control plane that runs and manages databases, message queues and other data infrastructure on K8s. Stars:`2.3K`.
- [kala](https://github.com/ajvb/kala) - Simplistic, modern, and performant job scheduler. Stars:`2.1K`.
- [fac](https://github.com/mkchoi212/fac) - Command-line user interface to fix git merge conflicts. Stars:`1.8K`.
- [StatusOK](https://github.com/sanathp/statusok) - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. Stars:`1.6K`.
- [ghorg](https://github.com/gabrie30/ghorg) - Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Gitea, and Bitbucket. Stars:`1.6K`.
- [go-selfupdate](https://github.com/sanbornm/go-selfupdate) - Enable your Go applications to self update. Stars:`1.6K`.
- [tlm](https://github.com/yusufcanb/tlm) - Local cli copilot, powered by CodeLLaMa Stars:`1.3K`.
- [uTask](https://github.com/ovh/utask) - Automation engine that models and executes business processes declared in yaml. Stars:`1.2K`.
- [s3gof3r](https://github.com/rlmcpherson/s3gof3r) - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. Stars:`1.1K`.
- [PipeCD](https://github.com/pipe-cd/pipecd) - A GitOps-style continuous delivery platform that provides consistent deployment and operations experience for any applications. Stars:`1.1K`.

**[⬆ back to top](#contents)**

### Other Software

- [croc](https://github.com/schollz/croc) - Easily and securely send files or folders from one computer to another. Stars:`28.5K`.
- [restic](https://github.com/restic/restic) - De-duplicating backup program. Stars:`27.2K`.
- [Seaweed File System](https://github.com/chrislusf/seaweedfs) - Fast, Simple and Scalable Distributed File System with O(1) disk seek. Stars:`23.5K`.
- [Gor](https://github.com/buger/gor) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. Stars:`18.7K`.
- [JuiceFS](https://github.com/juicedata/juicefs) - Distributed POSIX file system built on top of Redis and AWS S3. Stars:`11.1K`.
- [toxiproxy](https://github.com/shopify/toxiproxy) - Proxy to simulate network and system conditions for automated tests. Stars:`11.0K`.
- [Comcast](https://github.com/tylertreat/Comcast) - Simulate bad network connections. Stars:`10.3K`.
- [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul. Stars:`8.4K`.
- [LiteIDE](https://github.com/visualfc/liteide) - LiteIDE is a simple, open source, cross-platform Go IDE. Stars:`7.6K`.
- [scc](https://github.com/boyter/scc) - Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. Stars:`6.9K`.
- [drive](https://github.com/odeke-em/drive) - Google Drive client for the commandline. Stars:`6.7K`.
- [nes](https://github.com/fogleman/nes) - Nintendo Entertainment System (NES) emulator written in Go. Stars:`5.5K`.
- [Duplicacy](https://github.com/gilbertchen/duplicacy) - A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. Stars:`5.3K`.
- [blocky](https://github.com/0xERR0R/blocky) - Fast and lightweight DNS proxy as ad-blocker for local network with many features. Stars:`4.9K`.
- [myLG](https://github.com/mehrdadrad/mylg) - Command Line Network Diagnostic tool written in Go. Stars:`2.7K`.
- [GoBoy](https://github.com/Humpheh/goboy) - Nintendo Game Boy Color emulator written in Go. Stars:`2.6K`.
- [Stack Up](https://github.com/pressly/sup) - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. Stars:`2.5K`.
- [lgo](https://github.com/yunabe/lgo) - Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. Stars:`2.4K`.
- [Documize](https://github.com/documize/community) - Modern wiki software that integrates data from SaaS tools. Stars:`2.2K`.
- [sonic](https://github.com/go-sonic/sonic) - Sonic is a Go Blogging Platform. Simple and Powerful. Stars:`2.0K`.
- [Circuit](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. Stars:`2.0K`.
- [Gokapi](https://github.com/Forceu/gokapi) - Lightweight server to share files, which expire after a set amount of downloads or days. Similar to Firefox Send, but without public upload. Stars:`1.8K`.
- [borg](https://github.com/crufter/borg) - Terminal based search engine for bash snippets. Stars:`1.6K`.
- [portal](https://github.com/SpatiumPortae/portal) - Portal is a quick and easy command-line file transfer utility from any computer to another. Stars:`1.5K`.
- [Plik](https://github.com/root-gg/plik) - Plik is a temporary file upload system (Wetransfer like) in Go. Stars:`1.5K`.
- [shell2http](https://github.com/msoap/shell2http) - Executing shell commands via http server (for prototyping or remote control). Stars:`1.4K`.
- [vFlow](https://github.com/VerizonDigital/vflow) - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. Stars:`1.1K`.
- [peg](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. Stars:`1.0K`.

**[⬆ back to top](#contents)**

# Resources

_Where to discover new Go libraries._

**[⬆ back to top](#contents)**

## Benchmarks

- [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) - Go web framework benchmark. Stars:`2.1K`.
- [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router benchmark and comparison. Stars:`1.7K`.
- [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods. Stars:`1.6K`.
- [skynet](https://github.com/atemerev/skynet) - Skynet 1M threads microbenchmark. Stars:`1.0K`.

**[⬆ back to top](#contents)**

## Conferences


**[⬆ back to top](#contents)**

## E-Books

### E-books for purchase


### Free e-books

- [GoBooks](https://github.com/dariubs/GoBooks) - A curated list of Go books. Stars:`17.4K`.
- [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) Stars:`9.5K`.
- [Go AST Book (Chinese)](https://github.com/chai2010/go-ast-book) - A book focusing on Go `go/*` packages. Stars:`5.4K`.
- [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/) Stars:`3.2K`.

**[⬆ back to top](#contents)**

## Gophers

- [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) - Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. Stars:`3.6K`.
- [gophers](https://github.com/egonelbre/gophers) - Free gophers. Stars:`3.5K`.
- [gophers](https://github.com/ashleymcnamara/gophers) - Gopher artworks by Ashley McNamara. Stars:`3.0K`.

**[⬆ back to top](#contents)**

## Meetups


_Add the group of your city/country here (send **PR**)_

**[⬆ back to top](#contents)**

## Style Guides

- [CockroachDB](https://github.com/cockroachdb/cockroach/blob/master/docs/style.md) Stars:`30.4K`.
- [Uber](https://github.com/uber-go/guide/blob/master/style.md) Stars:`16.1K`.
- [Hyperledger](https://github.com/hyperledger/fabric/blob/release-1.4/docs/source/style-guides/go-style.rst) Stars:`15.9K`.
- [Magnetico](https://github.com/boramalper/magnetico/wiki/magnetico-Design-Specification) Stars:`3.1K`.
- [bahlo/go-styleguide](https://github.com/bahlo/go-styleguide) Stars:`1.5K`.

**[⬆ back to top](#contents)**

## Social Media

### Twitter


**[⬆ back to top](#contents)**

### Reddit


**[⬆ back to top](#contents)**

## Websites

- [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki. Stars:`125.1K`.
- [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) - Curated list of awesome remote jobs. A lot of them are looking for Go hackers. Stars:`36.1K`.
- [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) - List of other amazingly awesome lists. Stars:`32.3K`.

**[⬆ back to top](#contents)**

### Tutorials

- [Build a Database in 1000 lines of code]( https://link.medium.com/O9YQlx89Htb) - Build a NoSQL Database From Zero in 1000 Lines of Code.
- [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) - Golang ebook intro how to build a web app with golang. Stars:`43.4K`.
- [go-patterns](https://github.com/tmrts/go-patterns) - Curated list of Go design patterns, recipes and idioms. Stars:`25.7K`.
- [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) - Learn Go with test-driven development. Stars:`22.4K`.
- [Learn Go with 1000+ Exercises](https://github.com/inancgumus/learngo) - Learn Go with thousands of examples, exercises, and quizzes. Stars:`19.0K`.
- [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) - Go's reference card. Stars:`8.5K`.
- [go-clean-template](https://github.com/evrone/go-clean-template) - Clean Architecture template for Golang services. Stars:`6.4K`.
- [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) - Examples of Golang compared to Node.js for learning. Stars:`4.7K`.
- [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) - A little e-book on Ethereum Development with Go. Stars:`1.7K`.
- [golang-examples](https://github.com/SimonWaldherr/golang-examples) - Many examples to learn Golang. Stars:`1.6K`.
- [Working with Go](https://github.com/mkaz/working-with-go) - Intro to go for experienced programmers. Stars:`1.1K`.

**[⬆ back to top](#contents)**

### Guided Learning


**[⬆ back to top](#contents)**
