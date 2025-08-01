# Awesome Go

<a href="https://awesome-go.com/"><img align="right" src="https://github.com/avelino/awesome-go/raw/main/tmpl/assets/logo.png" alt="awesome-go" title="awesome-go" /></a>

[![Build Status](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml/badge.svg?branch=main)](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml?query=branch%3Amain)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)
[![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](https://gophers.slack.com/messages/awesome)
[![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)
[![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/avelino/awesome-go/)
[![Last Commit](https://img.shields.io/github/last-commit/avelino/awesome-go)](https://github.com/avelino/awesome-go/commits/main)

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

- [Awesome Go](#awesome-go)
  - [Contents](#contents)
  - [Actor Model](#actor-model)
  - [Artificial Intelligence](#artificial-intelligence)
  - [Audio and Music](#audio-and-music)
  - [Authentication and OAuth](#authentication-and-oauth)
  - [Blockchain](#blockchain)
  - [Bot Building](#bot-building)
  - [Build Automation](#build-automation)
  - [Command Line](#command-line)
    - [Advanced Console UIs](#advanced-console-uis)
    - [Standard CLI](#standard-cli)
  - [Configuration](#configuration)
  - [Continuous Integration](#continuous-integration)
  - [CSS Preprocessors](#css-preprocessors)
  - [Data Integration Frameworks](#data-integration-frameworks)
  - [Data Structures and Algorithms](#data-structures-and-algorithms)
    - [Bit-packing and Compression](#bit-packing-and-compression)
    - [Bit Sets](#bit-sets)
    - [Bloom and Cuckoo Filters](#bloom-and-cuckoo-filters)
    - [Data Structure and Algorithm Collections](#data-structure-and-algorithm-collections)
    - [Iterators](#iterators)
    - [Maps](#maps)
    - [Miscellaneous Data Structures and Algorithms](#miscellaneous-data-structures-and-algorithms)
    - [Nullable Types](#nullable-types)
    - [Queues](#queues)
    - [Sets](#sets)
    - [Text Analysis](#text-analysis)
    - [Trees](#trees)
    - [Pipes](#pipes)
  - [Database](#database)
    - [Caches](#caches)
    - [Databases Implemented in Go](#databases-implemented-in-go)
    - [Database Schema Migration](#database-schema-migration)
    - [Database Tools](#database-tools)
    - [SQL Query Builders](#sql-query-builders)
  - [Database Drivers](#database-drivers)
    - [Interfaces to Multiple Backends](#interfaces-to-multiple-backends)
    - [Relational Database Drivers](#relational-database-drivers)
    - [NoSQL Database Drivers](#nosql-database-drivers)
    - [Search and Analytic Databases](#search-and-analytic-databases)
  - [Date and Time](#date-and-time)
  - [Distributed Systems](#distributed-systems)
  - [Dynamic DNS](#dynamic-dns)
  - [Email](#email)
  - [Embeddable Scripting Languages](#embeddable-scripting-languages)
  - [Error Handling](#error-handling)
  - [File Handling](#file-handling)
  - [Financial](#financial)
  - [Forms](#forms)
  - [Functional](#functional)
  - [Game Development](#game-development)
  - [Generators](#generators)
  - [Geographic](#geographic)
  - [Go Compilers](#go-compilers)
  - [Goroutines](#goroutines)
  - [GUI](#gui)
  - [Hardware](#hardware)
  - [Images](#images)
  - [IoT (Internet of Things)](#iot-internet-of-things)
  - [Job Scheduler](#job-scheduler)
  - [JSON](#json)
  - [Logging](#logging)
  - [Machine Learning](#machine-learning)
  - [Messaging](#messaging)
  - [Microsoft Office](#microsoft-office)
    - [Microsoft Excel](#microsoft-excel)
    - [Microsoft Word](#microsoft-word)
  - [Miscellaneous](#miscellaneous)
    - [Dependency Injection](#dependency-injection)
    - [Project Layout](#project-layout)
    - [Strings](#strings)
    - [Uncategorized](#uncategorized)
  - [Natural Language Processing](#natural-language-processing)
    - [Language Detection](#language-detection)
    - [Morphological Analyzers](#morphological-analyzers)
    - [Slugifiers](#slugifiers)
    - [Tokenizers](#tokenizers)
    - [Translation](#translation)
    - [Transliteration](#transliteration)
  - [Networking](#networking)
    - [HTTP Clients](#http-clients)
  - [OpenGL](#opengl)
  - [ORM](#orm)
  - [Package Management](#package-management)
  - [Performance](#performance)
  - [Query Language](#query-language)
  - [Reflection](#reflection)
  - [Resource Embedding](#resource-embedding)
  - [Science and Data Analysis](#science-and-data-analysis)
  - [Security](#security)
  - [Serialization](#serialization)
  - [Server Applications](#server-applications)
  - [Stream Processing](#stream-processing)
  - [Template Engines](#template-engines)
  - [Testing](#testing)
    - [Testing Frameworks](#testing-frameworks)
    - [Mock](#mock)
    - [Fuzzing and delta-debugging/reducing/shrinking](#fuzzing-and-delta-debuggingreducingshrinking)
    - [Selenium and browser control tools](#selenium-and-browser-control-tools)
    - [Fail injection](#fail-injection)
  - [Text Processing](#text-processing)
    - [Formatters](#formatters)
    - [Markup Languages](#markup-languages)
    - [Parsers/Encoders/Decoders](#parsersencodersdecoders)
    - [Regular Expressions](#regular-expressions)
    - [Sanitation](#sanitation)
    - [Scrapers](#scrapers)
    - [RSS](#rss)
    - [Utility/Miscellaneous](#utilitymiscellaneous)
  - [Third-party APIs](#third-party-apis)
  - [Utilities](#utilities)
  - [UUID](#uuid)
  - [Validation](#validation)
  - [Version Control](#version-control)
  - [Video](#video)
  - [Web Frameworks](#web-frameworks)
    - [Middlewares](#middlewares)
      - [Actual middlewares](#actual-middlewares)
      - [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
    - [Routers](#routers)
  - [WebAssembly](#webassembly)
  - [Webhooks Server](#webhooks-server)
  - [Windows](#windows)
  - [Workflow Frameworks](#workflow-frameworks)
  - [XML](#xml)
  - [Zero Trust](#zero-trust)
  - [Code Analysis](#code-analysis)
  - [Editor Plugins](#editor-plugins)
  - [Go Generate Tools](#go-generate-tools)
  - [Go Tools](#go-tools)
  - [Software Packages](#software-packages)
    - [DevOps Tools](#devops-tools)
    - [Other Software](#other-software)
- [Resources](#resources)
  - [Benchmarks](#benchmarks)
  - [Conferences](#conferences)
  - [E-Books](#e-books)
    - [E-books for purchase](#e-books-for-purchase)
    - [Free e-books](#free-e-books)
  - [Gophers](#gophers)
  - [Meetups](#meetups)
  - [Style Guides](#style-guides)
  - [Social Media](#social-media)
    - [Twitter](#twitter)
    - [Reddit](#reddit)
  - [Websites](#websites)
    - [Tutorials](#tutorials)
    - [Guided Learning](#guided-learning)

**[⬆ back to top](#contents)**

## Actor Model

_Libraries for building actor-based programs._

- [ProtoActor](https://github.com/asynkron/protoactor-go) - Proto Actor - Ultra fast distributed actors for Go, C# and Java/Kotlin. Stars:`5.3K`.
- [Hollywood](https://github.com/anthdm/hollywood) - Blazingly fast and light-weight Actor engine written in Golang. Stars:`2.0K`.
- [Goakt](https://github.com/Tochemey/goakt) - Fast and Distributed Actor framework using protocol buffers as message for Golang. Stars:`272`.
- [Ergo](https://github.com/ergo-services/ergo) - An actor-based Framework with network transparency for creating event-driven architecture in Golang. Inspired by Erlang.

## Artificial Intelligence

_Libraries for building programs that leverage AI._

- [Ollama](https://github.com/jmorganca/ollama) - Run large language models locally. Stars:`148.3K`.
- [LocalAI](https://github.com/mudler/LocalAI) - Open Source OpenAI alternative, self-host AI models. Stars:`34.3K`.
- [langchaingo](https://github.com/tmc/langchaingo) - LangChainGo is a framework for developing applications powered by language models. Stars:`7.3K`.
- [chromem-go](https://github.com/philippgille/chromem-go) - Embeddable vector database for Go with Chroma-like interface and zero third-party dependencies. In-memory with optional persistence. Stars:`648`.
- [OllamaFarm](https://github.com/presbrey/ollamafarm) - Manage, load-balance, and failover packs of Ollamas Stars:`81`.
- [fun](https://gitlab.com/tozd/go/fun) - The simplest but powerful way to use large language models (LLMs) in Go.

**[⬆ back to top](#contents)**

## Audio and Music

_Libraries for manipulating audio._

- [Oto](https://github.com/hajimehoshi/oto) - A low-level library to play sound on multiple platforms. Stars:`1.7K`.
- [PortAudio](https://github.com/gordonklaus/portaudio) - Go bindings for the PortAudio audio I/O library. Stars:`783`.
- [GoAudio](https://github.com/DylanMeeus/GoAudio) - Native Go Audio Processing Library. Stars:`378`.
- [id3v2](https://github.com/bogem/id3v2) - ID3 decoding and encoding library for Go. Stars:`350`.
- [malgo](https://github.com/gen2brain/malgo) - Mini audio library. Stars:`342`.
- [flac](https://github.com/mewkiz/flac) - Native Go FLAC encoder/decoder with support for FLAC streams. Stars:`337`.
- [gaad](https://github.com/Comcast/gaad) - Native Go AAC bitstream parser. Stars:`131`.
- [minimp3](https://github.com/tosone/minimp3) - Lightweight MP3 decoder library. Stars:`129`.
- [gosamplerate](https://github.com/dh1tw/gosamplerate) - libsamplerate bindings for go. Stars:`36`.

**[⬆ back to top](#contents)**

## Authentication and Authorization

_Libraries for implementing authentication and authorization._

- [casbin](https://github.com/hsluoyz/casbin) - Authorization library that supports access control models like ACL, RBAC, and ABAC. Stars:`18.9K`.
- [jwt-go](https://github.com/golang-jwt/jwt) - A full featured implementation of JSON Web Tokens (JWT). This library supports the parsing and verification as well as the generation and signing of JWTs. Stars:`8.3K`.
- [goth](https://github.com/markbates/goth) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. Stars:`6.1K`.
- [spicedb](https://github.com/authzed/spicedb) - A Zanzibar-inspired database that enables fine-grained authorization. Stars:`5.9K`.
- [oauth2](https://github.com/golang/oauth2) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine, and App Engine support. Stars:`5.7K`.
- [keto](https://github.com/ory/keto) - Open Source (Go) implementation of "Zanzibar: Google's Consistent, Global Authorization System". Ships gRPC, REST APIs, newSQL, and an easy and granular permission language. Supports ACL, RBAC, and other access models. Stars:`5.1K`.
- [authboss](https://github.com/volatiletech/authboss) - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure it, and start building your app without having to build an authentication system each time. Stars:`4.0K`.
- [openfga](https://github.com/openfga/openfga) - Implementation of fine-grained authorization based on the "Zanzibar: Google's Consistent, Global Authorization System" paper. Backed by [CNCF](https://www.cncf.io/). Stars:`3.9K`.
- [scs](https://github.com/alexedwards/scs) - Session Manager for HTTP servers. Stars:`2.4K`.
- [jwx](https://github.com/lestrrat-go/jwx) - Go module implementing various JWx (JWA/JWE/JWK/JWS/JWT, otherwise known as JOSE) technologies Stars:`2.2K`.
- [osin](https://github.com/openshift/osin) - Golang OAuth2 server library. Stars:`1.9K`.
- [loginsrv](https://github.com/tarent/loginsrv) - JWT login microservice with pluggable backends such as OAuth2 (Github), htpasswd, osiam. Stars:`1.9K`.
- [gologin](https://github.com/dghubble/gologin) - chainable handlers for login with OAuth1 and OAuth2 authentication providers. Stars:`1.9K`.
- [gorbac](https://github.com/mikespook/gorbac) - provides a lightweight role-based access control (RBAC) implementation in Golang. Stars:`1.6K`.
- [oidc](https://github.com/zitadel/oidc) - Easy to use OpenID Connect client and server library written for Go and certified by the OpenID Foundation Stars:`1.6K`.
- [paseto](https://github.com/o1egl/paseto) - Golang implementation of Platform-Agnostic Security Tokens (PASETO). Stars:`902`.
- [jwt](https://github.com/cristalhq/jwt) - Safe, simple, and fast JSON Web Tokens for Go. Stars:`687`.
- [go-guardian](https://github.com/shaj13/go-guardian) - Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token, and Certificate based authentication. Stars:`589`.
- [go-jose](https://github.com/go-jose/go-jose) - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. Stars:`424`.
- [jwt](https://github.com/pascaldekloe/jwt) - Lightweight JSON Web Token (JWT) library. Stars:`363`.
- [jeff](https://github.com/abraithwaite/jeff) - Simple, flexible, secure, and idiomatic web session management with pluggable backends. Stars:`269`.
- [gosession](https://github.com/Kwynto/gosession) - This is quick session for net/http in GoLang. This package is perhaps the best implementation of the session mechanism, or at least it tries to become one. Stars:`258`.
- [jwt-auth](https://github.com/adam-hanna/jwt-auth) - JWT middleware for Golang http servers with many configuration options. Stars:`234`.
- [goiabada](https://github.com/leodip/goiabada) - An open-source authentication and authorization server supporting OAuth2 and OpenID Connect. Stars:`157`.
- [otpgen](https://github.com/grijul/otpgen) - Library to generate TOTP/HOTP codes. Stars:`140`.
- [sessionup](https://github.com/swithek/sessionup) - Simple, yet effective HTTP session management and identification package. Stars:`128`.
- [sjwt](https://github.com/brianvoe/sjwt) - Simple jwt generator and parser. Stars:`123`.
- [session](https://github.com/icza/session) - Go session management for web servers (including support for Google App Engine - GAE). Stars:`118`.
- [branca](https://github.com/essentialkaos/branca) - branca token [specification implementation](https://github.com/tuupola/branca-spec) for Golang 1.15+. Stars:`91`.
- [securecookie](https://github.com/chmike/securecookie) - Efficient secure cookie encoding/decoding. Stars:`82`.
- [sessions](https://github.com/adam-hanna/sessions) - Dead simple, highly performant, highly customizable sessions service for go http servers. Stars:`78`.
- [otpgo](https://github.com/jltorresm/otpgo) - Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. Stars:`76`.
- [scope](https://github.com/SonicRoshan/scope) - Easily Manage OAuth2 Scopes In Go. Stars:`41`.
- [cookiestxt](https://github.com/mengzhuo/cookiestxt) - provides a parser of cookies.txt file format. Stars:`20`.
- [permissions](https://github.com/xyproto/permissions) - Library for keeping track of users, login states, and permissions. Uses secure cookies and bcrypt. Stars:`10`.
- [x509proxy](https://github.com/vkuznet/x509proxy) - Library to handle X509 proxy certificates. Stars:`7`.

**[⬆ back to top](#contents)**

## Blockchain

_Tools for building blockchains._

- [go-ethereum](https://github.com/ethereum/go-ethereum) - Official Go implementation of the Ethereum protocol. Stars:`49.4K`.
- [kubo](https://github.com/ipfs/kubo) - A blockchain framework implemented in Go. It provides content-addressable storage which can be used for decentralized storage in DApps. It is based on the IPFS protocol. Stars:`16.6K`.
- [lnd](https://github.com/lightningnetwork/lnd) - A complete implementation of a Lightning Network node. Stars:`8.0K`.
- [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) - A Framework for Building Public Blockchains in the Cosmos Ecosystem. Stars:`6.7K`.
- [tendermint](https://github.com/tendermint/tendermint) - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. Stars:`5.8K`.
- [solana-go](https://github.com/gagliardetto/solana-go) - Go library to interface with Solana JSON RPC and WebSocket interfaces. Stars:`1.4K`.
- [gno](https://github.com/gnolang/gno) - A comprehensive smart contract suite built with Golang and Gnolang, a deterministic, purpose-built Go variant for blockchains. Stars:`980`.
- [cometbft](https://github.com/cometbft/cometbft) - A distributed, Byzantine fault-tolerant, deterministic state machine replication engine. It is a fork of Tendermint Core and implements the Tendermint consensus algorithm. Stars:`765`.
- [gossamer](https://github.com/ChainSafe/gossamer) - A Go implementation of the Polkadot Host. Stars:`451`.
- [gosemble](https://github.com/LimeChain/gosemble) - A Go-based framework for building Polkadot/Substrate-compatible runtimes. Stars:`14`.

**[⬆ back to top](#contents)**

## Bot Building

_Libraries for building and working with bots._

- [telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) - Simple and clean Telegram bot client. Stars:`6.2K`.
- [telebot](https://github.com/tucnak/telebot) - Telegram bot framework is written in Go. Stars:`4.4K`.
- [wayback](https://github.com/wabarc/wayback) - A bot for Telegram, Mastodon, Slack, and other messaging platforms archives webpages. Stars:`2.0K`.
- [bot](https://github.com/go-telegram/bot) - Zero-dependencies Telegram Bot library with additional UI components Stars:`1.3K`.
- [telego](https://github.com/mymmrac/telego) - Telegram Bot API library for Golang with full one-to-one API implementation. Stars:`747`.
- [arikawa](https://github.com/diamondburned/arikawa) - A library and framework for the Discord API. Stars:`537`.
- [echotron](https://github.com/NicoNex/echotron) - An elegant and concurrent library for Telegram Bots in Go. Stars:`406`.
- [go-twitch-irc](https://github.com/gempir/go-twitch-irc) - Library to write bots for twitch.tv chat Stars:`378`.
- [go-sarah](https://github.com/oklahomer/go-sarah) - Framework to build a bot for desired chat services including LINE, Slack, Gitter, and more. Stars:`264`.
- [slack-bot](https://github.com/innogames/slack-bot) - Ready to use Slack Bot for lazy developers: Custom commands, Jenkins, Jira, Bitbucket, Github... Stars:`199`.
- [go-tg](https://github.com/mr-linch/go-tg) - Generated from official docs Go client library for accessing Telegram Bot API, with batteries for building complex bots included. Stars:`120`.
- [slacker](https://github.com/slack-io/slacker) - Easy to use framework to create Slack bots. Stars:`55`.
- [micha](https://github.com/onrik/micha) - Go Library for Telegram bot api. Stars:`30`.
- [go-joe](https://joe-bot.net) - A general-purpose bot library inspired by Hubot but written in Go.

**[⬆ back to top](#contents)**

## Build Automation

_Libraries and tools help with build automation._

- [air](https://github.com/cosmtrek/air) - Air - Live reload for Go apps. Stars:`21.3K`.
- [Task](https://github.com/go-task/task) - simple "Make" alternative. Stars:`13.3K`.
- [realize](https://github.com/tockins/realize) - Go build a system with file watchers and live to reload. Run, build and watch file changes with custom paths. Stars:`4.5K`.
- [mage](https://github.com/magefile/mage) - Mage is a make/rake-like build tool using Go. Stars:`4.4K`.
- [mmake](https://github.com/tj/mmake) - Modern Make. Stars:`1.7K`.
- [xc](https://github.com/joerdav/xc) - Task runner with README.md defined tasks, executable markdown. Stars:`1.3K`.
- [goyek](https://github.com/goyek/goyek) - Create build pipelines in Go. Stars:`630`.
- [taskctl](https://github.com/taskctl/taskctl) - Concurrent task runner. Stars:`314`.
- [1build](https://github.com/gopinath-langote/1build) - Command line tool to frictionlessly manage project-specific commands. Stars:`235`.
- [gaper](https://github.com/maxclaus/gaper) - Builds and restarts a Go project when it crashes or some watched file changes. Stars:`83`.
- [anko](https://github.com/GuilhermeCaruso/anko) - Simple application watcher for multiple programming languages. Stars:`43`.
- [gob](https://github.com/kcmvp/gob) - [Gradle](https://docs.gradle.org/)/[Maven](https://maven.apache.org/) like build tool for Go projects. Stars:`13`.
- [gilbert](https://go-gilbert.github.io) - Build system and task runner for Go projects.

**[⬆ back to top](#contents)**

## Command Line

### Advanced Console UIs

_Libraries for building Console Applications and Console User Interfaces._

- [bubbletea](https://github.com/charmbracelet/bubbletea) - Go framework to build terminal apps, based on The Elm Architecture. Stars:`33.8K`.
- [fx](https://github.com/antonmedv/fx) - Terminal JSON viewer & processor. Stars:`19.7K`.
- [termui](https://github.com/gizak/termui) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). Stars:`13.4K`.
- [gocui](https://github.com/jroimartin/gocui) - Minimalist Go library aimed at creating Console User Interfaces. Stars:`10.3K`.
- [lipgloss](https://github.com/charmbracelet/lipgloss) - Declaratively define styles for color, format and layout in the terminal. Stars:`9.3K`.
- [bubbles](https://github.com/charmbracelet/bubbles) - TUI components for bubbletea. Stars:`6.7K`.
- [go-prompt](https://github.com/c-bata/go-prompt) - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). Stars:`5.4K`.
- [pterm](https://github.com/pterm/pterm) - A library to beautify console output on every platform with many combinable components. Stars:`5.2K`.
- [termbox-go](https://github.com/nsf/termbox-go) - Termbox is a library for creating cross-platform text-based interfaces. Stars:`4.7K`.
- [progressbar](https://github.com/schollz/progressbar) - Basic thread-safe progress bar that works in every OS. Stars:`4.5K`.
- [termdash](https://github.com/mum4k/termdash) - Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). Stars:`2.9K`.
- [asciigraph](https://github.com/guptarohit/asciigraph) - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. Stars:`2.9K`.
- [spinner](https://github.com/briandowns/spinner) - Go package to easily provide a terminal spinner with options. Stars:`2.5K`.
- [mpb](https://github.com/vbauerster/mpb) - Multi progress bar for terminal applications. Stars:`2.4K`.
- [uiprogress](https://github.com/gosuri/uiprogress) - Flexible library to render progress bars in terminal applications. Stars:`2.1K`.
- [termenv](https://github.com/muesli/termenv) - Advanced ANSI style & color support for your terminal applications. Stars:`1.9K`.
- [uilive](https://github.com/gosuri/uilive) - Library for updating terminal output in real time. Stars:`1.7K`.
- [gookit/color](https://github.com/gookit/color) - Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. Stars:`1.6K`.
- [aurora](https://github.com/logrusorgru/aurora) - ANSI terminal colors that support fmt.Printf/Sprintf. Stars:`1.5K`.
- [go-isatty](https://github.com/mattn/go-isatty) - isatty for golang. Stars:`866`.
- [go-colorable](https://github.com/mattn/go-colorable) - Colorable writer for windows. Stars:`795`.
- [uitable](https://github.com/gosuri/uitable) - Library to improve readability in terminal apps using tabular data. Stars:`741`.
- [box-cli-maker](https://github.com/Delta456/box-cli-maker) - Make Highly Customized Boxes for your CLI. Stars:`570`.
- [gommon/color](https://github.com/labstack/gommon/tree/master/color) - Style terminal text. Stars:`556`.
- [simpletable](https://github.com/alexeyco/simpletable) - Simple tables in a terminal with Go. Stars:`544`.
- [bubble-table](https://github.com/Evertras/bubble-table) - An interactive table component for bubbletea. Stars:`508`.
- [chalk](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output. Stars:`466`.
- [yacspin](https://github.com/theckman/yacspin) - Yet Another CLi Spinner package, for working with terminal spinners. Stars:`450`.
- [tabby](https://github.com/cheynewallace/tabby) - A tiny library for super simple Golang tables. Stars:`359`.
- [go-colortext](https://github.com/daviddengcn/go-colortext) - Go library for color output in terminals. Stars:`216`.
- [cfmt](https://github.com/mingrammer/cfmt) - Contextual fmt inspired by bootstrap color classes. Stars:`107`.
- [cfmt](https://github.com/i582/cfmt) - Simple and convenient formatted stylized output fully compatible with fmt library. Stars:`70`.
- [marker](https://github.com/cyucelen/marker) - Easiest way to match and mark strings for colorful terminal outputs. Stars:`51`.
- [ctc](https://github.com/wzshiming/ctc) - The non-invasive cross-platform terminal color library does not need to modify the Print method. Stars:`50`.
- [table](https://github.com/tomlazar/table) - Small library for terminal color based tables. Stars:`50`.
- [go-palette](https://github.com/abusomani/go-palette) - Go library that provides elegant and convenient style definitions using ANSI colors. Fully compatible & wraps the [fmt library](https://pkg.go.dev/fmt) for nice terminal layouts. Stars:`20`.
- [go-ataman](https://github.com/workanator/go-ataman) - Go library for rendering ANSI colored text templates in terminals. Stars:`19`.
- [crab-config-files-templating](https://github.com/alfiankan/crab-config-files-templating) - Dynamic configuration file templating tool for kubernetes manifest or general configuration files. Stars:`8`.

**[⬆ back to top](#contents)**

### Standard CLI

_Libraries for building standard or basic Command Line applications._

- [cobra](https://github.com/spf13/cobra) - Commander for modern Go CLI interactions. Stars:`41.3K`.
- [urfave/cli](https://github.com/urfave/cli) - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). Stars:`23.4K`.
- [elvish](https://github.com/elves/elvish) - An expressive programming language and a versatile interactive shell. Stars:`6.1K`.
- [kingpin](https://github.com/alecthomas/kingpin) - Command line and flag parser supporting sub commands (superseded by `kong`; see below). Stars:`3.5K`.
- [Dnote](https://github.com/dnote/dnote) - A simple command line notebook with multi-device sync. Stars:`2.9K`.
- [go-flags](https://github.com/jessevdk/go-flags) - go command line option parser. Stars:`2.7K`.
- [pflag](https://github.com/spf13/pflag) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. Stars:`2.6K`.
- [go-arg](https://github.com/alexflint/go-arg) - Struct-based argument parsing in Go. Stars:`2.2K`.
- [ops](https://github.com/nanovms/ops) - Unikernel Builder/Orchestrator. Stars:`1.4K`.
- [carapace-bin](https://github.com/rsteube/carapace-bin) - Multi-shell multi-command argument completer. Stars:`1.4K`.
- [liner](https://github.com/peterh/liner) - Go readline-like library for command-line interfaces. Stars:`1.1K`.
- [complete](https://github.com/posener/complete) - Write bash completions in Go + Go command bash completion. Stars:`939`.
- [mow.cli](https://github.com/jawher/mow.cli) - Go library for building CLI applications with sophisticated flag and argument parsing and validation. Stars:`886`.
- [flaggy](https://github.com/integrii/flaggy) - A robust and idiomatic flags package with excellent subcommand support. Stars:`864`.
- [carapace](https://github.com/rsteube/carapace) - Command argument completion generator for spf13/cobra. Stars:`775`.
- [mkideal/cli](https://github.com/mkideal/cli) - Feature-rich and easy to use command-line package based on golang struct tags. Stars:`729`.
- [argparse](https://github.com/akamensky/argparse) - Command line argument parser inspired by Python's argparse module. Stars:`623`.
- [wmenu](https://github.com/dixonwille/wmenu) - Easy to use menu structure for cli applications that prompt users to make choices. Stars:`224`.
- [climax](https://github.com/tucnak/climax) - Alternative CLI with "human face", in spirit of Go command. Stars:`218`.
- [clîr](https://github.com/leaanthony/clir) - A Simple and Clear CLI library. Dependency free. Stars:`193`.
- [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli) - cli application framework with auto configuration and dependency injection. Stars:`181`.
- [commandeer](https://github.com/jaffee/commandeer) - Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. Stars:`176`.
- [sflags](https://github.com/octago/sflags) - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin, and other libraries. Stars:`162`.
- [job](https://github.com/liujianping/job) - JOB, make your short-term command as a long-term job. Stars:`148`.
- [cmdr](https://github.com/hedzr/cmdr) - A POSIX/GNU style, getopt-like command-line UI Go library. Stars:`140`.
- [teris-io/cli](https://github.com/teris-io/cli) - Simple and complete API for building command line interfaces in Go. Stars:`131`.
- [ukautz/clif](https://github.com/ukautz/clif) - Small command line interface framework. Stars:`129`.
- [acmd](https://github.com/cristalhq/acmd) - Simple, useful, and opinionated CLI package in Go. Stars:`128`.
- [env](https://github.com/codingconcepts/env) - Tag-based environment configuration for structs. Stars:`121`.
- [readline](https://github.com/reeflective/readline) - Shell library with modern and easy to use UI features. Stars:`119`.
- [version](https://github.com/mszostok/version) - Collects and displays CLI version information in multiple formats along with upgrade notice. Stars:`109`.
- [console](https://github.com/reeflective/console) Closed-loop application library for Cobra commands, with oh-my-posh prompts, and more. Stars:`97`.
- [wlog](https://github.com/dixonwille/wlog) - Simple logging interface that supports cross-platform color and concurrency. Stars:`67`.
- [command-chain](https://github.com/rainu/go-command-chain) - A go library for configure and run command chains - such as pipelining in unix shells. Stars:`67`.
- [gocmd](https://github.com/devfacet/gocmd) - Go library for building command line applications. Stars:`66`.
- [strumt](https://github.com/antham/strumt) - Library to create prompt chain. Stars:`62`.
- [go-getoptions](https://github.com/DavidGamba/go-getoptions) - Go option parser inspired by the flexibility of Perl’s GetOpt::Long. Stars:`61`.
- [flagvar](https://github.com/sgreben/flagvar) - A collection of flag argument types for Go's standard `flag` package. Stars:`45`.
- [cmd](https://github.com/posener/cmd) - Extends the standard `flag` package to support sub commands and more in idiomatic way. Stars:`44`.
- [mcli](https://github.com/jxskiss/mcli) - A minimal but very powerful cli library for Go. Stars:`41`.
- [argv](https://github.com/cosiner/argv) - Go library to split command line string as arguments array using the bash syntax. Stars:`41`.
- [go-readline-ny](https://github.com/nyaosorg/go-readline-ny) - A customizable line-editing library with Emacs keybindings, Unicode support, completion, and syntax highlighting. Used in NYAGOS shell. Stars:`29`.
- [hashicorp/cli](https://github.com/hashicorp/cli) - Go library for implementing command-line interfaces. Stars:`27`.
- [sand](https://github.com/Zaba505/sand) - Simple API for creating interpreters and so much more. Stars:`27`.
- [ts](https://github.com/liujianping/ts) - Timestamp convert & compare tool. Stars:`22`.
- [carapace-spec](https://github.com/rsteube/carapace-spec) - Define simple completions using a spec file. Stars:`22`.
- [subcmd](https://github.com/bobg/subcmd) - Another approach to parsing and running subcommands. Works alongside the standard `flag` package. Stars:`8`.
- [getopt](https://github.com/jon-codes/getopt) - An accurate Go `getopt`, validated against the GNU libc implementation. Stars:`7`.
- [goopt](https://github.com/napalu/goopt) - A declarative, struct-tag based CLI framework for Go, with a broad feature set such as hierarchical commands/flags, i18n, shell completion, and validation. Stars:`4`.

**[⬆ back to top](#contents)**

## Configuration

_Libraries for configuration parsing._

- [viper](https://github.com/spf13/viper) - Go configuration with fangs. Stars:`29.0K`.
- [godotenv](https://github.com/joho/godotenv) - Go port of Ruby's dotenv library (Loads environment variables from `.env`). Stars:`9.6K`.
- [sonic](https://github.com/bytedance/sonic) - A blazingly fast JSON serializing & deserializing library. Stars:`8.4K`.
- [env](https://github.com/caarlos0/env) - Parse environment variables to Go structs (with defaults). Stars:`5.7K`.
- [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) - Go library for managing configuration data from environment variables. Stars:`5.3K`.
- [ini](https://github.com/go-ini/ini) - Go package to read and write INI files. Stars:`3.5K`.
- [koanf](https://github.com/knadh/koanf) - Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. Stars:`3.4K`.
- [kong](https://github.com/alecthomas/kong) - Command-line parser with support for arbitrarily complex command-line structures and additional sources of configuration such as YAML, JSON, TOML, etc (successor to `kingpin`). Stars:`2.7K`.
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Minimalistic configuration reader (from files, ENV, and wherever you want). Stars:`1.9K`.
- [xdg](https://github.com/adrg/xdg) - Go implementation of the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/latest/) and [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories). Stars:`854`.
- [konfig](https://github.com/lalamove/konfig) - Composable, observable and performant config handling for Go for the distributed processing era. Stars:`645`.
- [aconfig](https://github.com/cristalhq/aconfig) - Simple, useful and opinionated config loader. Stars:`596`.
- [gookit/config](https://github.com/gookit/config) - application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. Stars:`561`.
- [confita](https://github.com/heetch/confita) - Load configuration in cascade from multiple backends into a struct. Stars:`504`.
- [fig](https://github.com/kkyr/fig) - Tiny library for reading configuration from a file and from environment variables (with validation & defaults). Stars:`383`.
- [GoLobby/Config](https://github.com/golobby/config) - GoLobby Config is a lightweight yet powerful configuration manager for the Go programming language. Stars:`366`.
- [hjson](https://github.com/hjson/hjson-go) - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. Stars:`338`.
- [konf](https://github.com/nil-go/konf) - The simplest API for reading/watching config from file, env, flag and clouds (e.g. AWS, Azure, GCP). Stars:`338`.
- [config](https://github.com/JeremyLoy/config) - Cloud native application configuration. Bind ENV to structs in only two lines. Stars:`338`.
- [store](https://github.com/tucnak/store) - Lightweight configuration manager for Go. Stars:`274`.
- [envconfig](https://github.com/vrischmann/envconfig) - Read your configuration from environment variables. Stars:`244`.
- [joshbetz/config](https://github.com/joshbetz/config) - Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. Stars:`214`.
- [zerocfg](https://github.com/chaindead/zerocfg) - Zero-effort, concise configuration management that avoids boilerplate and repetitive code, supports multiple sources with priority overrides. Stars:`192`.
- [harvester](https://github.com/beatlabs/harvester) - Harvester, a easy to use static and dynamic configuration package supporting seeding, env vars and Consul integration. Stars:`134`.
- [onion](https://github.com/goraz/onion) - Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP. Stars:`118`.
- [configuration](https://github.com/BoRuDar/configuration) - Library for initializing configuration structs from env variables, files, flags and 'default' tag. Stars:`108`.
- [envh](https://github.com/antham/envh) - Helpers to manage environment variables. Stars:`100`.
- [configuro](https://github.com/sherifabdlnaby/configuro) - opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications. Stars:`94`.
- [hocon](https://github.com/gurkankaymak/hocon) - Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files. Stars:`85`.
- [env](https://github.com/junk1tm/env) - A lightweight package for loading environment variables into structs. Stars:`72`.
- [uConfig](https://github.com/omeid/uconfig) - Lightweight, zero-dependency, and extendable configuration management. Stars:`72`.
- [gofigure](https://github.com/ian-kent/gofigure) - Go application configuration made easy. Stars:`68`.
- [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) - Go package that fetches parameters from AWS System Manager - Parameter Store. Stars:`62`.
- [config](https://github.com/num30/config) - configure your app using file, environment variables, or flags in two lines of code Stars:`59`.
- [configure](https://github.com/paked/configure) - Provides configuration through multiple sources, including JSON, flags and environment variables. Stars:`55`.
- [gone/jconf](https://github.com/One-com/gone/tree/master/jconf) - Modular JSON configuration. Keep your config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization. Stars:`50`.
- [go-cfg](https://github.com/dsbasko/go-cfg) - The library provides a unified way to read configuration data into a structure from various sources, such as env, flags, and configuration files (.json, .yaml, .toml, .env). Stars:`47`.
- [go-up](https://github.com/ufoscout/go-up) - A simple configuration library with recursive placeholders resolution and no magic. Stars:`43`.
- [genv](https://github.com/sakirsensoy/genv) - Read environment variables easily with dotenv support. Stars:`42`.
- [ingo](https://github.com/schachmat/ingo) - Flags persisted in an ini-like config file. Stars:`40`.
- [confiq](https://github.com/greencoda/confiq) - Structured data format to config struct decoder library for Go - supporting multiple data formats Stars:`39`.
- [mini](https://github.com/sasbury/mini) - Golang package for parsing ini-style configuration files. Stars:`37`.
- [conflate](https://github.com/the4thamigo-uk/conflate) - Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. Stars:`36`.
- [bcl](https://github.com/wkhere/bcl) - BCL is a configuration language similar to HCL. Stars:`27`.
- [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) - Go utility for loading configuration parameters from AWS SSM (Parameter Store). Stars:`22`.
- [nasermirzaei89/env](https://github.com/nasermirzaei89/env) - Simple useful package for read environment variables. Stars:`21`.
- [go-array](https://github.com/deatil/go-array) - A Go package that read or set data from map, slice or json. Stars:`20`.
- [ini](https://github.com/wlevene/ini) - INI Parser & Write Library, Unmarshal to Struct, Marshal to Json, Write File, watch file. Stars:`18`.
- [yamagiconf](https://github.com/romshark/yamagiconf) - The "safe subset" of YAML for Go configs. Stars:`18`.
- [go-ini](https://github.com/subpop/go-ini) - A Go package that marshals and unmarshals INI-files. Stars:`15`.
- [envyaml](https://github.com/yuseferi/envyaml) - Yaml with environment variables reader. it helps to have secrets as environment variable but load them configs as structured Yaml. Stars:`13`.
- [typenv](https://github.com/diegomarangoni/typenv) - Minimalistic, zero dependency, typed environment variables library. Stars:`12`.
- [go-conf](https://github.com/ThomasObenaus/go-conf) - Simple library for application configuration based on annotated structs. It supports reading the configuration from environment variables, config files and command line parameters. Stars:`11`.
- [piper](https://github.com/Yiling-J/piper) - Viper wrapper with config inheritance and key generation. Stars:`10`.
- [gonfig](https://github.com/milad-abbasi/gonfig) - Tag-based configuration parser which loads values from different providers into typesafe struct. Stars:`9`.
- [nfigure](https://github.com/muir/nfigure) - Per-library struct-tag based configuration from command lines (Posix & Go-style); environment, JSON, YAML Stars:`8`.
- [swap](https://github.com/oblq/swap) - Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). Stars:`8`.
- [env](https://github.com/syntaqx/env) - An environment utility package with support for unmarshaling into structs Stars:`7`.
- [hedzr/store](https://github.com/hedzr/store) - Extensible, high-performance configuration management library, optimized for hierarchical data. Stars:`7`.
- [GoCfg](https://github.com/Jagerente/gocfg) - Config manager with Struct Tags based contracts, custom value providers, parsers, and documentation generation. Customizable yet simple. Stars:`6`.

**[⬆ back to top](#contents)**

## Continuous Integration

_Tools for help with continuous integration._

- [drone](https://github.com/drone/drone) - Drone is a Continuous Integration platform built on Docker, written in Go. Stars:`33.0K`.
- [woodpecker](https://github.com/woodpecker-ci/woodpecker) - Woodpecker is a community fork of the Drone CI system. Stars:`5.3K`.
- [CDS](https://github.com/ovh/cds) - Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. Stars:`4.7K`.
- [muffet](https://github.com/raviqqe/muffet) - Fast website link checker in Go, see [alternatives](https://github.com/lycheeverse/lychee#features). Stars:`2.6K`.
- [abstruse](https://github.com/bleenco/abstruse) - Abstruse is a distributed CI platform. Stars:`952`.
- [goveralls](https://github.com/mattn/goveralls) - Go integration for Coveralls.io continuous code coverage tracking system. Stars:`791`.
- [gotestfmt](https://github.com/GoTestTools/gotestfmt) - go test output for humans. Stars:`573`.
- [go-test-coverage](https://github.com/vladopajic/go-test-coverage) - Tool and GitHub action which reports issues when test coverage is below set threshold. Stars:`161`.
- [overalls](https://github.com/go-playground/overalls) - Multi-Package go project coverprofile for tools like goveralls. Stars:`115`.
- [gomason](https://github.com/nikogura/gomason) - Test, Build, Sign, and Publish your go binaries from a clean workspace. Stars:`65`.
- [dot](https://github.com/opnlabs/dot) - A minimal, local first continuous integration system that uses Docker to run jobs concurrently in stages. Stars:`24`.
- [roveralls](https://github.com/LawrenceWoodman/roveralls) - Recursive coverage testing tool. Stars:`20`.
- [go-beautiful-html-coverage](https://github.com/gha-common/go-beautiful-html-coverage) - A GitHub Action to track code coverage in your pull requests, with a beautiful HTML preview, for free. Stars:`17`.
- [go-fuzz-action](https://github.com/jidicula/go-fuzz-action) - Use Go 1.18's built-in fuzz testing in GitHub Actions. Stars:`17`.
- [go-semver-release](https://github.com/s0ders/go-semver-release) - Automate the semantic versioning of Git repositories. Stars:`17`.
- [Bencher](https://bencher.dev/) - A suite of continuous benchmarking tools designed to catch performance regressions in CI.

**[⬆ back to top](#contents)**

## CSS Preprocessors

_Libraries for preprocessing CSS files._

- [go-libsass](https://github.com/wellington/go-libsass) - Go wrapper to the 100% Sass compatible libsass project. Stars:`212`.

**[⬆ back to top](#contents)**

## Data Integration Frameworks

_Frameworks for performing ELT / ETL_

- [Benthos](https://github.com/benthosdev/benthos) - A message streaming bridge between a range of protocols. Stars:`8.4K`.
- [omniparser](https://github.com/jf-tech/omniparser) - A versatile ETL library that parses text input (CSV/txt/JSON/XML/EDI/X12/EDIFACT/etc) in streaming fashion and transforms data into JSON output using data-driven schema. Stars:`1.1K`.
- [CloudQuery](http://github.com/cloudquery/cloudquery) - A high-performance ELT data integration framework with pluggable architecture.

**[⬆ back to top](#contents)**

## Data Structures and Algorithms

### Bit-packing and Compression

- [roaring](https://github.com/RoaringBitmap/roaring) - Go package implementing compressed bitsets. Stars:`2.7K`.
- [binpacker](https://github.com/zhuangsirui/binpacker) - Binary packer and unpacker helps user build custom binary stream. Stars:`232`.
- [bit](https://github.com/yourbasic/bit) - Golang set data structure with bonus bit-twiddling functions. Stars:`164`.
- [crunch](https://github.com/superwhiskers/crunch) - Go package implementing buffers for handling various datatypes easily. Stars:`99`.
- [bingo](https://github.com/iancmcc/bingo) - Fast, zero-allocation, lexicographical-order-preserving packing of native types to bytes. Stars:`46`.
- [go-ef](https://github.com/amallia/go-ef) - A Go implementation of the Elias-Fano encoding. Stars:`40`.

### Bit Sets

- [bitset](https://github.com/bits-and-blooms/bitset) - Go package implementing bitsets. Stars:`1.4K`.
- [bitmap](https://github.com/kelindar/bitmap) - Dense, zero-allocation, SIMD-enabled bitmap/bitset in Go. Stars:`347`.

### Bloom and Cuckoo Filters

- [bloom](https://github.com/bits-and-blooms/bloom) - Go package implementing Bloom filters. Stars:`2.6K`.
- [boomfilters](https://github.com/tylertreat/BoomFilters) - Probabilistic data structures for processing continuous, unbounded streams. Stars:`1.6K`.
- [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. Stars:`1.2K`.
- [cuckoo-filter](https://github.com/linvon/cuckoo-filter) - Cuckoo filter: a comprehensive cuckoo filter, which is configurable and space optimized compared with other implements, and all features mentioned in original paper are available. Stars:`301`.
- [bloom](https://github.com/zhenjl/bloom) - Bloom filters implemented in Go. Stars:`147`.
- [ring](https://github.com/TheTannerRyan/ring) - Go implementation of a high performance, thread safe bloom filter. Stars:`141`.
- [bloom](https://github.com/yourbasic/bloom) - Golang Bloom filter implementation. Stars:`87`.
- [bloomfilter](https://github.com/OldPanda/bloomfilter) - Yet another Bloomfilter implementation in Go, compatible with Java's Guava library. Stars:`18`.

### Data Structure and Algorithm Collections

- [gods](https://github.com/emirpasic/gods) - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. Stars:`17.0K`.
- [go-datastructures](https://github.com/Workiva/go-datastructures) - Collection of useful, performant, and thread-safe data structures. Stars:`7.8K`.
- [gostl](https://github.com/liyue201/gostl) - Data structure and algorithm library for go, designed to provide functions similar to C++ STL. Stars:`1.1K`.
- [algorithms](https://github.com/shady831213/algorithms) - Algorithms and data structures.CLRS study. Stars:`824`.

### Iterators

- [iter](https://github.com/disksing/iter) - Go implementation of C++ STL iterators and algorithms. Stars:`191`.
- [goterator](https://github.com/yaa110/goterator) - Iterator implementation to provide map and reduce functionalities. Stars:`17`.

### Maps

See also [Database](#database) for more complex key-value stores, and [Trees](#trees) for
additional ordered map implementations.

- [cmap](https://github.com/lrita/cmap) - a thread-safe concurrent map for go, support using `interface{}` as key and auto scale up shards. Stars:`96`.
- [goradd/maps](https://github.com/goradd/maps) - Go 1.18+ generic map interface for maps; safe maps; ordered maps; ordered, safe maps; etc. Stars:`50`.
- [dict](https://github.com/srfrog/dict) - Python-like dictionaries (dict) for Go. Stars:`47`.
- [go-shelve](https://github.com/lucmq/go-shelve) - A persistent, map-like object for the Go programming language. Supports multiple embedded key-value stores. Stars:`10`.

### Miscellaneous Data Structures and Algorithms

- [gota](https://github.com/kniren/gota) - Implementation of dataframes, series, and data wrangling methods for Go. Stars:`3.2K`.
- [hyperloglog](https://github.com/axiomhq/hyperloglog) - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. Stars:`986`.
- [go-geoindex](https://github.com/hailocab/go-geoindex) - In-memory geo index. Stars:`358`.
- [go-rquad](https://github.com/aurelien-rainone/go-rquad) - Region quadtrees with efficient point location and neighbour finding. Stars:`137`.
- [gogu](https://github.com/esimov/gogu) - A comprehensive, reusable and efficient concurrent-safe generics utility functions and data structures library. Stars:`105`.
- [go-rampart](https://github.com/francesconi/go-rampart) - Determine how intervals relate to each other. Stars:`100`.
- [go-tuple](https://github.com/barweiss/go-tuple) - Generic tuple implementation for Go 1.18+. Stars:`93`.
- [go-generics](https://github.com/bobg/go-generics) - Generic slice, map, set, iterator, and goroutine utilities. Stars:`83`.
- [fsm](https://github.com/cocoonspace/fsm) - Finite-State Machine package. Stars:`80`.
- [hide](https://github.com/emvi/hide) - ID type with marshalling to/from hash to prevent sending IDs to clients. Stars:`70`.
- [count-min-log](https://github.com/seiflotfy/count-min-log) - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). Stars:`67`.
- [concurrent-writer](https://github.com/free/concurrent-writer) - Highly concurrent drop-in replacement for `bufio.Writer`. Stars:`57`.
- [genfuncs](https://github.com/nwillc/genfuncs) - Go 1.18+ generics package inspired by Kotlin's Sequence and Map. Stars:`52`.
- [go18ds](https://github.com/daichi-m/go18ds) - Go Data Structures using Go 1.18 generics. Stars:`46`.
- [quadtree](https://github.com/s0rg/quadtree) - Generic, zero-alloc, 100%-test covered quadtree. Stars:`38`.
- [slices](https://github.com/twharmon/slices) - Pure, generic functions for slices. Stars:`22`.
- [gofal](https://github.com/xxjwxc/gofal) - fractional api for Go. Stars:`19`.

### Nullable Types

- [nan](https://github.com/kak-tus/nan) - Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers. Stars:`87`.
- [typ](https://github.com/gurukami/typ) - Null Types, Safe primitive type conversion and fetching value from complex structures. Stars:`46`.
- [null](https://github.com/emvi/null) - Nullable Go types that can be marshalled/unmarshalled to/from JSON. Stars:`38`.

### Queues

- [hatchet](https://github.com/hatchet-dev/hatchet) - Distributed, Fault-tolerant task queue. Stars:`5.9K`.
- [deque](https://github.com/gammazero/deque) - Fast ring-buffer deque (double-ended queue). Stars:`694`.
- [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) - Concurrent FIFO queue. Stars:`431`.
- [queue](https://github.com/adrianbrad/queue) - Multiple thread-safe, generic queue implementations for Go. Stars:`312`.
- [deque](https://github.com/edwingeng/deque) - A highly optimized double-ended queue. Stars:`201`.
- [memlog](https://github.com/embano1/memlog) - An easy to use, lightweight, thread-safe and append-only in-memory data structure inspired by Apache Kafka. Stars:`133`.
- [dqueue](https://github.com/vodolaz095/dqueue) - Simple, in memory, zero dependency and battle tested, thread-safe deferred queue.

### Sets

- [golang-set](https://github.com/deckarep/golang-set) - Thread-Safe and Non-Thread-Safe high-performance sets for Go. Stars:`4.6K`.
- [goset](https://github.com/zoumo/goset) - A useful Set collection implementation for Go. Stars:`52`.
- [set](https://github.com/StudioSol/set) - Simple set data structure implementation in Go using LinkedHashMap. Stars:`29`.
- [dsu](https://github.com/ihebu/dsu) - Disjoint Set data structure implementation in Go. Stars:`15`.

### Text Analysis

- [bleve](https://github.com/blevesearch/bleve) - Modern text indexing library for go. Stars:`10.5K`.
- [trie](https://github.com/derekparker/trie) - Trie implementation in Go. Stars:`776`.
- [go-edlib](https://github.com/hbollon/go-edlib) - Go string comparison and edit distance algorithms library (Levenshtein, LCS, Hamming, Damerau levenshtein, Jaro-Winkler, etc.) compatible with Unicode. Stars:`522`.
- [levenshtein](https://github.com/agnivade/levenshtein) - Implementation to calculate levenshtein distance in Go. Stars:`424`.
- [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) - Go implementation of Adaptive Radix Tree. Stars:`395`.
- [levenshtein](https://github.com/agext/levenshtein) - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. Stars:`89`.
- [ptrie](https://github.com/viant/ptrie) - An implementation of prefix tree. Stars:`43`.
- [mspm](https://github.com/BlackRabbitt/mspm) - Multi-String Pattern Matching Algorithm for information retrieval. Stars:`27`.
- [parsefields](https://github.com/MonaxGT/parsefields) - Tools for parse JSON-like logs for collecting unique fields and events. Stars:`7`.

### Trees

- [skiplist](https://github.com/MauriceGit/skiplist) - Very fast Go Skiplist implementation. Stars:`287`.
- [skiplist](https://github.com/gansidui/skiplist) - Skiplist implementation in Go. Stars:`84`.
- [treemap](https://github.com/igrmk/treemap) - Generic key-sorted map using a red-black tree under the hood. Stars:`64`.
- [merkle](https://github.com/bobg/merkle) - Space-efficient computation of Merkle root hashes and inclusion proofs. Stars:`19`.
- [hashsplit](http://github.com/bobg/hashsplit) - Split byte streams into chunks, and arrange chunks into trees, with boundaries determined by content, not position.
- [treap](https://github.com/perdata/treap) - Persistent, fast ordered map using tree heaps.

### Pipes

- [pipeline](https://github.com/hyfather/pipeline) - An implementation of pipelines with fan-in and fan-out. Stars:`58`.
- [ordered-concurrently](https://github.com/tejzpr/ordered-concurrently) - Go module that processes work concurrently and returns output in a channel in the order of input. Stars:`41`.
- [parapipe](https://github.com/nazar256/parapipe) - FIFO Pipeline which parallels execution on each stage while maintaining the order of messages and results. Stars:`37`.
- [pipelines](https://github.com/nxdir-s/pipelines) - Generic pipeline functions for concurrent processing. Stars:`11`.

**[⬆ back to top](#contents)**

## Database

### Caches

_Data stores with expiring records, in-memory distributed data stores, or in-memory subsets of file-based databases._

- [groupcache](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. Stars:`13.2K`.
- [BigCache](https://github.com/allegro/bigcache) - Efficient key/value cache for gigabytes of data. Stars:`7.9K`.
- [ristretto](https://github.com/dgraph-io/ristretto) -  A high performance memory-bound Go cache. Stars:`6.3K`.
- [gocache](https://github.com/eko/gocache) - A complete Go cache library with multiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. Stars:`2.7K`.
- [GCache](https://github.com/bluele/gcache) - Cache library with support for expirable Cache, LFU, LRU and ARC. Stars:`2.7K`.
- [otter](https://github.com/maypok86/otter) - A high performance lockless cache for Go. Many times faster than Ristretto and friends. Stars:`2.3K`.
- [fastcache](https://github.com/VictoriaMetrics/fastcache) - fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. Stars:`2.2K`.
- [cache2go](https://github.com/muesli/cache2go) - In-memory key:value cache which supports automatic invalidation based on timeouts. Stars:`2.2K`.
- [sturdyc](https://github.com/viccon/sturdyc) - A caching library with advanced concurrency features designed to make I/O heavy applications robust and highly performant. Stars:`1.2K`.
- [ttlcache](https://github.com/jellydator/ttlcache) - An in-memory cache with item expiration and generics. Stars:`1.1K`.
- [EchoVault](https://github.com/EchoVault/EchoVault) - Embeddable Distributed in-memory data store compatible with Redis clients. Stars:`496`.
- [cachego](https://github.com/faabiosr/cachego) - Golang Cache component for multiple drivers. Stars:`373`.
- [theine](https://github.com/Yiling-J/theine-go) - High performance, near optimal in-memory cache with proactive TTL expiration and generics. Stars:`338`.
- [go-freelru](https://github.com/elastic/go-freelru) A GC-less, fast and generic LRU hashmap library with optional locking, sharding, eviction and expiration. Stars:`246`.
- [pocache](https://github.com/naughtygopher/pocache) - Pocache is a minimal cache package which focuses on a preemptive optimistic caching strategy. Stars:`226`.
- [cache](https://github.com/akyoto/cache) - In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. Stars:`194`.
- [bcache](https://github.com/iwanbk/bcache) - Eventually consistent distributed in-memory cache Go library. Stars:`160`.
- [go-cache](https://github.com/viney-shih/go-cache) - A flexible multi-layer Go caching library to deal with in-memory and shared cache by adopting Cache-Aside pattern. Stars:`157`.
- [imcache](https://github.com/erni27/imcache) - A generic in-memory cache Go library. It supports expiration, sliding expiration, max entries limit, eviction callbacks and sharding. Stars:`125`.
- [go-mcache](https://github.com/OrlovEvgeny/go-mcache) - Fast in-memory key:value store/cache library. Pointer caches. Stars:`98`.
- [timedmap](https://github.com/zekroTJA/timedmap) - Map with expiring key-value pairs. Stars:`75`.
- [couchcache](https://github.com/codingsince1985/couchcache) - RESTful caching micro-service backed by Couchbase server. Stars:`67`.
- [2q](https://github.com/floatdrop/2q) - 2Q in-memory cache implementation. Stars:`47`.
- [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) - BigCache with clustering support and individual item expiration. Stars:`45`.
- [icache](https://github.com/mdaliyan/icache) - A High Performance, Generic, thread-safe, zero-dependency cache package. Stars:`21`.
- [coherence-go-client](https://github.com/oracle/coherence-go-client) - Full implementation of Oracle Coherence cache API for Go applications using gRPC as network transport. Stars:`13`.
- [gdcache](https://github.com/ulovecode/gdcache) - A pure non-intrusive cache library implemented by golang, you can use it to implement your own distributed cache. Stars:`13`.
- [nscache](https://github.com/no-src/nscache) - A Go caching framework that supports multiple data source drivers. Stars:`12`.
- [gocache](https://github.com/yuseferi/gocache) - A data race free Go ache library with high performance and auto pruge functionality Stars:`12`.
- [ttlcache](https://github.com/cheshir/ttlcache) - In-memory key value storage with TTL for each record. Stars:`10`.

### Databases Implemented in Go

- [prometheus](https://github.com/prometheus/prometheus) - Monitoring system and time series database. Stars:`59.7K`.
- [tidb](https://github.com/pingcap/tidb) - TiDB is a distributed SQL database. Inspired by the design of Google F1. Stars:`38.8K`.
- [Milvus](https://github.com/milvus-io/milvus) - Milvus is a vector database for embedding management, analytics and search. Stars:`36.3K`.
- [cockroach](https://github.com/cockroachdb/cockroach) - Scalable, Geo-Replicated, Transactional Datastore. Stars:`31.1K`.
- [influxdb](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics. Stars:`30.4K`.
- [dgraph](https://github.com/dgraph-io/dgraph) - Scalable, Distributed, Low Latency, High Throughput Graph Database. Stars:`21.1K`.
- [dolt](https://github.com/dolthub/dolt) - Dolt – It's Git for Data. Stars:`18.9K`.
- [rqlite](https://github.com/rqlite/rqlite) - The lightweight, distributed, relational database built on SQLite. Stars:`16.8K`.
- [badger](https://github.com/dgraph-io/badger) - Fast key-value store in Go. Stars:`14.9K`.
- [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) - fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. Stars:`14.7K`.
- [DiceDB](https://github.com/DiceDB/dice) - An open-source, fast, reactive, in-memory database optimized for modern hardware. Higher throughput and lower median latencies, making it ideal for modern workloads. Stars:`10.6K`.
- [bbolt](https://github.com/etcd-io/bbolt) - An embedded key/value database for Go. Stars:`9.0K`.
- [immudb](https://github.com/codenotary/immudb) - immudb is a lightweight, high-speed immutable database for systems and applications written in Go. Stars:`8.8K`.
- [goleveldb](https://github.com/syndtr/goleveldb) - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. Stars:`6.3K`.
- [pebble](https://github.com/cockroachdb/pebble) - RocksDB/LevelDB inspired key-value database in Go. Stars:`5.4K`.
- [rosedb](https://github.com/roseduan/rosedb) - An embedded k-v database based on LSM+WAL, supports string, list, hash, set, zset. Stars:`4.8K`.
- [buntdb](https://github.com/tidwall/buntdb) - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. Stars:`4.8K`.
- [ledisdb](https://github.com/siddontang/ledisdb) - Ledisdb is a high performance NoSQL like Redis based on LevelDB. Stars:`4.1K`.
- [redka](https://github.com/nalgeon/redka) - Redis re-implemented with SQLite. Stars:`4.1K`.
- [godis](https://github.com/hdt3213/godis) - A Golang implemented high-performance Redis server and cluster. Stars:`3.7K`.
- [nutsdb](https://github.com/xujiajun/nutsdb) - Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as list, set, sorted set. Stars:`3.5K`.
- [LinDB](https://github.com/lindb/lindb) - LinDB is a scalable, high performance, high availability distributed time series database. Stars:`3.0K`.
- [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang. Stars:`2.7K`.
- [lotusdb](https://github.com/flower-corp/lotusdb) - Fast k/v database compatible with lsm and b+tree. Stars:`2.2K`.
- [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) - CovenantSQL is a SQL database on blockchain. Stars:`1.5K`.
- [column](https://github.com/kelindar/column) - High-performance, columnar, embeddable in-memory store with bitmap indexing and transactions. Stars:`1.5K`.
- [diskv](https://github.com/peterbourgon/diskv) - Home-grown disk-backed key-value store. Stars:`1.4K`.
- [pogreb](https://github.com/akrylysov/pogreb) - Embedded key-value store for read-heavy workloads. Stars:`1.3K`.
- [Databunker](https://github.com/paranoidguy/databunker) - Personally identifiable information (PII) storage service built to comply with GDPR and CCPA. Stars:`1.3K`.
- [objectbox-go](https://github.com/objectbox/objectbox-go) - High-performance embedded Object Database (NoSQL) with Go API. Stars:`1.2K`.
- [eliasdb](https://github.com/krotik/eliasdb) - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. Stars:`1.0K`.
- [moss](https://github.com/couchbase/moss) - Moss is a simple LSM key-value storage engine written in 100% Go. Stars:`1.0K`.
- [clover](https://github.com/ostafen/clover) - A lightweight document-oriented NoSQL database written in pure Golang. Stars:`767`.
- [levigo](https://github.com/jmhodges/levigo) - Levigo is a Go wrapper for LevelDB. Stars:`419`.
- [pudge](https://github.com/recoilme/pudge) - Fast and simple key/value store written using Go's standard library. Stars:`373`.
- [Vasto](https://github.com/chrislusf/vasto) - A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. Stars:`263`.
- [piladb](https://github.com/fern4lvarez/piladb) - Lightweight RESTful database engine based on stack data structures. Stars:`207`.
- [libradb](https://github.com/amit-davidson/LibraDB) - LibraDB is a simple database with less than 1000 lines of code for learning. Stars:`189`.
- [unitdb](https://github.com/unit-io/unitdb) - Fast timeseries database for IoT, realtime messaging applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application. Stars:`124`.
- [hare](https://github.com/jameycribbs/hare) - A simple database management system that stores each table as a text file of line-delimited JSON. Stars:`97`.
- [Coffer](https://github.com/claygod/coffer) - Simple ACID key-value database that supports transactions. Stars:`40`.
- [rotom](https://github.com/xgzlucario/rotom) - A tiny Redis server built with Golang, compatible with RESP protocols. Stars:`39`.
- [tempdb](https://github.com/rafaeljesus/tempdb) - Key-value store for temporary items. Stars:`19`.
- [Bitcask](https://git.mills.io/prologic/bitcask) - Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).

### Database Schema Migration

- [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library. Stars:`17.1K`.
- [bytebase](https://github.com/bytebase/bytebase) - Safe database schema change and version control for DevOps teams. Stars:`12.8K`.
- [goose](https://github.com/pressly/goose) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. Stars:`8.9K`.
- [atlas](https://github.com/ariga/atlas) - A Database Toolkit. A CLI designed to help companies better work with their data. Stars:`7.1K`.
- [dbmate](https://github.com/amacneil/dbmate) - A lightweight, framework-agnostic database migration tool. Stars:`6.2K`.
- [sql-migrate](https://github.com/rubenv/sql-migrate) - Database migration tool. Allows embedding migrations into the application using go-bindata. Stars:`3.4K`.
- [soda](https://github.com/gobuffalo/pop/tree/master/soda) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. Stars:`1.5K`.
- [skeema](https://github.com/skeema/skeema) - Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. Stars:`1.3K`.
- [gormigrate](https://github.com/go-gormigrate/gormigrate) - Database schema migration helper for Gorm ORM. Stars:`1.1K`.
- [goavro](https://github.com/linkedin/goavro) - A Go package that encodes and decodes Avro data. Stars:`1.0K`.
- [migrator](https://github.com/lopezator/migrator) - Dead simple Go database migration library. Stars:`178`.
- [darwin](https://github.com/GuiaBolso/darwin) - Database schema evolution library for Go. Stars:`149`.
- [sqlize](https://github.com/sunary/sqlize) - Database migration generator. Allows generate sql migration from model and existing sql by differ them. Stars:`121`.
- [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) - A Go package to help write migrations with go-pg/pg. Stars:`86`.
- [avro](https://github.com/khezen/avro) - Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. Stars:`47`.
- [schema](https://github.com/adlio/schema) - Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. Stars:`42`.
- [go-fixtures](https://github.com/RichardKnop/go-fixtures) - Django style fixtures for Golang's excellent built-in database/sql library. Stars:`30`.
- [migrator](https://github.com/larapulse/migrator) - MySQL database migrator designed to run migrations to your features and manage database schema update with intuitive go code. Stars:`25`.
- [libschema](https://github.com/muir/libschema) - Define your migrations separately in each library. Migrations for open source libraries. MySQL & PostgreSQL. Stars:`17`.
- [gorm-seeder](https://github.com/Kachit/gorm-seeder) - Simple database seeder for Gorm ORM. Stars:`16`.
- [go-pg-migrate](https://github.com/lawzava/go-pg-migrate) - CLI-friendly package for go-pg migrations management. Stars:`11`.
- [godfish](https://github.com/rafaelespinoza/godfish) - Database migration manager, works with native query language. Support for cassandra, mysql, postgres, sqlite3. Stars:`8`.

### Database Tools

- [vitess](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. Stars:`19.9K`.
- [pgweb](https://github.com/sosedoff/pgweb) - Web-based PostgreSQL database browser. Stars:`9.0K`.
- [go-mysql](https://github.com/siddontang/go-mysql) - Go toolset to handle MySQL protocol and replication. Stars:`4.8K`.
- [pREST](https://github.com/prest/prest) - Simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new. Stars:`4.4K`.
- [chproxy](https://github.com/Vertamedia/chproxy) - HTTP proxy for ClickHouse database. Stars:`1.4K`.
- [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable) - Advanced scheduling for PostgreSQL. Stars:`1.2K`.
- [rdb](https://github.com/HDT3213/rdb) - Redis RDB file parser for secondary development and memory analysis. Stars:`555`.
- [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) - Collects small inserts and sends big requests to ClickHouse servers. Stars:`502`.
- [wescale](https://github.com/wesql/wescale) - WeScale is a database proxy designed to enhance the scalability, performance, security, and resilience of your applications. Stars:`306`.
- [gatewayd](https://github.com/gatewayd-io/gatewayd) - Cloud-native database gateway and framework for building data-driven applications. Like API gateways, for databases. Stars:`264`.
- [onedump](https://github.com/liweiyi88/onedump) - Database backup from different drivers to different destinations with one command and configuration. Stars:`207`.
- [octillery](https://github.com/knocknote/octillery) - Go package for sharding databases ( Supports every ORM or raw SQL ). Stars:`198`.
- [dbbench](https://github.com/sj14/dbbench) - Database benchmarking tool with support for several databases and scripts. Stars:`110`.
- [gorm-multitenancy](https://github.com/bartventer/gorm-multitenancy) - Multi-tenancy support for GORM managed databases. Stars:`62`.
- [dg](https://github.com/codingconcepts/dg) - A fast data generator that produces CSV files from generated relational data. Stars:`35`.
- [prep](https://github.com/hexdigest/prep) - Use prepared SQL statements without changing your code. Stars:`35`.
- [database-gateway](https://github.com/kazhuravlev/database-gateway) - Running SQL in production with ACLs, logs, and shared links. Stars:`28`.
- [rwdb](https://github.com/andizzle/rwdb) - rwdb provides read replica capability for multiple database servers setup. Stars:`20`.
- [hasql](https://golang.yandex/hasql) - Library for accessing multi-host SQL database installations.

### SQL Query Builders

_Libraries for building and using SQL._

- [sqlc](https://github.com/kyleconroy/sqlc) - Generate type-safe code from SQL. Stars:`15.6K`.
- [Squirrel](https://github.com/Masterminds/squirrel) - Go library that helps you build SQL queries. Stars:`7.5K`.
- [xo](https://github.com/knq/xo) - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. Stars:`3.8K`.
- [jet](https://github.com/go-jet/jet) - Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. Stars:`3.3K`.
- [goqu](https://github.com/doug-martin/goqu) - Idiomatic SQL builder and query library. Stars:`2.5K`.
- [gendry](https://github.com/didi/gendry) - Non-invasive SQL builder and powerful data binder. Stars:`1.6K`.
- [Dotsql](https://github.com/gchaincl/dotsql) - Go library that helps you keep sql files in one place and use them with ease. Stars:`749`.
- [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) - Powerful data retrieval methods as well as DB-agnostic query building capabilities. Stars:`654`.
- [sqlingo](https://github.com/lqs/sqlingo) - A lightweight DSL to build SQL in Go. Stars:`430`.
- [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go. Stars:`413`.
- [sqrl](https://github.com/elgris/sqrl) - SQL query builder, fork of Squirrel with improved performance. Stars:`281`.
- [sq](https://github.com/bokwoon95/go-structured-query) - Type-safe SQL builder and struct mapper for Go. Stars:`201`.
- [sqlf](https://github.com/leporo/sqlf) - Fast SQL query builder. Stars:`181`.
- [buildsqlx](https://github.com/arthurkushman/buildsqlx) - Go database query builder library for PostgreSQL. Stars:`181`.
- [bqb](https://github.com/nullism/bqb) - Lightweight and easy to learn query builder. Stars:`174`.
- [igor](https://github.com/galeone/igor) - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. Stars:`126`.
- [builq](https://github.com/cristalhq/builq) - Easily build SQL queries in Go. Stars:`94`.
- [godbal](https://github.com/xujiajun/godbal) - Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. Stars:`59`.
- [gosql](https://github.com/twharmon/gosql) - SQL Query builder with better null values support. Stars:`36`.
- [qry](https://github.com/HnH/qry) - Tool that generates constants from files with raw SQL queries. Stars:`35`.
- [Hotcoal](https://github.com/motrboat/hotcoal) - Secure your handcrafted SQL against injection. Stars:`24`.
- [ormlite](https://github.com/pupizoid/ormlite) - Lightweight package containing some ORM-like features and helpers for sqlite databases. Stars:`16`.
- [sg](https://github.com/go-the-way/sg) - A SQL Gen for generating standard SQLs(supports: CRUD) written in Go. Stars:`7`.
- [patcher](https://github.com/Jacobbrewer1/patcher) - Powerful SQL Query builder that automatically generates SQL queries from structs. Stars:`6`.
- [Squalus](https://gitlab.com/qosenergy/squalus) - Thin layer over the Go SQL package that makes it easier to perform queries.

**[⬆ back to top](#contents)**

## Database Drivers

### Interfaces to Multiple Backends

- [cayley](https://github.com/google/cayley) - Graph database with support for multiple backends. Stars:`14.9K`.
- [gokv](https://github.com/philippgille/gokv) - Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). Stars:`804`.
- [go-transaction-manager](https://github.com/avito-tech/go-transaction-manager) - Transaction manager with multiple adapters (sql, sqlx, gorm, mongo, ...) controls transaction boundaries. Stars:`321`.
- [dsc](https://github.com/viant/dsc) - Datastore connectivity for SQL, NoSQL, structured files. Stars:`34`.
- [dynamo](https://github.com/fogfish/dynamo) - A simple key-value abstraction to store algebraic and linked-data data types at AWS storage services: AWS DynamoDB and AWS S3. Stars:`21`.

### Relational Database Drivers

- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go. Stars:`15.0K`.
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver supporting features beyond those exposed by database/sql. Stars:`12.2K`.
- [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql. Stars:`9.6K`.
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that uses database/sql. Stars:`8.6K`.
- [go-mssqldb](https://github.com/denisenkom/go-mssqldb) - Microsoft MSSQL driver for Go. Stars:`1.9K`.
- [go-sqlite3](https://github.com/ncruces/go-sqlite3) - This Go module is compatible with the database/sql driver. It allows embedding SQLite into your application, provides direct access to its C API, supports SQLite VFS, and also includes a GORM driver. Stars:`749`.
- [sqlhooks](https://github.com/qustavo/sqlhooks) - Attach hooks to any database/sql driver. Stars:`658`.
- [go-oci8](https://github.com/mattn/go-oci8) - Oracle driver for go that uses database/sql. Stars:`635`.
- [godror](https://github.com/godror/godror) - Oracle driver for Go, using the ODPI-C driver. Stars:`559`.
- [Sqinn-Go](https://github.com/cvilsmeier/sqinn-go) - SQLite with pure Go. Stars:`452`.
- [KSQL](https://github.com/VinGarcia/ksql) - A Simple and Powerful Golang SQL Library Stars:`340`.
- [surrealdb.go](https://github.com/surrealdb/surrealdb.go) - SurrealDB Driver for Go. Stars:`282`.
- [firebirdsql](https://github.com/nakagami/firebirdsql) - Firebird RDBMS SQL driver for Go. Stars:`245`.
- [ydb-go-sdk](https://github.com/ydb-platform/ydb-go-sdk) - native and database/sql driver YDB (Yandex Database) Stars:`167`.
- [go-rqlite](https://github.com/rqlite/gorqlite) - A Go client for rqlite, providing easy-to-use abstractions for working with the rqlite API. Stars:`167`.
- [go-adodb](https://github.com/mattn/go-adodb) - Microsoft ActiveX Object DataBase driver for go that uses database/sql. Stars:`149`.
- [avatica](https://github.com/apache/calcite-avatica-go) - Apache Avatica/Phoenix SQL driver for database/sql. Stars:`122`.
- [gofreetds](https://github.com/minus5/gofreetds) - Microsoft MSSQL driver. Go wrapper over [FreeTDS](https://www.freetds.org). Stars:`113`.
- [bgc](https://github.com/viant/bgc) - Datastore Connectivity for BigQuery for go. Stars:`21`.
- [pig](https://github.com/alexeyco/pig) - Simple [pgx](https://github.com/jackc/pgx) wrapper to execute and [scan](https://github.com/georgysavva/scany) query results easily. Stars:`17`.
- [sqlite](https://pkg.go.dev/modernc.org/sqlite) - Package sqlite is a sql/database driver using a CGo-free port of the C SQLite3 library.

### NoSQL Database Drivers

- [redis](https://github.com/redis/go-redis) - Redis client for Golang. Stars:`21.1K`.
- [redigo](https://github.com/gomodule/redigo) - Redigo is a Go client for the Redis database. Stars:`9.8K`.
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver for the Go language. Stars:`8.4K`.
- [mgo](https://github.com/globalsign/mgo) - (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. Stars:`2.0K`.
- [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language. Stars:`1.8K`.
- [gorethink](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB. Stars:`1.7K`.
- [qmgo](https://github.com/qiniu/qmgo) - The MongoDB driver for Go. It‘s based on official MongoDB driver but easier to use like Mgo. Stars:`1.3K`.
- [mgm](https://github.com/kamva/mgm) - MongoDB model-based ODM for Go (based on official MongoDB driver). Stars:`762`.
- [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) - Aerospike client in Go language. Stars:`449`.
- [redeo](https://github.com/bsm/redeo) - Redis-protocol compatible TCP servers/services. Stars:`446`.
- [neoism](https://github.com/jmcvetta/neoism) - Neo4j client for Golang. Stars:`390`.
- [gocb](https://github.com/couchbase/gocb) - Official Couchbase Go SDK. Stars:`373`.
- [go-rejson](https://github.com/nitishm/go-rejson) - Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. Stars:`343`.
- [Kivik](https://github.com/go-kivik/kivik) - Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. Stars:`325`.
- [go-couchbase](https://github.com/couchbase/go-couchbase) - Couchbase client in Go. Stars:`323`.
- [go-mongox](https://github.com/chenmingyong0423/go-mongox) - A Go Mongo library based on the official driver, featuring streamlined document operations, generic binding of structs to collections, built-in CRUD, aggregation, automated field updates, struct validation, hooks, and plugin-based programming. Stars:`202`.
- [godis](https://github.com/piaohao/godis) - redis client implement by golang, inspired by jedis. Stars:`113`.
- [arangolite](https://github.com/solher/arangolite) - Lightweight golang driver for ArangoDB. Stars:`72`.
- [go-pilosa](https://github.com/pilosa/go-pilosa) - Go client library for Pilosa. Stars:`57`.
- [forestdb](https://github.com/couchbase/goforestdb) - Go bindings for ForestDB. Stars:`36`.
- [goriak](https://github.com/zegl/goriak) - Go language driver for Riak KV. Stars:`30`.
- [neo4j](https://github.com/cihangir/neo4j) - Neo4j Rest API Bindings for Golang. Stars:`29`.
- [gocosmos](https://github.com/btnguyen2k/gocosmos) - REST client and standard `database/sql` driver for Azure Cosmos DB. Stars:`22`.
- [gomemcached](https://github.com/aliexpressru/gomemcached) - A binary Memcached client for Go with support for sharding using consistent hashing, along with SASL. Stars:`20`.
- [xredis](https://github.com/shomali11/xredis) - Typesafe, customizable, clean & easy to use Redis client. Stars:`19`.
- [godscache](https://github.com/defcronyke/godscache) - A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. Stars:`11`.
- [asc](https://github.com/viant/asc) - Datastore Connectivity for Aerospike for go. Stars:`11`.
- [gocql](https://gocql.github.io) - Go language driver for Apache Cassandra.
- [rueidis](http://github.com/rueian/rueidis) - Fast Redis RESP3 client with auto pipelining and server-assisted client side caching.

### Search and Analytic Databases

- [elastic](https://github.com/olivere/elastic) - Elasticsearch client for Go. Stars:`7.5K`.
- [go-elasticsearch](https://github.com/elastic/go-elasticsearch) - Official Elasticsearch client for Go. Stars:`5.9K`.
- [clickhouse-go](https://github.com/ClickHouse/clickhouse-go/) - ClickHouse SQL client for Go with a `database/sql` compatibility. Stars:`3.1K`.
- [elasticsql](https://github.com/cch123/elasticsql) - Convert sql to elasticsearch dsl in Go. Stars:`1.2K`.
- [zoekt](https://github.com/sourcegraph/zoekt) - Fast trigram based code search. Stars:`1.0K`.
- [elastigo](https://github.com/mattbaird/elastigo) - Elasticsearch client library. Stars:`941`.
- [skizze](https://github.com/seiflotfy/skizze) - probabilistic data-structures service and storage. Stars:`92`.
- [effdsl](https://github.com/sdqri/effdsl) - Elasticsearch query builder for Go. Stars:`33`.
- [goes](https://github.com/OwnLocal/goes) - Library to interact with Elasticsearch. Stars:`31`.

**[⬆ back to top](#contents)**

## Date and Time

_Libraries for working with dates and times._

- [carbon](https://github.com/dromara/carbon) - A simple, semantic and developer-friendly time package for golang. Stars:`5.1K`.
- [now](https://github.com/jinzhu/now) - Now is a time toolkit for golang. Stars:`4.6K`.
- [dateparse](https://github.com/araddon/dateparse) - Parse date's without knowing format in advance. Stars:`2.1K`.
- [carbon](https://github.com/uniplaces/carbon) - Simple Time extension with a lot of util methods, ported from PHP Carbon library. Stars:`783`.
- [durafmt](https://github.com/hako/durafmt) - Time duration formatting library for Go. Stars:`504`.
- [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) - The implementation of the Persian (Solar Hijri) Calendar in Go (golang). Stars:`227`.
- [gostradamus](https://github.com/bykof/gostradamus) - A Go package for working with dates. Stars:`210`.
- [timeutil](https://github.com/leekchan/timeutil) - Useful extensions (Timedelta, Strftime, ...) to the golang's time package. Stars:`192`.
- [go-sunrise](https://github.com/nathan-osman/go-sunrise) - Calculate the sunrise and sunset times for a given location. Stars:`169`.
- [iso8601](https://github.com/relvacode/iso8601) - Efficiently parse ISO8601 date-times without regex. Stars:`151`.
- [date](https://github.com/rickb777/date) - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. Stars:`140`.
- [go-str2duration](https://github.com/xhit/go-str2duration) - Convert string to duration. Support time.Duration returned string and more. Stars:`116`.
- [timespan](https://github.com/SaidinWoT/timespan) - For interacting with intervals of time, defined as a start time and a duration. Stars:`83`.
- [feiertage](https://github.com/wlbr/feiertage) - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... Stars:`49`.
- [go-anytime](https://github.com/ijt/go-anytime) - Parse dates/times like "next dec 22nd at 3pm" and ranges like "from today until next thursday" without knowing the format in advance. Stars:`30`.
- [kair](https://github.com/GuilhermeCaruso/kair) - Date and Time - Golang Formatting Library. Stars:`24`.
- [cronrange](https://github.com/1set/cronrange) - Parses Cron-style time range expressions, checks if the given time is within any ranges. Stars:`17`.
- [approx](https://github.com/goschtalt/approx) - A Duration extension supporting parsing/printing durations in days, weeks and years. Stars:`15`.
- [strftime](https://github.com/awoodbeck/strftime) - C99-compatible strftime formatter. Stars:`13`.
- [tuesday](https://github.com/osteele/tuesday) - Ruby-compatible Strftime function. Stars:`12`.
- [go-week](https://github.com/stoewer/go-week) - An efficient package to work with ISO8601 week dates. Stars:`10`.
- [go-datebin](https://github.com/deatil/go-datebin) - A simple datetime parse pkg. Stars:`4`.

**[⬆ back to top](#contents)**

## Distributed Systems

_Packages that help with building Distributed Systems._

- [go-zero](https://github.com/tal-tech/go-zero) - A web and rpc framework. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity. Stars:`31.5K`.
- [go-kit](https://github.com/go-kit/kit) - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. Stars:`27.2K`.
- [Kratos](https://github.com/go-kratos/kratos) - A modular-designed and easy-to-use microservices framework in Go. Stars:`24.7K`.
- [go-micro](https://github.com/micro/go-micro) - A distributed systems development framework. Stars:`22.4K`.
- [grpc-go](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC. Stars:`22.1K`.
- [NATS](https://github.com/nats-io/nats-server) - NATS is a simple, secure, and Stars:`17.6K`.
- [micro](https://github.com/micro/micro) - A distributed systems runtime for the cloud and beyond. Stars:`12.3K`.
- [Kitex](https://github.com/cloudwego/kitex) - A high-performance and strong-extensibility Golang RPC framework that helps developers build microservices. If the performance and extensibility are the main concerns when you develop microservices, Kitex can be a good choice. Stars:`7.6K`.
- [lura](https://github.com/luraproject/lura) - Ultra performant API Gateway framework with middlewares. Stars:`6.6K`.
- [dragonboat](https://github.com/lni/dragonboat) - A feature complete and high performance multi-group Raft library in Go. Stars:`5.2K`.
- [evans](https://github.com/ktr0731/evans) - Evans: more expressive universal gRPC client. Stars:`4.4K`.
- [emitter-io](https://github.com/emitter-io/emitter) - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. Stars:`4.0K`.
- [gleam](https://github.com/chrislusf/gleam) - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. Stars:`3.5K`.
- [glow](https://github.com/chrislusf/glow) - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. Stars:`3.2K`.
- [Dragonfly](https://github.com/dragonflyoss/Dragonfly2) - Provide efficient, stable and secure file distribution and image acceleration based on p2p technology to be the best practice and standard solution in cloud native architectures. Stars:`2.7K`.
- [liftbridge](https://github.com/liftbridge-io/liftbridge) - Lightweight, fault-tolerant message streams for NATS. Stars:`2.6K`.
- [go-eagle](https://github.com/go-eagle/eagle) - A Go framework for the API or Microservice with handy scaffolding tools. Stars:`2.3K`.
- [mochi mqtt](https://github.com/mochi-co/mqtt) - Fully spec compliant, embeddable high-performance MQTT v5/v3 broker for IoT, smarthome, and pubsub. Stars:`1.6K`.
- [hprose](https://github.com/hprose/hprose-golang) - Very newbility RPC Library, support 25+ languages now. Stars:`1.3K`.
- [go-doudou](https://github.com/unionj-cloud/go-doudou) - A gossip protocol and OpenAPI 3.0 spec based decentralized microservice framework. Built-in go-doudou cli focusing on low-code and rapid dev can power up your productivity. Stars:`1.2K`.
- [arpc](https://github.com/lesismal/arpc) - More effective network communication, support two-way-calling, notify, broadcast. Stars:`1.1K`.
- [K8gb](https://github.com/k8gb-io/k8gb) - A cloud native Kubernetes Global Balancer. Stars:`1.1K`.
- [go-health](https://github.com/InVisionApp/go-health) - Library for enabling asynchronous dependency health checks in your service. Stars:`751`.
- [consistent](https://github.com/buraksezer/consistent) - Consistent hashing with bounded loads. Stars:`741`.
- [gorpc](https://github.com/valyala/gorpc) - Simple, fast and scalable RPC library for high load. Stars:`708`.
- [go-sundheit](https://github.com/AppsFlyer/go-sundheit) - A library built to provide support for defining async service health checks for golang services. Stars:`557`.
- [digota](https://github.com/digota/digota) - grpc ecommerce microservice. Stars:`516`.
- [go-jump](https://github.com/dgryski/go-jump) - Port of Google's "Jump" Consistent Hash function. Stars:`387`.
- [jsonrpc](https://github.com/ybbus/jsonrpc) - JSON-RPC 2.0 HTTP client implementation. Stars:`356`.
- [dht](https://github.com/anacrolix/dht) - BitTorrent Kademlia DHT implementation. Stars:`333`.
- [jsonrpc](https://github.com/osamingo/jsonrpc) - The jsonrpc package helps implement of JSON-RPC 2.0. Stars:`189`.
- [doublejump](https://github.com/edwingeng/doublejump) - A revamped Google's jump consistent hash. Stars:`107`.
- [dot](https://github.com/dotchain/dot/) - distributed sync using operational transformation/OT. Stars:`86`.
- [celeriac](https://github.com/svcavallar/celeriac.v1) - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. Stars:`73`.
- [capillaries](https://github.com/capillariesio/capillaries) - distributed batch data processing framework. Stars:`67`.
- [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) - MySQL based distributed lock. Stars:`64`.
- [flowgraph](https://github.com/vectaport/flowgraph) - flow-based programming package. Stars:`63`.
- [drmaa](https://github.com/dgruber/drmaa) - Job submission library for cluster schedulers based on the DRMAA standard. Stars:`49`.
- [go-pdu](https://github.com/pdupub/go-pdu) - A decentralized identity-based social network. Stars:`48`.
- [committer](https://github.com/vadiminshakov/committer) - A distributed transactions management system (2PC/3PC implementation). Stars:`38`.
- [consistenthash](https://github.com/mbrostami/consistenthash) - Consistent hashing with configurable replicas. Stars:`28`.
- [gmsec](https://github.com/gmsec/micro) - A Go distributed systems development framework. Stars:`25`.
- [dynatomic](https://github.com/tylfin/dynatomic) - A library for using DynamoDB as an atomic counter. Stars:`17`.
- [failured](https://github.com/andy2046/failured) - adaptive accrual failure detector for distributed systems. Stars:`15`.
- [bedrock](https://github.com/z5labs/bedrock) - Provides a minimal, modular and composable foundation for quickly developing services and more use case specific frameworks in Go. Stars:`13`.
- [dynamolock](https://cirello.io/dynamolock) - DynamoDB-backed distributed locking implementation.
  performant communications system for digital systems, services, and devices.
- [raft](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol, by HashiCorp. Stars:`8.7K`.
- [rpcx](https://github.com/smallnest/rpcx) - Distributed pluggable RPC service framework like alibaba Dubbo. Stars:`8.2K`.
- [torrent](https://github.com/anacrolix/torrent) - BitTorrent client package. Stars:`5.8K`.
- [sponge](https://github.com/zhufuyi/sponge) - A distributed development framework that integrates automatic code generation, gin and grpc frameworks, base development frameworks. Stars:`2.3K`.
- [redis-lock](https://github.com/bsm/redislock) - Simplified distributed locking implementation using Redis. Stars:`1.7K`.
- [rain](https://github.com/cenkalti/rain) - BitTorrent client and library. Stars:`1.1K`.
- [trpc-go](https://github.com/trpc-group/trpc-go) - The Go language implementation of tRPC, which is a pluggable, high-performance RPC framework. Stars:`1.0K`.
- [raft](https://github.com/etcd-io/raft) - Go implementation of the Raft consensus protocol, by CoreOS. Stars:`885`.
- [ringpop-go](https://github.com/uber/ringpop-go) - Scalable, fault-tolerant application-layer sharding for Go applications. Stars:`854`.
- [Temporal](https://github.com/temporalio/sdk-go) - Durable execution system for making code fault-tolerant and simple. Stars:`687`.
- [opentelemetry-go-auto-instrumentation](https://github.com/alibaba/opentelemetry-go-auto-instrumentation) - OpenTelemetry Compile-Time Instrumentation for Golang. Stars:`638`.
- [sleuth](https://github.com/ursiform/sleuth) - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). Stars:`381`.
- [Tarmac](https://github.com/tarmac-project/tarmac) - Framework for writing functions, microservices, or monoliths with WebAssembly Stars:`335`.
- [outboxer](https://github.com/italolelis/outboxer) - Outboxer is a go library that implements the outbox pattern. Stars:`163`.
- [Semaphore](https://github.com/jexia/semaphore) - A straightforward (micro) service orchestrator. Stars:`94`.
- [outbox](https://github.com/oagudo/outbox) - Lightweight library for the transactional outbox pattern in Go, not tied to any specific relational database or broker. Stars:`79`.
- [resgate](https://resgate.io/) - Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
- [pjrpc](https://gitlab.com/pjrpc/pjrpc) - Golang JSON-RPC Server-Client with Protobuf spec.
- [pglock](https://cirello.io/pglock) - PostgreSQL-backed distributed locking implementation.

**[⬆ back to top](#contents)**

## Dynamic DNS

_Tools for updating dynamic DNS records._

- [GoDNS](https://github.com/timothyye/godns) - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. Stars:`1.6K`.
- [DDNS](https://github.com/skibish/ddns) - Personal DDNS client with Digital Ocean Networking DNS as backend. Stars:`263`.
- [dyndns](https://gitlab.com/alcastle/dyndns) - Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.

**[⬆ back to top](#contents)**

## Email

_Libraries and tools that implement email creation and sending._

- [MailHog](https://github.com/mailhog/MailHog) - Email and SMTP testing with web and API interface. Stars:`15.1K`.
- [Mailpit](https://github.com/axllent/mailpit) - Email and SMTP testing tool for developers. Stars:`7.4K`.
- [Maddy](https://github.com/foxcpp/maddy) - All-in-one (SMTP, IMAP, DKIM, DMARC, MTA-STS, DANE) email server Stars:`5.6K`.
- [mox](https://github.com/mjl-/mox) - Modern full-featured secure mail server for low-maintenance, self-hosted email. Stars:`5.1K`.
- [hermes](https://github.com/matcornic/hermes) - Golang package that generates clean, responsive HTML e-mails. Stars:`2.9K`.
- [email](https://github.com/jordan-wright/email) - A robust and flexible email library for Go. Stars:`2.7K`.
- [go-imap](https://github.com/emersion/go-imap) - IMAP library for clients and servers. Stars:`2.2K`.
- [email-verifier](https://github.com/AfterShip/email-verifier) - A Go library for email verification without sending any emails. Stars:`1.4K`.
- [go-mail](https://github.com/wneessen/go-mail) - A simple Go library for sending mails in Go. Stars:`1.1K`.
- [SendGrid](https://github.com/sendgrid/sendgrid-go) - SendGrid's Go library for sending email. Stars:`1.0K`.
- [mailgun-go](https://github.com/mailgun/mailgun-go) - Go library for sending mail with the Mailgun API. Stars:`732`.
- [go-simple-mail](https://github.com/xhit/go-simple-mail) - Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. Stars:`681`.
- [go-message](https://github.com/emersion/go-message) - Streaming library for the Internet Message Format and mail messages. Stars:`420`.
- [douceur](https://github.com/aymerick/douceur) - CSS inliner for your HTML emails. Stars:`252`.
- [Hectane](https://github.com/hectane/hectane) - Lightweight SMTP client providing an HTTP API. Stars:`224`.
- [go-premailer](https://github.com/vanng822/go-premailer) - Inline styling for HTML mail in Go. Stars:`171`.
- [smtpmock](https://github.com/mocktools/go-smtp-mock) - Lightweight configurable multithreaded fake SMTP server. Mimic any SMTP behaviour for your test environment. Stars:`153`.
- [mailchain](https://github.com/mailchain/mailchain) - Send encrypted emails to blockchain addresses written in Go. Stars:`143`.
- [truemail-go](https://github.com/truemail-rb/truemail-go) - Configurable Golang email validator/verifier. Verify email via Regex, DNS, SMTP and even more. Stars:`123`.
- [go-dkim](https://github.com/toorop/go-dkim) - DKIM library, to sign & verify email. Stars:`97`.
- [smtp](https://github.com/mailhog/smtp) - SMTP server protocol state machine. Stars:`75`.
- [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) - Golang library for providing a canonical representation of email address. Stars:`72`.
- [go-email-validator](https://github.com/go-email-validator/go-email-validator) - Modular email validator for syntax, disposable, smtp, etc... checking. Stars:`51`.
- [mailx](https://github.com/valord577/mailx) - Mailx is a library that makes it easier to send email via SMTP. It is an enhancement of the golang standard library `net/smtp`. Stars:`19`.
- [chasquid](https://blitiri.com.ar/p/chasquid) - SMTP server written in Go.

**[⬆ back to top](#contents)**

## Embeddable Scripting Languages

_Embedding other languages inside your go code._

- [FrankenPHP](https://github.com/dunglas/frankenphp) - PHP embedded in Go, with a `net/http` handler. Stars:`9.8K`.
- [expr](https://github.com/antonmedv/expr) - Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing. Stars:`7.1K`.
- [gopher-lua](https://github.com/yuin/gopher-lua) - Lua 5.1 VM and compiler written in Go. Stars:`6.7K`.
- [goja](https://github.com/dop251/goja) - ECMAScript 5.1(+) implementation in Go. Stars:`6.3K`.
- [tengo](https://github.com/d5/tengo) - Bytecode compiled script language for Go. Stars:`3.7K`.
- [go-lua](https://github.com/Shopify/go-lua) - Port of the Lua 5.2 VM to pure Go. Stars:`3.3K`.
- [cel-go](https://github.com/google/cel-go) - Fast, portable, non-Turing complete expression evaluation with gradual typing. Stars:`2.6K`.
- [starlark-go](https://github.com/google/starlark-go) - Go implementation of Starlark: Python-like language with deterministic evaluation and hermetic execution. Stars:`2.5K`.
- [metacall](https://github.com/metacall/core) - Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, WebAssembly, Java, Cobol and more. Stars:`1.7K`.
- [Wa/凹语言](https://github.com/wa-lang/wa) - The Wa Programming Language embedded in Go. Stars:`1.6K`.
- [go-python](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API. Stars:`1.5K`.
- [anko](https://github.com/mattn/anko) - Scriptable interpreter written in Go. Stars:`1.5K`.
- [go-php](https://github.com/deuill/go-php) - PHP bindings for Go. Stars:`943`.
- [gval](https://github.com/PaesslerAG/gval) - A highly customizable expression language written in Go. Stars:`797`.
- [go-duktape](https://github.com/olebedev/go-duktape) - Duktape JavaScript engine bindings for Go. Stars:`778`.
- [prolog](https://github.com/ichiban/prolog) - Embeddable Prolog. Stars:`678`.
- [golua](https://github.com/aarzilli/golua) - Go bindings for Lua C API. Stars:`675`.
- [gisp](https://github.com/jcla1/gisp) - Simple LISP in Go. Stars:`525`.
- [gentee](https://github.com/gentee/gentee) - Embeddable scripting programming language. Stars:`139`.
- [binder](https://github.com/alexeyco/binder) - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). Stars:`79`.
- [ecal](https://github.com/krotik/ecal) - A simple embeddable scripting language which supports concurrent event processing. Stars:`48`.
- [purl](https://github.com/ian-kent/purl) - Perl 5.18.2 embedded in Go. Stars:`41`.
- [starlet](https://github.com/1set/starlet) - Go wrapper for [starlark-go](https://github.com/google/starlark-go) that simplifies script execution, offers data conversion, and useful Starlark libraries and extensions. Stars:`31`.
- [ngaro](https://github.com/db47h/ngaro) - Embeddable Ngaro VM implementation enabling scripting in Retro. Stars:`30`.
- [goal](https://codeberg.org/anaseto/goal) - An embeddable scripting array language.

**[⬆ back to top](#contents)**

## Error Handling

_Libraries for handling errors._

- [go-multierror](https://github.com/hashicorp/go-multierror) - Go (golang) package for representing a list of errors as a single error. Stars:`2.5K`.
- [errors](https://github.com/cockroachdb/errors) - Go error library with error portability over the network. Stars:`2.3K`.
- [eris](https://github.com/rotisserie/eris) - A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. Stars:`1.7K`.
- [errorx](https://github.com/joomcode/errorx) - A feature rich error package with stack traces, composition of errors and more. Stars:`1.3K`.
- [multierr](https://github.com/uber-go/multierr) - Package for representing a list of errors as a single error. Stars:`1.1K`.
- [tracerr](https://github.com/ztrue/tracerr) - Golang errors with stack trace and source fragments. Stars:`1.1K`.
- [oops](https://github.com/samber/oops) - Error handling with context, stack trace and source fragments. Stars:`716`.
- [errlog](https://github.com/snwfdhmp/errlog) - Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. Stars:`461`.
- [emperror](https://github.com/emperror/emperror) - Error handling tools and best practices for Go libraries and applications. Stars:`359`.
- [Fault](https://github.com/Southclaws/fault) - An ergonomic mechanism for wrapping errors in order to facilitate structured metadata and context for error values. Stars:`289`.
- [errors](https://github.com/emperror/errors) - Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. Stars:`201`.
- [errors](https://github.com/naughtygopher/errors) - Drop-in replacement for builtin Go errors. This is a minimal error handling package with custom error types, user friendly messages, Unwrap & Is. With very easy to use and straightforward helper functions. Stars:`74`.
- [exception](https://github.com/rbrahul/exception) - A simple utility package for exception handling with try-catch in Golang. Stars:`37`.
- [Falcon](https://github.com/SonicRoshan/falcon) - A Simple Yet Highly Powerful Package For Error Handling. Stars:`11`.
- [errors](https://github.com/PumpkinSeed/errors) - The most simple error wrapper with awesome performance and minimal memory overhead. Stars:`7`.
- [errors](https://github.com/neuronlabs/errors) - Simple golang error handling with classification primitives. Stars:`6`.
- [metaerr](https://github.com/quantumcycle/metaerr) - A library to create your custom error builders producing structured errors with metadata from different sources and optional stacktraces. Stars:`4`.
- [errors](https://gitlab.com/tozd/go/errors) - Providing errors with a stack trace and optional structured details. Compatible with github.com/pkg/errors API but does not use it internally.

**[⬆ back to top](#contents)**

## File Handling

_Libraries for handling files and file systems._

- [pdfcpu](https://github.com/pdfcpu/pdfcpu) - PDF processor. Stars:`7.8K`.
- [afero](https://github.com/spf13/afero) - FileSystem Abstraction System for Go. Stars:`6.3K`.
- [gdu](https://github.com/dundee/gdu) - Disk usage analyzer with console interface. Stars:`4.7K`.
- [go-wkhtmltopdf](https://github.com/SebastiaanKlippert/go-wkhtmltopdf) - A package to convert an HTML template to a PDF file. Stars:`1.1K`.
- [notify](https://github.com/rjeczalik/notify) - File system event notification library with simple API, similar to os/signal. Stars:`915`.
- [gofs](https://github.com/no-src/gofs) - A cross-platform real-time file synchronization tool out of the box. Stars:`514`.
- [vfs](https://github.com/C2FO/vfs) - A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. Stars:`347`.
- [afs](https://github.com/viant/afs) - Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. Stars:`337`.
- [iso9660](https://github.com/kdomanski/iso9660) - A package for reading and creating ISO9660 disk images Stars:`276`.
- [go-exiftool](https://github.com/barasher/go-exiftool) - Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). Stars:`270`.
- [go-csv-tag](https://github.com/artonge/go-csv-tag) - Load csv file using tag. Stars:`127`.
- [parquet](https://github.com/parsyl/parquet) - Read and write [parquet](https://parquet.apache.org) files. Stars:`119`.
- [skywalker](https://github.com/dixonwille/skywalker) - Package to allow one to concurrently go through a filesystem with ease. Stars:`103`.
- [opc](https://github.com/qmuntal/opc) - Load Open Packaging Conventions (OPC) files for Go. Stars:`75`.
- [gulter](https://github.com/adelowo/gulter) - A simple HTTP middleware to automatically handle all your file upload needs Stars:`67`.
- [baraka](https://github.com/xis/baraka) - A library to process http file uploads easily. Stars:`60`.
- [go-gtfs](https://github.com/artonge/go-gtfs) - Load gtfs files in go. Stars:`48`.
- [gut/yos](https://github.com/1set/gut) - Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links. Stars:`27`.
- [todotxt](https://github.com/1set/todotxt) - Go library for Gina Trapani's [_todo.txt_](http://todotxt.org/) files, supports parsing and manipulating of task lists in the [_todo.txt_ format](https://github.com/todotxt/todo.txt). Stars:`25`.
- [higgs](https://github.com/dastoori/higgs) - A tiny cross-platform Go library to hide/unhide files and directories. Stars:`24`.
- [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) - Copy files for humans. Stars:`22`.
- [pathtype](https://github.com/jonchun/pathtype) - Treat paths as their own type instead of using strings. Stars:`13`.
- [fastwalk](https://github.com/charlievieth/fastwalk) - Fast parallel directory traversal library (used by [fzf](https://github.com/junegunn/fzf)).
- [copy](https://github.com/otiai10/copy) - Copy directory recursively.
- [checksum](https://github.com/codingsince1985/checksum) - Compute message digest, like MD5, SHA256, SHA1, CRC or BLAKE2s, for large files.
- [flop](https://github.com/homedepot/flop) - File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).

**[⬆ back to top](#contents)**

## Financial

_Packages for accounting and finance._

- [decimal](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers. Stars:`6.9K`.
- [ticker](https://github.com/achannarasappa/ticker) - Terminal stock watcher and stock position tracker. Stars:`5.4K`.
- [go-money](https://github.com/rhymond/go-money) - Implementation of Fowler's Money pattern. Stars:`1.8K`.
- [bbgo](https://github.com/c9s/bbgo) - A crypto trading bot framework written in Go. Including common crypto exchange API, standard indicators, back-testing and many built-in strategies. Stars:`1.4K`.
- [ledger](https://github.com/formancehq/ledger) - A programmable financial ledger that provides a foundation for money-moving applications. Stars:`997`.
- [accounting](https://github.com/leekchan/accounting) - money and currency formatting for golang. Stars:`899`.
- [techan](https://github.com/sdcoffey/techan) - Technical analysis library with advanced market analysis and trading strategies. Stars:`868`.
- [currency](https://github.com/bojanz/currency) - Handles currency amounts, provides currency information and formatting. Stars:`597`.
- [orderbook](https://github.com/i25959341/orderbook) - Matching Engine for Limit Order Book in Golang. Stars:`505`.
- [ach](https://github.com/moov-io/ach) - A reader, writer, and validator for Automated Clearing House (ACH) files. Stars:`498`.
- [go-finance](https://github.com/alpeb/go-finance) - Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. Stars:`186`.
- [decimal](https://github.com/govalues/decimal) - Immutable decimal numbers with panic-free arithmetic. Stars:`179`.
- [udecimal](https://github.com/quagmt/udecimal) - High performance, high precision, zero allocation fixed-point decimal library for financial applications. Stars:`149`.
- [ofxgo](https://github.com/aclindsa/ofxgo) - Query OFX servers and/or parse the responses (with example command-line client). Stars:`144`.
- [sleet](https://github.com/BoltApp/sleet) - One unified interface for multiple Payment Service Providers (PsP) to process online payment. Stars:`140`.
- [transaction](https://github.com/claygod/transaction) - Embedded transactional database of accounts, running in multithreaded mode. Stars:`136`.
- [vat](https://github.com/dannyvankooten/vat) - VAT number validation & EU VAT rates. Stars:`115`.
- [go-finnhub](https://github.com/m1/go-finnhub) - Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges. Stars:`94`.
- [payme](https://github.com/jovandeginste/payme) - QR code generator (ASCII & PNG) for SEPA payments. Stars:`88`.
- [currency](https://github.com/naughtygopher/currency) - High performant & accurate currency computation package. Stars:`62`.
- [money](https://github.com/govalues/money) - Immutable monetary amounts and exchange rates with panic-free arithmetic. Stars:`38`.
- [fpmoney](https://github.com/nikolaydubina/fpmoney) - Fast and simple ISO4217 fixed-point decimal money. Stars:`34`.
- [fpdecimal](https://github.com/nikolaydubina/fpdecimal) - Fast and precise serialization and arithmetic for small fixed-point decimals Stars:`32`.
- [go-finance](https://github.com/pieterclaerhout/go-finance) - Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers. Stars:`29`.
- [go-nowpayments](https://github.com/matm/go-nowpayments) - Library for the crypto NOWPayments API. Stars:`3`.
- [swift](https://code.pfad.fr/swift/) - Offline validity check of IBAN (International Bank Account Number) and retrieval of BIC (for some countries).

**[⬆ back to top](#contents)**

## Forms

_Libraries for working with forms._

- [nosurf](https://github.com/justinas/nosurf) - CSRF protection middleware for Go. Stars:`1.7K`.
- [gorilla/csrf](https://github.com/gorilla/csrf) - CSRF protection for Go web applications & services. Stars:`1.1K`.
- [form](https://github.com/go-playground/form) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. Stars:`855`.
- [httpin](https://github.com/ggicci/httpin) - Decode an HTTP request into a custom struct, including querystring, forms, HTTP headers, etc. Stars:`368`.
- [conform](https://github.com/leebenson/conform) - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. Stars:`325`.
- [formam](https://github.com/monoculum/formam) - decode form's values into a struct. Stars:`192`.
- [forms](https://github.com/albrow/forms) - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. Stars:`142`.
- [qs](https://github.com/sonh/qs) - Go module for encoding structs into URL query parameters. Stars:`78`.
- [checker](https://github.com/cinar/checker) - Checker helps validating user input through rules defined in struct tags or directly through functions. Stars:`44`.
- [bind](https://github.com/robfig/bind) - Bind form data to any Go values. Stars:`32`.
- [queryparam](https://github.com/tomwright/queryparam) - Decode `url.Values` into usable struct values of standard or custom types. Stars:`19`.
- [gbind](https://github.com/bdjimmy/gbind) - Bind data to any Go value. Can use built-in and custom expression binding capabilities; supports data validation Stars:`10`.

**[⬆ back to top](#contents)**

## Functional

_Packages to support functional programming in Go._

- [mo](https://github.com/samber/mo) - Monads and popular FP abstractions, based on Go 1.18+ Generics (Option, Result, Either...). Stars:`3.0K`.
- [go-underscore](https://github.com/tobyhede/go-underscore) - Useful collection of helpfully functional Go collection utilities. Stars:`1.3K`.
- [go-functional](https://github.com/BooleanCat/go-functional) - Functional programming in Go using generics Stars:`509`.
- [fpGo](https://github.com/TeaEntityLab/fpGo) - Monad, Functional Programming features for Golang. Stars:`351`.
- [fp-go](https://github.com/repeale/fp-go) - Collection of Functional Programming helpers powered by Golang 1.18+ generics. Stars:`315`.
- [gofp](https://github.com/rbrahul/gofp) - A lodash like powerful utility library for Golang. Stars:`151`.
- [fuego](https://github.com/seborama/fuego) - Functional Experiment in Go. Stars:`146`.
- [underscore](https://github.com/rjNemo/underscore) - Functional programming helpers for Go 1.18 and beyond. Stars:`117`.
- [valor](https://github.com/phelmkamp/valor) - Generic option and result types that optionally contain a value. Stars:`18`.

**[⬆ back to top](#contents)**

## Game Development

_Awesome game development libraries._

- [Ebitengine](https://github.com/hajimehoshi/ebiten) - dead simple 2D game engine in Go. Stars:`12.3K`.
- [Leaf](https://github.com/name5566/leaf) - Lightweight game server framework. Stars:`5.4K`.
- [nano](https://github.com/lonng/nano) - Lightweight, facility, high performance golang based game server framework. Stars:`3.1K`.
- [g3n](https://github.com/g3n/engine) - Go 3D Game Engine. Stars:`3.0K`.
- [goworld](https://github.com/xiaonanln/goworld) - Scalable game server engine, featuring space-entity framework and hot-swapping. Stars:`2.7K`.
- [Pitaya](https://github.com/topfreegames/pitaya) - Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. Stars:`2.6K`.
- [go-sdl2](https://github.com/veandco/go-sdl2) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). Stars:`2.3K`.
- [raylib-go](https://github.com/gen2brain/raylib-go) - Go bindings for [raylib](https://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. Stars:`2.1K`.
- [engo](https://github.com/EngoEngine/engo) - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. Stars:`1.8K`.
- [Oak](https://github.com/oakmound/oak) - Pure Go game engine. Stars:`1.6K`.
- [termloop](https://github.com/JoelOtter/termloop) - Terminal-based game engine for Go, built on top of Termbox. Stars:`1.5K`.
- [gonet](https://github.com/xtaci/gonet) - Game server skeleton implemented with golang. Stars:`1.3K`.
- [Harfang3D](https://github.com/harfang3d/harfang3d) - 3D engine for the Go language, works on Windows and Linux ([Harfang on Go.dev](https://github.com/harfang3d/harfang-go)). Stars:`621`.
- [go-astar](https://github.com/beefsack/go-astar) - Go implementation of the A\* path finding algorithm. Stars:`617`.
- [Pixel](https://github.com/gopxl/pixel) - Hand-crafted 2D game library in Go. Stars:`341`.
- [go3d](https://github.com/ungerik/go3d) - Performance oriented 2D/3D math package for Go. Stars:`325`.
- [tile](https://github.com/kelindar/tile) - Data-oriented and cache-friendly 2D Grid library (TileMap), includes pathfinding, observers and import/export. Stars:`198`.
- [ecs](https://github.com/andygeiss/ecs) - Build your own Game-Engine based on the Entity Component System concept in Golang. Stars:`149`.
- [Ark](https://github.com/mlange-42/ark) - Archetype-based Entity Component System (ECS) for Go. Stars:`94`.
- [prototype](https://github.com/gonutz/prototype) - Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API. Stars:`91`.
- [fantasyname](https://github.com/s0rg/fantasyname) - Fantasy names generator. Stars:`36`.
- [grid](https://github.com/s0rg/grid) - Generic 2D grid with ray-casting, shadow-casting and path finding. Stars:`22`.

**[⬆ back to top](#contents)**

## Generators

_Tools that generate Go code._

- [oapi-codegen](https://github.com/deepmap/oapi-codegen) - This package contains a set of utilities for generating Go boilerplate code for services based on OpenAPI 3.0 API definitions. Stars:`7.4K`.
- [go-linq](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go. Stars:`3.6K`.
- [jennifer](https://github.com/dave/jennifer) - Generate arbitrary Go code without templates. Stars:`3.5K`.
- [goderive](https://github.com/awalterschulze/goderive) - Derives functions from input types Stars:`1.3K`.
- [GoWrap](https://github.com/hexdigest/gowrap) - Generate decorators for Go interfaces using simple templates. Stars:`1.2K`.
- [go-enum](https://github.com/abice/go-enum) - Code generation for enums from code comments. Stars:`851`.
- [goverter](https://github.com/jmattheis/goverter) - Generate converters by defining an interface. Stars:`734`.
- [interfaces](https://github.com/rjeczalik/interfaces) - Command line tool for generating interface definitions. Stars:`431`.
- [copygen](https://github.com/switchupcb/copygen) - Generate any code based on Go types, including type-to-type converters (copy code) without reflection by default. Stars:`396`.
- [convergen](https://github.com/reedom/convergen) - Feature rich type-to-type copy code generator. Stars:`48`.
- [generis](https://github.com/senselogic/GENERIS) - Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. Stars:`47`.
- [typeregistry](https://github.com/xiaoxin01/typeregistry) - A library to create type dynamically. Stars:`23`.
- [go-enum-encoding](https://github.com/nikolaydubina/go-enum-encoding) - Code generation for enum encoding from code comments. Stars:`15`.

**[⬆ back to top](#contents)**

## Geographic

_Geographic tools and servers_

- [Tile38](https://github.com/tidwall/tile38) - Geolocation DB with spatial index and realtime geofencing. Stars:`9.4K`.
- [S2 geometry](https://github.com/golang/geo) - S2 geometry library in Go. Stars:`1.8K`.
- [mbtileserver](https://github.com/consbio/mbtileserver) - A simple Go-based server for map tiles stored in mbtiles format. Stars:`740`.
- [geoos](https://github.com/spatial-go/geoos) - A library provides spatial data and geometric algorithms. Stars:`524`.
- [osm](https://github.com/paulmach/osm) - Library for reading, writing and working with OpenStreetMap data and APIs. Stars:`414`.
- [H3](https://github.com/uber/h3-go) - Go bindings for H3, a hierarchical hexagonal geospatial indexing system. Stars:`358`.
- [godal](https://github.com/airbusgeo/godal) - Go wrapper for GDAL. Stars:`168`.
- [simplefeatures](https://github.com/peterstace/simplefeatures) - simplesfeatures is a 2D geometry library that provides Go types that model geometries, as well as algorithms that operate on them. Stars:`150`.
- [WGS84](https://github.com/wroge/wgs84) - Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). Stars:`140`.
- [geoserver](https://github.com/hishamkaram/geoserver) - geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. Stars:`93`.
- [gismanager](https://github.com/hishamkaram/gismanager) - Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. Stars:`54`.
- [pbf](https://github.com/maguro/pbf) - OpenStreetMap PBF golang encoder/decoder. Stars:`50`.
- [S2 geojson](https://github.com/pantrif/s2-geojson) - Convert geojson to s2 cells & demonstrating some S2 geometry features on map. Stars:`35`.
- [Web-Mercator-Projection](https://github.com/jorelosorio/web-mercator-projection) A project to easily use and convert LonLat, Point and Tile to display info, markers, etc, in a map using the Web Mercator Projection. Stars:`9`.
- [H3 GeoJSON](https://github.com/mmadfox/go-geojson2h3) - Conversion utilities between H3 indexes and GeoJSON. Stars:`5`.
- [H3GeoDist](https://github.com/mmadfox/go-h3geo-dist) - Distribution of Uber H3geo cells by virtual nodes.

**[⬆ back to top](#contents)**

## Go Compilers

_Tools for compiling Go to other languages and vice-versa._


- [gopherjs](https://github.com/gopherjs/gopherjs) - Compiler from Go to JavaScript. Stars:`13.0K`.
- [bunster](https://github.com/yassinebenaid/bunster) - Compile shell scripts to Go. Stars:`2.5K`.
- [c4go](https://github.com/Konstantin8105/c4go) - Transpile C code to Go code. Stars:`374`.
- [go2hx](https://github.com/go2hx/go2hx) - Compiler from Go to Haxe to Javascript/C++/Java/C#. Stars:`138`.
- [esp32](https://github.com/andygeiss/esp32-transpiler) - Transpile Go into Arduino code. Stars:`93`.
- [f4go](https://github.com/Konstantin8105/f4go) - Transpile FORTRAN 77 code to Go code. Stars:`48`.

**[⬆ back to top](#contents)**

## Goroutines

_Tools for managing and working with Goroutines._

- [ants](https://github.com/panjf2000/ants) - A high-performance and low-cost goroutine pool in Go. Stars:`13.9K`.
- [conc](https://github.com/sourcegraph/conc) - `conc` is your toolbelt for structured concurrency in go, making common tasks easier and safer. Stars:`10.0K`.
- [tunny](https://github.com/Jeffail/tunny) - Goroutine pool for golang. Stars:`4.0K`.
- [goworker](https://github.com/benmanns/goworker) - goworker is a Go-based background worker. Stars:`2.8K`.
- [pond](https://github.com/alitto/pond) - Minimalistic and High-performance goroutine worker pool written in Go. Stars:`1.9K`.
- [rill](https://github.com/destel/rill) - Go toolkit for clean, composable, channel-based concurrency. Stars:`1.7K`.
- [workerpool](https://github.com/gammazero/workerpool) - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. Stars:`1.4K`.
- [grpool](https://github.com/ivpusic/grpool) - Lightweight Goroutine pool. Stars:`744`.
- [pool](https://github.com/go-playground/pool) - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. Stars:`729`.
- [gowp](https://github.com/xxjwxc/gowp) - gowp is concurrency limiting goroutine pool. Stars:`526`.
- [flowmatic](https://github.com/carlmjohnson/flowmatic) - Structured concurrency made easy. Stars:`390`.
- [routine](https://github.com/timandy/routine) - `routine` is a `ThreadLocal` for go library. It encapsulates and provides some easy-to-use, non-competitive, high-performance `goroutine` context access interfaces, which can help you access coroutine context information more gracefully. Stars:`269`.
- [go-floc](https://github.com/workanator/go-floc) - Orchestrate goroutines with ease. Stars:`268`.
- [async](https://github.com/reugn/async) - An alternative sync library for Go (Future, Promise, Locks). Stars:`253`.
- [go-actor](https://github.com/vladopajic/go-actor) - A tiny library for writing concurrent programs using actor model. Stars:`239`.
- [go-flow](https://github.com/kamildrazkiewicz/go-flow) - Control goroutines execution order. Stars:`222`.
- [artifex](https://github.com/borderstech/artifex) - Simple in-memory job queue for Golang using worker-based dispatching. Stars:`212`.
- [semaphore](https://github.com/marusama/semaphore) - Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). Stars:`177`.
- [go-workers](https://github.com/catmullet/go-workers) - Easily and safely run workers for large data processing pipelines. Stars:`166`.
- [neilotoole/errgroup](https://github.com/neilotoole/errgroup) - Drop-in alternative to `sync/errgroup`, limited to a pool of N worker goroutines. Stars:`165`.
- [cyclicbarrier](https://github.com/marusama/cyclicbarrier) - CyclicBarrier for golang. Stars:`153`.
- [async](https://github.com/studiosol/async) - A safe way to execute functions asynchronously, recovering them in case of panic. Stars:`139`.
- [gollback](https://github.com/vardius/gollback) - asynchronous simple function utilities, for managing execution of closures and callbacks. Stars:`123`.
- [Hunch](https://github.com/AaronJan/Hunch) - Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. Stars:`105`.
- [threadpool](https://github.com/shettyh/threadpool) - Golang threadpool implementation. Stars:`104`.
- [semaphore](https://github.com/kamilsk/semaphore) - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. Stars:`103`.
- [worker-pool](https://github.com/vardius/worker-pool) - goworker is a Go simple async worker pool. Stars:`92`.
- [gpool](https://github.com/Sherifabdlnaby/gpool) - manages a resizeable pool of context-aware goroutines to bound concurrency. Stars:`90`.
- [goccm](https://github.com/zenthangplus/goccm) - Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently. Stars:`73`.
- [gowl](https://github.com/hamed-yousefi/gowl) - Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status. Stars:`70`.
- [nursery](https://github.com/arunsworld/nursery) - Structured concurrency in Go. Stars:`69`.
- [routine](https://github.com/x-mod/routine) - go routine control with context, support: Main, Go, Pool and some useful Executors. Stars:`62`.
- [gohive](https://github.com/loveleshsharma/gohive) - A highly performant and easy to use Goroutine pool for Go. Stars:`53`.
- [kyoo](https://github.com/dirkaholic/kyoo) - Provides an unlimited job queue and concurrent worker pools. Stars:`50`.
- [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) - Like `sync.WaitGroup` with error handling and concurrency control. Stars:`49`.
- [parallel-fn](https://github.com/rafaeljesus/parallel-fn) - Run functions in parallel. Stars:`37`.
- [go-trylock](https://github.com/subchen/go-trylock) - TryLock support on read-write lock for Golang. Stars:`37`.
- [channelify](https://github.com/ddelizia/channelify) - Transform your function to return channels for easy and powerful parallel processing. Stars:`33`.
- [stl](https://github.com/ssgreg/stl) - Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. Stars:`30`.
- [execpool](https://github.com/hexdigest/execpool) - A pool built around exec.Cmd that spins up a given number of processes in advance and attaches stdin and stdout to them when needed. Very similar to FastCGI or Apache Prefork MPM but works for any command. Stars:`29`.
- [breaker](https://github.com/kamilsk/breaker) - Flexible mechanism to make execution flow interruptible. Stars:`20`.
- [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) - Concurrency limiter with support for timeouts, dynamic priority and context cancellation of goroutines. Stars:`18`.
- [conexec](https://github.com/ITcathyh/conexec) - A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency. Stars:`18`.
- [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) - Manage a pool of goroutines using this lightweight library with a simple API. Stars:`16`.
- [async](https://github.com/yaitoo/async) - An asynchronous task package with async/await style for Go. Stars:`15`.
- [go-workerpool](https://github.com/zenthangplus/go-workerpool) - Inspired from Java Thread Pool, Go WorkerPool aims to control heavy Go Routines. Stars:`12`.
- [async-job](https://github.com/lab210-dev/async-job) - AsyncJob is an asynchronous queue job manager with light code, clear and speed. Stars:`12`.
- [hands](https://github.com/duanckham/hands) - A process controller used to control the execution and return strategies of multiple goroutines. Stars:`10`.
- [go-accumulator](https://github.com/nar10z/go-accumulator) - Solution for accumulation of events and their subsequent processing. Stars:`8`.
- [oversight](https://pkg.go.dev/cirello.io/oversight) - Oversight is a complete implementation of the Erlang supervision trees.

**[⬆ back to top](#contents)**

## GUI

_Libraries for building GUI Applications._

_Toolkits_

- [fyne](https://github.com/fyne-io/fyne) - Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android. Stars:`26.9K`.
- [webview](https://github.com/zserge/webview) - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). Stars:`13.2K`.
- [qt](https://github.com/therecipe/qt) - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). Stars:`10.7K`.
- [app](https://github.com/murlokswarm/app) - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. Stars:`8.7K`.
- [ui](https://github.com/andlabs/ui) - Platform-native GUI library for Go. Cross platform. Stars:`8.4K`.
- [walk](https://github.com/lxn/walk) - Windows application library kit for Go. Stars:`7.0K`.
- [DarwinKit](https://github.com/progrium/darwinkit) - Build native macOS applications using Go. Stars:`5.3K`.
- [go-sciter](https://github.com/sciter-sdk/go-sciter) - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. Stars:`2.6K`.
- [gotk3](https://github.com/gotk3/gotk3) - Go bindings for GTK3. Stars:`2.2K`.
- [Cogent Core](https://github.com/cogentcore/core) - A framework for building 2D and 3D apps that run on macOS, Windows, Linux, iOS, Android, and the web. Stars:`2.1K`.
- [Spot](https://github.com/roblillack/spot) - Reactive, cross-platform desktop GUI toolkit. Stars:`1.2K`.
- [energy](https://github.com/energye/energy) - Cross-platform based on LCL(Native System UI Control Library) and CEF(Chromium Embedded Framework) (Windows/ macOS / Linux) Stars:`502`.
- [gowd](https://github.com/dtylman/gowd) - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. Stars:`439`.
- [cimgui-go](https://github.com/AllenDang/cimgui-go) - Auto generated Go wrapper for [Dear ImGui](https://github.com/ocornut/imgui) via [cimgui](https://github.com/cimgui/cimgui). Stars:`434`.
- [unison](https://github.com/richardwilkes/unison) - A unified graphical user experience toolkit for Go desktop applications. macOS, Windows, and Linux are supported. Stars:`280`.
- [goradd/html5tag](https://github.com/goradd/html5tag) - Library for outputting HTML5 tags. Stars:`13`.
- [Goey](https://bitbucket.org/rj/goey/src/master/) - Cross platform UI toolkit aggregator for Windows / Linux / Mac. GTK, Cocoa, Windows API
- [go-gtk](https://mattn.github.io/go-gtk/) - Go bindings for GTK.
- [Wails](https://wails.io) - Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
- [gio](https://gioui.org) - Gio is a library for writing cross-platform immediate mode GUI-s in Go. Gio supports all the major platforms: Linux, macOS, Windows, Android, iOS, FreeBSD, OpenBSD and WebAssembly.

_Interaction_

- [robotgo](https://github.com/go-vgo/robotgo) - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. Stars:`10.3K`.
- [systray](https://github.com/getlantern/systray) - Cross platform Go library to place an icon and menu in the notification area. Stars:`3.5K`.
- [zenity](https://github.com/ncruces/zenity) - Cross-platform Go library and CLI to create simple dialogs that interact graphically with the user. Stars:`830`.
- [gosx-notifier](https://github.com/deckarep/gosx-notifier) - OSX Desktop Notifications library for Go. Stars:`589`.
- [trayhost](https://github.com/shurcooL/trayhost) - Cross-platform Go library to place an icon in the host operating system's taskbar. Stars:`256`.
- [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) - OSX Sleep/Wake notifications in golang. Stars:`38`.
- [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) - OSX library to notify about any (pluggable) activity on your machine. Stars:`31`.
- [AppIndicator Go](https://github.com/gopherlibs/appindicator) - Go bindings for libappindicator3 C library. Stars:`6`.

**[⬆ back to top](#contents)**

## Hardware

_Libraries, tools, and tutorials for interacting with hardware._

- [arduino-cli](https://github.com/arduino/arduino-cli) - Official Arduino CLI and library. Can run standalone, or be incorporated into larger Go projects. Stars:`4.6K`.
- [go-rpio](https://github.com/stianeikeland/go-rpio) - GPIO for Go, doesn't require cgo. Stars:`2.2K`.
- [ghw](https://github.com/jaypipes/ghw) - Golang hardware discovery/inspection library. Stars:`1.8K`.
- [emgo](https://github.com/ziutek/emgo) - Go-like language for programming embedded systems (e.g. STM32 MCU). Stars:`1.1K`.
- [sysinfo](https://github.com/zcalusic/sysinfo) - A pure Go library providing Linux OS / kernel / hardware system information. Stars:`551`.
- [goroslib](https://github.com/aler9/goroslib) - Robot Operating System (ROS) library for Go. Stars:`367`.
- [go-osc](https://github.com/hypebeast/go-osc) - Open Sound Control (OSC) bindings for Go. Stars:`215`.
- [joystick](https://github.com/0xcafed00d/joystick) - a polled API to read the state of an attached joystick. Stars:`69`.

**[⬆ back to top](#contents)**

## Images

_Libraries for manipulating images._

- [gocv](https://github.com/hybridgroup/gocv) - Go package for computer vision using OpenCV 3.3+. Stars:`7.2K`.
- [imaginary](https://github.com/h2non/imaginary) - Fast and simple HTTP microservice for image resizing. Stars:`5.9K`.
- [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package. Stars:`5.5K`.
- [gg](https://github.com/fogleman/gg) - 2D rendering in pure Go. Stars:`4.6K`.
- [bild](https://github.com/anthonynsimon/bild) - Collection of image processing algorithms in pure Go. Stars:`4.1K`.
- [gowitness](https://github.com/sensepost/gowitness) - Screenshoting webpages using go and headless chrome on command line. Stars:`3.9K`.
- [imagor](https://github.com/cshum/imagor) - Fast, secure image processing server and Go library, using libvips. Stars:`3.7K`.
- [ln](https://github.com/fogleman/ln) - 3D line art rendering in Go. Stars:`3.3K`.
- [bimg](https://github.com/h2non/bimg) - Small package for fast and efficient image processing using libvips. Stars:`2.9K`.
- [picfit](https://github.com/thoas/picfit) - An image resizing server written in Go. Stars:`2.3K`.
- [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation. Stars:`2.2K`.
- [pt](https://github.com/fogleman/pt) - Path tracing engine written in Go. Stars:`2.1K`.
- [smartcrop](https://github.com/muesli/smartcrop) - Finds good crops for arbitrary images and crop sizes. Stars:`1.8K`.
- [imagick](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API. Stars:`1.8K`.
- [gift](https://github.com/disintegration/gift) - Package of image processing filters. Stars:`1.8K`.
- [canvas](https://github.com/tdewolff/canvas) - Vector graphics to PDF, SVG or rasterized image. Stars:`1.6K`.
- [govips](https://github.com/davidbyttow/govips) - A lightning fast image processing and resizing library for Go. Stars:`1.5K`.
- [geopattern](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string. Stars:`1.3K`.
- [stegify](https://github.com/DimitarPetrov/stegify) - Go tool for LSB steganography, capable of hiding any file within an image. Stars:`1.2K`.
- [image2ascii](https://github.com/qeesung/image2ascii) - Convert image to ASCII. Stars:`929`.
- [goimagehash](https://github.com/corona10/goimagehash) - Go Perceptual image hashing package. Stars:`809`.
- [go-qrcode](https://github.com/yeqown/go-qrcode) - Generate QR codes with personalized styles, allowing adjustments to color, block size, shape, and icons. Stars:`737`.
- [govatar](https://github.com/o1egl/govatar) - Library and CMD tool for generating funny avatars. Stars:`607`.
- [draft](https://github.com/lucasepe/draft) - Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax. Stars:`581`.
- [mort](https://github.com/aldor007/mort) - Storage and image processing server written in Go. Stars:`514`.
- [go-nude](https://github.com/koyachi/go-nude) - Nudity detection with Go. Stars:`420`.
- [steganography](https://github.com/auyer/steganography) - Pure Go Library for LSB steganography. Stars:`340`.
- [go-webp](https://github.com/kolesa-team/go-webp) - Library for encode and decode webp pictures, using libwebp. Stars:`268`.
- [transformimgs](https://github.com/Pixboost/transformimgs) - Transformimgs resizes and optimises images for Web using next-generation formats. Stars:`267`.
- [gltf](https://github.com/qmuntal/gltf) - Efficient and robust glTF 2.0 reader, writer and validator. Stars:`261`.
- [mergi](https://github.com/noelyahan/mergi) - Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). Stars:`241`.
- [darkroom](https://github.com/gojek/darkroom) - An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. Stars:`235`.
- [rez](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD. Stars:`212`.
- [img](https://github.com/hawx/img) - Selection of image manipulation tools. Stars:`156`.
- [go-cairo](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library. Stars:`149`.
- [cameron](https://github.com/aofei/cameron) - An avatar generator for Go. Stars:`129`.
- [color-extractor](https://github.com/marekm4/color-extractor) - Dominant color extractor with no external dependencies. Stars:`100`.
- [gridder](https://github.com/shomali11/gridder) - A Grid based 2D Graphics library. Stars:`80`.
- [webp-server](https://github.com/mehdipourfar/webp-server) - Simple and minimal image server capable of storing, resizing, converting and caching images. Stars:`77`.
- [go-gd](https://github.com/bolknote/go-gd) - Go binding for GD library. Stars:`59`.
- [goimghdr](https://github.com/corona10/goimghdr) - The imghdr module determines the type of image contained in a file for Go. Stars:`40`.
- [go-webcolors](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go. Stars:`29`.
- [scout](https://github.com/jonoton/scout) - Scout is a standalone open source software solution for DIY video security. Stars:`23`.
- [mpo](https://github.com/donatj/mpo) - Decoder and conversion tool for MPO 3D Photos. Stars:`21`.

**[⬆ back to top](#contents)**

## IoT (Internet of Things)

_Libraries for programming devices of the IoT._

- [gobot](https://github.com/hybridgroup/gobot/) - Gobot is a framework for robotics, physical computing, and the Internet of Things. Stars:`9.2K`.
- [flogo](https://github.com/tibcosoftware/flogo) - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. Stars:`2.5K`.
- [ekuiper](https://github.com/lf-edge/ekuiper) - Lightweight data stream processing engine for IoT edge. Stars:`1.6K`.
- [shifu](https://github.com/Edgenesis/shifu) - Kubernetes native IoT development framework. Stars:`1.4K`.
- [rulego](https://github.com/rulego/rulego) - RuleGo is a lightweight, high-performance, embedded, orchestrable component-based rule engine for IoT edge. Stars:`1.2K`.
- [gatt](https://github.com/paypal/gatt) - Gatt is a Go package for building Bluetooth Low Energy peripherals. Stars:`1.2K`.
- [connectordb](https://github.com/connectordb/connectordb) - Open-Source Platform for Quantified Self & IoT. Stars:`418`.
- [devices](https://github.com/goiot/devices) - Suite of libraries for IoT devices, experimental for x/exp/io. Stars:`265`.
- [huego](https://github.com/amimof/huego) - An extensive Philips Hue client library for Go. Stars:`257`.
- [sensorbee](https://github.com/sensorbee/sensorbee) - Lightweight stream processing engine for IoT. Stars:`231`.
- [smart-home](https://github.com/e154/smart-home) - Software package for IoT automation. Stars:`92`.
- [iot](https://github.com/vaelen/iot/) - IoT is a simple framework for implementing a Google IoT Core device. Stars:`66`.
- [eywa](https://github.com/xcodersun/eywa) - Project Eywa is essentially a connection manager that keeps track of connected devices. Stars:`65`.
- [mainflux](https://github.com/Mainflux/mainflux) - Industrial IoT Messaging and Device Management Server. Stars:`54`.
- [periph](https://periph.io/) - Peripherals I/O to interface with low-level board facilities.

**[⬆ back to top](#contents)**

## Job Scheduler

_Libraries for scheduling jobs._

- [gocron](https://github.com/go-co-op/gocron) - Easy and fluent Go job scheduling. This is an actively maintained fork of [jasonlvhit/gocron](https://github.com/jasonlvhit/gocron). Stars:`6.4K`.
- [go-quartz](https://github.com/reugn/go-quartz) - Simple, zero-dependency scheduling library for Go. Stars:`1.9K`.
- [JobRunner](https://github.com/bamzi/jobrunner) - Smart and featureful cron job scheduler with job queuing and live monitoring built in. Stars:`1.1K`.
- [gron](https://github.com/roylee0704/gron) - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. Stars:`1.0K`.
- [scheduler](https://github.com/carlescere/scheduler) - Cronjobs scheduling made easy. Stars:`470`.
- [gronx](https://github.com/adhocore/gronx) - Cron expression parser, task runner and daemon consuming crontab like task list. Stars:`458`.
- [goflow](https://github.com/fieldryand/goflow) - A simple but powerful DAG scheduler and dashboard. Stars:`446`.
- [tasks](https://github.com/madflojo/tasks) - An easy to use in-process scheduler for recurring tasks in Go. Stars:`316`.
- [go-cron](https://github.com/rk/go-cron) - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. Stars:`238`.
- [cheek](https://github.com/bart6114/cheek) - A simple crontab like scheduler that aims to offer a KISS approach to job scheduling. Stars:`194`.
- [clockwerk](https://github.com/onatm/clockwerk) - Go package to schedule periodic jobs using a simple, fluent syntax. Stars:`181`.
- [leprechaun](https://github.com/kilgaloon/leprechaun) - Job scheduler that supports webhooks, crons and classic scheduling. Stars:`106`.
- [cdule](https://github.com/deepaksinghvi/cdule) - Job scheduler library with database support Stars:`56`.
- [sched](https://github.com/romshark/sched) - A job scheduler with the ability to fast-forward time. Stars:`28`.
- [cronticker](https://github.com/krayzpipes/cronticker) - A ticker implementation to support cron schedules. Stars:`18`.

**[⬆ back to top](#contents)**

## JSON

_Libraries for working with JSON._

- [GJSON](https://github.com/tidwall/gjson) - Get a JSON value with one line of code. Stars:`15.1K`.
- [gabs](https://github.com/Jeffail/gabs) - For parsing, creating and editing unknown or dynamic JSON in Go. Stars:`3.5K`.
- [gojson](https://github.com/ChimeraCoder/gojson) - Automatically generate Go (golang) struct definitions from example JSON. Stars:`2.7K`.
- [fastjson](https://github.com/valyala/fastjson) - Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection. Stars:`2.4K`.
- [OjG](https://github.com/ohler55/ojg) - Optimized JSON for Go is a high performance parser with a variety of additional JSON tools including JSONPath. Stars:`907`.
- [jsondiff](https://github.com/wI2L/jsondiff) - JSON diff library for Go based on RFC6902 (JSON Patch). Stars:`592`.
- [marshmallow](https://github.com/PerimeterX/marshmallow) - Performant JSON unmarshalling for flexible use cases. Stars:`386`.
- [kazaam](https://github.com/Qntfy/kazaam) - API for arbitrary transformation of JSON documents. Stars:`286`.
- [ajson](https://github.com/spyzhov/ajson) - Abstract JSON for golang with JSONPath support. Stars:`276`.
- [jsonvalue](https://github.com/Andrew-M-C/go.jsonvalue) - A fast and convenient library for unstructured JSON data, replacing `encoding/json`. Stars:`197`.
- [gojq](https://github.com/elgs/gojq) - JSON query in Golang. Stars:`189`.
- [jettison](https://github.com/wI2L/jettison) - Fast and flexible JSON encoder for Go. Stars:`178`.
- [json2go](https://github.com/m-zajac/json2go) - Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. Stars:`136`.
- [gjo](https://github.com/skanehira/gjo) - Small utility to create JSON objects. Stars:`130`.
- [JayDiff](https://github.com/yazgazan/jaydiff) - JSON diff utility written in Go. Stars:`109`.
- [jsongo](https://github.com/ricardolonga/jsongo) - Fluent API to make it easier to create Json objects. Stars:`108`.
- [jscan](https://github.com/romshark/jscan) - High performance zero-allocation JSON iterator. Stars:`95`.
- [ujson](https://github.com/olvrng/ujson) - Fast and minimal JSON parser and transformer that works on unstructured JSON. Stars:`82`.
- [jsonf](https://github.com/miolini/jsonf) - Console tool for highlighted formatting and struct query fetching JSON. Stars:`65`.
- [go-respond](https://github.com/nicklaw5/go-respond) - Go package for handling common HTTP JSON responses. Stars:`54`.
- [ask](https://github.com/simonnilsson/ask) - Easy access to nested values in maps and slices. Works in combination with encoding/json and other packages that "Unmarshal" arbitrary data into Go data-types. Stars:`53`.
- [mp](https://github.com/sanbornm/mp) - Simple cli email parser. It currently takes stdin and outputs JSON. Stars:`47`.
- [jsoncolor](https://github.com/neilotoole/jsoncolor) - Drop-in replacement for `encoding/json` that outputs colorized JSON. Stars:`47`.
- [vjson](https://github.com/miladibra10/vjson) - Go package for validating JSON objects with declaring a JSON schema with fluent API. Stars:`40`.
- [gojmapr](https://github.com/limiu82214/gojmapr) - Get simple struct from complex json by json path. Stars:`22`.
- [mapslice-json](https://github.com/mickep76/mapslice-json) - Go MapSlice for ordered marshal/ unmarshal of maps in JSON. Stars:`20`.
- [go-jsonerror](https://github.com/ddymko/go-jsonerror) - Go-JsonError is meant to allow us to easily create json response errors that follow the JsonApi spec. Stars:`16`.
- [dynjson](https://github.com/cocoonspace/dynjson) - Client-customizable JSON formats for dynamic APIs. Stars:`16`.
- [epoch](https://github.com/vtopc/epoch) - Contains primitives for marshaling/unmarshalling Unix timestamp/epoch to/from build-in time.Time type in JSON. Stars:`16`.
- [jsonhal](https://github.com/RichardKnop/jsonhal) - Simple Go package to make custom structs marshal into HAL compatible JSON responses. Stars:`14`.
- [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) - Go bindings based on the JSON API errors reference. Stars:`13`.
- [jsonic](https://github.com/sinhashubham95/jsonic) - Utilities to handle and query JSON without defining structs in a type safe manner. Stars:`11`.
- [jzon](https://github.com/zerosnake0/jzon) - JSON library with standard compatible API/behavior. Stars:`11`.
- [ej](https://github.com/lucassscaravelli/ej) - Write and read JSON from different sources succinctly. Stars:`10`.
- [htmljson](https://github.com/nikolaydubina/htmljson) - Rich rendering of JSON as HTML in Go. Stars:`9`.
- [omg.jsonparser](https://github.com/dedalqq/omg.jsonparser) - Simple JSON parser with validation by condition via golang struct fields tags. Stars:`6`.
- [jsonhandlers](https://github.com/abusomani/jsonhandlers) - JSON library to expose simple handlers that lets you easily read and write json from various sources. Stars:`2`.
- [JSON-to-Proto](https://json-to-proto.github.io/) - Convert JSON to Protobuf online.
- [JSON-to-Go](https://mholt.github.io/json-to-go/) - Convert JSON to Go struct.

**[⬆ back to top](#contents)**

## Logging

_Libraries for generating and working with log files._

- [logrus](https://github.com/Sirupsen/logrus) - Structured logger for Go. Stars:`25.4K`.
- [zap](https://github.com/uber-go/zap) - Fast, structured, leveled logging in Go. Stars:`23.4K`.
- [zerolog](https://github.com/rs/zerolog) - Zero-allocation JSON logger. Stars:`11.7K`.
- [spew](https://github.com/davecgh/go-spew) - Implements a deep pretty printer for Go data structures to aid in debugging. Stars:`6.3K`.
- [lumberjack](https://github.com/natefinch/lumberjack) - Simple rolling logger, implements io.WriteCloser. Stars:`5.2K`.
- [glog](https://github.com/golang/glog) - Leveled execution logs for Go. Stars:`3.6K`.
- [tail](https://github.com/hpcloud/tail) - Go package striving to emulate the features of the BSD tail program. Stars:`2.8K`.
- [pp](https://github.com/k0kubun/pp) - Colored pretty printer for Go language. Stars:`2.0K`.
- [seelog](https://github.com/cihub/seelog) - Logging functionality with flexible dispatching, filtering, and formatting. Stars:`1.6K`.
- [log](https://github.com/apex/log) - Structured logging package for Go. Stars:`1.4K`.
- [log15](https://github.com/inconshreveable/log15) - Simple, powerful logging for Go. Stars:`1.1K`.
- [tint](https://github.com/lmittmann/tint) - A slog.Handler that writes tinted logs. Stars:`1.1K`.
- [sentry-go](https://github.com/getsentry/sentry-go) - Sentry SDK for Go. Helps monitor and track errors with real-time alerts and performance monitoring. Stars:`980`.
- [phuslu/log](https://github.com/phuslu/log) - High performance structured logging. Stars:`785`.
- [lazyjournal](https://github.com/Lifailon/lazyjournal) - A TUI for reading and filtering logs from journalctl, file system, Docker and Podman containers, as well Kubernetes pods. Stars:`619`.
- [slog-multi](https://github.com/samber/slog-multi) - Chain of slog.Handler (pipeline, fanout...). Stars:`516`.
- [slog](https://github.com/gookit/slog) - Lightweight, configurable, extensible logger for Go. Stars:`486`.
- [onelog](https://github.com/francoispqt/onelog) - Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenarios. Also, it is one of the logger with the lowest allocation. Stars:`411`.
- [httpretty](https://github.com/henvic/httpretty) - Pretty-prints your regular HTTP requests on your terminal for debugging (similar to http.DumpRequest). Stars:`408`.
- [sqldb-logger](https://github.com/simukti/sqldb-logger) - A logger for Go SQL database driver without modify existing \*sql.DB stdlib usage. Stars:`379`.
- [logutils](https://github.com/hashicorp/logutils) - Utilities for slightly better logging in Go (Golang) extending the standard logger. Stars:`367`.
- [logxi](https://github.com/mgutz/logxi) - 12-factor app logger that is fast and makes you happy. Stars:`358`.
- [rollingwriter](https://github.com/arthurkiller/rollingWriter) - RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. Stars:`298`.
- [log](https://github.com/go-playground/log) - Simple, configurable and scalable Structured Logging for Go. Stars:`295`.
- [go-logger](https://github.com/apsdehal/go-logger) - Simple logger of Go Programs, with level handlers. Stars:`289`.
- [logur](https://github.com/logur/logur) - An opinionated logger interface and collection of logging best practices with adapters and integrations for well-known libraries ([logrus](https://github.com/sirupsen/logrus), [go-kit log](https://github.com/go-kit/kit/tree/master/log), [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog), etc). Stars:`205`.
- [glg](https://github.com/kpango/glg) - glg is simple and fast leveled logging library for Go. Stars:`190`.
- [slog-formatter](https://github.com/samber/slog-formatter) - Common formatters for slog and helpers to build your own. Stars:`180`.
- [logger](https://github.com/azer/logger) - Minimalistic logging library for Go. Stars:`156`.
- [xlog](https://github.com/rs/xlog) - Structured logger for `net/context` aware HTTP handlers with flexible dispatching. Stars:`139`.
- [ozzo-log](https://github.com/go-ozzo/ozzo-log) - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). Stars:`123`.
- [go-cronowriter](https://github.com/utahta/go-cronowriter) - Simple writer that rotate log files automatically based on current date and time, like cronolog. Stars:`57`.
- [gone/log](https://github.com/One-com/gone/tree/master/log) - Fast, extendable, full-featured, std-lib source compatible log library. Stars:`50`.
- [stdlog](https://github.com/alexcesaro/log) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. Stars:`48`.
- [noodlog](https://github.com/gyozatech/noodlog) - Parametrized JSON logging library which lets you obfuscate sensitive data and marshal any kind of content. No more printed pointers instead of values, nor escape chars for the JSON strings. Stars:`44`.
- [journald](https://github.com/ssgreg/journald) - Go implementation of systemd Journal's native API for logging. Stars:`43`.
- [go-log](https://github.com/ian-kent/go-log) - Log4j implementation in Go. Stars:`42`.
- [logex](https://github.com/chzyer/logex) - Golang log lib, supports tracking and level, wrap by standard log lib. Stars:`42`.
- [go-log](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers. Stars:`35`.
- [mlog](https://github.com/jbrodriguez/mlog) - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. Stars:`33`.
- [distillog](https://github.com/amoghe/distillog) - distilled levelled logging (think of it as stdlib + log levels). Stars:`31`.
- [logrusly](https://github.com/sebest/logrusly) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). Stars:`29`.
- [zkits-logger](https://github.com/edoger/zkits-logger) - A powerful zero-dependency JSON logger. Stars:`28`.
- [zax](https://github.com/yuseferi/zax) - Integrate Context with Zap logger, which leads to more flexibility in Go logging. Stars:`27`.
- [log](https://github.com/teris-io/log) - Structured log interface for Go cleanly separates logging facade from its implementation. Stars:`26`.
- [xylog](https://github.com/xybor-x/xylog) - Leveled and structured logging, dynamic fields, high performance, zone management, simple configuration, and readable syntax. Stars:`17`.
- [log](https://github.com/heartwilltell/log) - Simple leveled logging wrapper around standard log package. Stars:`16`.
- [kemba](https://github.com/clok/kemba) - A tiny debug logging tool inspired by [debug](https://github.com/visionmedia/debug), great for CLI tools and applications. Stars:`16`.
- [logrusiowriter](https://github.com/cabify/logrusiowriter) - `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. Stars:`16`.
- [glo](https://github.com/lajosbencz/glo) - PHP Monolog inspired logging facility with identical severity levels. Stars:`16`.
- [go-log](https://github.com/subchen/go-log) - Simple and configurable Logging in Go, with level, formatters and writers. Stars:`14`.
- [logo](https://github.com/mbndr/logo) - Golang logger to different configurable writers. Stars:`12`.
- [go-log](https://github.com/pieterclaerhout/go-log) - A logging library with stack traces, object dumping and optional timestamps. Stars:`11`.
- [logdump](https://github.com/ewwwwwqm/logdump) - Package for multi-level logging. Stars:`11`.
- [log](https://github.com/aerogo/log) - An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). Stars:`10`.
- [xlog](https://github.com/xfxdev/xlog) - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. Stars:`8`.
- [zl](https://github.com/nkmr-jp/zl) - High Developer Experience, zap based logger. It offers rich functionality but is easy to configure. Stars:`8`.
- [structy/log](https://github.com/structy/log) - A simple to use log system, minimalist but with features for debugging and differentiation of messages. Stars:`5`.
- [slf4g](https://github.com/echocat/slf4g) - Simple Logging Facade for Golang: Simple structured logging; but powerful, extendable and customizable, with huge amount of learnings from decades of past logging frameworks. Stars:`4`.
- [log](https://github.com/no-src/log) - A simple logging framework out of the box. Stars:`3`.
- [yell](https://github.com/jfcg/yell) - Yet another minimalistic logging library. Stars:`1`.
- [slogor](https://gitlab.com/greyxor/slogor) - A colorful slog handler.

**[⬆ back to top](#contents)**

## Machine Learning

_Libraries for Machine Learning._

- [GoLearn](https://github.com/sjwhitworth/golearn) - General Machine Learning library for Go. Stars:`9.4K`.
- [gorgonia](https://github.com/gorgonia/gorgonia) - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. Stars:`5.8K`.
- [gosseract](https://github.com/otiai10/gosseract) - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. Stars:`2.9K`.
- [m2cgen](https://github.com/BayesWitnesses/m2cgen) - A CLI tool to transpile trained classic ML models into a native Go code with zero dependencies, written in Python with Go language support. Stars:`2.9K`.
- [tfgo](https://github.com/galeone/tfgo) - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. Stars:`2.5K`.
- [goml](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go. Stars:`1.6K`.
- [eaopt](https://github.com/MaxHalford/eaopt) - An evolutionary optimization library. Stars:`902`.
- [GoMLX](https://github.com/gomlx/gomlx) - An accelerated Machine Learning framework for Go. Stars:`852`.
- [onnx-go](https://github.com/owulveryck/onnx-go) - Go Interface to Open Neural Network Exchange (ONNX). Stars:`815`.
- [bayesian](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang. Stars:`810`.
- [CloudForest](https://github.com/ryanbressler/CloudForest) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. Stars:`744`.
- [ocrserver](https://github.com/otiai10/ocrserver) - A simple OCR API server, seriously easy to be deployed by Docker and Heroku. Stars:`743`.
- [gobrain](https://github.com/goml/gobrain) - Neural Networks written in go. Stars:`567`.
- [go-deep](https://github.com/patrikeh/go-deep) - A feature-rich neural network library in Go. Stars:`554`.
- [hugot](https://github.com/knights-analytics/hugot) - Huggingface transformer pipelines for golang with onnxruntime. Stars:`434`.
- [regommend](https://github.com/muesli/regommend) - Recommendation & collaborative filtering engine. Stars:`313`.
- [Goptuna](https://github.com/c-bata/goptuna) - Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. Stars:`269`.
- [goga](https://github.com/tomcraven/goga) - Genetic algorithm library for Go. Stars:`224`.
- [goRecommend](https://github.com/timkaye11/goRecommend) - Recommendation Algorithms library written in Go. Stars:`206`.
- [go-galib](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang. Stars:`199`.
- [shield](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go. Stars:`158`.
- [go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing) - Fast and convenient feature processing for low latency machine learning in Go. Stars:`124`.
- [go-fann](https://github.com/white-pony/go-fann) - Go bindings for Fast Artificial Neural Networks(FANN) library. Stars:`115`.
- [goscore](https://github.com/asafschers/goscore) - Go Scoring API for PMML. Stars:`102`.
- [GoMind](https://github.com/surenderthakran/gomind) - A simplistic Neural Network Library in Go. Stars:`98`.
- [fonet](https://github.com/Fontinalis/fonet) - A Deep Neural Network library written in Go. Stars:`86`.
- [gonet](https://github.com/dathoangnd/gonet) - Neural Network for Go. Stars:`83`.
- [libsvm](https://github.com/datastream/libsvm) - libsvm golang version derived work based on LIBSVM 3.14. Stars:`73`.
- [neural-go](https://github.com/schuyler/neural-go) - Multilayer perceptron network implemented in Go, with training via backpropagation. Stars:`70`.
- [go-pr](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang. Stars:`67`.
- [Varis](https://github.com/Xamber/Varis) - Golang Neural Network. Stars:`55`.
- [randomforest](https://github.com/malaschitz/randomForest) - Easy to use Random Forest library for Go. Stars:`54`.
- [golinear](https://github.com/danieldk/golinear) - liblinear bindings for Go. Stars:`45`.
- [go-cluster](https://github.com/e-XpertSolutions/go-cluster) - Go implementation of the k-modes and k-prototypes clustering algorithms. Stars:`44`.
- [godist](https://github.com/e-dard/godist) - Various probability distributions, and associated methods. Stars:`43`.
- [ddt](https://github.com/sgrodriguez/ddt) - Dynamic decision tree, create trees defining customizable rules. Stars:`40`.
- [evoli](https://github.com/khezen/evoli) - Genetic Algorithm and Particle Swarm Optimization library. Stars:`34`.
- [probab](https://github.com/ThePaw/probab) - Probability distribution functions. Bayesian inference. Written in pure Go. Stars:`20`.
- [catboost-cgo](https://github.com/mirecl/catboost-cgo) - Fast, scalable, high performance Gradient Boosting on Decision Trees library. Golang using Cgo for blazing fast inference CatBoost Model. Stars:`17`.
- [gorse](https://github.com/zhenghaoz/gorse) - An offline recommender system backend based on collaborative filtering written in Go. Stars:`5`.

**[⬆ back to top](#contents)**

## Messaging

_Libraries that implement messaging systems._

- [Asynq](https://github.com/hibiken/asynq) - A simple, reliable, and efficient distributed task queue for Go built on top of Redis. Stars:`11.7K`.
- [Centrifugo](https://github.com/centrifugal/centrifugo) - Real-time messaging (Websockets or SockJS) server in Go. Stars:`9.2K`.
- [gorush](https://github.com/appleboy/gorush) - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). Stars:`8.4K`.
- [machinery](https://github.com/RichardKnop/machinery) - Asynchronous task queue/job queue based on distributed message passing. Stars:`7.8K`.
- [NATS Go Client](https://github.com/nats-io/nats.go) - Go client for the NATS Stars:`6.1K`.
- [go-socket.io](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework. Stars:`5.8K`.
- [Mercure](https://github.com/dunglas/mercure) - Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). Stars:`5.0K`.
- [Confluent Kafka Golang Client](https://github.com/confluentinc/confluent-kafka-go) - confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform. Stars:`5.0K`.
- [melody](https://github.com/olahol/melody) - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. Stars:`4.0K`.
- [APNs2](https://github.com/sideshow/apns2) - HTTP/2 Apple Push Notification provider for Go - Send push notifications to iOS, tvOS, Safari and OSX apps. Stars:`3.1K`.
- [go-nsq](https://github.com/nsqio/go-nsq) - the official Go package for NSQ. Stars:`2.6K`.
- [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) - gopush-cluster is a go push server cluster. Stars:`2.1K`.
- [EventBus](https://github.com/asaskevich/EventBus) - The lightweight event bus with async compatibility. Stars:`1.9K`.
- [amqp](https://github.com/rabbitmq/amqp091-go) - Go RabbitMQ Client Library. Stars:`1.8K`.
- [Beaver](https://github.com/Clivern/Beaver) - A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. Stars:`1.6K`.
- [Chanify](https://github.com/chanify/chanify) - A push notification server send message to your iOS devices. Stars:`1.3K`.
- [dbus](https://github.com/godbus/dbus) - Native Go bindings for D-Bus. Stars:`1.1K`.
- [Gollum](https://github.com/trivago/gollum) - A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. Stars:`941`.
- [mangos](https://github.com/nanomsg/mangos) - Pure go implementation of the Nanomsg ("Scalability Protocols") with transport interoperability. Stars:`716`.
- [golongpoll](https://github.com/jcuga/golongpoll) - HTTP longpoll server library that makes web pub-sub simple. Stars:`663`.
- [emitter](https://github.com/olebedev/emitter) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. Stars:`524`.
- [Glue](https://github.com/desertbit/glue) - Robust Go and Javascript Socket Library (Alternative to Socket.io). Stars:`420`.
- [Bus](https://github.com/mustafaturan/bus) - Minimalist message bus implementation for internal communication. Stars:`349`.
- [messagebus](https://github.com/vardius/message-bus) - messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. Stars:`278`.
- [Go-MediatR](https://github.com/mehdihadeli/Go-MediatR) - A library for handling mediator patterns and simplified CQRS patterns within an event-driven architecture, inspired by csharp MediatR library. Stars:`249`.
- [guble](https://github.com/smancke/guble) - Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. Stars:`160`.
- [hub](https://github.com/leandro-lugaresi/hub) - A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. Stars:`147`.
- [backlite](https://github.com/mikestefanello/backlite) - Type-safe, persistent, embedded task queues and background job runner w/ SQLite. Stars:`105`.
- [go-mq](https://github.com/cheshir/go-mq) - RabbitMQ client with declarative configuration. Stars:`90`.
- [drone-line](https://github.com/appleboy/drone-line) - Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. Stars:`80`.
- [go-notify](https://github.com/TheCreeper/go-notify) - Native implementation of the freedesktop notification spec. Stars:`71`.
- [Commander](https://github.com/jeroenrinzema/commander) - A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. Stars:`67`.
- [go-res](https://github.com/jirenius/go-res) - Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate. Stars:`66`.
- [event](https://github.com/agoalofalife/event) - Implementation of the pattern observer. Stars:`59`.
- [hare](https://github.com/leozz37/hare) - A user friendly library for sending messages and listening to TCP sockets. Stars:`53`.
- [GoEventBus](https://github.com/Raezil/GoEventBus) - A blazing‑fast, in‑memory, lock‑free event bus library Stars:`35`.
- [ami](https://github.com/kak-tus/ami) - Go client to reliable queues based on Redis Cluster Streams. Stars:`33`.
- [hypermatch](https://github.com/SchwarzIT/hypermatch) - A very fast and efficient Go library for matching events to a large set of rules Stars:`28`.
- [gosd](https://github.com/alexsniffin/gosd) - A library for scheduling when to dispatch a message to a channel. Stars:`26`.
- [go-vitotrol](https://github.com/maxatome/go-vitotrol) - Client library to Viessmann Vitotrol web service. Stars:`23`.
- [jazz](https://github.com/socifi/jazz) - A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. Stars:`18`.
- [gaurun-client](https://github.com/osamingo/gaurun-client) - Gaurun Client written in Go. Stars:`11`.
- [go-eventbus](https://github.com/stanipetrosyan/go-eventbus) - Simple Event Bus package for Go. Stars:`7`.
  messaging system.
- [sarama](https://github.com/Shopify/sarama) - Go library for Apache Kafka. Stars:`12.1K`.
- [Watermill](https://github.com/ThreeDotsLabs/watermill) - Working efficiently with message streams. Building event driven applications, enabling event sourcing, RPC over messages, sagas. Can use conventional pub/sub implementations like Kafka or RabbitMQ, but also HTTP or MySQL binlog. Stars:`8.7K`.
- [Uniqush-Push](https://github.com/uniqush/uniqush-push) - Redis backed unified push service for server-side notifications to mobile devices. Stars:`1.6K`.
- [zmq4](https://github.com/pebbe/zmq4) - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). Stars:`1.2K`.
- [pubsub](https://github.com/tuxychandru/pubsub) - Simple pubsub package for go. Stars:`442`.
- [Quamina](https://github.com/timbray/quamina) - Fast pattern-matching for filtering messages and events. Stars:`432`.
- [rabtap](https://github.com/jandelgado/rabtap) - RabbitMQ swiss army knife cli app. Stars:`270`.
- [redisqueue](https://github.com/robinjoseph08/redisqueue) - redisqueue provides a producer and consumer of a queue that uses Redis streams. Stars:`134`.
- [Ratus](https://github.com/hyperonym/ratus) - Ratus is a RESTful asynchronous task queue server. Stars:`122`.
- [rabbitroutine](https://github.com/furdarius/rabbitroutine) - Lightweight library that handles RabbitMQ auto-reconnect and publishing retries. The library takes into account the need to re-declare entities in RabbitMQ after reconnection. Stars:`112`.
- [oplog](https://github.com/dailymotion/oplog) - Generic oplog/replication system for REST APIs. Stars:`110`.
- [rabbus](https://github.com/rafaeljesus/rabbus) - A tiny wrapper over amqp exchanges and queues. Stars:`98`.
- [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) - A tiny wrapper around NSQ topic and channel. Stars:`80`.
- [RapidMQ](https://github.com/sybrexsys/RapidMQ) - RapidMQ is a lightweight and reliable library for managing of the local messages queue. Stars:`69`.
- [rmqconn](https://github.com/sbabiv/rmqconn) - RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. Stars:`23`.

**[⬆ back to top](#contents)**

## Microsoft Office

- [unioffice](https://github.com/unidoc/unioffice) - Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. Stars:`4.6K`.

### Microsoft Excel

_Libraries for working with Microsoft Excel._

- [excelize](https://github.com/xuri/excelize) - Golang library for reading and writing Microsoft Excel&trade; (XLSX) files. Stars:`19.4K`.
- [xlsx](https://github.com/tealeg/xlsx) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. Stars:`6.0K`.
- [go-excel](https://github.com/szyhf/go-excel) - A simple and light reader to read a relate-db-like excel as a table. Stars:`195`.
- [xlsx](https://github.com/plandem/xlsx) - Fast and safe way to read/update your existing Microsoft Excel files in Go programs. Stars:`176`.
- [exl](https://github.com/go-the-way/exl) - Excel binding to struct written in Go.(Only supports Go1.18+) Stars:`32`.
- [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) - Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. Stars:`21`.

### Microsoft Word

_Libraries for working with Microsoft Word._

- [godocx](https://github.com/gomutex/godocx) - Library for reading and writing Microsoft Word (Docx) files. Stars:`196`.

**[⬆ back to top](#contents)**

## Miscellaneous

### Dependency Injection

_Libraries for working with dependency injection._

- [google/wire](https://github.com/google/wire) - Automated Initialization in Go. Stars:`14.0K`.
- [fx](https://github.com/uber-go/fx) - A dependency injection based application framework for Go (built on top of dig). Stars:`6.7K`.
- [dig](https://github.com/uber-go/dig) - A reflection based dependency injection toolkit for Go. Stars:`4.2K`.
- [do](https://github.com/samber/do) - A dependency injection framework based on Generics. Stars:`2.2K`.
- [GoLobby/Container](https://github.com/golobby/container) - GoLobby Container is a lightweight yet powerful IoC dependency injection container for the Go programming language. Stars:`598`.
- [goioc/di](https://github.com/goioc/di) - Spring-inspired Dependency Injection Container. Stars:`367`.
- [di](https://github.com/goava/di) - A dependency injection container for go programming language. Stars:`238`.
- [kod](https://github.com/go-kod/kod) - A generics based dependency injection framework for Go. Stars:`185`.
- [dingo](https://github.com/i-love-flamingo/dingo) - A dependency injection toolkit for Go, based on Guice. Stars:`185`.
- [alice](https://github.com/magic003/alice) - Additive dependency injection container for Golang. Stars:`50`.
- [gontainer](https://github.com/NVIDIA/gontainer) - A dependency injection service container for Go projects. Stars:`50`.
- [wire](https://github.com/Fs02/wire) - Strict Runtime Dependency Injection for Golang. Stars:`39`.
- [linker](https://github.com/logrange/linker) - A reflection based dependency injection and inversion of control library with components lifecycle support. Stars:`35`.
- [nject](https://github.com/muir/nject) - A type safe, reflective framework for libraries, tests, http endpoints, and service startup. Stars:`30`.
- [componego](https://github.com/componego/componego) - A dependency injection framework based on components, allowing dynamic dependency replacement without duplicating code in tests. Stars:`28`.
- [ore](https://github.com/firasdarwish/ore) - Lightweight, generic & simple dependency injection (DI) container. Stars:`23`.
- [gocontainer](https://github.com/vardius/gocontainer) - Simple Dependency Injection Container. Stars:`21`.
- [gontainer/gontainer](https://github.com/gontainer/gontainer) - A YAML-based Dependency Injection container for GO. It supports dependencies' scopes, and auto-detection of circular dependencies. Gontainer is concurrent-safe. Stars:`16`.
- [autowire](https://github.com/tiendc/autowire) - Dependency injection using Generics and reflection. Stars:`11`.
- [kinit](https://github.com/go-kata/kinit) - Customizable dependency injection container with the global mode, cascade initialization and panic-safe finalization. Stars:`10`.
- [HnH/di](https://github.com/HnH/di) - DI container library that is focused on clean API and flexibility. Stars:`8`.
- [boot-go](http://github.com/boot-go/boot) - Component-based development with dependency injection using reflections for Go developers.
- [cosban/di](https://gitlab.com/cosban/di) - A code generation based dependency injection wiring tool.

**[⬆ back to top](#contents)**

### Project Layout

_**Unofficial** set of patterns for structuring projects._

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout) - Set of common historical and emerging project layout patterns in the Go ecosystem. Note: despite the org-name they do not represent official golang standards, see [this issue](https://github.com/golang-standards/project-layout/issues/117) for more information. Nonetheless, some may find the layout useful. Stars:`53.3K`.
- [go-blueprint](https://github.com/Melkeydev/go-blueprint) - Allows users to spin up a quick Go project using a popular framework. Stars:`8.0K`.
- [ardanlabs/service](https://github.com/ardanlabs/service) - A [starter kit](https://github.com/ardanlabs/service/wiki) for building production grade scalable web service applications. Stars:`3.8K`.
- [goxygen](https://github.com/shpota/goxygen) - Generate a modern Web project with Go and Angular, React, or Vue in seconds. Stars:`3.6K`.
- [pagoda](https://github.com/mikestefanello/pagoda) - Rapid, easy full-stack web development starter kit built in Go. Stars:`2.8K`.
- [nunu](https://github.com/go-nunu/nunu) - Nunu is a scaffolding tool for building Go applications. Stars:`2.3K`.
- [modern-go-application](https://github.com/sagikazarmark/modern-go-application) - Go application boilerplate and example applying modern practices. Stars:`1.9K`.
- [goapp](https://github.com/naughtygopher/goapp) - An opinionated guideline to structure & develop a Go web application/service. Stars:`990`.
- [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) - A Go application boilerplate template for quick starting projects following production best practices. Stars:`720`.
- [go-starter](https://github.com/allaboutapps/go-starter) - An opinionated production-ready RESTful JSON backend template, highly integrated with VSCode DevContainers. Stars:`569`.
- [golang-templates/seed](https://github.com/golang-templates/seed) - Go application GitHub repository template. Stars:`534`.
- [go-todo-backend](https://github.com/Fs02/go-todo-backend) - Go Todo Backend example using modular project layout for product microservice. Stars:`331`.
- [scaffold](https://github.com/catchplay/scaffold) - Scaffold generates a starter Go project layout. Lets you focus on business logic implemented. Stars:`147`.
- [go-sample](https://github.com/zitryss/go-sample) - A sample layout for Go application projects with the real code. Stars:`138`.
- [kickstart.go](https://github.com/raeperd/kickstart.go) - Minimalistic single-file Go HTTP server template without third-party dependencies. Stars:`86`.
- [gobase](https://github.com/wajox/gobase) - A simple skeleton for golang application with basic setup for real golang application. Stars:`66`.
- [go-module](https://github.com/octomation/go-module) - Template for a typical module written on Go. Stars:`35`.
- [wangyoucao577/go-project-layout](https://github.com/wangyoucao577/go-project-layout) - Set of practices and discussions on how to structure Go project layout. Stars:`26`.
- [insidieux/inizio](https://github.com/insidieux/inizio) - Golang project layout generator with plugins. Stars:`18`.

**[⬆ back to top](#contents)**

### Strings

_Libraries for working with strings._

- [xstrings](https://github.com/huandu/xstrings) - Collection of useful string functions ported from other languages. Stars:`1.4K`.
- [sttr](https://github.com/abhimanyu003/sttr) - cross-platform, cli app to perform various operations on string. Stars:`1.2K`.
- [gobeam/Stringy](https://github.com/gobeam/Stringy) - String manipulation library to convert string to camel case, snake case, kebab case / slugify etc. Stars:`250`.
- [strutil](https://github.com/ozgio/strutil) - String utilities. Stars:`207`.
- [caps](https://github.com/chanced/caps) - A case conversion library. Stars:`59`.
- [bexp](https://github.com/happy-sdk/happy/tree/main/pkg/strings/bexp) - Go implementation of Brace Expansion mechanism to generate arbitrary strings. Stars:`21`.
- [strcase](https://github.com/charlievieth/strcase) - Case-insensitive implementation of the standard library's strings/bytes packages. Stars:`8`.
- [go-formatter](https://gitlab.com/tymonx/go-formatter) - Implements **replacement fields** surrounded by curly braces `{}` format strings.

**[⬆ back to top](#contents)**

### Uncategorized

_These libraries were placed here because none of the other categories seemed to fit._

- [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). Stars:`11.4K`.
- [gatus](https://github.com/TwinProduction/gatus) - Automated service health dashboard. Stars:`7.9K`.
- [gofakeit](https://github.com/brianvoe/gofakeit) - Random data generator written in go. Stars:`5.1K`.
- [go-resiliency](https://github.com/eapache/go-resiliency) - Resiliency patterns for golang. Stars:`2.3K`.
- [base64Captcha](https://github.com/mojocn/base64Captcha) - Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. Stars:`2.3K`.
- [gosms](https://github.com/haxpax/gosms) - Your own local SMS gateway in Go that can be used to send SMS. Stars:`1.5K`.
- [shoutrrr](https://github.com/containrrr/shoutrrr) - Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others. Stars:`1.3K`.
- [go-commons-pool](https://github.com/jolestar/go-commons-pool) - Generic object pool for Golang. Stars:`1.2K`.
- [llvm](https://github.com/llir/llvm) - Library for interacting with LLVM IR in pure Go. Stars:`1.2K`.
- [stateless](https://github.com/qmuntal/stateless) - A fluent library for creating state machines. Stars:`1.1K`.
- [shortid](https://github.com/teris-io/shortid) - Distributed generation of super short, unique, non-sequential, URL friendly IDs. Stars:`950`.
- [health](https://github.com/alexliesenfeld/health) - A simple and flexible health check library for Go. Stars:`814`.
- [xz](https://github.com/ulikunitz/xz) - Pure golang package for reading and writing xz-compressed files. Stars:`519`.
- [banner](https://github.com/dimiro1/banner) - Add beautiful banners into your Go applications. Stars:`457`.
- [health](https://github.com/dimiro1/health) - Easy to use, extensible health check library. Stars:`450`.
- [gountries](https://github.com/pariz/gountries) - Package that exposes country and subdivision data. Stars:`419`.
- [lk](https://github.com/hyperboloide/lk) - A simple licensing library for golang. Stars:`390`.
- [conv](https://github.com/cstockton/go-conv) - Package conv provides fast and intuitive conversions across Go types. Stars:`381`.
- [ffmt](https://github.com/go-ffmt/ffmt) - Beautify data display for Humans. Stars:`312`.
- [gtree](https://github.com/ddddddO/gtree) - Provide CLI, Package and Web for tree output and directories creation from Markdown or programmatically. Stars:`304`.
- [go-unarr](https://github.com/gen2brain/go-unarr) - Decompression library for RAR, TAR, ZIP and 7z archives. Stars:`297`.
- [archives](https://github.com/mholt/archives) - a cross-platform, multi-format Go library for working with archives and compression formats with a unified API and as virtual file systems compatible with io/fs. Stars:`280`.
- [healthcheck](https://github.com/etherlabsio/healthcheck) - An opinionated and concurrent health-check HTTP handler for RESTful services. Stars:`274`.
- [antch](https://github.com/antchfx/antch) - A fast, powerful and extensible web crawling & scraping framework. Stars:`263`.
- [battery](https://github.com/distatus/battery) - Cross-platform, normalized battery information library. Stars:`262`.
- [bitio](https://github.com/icza/bitio) - Highly optimized bit-level Reader and Writer for Go. Stars:`250`.
- [stats](https://github.com/go-playground/stats) - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... Stars:`172`.
- [turtle](https://github.com/hackebrot/turtle) - Emojis for Go. Stars:`166`.
- [captcha](https://github.com/steambap/captcha) - Package captcha provides an easy to use, unopinionated API for captcha generation. Stars:`160`.
- [gommit](https://github.com/antham/gommit) - Analyze git commit messages to ensure they follow defined patterns. Stars:`115`.
- [indigo](https://github.com/osamingo/indigo) - Distributed unique ID generator of using Sonyflake and encoded by Base58. Stars:`111`.
- [gotoprom](https://github.com/cabify/gotoprom) - Type-safe metrics builder wrapper library for the official Prometheus client. Stars:`108`.
- [faker](https://github.com/pioz/faker) - Random fake data and struct generator for Go. Stars:`100`.
- [persian](https://github.com/mavihq/persian) - Some utilities for Persian language in go. Stars:`91`.
- [morse](https://github.com/alwindoss/morse) - Library to convert to and from morse code. Stars:`85`.
- [pdfgen](https://github.com/hyperboloide/pdfgen) - HTTP service to generate PDF from Json requests. Stars:`75`.
- [xkg](https://github.com/go-xkg/xkg) - X Keyboard Grabber. Stars:`61`.
- [url-shortener](https://github.com/pantrif/url-shortener) - A modern, powerful, and robust URL shortener microservice with mysql support. Stars:`50`.
- [browscap_go](https://github.com/digitalcrab/browscap_go) - GoLang Library for [Browser Capabilities Project](https://browscap.org/). Stars:`49`.
- [datacounter](https://github.com/miolini/datacounter) - Go counters for readers/writer/http.ResponseWriter. Stars:`49`.
- [xdg](https://github.com/rkoesters/xdg) - FreeDesktop.org (xdg) Specs implemented in Go. Stars:`48`.
- [sandid](https://github.com/aofei/sandid) - Every grain of sand on earth has its own ID. Stars:`47`.
- [autoflags](https://github.com/artyom/autoflags) - Go package to automatically define command line flags from struct fields. Stars:`43`.
- [numa](https://github.com/lrita/numa) - NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. Stars:`35`.
- [gosh](https://github.com/osamingo/gosh) - Provide Go Statistics Handler, Struct, Measure Method. Stars:`35`.
- [metrics](https://github.com/pascaldekloe/metrics) - Library for metrics instrumentation and Prometheus exposition. Stars:`27`.
- [shellwords](https://github.com/Wing924/shellwords) - A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. Stars:`26`.
- [avgRating](https://github.com/kirillDanshin/avgRating) - Calculate average score and rating based on Wilson Score Equation. Stars:`17`.
- [anagent](https://github.com/mudler/anagent) - Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. Stars:`16`.
- [varint](https://github.com/chmike/varint) - A faster varying length integer encoder/decoder than the one provided in the standard library. Stars:`15`.
- [go-commandbus](https://github.com/lana/go-commandbus) - A slight and pluggable command-bus for Go. Stars:`15`.
- [hostutils](https://github.com/Wing924/hostutils) - A golang library for packing and unpacking FQDNs list. Stars:`13`.
- [openapi](https://github.com/neotoolkit/openapi) - OpenAPI 3.x parser. Stars:`11`.
- [sitemap-format](https://github.com/mingard/sitemap-format) - A simple sitemap generator, with a little syntactic sugar. Stars:`6`.
- [VarHandler](https://github.com/azr/generators/tree/master/varhandler) - Generate boilerplate http input and output handling. Stars:`6`.
- [basexx](https://github.com/bobg/basexx) - Convert to, from, and between digit strings in various number bases. Stars:`5`.
- [common](https://github.com/kubeservice-stack/common) - A library for server framework. Stars:`4`.
- [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.

**[⬆ back to top](#contents)**

## Natural Language Processing

_Libraries for working with human languages._

See also [Text Processing](#text-processing) and [Text Analysis](#text-analysis).

### Language Detection

- [lingua-go](https://github.com/pemistahl/lingua-go) - An accurate natural language detection library, suitable for long and short text alike. Supports detecting multiple languages in mixed-language text. Stars:`1.3K`.
- [whatlanggo](https://github.com/abadojack/whatlanggo) - Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). Stars:`664`.
- [getlang](https://github.com/rylans/getlang) - Fast natural language detection package. Stars:`174`.
- [guesslanguage](https://github.com/endeveit/guesslanguage) - Functions to determine the natural language of a unicode text. Stars:`58`.
- [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) - Language Detection API Go Client. Supports batch requests, short phrase or single word language detection. Stars:`26`.

### Morphological Analyzers

- [spaGO](https://github.com/nlpodyssey/spago) - Self-contained Machine Learning and Natural Language Processing library in Go. Stars:`1.8K`.
- [kagome](https://github.com/ikawaha/kagome) - JP morphological analyzer written in pure Go. Stars:`888`.
- [nlp](https://github.com/james-bowman/nlp) - Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). Stars:`459`.
- [nlp](https://github.com/Shixzie/nlp) - Extract values from strings and fill your structs with nlp. Stars:`389`.
- [RAKE.go](https://github.com/afjoseph/RAKE.Go) - Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). Stars:`120`.
- [go-stem](https://github.com/agonopol/go-stem) - Implementation of the porter stemming algorithm. Stars:`82`.
- [go2vec](https://github.com/danieldk/go2vec) - Reader and utility functions for word2vec embeddings. Stars:`56`.
- [govader](https://github.com/jonreiter/govader) - Go implementation of [VADER Sentiment Analysis](https://github.com/cjhutto/vaderSentiment). Stars:`50`.
- [porter2](https://github.com/zhenjl/porter2) - Really fast Porter 2 stemmer. Stars:`46`.
- [snowball](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). Stars:`38`.
- [paicehusk](https://github.com/rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm. Stars:`29`.
- [golibstemmer](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2. Stars:`21`.
- [libtextcat](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. Stars:`13`.
- [porter](https://github.com/a2800276/porter) - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. Stars:`12`.
- [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) - Sentiment analyzer using sentiwordnet lexicon in Go. Stars:`11`.
- [govader-backend](https://github.com/PIMPfiction/govader_backend) - Microservice implementation of [GoVader](https://github.com/jonreiter/govader). Stars:`6`.
- [spelling-corrector](https://github.com/jorelosorio/spellingcorrector) - A spelling corrector for the Spanish language or create your own. Stars:`2`.

### Slugifiers

- [slug](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support. Stars:`1.3K`.
- [go-slugify](https://github.com/mozillazg/go-slugify) - Make pretty slug with multiple languages support. Stars:`97`.
- [Slugify](https://github.com/avelino/slugify) - Go slugify application that handles string. Stars:`34`.

### Tokenizers

- [prose](https://github.com/jdkato/prose) - Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. Stars:`3.1K`.
- [gse](https://github.com/go-ego/gse) - Go efficient text segmentation; support english, chinese, japanese and other. Stars:`2.7K`.
- [gojieba](https://github.com/yanyiwu/gojieba) - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. Stars:`2.5K`.
- [sentences](https://github.com/neurosnap/sentences) - Sentence tokenizer: converts text into a list of sentences. Stars:`453`.
- [segment](https://github.com/blevesearch/segment) - Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](https://www.unicode.org/reports/tr29/) Stars:`89`.
- [textcat](https://github.com/pebbe/textcat) - Go package for n-gram based text categorization, with support for utf-8 and raw text. Stars:`73`.
- [MMSEGO](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. Stars:`62`.
- [stemmer](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers. Stars:`53`.
- [gotokenizer](https://github.com/xujiajun/gotokenizer) - A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation) Stars:`21`.
- [shamoji](https://github.com/osamingo/shamoji) - The shamoji is word filtering package written in Go. Stars:`13`.

### Translation

- [go-i18n](https://github.com/nicksnyder/go-i18n/) - Package and an accompanying tool to work with localized text. Stars:`3.3K`.
- [go-pinyin](https://github.com/mozillazg/go-pinyin) - CN Hanzi to Hanyu Pinyin converter. Stars:`1.7K`.
- [gotext](https://github.com/leonelquinteros/gotext) - GNU gettext utilities for Go. Stars:`473`.
- [ctxi18n](https://github.com/invopop/ctxi18n/) - Context aware i18n with a short and consise API, pluralization, interpolation, and `fs.FS` support. YAML locale definitions are based on [Rails i18n](https://guides.rubyonrails.org/i18n.html). Stars:`74`.
- [spreak](https://github.com/vorlif/spreak) - Flexible translation and humanization library for Go, based on the concepts behind gettext. Stars:`64`.
- [iuliia-go](https://github.com/mehanizm/iuliia-go) - Transliterate Cyrillic → Latin in every possible way. Stars:`54`.
- [go-mystem](https://github.com/dveselov/mystem) - CGo bindings to Yandex.Mystem - russian morphology analyzer. Stars:`34`.
- [t](https://github.com/youthlin/t) - Another i18n pkg for golang, which follows GNU gettext style and supports .po/.mo files: `t.T (gettext)`, `t.N (ngettext)`, etc. And it contains a cmd tool [xtemplate](https://github.com/youthlin/t/blob/main/cmd/xtemplate), which can extract messages as a pot file from text/html template. Stars:`20`.
- [go-words](https://github.com/saleh-rahimzadeh/go-words) - A words table and text resource library for Golang projects. Stars:`7`.

### Transliteration

- [go-unidecode](https://github.com/mozillazg/go-unidecode) - ASCII transliterations of Unicode text. Stars:`141`.
- [gounidecode](https://github.com/fiam/gounidecode) - Unicode transliterator (also known as unidecode) for Go. Stars:`80`.
- [transliterator](https://github.com/alexsergivan/transliterator) - Provides one-way string transliteration with supporting of language-specific transliteration rules. Stars:`46`.
- [enca](https://github.com/endeveit/enca) - Minimal cgo bindings for [libenca](https://cihar.com/software/enca/), which detects character encodings. Stars:`18`.

**[⬆ back to top](#contents)**

## Networking

_Libraries for working with various layers of the network._

- [fasthttp](https://github.com/valyala/fasthttp) - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. Stars:`22.8K`.
- [webrtc](https://github.com/pions/webrtc) - A pure Go implementation of the WebRTC API. Stars:`15.1K`.
- [kcptun](https://github.com/xtaci/kcptun) - Extremely simple & fast udp tunnel based on KCP protocol. Stars:`14.2K`.
- [cloudflared](https://github.com/cloudflare/cloudflared) - Cloudflare Tunnel client (formerly Argo Tunnel). Stars:`11.2K`.
- [quic-go](https://github.com/lucas-clemente/quic-go) - An implementation of the QUIC protocol in pure Go. Stars:`10.9K`.
- [gnet](https://github.com/panjf2000/gnet) - `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. Stars:`10.7K`.
- [dns](https://github.com/miekg/dns) - Go library for working with DNS. Stars:`8.4K`.
- [gopacket](https://github.com/google/gopacket) - Go library for packet processing with libpcap bindings. Stars:`6.6K`.
- [GoProxy](https://github.com/elazarl/goproxy) - A library to create a customized HTTP/HTTPS proxy server using Go. Stars:`6.4K`.
- [netpoll](https://github.com/cloudwego/netpoll) - A high-performance non-blocking I/O networking framework, which focused on RPC scenarios, developed by ByteDance. Stars:`4.4K`.
- [kcp-go](https://github.com/xtaci/kcp-go) - KCP - Fast and Reliable ARQ Protocol. Stars:`4.3K`.
- [HTTPLab](https://github.com/gchaincl/httplab) - HTTPLabs let you inspect HTTP requests and forge responses. Stars:`4.1K`.
- [tun2socks](https://github.com/xjasonlyu/tun2socks) - A pure go implementation of tun2socks powered by [gVisor](https://gvisor.dev/) TCP/IP stack. Stars:`4.0K`.
- [ssh](https://github.com/gliderlabs/ssh) - Higher-level API for building SSH servers (wraps crypto/ssh). Stars:`4.0K`.
- [gobgp](https://github.com/osrg/gobgp) - BGP implemented in the Go Programming Language. Stars:`3.8K`.
- [fortio](https://github.com/fortio/fortio) - Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. Stars:`3.6K`.
- [nbio](https://github.com/lesismal/nbio) - Pure Go 1000k+ connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use. Stars:`2.5K`.
- [water](https://github.com/songgao/water) - Simple TUN/TAP library. Stars:`2.1K`.
- [gev](https://github.com/Allenxuxu/gev) - gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. Stars:`1.8K`.
- [go-getter](https://github.com/hashicorp/go-getter) - Go library for downloading files or directories from various sources using a URL. Stars:`1.7K`.
- [sftp](https://github.com/pkg/sftp) - Package sftp implements the SSH File Transfer Protocol as described in <https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt>. Stars:`1.6K`.
- [gws](https://github.com/lxzan/gws) - High-Performance WebSocket Server & Client With AsyncIO Supporting . Stars:`1.6K`.
- [grab](https://github.com/cavaliercoder/grab) - Go package for managing file downloads. Stars:`1.5K`.
- [NFF-Go](https://github.com/intel-go/nff-go) - Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). Stars:`1.4K`.
- [ftp](https://github.com/jlaffaye/ftp) - Package ftp implements a FTP client as described in [RFC 959](https://tools.ietf.org/html/rfc959). Stars:`1.4K`.
- [mdns](https://github.com/hashicorp/mdns) - Simple mDNS (Multicast DNS) client/server library in Golang. Stars:`1.3K`.
- [gosnmp](https://github.com/soniah/gosnmp) - Native Go library for performing SNMP actions. Stars:`1.2K`.
- [nodepass](https://github.com/yosebyte/nodepass) - A secure, efficient TCP/UDP tunneling solution that delivers fast, reliable access across network restrictions using pre-established TLS/TCP connections. Stars:`1.2K`.
- [gmqtt](https://github.com/DrmagicE/gmqtt) - Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. Stars:`1.0K`.
- [sdns](https://github.com/semihalev/sdns) - A high-performance, recursive DNS resolver server with DNSSEC support, focused on preserving privacy. Stars:`1.0K`.
- [vssh](https://github.com/yahoo/vssh) - Go library for building network and server automation over SSH protocol. Stars:`974`.
- [cidranger](https://github.com/yl2chen/cidranger) - Fast IP to CIDR lookup for Go. Stars:`947`.
- [easytcp](https://github.com/DarthPestilane/easytcp) - A light-weight TCP framework written in Go (Golang), built with message router. EasyTCP helps you build a TCP server easily fast and less painful. Stars:`823`.
- [gaio](https://github.com/xtaci/gaio) - High performance async-io networking for Golang in proactor mode. Stars:`794`.
- [go-stun](https://github.com/ccding/go-stun) - Go implementation of the STUN client (RFC 3489 and RFC 5389). Stars:`710`.
- [lhttp](https://github.com/fanux/lhttp) - Powerful websocket framework, build your IM server more easily. Stars:`690`.
- [peerdiscovery](https://github.com/schollz/peerdiscovery) - Pure Go library for cross-platform local peer discovery using UDP multicast. Stars:`660`.
- [gotcp](https://github.com/gansidui/gotcp) - Go package for quickly writing tcp applications. Stars:`515`.
- [stun](https://github.com/go-rtc/stun) - Go implementation of RFC 5389 STUN protocol. Stars:`494`.
- [gopcap](https://github.com/akrennmair/gopcap) - Go wrapper for libpcap. Stars:`492`.
- [ftpserverlib](https://github.com/fclairamb/ftpserverlib) - Fully featured FTP server library. Stars:`453`.
- [winrm](https://github.com/masterzen/winrm) - Go WinRM client to remotely execute commands on Windows machines. Stars:`446`.
- [raw](https://github.com/mdlayher/raw) - Package raw enables reading and writing data at the device driver level for a network interface. Stars:`427`.
- [arp](https://github.com/mdlayher/arp) - Package arp implements the ARP protocol, as described in RFC 826. Stars:`375`.
- [dnsmonster](https://github.com/mosajjal/dnsmonster) - Passive DNS Capture/Monitoring Framework. Stars:`334`.
- [ethernet](https://github.com/mdlayher/ethernet) - Package ethernet implements marshaling and unmarshalling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. Stars:`283`.
- [gNxI](https://github.com/google/gnxi) - A collection of tools for Network Management that use the gNMI and gNOI protocols. Stars:`277`.
- [buffstreams](https://github.com/stabbycutyou/buffstreams) - Streaming protocolbuffer data over TCP made easy. Stars:`254`.
- [jazigo](https://github.com/udhos/jazigo) - Jazigo is a tool written in Go for retrieving configuration for multiple network devices. Stars:`222`.
- [utp](https://github.com/anacrolix/utp) - Go uTP micro transport protocol implementation. Stars:`178`.
- [tcpack](https://github.com/lim-yoona/tcpack) - tcpack is an application protocol based on TCP to Pack and Unpack bytes stream in go program. Stars:`171`.
- [xtcp](https://github.com/xfxdev/xtcp) - TCP Server Framework with simultaneous full duplex communication, graceful shutdown, and custom protocol. Stars:`158`.
- [canopus](https://github.com/zubairhamed/canopus) - CoAP Client/Server implementation (RFC 7252). Stars:`156`.
- [psql-wire](https://github.com/jeroenrinzema/psql-wire) - PostgreSQL server wire protocol. Build your own server and start serving connections.. Stars:`153`.
- [sslb](https://github.com/eduardonunesp/sslb) - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. Stars:`151`.
- [iplib](https://github.com/c-robinson/iplib) - Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html) Stars:`149`.
- [event](https://github.com/cheng-zhongliang/event) - Simple I/O event notification library written in Golang. Stars:`120`.
- [gldap](https://github.com/jimlambrt/gldap) - gldap provides an ldap server implementation and you provide handlers for its ldap operations. Stars:`116`.
- [natiu-mqtt](https://github.com/soypat/natiu-mqtt) - A dead-simple, non-allocating, low level implementation of MQTT well suited for embedded systems. Stars:`97`.
- [go-powerdns](https://github.com/joeig/go-powerdns) - PowerDNS API bindings for Golang. Stars:`97`.
- [fullproxy](https://github.com/shoriwe/fullproxy) - A fully featured scriptable and daemon configurable proxy and pivoting toolkit with SOCKS5, HTTP, raw ports and reverse proxy protocols. Stars:`82`.
- [ether](https://github.com/songgao/ether) - Cross-platform Go package for sending and receiving ethernet frames. Stars:`80`.
- [dhcp6](https://github.com/mdlayher/dhcp6) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. Stars:`79`.
- [bart](https://github.com/gaissmai/bart) - Package bart provides a fast routing table algorithm. Stars:`78`.
- [fwdctl](https://github.com/alegrey91/fwdctl) - A simple and intuitive CLI to manage IPTables forwards in your Linux server. Stars:`67`.
- [portproxy](https://github.com/aybabtme/portproxy) - Simple TCP proxy which adds CORS support to API's which don't support it. Stars:`58`.
- [linkio](https://github.com/ian-kent/linkio) - Network link speed simulation for Reader/Writer interfaces. Stars:`54`.
- [httpproxy](https://github.com/wzshiming/httpproxy) - HTTP proxy handler and dialer. Stars:`31`.
- [publicip](https://github.com/polera/publicip) - Package publicip returns your public facing IPv4 address (internet egress). Stars:`29`.
- [graval](https://github.com/koofr/graval) - Experimental FTP server framework. Stars:`28`.
- [gnet](https://github.com/fish-tennis/gnet) - `gnet` is a high-performance networking framework,especially for game servers. Stars:`22`.
- [goshark](https://github.com/sunwxg/goshark) - Package goshark use tshark to decode IP packet and create data struct to analyse packet. Stars:`19`.
- [llb](https://github.com/kirillDanshin/llb) - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. Stars:`16`.
- [tspool](https://github.com/two/tspool) - A TCP Library use worker pool to improve performance and protect your server. Stars:`14`.
- [go-multiproxy](https://github.com/presbrey/go-multiproxy) - Library for making HTTP requests through a pool of proxies offering fault tolerance, load balancing, automatic retries, cookie management, and more, via http.Get/Post replacement or http.Client RoundTripper drop-in Stars:`14`.
- [go-sse](https://github.com/lampctl/go-sse) - Go client and server implementation of HTML server-sent events. Stars:`14`.
- [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
- [net](https://golang.org/x/net) - This repository holds supplementary Go networking libraries.

**[⬆ back to top](#contents)**

### HTTP Clients

_Libraries for making HTTP requests._

- [resty](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client. Stars:`11.1K`.
- [req](https://github.com/imroc/req) - Simple Go HTTP client with Black Magic (Less code and More efficiency). Stars:`4.6K`.
- [heimdall](https://github.com/gojektech/heimdall) - An enhanced http client with retry and hystrix capabilities. Stars:`2.7K`.
- [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) - Retryable HTTP client in Go. Stars:`2.2K`.
- [grequests](https://github.com/levigross/grequests) - A Go "clone" of the great and famous Requests library. Stars:`2.2K`.
- [sling](https://github.com/dghubble/sling) - Sling is a Go HTTP client library for creating and sending API requests. Stars:`1.7K`.
- [requests](https://github.com/carlmjohnson/requests) - HTTP requests for Gophers. Uses context.Context and doesn't hide the underlying net/http.Client, making it compatible with standard Go APIs. Also includes testing tools. Stars:`1.6K`.
- [tls-client](https://github.com/bogdanfinn/tls-client) - net/http.Client like HTTP Client with options to select specific client TLS Fingerprints to use for requests. Stars:`1.2K`.
- [gentleman](https://github.com/h2non/gentleman) - Full-featured plugin-driven HTTP client library. Stars:`1.1K`.
- [pester](https://github.com/sethgrid/pester) - Go HTTP client calls with retries, backoff, and concurrency. Stars:`652`.
- [go-cleanhttp](https://github.com/hashicorp/go-cleanhttp) - Get easily stdlib HTTP client, which does not share any state with other clients. Stars:`404`.
- [azuretls-client](https://github.com/Noooste/azuretls-client) -  An easy-to-use HTTP client 100% in Go to spoof TLS/JA3 and HTTP2 fingerprint Stars:`297`.
- [request](https://github.com/monaco-io/request) - HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency. Stars:`293`.
- [fast-shot](https://github.com/opus-domini/fast-shot) - Hit your API targets with rapid-fire precision using Go's fastest and simple HTTP Client. Stars:`87`.
- [go-otelroundtripper](https://github.com/NdoleStudio/go-otelroundtripper) - Go http.RoundTripper that emits open telemetry metrics for HTTP requests. Stars:`85`.
- [go-http-client](https://github.com/bozd4g/go-http-client) - Make http calls simply and easily. Stars:`83`.
- [go-zoox/fetch](https://github.com/go-zoox/fetch) - A Powerful, Lightweight, Easy Http Client, inspired by Web Fetch API. Stars:`81`.
- [rq](https://github.com/ddo/rq) - A nicer interface for golang stdlib HTTP client. Stars:`53`.
- [httpretry](https://github.com/ybbus/httpretry) - Enriches the default go HTTP client with retry functionality. Stars:`52`.
- [axios4go](https://github.com/rezmoss/axios4go) - A Go HTTP client library inspired by Axios, providing a simple and intuitive API for making HTTP requests. Stars:`28`.
- [go-ipmux](https://github.com/optimus-hft/go-ipmux) - A library for Multiplexing HTTP requests based on multiple Source IPs. Stars:`27`.
- [go-req](https://github.com/wenerme/go-req) - Declarative golang HTTP client. Stars:`22`.

**[⬆ back to top](#contents)**

## OpenGL

_Libraries for using OpenGL in Go._

- [glfw](https://github.com/go-gl/glfw) - Go bindings for GLFW 3. Stars:`1.6K`.
- [gl](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow). Stars:`1.2K`.
- [mathgl](https://github.com/go-gl/mathgl) - Pure Go math package specialized for 3D math, with inspiration from GLM. Stars:`582`.
- [goxjs/gl](https://github.com/goxjs/gl) - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). Stars:`178`.
- [goxjs/glfw](https://github.com/goxjs/glfw) - Go cross-platform glfw library for creating an OpenGL context and receiving events. Stars:`83`.
- [go-glmatrix](https://github.com/technohippy/go-glmatrix) - Go port of [glMatrix](https://glmatrix.net/) library. Stars:`11`.

**[⬆ back to top](#contents)**

## ORM

_Libraries that implement Object-Relational Mapping or datamapping techniques._

- [GORM](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang, aims to be developer friendly. Stars:`38.6K`.
- [ent](https://github.com/facebook/ent) - An entity framework for Go. Simple, yet powerful ORM for modeling and querying data. Stars:`16.4K`.
- [SQLBoiler](https://github.com/volatiletech/sqlboiler) - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. Stars:`6.9K`.
- [bun](https://github.com/uptrace/bun) - SQL-first Golang ORM. Successor of go-pg. Stars:`4.3K`.
- [gorp](https://github.com/go-gorp/gorp) - Go Relational Persistence, ORM-ish library for Go. Stars:`3.7K`.
- [upper.io/db](https://github.com/upper/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. Stars:`3.6K`.
- [gormt](https://github.com/xxjwxc/gormt) - Mysql database to golang gorm struct. Stars:`2.4K`.
- [Prisma](https://github.com/prisma/prisma-client-go) - Prisma Client Go, Typesafe database access for Go. Stars:`2.3K`.
- [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) - A flexible and powerful SQL string builder library plus a zero-config ORM. Stars:`1.6K`.
- [pop/soda](https://github.com/gobuffalo/pop) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. Stars:`1.5K`.
- [reform](https://github.com/go-reform/reform) - Better ORM for Go, based on non-empty interfaces and code generation. Stars:`1.4K`.
- [bob](https://github.com/stephenafamo/bob) - SQL query builder and ORM/Factory generator for Go. Successor of SQLBoiler. Stars:`1.4K`.
- [rel](https://github.com/go-rel/rel) - Modern Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API. Stars:`783`.
- [Zoom](https://github.com/albrow/zoom) - Blazing-fast datastore and querying engine built on Redis. Stars:`313`.
- [go-sql](https://github.com/rushteam/gosql) - A easy ORM for mysql. Stars:`180`.
- [grimoire](https://github.com/Fs02/grimoire) - Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). Stars:`163`.
- [golobby/orm](https://github.com/golobby/orm) - Simple, fast, type-safe, generic orm for developer happiness. Stars:`160`.
- [go-store](https://github.com/gosuri/go-store) - Simple and fast Redis backed key-value store library for Go. Stars:`113`.
- [go-firestorm](https://github.com/jschoedt/go-firestorm) - A simple ORM for Google/Firebase Cloud Firestore. Stars:`52`.
- [cacheme](https://github.com/Yiling-J/cacheme-go) - Schema based, typed Redis caching/memoize framework for Go. Stars:`24`.
- [CQL](https://github.com/FrancoLiberali/cql) - Built on top of GORM, adds compile-time verified queries based on auto-generated code. Stars:`17`.
- [go-dbw](https://github.com/hashicorp/go-dbw) - A simple package that encapsulates database operations. Stars:`17`.
- [lore](https://github.com/abrahambotros/lore) - Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. Stars:`14`.
- [marlow](https://github.com/marlow/marlow) - Generated ORM from project structs for compile time safety assurances. Stars:`12`.
- [XORM](https://gitea.com/xorm/xorm) - Simple and powerful ORM for Go. (Support: MySQL, MyMysql, PostgreSQL, Tidb, SQLite3, MsSql and Oracle).

**[⬆ back to top](#contents)**

## Package Management

_Official tooling for dependency and package management_

- [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) - Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

_Unofficial libraries for package and dependency management._

- [glide](https://github.com/Masterminds/glide) - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. Stars:`8.1K`.
- [syft](https://github.com/anchore/syft) - A CLI tool and Go library for generating a Software Bill of Materials (SBOM) from container images and filesystems. Stars:`7.4K`.
- [godep](https://github.com/tools/godep) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. Stars:`5.5K`.
- [govendor](https://github.com/kardianos/govendor) - Go Package Manager. Go vendor tool that works with the standard vendor file. Stars:`4.9K`.
- [gopm](https://github.com/gpmgo/gopm) - Go Package Manager. Stars:`2.5K`.
- [gpm](https://github.com/pote/gpm) - Barebones dependency manager for Go. Stars:`1.2K`.
- [goop](https://github.com/nitrous-io/goop) - Simple dependency manager for Go (golang), inspired by Bundler. Stars:`775`.
- [modgv](https://github.com/lucasepe/modgv) - Converts 'go mod graph' output into Graphviz's DOT language. Stars:`493`.
- [gup](https://github.com/nao1215/gup) - Update binaries installed by "go install". Stars:`431`.
- [johnny-deps](https://github.com/VividCortex/johnny-deps) - Minimal dependency version using Git. Stars:`213`.
- [mvn-golang](https://github.com/raydac/mvn-golang) - plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. Stars:`165`.
- [VenGO](https://github.com/DamnWidget/VenGO) - create and manage exportable isolated go virtual environments. Stars:`126`.
- [gop](https://github.com/lunny/gop) - Build and manage your Go applications out of GOPATH. Stars:`48`.

**[⬆ back to top](#contents)**

## Performance

- [jaeger](https://github.com/jaegertracing/jaeger) - A distributed tracing system. Stars:`21.7K`.
- [pixie](https://github.com/pixie-labs/pixie) - No instrumentation tracing for Golang applications via eBPF. Stars:`6.1K`.
- [statsviz](https://github.com/arl/statsviz) - Live visualization of your Go application runtime statistics. Stars:`3.3K`.
- [profile](https://github.com/pkg/profile) - Simple profiling support package for Go. Stars:`2.0K`.
- [go-instrument](https://github.com/nikolaydubina/go-instrument) - Automatically add spans to all methods and functions. Stars:`272`.
- [mm-go](https://github.com/joetifa2003/mm-go) - Generic manual memory management for golang. Stars:`175`.
- [tracer](https://github.com/kamilsk/tracer) - Simple, lightweight tracing. Stars:`89`.

**[⬆ back to top](#contents)**

## Query Language

- [gqlgen](https://github.com/99designs/gqlgen) - go generate based graphql server library. Stars:`10.4K`.
- [graphql-go](https://github.com/graphql-go/graphql) - Implementation of GraphQL for Go. Stars:`10.1K`.
- [dasel](https://github.com/tomwright/dasel) - Query and update data structures using selectors from the command line. Comparable to jq/yq but supports JSON, YAML, TOML and XML with zero runtime dependencies. Stars:`7.5K`.
- [graphql](https://github.com/neelance/graphql-go) - GraphQL server with a focus on ease of use. Stars:`4.7K`.
- [gojsonq](https://github.com/thedevsaddam/gojsonq) - A simple Go package to Query over JSON Data. Stars:`2.2K`.
- [rql](https://github.com/a8m/rql) - Resource Query Language for REST API. Stars:`359`.
- [jsonql](https://github.com/elgs/jsonql) - JSON query expression library in Golang. Stars:`278`.
- [jsonslice](https://github.com/bhmj/jsonslice) - Jsonpath queries with advanced filters. Stars:`92`.
- [rqp](https://github.com/timsolov/rest-query-parser) - Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query. Stars:`85`.
- [mql](https://github.com/hashicorp/mql) - Model Query Language (mql) is a query language for your database models. Stars:`64`.
- [goven](https://github.com/SeldonIO/goven) - A drop-in query language for any database schema. Stars:`62`.
- [graphql](https://github.com/tmc/graphql) - graphql parser + utilities. Stars:`57`.
- [api-fu](https://github.com/ccbrown/api-fu) - Comprehensive GraphQL implementation. Stars:`57`.
- [straf](https://github.com/SonicRoshan/straf) - Easily Convert Golang structs to GraphQL objects. Stars:`40`.
- [jsonpath](https://github.com/AsaiYusuke/jsonpath) - A query library for retrieving part of JSON based on JSONPath syntax. Stars:`27`.
- [gws](https://github.com/Zaba505/gws) - Apollos' "GraphQL over Websocket" client and server implementation. Stars:`6`.
- [grapher](https://github.com/reaganiwadha/grapher) - A GraphQL field builder utilizing Go generics with extra utilities and features. Stars:`4`.

**[⬆ back to top](#contents)**

## Reflection

- [Deepcopier](https://github.com/ulule/deepcopier) - Simple struct copying for Go. Stars:`458`.
- [go-deepcopy](https://github.com/tiendc/go-deepcopy) - Fast deep copy library. Stars:`101`.
- [gotype](https://github.com/wzshiming/gotype) - Golang source code parsing, usage like reflect package. Stars:`63`.
- [copy](https://github.com/gotidy/copy) - Package for fast copying structs of different types. Stars:`51`.
- [gpath](https://github.com/tenntenn/gpath) - Library to simplify access struct fields with Go's expression in reflection. Stars:`40`.
- [goenum](https://github.com/lvyahui8/goenum) - A common enumeration struct based on generics and reflection that allows you to quickly define enumerations and use a set of useful default methods. Stars:`16`.
- [reflectutils](https://github.com/muir/reflectutils) - Helpers for working with reflection: struct tag parsing; recursive walking; fill value from string. Stars:`8`.
- [reflectpro](https://github.com/gontainer/reflectpro) - Callers, copiers, getters and setters for go. Stars:`7`.
- [objwalker](https://github.com/rekby/objwalker) - Walk by go objects with reflection. Stars:`3`.

**[⬆ back to top](#contents)**

## Resource Embedding

- [vfsgen](https://github.com/shurcooL/vfsgen) - Generates a vfsdata.go file that statically implements the given virtual filesystem. Stars:`982`.
- [debme](https://github.com/leaanthony/debme) - Create an `embed.FS` from an existing `embed.FS` subdirectory. Stars:`34`.
- [rebed](https://github.com/soypat/rebed) - Recreate folder structures and files from Go 1.16's `embed.FS` type Stars:`29`.
- [embed](https://pkg.go.dev/embed) - Package embed provides access to files embedded in the running Go program.

**[⬆ back to top](#contents)**

## Science and Data Analysis

_Libraries for scientific computing and data analyzing._

- [gonum](https://github.com/gonum/gonum) - Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. Stars:`8.1K`.
- [stats](https://github.com/montanaflynn/stats) - Statistics package with common functions missing from the Golang standard library. Stars:`3.0K`.
- [gonum/plot](https://github.com/gonum/plot) - gonum/plot provides an API for building and drawing plots in Go. Stars:`2.9K`.
- [gosl](https://github.com/cpmech/gosl) - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. Stars:`1.9K`.
- [streamtools](https://github.com/nytlabs/streamtools) - general purpose, graphical tool for dealing with streams of data. Stars:`1.3K`.
- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Dataframes for machine-learning and statistics (similar to pandas). Stars:`1.3K`.
- [orb](https://github.com/paulmach/orb) - 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. Stars:`1.0K`.
- [go-dsp](https://github.com/mjibson/go-dsp) - Digital Signal Processing for Go. Stars:`885`.
- [chart](https://github.com/vdobler/chart) - Simple Chart Plotting library for Go. Supports many graphs types. Stars:`774`.
- [goraph](https://github.com/gyuho/goraph) - Pure Go graph theory library(data structure, algorithm visualization). Stars:`743`.
- [graph](https://github.com/yourbasic/graph) - Library of basic graph algorithms. Stars:`731`.
- [Poly](https://github.com/bebop/poly) - A Go package for engineering organisms. Stars:`709`.
- [ewma](https://github.com/VividCortex/ewma) - Exponentially-weighted moving averages. Stars:`448`.
- [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) - Calendar heatmap in plain Go inspired by Github contribution activity. Stars:`408`.
- [go-hep](https://github.com/go-hep/hep) - A set of libraries and tools for performing High Energy Physics analyses with ease. Stars:`248`.
- [TextRank](https://github.com/DavidBelicza/TextRank) - TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. Stars:`217`.
- [sparse](https://github.com/james-bowman/sparse) - Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. Stars:`165`.
- [go-estimate](https://github.com/milosgajdos/go-estimate) - State estimation and filtering algorithms in Go. Stars:`117`.
- [gograph](https://github.com/hmdsefi/gograph) - A golang generic graph library that provides mathematical graph-theory and algorithms. Stars:`89`.
- [pagerank](https://github.com/alixaxel/pagerank) - Weighted PageRank algorithm implemented in Go. Stars:`86`.
- [jsonl-graph](https://github.com/nikolaydubina/jsonl-graph) - Tool to manipulate JSONL graphs with graphviz support. Stars:`75`.
- [geom](https://github.com/skelterjohn/geom) - 2D geometry for golang. Stars:`55`.
- [evaler](https://github.com/soniah/evaler) - Simple floating point arithmetic expression evaluator. Stars:`51`.
- [decimal](https://github.com/db47h/decimal) - Package decimal implements arbitrary-precision decimal floating-point arithmetic. Stars:`45`.
- [triangolatte](https://github.com/tchayen/triangolatte) - 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. Stars:`37`.
- [goent](https://github.com/kzahedi/goent) - GO Implementation of Entropy Measures. Stars:`35`.
- [piecewiselinear](https://github.com/sgreben/piecewiselinear) - Tiny linear interpolation library. Stars:`28`.
- [godesim](https://github.com/soypat/godesim) - Extended/multivariable ODE solver framework for event-based simulations with simple API. Stars:`23`.
- [GoStats](https://github.com/OGFris/GoStats) - GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. Stars:`22`.
- [ode](https://github.com/ChristopherRabotin/ode) - Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. Stars:`22`.
- [PiHex](https://github.com/claygod/PiHex) - Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. Stars:`20`.
- [assocentity](https://github.com/ndabAP/assocentity) - Package assocentity returns the average distance from words to a given entity. Stars:`17`.
- [topk](https://github.com/keilerkonzept/topk) - Sliding-window and regular top-K sketches, based on the HeavyKeeper algorithm. Stars:`15`.
- [rootfinding](https://github.com/khezen/rootfinding) - root-finding algorithms library for finding roots of quadratic functions. Stars:`11`.
- [bradleyterry](https://github.com/seanhagen/bradleyterry) - Provides a Bradley-Terry Model for pairwise comparisons. Stars:`11`.
- [go-gt](https://github.com/ThePaw/go-gt) - Graph theory algorithms written in "Go" language. Stars:`11`.

**[⬆ back to top](#contents)**

## Security

_Libraries that are used to help make your application more secure._

- [age](https://github.com/FiloSottile/age) - A simple, modern and secure encryption tool (and Go library) with small explicit keys, no config options, and UNIX-style composability. Stars:`19.4K`.
- [lego](https://github.com/go-acme/lego) - Pure Go ACME client library and CLI tool (for use with Let's Encrypt). Stars:`8.8K`.
- [CertMagic](https://github.com/caddyserver/certmagic) - Mature, robust, and powerful ACME client integration for fully-managed TLS certificate issuance and renewal. Stars:`5.3K`.
- [Cameradar](https://github.com/Ullaakut/cameradar) - Tool and library to remotely hack RTSP streams from surveillance cameras. Stars:`4.6K`.
- [Coraza](https://github.com/corazawaf/coraza) - Enterprise-ready, modsecurity and OWASP CRS compatible WAF library. Stars:`2.8K`.
- [memguard](https://github.com/awnumar/memguard) - A pure Go library for handling sensitive values in memory. Stars:`2.6K`.
- [secure](https://github.com/unrolled/secure) - HTTP middleware for Go that facilitates some quick security wins. Stars:`2.3K`.
- [acmetool](https://github.com/hlandau/acme) - ACME (Let's Encrypt) client tool with automatic renewal. Stars:`2.1K`.
- [themis](https://github.com/cossacklabs/themis) - high-level cryptographic library for solving typical data security tasks (secure data storage, secure messaging, zero-knowledge proof authentication), available for 14 languages, best fit for multi-platform apps. Stars:`1.9K`.
- [acra](https://github.com/cossacklabs/acra) - Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. Stars:`1.4K`.
- [beelzebub](https://github.com/mariocandela/beelzebub) - A secure low code honeypot framework, leveraging AI for System Virtualization. Stars:`1.3K`.
- [dongle](https://github.com/golang-module/dongle) - A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption. Stars:`990`.
- [booster](https://github.com/anatol/booster) - Fast initramfs generator with full-disk encryption support. Stars:`579`.
- [nacl](https://github.com/kevinburke/nacl) - Go implementation of the NaCL set of API's. Stars:`548`.
- [go-password-validator](https://github.com/lane-c-wagner/go-password-validator) - Password validator based on raw cryptographic entropy values. Stars:`541`.
- [ssh-vault](https://github.com/ssh-vault/ssh-vault) - encrypt/decrypt using ssh keys. Stars:`473`.
- [teler-waf](https://github.com/kitabisa/teler-waf) - teler-waf is a Go HTTP middleware that provide teler IDS functionality to protect against web-based attacks and improve the security of Go-based web applications. It is highly configurable and easy to integrate into existing Go applications. Stars:`383`.
- [go-yara](https://github.com/hillu/go-yara) - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". Stars:`375`.
- [optimus-go](https://github.com/pjebs/optimus-go) - ID hashing and Obfuscation using Knuth's Algorithm. Stars:`361`.
- [firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) - A rest application to dynamically update firewalld rules on a linux server. Stars:`343`.
- [BadActor](https://github.com/jaredfolkins/badactor) - In-memory, application-driven jailer built in the spirit of fail2ban. Stars:`324`.
- [go-peer](https://github.com/number571/go-peer) - A software library for creating secure and anonymous decentralized systems. Stars:`295`.
- [passlib](https://github.com/hlandau/passlib) - Futureproof password hashing library. Stars:`293`.
- [simple-scrypt](https://github.com/elithrar/simple-scrypt) - Scrypt package with a simple, obvious API and automatic cost calibration built-in. Stars:`201`.
- [luks.go](https://github.com/anatol/luks.go) - Pure Golang library to manage LUKS partitions. Stars:`91`.
- [argon2pw](https://github.com/raja/argon2pw) - Argon2 password hash generation with constant-time password comparison. Stars:`91`.
- [passwap](https://github.com/zitadel/passwap) - Provides a unified implementation between different password hashing algorithms Stars:`67`.
- [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) - A probably paranoid package for securely hashing and encrypting passwords. Stars:`60`.
- [go-generate-password](https://github.com/m1/go-generate-password) - Password generator that can be used on the cli or as a library. Stars:`55`.
- [go-htpasswd](https://github.com/tg123/go-htpasswd) - Apache htpasswd Parser for Go. Stars:`44`.
- [certificates](https://github.com/mvmaasakkers/certificates) - An opinionated tool for generating tls certificates. Stars:`38`.
- [secureio](https://github.com/xaionaro-go/secureio) - An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519. Stars:`33`.
- [secret](https://github.com/rsjethani/secret) - Prevent your secrets from leaking into logs, std\* etc. Stars:`31`.
- [sslmgr](https://github.com/adrianosela/sslmgr) - SSL certificates made easy with a high level wrapper around acme/autocert. Stars:`30`.
- [argon2-hashing](https://github.com/andskur/argon2-hashing) - light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package. Stars:`23`.
- [goArgonPass](https://github.com/dwin/goArgonPass) - Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. Stars:`20`.
- [qrand](https://github.com/bitfield/qrand) - Client for the ANU Quantum Numbers (AQN) API, providing quantum-mechanically secure random data. Stars:`16`.
- [multikey](https://github.com/adrianosela/multikey) - An n-out-of-N keys encryption/decryption framework based on Shamir's Secret Sharing algorithm. Stars:`9`.
- [Interpol](https://github.com/avahidi/interpol) - Rule-based data generator for fuzzing and penetration testing. Stars:`7`.
- [entpassgen](https://github.com/andreimerlescu/entpassgen) - Entropy Password Generator with extensive command line arguments to generate random strings securely including digits, passwords, and passwords built using obscure dictionary words mixed with symbols and digits. Stars:`5`.
- [encid](https://github.com/bobg/encid) - Encode and decode encrypted integer IDs. Stars:`5`.
- [autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert) - Auto provision Let's Encrypt certificates and start a TLS server.
- [acopw-go](https://sr.ht/~jamesponddotco/acopw-go/) - Small cryptographically secure password generator package for Go.

**[⬆ back to top](#contents)**

## Serialization

_Libraries and tools for binary serialization._

- [jsoniter](https://github.com/json-iterator/go) - High-performance 100% compatible drop-in replacement of "encoding/json". Stars:`13.8K`.
- [goprotobuf](https://github.com/golang/protobuf) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. Stars:`10.0K`.
- [mapstructure](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures. Stars:`8.0K`.
- [gogoprotobuf](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets. Stars:`5.7K`.
- [go-codec](https://github.com/ugorji/go) - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. Stars:`1.9K`.
- [csvutil](https://github.com/jszwec/csvutil) - High Performance, idiomatic CSV record encoding and decoding to native Go structures. Stars:`1.0K`.
- [cbor](https://github.com/fxamacker/cbor) - Small, safe, and easy CBOR encoding and decoding library. Stars:`907`.
- [colfer](https://github.com/pascaldekloe/colfer) - Code generation for the Colfer binary format. Stars:`749`.
- [go-capnproto](https://github.com/glycerine/go-capnproto) - Cap'n Proto library and parser for go. Stars:`287`.
- [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. Stars:`167`.
- [structomap](https://github.com/tuvistavie/structomap) - Library to easily and dynamically generate maps from static structures. Stars:`144`.
- [binstruct](https://github.com/ghostiam/binstruct) - Golang binary decoder for mapping data into the structure. Stars:`106`.
- [bambam](https://github.com/glycerine/bambam) - generator for Cap'n Proto schemas from go. Stars:`64`.
- [bel](https://github.com/32leaves/bel) - Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. Stars:`44`.
- [fwencoder](https://github.com/o1egl/fwencoder) - Fixed width file parser (encoding and decoding library) for Go. Stars:`27`.
- [elastic](https://github.com/epiclabs-io/elastic) - Convert slices, maps or any other unknown value across different types at run-time, no matter what. Stars:`25`.
- [gotiny](https://github.com/raszia/gotiny) - Efficient Go serialization library, gotiny is almost as fast as serialization libraries that generate code. Stars:`21`.
- [go-csvlib](https://github.com/tiendc/go-csvlib) - High level and rich functionalities CSV serialization/deserialization library. Stars:`19`.
- [pletter](https://github.com/vimeda/pletter) - A standard way to wrap a proto message for message brokers. Stars:`19`.
- [fixedwidth](https://github.com/huydang284/fixedwidth) - Fixed-width text formatting (UTF-8 supported). Stars:`9`.
- [unitpacking](https://github.com/recolude/unitpacking) - Library to pack unit vectors into as fewest bytes as possible. Stars:`7`.

**[⬆ back to top](#contents)**

## Server Applications

- [Caddy](https://github.com/caddyserver/caddy) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use. Stars:`65.9K`.
- [minio](https://github.com/minio/minio) - Minio is a distributed object storage server. Stars:`54.3K`.
- [etcd](https://github.com/etcd-io/etcd) - Highly-available key value store for shared configuration and service discovery. Stars:`49.9K`.
- [pocketbase](https://github.com/pocketbase/pocketbase) - PocketBase is a realtime backend in 1 file consisting of embedded database (SQLite) with realtime subscriptions, built-in auth management and much more. Stars:`49.6K`.
- [SFTPGo](https://github.com/drakkan/sftpgo) - Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support. It can serve local filesystem and Cloud Storage backends such as S3 and Google Cloud Storage. Stars:`10.8K`.
- [RoadRunner](https://github.com/spiral/roadrunner) - High-performance PHP application server, load-balancer and process manager. Stars:`8.2K`.
- [Easegress](https://github.com/megaease/easegress) - A cloud native high availability/performance traffic orchestration system with observability and extensibility. Stars:`5.8K`.
- [flipt](https://github.com/markphelps/flipt) - A self contained feature flag solution written in Go and Vue.js Stars:`4.5K`.
- [Wish](https://github.com/charmbracelet/wish) - Make SSH apps, just like that! Stars:`4.3K`.
- [Fider](https://github.com/getfider/fider) - Fider is an open platform to collect and organize customer feedback. Stars:`3.7K`.
- [devd](https://github.com/cortesi/devd) - Local webserver for developers. Stars:`3.5K`.
- [algernon](https://github.com/xyproto/algernon) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. Stars:`2.9K`.
- [Flagr](https://github.com/checkr/flagr) - Flagr is an open-source feature flagging and A/B testing service. Stars:`2.5K`.
- [Trickster](https://github.com/tricksterproxy/trickster) - HTTP reverse proxy cache and time series accelerator. Stars:`2.0K`.
- [discovery](https://github.com/Bilibili/discovery) - A registry for resilient mid-tier load balancing and failover. Stars:`1.8K`.
- [go-feature-flag](https://github.com/thomaspoignant/go-feature-flag) - A simple, complete and lightweight self-hosted feature flag solution 100% Open Source. Stars:`1.7K`.
- [Clace](https://github.com/claceio/clace) - Clace makes internal tool deployment and management easy by implementing an app server for containerized webapps. Stars:`642`.
- [Euterpe](https://github.com/ironsmile/euterpe) - Self-hosted music streaming server with built-in web UI and REST API. Stars:`557`.
- [wd-41](https://github.com/baalimago/wd-41) - A (w)eb (d)evelopment server with automatic live-reload on file changes. Stars:`149`.
- [dudeldu](https://github.com/krotik/dudeldu) - A simple SHOUTcast server. Stars:`146`.
- [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) - Simple Reverse Proxy with Caching, written in Go, using Redis. Stars:`143`.
- [cortex-tenant](https://github.com/blind-oracle/cortex-tenant) - Prometheus remote write proxy that adds add Cortex tenant ID header based on metric labels. Stars:`125`.
- [lets-proxy2](https://github.com/rekby/lets-proxy2) - Reverse proxy for handle https with issue certificates in fly from lets-encrypt. Stars:`100`.
- [psql-streamer](https://github.com/blind-oracle/psql-streamer) - Stream database events from PostgreSQL to Kafka. Stars:`60`.
- [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) - Nginx log parser and exporter to Prometheus. Stars:`41`.
- [protoxy](https://github.com/camgraff/protoxy) - A proxy server that converts JSON request bodies to Protocol Buffers. Stars:`34`.
- [Moxy](https://github.com/sinhashubham95/moxy) - Moxy is a simple mocker and proxy application server, you can create mock endpoints as well as proxy requests in case no mock exists for the endpoint. Stars:`16`.
- [gondola](https://github.com/bmf-san/gondola) - A YAML based golang reverse proxy. Stars:`9`.
- [riemann-relay](https://github.com/blind-oracle/riemann-relay) - Relay to load-balance Riemann events and/or convert them to Carbon. Stars:`2`.
- [nsq](https://nsq.io/) - A realtime distributed messaging platform.
- [Engity's Bifröst](https://bifroest.engity.org/) - Highly customizable SSH server with several ways to authorize a user how to execute its session (local or in containers).
- [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.

**[⬆ back to top](#contents)**

## Stream Processing

_Libraries and tools for stream processing and reactive programming._

- [go-streams](https://github.com/reugn/go-streams) - Go stream processing library. Stars:`2.1K`.
- [machine](https://github.com/whitaker-io/machine) - Go library for writing and generating stream workers with built in metrics and traceability. Stars:`161`.
- [go-etl](https://github.com/Breeze0806/go-etl) - A lightweight toolkit for data source extraction, transformation, and loading (ETL). Stars:`153`.
- [stream](https://github.com/youthlin/stream) - Go Stream, like Java 8 Stream: Filter/Map/FlatMap/Peek/Sorted/ForEach/Reduce... Stars:`95`.
- [goio](https://github.com/primetalk/goio) - An implementation of IO, Stream, Fiber for Golang, inspired by awesome Scala libraries cats and fs2. Stars:`86`.
- [nibbler](https://github.com/naughtygopher/nibbler) - A lightweight package for micro batch processing. Stars:`15`.

**[⬆ back to top](#contents)**

## Template Engines

_Libraries and tools for templating and lexing._

- [templ](https://github.com/a-h/templ) - A HTML templating language that has great developer tooling. Stars:`9.6K`.
- [quicktemplate](https://github.com/valyala/quicktemplate) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. Stars:`3.2K`.
- [pongo2](https://github.com/flosch/pongo2) - Django-like template-engine for Go. Stars:`3.0K`.
- [maroto](https://github.com/johnfercher/maroto) - A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. Stars:`2.5K`.
- [jet](https://github.com/CloudyKit/jet) - Jet template engine. Stars:`1.3K`.
- [fasttemplate](https://github.com/valyala/fasttemplate) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](https://golang.org/pkg/text/template/). Stars:`878`.
- [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang. Stars:`873`.
- [raymond](https://github.com/aymerick/raymond) - Complete handlebars implementation in Go. Stars:`639`.
- [ego](https://github.com/benbjohnson/ego) - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. Stars:`584`.
- [goview](https://github.com/foolin/goview) - Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. Stars:`457`.
- [liquid](https://github.com/osteele/liquid) - Go implementation of Shopify Liquid templates. Stars:`310`.
- [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). Stars:`175`.
- [sprout](https://github.com/go-sprout/sprout) - Useful template functions for Go templates. Stars:`174`.
- [got](https://github.com/goradd/got) - A Go code generator inspired by Hero and Fasttemplate. Has include files, custom tag definitions, injected Go code, language translation, and more. Stars:`38`.
- [tbd](https://github.com/lucasepe/tbd) - A really simple way to create text templates with placeholders - exposes extra builtin Git repo metadata. Stars:`27`.
- [htmgo](https://htmgo.dev) - build simple and scalable systems with go + htmx
- [gomponents](https://www.gomponents.com) - HTML 5 components in pure Go, that look something like this: `func(name string) g.Node { return Div(Class("headline"), g.Textf("Hi %v!", name)) }`.
- [extemplate](https://git.sr.ht/~dvko/extemplate) - Tiny wrapper around html/template to allow for easy file-based template inheritance.

**[⬆ back to top](#contents)**

## Testing

_Libraries for testing codebases and generating test data._

### Testing Frameworks

- [Testify](https://github.com/stretchr/testify) - Sacred extension to the standard go testing package. Stars:`25.0K`.
- [keploy](https://github.com/keploy/keploy) - Generate Testcase and Data Mocks from API calls automatically. Stars:`10.3K`.
- [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload. Stars:`8.4K`.
- [go-cmp](https://github.com/google/go-cmp) - Package for comparing Go values in tests. Stars:`4.5K`.
- [testcontainers-go](https://github.com/testcontainers/testcontainers-go) - A Go package that makes it simple to create and clean up container-based dependencies for automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers that should be run as part of a test and clean up those resources when the test is done. Stars:`4.2K`.
- [httpexpect](https://github.com/gavv/httpexpect) - Concise, declarative, and easy to use end-to-end HTTP and REST API testing. Stars:`2.7K`.
- [godog](https://github.com/cucumber/godog) - Cucumber BDD framework for Go. Stars:`2.5K`.
- [is](https://github.com/matryer/is) - Professional lightweight testing mini-framework for Go. Stars:`1.9K`.
- [gnomock](https://github.com/orlangure/gnomock) - integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks. Stars:`1.5K`.
- [go-vcr](https://github.com/dnaeon/go-vcr) - Record and replay your HTTP interactions for fast, deterministic and accurate tests. Stars:`1.3K`.
- [testfixtures](https://github.com/go-testfixtures/testfixtures) - A helper for Rails' like test fixtures to test database applications. Stars:`1.2K`.
- [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) - Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. Stars:`990`.
- [goblin](https://github.com/franela/goblin) - Mocha like testing framework of Go. Stars:`890`.
- [goc](https://github.com/qiniu/goc) - Goc is a comprehensive coverage testing system for The Go Programming Language. Stars:`865`.
- [baloo](https://github.com/h2non/baloo) - Expressive and versatile end-to-end HTTP API testing made easy. Stars:`779`.
- [go-mutesting](https://github.com/zimmski/go-mutesting) - Mutation testing for Go source code. Stars:`657`.
- [gotest.tools](https://github.com/gotestyourself/gotest.tools) - A collection of packages to augment the go testing package and support common patterns. Stars:`564`.
- [go-testdeep](https://github.com/maxatome/go-testdeep) - Extremely flexible golang deep comparison, extends the go testing package. Stars:`449`.
- [gofight](https://github.com/appleboy/gofight) - API Handler Testing for Golang Router framework. Stars:`440`.
- [testza](https://github.com/MarvinJWendt/testza) - Full-featured test framework with nice colorized output. Stars:`421`.
- [cupaloy](https://github.com/bradleyjkemp/cupaloy) - Simple snapshot testing addon for your test framework. Stars:`318`.
- [frisby](https://github.com/verdverm/frisby) - REST API testing framework. Stars:`276`.
- [got](https://github.com/ysmood/got) - An enjoyable golang test framework. Stars:`269`.
- [endly](https://github.com/viant/endly) - Declarative end to end functional testing. Stars:`268`.
- [go-hit](https://github.com/Eun/go-hit) - Hit is an http integration test framework written in golang. Stars:`260`.
- [go-carpet](https://github.com/msoap/go-carpet) - Tool for viewing test coverage in terminal. Stars:`249`.
- [commander](https://github.com/SimonBaeumer/commander) - Tool for testing cli applications on windows, linux and osx. Stars:`228`.
- [charlatan](https://github.com/percolate/charlatan) - Tool to generate fake interface implementations for tests. Stars:`204`.
- [dbcleaner](https://github.com/khaiql/dbcleaner) - Clean database for testing purpose, inspired by `database_cleaner` in Ruby. Stars:`164`.
- [jsonassert](https://github.com/kinbiko/jsonassert) - Package for verifying that your JSON payloads are serialized correctly. Stars:`138`.
- [testcase](https://github.com/adamluzsi/testcase) - Idiomatic testing framework for Behavior Driven Development. Stars:`123`.
- [be](https://github.com/carlmjohnson/be) - The minimalist generic test assertion library. Stars:`108`.
- [gocrest](https://github.com/corbym/gocrest) - Composable hamcrest-like matchers for Go assertions. Stars:`106`.
- [wstest](https://github.com/posener/wstest) - Websocket client for unit-testing a websocket http.Handler. Stars:`102`.
- [Gont](https://github.com/stv0g/gont) - Go network testing toolkit for testing building complex network topologies using Linux namespaces. Stars:`85`.
- [gherkingen](https://github.com/hedhyw/gherkingen) - BDD boilerplate generator and framework. Stars:`84`.
- [testcerts](https://github.com/madflojo/testcerts) - Dynamically generate self-signed certificates and certificate authorities within your test functions. Stars:`75`.
- [assert](https://github.com/go-playground/assert) - Basic Assertion Library used along side native go testing, with building blocks for custom assertions. Stars:`65`.
- [restit](https://github.com/yookoala/restit) - Go micro framework to help writing RESTful API integration test. Stars:`55`.
- [gospecify](https://github.com/stesla/gospecify) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. Stars:`52`.
- [gomatch](https://github.com/jfilipczyk/gomatch) - library created for testing JSON against patterns. Stars:`48`.
- [dsunit](https://github.com/viant/dsunit) - Datastore testing for SQL, NoSQL, structured files. Stars:`45`.
- [fixenv](https://github.com/rekby/fixenv) - Fixture manage engine, inspired by pytest fixtures. Stars:`32`.
- [Hamcrest](https://github.com/rdrdr/hamcrest) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. Stars:`30`.
- [schema](https://github.com/jgroeneveld/schema) - Quick and easy expression matching for JSON schemas used in requests and responses. Stars:`20`.
- [flute](https://github.com/suzuki-shunsuke/flute) - HTTP client testing framework. Stars:`20`.
- [testsql](https://github.com/zhulongcheng/testsql) - Generate test data from SQL files before testing and clear it after finished. Stars:`17`.
- [dft](https://github.com/abecodes/dft) - Lightweight, zero dependency docker containers for testing (or more). Stars:`16`.
- [gogiven](https://github.com/corbym/gogiven) - YATSPEC-like BDD testing framework for Go. Stars:`15`.
- [biff](https://github.com/fulldump/biff) - Bifurcation testing framework, BDD compatible. Stars:`14`.
- [gosuite](https://github.com/pavlo/gosuite) - Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. Stars:`12`.
- [stop-and-go](https://github.com/elgohr/stop-and-go) - Testing helper for concurrency. Stars:`12`.
- [Tt](https://github.com/vcaesar/tt) - Simple and colorful test tools. Stars:`9`.
- [trial](https://github.com/jgroeneveld/trial) - Quick and easy extendable assertions without introducing much boilerplate. Stars:`6`.
- [go-testpredicate](https://github.com/maargenton/go-testpredicate) - Test predicate style assertions library with extensive diagnostics output. Stars:`5`.
- [envite](https://github.com/PerimeterX/envite) - Dev and testing environment management framework. Stars:`4`.
- [go-mysql-test-container](https://github.com/arikama/go-mysql-test-container) - Golang MySQL testcontainer to help with MySQL integration testing. Stars:`3`.
- [arch-go](https://github.com/fdaines/arch-go) - Architecture testing tool for Go projects. Stars:`2`.
- [omg.testingtools](https://github.com/dedalqq/omg.testingtools) - The simple library for change a values of private fields for testing. Stars:`1`.
- [gomega](https://onsi.github.io/gomega/) - Rspec like matcher/assertion library.
- [go-snaps](http://github.com/gkampitakis/go-snaps) - Jest-like snapshot testing in Golang.
- [apitest](https://apitest.dev) - Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
- [gocheck](https://labix.org/gocheck) - More advanced testing framework alternative to gotest.
- [ginkgo](https://onsi.github.io/ginkgo/) - BDD Testing Framework for Go.

### Mock

- [mockery](https://github.com/vektra/mockery) - Tool to generate Go interfaces. Stars:`6.7K`.
- [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database interactions. Stars:`6.4K`.
- [gomock](https://github.com/uber-go/mock) - Mocking framework for the Go programming language. Stars:`2.9K`.
- [hoverfly](https://github.com/SpectoLabs/hoverfly) - HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. Stars:`2.4K`.
- [gock](https://github.com/h2non/gock) - Versatile HTTP mocking made easy. Stars:`2.2K`.
- [moq](https://github.com/matryer/moq) - Utility that generates a struct from any interface. The struct can be used in test code as a mock of the interface. Stars:`2.1K`.
- [httpmock](https://github.com/jarcoal/httpmock) - Easy mocking of HTTP responses from external resources. Stars:`2.0K`.
- [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects. Stars:`1.1K`.
- [go-txdb](https://github.com/DATA-DOG/go-txdb) - Single transaction based database driver mainly for testing purposes. Stars:`722`.
- [minimock](https://github.com/gojuno/minimock) - Mock generator for Go interfaces. Stars:`714`.
- [pgxmock](https://github.com/pashagolub/pgxmock) - A mock library implementing [pgx - PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx/). Stars:`503`.
- [xgo](https://github.com/xhd2015/xgo) - A general pureposed function mocking library. Stars:`416`.
- [govcr](https://github.com/seborama/govcr) - HTTP mock for Golang: record and replay HTTP interactions for offline testing. Stars:`191`.
- [go-localstack](https://github.com/elgohr/go-localstack) - Tool for using localstack in AWS testing. Stars:`83`.
- [timex](https://github.com/cabify/timex) - A test-friendly replacement for the native `time` package. Stars:`71`.
- [mockhttp](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter. Stars:`23`.
- [mooncake](https://github.com/GuilhermeCaruso/mooncake) - A simple way to generate mocks for multiple purposes. Stars:`17`.
- [genmock](https://gitlab.com/so_literate/genmock) - Go mocking system with code generator for building calls of the interface methods.

### Fuzzing and delta-debugging/reducing/shrinking

- [go-fuzz](https://github.com/dvyukov/go-fuzz) - Randomized testing system. Stars:`4.8K`.
- [gofuzz](https://github.com/google/gofuzz) - Library for populating go objects with random values. Stars:`1.5K`.
- [Tavor](https://github.com/zimmski/tavor) - Generic fuzzing and delta-debugging framework. Stars:`247`.

### Selenium and browser control tools

- [chromedp](https://github.com/knq/chromedp) - a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. Stars:`12.1K`.
- [rod](https://github.com/go-rod/rod) - A Devtools driver to make web automation and scraping easy. Stars:`6.1K`.
- [playwright-go](https://github.com/mxschmitt/playwright-go) - browser automation library to control Chromium, Firefox and WebKit with a single API. Stars:`2.8K`.
- [selenoid](https://github.com/aerokube/selenoid) - alternative Selenium hub server that launches browsers within containers. Stars:`2.6K`.
- [cdp](https://github.com/mafredri/cdp) - Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. Stars:`758`.
- [ggr](https://github.com/aerokube/ggr) - a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. Stars:`315`.

### Fail injection

- [failpoint](https://github.com/pingcap/failpoint) - An implementation of [failpoints](https://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. Stars:`854`.

**[⬆ back to top](#contents)**

## Text Processing

_Libraries for parsing and manipulating texts._

See also [Natural Language Processing](#natural-language-processing) and [Text Analysis](#text-analysis).

### Formatters

- [go-humanize](https://github.com/dustin/go-humanize) - Formatters for time, numbers, and memory size to human readable format. Stars:`4.6K`.
- [sq](https://github.com/neilotoole/sq) - Convert data from SQL databases or document formats like CSV or Excel into formats such as JSON, Excel, CSV, HTML, Markdown, XML, and YAML. Stars:`2.3K`.
- [bytes](https://github.com/labstack/gommon/tree/master/bytes) - Formats and parses numeric byte values (10K, 2M, 3G, etc.). Stars:`556`.
- [gotabulate](https://github.com/bndr/gotabulate) - Easily pretty-print your tabular data with Go. Stars:`337`.
- [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) - Fixed-width text formatting (encoder/decoder with reflection). Stars:`86`.
- [align](https://github.com/Guitarbum722/align) - A general purpose application that aligns text. Stars:`84`.
- [address](https://github.com/bojanz/address) - Handles address representation, validation and formatting. Stars:`77`.
- [textwrap](https://github.com/isbm/textwrap) - Wraps text at end of lines. Implementation of `textwrap` module from Python. Stars:`5`.

### Markup Languages

- [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go. Stars:`5.6K`.
- [toml](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection). Stars:`4.8K`.
- [goldmark](https://github.com/yuin/goldmark) - A Markdown parser written in Go. Easy to extend, standard (CommonMark) compliant, well structured. Stars:`4.2K`.
- [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) - Convert HTML to Markdown. Even works with entire websites and can be extended through rules. Stars:`3.0K`.
- [go-toml](https://github.com/pelletier/go-toml) - Go library for the TOML format with query support and handy cli tools. Stars:`1.8K`.
- [htmlquery](https://github.com/antchfx/htmlquery) - An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. Stars:`768`.
- [mxj](https://github.com/clbanning/mxj) - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. Stars:`628`.
- [goq](https://github.com/andrewstuart/goq) - Declarative unmarshalling of HTML using struct tags with jQuery syntax (uses GoQuery). Stars:`267`.
- [bafi](https://github.com/mmalcek/bafi) - Universal JSON, BSON, YAML, XML translator to ANY format using templates. Stars:`108`.
- [go-output-format](https://github.com/drewstinnett/go-output-format) - Output go structures into multiple formats (YAML/JSON/etc) in your command line app. Stars:`18`.
- [bbConvert](https://github.com/CalebQ42/bbConvert) - Converts bbCode to HTML that allows you to add support for custom bbCode tags. Stars:`12`.
- [htmlyaml](https://github.com/nikolaydubina/htmlyaml) -  Rich rendering of YAML as HTML in Go Stars:`5`.
- [htree](https://github.com/bobg/htree) - Traverse, navigate, filter, and otherwise process trees of [html.Node](https://pkg.go.dev/golang.org/x/net/html#Node) objects. Stars:`4`.

### Parsers/Encoders/Decoders

- [sh](https://github.com/mvdan/sh) - Shell parser and formatter. Stars:`7.9K`.
- [gofeed](https://github.com/mmcdole/gofeed) - Parse RSS and Atom feeds in Go. Stars:`2.7K`.
- [go-querystring](https://github.com/google/go-querystring) - Go library for encoding structs into URL query parameters. Stars:`2.1K`.
- [when](https://github.com/olebedev/when) - Natural EN and RU language date/time parser with pluggable rules. Stars:`1.4K`.
- [godump (goforj)](https://github.com/goforj/godump) - Pretty-print Go structs with Laravel/Symfony-style dumps, full type info, colorized CLI output, cycle detection, and private field access. Stars:`1.3K`.
- [commonregex](https://github.com/mingrammer/commonregex) - A collection of common regular expressions for Go. Stars:`900`.
- [gographviz](https://github.com/awalterschulze/gographviz) - Parses the Graphviz DOT language. Stars:`565`.
- [go-nmea](https://github.com/adrianmo/go-nmea) - NMEA parser library for the Go language. Stars:`245`.
- [godump](https://github.com/yassinebenaid/godump) - Pretty print any GO variable with ease, an alternative to Go's `fmt.Printf("%#v")`. Stars:`217`.
- [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) - Editorconfig file parser and manipulator for Go. Stars:`147`.
- [tokenizer](https://github.com/bzick/tokenizer) - Parse any string, slice or infinite buffer to any tokens. Stars:`129`.
- [go-vcard](https://github.com/emersion/go-vcard) - Parse and format vCard. Stars:`118`.
- [sdp](https://github.com/gortc/sdp) - SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. Stars:`113`.
- [did](https://github.com/ockam-network/did) - DID (Decentralized Identifiers) Parser and Stringer in Go. Stars:`93`.
- [allot](https://github.com/sbstjn/allot) - Placeholder and wildcard text parsing for CLI tools and bots. Stars:`59`.
- [vdf](https://github.com/andygrunwald/vdf) - A Lexer and Parser for Valves Data Format (known as vdf) written in Go. Stars:`58`.
- [normalize](https://github.com/avito-tech/normalize) - Sanitize, normalize and compare fuzzy text. Stars:`52`.
- [gonameparts](https://github.com/polera/gonameparts) - Parses human names into individual name parts. Stars:`43`.
- [go-fasttld](https://github.com/elliotwutingfeng/go-fasttld) - High performance effective top level domains (eTLD) extraction module. Stars:`39`.
- [xj2go](https://github.com/stackerzzq/xj2go) - Convert xml or json to go struct. Stars:`36`.
- [codetree](https://github.com/aerogo/codetree) - Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. Stars:`25`.
- [doi](https://github.com/hscells/doi) - Document object identifier (doi) parser in Go. Stars:`11`.
- [parseargs-go](https://github.com/nproc/parseargs-go) - string argument parser that understands quotes and backslashes. Stars:`10`.
- [ltsv](https://github.com/Wing924/ltsv) - High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go. Stars:`10`.
- [prattle](https://github.com/askeladdk/prattle) - Scan and parse LL(1) grammars simply and efficiently. Stars:`7`.
- [parth](https://github.com/codemodus/parth) - URL path segmentation parsing.
- [encdec](https://github.com/mickep76/encdec) - Package provides a generic interface to encoders and decoders.

### Regular Expressions

- [rex](https://github.com/hedhyw/rex) - Regular expressions builder. Stars:`206`.
- [regroup](https://github.com/oriser/regroup) - Match regex expression named groups into go struct using struct tags and automatic parsing. Stars:`147`.
- [goregen](https://github.com/zach-klippenstein/goregen) - Library for generating random strings from regular expressions. Stars:`92`.
- [go-wildcard](https://github.com/IGLOU-EU/go-wildcard) - Simple and lightweight wildcard pattern matching. Stars:`87`.
- [genex](https://github.com/alixaxel/genex) - Count and expand Regular Expressions into all matching Strings. Stars:`76`.

### Sanitation

- [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer. Stars:`3.5K`.
- [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) - A sanitization-based swear filter for Go. Stars:`68`.

### Scrapers

- [colly](https://github.com/asciimoo/colly) - Fast and Elegant Scraping Framework for Gophers. Stars:`24.5K`.
- [GoQuery](https://github.com/PuerkitoBio/goquery) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language. Stars:`14.6K`.
- [xurls](https://github.com/mvdan/xurls) - Extract urls from text. Stars:`1.2K`.
- [dataflowkit](https://github.com/slotix/dataflowkit) - Web scraping Framework to turn websites into structured data. Stars:`688`.
- [pagser](https://github.com/foolin/pagser) - Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler. Stars:`109`.
- [Tagify](https://github.com/zoomio/tagify) - Produces a set of tags from given source. Stars:`39`.
- [go-recipe](https://github.com/kkyr/go-recipe) - A package for scraping recipes from websites. Stars:`30`.
- [walker](https://github.com/cyucelen/walker) - Seamlessly fetch paginated data from any source. Simple and high performance API scraping included. Stars:`14`.
- [go-sitemap-parser](https://github.com/aafeher/go-sitemap-parser) - Go language library for parsing Sitemaps. Stars:`4`.

### RSS

- [podcast](https://github.com/eduncan911/podcast) - iTunes Compliant and RSS 2.0 Podcast Generator in Golang Stars:`136`.

### Utility/Miscellaneous

- [w2vgrep](https://github.com/arunsupe/semantic-grep) - A semantic grep tool using word embeddings to find semantically similar matches. For example, searching for "death" will find "dead", "killing", "murder". Stars:`1.2K`.
- [go-runewidth](https://github.com/mattn/go-runewidth) - Functions to get fixed width of the character or string. Stars:`646`.
- [radix](https://github.com/yourbasic/radix) - Fast string sorting algorithm. Stars:`193`.
- [go-zero-width](https://github.com/trubitsyn/go-zero-width) - Zero-width character detection and removal for Go. Stars:`110`.
- [petrovich](https://github.com/striker2000/petrovich) - Petrovich is the library which inflects Russian names to given grammatical case. Stars:`50`.
- [kace](https://github.com/codemodus/kace) - Common case conversions covering common initialisms. Stars:`20`.
- [TySug](https://github.com/Dynom/TySug) - Alternative suggestions with respect to keyboard layouts. Stars:`19`.

**[⬆ back to top](#contents)**

## Third-party APIs

_Libraries for accessing third party APIs._

- [github](https://github.com/google/go-github) - Go library for accessing the GitHub REST API v3. Stars:`10.9K`.
- [go-openai](https://github.com/sashabaranov/go-openai) - OpenAI ChatGPT, DALL·E, Whisper API library for Go. Stars:`10.2K`.
- [discordgo](https://github.com/bwmarrin/discordgo) - Go bindings for the Discord Chat API. Stars:`5.5K`.
- [slack](https://github.com/slack-go/slack) - Slack API in Go. Stars:`4.8K`.
- [google](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go. Stars:`4.3K`.
- [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library. Stars:`4.3K`.
- [aws-sdk-go](https://github.com/aws/aws-sdk-go-v2) - The official AWS SDK for the Go programming language. Stars:`3.2K`.
- [minio-go](https://github.com/minio/minio-go) - Minio Go Library for Amazon S3 compatible cloud storage. Stars:`2.8K`.
- [stripe](https://github.com/stripe/stripe-go) - Go client for the Stripe API. Stars:`2.4K`.
- [go-jira](https://github.com/andygrunwald/go-jira) - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira) Stars:`1.6K`.
- [facebook](https://github.com/huandu/facebook) - Go Library that supports the Facebook Graph API. Stars:`1.4K`.
- [githubql](https://github.com/shurcooL/githubql) - Go library for accessing the GitHub GraphQL API v4. Stars:`1.2K`.
- [anaconda](https://github.com/ChimeraCoder/anaconda) - Go client library for the Twitter 1.1 API. Stars:`1.1K`.
- [webhooks](https://github.com/go-playground/webhooks) - Webhook receiver for GitHub and Bitbucket. Stars:`992`.
- [paypal](https://github.com/logpacker/PayPal-Go-SDK) - Wrapper for PayPal payment API. Stars:`742`.
- [geo-golang](https://github.com/codingsince1985/geo-golang) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](https://developer.mapquest.com/documentation/api/geocoding/), [Nominatim](https://nominatim.org/release-docs/latest/api/Overview/), [OpenCage](https://opencagedata.com/api), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. Stars:`535`.
- [lark](https://github.com/chyroc/lark) - [Feishu](https://open.feishu.cn/)/[Lark](https://open.larksuite.com/) Open API Go SDK, Support ALL Open API and Event Callback. Stars:`452`.
- [openaigo](https://github.com/otiai10/openaigo) - OpenAI GPT3/GPT3.5 ChatGPT API client library for Go. Stars:`296`.
- [ethrpc](https://github.com/onrik/ethrpc) - Go bindings for Ethereum JSON RPC API. Stars:`273`.
- [Trello](https://github.com/adlio/trello) - Go wrapper for the Trello API. Stars:`225`.
- [go-lark](https://github.com/go-lark/lark) - An easy-to-use unofficial SDK for [Feishu](https://open.feishu.cn/) and [Lark](https://open.larksuite.com/) Open Platform. Stars:`225`.
- [go-marathon](https://github.com/gambol99/go-marathon) - Go library for interacting with Mesosphere's Marathon PAAS. Stars:`200`.
- [wit-go](https://github.com/wit-ai/wit-go) - Go client for wit.ai HTTP API. Stars:`169`.
- [go-atlassian](https://github.com/ctreminiom/go-atlassian) - Go library for accessing the [Atlassian Cloud](https://www.atlassian.com/enterprise/cloud) services (Jira, Jira Service Management, Jira Agile, Confluence, Admin Cloud) Stars:`167`.
- [simples3](https://github.com/rhnvrm/simples3) - Simple no frills AWS S3 Library using REST with V4 Signing written in Go. Stars:`166`.
- [gosip](https://github.com/koltyakov/gosip) - Client library for SharePoint. Stars:`159`.
- [pushover](https://github.com/gregdel/pushover) - Go wrapper for the Pushover API. Stars:`146`.
- [go-trending](https://github.com/andygrunwald/go-trending) - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. Stars:`145`.
- [Medium](https://github.com/Medium/medium-sdk-go) - Golang SDK for Medium's OAuth2 API. Stars:`142`.
- [golang-tmdb](https://github.com/cyruzin/golang-tmdb) - Golang wrapper for The Movie Database API v3. Stars:`138`.
- [gostorm](https://github.com/jsgilmore/gostorm) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. Stars:`129`.
- [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP. Stars:`111`.
- [hipchat](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API. Stars:`104`.
- [disgo](https://github.com/switchupcb/disgo) - Go API Wrapper for the Discord API. Stars:`104`.
- [go-gerrit](https://github.com/andygrunwald/go-gerrit) - Go client library for [Gerrit Code Review](https://www.gerritcodereview.com/). Stars:`100`.
- [cachet](https://github.com/andygrunwald/cachet) - Go client library for [Cachet (open source status page system)](https://cachethq.io/). Stars:`90`.
- [go-redoc](https://github.com/mvrilo/go-redoc) - Embedded OpenAPI/Swagger documentation ui for Go using [ReDoc](https://redocly.com/). Stars:`87`.
- [GoFreeDB](https://github.com/FreeLeh/GoFreeDB) - Golang library providing common and simple database abstractions on top of Google Sheets. Stars:`86`.
- [go-postman-collection](https://github.com/rbretecher/go-postman-collection) - Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia). Stars:`86`.
- [gogtrends](https://github.com/groovili/gogtrends) - Google Trends Unofficial API. Stars:`85`.
- [igdb](https://github.com/Henry-Sarabia/igdb) - Go client for the [Internet Game Database API](https://api.igdb.com/). Stars:`84`.
- [airtable](https://github.com/mehanizm/airtable) - Go client library for the [Airtable API](https://airtable.com/api). Stars:`82`.
- [go-unsplash](https://github.com/hbagdi/go-unsplash) - Go client library for the [Unsplash.com](https://unsplash.com) API. Stars:`76`.
- [ynab](https://github.com/brunomvsouza/ynab.go) - Go wrapper for the YNAB API. Stars:`74`.
- [circleci](https://github.com/jszwedko/go-circleci) - Go client library for interacting with CircleCI's API. Stars:`65`.
- [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) - Go MusicBrainz WS2 client library. Stars:`61`.
- [mixpanel](https://github.com/dukex/mixpanel) - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. Stars:`60`.
- [uptimerobot](https://github.com/bitfield/uptimerobot) - Go wrapper and command-line client for the Uptime Robot v2 API. Stars:`57`.
- [clarifai](https://github.com/samuelcouch/clarifai) - Go client library for interfacing with the Clarifai API. Stars:`55`.
- [megos](https://github.com/andygrunwald/megos) - Client library for accessing an [Apache Mesos](https://mesos.apache.org/) cluster. Stars:`54`.
- [go-xkcd](https://github.com/nishanths/go-xkcd) - Go client for the xkcd API. Stars:`51`.
- [gads](https://github.com/emiddleton/gads) - Google Adwords Unofficial API. Stars:`51`.
- [fcm](https://github.com/maddevsio/fcm) - Go library for Firebase Cloud Messaging. Stars:`51`.
- [spotify](https://github.com/rapito/go-spotify) - Go Library to access Spotify WEB API. Stars:`51`.
- [swag](https://github.com/zc2638/swag) - No comments, simple go wrapper to create swagger 2.0 compatible APIs. Support most routing frameworks, such as built-in, gin, chi, mux, echo, httprouter, fasthttp and more. Stars:`49`.
- [go-salesforce](https://github.com/k-capehart/go-salesforce) - Go client library for interacting with the [Salesforce REST API](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/api_rest/resources_list.htm). Stars:`44`.
- [patreon-go](https://github.com/mxpv/patreon-go) - Go library for Patreon API. Stars:`44`.
- [go-myanimelist](https://github.com/nstratos/go-myanimelist) - Go client library for accessing the [MyAnimeList API](https://myanimelist.net/apiconfig/references/api/v2). Stars:`41`.
- [golyrics](https://github.com/mamal72/golyrics) - Golyrics is a Go library to fetch music lyrics data from the Wikia website. Stars:`41`.
- [lastpass-go](https://github.com/ansd/lastpass-go) - Go client library for the [LastPass](https://www.lastpass.com/) API. Stars:`35`.
- [steam](https://github.com/sostronk/go-steam) - Go Library to interact with Steam game servers. Stars:`33`.
- [gcm](https://github.com/Aorioli/gcm) - Go library for Google Cloud Messaging. Stars:`31`.
- [jokeapi-go](https://github.com/icelain/jokeapi) - Go client for [JokeAPI](https://sv443.net/jokeapi/v2/). Stars:`25`.
- [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) - Go client library for interacting with Coinpaprika's API. Stars:`25`.
- [shopify](https://github.com/rapito/go-shopify) - Go Library to make CRUD request to the Shopify API. Stars:`24`.
- [device-check-go](https://github.com/rinchsan/device-check-go) - Go client library for interacting with [iOS DeviceCheck API](https://developer.apple.com/documentation/devicecheck) v1. Stars:`24`.
- [go-imgur](https://github.com/koffeinsource/go-imgur) - Go client library for [imgur](https://imgur.com) Stars:`23`.
- [aws-encryption-sdk-go](https://github.com/chainifynet/aws-encryption-sdk-go) - Unofficial Go SDK implementation of the [AWS Encryption SDK](https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/index.html). Stars:`21`.
- [brewerydb](https://github.com/naegelejd/brewerydb) - Go library for accessing the BreweryDB API. Stars:`20`.
- [textbelt](https://github.com/dietsche/textbelt) - Go client for the textbelt.com txt messaging API. Stars:`19`.
- [go-openproject](https://github.com/manuelbcd/go-openproject) - Go client library for interacting with [OpenProject](https://docs.openproject.org/api/) API. Stars:`19`.
- [codeship-go](https://github.com/codeship/codeship-go) - Go client library for interacting with Codeship's API v2. Stars:`18`.
- [go-aws-news](https://github.com/circa10a/go-aws-news) - Go application and library to fetch what's new from AWS. Stars:`17`.
- [goami2](https://github.com/staskobzar/goami2) - AMI v2 library for Asterisk PBX. Stars:`17`.
- [go-hacknews](https://github.com/PaulRosset/go-hacknews) - Tiny Go client for HackerNews API. Stars:`17`.
- [gopaapi5](https://github.com/utekaravinash/gopaapi5) - Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/). Stars:`16`.
- [bqwriter](https://github.com/OTA-Insight/bqwriter) - High Level Go Library to write data into [Google BigQuery](https://cloud.google.com/bigquery) at a high throughout. Stars:`16`.
- [google-analytics](https://github.com/chonthu/go-google-analytics) - Simple wrapper for easy google analytics reporting. Stars:`15`.
- [go-here](https://github.com/abdullahselek/go-here) - Go client library around the HERE location based APIs. Stars:`13`.
- [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) - Go library for the [RAWG Video Games Database](https://rawg.io/) API Stars:`13`.
- [ip2location-io-go](https://github.com/ip2location/ip2location-io-go) - Go wrapper for the IP2Location.io API [IP2Location.io](https://www.ip2location.io/). Stars:`13`.
- [go-sophos](https://github.com/esurdam/go-sophos) - Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. Stars:`12`.
- [gomalshare](https://github.com/MonaxGT/gomalshare) - Go library MalShare API [malshare.com](https://www.malshare.com/) Stars:`12`.
- [goagi](https://github.com/staskobzar/goagi) - Go library to build Asterisk PBX agi/fastagi applications. Stars:`11`.
- [go-swagger-ui](https://github.com/esurdam/go-swagger-ui) - Go library containing precompiled [Swagger UI](https://swagger.io/tools/swagger-ui/) for serving swagger json. Stars:`11`.
- [smite](https://github.com/sergiotapia/smitego) - Go package to wraps access to the Smite game API. Stars:`11`.
- [libgoffi](https://github.com/clevabit/libgoffi) - Library adapter toolbox for native [libffi](https://sourceware.org/libffi/) integration Stars:`10`.
- [go-hibp](https://github.com/wneessen/go-hibp) - Simple Go binding to the "Have I Been Pwned" APIs. Stars:`10`.
- [go-chronos](https://github.com/axelspringer/go-chronos) - Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler Stars:`8`.
- [tumblr](https://github.com/mattcunningham/gumblr) - Go wrapper for the Tumblr v2 API. Stars:`8`.
- [newsapi-go](https://github.com/jellydator/newsapi-go) - Go client for [NewsAPI](https://newsapi.org/). Stars:`7`.
- [zooz](https://github.com/gojuno/go-zooz) - Go client for the Zooz API. Stars:`7`.
- [appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go) - Unofficial Golang SDK for AppStore Connect API. Stars:`6`.
- [go-restcountries](https://github.com/chriscross0/go-restcountries) - Go library for the [REST Countries API](https://countrylayer.com/). Stars:`3`.
- [dusupay-sdk-go](https://github.com/Kachit/dusupay-sdk-go) - Unofficial Dusupay payment gateway API Client for Go Stars:`3`.
- [gopensky](https://github.com/navidys/gopensky) - Go client implementation for [OpenSKY Network](https://opensky-network.org/) live's API (airspace ADS-B and Mode S data). Stars:`2`.
- [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) - Go wrapper for the TripAdvisor API. Stars:`2`.
- [vl-go](https://github.com/verifid/vl-go) - Go client library around the VerifID identity verification layer API. Stars:`2`.
- [fasapay-sdk-go](https://github.com/Kachit/fasapay-sdk-go) - Unofficial Fasapay payment gateway XML API Client for Golang. Stars:`2`.
- [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) - The Playlyfe Rest API Go SDK. Stars:`2`.
- [go-telegraph](https://gitlab.com/toby3d/telegraph) - Telegraph publishing platform API client.
- [go-yapla](https://gitlab.com/adrienK/go-yapla) - Go client library for the Yapla v2.0 API.

**[⬆ back to top](#contents)**

## Utilities

_General utilities and tools to make your life easier._

- [fzf](https://github.com/junegunn/fzf) - Command-line fuzzy finder written in Go. Stars:`72.5K`.
- [dive](https://github.com/wagoodman/dive) - A tool for exploring each layer in a Docker image. Stars:`51.5K`.
- [hub](https://github.com/github/hub) - wrap git commands with additional functionality to interact with github from the terminal. Stars:`22.9K`.
- [lo](https://github.com/samber/lo) - A Lodash like Go library based on Go 1.18+ Generics (map, filter, contains, find...) Stars:`19.8K`.
- [sqlx](https://github.com/jmoiron/sqlx) - provides a set of extensions on top of the excellent built-in database/sql package. Stars:`17.1K`.
- [ctop](https://github.com/bcicen/ctop) - [Top-like](https://ctop.sh) interface (e.g. htop) for container metrics. Stars:`16.7K`.
- [goreleaser](https://github.com/goreleaser/goreleaser) - Deliver Go binaries as fast and easily as possible. Stars:`15.0K`.
- [wuzz](https://github.com/asciimoo/wuzz) - Interactive cli tool for HTTP inspection. Stars:`10.7K`.
- [usql](https://github.com/knq/usql) - usql is a universal command-line interface for SQL databases. Stars:`9.5K`.
- [peco](https://github.com/peco/peco) - Simplistic interactive filtering tool. Stars:`7.8K`.
- [lancet](https://github.com/duke-git/lancet) - A comprehensive, efficient, and reusable util function library of go. Stars:`5.2K`.
- [go-funk](https://github.com/thoas/go-funk) - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). Stars:`4.9K`.
- [godropbox](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications from Dropbox. Stars:`4.2K`.
- [minify](https://github.com/tdewolff/minify) - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. Stars:`3.9K`.
- [panicparse](https://github.com/maruel/panicparse) - Groups similar goroutines and colorizes stack dump. Stars:`3.7K`.
- [goreporter](https://github.com/wgliang/goreporter) - Golang tool that does static analysis, unit testing, code review and generate code quality report. Stars:`3.1K`.
- [mc](https://github.com/minio/mc) - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. Stars:`3.1K`.
- [mergo](https://github.com/imdario/mergo) - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. Stars:`3.0K`.
- [retry-go](https://github.com/avast/retry-go) - Simple library for retry mechanism. Stars:`2.7K`.
- [create-go-app](https://github.com/create-go-app/cli) - A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command. Stars:`2.7K`.
- [EaseProbe](https://github.com/megaease/easeprobe) - A simple, standalone, and lightWeight tool that can do health/status checking daemon, support HTTP/TCP/SSH/Shell/Client/... probes, and Slack/Discord/Telegram/SMS... notification. Stars:`2.3K`.
- [filetype](https://github.com/h2non/filetype) - Small package to infer the file type checking the magic numbers signature. Stars:`2.2K`.
- [Storm](https://github.com/asdine/storm) - Simple and powerful toolkit for BoltDB. Stars:`2.1K`.
- [jump](https://github.com/gsamokovarov/jump) - Jump helps you navigate faster by learning your habits. Stars:`1.9K`.
- [mimetype](https://github.com/gabriel-vasile/mimetype) - Package for MIME type detection based on magic numbers. Stars:`1.8K`.
- [Failsafe-go](https://github.com/failsafe-go/failsafe-go) - Fault tolerance and resilience patterns for Go. Stars:`1.8K`.
- [boilr](https://github.com/tmrts/boilr) - Blazingly fast CLI tool for creating projects from boilerplate templates. Stars:`1.7K`.
- [mole](https://github.com/davrodpin/mole) - cli app to easily create ssh tunnels. Stars:`1.7K`.
- [gitbatch](https://github.com/isacikgoz/gitbatch) - manage your git repositories in one place. Stars:`1.6K`.
- [scany](https://github.com/georgysavva/scany) - Library for scanning data from a database into Go structs and more. Stars:`1.5K`.
- [bed](https://github.com/itchyny/bed) - A Vim-like binary editor written in Go. Stars:`1.3K`.
- [sesh](https://github.com/joshmedeski/sesh) - Sesh is a CLI that helps you create and manage tmux sessions quickly and easily using zoxide. Stars:`1.2K`.
- [hostctl](https://github.com/guumaster/hostctl) - A CLI tool to manage /etc/hosts with easy commands. Stars:`1.2K`.
- [circuitbreaker](https://github.com/rubyist/circuitbreaker) - Circuit Breakers in Go. Stars:`1.1K`.
- [git-time-metric](https://github.com/git-time-metric/gtm) - Simple, seamless, lightweight time tracking for Git. Stars:`997`.
- [upterm](https://github.com/owenthereal/upterm) - A tool for developers to share terminal/tmux sessions securely over the web. It’s perfect for remote pair programming, accessing computers behind NATs/firewalls, remote debugging, and more. Stars:`949`.
- [immortal](https://github.com/immortal/immortal) - \*nix cross-platform (OS agnostic) supervisor. Stars:`824`.
- [circuit](https://github.com/cep21/circuit) - An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. Stars:`786`.
- [changie](https://github.com/miniscruff/changie) - Automated changelog tool for preparing releases with lots of customization options. Stars:`783`.
- [clockwork](https://github.com/jonboulle/clockwork) - A simple fake clock for golang. Stars:`710`.
- [clipboard](https://github.com/golang-design/clipboard) - 📋 cross-platform clipboard package in Go. Stars:`708`.
- [delve](https://github.com/derekparker/delve) - Go debugger. Stars:`654`.
- [ergo](https://github.com/cristianoliveira/ergo) - The management of multiple local services running over different ports made easy. Stars:`636`.
- [remote-touchpad](https://github.com/Unrud/remote-touchpad) - Control mouse and keyboard from a smartphone. Stars:`612`.
- [scan](https://github.com/blockloop/scan) - Scan golang `sql.Rows` directly to structs, slices, or primitive types. Stars:`597`.
- [limiters](https://github.com/mennanov/limiters) - Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. Stars:`577`.
- [mani](https://github.com/alajmo/mani) - CLI tool to help you manage multiple repositories. Stars:`568`.
- [htcat](https://github.com/htcat/htcat) - Parallel and Pipelined HTTP GET Utility. Stars:`556`.
- [koazee](https://github.com/wesovilabs/koazee) - Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. Stars:`531`.
- [gubrak](https://github.com/novalagung/gubrak) - Golang utility library with syntactic sugar. It's like lodash, but for golang. Stars:`500`.
- [godaemon](https://github.com/VividCortex/godaemon) - Utility to write daemons. Stars:`495`.
- [go-dry](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go. Stars:`487`.
- [countries](https://github.com/biter777/countries) - Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standards. Stars:`462`.
- [gopencils](https://github.com/bndr/gopencils) - Small and simple package to easily consume REST APIs. Stars:`453`.
- [request](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™. Stars:`425`.
- [go-rate](https://github.com/beefsack/go-rate) - Timed rate limiter for Go. Stars:`403`.
- [config-file-validator](https://github.com/Boeing/config-file-validator) - Cross Platform tool to validate configuration files. Stars:`384`.
- [grofer](https://github.com/pesos/grofer) - A system and resource monitoring tool written in Golang! Stars:`375`.
- [serve](https://github.com/syntaqx/serve) - A static http server anywhere you need. Stars:`347`.
- [retry](https://github.com/kamilsk/retry) - The most advanced functional mechanism to perform actions repetitively until successful. Stars:`345`.
- [rospo](https://github.com/ferama/rospo) - Simple and reliable ssh tunnels with embedded ssh server in Golang. Stars:`341`.
- [gotenv](https://github.com/subosito/gotenv) - Load environment variables from `.env` or any `io.Reader` in Go. Stars:`306`.
- [util](https://github.com/shomali11/util) - Collection of useful utility functions. (strings, concurrency, manipulations, ...). Stars:`299`.
- [wifiqr](https://github.com/reugn/wifiqr) - Wi-Fi QR Code Generator. Stars:`274`.
- [pattern-match](https://github.com/alexpantyukhin/go-pattern-match) - Pattern matching library. Stars:`250`.
- [go-trigger](https://github.com/sadlil/go-trigger) - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. Stars:`249`.
- [toolbox](https://github.com/viant/toolbox) - Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. Stars:`225`.
- [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) - XML Sitemap generator written in Go. Stars:`225`.
- [Death](https://github.com/vrecan/death) - Managing go application shutdown with signals. Stars:`198`.
- [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) - go:generate tool for wrapping symbols exported by golang plugins (1.8 only). Stars:`181`.
- [goval](https://github.com/maja42/goval) - Evaluate arbitrary expressions in Go. Stars:`172`.
- [moldova](https://github.com/StabbyCutyou/moldova) - Utility for generating random data based on an input template. Stars:`169`.
- [apm](https://github.com/topfreegames/apm) - Process manager for Golang applications with an HTTP API. Stars:`167`.
- [rerun](https://github.com/ivpusic/rerun) - Recompiling and rerunning go apps when source changes. Stars:`166`.
- [go-bsdiff](https://github.com/gabstv/go-bsdiff) - Pure Go bsdiff and bspatch libraries and CLI tools. Stars:`166`.
- [json-log-viewer](https://github.com/hedhyw/json-log-viewer) - Interactive viewer for JSON logs. Stars:`160`.
- [cmd](https://github.com/SimonBaeumer/cmd) - Library for executing shell commands on osx, windows and linux. Stars:`160`.
- [robustly](https://github.com/VividCortex/robustly) - Runs functions resiliently, catching and restarting panics. Stars:`159`.
- [chyle](https://github.com/antham/chyle) - Changelog generator using a git repository with multiple configuration possibilities. Stars:`158`.
- [cryptgo](https://github.com/Gituser143/cryptgo) - Crytpgo is a TUI based application written purely in Go to monitor and observe cryptocurrency prices in real time! Stars:`154`.
- [filter](https://github.com/gookit/filter) - provide filtering, sanitizing, and conversion of Go data. Stars:`151`.
- [nostromo](https://github.com/pokanop/nostromo) - CLI for building powerful aliases. Stars:`146`.
- [sorty](https://github.com/jfcg/sorty) - Fast Concurrent / Parallel Sorting. Stars:`137`.
- [onecache](https://github.com/adelowo/onecache) - Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). Stars:`136`.
- [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) - Mongodb Pagination for official mongodb/mongo-go-driver package which supports both normal queries and Aggregation pipelines. Stars:`132`.
- [lrserver](https://github.com/jaschaephraim/lrserver) - LiveReload server for Go. Stars:`128`.
- [gitcs](https://github.com/knbr13/gitcs/) - Git Commits Visualizer, CLI tool to visualize your Git commits on your local machine. Stars:`124`.
- [go-lock](https://github.com/viney-shih/go-lock) - go-lock is a lock library implementing read-write mutex and read-write trylock without starvation. Stars:`120`.
- [goseaweedfs](https://github.com/linxGnu/goseaweedfs) - SeaweedFS client library with almost full features. Stars:`116`.
- [cookie](https://github.com/syntaqx/cookie) - Cookie struct parsing and helper package. Stars:`111`.
- [mssqlx](https://github.com/linxGnu/mssqlx) - Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. Stars:`102`.
- [xferspdy](https://github.com/monmohan/xferspdy) - Xferspdy provides binary diff and patch library in golang. Stars:`102`.
- [mimemagic](https://github.com/zRedShift/mimemagic) - Pure Go ultra performant MIME sniffing library/utility. Stars:`100`.
- [go-health](https://github.com/Talento90/go-health) - Health package simplifies the way you add health check to your services. Stars:`96`.
- [go-lambda-cleanup](https://github.com/karl-cardenas-coding/go-lambda-cleanup) - A CLI for removing unused or previous versions of AWS Lambdas. Stars:`95`.
- [go-pattern-match](https://github.com/PhakornKiong/go-pattern-match) - A Pattern matching library inspired by ts-pattern. Stars:`93`.
- [countries](https://github.com/pioz/countries) - All you need when you are working with countries in Go. Stars:`92`.
- [equalizer](https://github.com/reugn/equalizer) - Quota manager and rate limiter collection for Go. Stars:`91`.
- [pgo](https://github.com/arthurkushman/pgo) - Convenient functions for PHP community. Stars:`88`.
- [repeat](https://github.com/ssgreg/repeat) - Go implementation of different backoff strategies useful for retrying operations and heartbeating. Stars:`85`.
- [handy](https://github.com/miguelpragier/handy) - Many utilities and helpers like string handlers/formatters and validators. Stars:`82`.
- [pm](https://github.com/VividCortex/pm) - Process (i.e. goroutine) manager with an HTTP API. Stars:`79`.
- [netbug](https://github.com/e-dard/netbug) - Easy remote profiling of your services. Stars:`72`.
- [UNIS](https://github.com/esemplastic/unis) - Common Architecture™ for String Utilities in Go. Stars:`70`.
- [multitick](https://github.com/VividCortex/multitick) - Multiplexor for aligned tickers. Stars:`70`.
- [goreadability](https://github.com/philipjkim/goreadability) - Webpage summary extractor using Facebook Open Graph and arc90's readability. Stars:`69`.
- [dbt](https://github.com/nikogura/dbt) - A framework for running self-updating signed binaries from a central, trusted repository. Stars:`68`.
- [backscanner](https://github.com/icza/backscanner) - A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. Stars:`67`.
- [scan](https://github.com/wroge/scan) - Scan sql rows into any type powered by generics. Stars:`67`.
- [go-astitodo](https://github.com/asticode/go-astitodo) - Parse TODOs in your GO code. Stars:`66`.
- [retry](https://github.com/thedevsaddam/retry) - Simple and easy retry mechanism package for Go. Stars:`66`.
- [golog](https://github.com/mlimaloureiro/golog) - Easy and lightweight CLI tool to time track your tasks. Stars:`63`.
- [minquery](https://github.com/icza/minquery) - MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). Stars:`62`.
- [shutdown](https://github.com/ztrue/shutdown) - App shutdown hooks for `os.Signal` handling. Stars:`60`.
- [copy-pasta](https://github.com/jutkko/copy-pasta) - Universal multi-workstation clipboard that uses S3 like backend for the storage. Stars:`54`.
- [golarm](https://github.com/msempere/golarm) - Fire alarms with system events. Stars:`54`.
- [cvt](https://github.com/shockerli/cvt) - Easy and safe convert any value to another type. Stars:`53`.
- [gofn](https://github.com/tiendc/gofn) - High performance utility functions written using Generics for Go 1.18+. Stars:`53`.
- [sshman](https://github.com/shoobyban/sshman) - SSH Manager for authorized_keys files on multiple remote servers. Stars:`51`.
- [slice](https://github.com/psampaz/slice) - Type-safe functions for common Go slice operations. Stars:`51`.
- [goback](https://github.com/carlescere/goback) - Go simple exponential backoff package. Stars:`50`.
- [retry-go](https://github.com/rafaeljesus/retry-go) - Retrying made simple and easy for golang. Stars:`50`.
- [go-httpheader](https://github.com/mozillazg/go-httpheader) - Go library for encoding structs into Header fields. Stars:`47`.
- [ghokin](https://github.com/antham/ghokin) - Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). Stars:`47`.
- [gostrutils](https://github.com/ik5/gostrutils) - Collections of string manipulation and conversion functions. Stars:`47`.
- [slicer](https://github.com/leaanthony/slicer) - Makes working with slices easier. Stars:`47`.
- [set](https://github.com/nofeaturesonlybugs/set) - Performant and flexible struct mapping and loose type conversion. Stars:`47`.
- [yogo](https://github.com/antham/yogo) - Check yopmail mails from command line. Stars:`46`.
- [pointer](https://github.com/xorcare/pointer) - Package pointer contains helper routines for simplifying the creation of optional fields of basic type. Stars:`45`.
- [evaluator](https://github.com/nullne/evaluator) - Evaluate an expression dynamically based on s-expression. It's simple and easy to extend. Stars:`42`.
- [go-qr](https://github.com/piglig/go-qr) - A native, high-quality and minimalistic QR code generator. Stars:`41`.
- [throttle](https://github.com/yudppp/throttle) - Throttle is an object that will perform exactly one action per duration. Stars:`40`.
- [just](https://github.com/kazhuravlev/just) - Just a collection of useful functions for working with generic data structures. Stars:`36`.
- [rclient](https://github.com/zpatrick/rclient) - Readable, flexible, simple-to-use client for REST APIs. Stars:`35`.
- [tome](https://github.com/cyruzin/tome) - Tome was designed to paginate simple RESTful APIs. Stars:`35`.
- [watchhttp](https://github.com/nikolaydubina/watchhttp) - Run command periodically and expose latest STDOUT or its rich delta as HTTP endpoint. Stars:`35`.
- [mimesniffer](https://github.com/aofei/mimesniffer) - A MIME type sniffer for Go. Stars:`34`.
- [git-tools](https://github.com/kazhuravlev/git-tools) - Tool to help manage git tags. Stars:`33`.
- [generate](https://github.com/go-playground/generate) - runs go generate recursively on a specified path or environment variable and can filter by regex. Stars:`30`.
- [debounce](https://github.com/floatdrop/debounce) - A zero-allocation debouncer written in Go. Stars:`29`.
- [goplaceholder](https://github.com/michiwend/goplaceholder) - a small golang lib to generate placeholder images. Stars:`29`.
- [graterm](https://github.com/skovtunenko/graterm) - Provides primitives to perform ordered (sequential/concurrent) GRAceful TERMination (aka shutdown) in Go application. Stars:`28`.
- [ptr](https://github.com/gotidy/ptr) - Package that provide functions for simplified creation of pointers from constants of basic types. Stars:`28`.
- [ctxutil](https://github.com/posener/ctxutil) - A collection of utility functions for contexts. Stars:`25`.
- [structs](https://github.com/PumpkinSeed/structs) - Implement simple functions to manipulate structs. Stars:`24`.
- [go-convert](https://github.com/Eun/go-convert) - Package go-convert enables you to convert a value into another type. Stars:`23`.
- [jsend](https://github.com/clevergo/jsend) - JSend's implementation written in Go. Stars:`22`.
- [healthcheck](https://github.com/kazhuravlev/healthcheck) - A simple yet powerful readiness test for Kubernetes. Stars:`21`.
- [go-type](https://github.com/mikekonan/go-types) - Library providing Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types. Stars:`21`.
- [filler](https://github.com/yaronsumel/filler) - small utility to fill structs using "fill" tag. Stars:`18`.
- [dlog](https://github.com/kirillDanshin/dlog) - Compile-time controlled logger to make your release smaller without removing debug calls. Stars:`17`.
- [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) - Go package for working with Problem Details. Stars:`17`.
- [rest-go](https://github.com/edermanoel94/rest-go) - A package that provide many helpful methods for working with rest api. Stars:`16`.
- [go-countries](https://github.com/mikekonan/go-countries) - Lightweight lookup over ISO-3166 codes. Stars:`16`.
- [okrun](https://github.com/xta/okrun) - go run error steamroller. Stars:`16`.
- [go-tripper](https://github.com/rajnandan1/go-tripper) - Tripper is a circuit breaker package for Go that allows you to circuit and control the status of circuits. Stars:`14`.
- [silk](https://github.com/chrispassas/silk) - Read silk netflow files. Stars:`14`.
- [go-clip](https://github.com/prashantgupta24/go-clip) - A minimalistic clipboard manager for Mac. Stars:`14`.
- [retry](https://github.com/shafreeck/retry) - A pretty simple library to ensure your work to be done. Stars:`13`.
- [blank](https://github.com/Henry-Sarabia/blank) - Verify or remove blanks and whitespace from strings. Stars:`13`.
- [command](https://github.com/txgruppi/command) - Command pattern for Go with thread safe serial and parallel dispatcher. Stars:`13`.
- [jet](https://github.com/NicoNex/jet) - Just Edit Text: a fast and powerful tool for finding and replacing file content and names using regular expressions. Stars:`12`.
- [loncha](https://github.com/kazu/loncha) - A high-performance slice Utilities. Stars:`11`.
- [nfdump](https://github.com/chrispassas/nfdump) - Read nfdump netflow files. Stars:`11`.
- [statiks](https://github.com/janiltonmaciel/statiks) - Fast, zero-configuration, static HTTP filer server. Stars:`11`.
- [bleep](https://github.com/sinhashubham95/bleep) - Perform any number of actions on any set of OS signals in Go. Stars:`11`.
- [retry](https://github.com/percolate/retry) - A simple but highly configurable retry package for Go. Stars:`11`.
- [goctx](https://github.com/zerosnake0/goctx) - Get your context value with high performance. Stars:`10`.
- [optional](https://github.com/kazhuravlev/optional) - Optional struct fields and vars. Stars:`10`.
- [sliceconv](https://github.com/Henry-Sarabia/sliceconv) - Slice conversion between primitive types. Stars:`9`.
- [go-events](https://github.com/deatil/go-events) - A go event and event'subscribe package, like wordpress hook functions. Stars:`7`.
- [lets-go](https://github.com/aplescia-chwy/lets-go) - Go module that provides common utilities for Cloud Native REST API development. Also contains AWS Specific utilities. Stars:`7`.
- [go-pkg](https://github.com/chenquan/go-pkg) - A go toolkit. Stars:`7`.
- [tik](https://github.com/andy2046/tik) - Simple and easy timing wheel package for Go. Stars:`5`.
- [olaf](https://github.com/btnguyen2k/olaf) - Twitter Snowflake implemented in Go. Stars:`5`.
- [xpool](https://github.com/peczenyj/xpool) - Yet another golang type safe object pool using generics. Stars:`4`.
- [go-actuator](https://github.com/sinhashubham95/go-actuator) - Production ready features for Go based web frameworks. Stars:`3`.
- [lang](https://github.com/maxbolgarin/lang) - Generic one-liners to work with variables, slices and maps without boilerplate code. Stars:`3`.

**[⬆ back to top](#contents)**

## UUID

_Libraries for working with UUIDs._

- [uuid](https://github.com/google/uuid) - Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. Stars:`5.8K`.
- [ulid](https://github.com/oklog/ulid) - Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). Stars:`4.8K`.
- [xid](https://github.com/rs/xid) - Xid is a globally unique id generator library, ready to be safely used directly in your server code. Stars:`4.2K`.
- [uuid](https://github.com/gofrs/uuid) - Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. Stars:`1.7K`.
- [wuid](https://github.com/edwingeng/wuid) - An extremely fast globally unique number generator. Stars:`541`.
- [sno](https://github.com/muyo/sno) - Compact, sortable and fast unique IDs with embedded metadata. Stars:`92`.
- [guid](https://github.com/sdrapkin/guid) - Fast cryptographically safe Guid generator for Go (~10x faster than `uuid`). Stars:`71`.
- [nanoid](https://github.com/aidarkhanov/nanoid) - A tiny and efficient Go unique string ID generator. Stars:`62`.
- [goid](https://github.com/jakehl/goid) - Generate and Parse RFC4122 compliant V4 UUIDs. Stars:`41`.
- [gouid](https://github.com/twharmon/gouid) - Generate cryptographically secure random string IDs with just one allocation. Stars:`26`.
- [uuid](https://github.com/agext/uuid) - Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. Stars:`18`.
- [uuidcheck](https://github.com/ashwingopalsamy/uuidcheck) - A tiny, dependency-free Go library that validates UUIDs against standard RFC 4122 formatting, converts UUIDv7() into UTC timestamps. Stars:`5`.
- [uniq](https://gitlab.com/skilstak/code/go/uniq) - No hassle safe, fast unique identifiers with commands.
- [fastuuid](https://github.com/rekby/fastuuid) - Fast generate UUIDv4 as string or bytes.

**[⬆ back to top](#contents)**

## Validation

_Libraries for validation._

- [validator](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. Stars:`18.8K`.
- [govalidator](https://github.com/asaskevich/govalidator) - Validators and sanitizers for strings, numerics, slices and structs. Stars:`6.2K`.
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. Stars:`4.0K`.
- [govalidator](https://github.com/thedevsaddam/govalidator) - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. Stars:`1.3K`.
- [validate](https://github.com/gookit/validate) - Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. Stars:`1.1K`.
- [Zog](https://github.com/Oudwins/zog) - A [Zod](https://github.com/colinhacks/zod) inspired schema builder for runtime value parsing and validation. Stars:`931`.
- [gody](https://github.com/guiferpa/gody) - :balloon: A lightweight struct validator for Go. Stars:`161`.
- [jio](https://github.com/faceair/jio) - jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). Stars:`118`.
- [checkdigit](https://github.com/osamingo/checkdigit) - Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). Stars:`111`.
- [govalid](https://github.com/twharmon/govalid) - Fast, tag-based validation for structs. Stars:`99`.
- [validate](https://github.com/gobuffalo/validate) - This package provides a framework for writing validations for Go applications. Stars:`93`.
- [valix](https://github.com/marrow16/valix) Go package for validating requests Stars:`31`.
- [go-validator](https://github.com/tiendc/go-validator) - Validation library using Generics. Stars:`31`.
- [hvalid](https://github.com/lyonnee/hvalid) hvalid is a lightweight validation library written in Go language. It provides a custom validator interface and a series of common validation functions to help developers quickly implement data validation. Stars:`12`.
- [Validator](https://github.com/go-the-way/validator) - A lightweight model validator written in Go.Contains VFs:Min, Max, MinLength, MaxLength, Length, Enum, Regex. Stars:`7`.
**[⬆ back to top](#contents)**

## Version Control

_Libraries for version control._

- [go-git](https://github.com/go-git/go-git) - highly extensible Git implementation in pure Go. Stars:`6.7K`.
- [hercules](https://github.com/src-d/hercules) - gaining advanced insights from Git repository history. Stars:`2.7K`.
- [git2go](https://github.com/libgit2/git2go) - Go bindings for libgit2. Stars:`2.0K`.
- [githooks](https://github.com/gabyx/githooks) - Per-repo and shared Git hooks with version control and auto update. Stars:`111`.
- [go-vcs](https://github.com/sourcegraph/go-vcs) - manipulate and inspect VCS repositories in Go. Stars:`80`.
- [froggit-go](https://github.com/jfrog/froggit-go) - Froggit-Go is a Go library, allowing to perform actions on VCS providers. Stars:`47`.
- [hgo](https://github.com/beyang/hgo) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories. Stars:`16`.
- [cli](https://gitlab.com/gitlab-org/cli) - An open-source GitLab command line tool bringing GitLab's cool features to your command line.

**[⬆ back to top](#contents)**

## Video

_Libraries for manipulating video._

- [goav](https://github.com/giorgisio/goav) - Comprehensive Go bindings for FFmpeg. Stars:`2.1K`.
- [m3u8](https://github.com/grafov/m3u8) - Parser and generator library of M3U8 playlists for Apple HLS. Stars:`1.3K`.
- [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries. Stars:`917`.
- [gortsplib](https://github.com/aler9/gortsplib) - Pure Go RTSP server and client library. Stars:`812`.
- [go-astisub](https://github.com/asticode/go-astisub) - Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). Stars:`648`.
- [go-astits](https://github.com/asticode/go-astits) - Parse and demux MPEG Transport Streams (.ts) natively in GO. Stars:`578`.
- [mp4ff](https://github.com/Eyevinn/mp4ff) - Library and tools for working with MP4 files containing video, audio, subtitles, or metadata. Stars:`539`.
- [go-astiav](https://github.com/asticode/go-astiav) - Better C bindings for ffmpeg in GO. Stars:`529`.
- [libvlc-go](https://github.com/adrg/libvlc-go) - Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). Stars:`476`.
- [v4l](https://github.com/korandiz/v4l) - Video capture library for Linux, written in Go. Stars:`85`.
- [go-m3u8](https://github.com/etherlabsio/go-m3u8) - Parser and generator library for Apple m3u8 playlists. Actively maintained version of quangngotan95/go-m3u8 with improvements and latest HLS playlist parsing compatibility. Stars:`35`.
- [go-mpd](https://github.com/unki2aut/go-mpd) - Parser and generator library for MPEG-DASH manifest files. Stars:`26`.

**[⬆ back to top](#contents)**

## Web Frameworks

_Full stack web frameworks._

- [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. Stars:`83.4K`.
- [Fiber](https://github.com/gofiber/fiber) - An Express.js inspired web framework build on Fasthttp. Stars:`37.3K`.
- [Beego](https://github.com/beego/beego) - beego is an open-source, high-performance web framework for the Go programming language. Stars:`32.2K`.
- [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework. Stars:`31.3K`.
- [Revel](https://github.com/revel/revel) - High-productivity web framework for the Go language. Stars:`13.2K`.
- [GoFr](https://github.com/gofr-dev/gofr) - Gofr is an opinionated microservice development framework. Stars:`12.7K`.
- [GoFrame](https://github.com/gogf/gf) - GoFrame is a modular, powerful, high-performance and enterprise-class application development framework of Golang. Stars:`12.6K`.
- [Hertz](https://github.com/cloudwego/hertz) - A high-performance and strong-extensibility Go HTTP framework that helps developers build microservices. Stars:`6.5K`.
- [Goa](https://github.com/goadesign/goa) - Goa provides a holistic approach for developing remote APIs and microservices in Go. Stars:`5.9K`.
- [goravel](https://github.com/goravel/goravel) - A Laravel-inspired web framework with ORM, authentication, queue, task scheduling, and more built-in features. Stars:`3.9K`.
- [Huma](https://github.com/danielgtaylor/huma/) - Framework for modern REST/GraphQL APIs with built-in OpenAPI 3, generated documentation, and a CLI. Stars:`3.3K`.
- [Goyave](https://github.com/go-goyave/goyave) - Feature-complete REST API framework aimed at clean code and fast development, with powerful built-in functionalities. Stars:`1.7K`.
- [Fuego](https://github.com/go-fuego/fuego) - The framework for busy Go developers! Web framework generating OpenAPI 3 spec from source code. Stars:`1.5K`.
- [Atreugo](https://github.com/savsgio/atreugo) - High performance and extensible micro web framework with zero memory allocations in hot paths. Stars:`1.3K`.
- [templui](https://github.com/axzilla/templui) - Modern UI Components for Go & Templ. Stars:`844`.
- [Yokai](https://github.com/ankorstore/yokai) - Simple, modular, and observable Go framework for backend applications. Stars:`740`.
- [iWF](https://github.com/indeedeng/iwf) - iWF is an all-in-one platform for developing long-running business processes. It offers a convenient abstraction for utilizing databases, ElasticSearch, message queues, durable timers, and more, with a clean, simple, and user-friendly interface. Stars:`583`.
- [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) - Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications. Stars:`568`.
- [rk-boot](https://github.com/rookie-ninja/rk-boot) - A bootstrapper library for building enterprise go microservice with Gin and gRPC quickly and easily. Stars:`545`.
- [Flamingo](https://github.com/i-love-flamingo/flamingo) - Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc. Stars:`540`.
- [Fastschema](https://github.com/fastschema/fastschema) - A flexible Go web framework and Headless CMS. Stars:`476`.
- [uAdmin](https://github.com/uadmin/uadmin) - Fully featured web framework for Golang, inspired by Django. Stars:`352`.
- [WebGo](https://github.com/naughtygopher/webgo) - A micro-framework to build web apps with handler chaining, middleware, and context injection. With standard library-compliant HTTP handlers (i.e., `http.HandlerFunc`).. Stars:`308`.
- [Ginrpc](https://github.com/xxjwxc/ginrpc) - Gin parameter automatic binding tool,gin rpc tools. Stars:`297`.
- [hiboot](https://github.com/hidevopsio/hiboot) - hiboot is a high performance web application framework with auto configuration and dependency injection support. Stars:`181`.
- [Gone](https://github.com/gone-io/gone) - A lightweight dependency injection and web framework inspired by Spring. Stars:`131`.
- [patron](https://github.com/beatlabs/patron) - Patron is a microservice framework following best cloud practices with a focus on productivity. Stars:`126`.
- [Microservice](https://github.com/claygod/microservice) - The framework for the creation of microservices, written in Golang. Stars:`118`.
- [rux](https://github.com/gookit/rux) - Simple and fast web framework for build golang HTTP applications. Stars:`97`.
- [Don](https://github.com/abemedia/go-don) - A highly performant and simple to use API framework. Stars:`54`.
- [Ronykit](https://github.com/clubpay/ronykit) - Web framework with pluggable architecture and very performant. Stars:`34`.
- [Lit](https://github.com/jvcoutinho/lit) - Highly performant declarative web framework for Golang, aiming for simplicity and quality of life. Stars:`30`.
- [golamb](https://github.com/twharmon/golamb) - Golamb makes it easier to write API endpoints for use with AWS Lambda and API Gateway. Stars:`7`.
- [Pnutmux](https://gitlab.com/fruitygo/pnutmux) - Pnutmux is a powerful Go web framework that uses regex for matching and handling HTTP requests. It offers features such as CORS handling, structured logging, URL parameters extraction, middlewares, and concurrency limiting.
- [Confetti Framework](https://confetti-framework.github.io/docs/) - Confetti is a Go web application framework with an expressive, elegant syntax. Confetti combines the elegance of Laravel and the simplicity of Go.

**[⬆ back to top](#contents)**

### Middlewares

#### Actual middlewares

- [Tollbooth](https://github.com/didip/tollbooth) - Rate limit HTTP request handler. Stars:`2.8K`.
- [CORS](https://github.com/rs/cors) - Easily add CORS capabilities to your API. Stars:`2.8K`.
- [Limiter](https://github.com/ulule/limiter) - Dead simple rate limit middleware for Go. Stars:`2.3K`.
- [go-server-timing](https://github.com/mitchellh/go-server-timing) - Add/parse Server-Timing header. Stars:`860`.
- [go-fault](https://github.com/github/go-fault) - Fault injection middleware for Go. Stars:`509`.
- [ln-paywall](https://github.com/philippgille/ln-paywall) - Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). Stars:`155`.
- [XFF](https://github.com/sebest/xff) - Handle `X-Forwarded-For` header and friends. Stars:`100`.
- [rk-grpc](https://github.com/rookie-ninja/rk-grpc) - Middleware for gRPC with logging, metrics, auth, tracing etc. Stars:`78`.
- [rk-gin](https://github.com/rookie-ninja/rk-gin) - Middleware for Gin framework with logging, metrics, auth, tracing etc. Stars:`50`.
- [formjson](https://github.com/rs/formjson) - Transparently handle JSON input as a standard form POST. Stars:`38`.
- [client-timing](https://github.com/posener/client-timing) - An HTTP client for Server-Timing header. Stars:`24`.
- [echo-middleware](https://github.com/faabiosr/echo-middleware) - Middleware for Echo framework with logging and metrics. Stars:`16`.
- [mid](https://github.com/bobg/mid) - Miscellaneous HTTP middleware features: idiomatic error return from handlers; receive/respond with JSON data; request tracing; and more. Stars:`10`.

#### Libraries for creating HTTP middlewares

- [negroni](https://github.com/urfave/negroni) - Idiomatic HTTP middleware for Golang. Stars:`7.5K`.
- [alice](https://github.com/justinas/alice) - Painless middleware chaining for Go. Stars:`3.3K`.
- [render](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, and HTML template responses. Stars:`2.0K`.
- [stats](https://github.com/thoas/stats) - Go middleware that stores various information about your web application. Stars:`593`.
- [interpose](https://github.com/carbocation/interpose) - Minimalist net/http middleware for golang. Stars:`293`.
- [renderer](https://github.com/thedevsaddam/renderer) - Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. Stars:`262`.
- [muxchain](https://github.com/stephens2424/muxchain) - Lightweight middleware for net/http. Stars:`206`.
- [gores](https://github.com/alioygur/gores) - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. Stars:`106`.
- [rye](https://github.com/InVisionApp/rye) - Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. Stars:`100`.
- [mediary](https://github.com/HereMobilityDevelopers/mediary) - add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses. Stars:`88`.
- [chain](https://github.com/codemodus/chain) - Handler wrapper chaining with scoped data (net/context-based "middleware"). Stars:`61`.
- [catena](https://github.com/codemodus/catena) - http.Handler wrapper catenation (same API as "chain"). Stars:`9`.

**[⬆ back to top](#contents)**

### Routers

- [mux](https://github.com/gorilla/mux) - Powerful URL router and dispatcher for golang. Stars:`21.5K`.
- [chi](https://github.com/go-chi/chi) - Small, fast and expressive HTTP router built on net/context. Stars:`20.2K`.
- [httprouter](https://github.com/julienschmidt/httprouter) - High performance router. Use this and the standard http handlers to form a very high performance web framework. Stars:`16.9K`.
- [gocraft/web](https://github.com/gocraft/web) - Mux and middleware package in Go. Stars:`1.5K`.
- [Bone](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer. Stars:`1.3K`.
- [Goji](https://github.com/goji/goji) - Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. Stars:`971`.
- [fasthttprouter](https://github.com/buaazp/fasthttprouter) - High performance router forked from `httprouter`. The first router fit for `fasthttp`. Stars:`871`.
- [httptreemux](https://github.com/dimfeld/httptreemux) - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. Stars:`619`.
- [xujiajun/gorouter](https://github.com/xujiajun/gorouter) - A simple and fast HTTP router for Go. Stars:`529`.
- [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. Stars:`456`.
- [lars](https://github.com/go-playground/lars) - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. Stars:`384`.
- [Siesta](https://github.com/VividCortex/siesta) - Composable framework to write middleware and handlers. Stars:`348`.
- [vestigo](https://github.com/husobee/vestigo) - Performant, stand-alone, HTTP compliant URL Router for go web applications. Stars:`267`.
- [gowww/router](https://github.com/gowww/router) - Lightning fast HTTP router fully compatible with the net/http.Handler interface. Stars:`185`.
- [GoRouter](https://github.com/vardius/gorouter) - GoRouter is a Server/API micro framework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. Stars:`153`.
- [pure](https://github.com/go-playground/pure) - Is a lightweight HTTP router that sticks to the std "net/http" implementation. Stars:`150`.
- [alien](https://github.com/gernest/alien) - Lightweight and fast http router from outer space. Stars:`135`.
- [violetear](https://github.com/nbari/violetear) - Go HTTP router. Stars:`108`.
- [Bxog](https://github.com/claygod/Bxog) - Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. Stars:`101`.
- [xmux](https://github.com/rs/xmux) - High performance muxer based on `httprouter` with `net/context` support. Stars:`99`.
- [goblin](https://github.com/bmf-san/goblin) - A golang http router based on trie tree. Stars:`81`.
- [ngamux](https://github.com/ngamux/ngamux) - Simple HTTP router for Go. Stars:`71`.
- [bellt](https://github.com/GuilhermeCaruso/bellt) - A simple Go HTTP router. Stars:`54`.
- [FastRouter](https://github.com/razonyang/fastrouter) - a fast, flexible HTTP router written in Go. Stars:`23`.
- [GoLobby/Router](https://github.com/golobby/router) - GoLobby Router is a lightweight yet powerful HTTP router for the Go programming language. Stars:`22`.
- [nchi](https://github.com/muir/nchi) - chi-like router built on httprouter with dependency injection based middleware wrappers Stars:`18`.
- [goroute](https://github.com/goroute/route) - Simple yet powerful HTTP request multiplexer. Stars:`9`.

**[⬆ back to top](#contents)**

## WebAssembly

- [tinygo](https://github.com/tinygo-org/tinygo) - Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. Stars:`16.5K`.
- [dom](https://github.com/dennwc/dom) - DOM library. Stars:`497`.
- [go-canvas](https://github.com/markfarnan/go-canvas) - Library to use HTML5 Canvas, with all drawing within go code. Stars:`264`.
- [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) - Run Go WASM tests in your browser. Stars:`200`.
- [webapi](https://github.com/gowebapi/webapi) - Bindings for DOM and HTML generated from WebIDL. Stars:`179`.
- [Extism Go SDK](https://github.com/extism/go-sdk) - Universal, cross-language WebAssembly framework for building plug-in systems and polyglot apps. Stars:`126`.
- [vert](https://github.com/norunners/vert) - Interop between Go and JS values. Stars:`104`.

**[⬆ back to top](#contents)**

## Webhooks Server

- [webhook](https://github.com/adnanh/webhook) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. Stars:`11.1K`.
- [WebhookX](https://github.com/webhookx-io/webhookx) - A webhooks gateway for message receiving, processing, and reliable delivering. Stars:`235`.
- [webhooked](https://github.com/42Atomys/webhooked) - A webhook receiver on steroids: handle, secure, format and store a Webhook payload has never been easier. Stars:`41`.

**[⬆ back to top](#contents)**

## Windows

- [go-ole](https://github.com/go-ole/go-ole) - Win32 OLE implementation for golang. Stars:`1.3K`.
- [d3d9](https://github.com/gonutz/d3d9) - Go bindings for Direct3D9. Stars:`161`.
- [gosddl](https://github.com/MonaxGT/gosddl) - Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. Stars:`11`.

**[⬆ back to top](#contents)**

## Workflow Frameworks

_Libraries for creating Workflows._

- [Dagu](https://github.com/dagu-go/dagu) - No-code workflow executor. it executes DAGs defined in a simple YAML format. Stars:`2.5K`.
- [go-taskflow](https://github.com/noneback/go-taskflow) - A taskflow-like General-purpose Task-parallel Programming Framework with integrated visualizer and profiler. Stars:`564`.
- [Cadence-client](https://github.com/uber-go/cadence-client) - A framework for authoring workflows and activities running on top of the Cadence orchestration engine made by Uber. Stars:`362`.
- [workflow](https://github.com/luno/workflow) - A tech stack agnostic Event Driven Workflow framework. Stars:`181`.
- [go-dag](https://github.com/rhosocial/go-dag) - A framework developed in Go that manages the execution of workflows described by directed acyclic graphs. Stars:`29`.

**[⬆ back to top](#contents)**

## XML

_Libraries and tools for manipulating XML._

- [zek](https://github.com/miku/zek) - Generate a Go struct from XML. Stars:`771`.
- [xpath](https://github.com/antchfx/xpath) - XPath package for Go. Stars:`724`.
- [xmlquery](https://github.com/antchfx/xmlquery) - xmlquery is Golang XPath package for XML query. Stars:`475`.
- [xml2map](https://github.com/sbabiv/xml2map) - XML to MAP converter written Golang. Stars:`64`.
- [xmlwriter](https://github.com/shabbyrobe/xmlwriter) - Procedural XML generation API based on libxml2's xmlwriter module. Stars:`28`.
- [XML-Comp](https://github.com/xml-comp/xml-comp) - Simple command line XML comparer that generates diffs of folders, files and tags. Stars:`21`.

## Zero Trust

_Libraries and tools to implement Zero Trust architectures._

- [Cosign](https://github.com/sigstore/cosign) - Container Signing, Verification and Storage in an OCI registry. Stars:`5.1K`.
- [OpenZiti](https://github.com/openziti/ziti) - A full, open source zero trust overlay network. Including numerous SDKs for numerous languages such as [golang](https://github.com/openziti/sdk-golang) allowing you to embed zero trust principles directly into your applications. The [OpenZiti Test Kitchen](https://github.com/openziti-test-kitchen) has numerous examples to draw inspiration from including a [zero trust ssh client - zssh](https://github.com/openziti-test-kitchen/zssh) Stars:`3.5K`.
- [Spire](https://github.com/spiffe/spire) - SPIRE (the SPIFFE Runtime Environment) is a toolchain of APIs for establishing trust between software systems across a wide variety of hosting platforms. Stars:`2.0K`.
- [in-toto](https://github.com/in-toto/in-toto-golang) - Go implementation of the in-toto (provides a framework to protect the integrity of the software supply chain) python reference implementation. Stars:`140`.
- [Spiffe-Vault](https://github.com/philips-labs/spiffe-vault) - Utilizes Spiffe JWT authentication with Hashicorp Vault for secretless authentication. Stars:`90`.

## Code Analysis

_Source code analysis tools, also known as Static Application Security Testing (SAST) Tools._

- [golangci-lint](https://github.com/golangci/golangci-lint) – A fast Go linters runner. It runs linters in parallel, uses caching, supports `yaml` config, has integrations with all major IDE and has dozens of linters included. Stars:`17.4K`.
- [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#. Stars:`6.5K`.
- [revive](https://github.com/mgechev/revive) – ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for `golint`. Stars:`5.3K`.
- [errcheck](https://github.com/kisielk/errcheck) - Errcheck is a program for checking for unchecked errors in Go programs. Stars:`2.4K`.
- [GoPlantUML](https://github.com/jfeliu007/goplantuml) - Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. Stars:`2.0K`.
- [go-critic](https://github.com/go-critic/go-critic) - source code linter that brings checks that are currently not implemented in other linters. Stars:`2.0K`.
- [golines](https://github.com/segmentio/golines) - Formatter that automatically shortens long lines in Go code. Stars:`1.1K`.
- [php-parser](https://github.com/z7zmey/php-parser) - A Parser for PHP written in Go. Stars:`955`.
- [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) - go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. Stars:`954`.
- [vacuum](https://github.com/daveshanley/vacuum) - An ultra-super-fast, lightweight OpenAPI linter and quality checking tool. Stars:`811`.
- [goast-viewer](https://github.com/yuroyoro/goast-viewer) - Web based Golang AST visualizer. Stars:`784`.
- [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) - An easy way to find outdated dependencies of your Go projects. Stars:`662`.
- [goreturns](https://github.com/sqs/goreturns) - Adds zero-value return statements to match the func return types. Stars:`534`.
- [Chronos](https://github.com/amit-davidson/Chronos) - Detects race conditions statically Stars:`435`.
- [todocheck](https://github.com/preslavmihaylov/todocheck) - Static code analyser which links TODO comments in code with issues in your issue tracker. Stars:`434`.
- [unconvert](https://github.com/mdempsky/unconvert) - Remove unnecessary type conversions from Go source. Stars:`384`.
- [dupl](https://github.com/mibk/dupl) - Tool for code clone detection. Stars:`352`.
- [wrapcheck](https://github.com/tomarrell/wrapcheck) - A linter to check that errors from external packages are wrapped. Stars:`344`.
- [tickgit](https://github.com/augmentable-dev/tickgit) - CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author. Stars:`324`.
- [gostatus](https://github.com/shurcooL/gostatus) - Command line tool, shows the status of repositories that contain Go packages. Stars:`244`.
- [apicompat](https://github.com/bradleyfalzon/apicompat) - Checks recent changes to a Go project for backwards incompatible changes. Stars:`180`.
- [testifylint](https://github.com/Antonboom/testifylint) – A linter that checks usage of [github.com/stretchr/testify](https://github.com/stretchr/testify). Stars:`147`.
- [go-checkstyle](https://github.com/qiniu/checkstyle) - checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. Stars:`131`.
- [asty](https://github.com/asty-org/asty) - Converts golang AST to JSON and JSON to AST. Stars:`88`.
- [lint](https://github.com/surullabs/lint) - Run linters as part of go test. Stars:`65`.
- [validate](https://github.com/mccoyst/validate) - Automatically validates struct fields with tags. Stars:`62`.
- [ChainJacking](https://github.com/Checkmarx/chainjacking) - Find which of your Go lang direct GitHub dependencies is susceptible to ChainJacking attack. Stars:`58`.
- [fatcontext](https://github.com/Crocmagnon/fatcontext) - Fatcontext detects nested contexts in loops or function literals. Stars:`51`.
- [usestdlibvars](https://github.com/sashamelentyev/usestdlibvars) - A linter that detect the possibility to use variables/constants from the Go standard library. Stars:`45`.
- [golang-ifood-sdk](https://github.com/arxdsilva/golang-ifood-sdk) - iFood API SDK. Stars:`13`.
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
- [blanket](https://gitlab.com/verygoodsoftwarenotvirus/blanket) - blanket is a tool that helps you catch functions which don't have direct unit tests in your Go packages.

**[⬆ back to top](#contents)**

## Editor Plugins

_Plugin for text editors and IDEs._

- [vim-go](https://github.com/fatih/vim-go) - Go development plugin for Vim. Stars:`16.2K`.
- [gocode](https://github.com/nsf/gocode) - Autocompletion daemon for the Go programming language. Stars:`5.0K`.
- [vscode-go](https://github.com/golang/vscode-go) - Extension for Visual Studio Code (VS Code) which provides support for the Go language. Stars:`4.1K`.
- [GoSublime](https://github.com/DisposaBoy/GoSublime) - Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. Stars:`3.4K`.
- [go-mode](https://github.com/dominikh/go-mode.el) - Go mode for GNU/Emacs. Stars:`1.4K`.
- [goimports-reviser](https://github.com/incu6us/goimports-reviser) - Formatting tool for imports. Stars:`679`.
- [coc-go language server extension for Vim/Neovim](https://github.com/josa42/coc-go) - This plugin adds [gopls](https://github.com/golang/tools/blob/master/gopls/README.md) features to Vim/Neovim. Stars:`576`.
- [Watch](https://github.com/eaburns/Watch) - Runs a command in an acme win on file changes. Stars:`202`.
- [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) - Vim plugin to highlight syntax errors on save. Stars:`90`.
- [gounit-vim](https://github.com/hexdigest/gounit-vim) - Vim plugin for generating Go tests based on the function's or method's signature. Stars:`25`.
- [Go Doc](https://github.com/msyrus/vscode-go-doc) - A Visual Studio Code extension for showing definition in output and generating go doc. Stars:`8`.
- [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof) - This extension adds benchmark profiling support for the Go language to VS Code.
- [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) - Go plugin for JetBrains IDEs.

**[⬆ back to top](#contents)**

## Go Generate Tools

- [gotests](https://github.com/cweill/gotests) - Generate Go tests from your source code. Stars:`5.1K`.
- [xgen](https://github.com/xuri/xgen) - XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator. Stars:`377`.
- [hasgo](https://github.com/DylanMeeus/hasgo) - Generate Haskell inspired functions for your slices. Stars:`143`.
- [gocontracts](https://github.com/Parquery/gocontracts) - brings design-by-contract to Go by synchronizing the code with the documentation. Stars:`114`.
- [gonerics](https://github.com/bouk/gonerics) - Idiomatic Generics in Go. Stars:`113`.
- [options-gen](https://github.com/kazhuravlev/options-gen) - Functional options described by Dave Cheney's post "Functional options for friendly APIs". Stars:`96`.
- [envdoc](https://github.com/g4s8/envdoc) -  generate documentation for environment variables from Go source files. Stars:`88`.
- [sqlgen](https://github.com/anqiansong/sqlgen) - Generate gorm, xorm, sqlx, bun, sql code from SQL file or DSN. Stars:`85`.
- [gounit](https://github.com/hexdigest/gounit) - Generate Go tests using your own templates. Stars:`85`.
- [generic](https://github.com/usk81/generic) - flexible data type for Go. Stars:`49`.
- [godal](https://github.com/mafulong/godal) - Generate orm models corresponding to golang by specifying sql ddl file, which can be used by gorm. Stars:`18`.
- [TOML-to-Go](https://xuri.me/toml-to-go) - Translates TOML into a Go type in the browser instantly.
- [re2dfa](https://gitlab.com/opennota/re2dfa) - Transform regular expressions into finite state machines and output Go source code.

**[⬆ back to top](#contents)**

## Go Tools

- [go-swagger](https://github.com/go-swagger/go-swagger) - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. Stars:`9.8K`.
- [go-callvis](https://github.com/TrueFurby/go-callvis) - Visualize call graph of your Go program using dot format. Stars:`6.3K`.
- [OctoLinker](https://github.com/OctoLinker/browser-extension) - Navigate through go files efficiently with the OctoLinker browser extension for GitHub. Stars:`5.3K`.
- [go-size-analyzer](https://github.com/Zxilly/go-size-analyzer) - Analyze and visualize the size of dependencies in compiled Golang binaries, providing insight into their impact on the final build. Stars:`1.6K`.
- [richgo](https://github.com/kyoh86/richgo) - Enrich `go test` outputs with text decorations. Stars:`862`.
- [MoniGO](https://github.com/iyashjayesh/monigo) - A performance monitoring library for Go applications. It provides real-time insights into application performance! 🚀 Stars:`275`.
- [rts](https://github.com/galeone/rts) - RTS: response to struct. Generates Go structs from server responses. Stars:`251`.
- [roumon](https://github.com/becheran/roumon) - Monitor current state of all active goroutines via a command line interface. Stars:`220`.
- [godbg](https://github.com/tylerwince/godbg) - Implementation of Rusts `dbg!` macro for quick and easy debugging during development. Stars:`206`.
- [typex](https://github.com/dtgorski/typex) - Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration. Stars:`205`.
- [gotestdox](https://github.com/bitfield/gotestdox) - Show Go test results as readable sentences. Stars:`146`.
- [gothanks](https://github.com/psampaz/gothanks) - GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers. Stars:`128`.
- [igo](https://github.com/rocketlaunchr/igo) - An igo to go transpiler (new language features for Go language!) Stars:`68`.
- [gotutor](https://github.com/ahmedakef/gotutor) - Online Go Debugger & Visualizer. Stars:`53`.
- [docs](https://github.com/go-oas/docs) - Automatically generate RESTful API documentation for GO projects - aligned with Open API Specification standard. Stars:`46`.
- [gomodrun](https://github.com/dustinblackman/gomodrun/) - Go tool that executes and caches binaries included in go.mod files. Stars:`38`.
- [decouple](https://github.com/bobg/decouple) - Find “overspecified” function parameters that could be generalized with interface types. Stars:`31`.
- [modver](https://github.com/bobg/modver) - Compare two versions of a Go module to check the version-number change required (major, minor, or patchlevel), according to [semver](https://semver.org/) rules. Stars:`21`.
- [textra](https://github.com/ravsii/textra) - Extract Go struct field names, types and tags for filtering and exporting. Stars:`7`.
- [gotemplate.io](https://gotemplate.io/) - Online tool to preview `text/template` templates live.
- [go-template-playground](https://bartventer.github.io/go-template-playground/) - An interactive environment to create and test Go templates.

**[⬆ back to top](#contents)**

## Software Packages

_Software written in Go._

**[⬆ back to top](#contents)**

### DevOps Tools

- [kubernetes](https://github.com/kubernetes/kubernetes) - Container Cluster Manager from Google. Stars:`116.6K`.
- [Moby](https://github.com/moby/moby) - Collaborative project for the container ecosystem to assemble container-based systems. Stars:`70.4K`.
- [traefik](https://github.com/containous/traefik) - Reverse proxy and load balancer with support for multiple backends. Stars:`55.8K`.
- [Gitea](https://github.com/go-gitea/gitea) - Fork of Gogs, entirely community driven. Stars:`49.8K`.
- [minikube](https://github.com/kubernetes/minikube) - Run Kubernetes locally. Stars:`30.7K`.
- [k3s](https://github.com/k3s-io/k3s) - Lightweight Kubernetes. Stars:`30.3K`.
- [k6](https://github.com/grafana/k6) - A modern load testing tool, using Go and JavaScript. Stars:`28.4K`.
- [Vegeta](https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000! Stars:`24.4K`.
- [Hey](https://github.com/rakyll/hey) - Hey is a tiny program that sends some load to a web application. Stars:`19.1K`.
- [Packer](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. Stars:`15.4K`.
- [kind](https://github.com/kubernetes-sigs/kind) - Kubernetes IN Docker - local clusters for testing Kubernetes. Stars:`14.4K`.
- [kubeshark](https://github.com/kubeshark/kubeshark) - API traffic analyzer for Kubernetes, inspired by Wireshark, purposely built for Kubernetes. Stars:`11.5K`.
- [GVM](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions. Stars:`11.1K`.
- [Flannel](https://github.com/flannel-io/flannel) - Flannel is a network fabric for containers, designed for Kubernetes. Stars:`9.2K`.
- [Ddosify](https://github.com/ddosify/ddosify) - High-performance load testing tool, written in Golang. Stars:`8.5K`.
- [ko](https://github.com/google/ko) - Command line tool for building and deploying Go applications on Kubernetes Stars:`8.1K`.
- [script](https://github.com/bitfield/script) - Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. Stars:`6.8K`.
- [KubeVela](https://github.com/kubevela/kubevela) - Cloud native application delivery. Stars:`6.8K`.
- [bombardier](https://github.com/codesenberg/bombardier) - Fast cross-platform HTTP benchmarking tool. Stars:`6.4K`.
- [k3d](https://github.com/k3d-io/k3d) - Little helper to run CNCF's k3s in Docker. Stars:`5.9K`.
- [podinfo](https://github.com/stefanprodan/podinfo) - Podinfo is a tiny web application made with Go that showcases best practices of running microservices in Kubernetes. Podinfo is used by CNCF projects like Flux and Flagger for end-to-end testing and workshops. Stars:`5.7K`.
- [Fleet device management](https://github.com/fleetdm/fleet) - Lightweight, programmable telemetry for servers and workstations. Stars:`5.3K`.
- [gaia](https://github.com/gaia-pipeline/gaia) - Build powerful pipelines in any programming language. Stars:`5.2K`.
- [Pomerium](https://github.com/pomerium/pomerium) - Pomerium is an identity-aware access proxy. Stars:`4.4K`.
- [tau](https://github.com/taubyte/tau) - Easily build Cloud Computing Platforms with features like Serverless WebAssembly Functions, Frontend Hosting, CI/CD, Object Storage, K/V Database, and Pub-Sub Messaging. Stars:`4.2K`.
- [s5cmd](https://github.com/peak/s5cmd) - Blazing fast S3 and local filesystem execution tool. Stars:`3.5K`.
- [bosun](https://github.com/bosun-monitor/bosun) - Time Series Alerting Framework. Stars:`3.4K`.
- [kubeblocks](https://github.com/apecloud/kubeblocks) - KubeBlocks is an open-source control plane that runs and manages databases, message queues and other data infrastructure on K8s. Stars:`2.8K`.
- [aptly](https://github.com/aptly-dev/aptly) - aptly is a Debian repository management tool. Stars:`2.7K`.
- [kala](https://github.com/ajvb/kala) - Simplistic, modern, and performant job scheduler. Stars:`2.2K`.
- [fac](https://github.com/mkchoi212/fac) - Command-line user interface to fix git merge conflicts. Stars:`1.9K`.
- [ghorg](https://github.com/gabrie30/ghorg) - Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Gitea, and Bitbucket. Stars:`1.8K`.
- [StatusOK](https://github.com/sanathp/statusok) - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. Stars:`1.6K`.
- [go-selfupdate](https://github.com/sanbornm/go-selfupdate) - Enable your Go applications to self update. Stars:`1.6K`.
- [tlm](https://github.com/yusufcanb/tlm) - Local cli copilot, powered by CodeLLaMa Stars:`1.4K`.
- [uTask](https://github.com/ovh/utask) - Automation engine that models and executes business processes declared in yaml. Stars:`1.3K`.
- [PipeCD](https://github.com/pipe-cd/pipecd) - A GitOps-style continuous delivery platform that provides consistent deployment and operations experience for any applications. Stars:`1.2K`.
- [KubeVPN](https://github.com/kubenetworks/kubevpn) - KubeVPN offers a Cloud-Native Dev Environment that seamlessly connects to your Kubernetes cluster network. Stars:`1.2K`.
- [KusionStack](https://github.com/KusionStack/kusion) - A unified programmable configuration techstack to deliver modern app in 'platform as code' and 'infra as code' approach. Stars:`1.2K`.
- [s3gof3r](https://github.com/rlmcpherson/s3gof3r) - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. Stars:`1.1K`.
- [kwatch](https://github.com/abahmed/kwatch) - Monitor & detect crashes in your Kubernetes(K8s) cluster instantly. Stars:`986`.
- [skm](https://github.com/TimothyYe/skm) - SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! Stars:`980`.
- [Scaleway-cli](https://github.com/scaleway/scaleway-cli) - Manage BareMetal Servers from Command Line (as easily as with Docker). Stars:`926`.
- [cassowary](https://github.com/rogerwelin/cassowary) - Modern cross-platform HTTP load-testing tool written in Go. Stars:`804`.
- [alaz](https://github.com/ddosify/alaz) - Effortless, Low-Overhead, eBPF-based Kubernetes Monitoring. Stars:`698`.
- [kool](https://github.com/kool-dev/kool) - Command line tool for managing Docker environments as an easy way. Stars:`691`.
- [aurora](https://github.com/xuri/aurora) - Cross-platform web-based Beanstalkd queue server console. Stars:`601`.
- [govvv](https://github.com/ahmetalpbalkan/govvv) - “go build” wrapper to easily add version information into Go binaries. Stars:`539`.
- [Pewpew](https://github.com/bengadbois/pewpew) - Flexible HTTP command line stress tester. Stars:`437`.
- [jcli](https://github.com/jenkins-zh/jenkins-cli) - Jenkins CLI allows you manage your Jenkins as an easy way. Stars:`405`.
- [gobrew](https://github.com/kevincobain2000/gobrew) - Go version manager. Super simple tool to install and manage Go versions. Install go without root. Gobrew doesn't require shell rehash. Stars:`397`.
- [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) - S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth). Stars:`377`.
- [gonative](https://github.com/inconshreveable/gonative) - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. Stars:`338`.
- [lstags](https://github.com/ivanilves/lstags) - Tool and API to sync Docker images across different registries. Stars:`337`.
- [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) - Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. Stars:`336`.
- [trubka](https://github.com/xitonix/trubka) - A CLI tool to manage and troubleshoot Apache Kafka clusters with the ability of generically publishing/consuming protocol buffer and plain text events to/from Kafka. Stars:`334`.
- [Mora](https://github.com/emicklei/mora) - REST server for accessing MongoDB documents and meta data. Stars:`316`.
- [Balerter](https://github.com/balerter/balerter) - A self-hosted script-based alerting manager. Stars:`306`.
- [manssh](https://github.com/xwjdsh/manssh) - manssh is a command line tool for managing your ssh alias config easily. Stars:`300`.
- [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) - Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. Stars:`278`.
- [dogo](https://github.com/liudng/dogo) - Monitoring changes in the source file and automatically compile and run (restart). Stars:`272`.
- [dish](https://github.com/thevxn/dish) - A lightweight, remotely configurable monitoring service. Stars:`247`.
- [godbg](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application. Stars:`228`.
- [abbreviate](https://github.com/dnnrly/abbreviate) - abbreviate is a tool turning long strings in to shorter ones with configurable separators, for example to embed branch names in to deployment stack IDs. Stars:`222`.
- [Blast](https://github.com/dave/blast) - A simple tool for API load testing and batch jobs. Stars:`220`.
- [kcli](https://github.com/cswank/kcli) - Command line tool for inspecting kafka topics/partitions/messages. Stars:`220`.
- [gobrew](https://github.com/cryptojuice/gobrew) - gobrew lets you easily switch between multiple versions of go. Stars:`191`.
- [ostent](https://github.com/ostrost/ostent) - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. Stars:`180`.
- [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) - A go library and an executable that produces valid Dockerfiles using various input channels. Stars:`177`.
- [grapes](https://github.com/yaronsumel/grapes) - Lightweight tool designed to distribute commands over ssh with ease. Stars:`169`.
- [winrm-cli](https://github.com/masterzen/winrm-cli) - Cli tool to remotely execute commands on Windows machines. Stars:`168`.
- [tf-profile](https://github.com/datarootsio/tf-profile) - Profiler for Terraform runs. Generate global stats, resource-level stats or visualizations. Stars:`158`.
- [drone-scp](https://github.com/appleboy/drone-scp) - Copy files and artifacts via SSH using a binary, docker or Drone CI. Stars:`156`.
- [Mantil](https://github.com/mantil-io/mantil) - Go specific framework for building serverless applications on AWS that enables you to focus on pure Go code while Mantil takes care of the infrastructure. Stars:`114`.
- [go-rocket-update](https://github.com/mouuff/go-rocket-update) - A simple way to make self updating Go applications - Supports Github and Gitlab. Stars:`107`.
- [decompose](https://github.com/s0rg/decompose) - tool to generate and process Docker containers connections graphs. Stars:`107`.
- [go-furnace](https://github.com/go-furnace/go-furnace) - Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. Stars:`99`.
- [Dropship](https://github.com/chrismckenzie/dropship) - Tool for deploying code via cdn. Stars:`66`.
- [docker-go-mingw](https://github.com/x1unix/docker-go-mingw) - Docker image for building Go binaries for Windows with MinGW toolchain. Stars:`52`.
- [httpref](https://github.com/dnnrly/httpref) - httpref is a handy CLI reference for HTTP methods, status codes, headers, and TCP and UDP ports. Stars:`41`.
- [drone-jenkins](https://github.com/appleboy/drone-jenkins) - Trigger downstream Jenkins jobs using a binary, docker or Drone CI. Stars:`38`.
- [awsenv](https://github.com/soniah/awsenv) - Small binary that loads Amazon (AWS) environment variables for a profile. Stars:`35`.
- [Rodent](https://github.com/alouche/rodent) - Rodent helps you manage Go versions, projects and track dependencies. Stars:`33`.
- [lwc](https://github.com/timdp/lwc) - A live-updating version of the UNIX wc command. Stars:`32`.
- [DepCharge](https://github.com/centerorbit/depcharge) - Helps orchestrating the execution of commands across the many dependencies in larger projects. Stars:`23`.
- [wait-for](https://github.com/dnnrly/wait-for) - Wait for something to happen (from the command line) before continuing. Easy orchestration of Docker services and other things. Stars:`17`.
- [sg](https://github.com/ChristopherRabotin/sg) - Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. Stars:`8`.
- [Docker](https://www.docker.com/) - Open platform for distributed applications for developers and sysadmins.
- [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator) - Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
- [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
- [Wide](https://wide.b3log.org/login) - Web-based IDE for Teams using Golang.

**[⬆ back to top](#contents)**

### Other Software

- [croc](https://github.com/schollz/croc) - Easily and securely send files or folders from one computer to another. Stars:`30.7K`.
- [restic](https://github.com/restic/restic) - De-duplicating backup program. Stars:`29.5K`.
- [Seaweed File System](https://github.com/chrislusf/seaweedfs) - Fast, Simple and Scalable Distributed File System with O(1) disk seek. Stars:`25.4K`.
- [Gor](https://github.com/buger/gor) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. Stars:`19.0K`.
- [JuiceFS](https://github.com/juicedata/juicefs) - Distributed POSIX file system built on top of Redis and AWS S3. Stars:`12.0K`.
- [toxiproxy](https://github.com/shopify/toxiproxy) - Proxy to simulate network and system conditions for automated tests. Stars:`11.4K`.
- [Comcast](https://github.com/tylertreat/Comcast) - Simulate bad network connections. Stars:`10.4K`.
- [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul. Stars:`8.4K`.
- [LiteIDE](https://github.com/visualfc/liteide) - LiteIDE is a simple, open source, cross-platform Go IDE. Stars:`7.7K`.
- [scc](https://github.com/boyter/scc) - Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. Stars:`7.5K`.
- [drive](https://github.com/odeke-em/drive) - Google Drive client for the commandline. Stars:`6.7K`.
- [nes](https://github.com/fogleman/nes) - Nintendo Entertainment System (NES) emulator written in Go. Stars:`5.6K`.
- [blocky](https://github.com/0xERR0R/blocky) - Fast and lightweight DNS proxy as ad-blocker for local network with many features. Stars:`5.5K`.
- [Duplicacy](https://github.com/gilbertchen/duplicacy) - A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. Stars:`5.5K`.
- [myLG](https://github.com/mehrdadrad/mylg) - Command Line Network Diagnostic tool written in Go. Stars:`2.7K`.
- [GoBoy](https://github.com/Humpheh/goboy) - Nintendo Game Boy Color emulator written in Go. Stars:`2.6K`.
- [Stack Up](https://github.com/pressly/sup) - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. Stars:`2.5K`.
- [lgo](https://github.com/yunabe/lgo) - Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. Stars:`2.5K`.
- [Documize](https://github.com/documize/community) - Modern wiki software that integrates data from SaaS tools. Stars:`2.3K`.
- [Gokapi](https://github.com/Forceu/gokapi) - Lightweight server to share files, which expire after a set amount of downloads or days. Similar to Firefox Send, but without public upload. Stars:`2.2K`.
- [sonic](https://github.com/go-sonic/sonic) - Sonic is a Go Blogging Platform. Simple and Powerful. Stars:`2.1K`.
- [Circuit](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. Stars:`2.0K`.
- [portal](https://github.com/SpatiumPortae/portal) - Portal is a quick and easy command-line file transfer utility from any computer to another. Stars:`1.7K`.
- [borg](https://github.com/crufter/borg) - Terminal based search engine for bash snippets. Stars:`1.6K`.
- [Plik](https://github.com/root-gg/plik) - Plik is a temporary file upload system (Wetransfer like) in Go. Stars:`1.6K`.
- [shell2http](https://github.com/msoap/shell2http) - Executing shell commands via http server (for prototyping or remote control). Stars:`1.4K`.
- [vFlow](https://github.com/VerizonDigital/vflow) - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. Stars:`1.1K`.
- [peg](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. Stars:`1.1K`.
- [Go Package Store](https://github.com/shurcooL/Go-Package-Store) - App that displays updates for the Go packages in your GOPATH. Stars:`897`.
- [GoNB](https://github.com/janpfeifer/gonb) - Interactive Go programming with Jupyter Notebooks (also works in VSCode, Binder and Google's Colab). Stars:`894`.
- [yai](https://github.com/ekkinox/yai) - AI powered terminal assistant. Stars:`815`.
- [Leaps](https://github.com/jeffail/leaps) - Pair programming service using Operational Transforms. Stars:`756`.
- [gfile](https://github.com/Antonito/gfile) - Securely transfer files between two computers, without any third party, over WebRTC. Stars:`752`.
- [sake](https://github.com/alajmo/sake) - sake is a command runner for local and remote hosts. Stars:`703`.
- [Guora](https://github.com/meloalright/guora) - A self-hosted Quora like web application written in Go. Stars:`675`.
- [gocc](https://github.com/goccmack/gocc) - Gocc is a compiler kit for Go written in Go. Stars:`644`.
- [Gebug](https://github.com/moshebe/gebug) - A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly. Stars:`632`.
- [Chapar](https://github.com/chapar-rest/chapar) - Chapar is a a cross-platform Postman alternative built with go, aims to help developers to test their api endpoints. it support http and grpc protocols. Stars:`601`.
- [mockingjay](https://github.com/quii/mockingjay-server) - Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. Stars:`564`.
- [woke](https://github.com/get-woke/woke) - Detect non-inclusive language in your source code. Stars:`490`.
- [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) - Video streaming torrent client. Stars:`473`.
- [hotswap](https://github.com/edwingeng/hotswap) - A complete solution to reload your go code without restarting your server, interrupting or blocking any ongoing procedure. Stars:`398`.
- [ide](https://github.com/thestrukture/ide) - Browser accessible IDE. Designed for Go with Go. Stars:`361`.
- [Neo-cowsay](https://github.com/Code-Hex/Neo-cowsay) - 🐮 cowsay is reborn. for a New Era. Stars:`331`.
- [crawley](https://github.com/s0rg/crawley) - Web scraper/crawler for cli. Stars:`303`.
- [wellington](https://github.com/wellington/wellington) - Sass project management tool, extends the language with sprite functions (like Compass). Stars:`302`.
- [Cherry](https://github.com/rafael-santiago/cherry) - Tiny webchat server in Go. Stars:`301`.
- [stew](https://github.com/marwanhawari/stew) - An independent package manager for compiled binaries. Stars:`278`.
- [tcpdog](https://github.com/mehrdadrad/tcpdog) - eBPF based TCP observability. Stars:`260`.
- [joincap](https://github.com/assafmo/joincap) - Command-line utility for merging multiple pcap files together. Stars:`217`.
- [Orbit](https://github.com/gulien/orbit) - A simple tool for running commands and generating files from templates. Stars:`187`.
- [fjira](https://github.com/mk-5/fjira) - A fuzzy-search based terminal UI application for Attlasian Jira Stars:`184`.
- [vaku](https://github.com/lingrino/vaku) - CLI & API for folder-based functions in Vault like copy, move, and search. Stars:`155`.
- [dp](https://github.com/scryinfo/dp) - Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. Stars:`80`.
- [boxed](https://github.com/tejo/boxed) - Dropbox based blog engine. Stars:`78`.
- [naclpipe](https://github.com/unix4fun/naclpipe) - Simple NaCL EC25519 based crypto pipe tool written in Go. Stars:`23`.
- [Snitch](https://github.com/lucasgomide/snitch) - Simple way to notify your team and many tools when someone has deployed any application via Tsuru. Stars:`16`.
- [GoDocTooltip](https://github.com/diankong/GoDocTooltip) - Chrome extension for Go Doc sites, which shows function description as tooltip at function list. Stars:`13`.
- [tinycare-tui](https://github.com/DMcP89/tinycare-tui) - Small terminal app that shows git commits from the last 24 hours and week, current weather, some self care advice, a joke, and you current todo list tasks. Stars:`13`.
- [hoofli](https://github.com/dnnrly/hoofli) - Generate PlantUML diagrams from Chrome or Firefox network inspections. Stars:`9`.
- [syncthing](https://syncthing.net/) - Open, decentralized file synchronization tool and protocol.
- [hugo](https://gohugo.io/) - Fast and Modern Static Website Engine.
- [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
- [tsuru](https://tsuru.io/) - Extensible and open source Platform as a Service software.
- [GoLand](https://jetbrains.com/go) - Full featured cross-platform Go IDE.
- [goblin](https://goblin.run) - Cloud builder for CLI's written in go lang
- [Wave Terminal](https://waveterm.dev) - Wave is an open-source, AI-native terminal built for seamless developer workflows with inline rendering, a modern UI, and persistent sessions.
- [limetext](https://limetext.github.io) - Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
- [Better Go Playground](https://goplay.tools) - Go playground with syntax highlight, code completion and other features.
- [Layli](https://layli.app) - Draw pretty layout diagrams as code.
- [zs](https://git.mills.io/prologic/zs) - an extremely minimal static site generator.

**[⬆ back to top](#contents)**

# Resources

_Where to discover new Go libraries._

**[⬆ back to top](#contents)**

## Benchmarks

- [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) - Go web framework benchmark. Stars:`2.1K`.
- [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router benchmark and comparison. Stars:`1.7K`.
- [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods. Stars:`1.6K`.
- [skynet](https://github.com/atemerev/skynet) - Skynet 1M threads microbenchmark. Stars:`1.1K`.
- [speedtest-resize](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language. Stars:`243`.
- [go-benchmarks](https://github.com/tylertreat/go-benchmarks) - Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. Stars:`148`.
- [golang-benchmarks](https://github.com/SimonWaldherr/golang-benchmarks) - a collection of golang benchmarks. Stars:`137`.
- [gospeed](https://github.com/feyeleanor/GoSpeed) - Go micro-benchmarks for calculating the speed of language constructs. Stars:`126`.
- [autobench](https://github.com/davecheney/autobench) - Framework to compare the performance between different Go versions. Stars:`99`.
- [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) - Benchmarks of common basic operations for the Go language. Stars:`61`.
- [go-ml-benchmarks](https://github.com/nikolaydubina/go-ml-benchmarks) - benchmarks for machine learning inference in Go. Stars:`32`.
- [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) - Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. Stars:`28`.
- [kvbench](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark. Stars:`26`.
- [go-json-benchmark](https://github.com/zerosnake0/go-json-benchmark) - Go JSON benchmark. Stars:`10`.
- [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) - Collection of benchmarks for popular Go database/SQL utilities.

**[⬆ back to top](#contents)**

## Conferences

- [GoCon](https://gocon.connpass.com/) - Tokyo, Japan.
- [GoDays](https://www.godays.io/) - Berlin, Germany.
- [GoLab](https://golab.io/) - Florence, Italy.
- [GopherChina](https://gopherchina.org) - Shanghai, China.
- [GopherCon](https://www.gophercon.com/) - Varied Locations Each Year, USA.
- [GopherCon Australia](https://gophercon.com.au/) - Sydney, Australia.
- [GopherCon Brazil](https://gopherconbr.org) - Florianópolis, Brazil.
- [GopherCon Europe](https://gophercon.eu/) - Berlin, Germany.
- [GopherCon India](https://gopherconindia.org/) - Pune, India.
- [GopherCon Israel](https://www.gophercon.org.il/) - Tel Aviv, Israel.
- [GopherCon Russia](https://www.gophercon-russia.ru) - Moscow, Russia.
- [GopherCon Singapore](https://gophercon.sg) - Mapletree Business City, Singapore.
- [GopherCon UK](https://www.gophercon.co.uk/) - London, UK.
- [GopherCon Vietnam](https://gophercon.vn/) - Ho Chi Minh City, Vietnam.
- [GoWest Conference](https://www.gowestconf.com/) - Lehi, USA.

**[⬆ back to top](#contents)**

## E-Books

### E-books for purchase

- [100 Go Mistakes: How to Avoid Them](https://www.manning.com/books/100-go-mistakes-how-to-avoid-them)
- [Black Hat Go](https://nostarch.com/blackhatgo) - Go programming for hackers and pentesters.
- [Build an Orchestrator in Go](https://www.manning.com/books/build-an-orchestrator-in-go)
- [Continuous Delivery in Go](https://www.manning.com/books/continuous-delivery-in-go) - This practical guide to continuous delivery shows you how to rapidly establish an automated pipeline that will improve your testing, code quality, and final product.
- [Creative DIY Microcontroller Project With TinyGo and WebAssembly](https://www.packtpub.com/product/creative-diy-microcontroller-projects-with-tinygo-and-webassembly/9781800560208) - An introduction into the TinyGo compiler with projects involving Arduino and WebAssembly.
- [Effective Go: Elegant, efficient, and testable code](https://www.manning.com/books/effective-go) - Unlock Go’s unique perspective on program design, and start writing simple, maintainable, and testable Go code.
- [For the Love of Go](https://bitfieldconsulting.com/books/love) - An introductory book for Go beginners.
- [Go in Practice, Second Edition](https://www.manning.com/books/go-in-practice-second-edition)  - Your practical guide on the ins-and-outs of Go development, covering the standard library and the most important tools from Go’s powerful ecosystem.
- [Know Go: Generics](https://bitfieldconsulting.com/books/generics) - A guide to understanding and using generics in Go.
- [Lets-Go](https://lets-go.alexedwards.net) - A step-by-step guide to creating fast, secure and maintanable web applications with Go.
- [Lets-Go-Further](https://lets-go-further.alexedwards.net) - Advanced patterns for building APIs and web applications in Go.
- [The Power of Go: Tests](https://bitfieldconsulting.com/books/tests) - A guide to testing in Go.
- [The Power of Go: Tools](https://bitfieldconsulting.com/books/tools) - A guide to writing command-line tools in Go.
- [Writing A Compiler In Go](https://compilerbook.com)
- [Writing An Interpreter In Go](https://interpreterbook.com) - Book that introduces dozens of techniques for writing idiomatic, expressive, and efficient Go code that avoids common pitfalls.

### Free e-books

- [GoBooks](https://github.com/dariubs/GoBooks) - A curated list of Go books. Stars:`18.3K`.
- [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) Stars:`9.5K`.
- [Go AST Book (Chinese)](https://github.com/chai2010/go-ast-book) - A book focusing on Go `go/*` packages. Stars:`5.5K`.
- [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/) Stars:`3.2K`.
- [Build a blockchain from scratch in Go with gRPC](https://github.com/volodymyrprokopyuk/go-blockchain) - The foundational and practical guide for effectively learning and progressively building a blockchain from scratch in Go with gRPC. Stars:`446`.
- [Go Succinctly](https://github.com/thedevsir/gosuccinctly) - in Persian. Stars:`22`.
- [Building Web Apps With Go](https://codegangsta.gitbooks.io/building-web-apps-with-go/content/)
- [Go Faster](https://leanpub.com/gofaster) - This book seeks to shorten your learning curve and help you become a proficient Go programmer, faster.
- [Go 101](https://go101.org) - A book focusing on Go syntax/semantics and all kinds of details.
- [Go with the domain](https://threedots.tech/go-with-the-domain/) - A book showing how to apply DDD, Clean Architecture, and CQRS by practical refactoring.
- [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
- [How To Code in Go eBook](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook) - A 600 page introduction to Go aimed at first time developers.
- [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
- [Network Programming With Go](https://jan.newmarch.name/golang/)
- [Practical Go Lessons](https://www.practical-go-lessons.com/)
- [Spaceship Go A Journey to the Standard Library](https://blasrodri.github.io/spaceship-go-gh-pages/)
- [The Go Programming Language](https://www.gopl.io/)
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/)
- [An Introduction to Programming in Go](http://www.golang-book.com/)

**[⬆ back to top](#contents)**

## Gophers

- [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) - Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. Stars:`3.8K`.
- [gophers](https://github.com/egonelbre/gophers) - Free gophers. Stars:`3.7K`.
- [gophers](https://github.com/ashleymcnamara/gophers) - Gopher artworks by Ashley McNamara. Stars:`3.0K`.
- [gopherize.me](https://github.com/matryer/gopherize.me) - Gopherize yourself. Stars:`737`.
- [gophericons](https://github.com/shalakhin/gophericons) Stars:`625`.
- [gopher-stickers](https://github.com/tenntenn/gopher-stickers) Stars:`601`.
- [gophers](https://github.com/sillecelik/go-gopher) - Gopher amigurumi toy pattern. Stars:`156`.
- [gopher-logos](https://github.com/GolangUA/gopher-logos) - adorable gopher logos. Stars:`131`.
- [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) - Go gopher Vector Data [.ai, .svg]. Stars:`73`.
- [gophers](https://github.com/rogeralsing/gophers) - random gopher graphics. Stars:`58`.
- [gophers](https://github.com/scraly/gophers) - Gophers by Aurélie Vache. Stars:`36`.

**[⬆ back to top](#contents)**

## Meetups

- [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
- [Belfast Gophers](https://www.meetup.com/Belfast-Gophers/)
- [Belgrade Golang Meetup](https://www.meetup.com/golang-serbia/)
- [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
- [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
- [Bärner Go Meetup - Berne, Switzerland](https://www.meetup.com/berner-go-meetup/)
- [Go Ireland - Dublin](https://www.meetup.com/goireland/)
- [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
- [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
- [Go Remote Meetup](https://www.meetup.com/Go-Remote-Meetup/)
- [Go Toronto](https://www.meetup.com/go-toronto/)
- [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
- [GoBandung](https://www.meetup.com/GoBandung/)
- [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
- [GoCracow - Krakow, Poland](https://www.meetup.com/GoCracow/)
- [GoJakarta](https://www.meetup.com/GoJakarta/)
- [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
- [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
- [Golang Athens](https://www.meetup.com/Athens-Gophers/)
- [Golang Baltimore, MD](https://www.meetup.com/BaltimoreGolang/)
- [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
- [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
- [Golang Boston](https://www.meetup.com/bostongo/)
- [Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)
- [Golang Cardiff, UK](https://www.meetup.com/Cardiff-Go-Meetup/)
- [Golang Copenhagen](https://www.meetup.com/Go-Cph/)
- [Golang Curitiba - Brazil](https://www.meetup.com/GolangCWB/)
- [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
- [Golang Dorset, UK](https://www.meetup.com/golang-dorset/)
- [Golang Estonia](https://www.meetup.com/Golang-Estonia/)
- [Golang Gurgaon, India](https://www.meetup.com/Gurgaon-Go-Meetup/)
- [Golang Hamburg - Germany](https://www.meetup.com/Go-User-Group-Hamburg/)
- [Golang Israel](https://www.meetup.com/Go-Israel/)
- [Golang Kathmandu](https://www.meetup.com/Golang-Kathmandu/)
- [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
- [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
- [Golang Marseille](https://www.meetup.com/fr-FR/Golang-Marseille/)
- [Golang Melbourne](https://www.meetup.com/golang-mel/)
- [Golang Milano](https://www.meetup.com/golang-milano/)
- [Golang North East](https://www.meetup.com/en-AU/Golang-North-East/)
- [Golang Paris](https://www.meetup.com/Golang-Paris/)
- [Golang Poland](https://www.meetup.com/Golang-Poland/)
- [Golang Pune](https://www.meetup.com/Golang-Pune/)
- [Golang Roma](https://www.meetup.com/golangroma/)
- [Golang Rotterdam](https://www.meetup.com/golang-rotterdam/)
- [Golang Singapore](https://www.meetup.com/golangsg/)
- [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
- [Golang Sydney, AU](https://www.meetup.com/golang-syd/)
- [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
- [Golang Taipei](https://www.meetup.com/golang-taipei-meetup/)
- [Golang Thessaloniki](https://www.meetup.com/thessaloniki-golang-meetup/)
- [Golang Torino](https://www.meetup.com/golang-torino/)
- [Golang Turkey](https://kommunity.com/goturkiye)
- [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
- [Golang Vienna, Austria](https://www.meetup.com/viennago/)
- [Golang Москва](https://www.meetup.com/Golang-Moscow/)
- [GoSF - San Francisco, CA](https://www.meetup.com/golangsf)
- [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
- [Lagos Gophers](https://www.meetup.com/GolangNigeria/)
- [Nairobi Gophers](https://www.meetup.com/nairobi-gophers/)
- [Seattle Go Programmers](https://www.meetup.com/golang/)
- [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
- [Utah Go User Group](https://www.meetup.com/utahgophers/)
- [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)
- [Zürich Gophers - Zurich, Switzerland](https://www.meetup.com/zurich-gophers/)

_Add the group of your city/country here (send **PR**)_

**[⬆ back to top](#contents)**

## Style Guides

- [CockroachDB](https://github.com/cockroachdb/cockroach/blob/master/docs/style.md) Stars:`31.1K`.
- [Uber](https://github.com/uber-go/guide/blob/master/style.md) Stars:`16.7K`.
- [Hyperledger](https://github.com/hyperledger/fabric/blob/release-1.4/docs/source/style-guides/go-style.rst) Stars:`16.2K`.
- [Magnetico](https://github.com/boramalper/magnetico/wiki/magnetico-Design-Specification) Stars:`3.1K`.
- [bahlo/go-styleguide](https://github.com/bahlo/go-styleguide) Stars:`1.5K`.
- [Trybe](https://github.com/betrybe/playbook-go/blob/main/README_EN.md) Stars:`309`.
- [GitLab](https://docs.gitlab.com/ee/development/go_guide/)
- [Google](https://google.github.io/styleguide/go/)
- [Thanos](https://thanos.io/tip/contributing/coding-style-guide.md/)

**[⬆ back to top](#contents)**

## Social Media

### Twitter

- [@GoDiscussions](https://twitter.com/GoDiscussions)
- [@golang](https://twitter.com/golang)
- [@golang_news](https://twitter.com/golang_news)
- [@golangch](https://twitter.com/golangch)
- [@golangweekly](https://twitter.com/golangweekly)

**[⬆ back to top](#contents)**

### Reddit

- [r/golang](https://www.reddit.com/r/golang/)

**[⬆ back to top](#contents)**

## Websites

- [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki. Stars:`129.2K`.
- [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) - Curated list of awesome remote jobs. A lot of them are looking for Go hackers. Stars:`39.1K`.
- [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) - List of other amazingly awesome lists. Stars:`32.7K`.
- [Awesome Golang Workshops](https://github.com/amit-davidson/awesome-golang-workshops) - A curated list of awesome golang workshops. Stars:`521`.
- [gocryforhelp](https://github.com/ninedraft/gocryforhelp) - Collection of Go projects that needs help. Good place to start your open-source way in Go. Stars:`40`.
- [awesome-go-extra](https://github.com/xwjdsh/awesome-go-extra) - Parse awesome-go README file and generate a new README file with repo info. Stars:`26`.
- [Go Blog](https://blog.golang.org) - The official Go blog.
- [Golang Nugget](https://golangnugget.com) - A weekly roundup of the best Go content, delivered to your inbox every Monday.
- [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
- [Go Code Club](https://www.youtube.com/watch?v=nvoIPQYdx9g&list=PLEcwzBXTPUE_YQR7R0BRtHBYJ0LN3Y0i3) - A group of Gophers read and discuss a different Go project every week.
- [Go Community on Hashnode](https://hashnode.com/n/go) - Community of Gophers on Hashnode.
- [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
- [Coding Mystery](https://codingmystery.com) - Solve exciting escape-room-inspired programming challenges using Go.
- [Go Proverbs](https://go-proverbs.github.io/) - Go Proverbs by Rob Pike.
- [Go Report Card](https://goreportcard.com) - A report card for your Go package.
- [go.dev](https://go.dev/) - A hub for Go developers.
- [Code with Mukesh](https://codewithmukesh.com/categories/golang) - Software Engineer and Blogs @ codewithmukesh.com.
- [Golang Developer Jobs](https://golangjob.xyz) - Developer Jobs exclusively for Golang related Roles.
- [Golang News](https://golangnews.com) - Links and news about Go programming.
- [CodinGame](https://www.codingame.com/) - Learn Go by solving interactive tasks using small games as practical examples.
- [Golang Weekly](https://discu.eu/weekly/golang/) - Each monday projects, tutorials and articles about Go.
- [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
- [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for #golang enthusiasts.
- [Gopher Community Chat](https://invite.slack.golangbridge.org) - Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
- [Gophercises](https://gophercises.com/) - Free coding exercises for budding gophers.
- [json2go](https://m-zajac.github.io/json2go) - Advanced JSON to Go struct conversion - online tool.
- [justforfunc](https://www.youtube.com/c/justforfunc) - Youtube channel dedicated to Go programming language tips and tricks, hosted by Francesc Campoy [@francesc](https://twitter.com/francesc).
- [Learn Go Programming](https://blog.learngoprogramming.com) - Learn Go concepts with illustrations.
- [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
- [pkg.go.dev](https://pkg.go.dev/) - Documentation for open source Go packages.
- [studygolang](https://studygolang.com) - The community of studygolang in China.
- [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.
- [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

**[⬆ back to top](#contents)**

### Tutorials

- [50 Shades of Go](https://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/) - Traps, Gotchas, and Common Mistakes for New Golang Devs.
- [A Comprehensive Guide to Structured Logging in Go](https://betterstack.com/community/guides/logging/logging-in-go/) - Delve deep into the world of structured logging in Go with a specific focus on recently accepted slog proposal which aims to bring high performance structured logging with levels to the standard library.
- [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo) - Building a Golang site for e-commerce (demo included).
- [A Tour of Go](https://tour.golang.org/) - Interactive tour of Go.
- [Build a Database in 1000 lines of code]( https://link.medium.com/O9YQlx89Htb) - Build a NoSQL Database From Zero in 1000 Lines of Code.
- [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) - Golang ebook intro how to build a web app with golang. Stars:`43.7K`.
- [go-patterns](https://github.com/tmrts/go-patterns) - Curated list of Go design patterns, recipes and idioms. Stars:`26.7K`.
- [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) - Learn Go with test-driven development. Stars:`23.0K`.
- [Learn Go with 1000+ Exercises](https://github.com/inancgumus/learngo) - Learn Go with thousands of examples, exercises, and quizzes. Stars:`19.4K`.
- [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) - Go's reference card. Stars:`8.7K`.
- [go-clean-template](https://github.com/evrone/go-clean-template) - Clean Architecture template for Golang services. Stars:`7.0K`.
- [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) - Examples of Golang compared to Node.js for learning. Stars:`4.7K`.
- [golang-examples](https://github.com/SimonWaldherr/golang-examples) - Many examples to learn Golang. Stars:`1.6K`.
- [Go in 7 days](https://github.com/harrytran103/7_days_of_go) - Learn everything about Go in 7 days (from a Nodejs developer). Stars:`155`.
- [Design Patterns in Go](https://github.com/shubhamzanwar/design-patterns) - Collection of programming design patterns implemented in Go. Stars:`128`.
- [Hex Monscape](https://github.com/Haraj-backend/hex-monscape) - Getting started guidelines in writing maintainable code using Hexagonal Architecture. Stars:`79`.
- [Golangbot](https://golangbot.com/learn-golang-series/) - Tutorials to get started with programming in Go.
- [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9) - How to cache slow database queries.
- [Go Tutorial](https://www.tutorialspoint.com/go/index.htm) - Learn Go programming.
- [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
- [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
- [Go By Example](https://gobyexample.com/) - Hands-on introduction to Go using annotated example programs.
- [Games With Go](https://www.youtube.com/watch?v=9D4yH7e_ea8&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x) - A video series teaching programming and game development.
- [Golang Tutorial Guide](https://www.freecodecamp.org/news/golang-tutorial-list-free-courses-learn-go-programming-language/) - A List of Free Courses to Learn the Go Programming Language.
- [CodeCrafters Golang Track](https://app.codecrafters.io/tracks/go) - Achieve mastery in advanced Go by building your own Redis, Docker, Git, and SQLite. Featuring goroutines, systems programming, file I/O, and more.
- [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30) - How to cancel MySQL queries.
- [GopherCoding](https://gophercoding.com/) - Collection of code snippets and tutorials to help tackle every day issues.
- [GopherSnippets](https://gophersnippets.com/) - Code snippets with tests and testable examples for the Go programming language.
- [Gosamples](https://gosamples.dev/) - Collection of code snippets that let you solve everyday code problems.
- [GraphQL with Go](https://hasura.io/learn/graphql/backend-stack/languages/go/) - Learn how to create a Go GraphQL server and client with code generation. Also includes creating REST endpoints.
- [Hackr.io](https://hackr.io/tutorials/learn-golang) - Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
- [Go Language Tutorial](https://www.javatpoint.com/go-tutorial) - Learn Go language Tutorial.
- [How to Benchmark: dbq vs sqlx vs GORM](https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5) - Learn how to benchmark in Go. As a case-study, we will benchmark dbq, sqlx and GORM.
- [How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker) - Learn how to use Docker for Go development and how to build production Docker images.
- [How to Implement Role-Based Access Control (RBAC) Authorization in Golang](https://www.permit.io/blog/role-based-access-control-rbac-authorization-in-golang) - A guide to implementing Role-Based Access Control (RBAC) in Golang, including code examples, covering various methods to secure app endpoints with role-based authorization.
- [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - Get started with Godog - a Behavior-driven development framework for building and testing Go applications.
- [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
- [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql) - We’ll write an API with the help of the powerful Gorilla Mux.
- [Learning Go by examples](https://dev.to/aurelievache/learning-go-by-examples-introduction-448n) - Series of articles in order to learn Golang language by concrete applications as example.
- [Microservices with Go](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_) - Dive deep into building microservices using Go, including gRPC.
- [package main](https://www.youtube.com/packagemain) - YouTube channel about Programming in Go.
- [Programming with Google Go](https://www.coursera.org/specializations/google-golang) - Coursera Specialization to learn about Go from scratch.
- [Scaling Go Applications](https://betterstack.com/community/guides/scaling-go/) - Everything about building, deploying and scaling Go applications in production.
- [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
- [Understanding Go in a visual way](https://dev.to/aurelievache/series/26234) - Learn Go visually
- [W3basic Go Tutorials](https://www.w3basic.com/golang/) - W3Basic provides an in-depth tutorial and well-organized content to learn Golang programming.
- [Your basic Go](https://yourbasic.org/golang) - Huge collection of tutorials and how to's.

**[⬆ back to top](#contents)**

### Guided Learning

- [The Go Interview Practice](https://github.com/RezaSi/go-interview-practice) - A GitHub repository offering coding challenges for Go technical interview preparation. Stars:`706`.
- [The Go Developer Roadmap](https://roadmap.sh/golang) - A visual roadmap that new Go developers can follow through to help them learn Go.
- [The Go Learning Path](https://tutorialedge.net/paths/golang/) - A guided learning path containing a mix of free and premium resources.
- [The Go Skill Tree](https://labex.io/skilltrees/go) - A structured learning path that combines both free and premium resources.

**[⬆ back to top](#contents)**
