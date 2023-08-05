# Awesome Go

<a href="https://awesome-go.com/"><img align="right" src="https://github.com/avelino/awesome-go/raw/main/tmpl/assets/logo.png" alt="awesome-go" title="awesome-go" /></a>

[![Build Status](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml/badge.svg?branch=main)](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml?query=branch%3Amain)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)
[![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](https://gophers.slack.com/messages/awesome)
[![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)
[![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/avelino/awesome-go/)

We use the _[Golang Bridge](https://github.com/gobridge/about-us/blob/master/README.md)_ community Slack for instant communication, follow the [form here to join](https://invite.slack.golangbridge.org/).

<a href="https://www.producthunt.com/posts/awesome-go?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-awesome-go" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=291535&theme=light" alt="awesome-go - Curated list awesome Go frameworks, libraries and software | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

**Sponsorships:**

_Special thanks to_

<div align="center">
<table cellpadding="5">
<tbody align="center">
<tr>
<td>
<a href="https://bit.ly/awesome-go-workos">
<img src="https://avelino.run/sponsors/workos-logo-white-bg.svg" width="200" alt="WorkOS"><br/>
<b>Your app, enterprise-ready.</b><br/>
<sub>Start selling to enterprise customers with just a few lines of code.</sub><br/>
<sup>Add Single Sign-On (and more) in minutes instead of months.</sup>
</a>
</td>
<td>
<a href="https://bit.ly/awesome-go-keygen">
<img src="https://avelino.run/sponsors/keygen-logo.png" width="200" alt="keygen"><br/>
<b>An open, source-available software licensing and distribution API.</b><br/>
<sub>Securely license and distribute Go applications with a single API.</sub><br>
<sup>Add auto updates with only a few lines of code.</sup>
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
  - [Resource Embedding](#resource-embedding)
  - [Science and Data Analysis](#science-and-data-analysis)
  - [Security](#security)
  - [Serialization](#serialization)
  - [Server Applications](#server-applications)
  - [Stream Processing](#stream-processing)
  - [Template Engines](#template-engines)
  - [Testing](#testing)
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
  - [Windows](#windows)
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
    - [Guided Learning Paths](#guided-learning)

**[⬆ back to top](#contents)**

## Audio and Music

_Libraries for manipulating audio._

- [Oto](https://github.com/hajimehoshi/oto) - A low-level library to play sound on multiple platforms. Stars:`1.3K`.
- [PortAudio](https://github.com/gordonklaus/portaudio) - Go bindings for the PortAudio audio I/O library. Stars:`608`.
- [id3v2](https://github.com/bogem/id3v2) - ID3 decoding and encoding library for Go. Stars:`293`.
- [GoAudio](https://github.com/DylanMeeus/GoAudio) - Native Go Audio Processing Library. Stars:`292`.
- [flac](https://github.com/mewkiz/flac) - Native Go FLAC encoder/decoder with support for FLAC streams. Stars:`255`.
- [malgo](https://github.com/gen2brain/malgo) - Mini audio library. Stars:`223`.
- [gaad](https://github.com/Comcast/gaad) - Native Go AAC bitstream parser. Stars:`112`.
- [minimp3](https://github.com/tosone/minimp3) - Lightweight MP3 decoder library. Stars:`107`.
- [gosamplerate](https://github.com/dh1tw/gosamplerate) - libsamplerate bindings for go. Stars:`29`.

**[⬆ back to top](#contents)**

## Authentication and OAuth

_Libraries for implementing authentication schemes._

- [casbin](https://github.com/hsluoyz/casbin) - Authorization library that supports access control models like ACL, RBAC, and ABAC. Stars:`15.4K`.
- [jwt-go](https://github.com/golang-jwt/jwt) - A full featured implementation of JSON Web Tokens (JWT). This library supports the parsing and verification as well as the generation and signing of JWTs. Stars:`5.1K`.
- [oauth2](https://github.com/golang/oauth2) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine, and App Engine support. Stars:`4.8K`.
- [goth](https://github.com/markbates/goth) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. Stars:`4.2K`.
- [keto](https://github.com/ory/keto) - Open Source (Go) implementation of "Zanzibar: Google's Consistent, Global Authorization System". Ships gRPC, REST APIs, newSQL, and an easy and granular permission language. Supports ACL, RBAC, and other access models. Stars:`4.2K`.
- [authboss](https://github.com/volatiletech/authboss) - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure it, and start building your app without having to build an authentication system each time. Stars:`3.5K`.
- [loginsrv](https://github.com/tarent/loginsrv) - JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. Stars:`1.9K`.
- [osin](https://github.com/openshift/osin) - Golang OAuth2 server library. Stars:`1.8K`.
- [gologin](https://github.com/dghubble/gologin) - chainable handlers for login with OAuth1 and OAuth2 authentication providers. Stars:`1.6K`.
- [scs](https://github.com/alexedwards/scs) - Session Manager for HTTP servers. Stars:`1.5K`.
- [gorbac](https://github.com/mikespook/gorbac) - provides a lightweight role-based access control (RBAC) implementation in Golang. Stars:`1.5K`.
- [paseto](https://github.com/o1egl/paseto) - Golang implementation of Platform-Agnostic Security Tokens (PASETO). Stars:`736`.
- [jwt](https://github.com/cristalhq/jwt) - Safe, simple, and fast JSON Web Tokens for Go. Stars:`611`.
- [permissions2](https://github.com/xyproto/permissions2) - Library for keeping track of users, login states, and permissions. Uses secure cookies and bcrypt. Stars:`486`.
- [go-guardian](https://github.com/shaj13/go-guardian) - Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token, and Certificate based authentication. Stars:`483`.
- [jwt](https://github.com/pascaldekloe/jwt) - Lightweight JSON Web Token (JWT) library. Stars:`337`.
- [jeff](https://github.com/abraithwaite/jeff) - Simple, flexible, secure, and idiomatic web session management with pluggable backends. Stars:`257`.
- [jwt-auth](https://github.com/adam-hanna/jwt-auth) - JWT middleware for Golang http servers with many configuration options. Stars:`229`.
- [go-jose](https://github.com/go-jose/go-jose) - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. Stars:`148`.
- [otpgen](https://github.com/grijul/otpgen) - Library to generate TOTP/HOTP codes. Stars:`130`.
- [sessionup](https://github.com/swithek/sessionup) - Simple, yet effective HTTP session management and identification package. Stars:`120`.
- [sjwt](https://github.com/brianvoe/sjwt) - Simple jwt generator and parser. Stars:`112`.
- [session](https://github.com/icza/session) - Go session management for web servers (including support for Google App Engine - GAE). Stars:`111`.
- [securecookie](https://github.com/chmike/securecookie) - Efficient secure cookie encoding/decoding. Stars:`75`.
- [sessions](https://github.com/adam-hanna/sessions) - Dead simple, highly performant, highly customizable sessions service for go http servers. Stars:`74`.
- [branca](https://github.com/essentialkaos/branca) - branca token [specification implementation](https://github.com/tuupola/branca-spec) for Golang 1.15+. Stars:`62`.
- [otpgo](https://github.com/jltorresm/otpgo) - Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go. Stars:`58`.
- [scope](https://github.com/SonicRoshan/scope) - Easily Manage OAuth2 Scopes In Go. Stars:`32`.
- [cookiestxt](https://github.com/mengzhuo/cookiestxt) - provides a parser of cookies.txt file format. Stars:`12`.
- [gosession](http://github.com/Kwynto/gosession) - This is quick session for net/http in GoLang. This package is perhaps the best implementation of the session mechanism, or at least it tries to become one.

**[⬆ back to top](#contents)**

## Blockchain

_Tools for building blockchains._

- [go-ethereum](https://github.com/ethereum/go-ethereum) - Official Go implementation of the Ethereum protocol. Stars:`43.2K`.
- [tendermint](https://github.com/tendermint/tendermint) - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. Stars:`5.5K`.
- [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) - A Framework for Building Public Blockchains in the Cosmos Ecosystem. Stars:`5.4K`.
- [solana-go](https://github.com/gagliardetto/solana-go) - Go library to interface with Solana JSON RPC and WebSocket interfaces. Stars:`439`.
- [gossamer](https://github.com/ChainSafe/gossamer) - A Go implementation of the Polkadot Host. Stars:`391`.
- [cometbft](https://github.com/cometbft/cometbft) - A distributed, Byzantine fault-tolerant, deterministic state machine replication engine. It is a fork of Tendermint Core and implements the Tendermint consensus algorithm. Stars:`327`.

**[⬆ back to top](#contents)**

## Bot Building

_Libraries for building and working with bots._

- [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) - Simple and clean Telegram bot client. Stars:`4.8K`.
- [olivia](https://github.com/olivia-ai/olivia) - A chatbot built with an artificial neural network. Stars:`3.6K`.
- [telebot](https://github.com/tucnak/telebot) - Telegram bot framework is written in Go. Stars:`3.2K`.
- [wayback](https://github.com/wabarc/wayback) - A bot for Telegram, Mastodon, Slack, and other messaging platforms archives webpages. Stars:`1.4K`.
- [Kelp](https://github.com/stellar/kelp) - official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. Stars:`1.0K`.
- [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) - A golang implementation of a console-based trading bot for cryptocurrency exchanges. Stars:`928`.
- [slacker](https://github.com/shomali11/slacker) - Easy to use framework to create Slack bots. Stars:`788`.
- [go-chat-bot](https://github.com/go-chat-bot/bot) - IRC, Slack & Telegram bot written in Go. Stars:`786`.
- [tbot](https://github.com/yanzay/tbot) - Telegram bot server with API similar to net/http. Stars:`336`.
- [go-twitch-irc](https://github.com/gempir/go-twitch-irc) - Library to write bots for twitch.tv chat Stars:`302`.
- [echotron](https://github.com/NicoNex/echotron) - An elegant and concurrent library for Telegram Bots in Go. Stars:`262`.
- [go-sarah](https://github.com/oklahomer/go-sarah) - Framework to build a bot for desired chat services including LINE, Slack, Gitter, and more. Stars:`259`.
- [telego](https://github.com/mymmrac/telego) - Telegram Bot API library for Golang with full one-to-one API implementation. Stars:`233`.
- [Tenyks](https://github.com/kyleterry/tenyks) - Service oriented IRC bot using Redis and JSON for messaging. Stars:`175`.
- [bot](https://github.com/go-telegram/bot) - Zero-dependencies Telegram Bot library with additional UI components Stars:`154`.
- [hanu](https://github.com/sbstjn/hanu) - Framework for writing Slack bots. Stars:`147`.
- [larry](https://github.com/ezeoleaf/larry) - Larry 🐦 is a really simple Twitter bot generator that tweets random repositories from Github built in Go. Stars:`145`.
- [slack-bot](https://github.com/innogames/slack-bot) - Ready to use Slack Bot for lazy developers: Custom commands, Jenkins, Jira, Bitbucket, Github... Stars:`141`.
- [go-tgbot](https://github.com/olebedev/go-tgbot) - Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router, and middleware. Stars:`119`.
- [margelet](https://github.com/zhulik/margelet) - Framework for building Telegram bots. Stars:`80`.
- [ephemeral-roles](https://github.com/ewohltman/ephemeral-roles) - A Discord bot for managing ephemeral roles based upon voice channel member presence. Stars:`80`.
- [slackscot](https://github.com/alexandre-normand/slackscot) - Another framework for building Slack bots. Stars:`55`.
- [govkbot](https://github.com/nikepan/govkbot) - Simple Go [VK](https://vk.com) bot library. Stars:`43`.
- [micha](https://github.com/onrik/micha) - Go Library for Telegram bot api. Stars:`25`.
- [teleterm](https://github.com/alfiankan/teleterm) - Telegram Bot Exec Terminal Command. Stars:`25`.
- [go-joe](https://joe-bot.net) - A general-purpose bot library inspired by Hubot but written in Go.

**[⬆ back to top](#contents)**

## Build Automation

_Libraries and tools help with build automation._

- [Task](https://github.com/go-task/task) - simple "Make" alternative. Stars:`8.2K`.
- [realize](https://github.com/tockins/realize) - Go build a system with file watchers and live to reload. Run, build and watch file changes with custom paths. Stars:`4.4K`.
- [mage](https://github.com/magefile/mage) - Mage is a make/rake-like build tool using Go. Stars:`3.6K`.
- [mmake](https://github.com/tj/mmake) - Modern Make. Stars:`1.7K`.
- [xc](https://github.com/joerdav/xc) - Task runner with README.md defined tasks, executable markdown. Stars:`876`.
- [goyek](https://github.com/goyek/goyek) - Create build pipelines in Go. Stars:`391`.
- [taskctl](https://github.com/taskctl/taskctl) - Concurrent task runner. Stars:`258`.
- [1build](https://github.com/gopinath-langote/1build) - Command line tool to frictionlessly manage project-specific commands. Stars:`214`.
- [gaper](https://github.com/maxcnunes/gaper) - Builds and restarts a Go project when it crashes or some watched file changes. Stars:`74`.
- [anko](https://github.com/GuilhermeCaruso/anko) - Simple application watcher for multiple programming languages. Stars:`30`.
- [gilbert](https://go-gilbert.github.io) - Build system and task runner for Go projects.

**[⬆ back to top](#contents)**

## Command Line

### Advanced Console UIs

_Libraries for building Console Applications and Console User Interfaces._

- [bubbletea](https://github.com/charmbracelet/bubbletea) - Go framework to build terminal apps, based on The Elm Architecture. Stars:`19.2K`.
- [termui](https://github.com/gizak/termui) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). Stars:`12.6K`.
- [gocui](https://github.com/jroimartin/gocui) - Minimalist Go library aimed at creating Console User Interfaces. Stars:`9.3K`.
- [lipgloss](https://github.com/charmbracelet/lipgloss) - Declaratively define styles for color, format and layout in the terminal. Stars:`6.1K`.
- [go-prompt](https://github.com/c-bata/go-prompt) - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). Stars:`5.0K`.
- [termbox-go](https://github.com/nsf/termbox-go) - Termbox is a library for creating cross-platform text-based interfaces. Stars:`4.6K`.
- [pterm](https://github.com/pterm/pterm) - A library to beautify console output on every platform with many combinable components. Stars:`3.9K`.
- [bubbles](https://github.com/charmbracelet/bubbles) - TUI components for bubbletea. Stars:`3.6K`.
- [progressbar](https://github.com/schollz/progressbar) - Basic thread-safe progress bar that works in every OS. Stars:`3.4K`.
- [termdash](https://github.com/mum4k/termdash) - Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). Stars:`2.3K`.
- [asciigraph](https://github.com/guptarohit/asciigraph) - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. Stars:`2.3K`.
- [spinner](https://github.com/briandowns/spinner) - Go package to easily provide a terminal spinner with options. Stars:`2.1K`.
- [mpb](https://github.com/vbauerster/mpb) - Multi progress bar for terminal applications. Stars:`2.1K`.
- [uiprogress](https://github.com/gosuri/uiprogress) - Flexible library to render progress bars in terminal applications. Stars:`2.0K`.
- [uilive](https://github.com/gosuri/uilive) - Library for updating terminal output in real time. Stars:`1.6K`.
- [termenv](https://github.com/muesli/termenv) - Advanced ANSI style & color support for your terminal applications. Stars:`1.4K`.
- [gookit/color](https://github.com/gookit/color) - Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. Stars:`1.4K`.
- [aurora](https://github.com/logrusorgru/aurora) - ANSI terminal colors that support fmt.Printf/Sprintf. Stars:`1.3K`.
- [go-isatty](https://github.com/mattn/go-isatty) - isatty for golang. Stars:`724`.
- [uitable](https://github.com/gosuri/uitable) - Library to improve readability in terminal apps using tabular data. Stars:`711`.
- [go-colorable](https://github.com/mattn/go-colorable) - Colorable writer for windows. Stars:`706`.
- [simpletable](https://github.com/alexeyco/simpletable) - Simple tables in a terminal with Go. Stars:`461`.
- [chalk](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output. Stars:`428`.
- [yacspin](https://github.com/theckman/yacspin) - Yet Another CLi Spinner package, for working with terminal spinners. Stars:`404`.
- [box-cli-maker](https://github.com/Delta456/box-cli-maker) - Make Highly Customized Boxes for your CLI. Stars:`371`.
- [tabby](https://github.com/cheynewallace/tabby) - A tiny library for super simple Golang tables. Stars:`331`.
- [bubble-table](https://github.com/Evertras/bubble-table) - An interactive table component for bubbletea. Stars:`231`.
- [go-colortext](https://github.com/daviddengcn/go-colortext) - Go library for color output in terminals. Stars:`214`.
- [cfmt](https://github.com/mingrammer/cfmt) - Contextual fmt inspired by bootstrap color classes. Stars:`95`.
- [tabular](https://github.com/InVisionApp/tabular) - Print ASCII tables from command line utilities without the need to pass large sets of data to the API. Stars:`70`.
- [cfmt](https://github.com/i582/cfmt) - Simple and convenient formatted stylized output fully compatible with fmt library. Stars:`58`.
- [table](https://github.com/tomlazar/table) - Small library for terminal color based tables . Stars:`44`.
- [ctc](https://github.com/wzshiming/ctc) - The non-invasive cross-platform terminal color library does not need to modify the Print method. Stars:`44`.
- [marker](https://github.com/cyucelen/marker) - Easiest way to match and mark strings for colorful terminal outputs. Stars:`40`.
- [colourize](https://github.com/TreyBastian/colourize) - Go library for ANSI colour text in terminals. Stars:`26`.
- [go-ataman](https://github.com/workanator/go-ataman) - Go library for rendering ANSI colored text templates in terminals. Stars:`14`.
- [go-palette](https://github.com/abusomani/go-palette) - Go library that provides elegant and convenient style definitions using ANSI colors. Fully compatible & wraps the [fmt library](https://pkg.go.dev/fmt) for nice terminal layouts. Stars:`10`.
- [crab-config-files-templating](https://github.com/alfiankan/crab-config-files-templating) - Dynamic configuration file templating tool for kubernetes manifest or general configuration files. Stars:`6`.
- [gommon/color](https://github.com/labstack/gommon/tree/master/color) - Style terminal text.

**[⬆ back to top](#contents)**

### Standard CLI

_Libraries for building standard or basic Command Line applications._

- [cobra](https://github.com/spf13/cobra) - Commander for modern Go CLI interactions. Stars:`32.8K`.
- [urfave/cli](https://github.com/urfave/cli) - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). Stars:`20.5K`.
- [elvish](https://github.com/elves/elvish) - An expressive programming language and a versatile interactive shell. Stars:`5.1K`.
- [survey](https://github.com/go-survey/survey) - Build interactive and accessible prompts with full support for windows and posix terminals. Stars:`3.9K`.
- [kingpin](https://github.com/alecthomas/kingpin) - Command line and flag parser supporting sub commands (superseded by `kong`; see below). Stars:`3.4K`.
- [Dnote](https://github.com/dnote/dnote) - A simple command line notebook with multi-device sync. Stars:`2.6K`.
- [go-flags](https://github.com/jessevdk/go-flags) - go command line option parser. Stars:`2.4K`.
- [pflag](https://github.com/spf13/pflag) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. Stars:`2.1K`.
- [mitchellh/cli](https://github.com/mitchellh/cli) - Go library for implementing command-line interfaces. Stars:`1.7K`.
- [go-arg](https://github.com/alexflint/go-arg) - Struct-based argument parsing in Go. Stars:`1.7K`.
- [ops](https://github.com/nanovms/ops) - Unikernel Builder/Orchestrator. Stars:`1.1K`.
- [liner](https://github.com/peterh/liner) - Go readline-like library for command-line interfaces. Stars:`983`.
- [complete](https://github.com/posener/complete) - Write bash completions in Go + Go command bash completion. Stars:`891`.
- [mow.cli](https://github.com/jawher/mow.cli) - Go library for building CLI applications with sophisticated flag and argument parsing and validation. Stars:`854`.
- [flaggy](https://github.com/integrii/flaggy) - A robust and idiomatic flags package with excellent subcommand support. Stars:`832`.
- [cli](https://github.com/mkideal/cli) - Feature-rich and easy to use command-line package based on golang struct tags. Stars:`689`.
- [argparse](https://github.com/akamensky/argparse) - Command line argument parser inspired by Python's argparse module. Stars:`537`.
- [carapace-bin](https://github.com/rsteube/carapace-bin) - Multi-shell multi-command argument completer. Stars:`257`.
- [climax](https://github.com/tucnak/climax) - Alternative CLI with "human face", in spirit of Go command. Stars:`207`.
- [wmenu](https://github.com/dixonwille/wmenu) - Easy to use menu structure for cli applications that prompt users to make choices. Stars:`194`.
- [commandeer](https://github.com/jaffee/commandeer) - Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. Stars:`169`.
- [clîr](https://github.com/leaanthony/clir) - A Simple and Clear CLI library. Dependency free. Stars:`154`.
- [sflags](https://github.com/octago/sflags) - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin, and other libraries. Stars:`146`.
- [job](https://github.com/liujianping/job) - JOB, make your short-term command as a long-term job. Stars:`128`.
- [flag](https://github.com/cosiner/flag) - Simple but powerful command line option parsing library for Go supporting subcommand. Stars:`126`.
- [cmdr](https://github.com/hedzr/cmdr) - A POSIX/GNU style, getopt-like command-line UI Go library. Stars:`123`.
- [ukautz/clif](https://github.com/ukautz/clif) - Small command line interface framework. Stars:`123`.
- [cli](https://github.com/teris-io/cli) - Simple and complete API for building command line interfaces in Go. Stars:`120`.
- [carapace](https://github.com/rsteube/carapace) - Command argument completion generator for spf13/cobra. Stars:`116`.
- [env](https://github.com/codingconcepts/env) - Tag-based environment configuration for structs. Stars:`106`.
- [acmd](https://github.com/cristalhq/acmd) - Simple, useful, and opinionated CLI package in Go. Stars:`91`.
- [version](https://github.com/mszostok/version) - Collects and displays CLI version information in multiple formats along with upgrade notice. Stars:`71`.
- [gocmd](https://github.com/devfacet/gocmd) - Go library for building command line applications. Stars:`65`.
- [wlog](https://github.com/dixonwille/wlog) - Simple logging interface that supports cross-platform color and concurrency. Stars:`59`.
- [strumt](https://github.com/antham/strumt) - Library to create prompt chain. Stars:`55`.
- [go-getoptions](https://github.com/DavidGamba/go-getoptions) - Go option parser inspired by the flexibility of Perl’s GetOpt::Long. Stars:`47`.
- [command-chain](https://github.com/rainu/go-command-chain) - A go library for configure and run command chains - such as pipelining in unix shells. Stars:`44`.
- [flagvar](https://github.com/sgreben/flagvar) - A collection of flag argument types for Go's standard `flag` package. Stars:`42`.
- [readline](https://github.com/reeflective/readline) Shell library with modern and easy to use UI features. Stars:`41`.
- [argv](https://github.com/cosiner/argv) - Go library to split command line string as arguments array using the bash syntax. Stars:`38`.
- [cmd](https://github.com/posener/cmd) - Extends the standard `flag` package to support sub commands and more in idiomatic way. Stars:`37`.
- [go-commander](https://github.com/yitsushi/go-commander) - Go library to simplify CLI workflow. Stars:`33`.
- [go-andotp](https://github.com/grijul/go-andotp) - A CLI program to encrypt/decrypt [andOTP](https://github.com/andOTP/andOTP) files. Can be used as a library as well. Stars:`23`.
- [sand](https://github.com/Zaba505/sand) - Simple API for creating interpreters and so much more. Stars:`20`.
- [ts](https://github.com/liujianping/ts) - Timestamp convert & compare tool. Stars:`17`.
- [mcli](https://github.com/jxskiss/mcli) - A minimal but very powerful cli library for Go. Stars:`13`.
- [carapace-spec](https://github.com/rsteube/carapace-spec) - Define simple completions using a spec file. Stars:`7`.
- [subcmd](https://github.com/bobg/subcmd) - Another approach to parsing and running subcommands. Works alongside the standard `flag` package. Stars:`4`.
- [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli) - cli application framework with auto configuration and dependency injection.

**[⬆ back to top](#contents)**

## Configuration

_Libraries for configuration parsing._

- [viper](https://github.com/spf13/viper) - Go configuration with fangs. Stars:`23.6K`.
- [godotenv](https://github.com/joho/godotenv) - Go port of Ruby's dotenv library (Loads environment variables from `.env`). Stars:`6.5K`.
- [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) - Go library for managing configuration data from environment variables. Stars:`4.6K`.
- [env](https://github.com/caarlos0/env) - Parse environment variables to Go structs (with defaults). Stars:`3.7K`.
- [ini](https://github.com/go-ini/ini) - Go package to read and write INI files. Stars:`3.3K`.
- [koanf](https://github.com/knadh/koanf) - Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. Stars:`1.9K`.
- [kong](https://github.com/alecthomas/kong) - Command-line parser with support for arbitrarily complex command-line structures and additional sources of configuration such as YAML, JSON, TOML, etc (successor to `kingpin`). Stars:`1.5K`.
- [cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Minimalistic configuration reader (from files, ENV, and wherever you want). Stars:`1.1K`.
- [konfig](https://github.com/lalamove/konfig) - Composable, observable and performant config handling for Go for the distributed processing era. Stars:`639`.
- [gookit/config](https://github.com/gookit/config) - application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. Stars:`475`.
- [confita](https://github.com/heetch/confita) - Load configuration in cascade from multiple backends into a struct. Stars:`471`.
- [aconfig](https://github.com/cristalhq/aconfig) - Simple, useful and opinionated config loader. Stars:`469`.
- [xdg](https://github.com/adrg/xdg) - Go implementation of the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) and [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories). Stars:`422`.
- [GoLobby/Config](https://github.com/golobby/config) - GoLobby Config is a lightweight yet powerful configuration manager for the Go programming language. Stars:`336`.
- [config](https://github.com/JeremyLoy/config) - Cloud native application configuration. Bind ENV to structs in only two lines. Stars:`329`.
- [hjson](https://github.com/hjson/hjson-go) - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. Stars:`303`.
- [fig](https://github.com/kkyr/fig) - Tiny library for reading configuration from a file and from environment variables (with validation & defaults). Stars:`275`.
- [store](https://github.com/tucnak/store) - Lightweight configuration manager for Go. Stars:`271`.
- [config](https://github.com/olebedev/config) - JSON or YAML configuration wrapper with environment variables and flags parsing. Stars:`265`.
- [envconfig](https://github.com/vrischmann/envconfig) - Read your configuration from environment variables. Stars:`234`.
- [joshbetz/config](https://github.com/joshbetz/config) - Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. Stars:`215`.
- [gcfg](https://github.com/go-gcfg/gcfg) - read INI-style configuration files into Go structs; supports user-defined types and subsections. Stars:`164`.
- [harvester](https://github.com/beatlabs/harvester) - Harvester, a easy to use static and dynamic configuration package supporting seeding, env vars and Consul integration. Stars:`126`.
- [onion](https://github.com/goraz/onion) - Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP. Stars:`113`.
- [envcfg](https://github.com/tomazk/envcfg) - Un-marshaling environment variables to Go structs. Stars:`100`.
- [envh](https://github.com/antham/envh) - Helpers to manage environment variables. Stars:`96`.
- [configuro](https://github.com/sherifabdlnaby/configuro) - opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications. Stars:`87`.
- [configuration](https://github.com/BoRuDar/configuration) - Library for initializing configuration structs from env variables, files, flags and 'default' tag. Stars:`84`.
- [xdg](https://github.com/OpenPeeDeeP/xdg) - Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). Stars:`76`.
- [gofigure](https://github.com/ian-kent/gofigure) - Go application configuration made easy. Stars:`66`.
- [uConfig](https://github.com/omeid/uconfig) - Lightweight, zero-dependency, and extendable configuration management. Stars:`62`.
- [hocon](https://github.com/gurkankaymak/hocon) - Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files. Stars:`60`.
- [configure](https://github.com/paked/configure) - Provides configuration through multiple sources, including JSON, flags and environment variables. Stars:`57`.
- [go-aws-ssm](https://github.com/PaddleHQ/go-aws-ssm) - Go package that fetches parameters from AWS System Manager - Parameter Store. Stars:`54`.
- [go-up](https://github.com/ufoscout/go-up) - A simple configuration library with recursive placeholders resolution and no magic. Stars:`42`.
- [env](https://github.com/junk1tm/env) - A lightweight package for loading environment variables into structs. Stars:`40`.
- [config](https://github.com/num30/config) - configure you app using file, environment variables, or flags in two lines of code Stars:`37`.
- [ingo](https://github.com/schachmat/ingo) - Flags persisted in an ini-like config file. Stars:`36`.
- [mini](https://github.com/sasbury/mini) - Golang package for parsing ini-style configuration files. Stars:`34`.
- [genv](https://github.com/sakirsensoy/genv) - Read environment variables easily with dotenv support. Stars:`33`.
- [conflate](https://github.com/the4thamigo-uk/conflate) - Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. Stars:`27`.
- [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) - Go utility for loading configuration parameters from AWS SSM (Parameter Store). Stars:`17`.
- [ini](https://github.com/wlevene/ini) - INI Parser & Write Library, Unmarshal to Struct,Marshal to Json,Write File,watch file. Stars:`12`.
- [envconf](https://github.com/ian-kent/envconf) - Configuration from environment. Stars:`11`.
- [nasermirzaei89/env](https://github.com/nasermirzaei89/env) - Simple useful package for read environment variables. Stars:`10`.
- [go-ini](https://github.com/subpop/go-ini) - A Go package that marshals and unmarshals INI-files. Stars:`9`.
- [typenv](https://github.com/diegomarangoni/typenv) - Minimalistic, zero dependency, typed environment variables library. Stars:`9`.
- [swap](https://github.com/oblq/swap) - Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env). Stars:`8`.
- [go-conf](https://github.com/ThomasObenaus/go-conf) - Simple library for application configuration based on annotated structs. It supports reading the configuration from environment variables, config files and command line parameters. Stars:`8`.
- [piper](https://github.com/Yiling-J/piper) - Viper wrapper with config inheritance and key generation. Stars:`7`.
- [gonfig](https://github.com/milad-abbasi/gonfig) - Tag-based configuration parser which loads values from different providers into typesafe struct. Stars:`6`.
- [nfigure](https://github.com/muir/nfigure) - Per-library struct-tag based configuration from command lines (Posix & Go-style); environment, JSON, YAML Stars:`5`.
- [goConfig](https://github.com/crgimenes/goConfig) - Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. Stars:`3`.
- [gone/jconf](https://github.com/One-com/gone/tree/master/jconf) - Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.

**[⬆ back to top](#contents)**

## Continuous Integration

_Tools for help with continuous integration._

- [drone](https://github.com/drone/drone) - Drone is a Continuous Integration platform built on Docker, written in Go. Stars:`27.1K`.
- [CDS](https://github.com/ovh/cds) - Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. Stars:`4.3K`.
- [woodpecker](https://github.com/woodpecker-ci/woodpecker) - Woodpecker is a community fork of the Drone CI system. Stars:`2.9K`.
- [goveralls](https://github.com/mattn/goveralls) - Go integration for Coveralls.io continuous code coverage tracking system. Stars:`765`.
- [gotestfmt](https://github.com/GoTestTools/gotestfmt) - go test output for humans. Stars:`408`.
- [overalls](https://github.com/go-playground/overalls) - Multi-Package go project coverprofile for tools like goveralls. Stars:`112`.
- [duci](https://github.com/duck8823/duci) - A simple ci server no needs domain specific languages. Stars:`74`.
- [gomason](https://github.com/nikogura/gomason) - Test, Build, Sign, and Publish your go binaries from a clean workspace. Stars:`57`.
- [roveralls](https://github.com/LawrenceWoodman/roveralls) - Recursive coverage testing tool. Stars:`19`.
- [go-test-coverage](https://github.com/vladopajic/go-test-coverage) - Tool and GitHub action which reports issues when test coverage is below set threshold. Stars:`15`.
- [go-fuzz-action](https://github.com/jidicula/go-fuzz-action) - Use Go 1.18's built-in fuzz testing in GitHub Actions. Stars:`9`.

**[⬆ back to top](#contents)**

## CSS Preprocessors

_Libraries for preprocessing CSS files._

- [gcss](https://github.com/yosssi/gcss) - Pure Go CSS Preprocessor. Stars:`483`.
- [go-libsass](https://github.com/wellington/go-libsass) - Go wrapper to the 100% Sass compatible libsass project. Stars:`196`.

**[⬆ back to top](#contents)**

## Data Structures and Algorithms

### Bit-packing and Compression

- [roaring](https://github.com/RoaringBitmap/roaring) - Go package implementing compressed bitsets. Stars:`2.1K`.
- [binpacker](https://github.com/zhuangsirui/binpacker) - Binary packer and unpacker helps user build custom binary stream. Stars:`210`.
- [bit](https://github.com/yourbasic/bit) - Golang set data structure with bonus bit-twiddling functions. Stars:`150`.
- [crunch](https://github.com/superwhiskers/crunch) - Go package implementing buffers for handling various datatypes easily. Stars:`87`.
- [go-ef](https://github.com/amallia/go-ef) - A Go implementation of the Elias-Fano encoding. Stars:`30`.
- [bingo](https://github.com/iancmcc/bingo) - Fast, zero-allocation, lexicographical-order-preserving packing of native types to bytes. Stars:`23`.

### Bit Sets

- [bitset](https://github.com/bits-and-blooms/bitset) - Go package implementing bitsets. Stars:`1.1K`.
- [bitmap](https://github.com/kelindar/bitmap) - Dense, zero-allocation, SIMD-enabled bitmap/bitset in Go. Stars:`224`.

### Bloom and Cuckoo Filters

- [bloom](https://github.com/bits-and-blooms/bloom) - Go package implementing Bloom filters. Stars:`2.0K`.
- [boomfilters](https://github.com/tylertreat/BoomFilters) - Probabilistic data structures for processing continuous, unbounded streams. Stars:`1.5K`.
- [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. Stars:`998`.
- [cuckoo-filter](https://github.com/linvon/cuckoo-filter) - Cuckoo filter: a comprehensive cuckoo filter, which is configurable and space optimized compared with other implements, and all features mentioned in original paper is available. Stars:`262`.
- [bloom](https://github.com/zhenjl/bloom) - Bloom filters implemented in Go. Stars:`148`.
- [ring](https://github.com/TheTannerRyan/ring) - Go implementation of a high performance, thread safe bloom filter. Stars:`130`.
- [bloom](https://github.com/yourbasic/bloom) - Golang Bloom filter implementation. Stars:`79`.
- [bloomfilter](https://github.com/OldPanda/bloomfilter) - Yet another Bloomfilter implementation in Go, compatible with Java's Guava library. Stars:`13`.

### Data Structure and Algorithm Collections

- [gods](https://github.com/emirpasic/gods) - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. Stars:`14.1K`.
- [go-datastructures](https://github.com/Workiva/go-datastructures) - Collection of useful, performant, and thread-safe data structures. Stars:`7.1K`.
- [gostl](https://github.com/liyue201/gostl) - Data structure and algorithm library for go, designed to provide functions similar to C++ STL. Stars:`898`.
- [algorithms](https://github.com/shady831213/algorithms) - Algorithms and data structures.CLRS study. Stars:`733`.

### Iterators

- [iter](https://github.com/disksing/iter) - Go implementation of C++ STL iterators and algorithms. Stars:`179`.
- [goterator](https://github.com/yaa110/goterator) - Iterator implementation to provide map and reduce functionalities. Stars:`14`.

### Maps

See also [Database](#database) for more complex key-value stores, and [Trees](#trees) for
additional ordered map implementations.

- [cmap](https://github.com/lrita/cmap) - a thread-safe concurrent map for go, support using `interface{}` as key and auto scale up shards. Stars:`61`.
- [dict](https://github.com/srfrog/dict) - Python-like dictionaries (dict) for Go. Stars:`39`.
- [goradd/maps](https://github.com/goradd/maps) - Go 1.18+ generic map interface for maps; safe maps; ordered maps; ordered, safe maps; etc. Stars:`22`.

### Miscellaneous Data Structures and Algorithms

- [gota](https://github.com/kniren/gota) - Implementation of dataframes, series, and data wrangling methods for Go. Stars:`2.7K`.
- [hyperloglog](https://github.com/axiomhq/hyperloglog) - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. Stars:`868`.
- [go-geoindex](https://github.com/hailocab/go-geoindex) - In-memory geo index. Stars:`350`.
- [hilbert](https://github.com/google/hilbert) - Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. Stars:`266`.
- [go-rquad](https://github.com/aurelien-rainone/go-rquad) - Region quadtrees with efficient point location and neighbour finding. Stars:`126`.
- [conjungo](https://github.com/InVisionApp/conjungo) - A small, powerful and flexible merge library. Stars:`116`.
- [go-rampart](https://github.com/francesconi/go-rampart) - Determine how intervals relate to each other. Stars:`90`.
- [gogu](https://github.com/esimov/gogu) - A comprehensive, reusable and efficient concurrent-safe generics utility functions and data structures library. Stars:`80`.
- [go-tuple](https://github.com/barweiss/go-tuple) - Generic tuple implementation for Go 1.18+. Stars:`63`.
- [count-min-log](https://github.com/seiflotfy/count-min-log) - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). Stars:`63`.
- [hide](https://github.com/emvi/hide) - ID type with marshalling to/from hash to prevent sending IDs to clients. Stars:`54`.
- [concurrent-writer](https://github.com/free/concurrent-writer) - Highly concurrent drop-in replacement for `bufio.Writer`. Stars:`49`.
- [genfuncs](https://github.com/nwillc/genfuncs) - Go 1.18+ generics package inspired by Kotlin's Sequence and Map. Stars:`46`.
- [fsm](https://github.com/cocoonspace/fsm) - Finite-State Machine package. Stars:`45`.
- [go18ds](https://github.com/daichi-m/go18ds) - Go Data Structures using Go 1.18 generics. Stars:`32`.
- [go-generics](https://github.com/bobg/go-generics) - Generic slice, map, set, iterator, and goroutine utilities. Stars:`30`.
- [quadtree](https://github.com/s0rg/quadtree) - Generic, zero-alloc, 100%-test covered quadtree. Stars:`27`.
- [gofal](https://github.com/xxjwxc/gofal) - fractional api for Go. Stars:`17`.
- [slices](https://github.com/srfrog/slices) - Functions that operate on slices; like `package strings` but adapted to work with slices. Stars:`13`.
- [slices](https://github.com/twharmon/slices) - Pure, generic functions for slices. Stars:`13`.

### Nullable Types

- [nan](https://github.com/kak-tus/nan) - Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers. Stars:`75`.
- [typ](https://github.com/gurukami/typ) - Null Types, Safe primitive type conversion and fetching value from complex structures. Stars:`42`.
- [null](https://github.com/emvi/null) - Nullable Go types that can be marshalled/unmarshalled to/from JSON. Stars:`31`.

### Queues

- [deque](https://github.com/gammazero/deque) - Fast ring-buffer deque (double-ended queue). Stars:`473`.
- [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) - Concurrent FIFO queue. Stars:`296`.
- [queue](https://github.com/adrianbrad/queue) - Multiple thread-safe, generic queue implementations for Go. Stars:`179`.
- [deque](https://github.com/edwingeng/deque) - A highly optimized double-ended queue. Stars:`135`.
- [memlog](https://github.com/embano1/memlog) - An easy to use, lightweight, thread-safe and append-only in-memory data structure inspired by Apache Kafka. Stars:`98`.

### Sets

- [golang-set](https://github.com/deckarep/golang-set) - Thread-Safe and Non-Thread-Safe high-performance sets for Go. Stars:`3.4K`.
- [goset](https://github.com/zoumo/goset) - A useful Set collection implementation for Go. Stars:`51`.
- [set](https://github.com/StudioSol/set) - Simple set data structure implementation in Go using LinkedHashMap. Stars:`23`.
- [dsu](https://github.com/ihebu/dsu) - Disjoint Set data structure implementation in Go. Stars:`11`.

### Text Analysis

- [bleve](https://github.com/blevesearch/bleve) - Modern text indexing library for go. Stars:`9.2K`.
- [trie](https://github.com/derekparker/trie) - Trie implementation in Go. Stars:`655`.
- [go-edlib](https://github.com/hbollon/go-edlib) - Go string comparison and edit distance algorithms library (Levenshtein, LCS, Hamming, Damerau levenshtein, Jaro-Winkler, etc.) compatible with Unicode. Stars:`396`.
- [levenshtein](https://github.com/agnivade/levenshtein) - Implementation to calculate levenshtein distance in Go. Stars:`295`.
- [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) - Go implementation of Adaptive Radix Tree. Stars:`293`.
- [levenshtein](https://github.com/agext/levenshtein) - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. Stars:`77`.
- [ptrie](https://github.com/viant/ptrie) - An implementation of prefix tree. Stars:`33`.
- [mspm](https://github.com/BlackRabbitt/mspm) - Multi-String Pattern Matching Algorithm for information retrieval. Stars:`24`.
- [parsefields](https://github.com/MonaxGT/parsefields) - Tools for parse JSON-like logs for collecting unique fields and events. Stars:`7`.

### Trees

- [skiplist](https://github.com/MauriceGit/skiplist) - Very fast Go Skiplist implementation. Stars:`253`.
- [skiplist](https://github.com/gansidui/skiplist) - Skiplist implementation in Go. Stars:`82`.
- [treemap](https://github.com/igrmk/treemap) - Generic key-sorted map using a red-black tree under the hood. Stars:`46`.
- [treap](https://github.com/perdata/treap) - Persistent, fast ordered map using tree heaps. Stars:`22`.
- [merkle](https://github.com/bobg/merkle) - Space-efficient computation of Merkle root hashes and inclusion proofs. Stars:`10`.
- [hashsplit](http://github.com/bobg/hashsplit) - Split byte streams into chunks, and arrange chunks into trees, with boundaries determined by content, not position.

### Pipes

- [pipeline](https://github.com/hyfather/pipeline) - An implementation of pipelines with fan-in and fan-out. Stars:`50`.
- [ordered-concurrently](https://github.com/tejzpr/ordered-concurrently) - Go module that processes work concurrently and returns output in a channel in the order of input. Stars:`23`.
- [parapipe](https://github.com/nazar256/parapipe) - FIFO Pipeline which parallels execution on each stage while maintaining the order of messages and results. Stars:`23`.

**[⬆ back to top](#contents)**

## Database

### Caches

_Data stores with expiring records, in-memory distributed data stores, or in-memory subsets of file-based databases._

- [groupcache](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. Stars:`12.2K`.
- [BigCache](https://github.com/allegro/bigcache) - Efficient key/value cache for gigabytes of data. Stars:`6.7K`.
- [GCache](https://github.com/bluele/gcache) - Cache library with support for expirable Cache, LFU, LRU and ARC. Stars:`2.4K`.
- [cache2go](https://github.com/muesli/cache2go) - In-memory key:value cache which supports automatic invalidation based on timeouts. Stars:`2.0K`.
- [gocache](https://github.com/eko/gocache) - A complete Go cache library with multiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. Stars:`1.9K`.
- [fastcache](https://github.com/VictoriaMetrics/fastcache) - fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. Stars:`1.8K`.
- [ttlcache](https://github.com/jellydator/ttlcache) - An in-memory cache with item expiration and generics. Stars:`684`.
- [cachego](https://github.com/faabiosr/cachego) - Golang Cache component for multiple drivers. Stars:`284`.
- [cache](https://github.com/akyoto/cache) - In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. Stars:`140`.
- [remember-go](https://github.com/rocketlaunchr/remember-go) - A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory). Stars:`133`.
- [bcache](https://github.com/iwanbk/bcache) - Eventually consistent distributed in-memory cache Go library. Stars:`125`.
- [go-cache](https://github.com/viney-shih/go-cache) - A flexible multi-layer Go caching library to deal with in-memory and shared cache by adopting Cache-Aside pattern. Stars:`111`.
- [theine](https://github.com/Yiling-J/theine-go) - High performance, near optimal in-memory cache with proactive TTL expiration and generics. Stars:`104`.
- [go-mcache](https://github.com/OrlovEvgeny/go-mcache) - Fast in-memory key:value store/cache library. Pointer caches. Stars:`88`.
- [imcache](https://github.com/erni27/imcache) - A generic in-memory cache Go library. It supports expiration, sliding expiration, max entries limit, eviction callbacks and sharding. Stars:`87`.
- [timedmap](https://github.com/zekroTJA/timedmap) - Map with expiring key-value pairs. Stars:`64`.
- [couchcache](https://github.com/codingsince1985/couchcache) - RESTful caching micro-service backed by Couchbase server. Stars:`60`.
- [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) - BigCache with clustering support and individual item expiration. Stars:`43`.
- [2q](https://github.com/floatdrop/2q) - 2Q in-memory cache implementation. Stars:`27`.
- [gdcache](https://github.com/ulovecode/gdcache) - A pure non-intrusive cache library implemented by golang, you can use it to implement your own distributed cache. Stars:`10`.
- [ttlcache](https://github.com/cheshir/ttlcache) - In-memory key value storage with TTL for each record. Stars:`9`.
- [nscache](https://github.com/no-src/nscache) - A Go caching framework that supports multiple data source drivers. Stars:`4`.
- [gocache](https://github.com/yuseferi/gocache) - A data race free Go ache library with high performance and auto pruge functionality Stars:`3`.

### Databases Implemented in Go

- [prometheus](https://github.com/prometheus/prometheus) - Monitoring system and time series database. Stars:`49.3K`.
- [tidb](https://github.com/pingcap/tidb) - TiDB is a distributed SQL database. Inspired by the design of Google F1. Stars:`34.5K`.
- [cockroach](https://github.com/cockroachdb/cockroach) - Scalable, Geo-Replicated, Transactional Datastore. Stars:`27.5K`.
- [influxdb](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics. Stars:`25.9K`.
- [Milvus](https://github.com/milvus-io/milvus) - Milvus is a vector database for embedding management, analytics and search. Stars:`21.9K`.
- [dgraph](https://github.com/dgraph-io/dgraph) - Scalable, Distributed, Low Latency, High Throughput Graph Database. Stars:`19.5K`.
- [dolt](https://github.com/dolthub/dolt) - Dolt – It's Git for Data. Stars:`15.4K`.
- [rqlite](https://github.com/rqlite/rqlite) - The lightweight, distributed, relational database built on SQLite. Stars:`13.9K`.
- [badger](https://github.com/dgraph-io/badger) - Fast key-value store in Go. Stars:`12.5K`.
- [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) - fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. Stars:`9.0K`.
- [immudb](https://github.com/codenotary/immudb) - immudb is a lightweight, high-speed immutable database for systems and applications written in Go. Stars:`8.3K`.
- [bbolt](https://github.com/etcd-io/bbolt) - An embedded key/value database for Go. Stars:`6.8K`.
- [goleveldb](https://github.com/syndtr/goleveldb) - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. Stars:`5.8K`.
- [buntdb](https://github.com/tidwall/buntdb) - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. Stars:`4.2K`.
- [ledisdb](https://github.com/siddontang/ledisdb) - Ledisdb is a high performance NoSQL like Redis based on LevelDB. Stars:`4.0K`.
- [rosedb](https://github.com/roseduan/rosedb) - An embedded k-v database based on LSM+WAL, supports string, list, hash, set, zset. Stars:`3.9K`.
- [nutsdb](https://github.com/xujiajun/nutsdb) - Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as list, set, sorted set. Stars:`3.0K`.
- [godis](https://github.com/hdt3213/godis) - A Golang implemented high-performance Redis server and cluster. Stars:`2.9K`.
- [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang. Stars:`2.7K`.
- [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) - CovenantSQL is a SQL database on blockchain. Stars:`1.4K`.
- [diskv](https://github.com/peterbourgon/diskv) - Home-grown disk-backed key-value store. Stars:`1.3K`.
- [column](https://github.com/kelindar/column) - High-performance, columnar, embeddable in-memory store with bitmap indexing and transactions. Stars:`1.3K`.
- [lotusdb](https://github.com/flower-corp/lotusdb) - Fast k/v database compatible with lsm and b+tree. Stars:`1.2K`.
- [Databunker](https://github.com/paranoidguy/databunker) - Personally identifiable information (PII) storage service built to comply with GDPR and CCPA. Stars:`1.1K`.
- [pogreb](https://github.com/akrylysov/pogreb) - Embedded key-value store for read-heavy workloads. Stars:`1.1K`.
- [eliasdb](https://github.com/krotik/eliasdb) - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. Stars:`972`.
- [moss](https://github.com/couchbase/moss) - Moss is a simple LSM key-value storage engine written in 100% Go. Stars:`928`.
- [objectbox-go](https://github.com/objectbox/objectbox-go) - High-performance embedded Object Database (NoSQL) with Go API. Stars:`892`.
- [clover](https://github.com/ostafen/clover) - A lightweight document-oriented NoSQL database written in pure Golang. Stars:`443`.
- [levigo](https://github.com/jmhodges/levigo) - Levigo is a Go wrapper for LevelDB. Stars:`417`.
- [pudge](https://github.com/recoilme/pudge) - Fast and simple key/value store written using Go's standard library. Stars:`343`.
- [Vasto](https://github.com/chrislusf/vasto) - A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. Stars:`249`.
- [dtf](https://github.com/dtm-labs/dtf) - A distributed transaction manager. Support XA, TCC, SAGA, Reliable Messages. Stars:`217`.
- [piladb](https://github.com/fern4lvarez/piladb) - Lightweight RESTful database engine based on stack data structures. Stars:`201`.
- [unitdb](https://github.com/unit-io/unitdb) - Fast timeseries database for IoT, realtime messaging applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application. Stars:`108`.
- [libradb](https://github.com/amit-davidson/LibraDB) - LibraDB is a simple database with less then 1000 lines of code for learning. Stars:`101`.
- [hare](https://github.com/jameycribbs/hare) - A simple database management system that stores each table as a text file of line-delimited JSON. Stars:`78`.
- [regatta](https://github.com/jamf/regatta) - Fast, simple, geo-distributed KV store built for cloud native era. Stars:`55`.
- [Coffer](https://github.com/claygod/coffer) - Simple ACID key-value database that supports transactions. Stars:`34`.
- [tempdb](https://github.com/rafaeljesus/tempdb) - Key-value store for temporary items. Stars:`17`.
- [Bitcask](https://git.mills.io/prologic/bitcask) - Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).

### Database Schema Migration

- [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library. Stars:`12.0K`.
- [bytebase](https://github.com/bytebase/bytebase) - Safe database schema change and version control for DevOps teams. Stars:`6.4K`.
- [goose](https://github.com/pressly/goose) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. Stars:`4.2K`.
- [dbmate](https://github.com/amacneil/dbmate) - A lightweight, framework-agnostic database migration tool. Stars:`3.6K`.
- [atlas](https://github.com/ariga/atlas) - A Database Toolkit. A CLI designed to help companies better work with their data. Stars:`3.4K`.
- [sql-migrate](https://github.com/rubenv/sql-migrate) - Database migration tool. Allows embedding migrations into the application using go-bindata. Stars:`2.9K`.
- [skeema](https://github.com/skeema/skeema) - Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. Stars:`1.2K`.
- [gormigrate](https://github.com/go-gormigrate/gormigrate) - Database schema migration helper for Gorm ORM. Stars:`919`.
- [goavro](https://github.com/linkedin/goavro) - A Go package that encodes and decodes Avro data. Stars:`901`.
- [migrator](https://github.com/lopezator/migrator) - Dead simple Go database migration library. Stars:`154`.
- [darwin](https://github.com/GuiaBolso/darwin) - Database schema evolution library for Go. Stars:`135`.
- [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) - A Go package to help write migrations with go-pg/pg. Stars:`83`.
- [sqlize](https://github.com/sunary/sqlize) - Database migration generator. Allows generate sql migration from model and existing sql by differ them. Stars:`64`.
- [avro](https://github.com/khezen/avro) - Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. Stars:`43`.
- [go-fixtures](https://github.com/RichardKnop/go-fixtures) - Django style fixtures for Golang's excellent built-in database/sql library. Stars:`29`.
- [schema](https://github.com/adlio/schema) - Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. Stars:`29`.
- [migrator](https://github.com/larapulse/migrator) - MySQL database migrator designed to run migrations to your features and manage database schema update with intuitive go code. Stars:`24`.
- [libschema](https://github.com/muir/libschema) - Define your migrations separately in each library. Migrations for open source libraries. MySQL & PostgreSQL. Stars:`13`.
- [go-pg-migrate](https://github.com/lawzava/go-pg-migrate) - CLI-friendly package for go-pg migrations management. Stars:`11`.
- [gorm-seeder](https://github.com/Kachit/gorm-seeder) - Simple database seeder for Gorm ORM. Stars:`9`.
- [godfish](https://github.com/rafaelespinoza/godfish) - Database migration manager, works with native query language. Support for cassandra, mysql, postgres, sqlite3. Stars:`5`.
- [soda](https://github.com/gobuffalo/pop/tree/master/soda) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.

### Database Tools

- [vitess](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. Stars:`16.6K`.
- [pgweb](https://github.com/sosedoff/pgweb) - Web-based PostgreSQL database browser. Stars:`8.0K`.
- [kingshard](https://github.com/flike/kingshard) - kingshard is a high performance proxy for MySQL powered by Golang. Stars:`6.2K`.
- [orchestrator](https://github.com/github/orchestrator) - MySQL replication topology manager & visualizer. Stars:`5.1K`.
- [go-mysql](https://github.com/siddontang/go-mysql) - Go toolset to handle MySQL protocol and replication. Stars:`4.2K`.
- [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) - Sync your MySQL data into Elasticsearch automatically. Stars:`4.0K`.
- [pREST](https://github.com/prest/prest) - Simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new. Stars:`3.9K`.
- [chproxy](https://github.com/Vertamedia/chproxy) - HTTP proxy for ClickHouse database. Stars:`1.1K`.
- [pg_timetable](https://github.com/cybertec-postgresql/pg_timetable) - Advanced scheduling for PostgreSQL. Stars:`895`.
- [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) - Collects small inserts and sends big requests to ClickHouse servers. Stars:`427`.
- [rdb](https://github.com/HDT3213/rdb) - Redis RDB file parser for secondary development and memory analysis. Stars:`273`.
- [octillery](https://github.com/knocknote/octillery) - Go package for sharding databases ( Supports every ORM or raw SQL ). Stars:`183`.
- [dbbench](https://github.com/sj14/dbbench) - Database benchmarking tool with support for several databases and scripts. Stars:`88`.
- [datagen](https://github.com/codingconcepts/datagen) - A fast data generator that's multi-table aware and supports multi-row DML. Stars:`56`.
- [prep](https://github.com/hexdigest/prep) - Use prepared SQL statements without changing your code. Stars:`33`.
- [onedump](https://github.com/liweiyi88/onedump) - Database backup from different drivers to different destinations with one command and configuration. Stars:`20`.
- [rwdb](https://github.com/andizzle/rwdb) - rwdb provides read replica capability for multiple database servers setup. Stars:`18`.
- [dynago](https://github.com/twharmon/dynago) - Simplify working with AWS DynamoDB. Stars:`8`.
- [hasql](https://golang.yandex/hasql) - Library for accessing multi-host SQL database installations.

### SQL Query Builders

_Libraries for building and using SQL._

- [sqlc](https://github.com/kyleconroy/sqlc) - Generate type-safe code from SQL. Stars:`8.7K`.
- [Squirrel](https://github.com/Masterminds/squirrel) - Go library that helps you build SQL queries. Stars:`5.9K`.
- [xo](https://github.com/knq/xo) - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. Stars:`3.4K`.
- [goqu](https://github.com/doug-martin/goqu) - Idiomatic SQL builder and query library. Stars:`1.9K`.
- [gendry](https://github.com/didi/gendry) - Non-invasive SQL builder and powerful data binder. Stars:`1.5K`.
- [jet](https://github.com/go-jet/jet) - Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. Stars:`1.4K`.
- [Dotsql](https://github.com/gchaincl/dotsql) - Go library that helps you keep sql files in one place and use them with ease. Stars:`677`.
- [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) - Powerful data retrieval methods as well as DB-agnostic query building capabilities. Stars:`606`.
- [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go. Stars:`387`.
- [sqlingo](https://github.com/lqs/sqlingo) - A lightweight DSL to build SQL in Go. Stars:`327`.
- [sqrl](https://github.com/elgris/sqrl) - SQL query builder, fork of Squirrel with improved performance. Stars:`257`.
- [sq](https://github.com/bokwoon95/go-structured-query) - Type-safe SQL builder and struct mapper for Go. Stars:`193`.
- [sqlf](https://github.com/leporo/sqlf) - Fast SQL query builder. Stars:`124`.
- [buildsqlx](https://github.com/arthurkushman/buildsqlx) - Go database query builder library for PostgreSQL. Stars:`121`.
- [igor](https://github.com/galeone/igor) - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. Stars:`111`.
- [bqb](https://github.com/nullism/bqb) - Lightweight and easy to learn query builder. Stars:`96`.
- [builq](https://github.com/cristalhq/builq) - Easily build SQL queries in Go. Stars:`66`.
- [godbal](https://github.com/xujiajun/godbal) - Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. Stars:`58`.
- [gosql](https://github.com/twharmon/gosql) - SQL Query builder with better null values support. Stars:`30`.
- [qry](https://github.com/HnH/qry) - Tool that generates constants from files with raw SQL queries. Stars:`29`.
- [ormlite](https://github.com/pupizoid/ormlite) - Lightweight package containing some ORM-like features and helpers for sqlite databases. Stars:`10`.
- [sg](https://github.com/go-the-way/sg) - A SQL Gen for generating standard SQLs(supports: CRUD) written in Go. Stars:`4`.
- [Squalus](https://gitlab.com/qosenergy/squalus) - Thin layer over the Go SQL package that makes it easier to perform queries.

**[⬆ back to top](#contents)**

## Database Drivers

### Interfaces to Multiple Backends

- [cayley](https://github.com/google/cayley) - Graph database with support for multiple backends. Stars:`14.6K`.
- [gokv](https://github.com/philippgille/gokv) - Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). Stars:`570`.
- [go-transaction-manager](https://github.com/avito-tech/go-transaction-manager) - Transaction manager with multiple adapters (sql, sqlx, gorm, mongo, ...) controls transaction boundaries. Stars:`64`.
- [dsc](https://github.com/viant/dsc) - Datastore connectivity for SQL, NoSQL, structured files. Stars:`27`.

### Relational Database Drivers

- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go. Stars:`13.6K`.
- [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql. Stars:`8.2K`.
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver supporting features beyond those exposed by database/sql. Stars:`7.5K`.
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that uses database/sql. Stars:`6.8K`.
- [go-mssqldb](https://github.com/denisenkom/go-mssqldb) - Microsoft MSSQL driver for Go. Stars:`1.7K`.
- [sqlhooks](https://github.com/qustavo/sqlhooks) - Attach hooks to any database/sql driver. Stars:`611`.
- [go-oci8](https://github.com/mattn/go-oci8) - Oracle driver for go that uses database/sql. Stars:`607`.
- [godror](https://github.com/godror/godror) - Oracle driver for Go, using the ODPI-C driver. Stars:`450`.
- [KSQL](https://github.com/VinGarcia/ksql) - A Simple and Powerful Golang SQL Library Stars:`234`.
- [firebirdsql](https://github.com/nakagami/firebirdsql) - Firebird RDBMS SQL driver for Go. Stars:`201`.
- [surrealdb.go](https://github.com/surrealdb/surrealdb.go) - SurrealDB Driver for Go. Stars:`177`.
- [Sqinn-Go](https://github.com/cvilsmeier/sqinn-go) - SQLite with pure Go. Stars:`139`.
- [go-adodb](https://github.com/mattn/go-adodb) - Microsoft ActiveX Object DataBase driver for go that uses database/sql. Stars:`133`.
- [avatica](https://github.com/apache/calcite-avatica-go) - Apache Avatica/Phoenix SQL driver for database/sql. Stars:`109`.
- [gofreetds](https://github.com/minus5/gofreetds) - Microsoft MSSQL driver. Go wrapper over [FreeTDS](https://www.freetds.org). Stars:`109`.
- [ydb-go-sdk](https://github.com/ydb-platform/ydb-go-sdk) - native and database/sql driver YDB (Yandex Database) Stars:`99`.
- [bgc](https://github.com/viant/bgc) - Datastore Connectivity for BigQuery for go. Stars:`20`.
- [pig](https://github.com/alexeyco/pig) - Simple [pgx](https://github.com/jackc/pgx) wrapper to execute and [scan](https://github.com/georgysavva/scany) query results easily. Stars:`13`.

### NoSQL Database Drivers

- [redis](https://github.com/redis/go-redis) - Redis client for Golang. Stars:`17.7K`.
- [redigo](https://github.com/gomodule/redigo) - Redigo is a Go client for the Redis database. Stars:`9.6K`.
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver for the Go language. Stars:`7.5K`.
- [mgo](https://github.com/globalsign/mgo) - (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. Stars:`2.0K`.
- [gorethink](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB. Stars:`1.6K`.
- [qmgo](https://github.com/qiniu/qmgo) - The MongoDB driver for Go. It‘s based on official MongoDB driver but easier to use like Mgo. Stars:`1.2K`.
- [mgm](https://github.com/kamva/mgm) - MongoDB model-based ODM for Go (based on official MongoDB driver). Stars:`677`.
- [redeo](https://github.com/bsm/redeo) - Redis-protocol compatible TCP servers/services. Stars:`426`.
- [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) - Aerospike client in Go language. Stars:`411`.
- [neoism](https://github.com/jmcvetta/neoism) - Neo4j client for Golang. Stars:`389`.
- [gocb](https://github.com/couchbase/gocb) - Official Couchbase Go SDK. Stars:`344`.
- [go-rejson](https://github.com/nitishm/go-rejson) - Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. Stars:`321`.
- [go-couchbase](https://github.com/couchbase/go-couchbase) - Couchbase client in Go. Stars:`315`.
- [Kivik](https://github.com/go-kivik/kivik) - Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. Stars:`276`.
- [godis](https://github.com/piaohao/godis) - redis client implement by golang, inspired by jedis. Stars:`109`.
- [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang. Stars:`76`.
- [arangolite](https://github.com/solher/arangolite) - Lightweight golang driver for ArangoDB. Stars:`73`.
- [go-pilosa](https://github.com/pilosa/go-pilosa) - Go client library for Pilosa. Stars:`58`.
- [forestdb](https://github.com/couchbase/goforestdb) - Go bindings for ForestDB. Stars:`33`.
- [goriak](https://github.com/zegl/goriak) - Go language driver for Riak KV. Stars:`30`.
- [neo4j](https://github.com/cihangir/neo4j) - Neo4j Rest API Bindings for Golang. Stars:`27`.
- [xredis](https://github.com/shomali11/xredis) - Typesafe, customizable, clean & easy to use Redis client. Stars:`19`.
- [gocosmos](https://github.com/btnguyen2k/gocosmos) - REST client and standard `database/sql` driver for Azure Cosmos DB. Stars:`17`.
- [godscache](https://github.com/defcronyke/godscache) - A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. Stars:`11`.
- [asc](https://github.com/viant/asc) - Datastore Connectivity for Aerospike for go. Stars:`9`.
- [gocql](https://gocql.github.io) - Go language driver for Apache Cassandra.
- [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language.
- [rueidis](http://github.com/rueian/rueidis) - Fast Redis RESP3 client with auto pipelining and server-assisted client side caching.

### Search and Analytic Databases

- [elastic](https://github.com/olivere/elastic) - Elasticsearch client for Go. Stars:`7.2K`.
- [go-elasticsearch](https://github.com/elastic/go-elasticsearch) - Official Elasticsearch client for Go. Stars:`5.1K`.
- [elasticsql](https://github.com/cch123/elasticsql) - Convert sql to elasticsearch dsl in Go. Stars:`1.1K`.
- [elastigo](https://github.com/mattbaird/elastigo) - Elasticsearch client library. Stars:`951`.
- [skizze](https://github.com/seiflotfy/skizze) - probabilistic data-structures service and storage. Stars:`88`.
- [goes](https://github.com/OwnLocal/goes) - Library to interact with Elasticsearch. Stars:`29`.

**[⬆ back to top](#contents)**

## Date and Time

_Libraries for working with dates and times._

- [now](https://github.com/jinzhu/now) - Now is a time toolkit for golang. Stars:`4.2K`.
- [carbon](https://github.com/golang-module/carbon) - A simple, semantic and developer-friendly golang package for datetime. Stars:`3.4K`.
- [dateparse](https://github.com/araddon/dateparse) - Parse date's without knowing format in advance. Stars:`1.9K`.
- [carbon](https://github.com/uniplaces/carbon) - Simple Time extension with a lot of util methods, ported from PHP Carbon library. Stars:`762`.
- [durafmt](https://github.com/hako/durafmt) - Time duration formatting library for Go. Stars:`465`.
- [timeutil](https://github.com/leekchan/timeutil) - Useful extensions (Timedelta, Strftime, ...) to the golang's time package. Stars:`193`.
- [gostradamus](https://github.com/bykof/gostradamus) - A Go package for working with dates. Stars:`181`.
- [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) - The implementation of the Persian (Solar Hijri) Calendar in Go (golang). Stars:`163`.
- [iso8601](https://github.com/relvacode/iso8601) - Efficiently parse ISO8601 date-times without regex. Stars:`126`.
- [date](https://github.com/rickb777/date) - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. Stars:`114`.
- [timespan](https://github.com/SaidinWoT/timespan) - For interacting with intervals of time, defined as a start time and a duration. Stars:`82`.
- [go-str2duration](https://github.com/xhit/go-str2duration) - Convert string to duration. Support time.Duration returned string and more. Stars:`77`.
- [go-sunrise](https://github.com/nathan-osman/go-sunrise) - Calculate the sunrise and sunset times for a given location. Stars:`74`.
- [feiertage](https://github.com/wlbr/feiertage) - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... Stars:`43`.
- [kair](https://github.com/GuilhermeCaruso/kair) - Date and Time - Golang Formatting Library. Stars:`25`.
- [cronrange](https://github.com/1set/cronrange) - Parses Cron-style time range expressions, checks if the given time is within any ranges. Stars:`18`.
- [NullTime](https://github.com/kirillDanshin/nulltime) - Nullable `time.Time`. Stars:`14`.
- [strftime](https://github.com/awoodbeck/strftime) - C99-compatible strftime formatter. Stars:`12`.
- [tuesday](https://github.com/osteele/tuesday) - Ruby-compatible Strftime function. Stars:`12`.
- [go-anytime](https://github.com/ijt/go-anytime) - Parse dates/times like "next dec 22nd at 3pm" and ranges like "from today until next thursday" without knowing the format in advance. Stars:`10`.
- [go-week](https://github.com/stoewer/go-week) - An efficient package to work with ISO8601 week dates. Stars:`8`.

**[⬆ back to top](#contents)**

## Distributed Systems

_Packages that help with building Distributed Systems._

- [go-kit](https://github.com/go-kit/kit) - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. Stars:`25.3K`.
- [go-zero](https://github.com/tal-tech/go-zero) - A web and rpc framework. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity. Stars:`25.1K`.
- [Kratos](https://github.com/go-kratos/kratos) - A modular-designed and easy-to-use microservices framework in Go. Stars:`21.1K`.
- [go-micro](https://github.com/micro/go-micro) - A distributed systems development framework. Stars:`20.7K`.
- [grpc-go](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC. Stars:`18.6K`.
- [NATS](https://github.com/nats-io/gnatsd) - Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. Stars:`13.3K`.
- [micro](https://github.com/micro/micro) - A distributed systems runtime for the cloud and beyond. Stars:`11.8K`.
- [rpcx](https://github.com/smallnest/rpcx) - Distributed pluggable RPC service framework like alibaba Dubbo. Stars:`7.7K`.
- [raft](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol, by HashiCorp. Stars:`7.3K`.
- [Kitex](https://github.com/cloudwego/kitex) - A high-performance and strong-extensibility Golang RPC framework that helps developers build microservices. If the performance and extensibility are the main concerns when you develop microservices, Kitex can be a good choice. Stars:`6.1K`.
- [lura](https://github.com/luraproject/lura) - Ultra performant API Gateway framework with middlewares. Stars:`5.7K`.
- [torrent](https://github.com/anacrolix/torrent) - BitTorrent client package. Stars:`4.9K`.
- [dragonboat](https://github.com/lni/dragonboat) - A feature complete and high performance multi-group Raft library in Go. Stars:`4.7K`.
- [emitter-io](https://github.com/emitter-io/emitter) - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. Stars:`3.6K`.
- [gleam](https://github.com/chrislusf/gleam) - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. Stars:`3.2K`.
- [glow](https://github.com/chrislusf/glow) - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. Stars:`3.1K`.
- [liftbridge](https://github.com/liftbridge-io/liftbridge) - Lightweight, fault-tolerant message streams for NATS. Stars:`2.5K`.
- [Dragonfly](https://github.com/dragonflyoss/Dragonfly2) - Provide efficient, stable and secure file distribution and image acceleration based on p2p technology to be the best practice and standard solution in cloud native architectures. Stars:`1.5K`.
- [hprose](https://github.com/hprose/hprose-golang) - Very newbility RPC Library, support 25+ languages now. Stars:`1.2K`.
- [go-doudou](https://github.com/unionj-cloud/go-doudou) - A gossip protocol and OpenAPI 3.0 spec based decentralized microservice framework. Built-in go-doudou cli focusing on low-code and rapid dev can power up your productivity. Stars:`1.2K`.
- [redis-lock](https://github.com/bsm/redislock) - Simplified distributed locking implementation using Redis. Stars:`1.1K`.
- [rain](https://github.com/cenkalti/rain) - BitTorrent client and library. Stars:`855`.
- [arpc](https://github.com/lesismal/arpc) - More effective network communication, support two-way-calling, notify, broadcast. Stars:`787`.
- [ringpop-go](https://github.com/uber/ringpop-go) - Scalable, fault-tolerant application-layer sharding for Go applications. Stars:`786`.
- [go-health](https://github.com/InVisionApp/go-health) - Library for enabling asynchronous dependency health checks in your service. Stars:`701`.
- [gorpc](https://github.com/valyala/gorpc) - Simple, fast and scalable RPC library for high load. Stars:`680`.
- [consistent](https://github.com/buraksezer/consistent) - Consistent hashing with bounded loads. Stars:`608`.
- [go-sundheit](https://github.com/AppsFlyer/go-sundheit) - A library built to provide support for defining async service health checks for golang services. Stars:`511`.
- [digota](https://github.com/digota/digota) - grpc ecommerce microservice. Stars:`479`.
- [sleuth](https://github.com/ursiform/sleuth) - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). Stars:`367`.
- [go-jump](https://github.com/dgryski/go-jump) - Port of Google's "Jump" Consistent Hash function. Stars:`367`.
- [Temporal](https://github.com/temporalio/sdk-go) - Durable execution system for making code fault-tolerant and simple. Stars:`350`.
- [raft](https://github.com/etcd-io/raft) - Go implementation of the Raft consensus protocol, by CoreOS. Stars:`297`.
- [jsonrpc](https://github.com/ybbus/jsonrpc) - JSON-RPC 2.0 HTTP client implementation. Stars:`280`.
- [dht](https://github.com/anacrolix/dht) - BitTorrent Kademlia DHT implementation. Stars:`268`.
- [jsonrpc](https://github.com/osamingo/jsonrpc) - The jsonrpc package helps implement of JSON-RPC 2.0. Stars:`180`.
- [outboxer](https://github.com/italolelis/outboxer) - Outboxer is a go library that implements the outbox pattern. Stars:`137`.
- [doublejump](https://github.com/edwingeng/doublejump) - A revamped Google's jump consistent hash. Stars:`88`.
- [Semaphore](https://github.com/jexia/semaphore) - A straightforward (micro) service orchestrator. Stars:`84`.
- [celeriac](https://github.com/svcavallar/celeriac.v1) - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. Stars:`71`.
- [go-mysql-lock](https://github.com/sanketplus/go-mysql-lock) - MySQL based distributed lock. Stars:`55`.
- [flowgraph](https://github.com/vectaport/flowgraph) - flow-based programming package. Stars:`52`.
- [go-pdu](https://github.com/pdupub/go-pdu) - A decentralized identity-based social network. Stars:`47`.
- [drmaa](https://github.com/dgruber/drmaa) - Job submission library for cluster schedulers based on the DRMAA standard. Stars:`44`.
- [gmsec](https://github.com/gmsec/micro) - A Go distributed systems development framework. Stars:`22`.
- [consistenthash](https://github.com/mbrostami/consistenthash) - Consistent hashing with configurable replicas. Stars:`19`.
- [dynatomic](https://github.com/tylfin/dynatomic) - A library for using DynamoDB as an atomic counter. Stars:`15`.
- [failured](https://github.com/andy2046/failured) - adaptive accrual failure detector for distributed systems. Stars:`10`.
- [pjrpc](https://gitlab.com/pjrpc/pjrpc) - Golang JSON-RPC Server-Client with Protobuf spec.
- [resgate](https://resgate.io/) - Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
- [dot](https://github.com/dotchain/dot/) - distributed sync using operational transformation/OT.
- [dynamolock](https://cirello.io/dynamolock) - DynamoDB-backed distributed locking implementation.
- [pglock](https://cirello.io/pglock) - PostgreSQL-backed distributed locking implementation.

**[⬆ back to top](#contents)**

## Dynamic DNS

_Tools for updating dynamic DNS records._

- [GoDNS](https://github.com/timothyye/godns) - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. Stars:`1.3K`.
- [DDNS](https://github.com/skibish/ddns) - Personal DDNS client with Digital Ocean Networking DNS as backend. Stars:`237`.
- [dyndns](https://gitlab.com/alcastle/dyndns) - Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.

**[⬆ back to top](#contents)**

## Email

_Libraries and tools that implement email creation and sending._

- [MailHog](https://github.com/mailhog/MailHog) - Email and SMTP testing with web and API interface. Stars:`12.4K`.
- [Mailpit](https://github.com/axllent/mailpit) - Email and SMTP testing tool for developers. Stars:`2.7K`.
- [hermes](https://github.com/matcornic/hermes) - Golang package that generates clean, responsive HTML e-mails. Stars:`2.7K`.
- [email](https://github.com/jordan-wright/email) - A robust and flexible email library for Go. Stars:`2.4K`.
- [go-imap](https://github.com/emersion/go-imap) - IMAP library for clients and servers. Stars:`1.8K`.
- [email-verifier](https://github.com/AfterShip/email-verifier) - A Go library for email verification without sending any emails. Stars:`923`.
- [SendGrid](https://github.com/sendgrid/sendgrid-go) - SendGrid's Go library for sending email. Stars:`914`.
- [mailgun-go](https://github.com/mailgun/mailgun-go) - Go library for sending mail with the Mailgun API. Stars:`656`.
- [go-simple-mail](https://github.com/xhit/go-simple-mail) - Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. Stars:`499`.
- [go-message](https://github.com/emersion/go-message) - Streaming library for the Internet Message Format and mail messages. Stars:`322`.
- [go-mail](https://github.com/wneessen/go-mail) - A simple Go library for sending mails in Go. Stars:`309`.
- [douceur](https://github.com/aymerick/douceur) - CSS inliner for your HTML emails. Stars:`230`.
- [Hectane](https://github.com/hectane/hectane) - Lightweight SMTP client providing an HTTP API. Stars:`220`.
- [mailchain](https://github.com/mailchain/mailchain) - Send encrypted emails to blockchain addresses written in Go. Stars:`139`.
- [go-premailer](https://github.com/vanng822/go-premailer) - Inline styling for HTML mail in Go. Stars:`113`.
- [go-dkim](https://github.com/toorop/go-dkim) - DKIM library, to sign & verify email. Stars:`90`.
- [smtpmock](https://github.com/mocktools/go-smtp-mock) - Lightweight configurable multithreaded fake SMTP server. Mimic any SMTP behaviour for your test environment. Stars:`83`.
- [smtp](https://github.com/mailhog/smtp) - SMTP server protocol state machine. Stars:`75`.
- [go-email-normalizer](https://github.com/dimuska139/go-email-normalizer) - Golang library for providing a canonical representation of email address. Stars:`57`.
- [truemail-go](https://github.com/truemail-rb/truemail-go) - Configurable Golang email validator/verifier. Verify email via Regex, DNS, SMTP and even more. Stars:`52`.
- [go-email-validator](https://github.com/go-email-validator/go-email-validator) - Modular email validator for syntax, disposable, smtp, etc... checking. Stars:`51`.
- [mailx](https://github.com/valord577/mailx) - Mailx is a library that makes it easier to send email via SMTP. It is an enhancement of the golang standard library `net/smtp`. Stars:`7`.
- [chasquid](https://blitiri.com.ar/p/chasquid) - SMTP server written in Go.

**[⬆ back to top](#contents)**

## Embeddable Scripting Languages

_Embedding other languages inside your go code._

- [gopher-lua](https://github.com/yuin/gopher-lua) - Lua 5.1 VM and compiler written in Go. Stars:`5.6K`.
- [goja](https://github.com/dop251/goja) - ECMAScript 5.1(+) implementation in Go. Stars:`4.1K`.
- [expr](https://github.com/antonmedv/expr) - Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing. Stars:`3.9K`.
- [tengo](https://github.com/d5/tengo) - Bytecode compiled script language for Go. Stars:`3.2K`.
- [go-lua](https://github.com/Shopify/go-lua) - Port of the Lua 5.2 VM to pure Go. Stars:`2.7K`.
- [starlark-go](https://github.com/google/starlark-go) - Go implementation of Starlark: Python-like language with deterministic evaluation and hermetic execution. Stars:`2.0K`.
- [cel-go](https://github.com/google/cel-go) - Fast, portable, non-Turing complete expression evaluation with gradual typing. Stars:`1.7K`.
- [go-python](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API. Stars:`1.5K`.
- [anko](https://github.com/mattn/anko) - Scriptable interpreter written in Go. Stars:`1.4K`.
- [metacall](https://github.com/metacall/core) - Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, WebAssembly, Java, Cobol and more. Stars:`1.3K`.
- [go-php](https://github.com/deuill/go-php) - PHP bindings for Go. Stars:`900`.
- [go-duktape](https://github.com/olebedev/go-duktape) - Duktape JavaScript engine bindings for Go. Stars:`779`.
- [gval](https://github.com/PaesslerAG/gval) - A highly customizable expression language written in Go. Stars:`639`.
- [golua](https://github.com/aarzilli/golua) - Go bindings for Lua C API. Stars:`617`.
- [gisp](https://github.com/jcla1/gisp) - Simple LISP in Go. Stars:`502`.
- [prolog](https://github.com/ichiban/prolog) - Embeddable Prolog. Stars:`497`.
- [gentee](https://github.com/gentee/gentee) - Embeddable scripting programming language. Stars:`112`.
- [binder](https://github.com/alexeyco/binder) - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). Stars:`66`.
- [purl](https://github.com/ian-kent/purl) - Perl 5.18.2 embedded in Go. Stars:`38`.
- [ecal](https://github.com/krotik/ecal) - A simple embeddable scripting language which supports concurrent event processing. Stars:`34`.
- [ngaro](https://github.com/db47h/ngaro) - Embeddable Ngaro VM implementation enabling scripting in Retro. Stars:`25`.

**[⬆ back to top](#contents)**

## Error Handling

_Libraries for handling errors._

- [errors](https://github.com/pkg/errors) - Package that provides simple error handling primitives. Stars:`8.1K`.
- [go-multierror](https://github.com/hashicorp/go-multierror) - Go (golang) package for representing a list of errors as a single error. Stars:`2.0K`.
- [errors](https://github.com/cockroachdb/errors) - Go error library with error portability over the network. Stars:`1.7K`.
- [eris](https://github.com/rotisserie/eris) - A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. Stars:`1.3K`.
- [errorx](https://github.com/joomcode/errorx) - A feature rich error package with stack traces, composition of errors and more. Stars:`993`.
- [tracerr](https://github.com/ztrue/tracerr) - Golang errors with stack trace and source fragments. Stars:`811`.
- [errlog](https://github.com/snwfdhmp/errlog) - Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. Stars:`452`.
- [emperror](https://github.com/emperror/emperror) - Error handling tools and best practices for Go libraries and applications. Stars:`311`.
- [errors](https://github.com/emperror/errors) - Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. Stars:`180`.
- [Fault](https://github.com/Southclaws/fault) - An ergonomic mechanism for wrapping errors in order to facilitate structured metadata and context for error values. Stars:`133`.
- [oops](https://github.com/samber/oops) - Error handling with context, stack trace and source fragments. Stars:`55`.
- [errors](https://github.com/bnkamalesh/errors) - Drop-in replacement for builtin Go errors. This is a minimal error handling package with custom error types, user friendly messages, Unwrap & Is. With very easy to use and straightforward helper functions. Stars:`52`.
- [exception](https://github.com/rbrahul/exception) - A simple utility package for exception handling with try-catch in Golang. Stars:`24`.
- [Falcon](https://github.com/SonicRoshan/falcon) - A Simple Yet Highly Powerful Package For Error Handling. Stars:`9`.
- [errors](https://github.com/PumpkinSeed/errors) - The most simple error wrapper with awesome performance and minimal memory overhead. Stars:`8`.
- [errors](https://github.com/neuronlabs/errors) - Simple golang error handling with classification primitives. Stars:`5`.

**[⬆ back to top](#contents)**

## File Handling

_Libraries for handling files and file systems._

- [afero](https://github.com/spf13/afero) - FileSystem Abstraction System for Go. Stars:`5.3K`.
- [pdfcpu](https://github.com/pdfcpu/pdfcpu) - PDF processor. Stars:`5.2K`.
- [gdu](https://github.com/dundee/gdu) - Disk usage analyzer with console interface. Stars:`2.7K`.
- [notify](https://github.com/rjeczalik/notify) - File system event notification library with simple API, similar to os/signal. Stars:`835`.
- [copy](https://github.com/otiai10/copy) - Copy directory recursively. Stars:`580`.
- [gofs](https://github.com/no-src/gofs) - A cross-platform real-time file synchronization tool out of the box. Stars:`311`.
- [afs](https://github.com/viant/afs) - Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. Stars:`255`.
- [vfs](https://github.com/C2FO/vfs) - A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. Stars:`248`.
- [bigfile](https://github.com/bigfile/bigfile) - A file transfer system, support to manage files with http api, rpc call and ftp client. Stars:`241`.
- [go-exiftool](https://github.com/barasher/go-exiftool) - Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). Stars:`177`.
- [go-csv-tag](https://github.com/artonge/go-csv-tag) - Load csv file using tag. Stars:`108`.
- [skywalker](https://github.com/dixonwille/skywalker) - Package to allow one to concurrently go through a filesystem with ease. Stars:`90`.
- [checksum](https://github.com/codingsince1985/checksum) - Compute message digest, like MD5, SHA256, SHA1, CRC or BLAKE2s, for large files. Stars:`88`.
- [parquet](https://github.com/parsyl/parquet) - Read and write [parquet](https://parquet.apache.org) files. Stars:`83`.
- [opc](https://github.com/qmuntal/opc) - Load Open Packaging Conventions (OPC) files for Go. Stars:`74`.
- [tarfs](https://github.com/posener/tarfs) - Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. Stars:`56`.
- [baraka](https://github.com/xis/baraka) - A library to process http file uploads easily. Stars:`50`.
- [go-gtfs](https://github.com/artonge/go-gtfs) - Load gtfs files in go. Stars:`37`.
- [flop](https://github.com/homedepot/flop) - File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html). Stars:`34`.
- [gut/yos](https://github.com/1set/gut) - Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links. Stars:`26`.
- [todotxt](https://github.com/1set/todotxt) - Go library for Gina Trapani's [_todo.txt_](http://todotxt.org/) files, supports parsing and manipulating of task lists in the [_todo.txt_ format](https://github.com/todotxt/todo.txt). Stars:`19`.
- [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) - Copy files for humans. Stars:`19`.
- [higgs](https://github.com/dastoori/higgs) - A tiny cross-platform Go library to hide/unhide files and directories. Stars:`17`.
- [pathtype](https://github.com/jonchun/pathtype) - Treat paths as their own type instead of using strings. Stars:`12`.
- [stl](https://gitlab.com/russoj88/stl) - Modules to read and write STL (stereolithography) files. Concurrent algorithm for reading.

**[⬆ back to top](#contents)**

## Financial

_Packages for accounting and finance._

- [decimal](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers. Stars:`5.2K`.
- [ticker](https://github.com/achannarasappa/ticker) - Terminal stock watcher and stock position tracker. Stars:`4.6K`.
- [go-money](https://github.com/rhymond/go-money) - Implementation of Fowler's Money pattern. Stars:`1.3K`.
- [bbgo](https://github.com/c9s/bbgo) - A crypto trading bot framework written in Go. Including common crypto exchange API, standard indicators, back-testing and many built-in strategies. Stars:`929`.
- [accounting](https://github.com/leekchan/accounting) - money and currency formatting for golang. Stars:`821`.
- [techan](https://github.com/sdcoffey/techan) - Technical analysis library with advanced market analysis and trading strategies. Stars:`736`.
- [currency](https://github.com/bojanz/currency) - Handles currency amounts, provides currency information and formatting. Stars:`411`.
- [ach](https://github.com/moov-io/ach) - A reader, writer, and validator for Automated Clearing House (ACH) files. Stars:`378`.
- [orderbook](https://github.com/i25959341/orderbook) - Matching Engine for Limit Order Book in Golang. Stars:`346`.
- [go-finance](https://github.com/alpeb/go-finance) - Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. Stars:`150`.
- [transaction](https://github.com/claygod/transaction) - Embedded transactional database of accounts, running in multithreaded mode. Stars:`123`.
- [sleet](https://github.com/BoltApp/sleet) - One unified interface for multiple Payment Service Providers (PsP) to process online payment. Stars:`122`.
- [ofxgo](https://github.com/aclindsa/ofxgo) - Query OFX servers and/or parse the responses (with example command-line client). Stars:`117`.
- [vat](https://github.com/dannyvankooten/vat) - VAT number validation & EU VAT rates. Stars:`101`.
- [go-finnhub](https://github.com/m1/go-finnhub) - Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges. Stars:`79`.
- [currency](https://github.com/bnkamalesh/currency) - High performant & accurate currency computation package. Stars:`52`.
- [fpdecimal](https://github.com/nikolaydubina/fpdecimal) - Fast and precise serialization and arithmetic for small fixed-point decimals Stars:`23`.
- [payme](https://github.com/jovandeginste/payme) - QR code generator (ASCII & PNG) for SEPA payments. Stars:`22`.
- [fpmoney](https://github.com/nikolaydubina/fpmoney) - Fast and simple ISO4217 fixed-point decimal money. Stars:`20`.
- [go-finance](https://github.com/pieterclaerhout/go-finance) - Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers. Stars:`13`.

**[⬆ back to top](#contents)**

## Forms

_Libraries for working with forms._

- [nosurf](https://github.com/justinas/nosurf) - CSRF protection middleware for Go. Stars:`1.4K`.
- [gorilla/csrf](https://github.com/gorilla/csrf) - CSRF protection for Go web applications & services. Stars:`900`.
- [binding](https://github.com/mholt/binding) - Binds form and JSON data from net/http Request to struct. Stars:`795`.
- [form](https://github.com/go-playground/form) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. Stars:`616`.
- [conform](https://github.com/leebenson/conform) - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. Stars:`308`.
- [httpin](https://github.com/ggicci/httpin) - Decode an HTTP request into a custom struct, including querystring, forms, HTTP headers, etc. Stars:`190`.
- [formam](https://github.com/monoculum/formam) - decode form's values into a struct. Stars:`181`.
- [forms](https://github.com/albrow/forms) - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. Stars:`133`.
- [qs](https://github.com/sonh/qs) - Go module for encoding structs into URL query parameters. Stars:`67`.
- [bind](https://github.com/robfig/bind) - Bind form data to any Go values. Stars:`28`.
- [queryparam](https://github.com/tomwright/queryparam) - Decode `url.Values` into usable struct values of standard or custom types. Stars:`17`.
- [gbind](https://github.com/bdjimmy/gbind) - Bind data to any Go value. Can use built-in and custom expression binding capabilities; supports data validation Stars:`8`.

**[⬆ back to top](#contents)**

## Functional

_Packages to support functional programming in Go._

- [mo](https://github.com/samber/mo) - Monads and popular FP abstractions, based on Go 1.18+ Generics (Option, Result, Either...). Stars:`1.7K`.
- [go-underscore](https://github.com/tobyhede/go-underscore) - Useful collection of helpfully functional Go collection utilities. Stars:`1.3K`.
- [fpGo](https://github.com/TeaEntityLab/fpGo) - Monad, Functional Programming features for Golang. Stars:`323`.
- [fp-go](https://github.com/repeale/fp-go) - Collection of Functional Programming helpers powered by Golang 1.18+ generics. Stars:`241`.
- [fuego](https://github.com/seborama/fuego) - Functional Experiment in Go. Stars:`141`.
- [gofp](https://github.com/rbrahul/gofp) - A lodash like powerful utility library for Golang. Stars:`139`.
- [underscore](https://github.com/rjNemo/underscore) - Functional programming helpers for Go 1.18 and beyond. Stars:`89`.
- [valor](https://github.com/phelmkamp/valor) - Generic option and result types that optionally contain a value. Stars:`14`.

**[⬆ back to top](#contents)**

## Game Development

_Awesome game development libraries._

- [Ebitengine](https://github.com/hajimehoshi/ebiten) - dead simple 2D game engine in Go. Stars:`8.5K`.
- [Leaf](https://github.com/name5566/leaf) - Lightweight game server framework. Stars:`4.9K`.
- [Pixel](https://github.com/faiface/pixel) - Hand-crafted 2D game library in Go. Stars:`4.3K`.
- [nano](https://github.com/lonng/nano) - Lightweight, facility, high performance golang based game server framework. Stars:`2.4K`.
- [g3n](https://github.com/g3n/engine) - Go 3D Game Engine. Stars:`2.4K`.
- [goworld](https://github.com/xiaonanln/goworld) - Scalable game server engine, featuring space-entity framework and hot-swapping. Stars:`2.4K`.
- [go-sdl2](https://github.com/veandco/go-sdl2) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). Stars:`2.0K`.
- [Pitaya](https://github.com/topfreegames/pitaya) - Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. Stars:`1.9K`.
- [engo](https://github.com/EngoEngine/engo) - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. Stars:`1.6K`.
- [Oak](https://github.com/oakmound/oak) - Pure Go game engine. Stars:`1.4K`.
- [termloop](https://github.com/JoelOtter/termloop) - Terminal-based game engine for Go, built on top of Termbox. Stars:`1.4K`.
- [gonet](https://github.com/xtaci/gonet) - Game server skeleton implemented with golang. Stars:`1.2K`.
- [raylib-go](https://github.com/gen2brain/raylib-go) - Go bindings for [raylib](https://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. Stars:`1.0K`.
- [Azul3D](https://github.com/azul3d/engine) - 3D game engine written in Go. Stars:`585`.
- [go-astar](https://github.com/beefsack/go-astar) - Go implementation of the A\* path finding algorithm. Stars:`558`.
- [Harfang3D](https://github.com/harfang3d/harfang3d) - 3D engine for the Go language, works on Windows and Linux ([Harfang on Go.dev](https://github.com/harfang3d/harfang-go)). Stars:`354`.
- [go3d](https://github.com/ungerik/go3d) - Performance oriented 2D/3D math package for Go. Stars:`285`.
- [tile](https://github.com/kelindar/tile) - Data-oriented and cache-friendly 2D Grid library (TileMap), includes pathfinding, observers and import/export. Stars:`108`.
- [prototype](https://github.com/gonutz/prototype) - Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API. Stars:`81`.
- [fantasyname](https://github.com/s0rg/fantasyname) - Fantasy names generator. Stars:`19`.

**[⬆ back to top](#contents)**

## Generators

_Tools that generate Go code._

- [go-linq](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go. Stars:`3.3K`.
- [jennifer](https://github.com/dave/jennifer) - Generate arbitrary Go code without templates. Stars:`3.0K`.
- [goderive](https://github.com/awalterschulze/goderive) - Derives functions from input types. Stars:`1.1K`.
- [GoWrap](https://github.com/hexdigest/gowrap) - Generate decorators for Go interfaces using simple templates. Stars:`785`.
- [go-enum](https://github.com/abice/go-enum) - Code generation for enums from code comments. Stars:`532`.
- [interfaces](https://github.com/rjeczalik/interfaces) - Command line tool for generating interface definitions. Stars:`398`.
- [goverter](https://github.com/jmattheis/goverter) - Generate converters by defining an interface. Stars:`287`.
- [copygen](https://github.com/switchupcb/copygen) - Generate type-to-type and type-based code without reflection. Stars:`272`.
- [gotype](https://github.com/wzshiming/gotype) - Golang source code parsing, usage like reflect package. Stars:`54`.
- [generis](https://github.com/senselogic/GENERIS) - Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. Stars:`42`.
- [typeregistry](https://github.com/xiaoxin01/typeregistry) - A library to create type dynamically. Stars:`20`.
- [convergen](https://github.com/reedom/convergen) - Feature rich type-to-type copy code generator. Stars:`7`.

**[⬆ back to top](#contents)**

## Geographic

_Geographic tools and servers_

- [Tile38](https://github.com/tidwall/tile38) - Geolocation DB with spatial index and realtime geofencing. Stars:`8.6K`.
- [S2 geometry](https://github.com/golang/geo) - S2 geometry library in Go. Stars:`1.6K`.
- [mbtileserver](https://github.com/consbio/mbtileserver) - A simple Go-based server for map tiles stored in mbtiles format. Stars:`512`.
- [osm](https://github.com/paulmach/osm) - Library for reading, writing and working with OpenStreetMap data and APIs. Stars:`280`.
- [H3](https://github.com/uber/h3-go) - Go bindings for H3, a hierarchical hexagonal geospatial indexing system. Stars:`234`.
- [godal](https://github.com/airbusgeo/godal) - Go wrapper for GDAL. Stars:`100`.
- [WGS84](https://github.com/wroge/wgs84) - Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). Stars:`99`.
- [simplefeatures](https://github.com/peterstace/simplefeatures) - simplesfeatures is a 2D geometry library that provides Go types that model geometries, as well as algorithms that operate on them. Stars:`89`.
- [geoserver](https://github.com/hishamkaram/geoserver) - geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. Stars:`83`.
- [gismanager](https://github.com/hishamkaram/gismanager) - Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. Stars:`50`.
- [pbf](https://github.com/maguro/pbf) - OpenStreetMap PBF golang encoder/decoder. Stars:`45`.
- [S2 geojson](https://github.com/pantrif/s2-geojson) - Convert geojson to s2 cells & demonstrating some S2 geometry features on map. Stars:`25`.
- [H3 GeoJSON](https://github.com/mmadfox/go-geojson2h3) - Conversion utilities between H3 indexes and GeoJSON. Stars:`3`.
- [Web-Mercator-Projection](https://github.com/jorelosorio/web-mercator-projection) A project to easily use and convert LonLat, Point and Tile to display info, markers, etc, in a map using the Web Mercator Projection. Stars:`3`.
- [H3GeoDist](https://github.com/mmadfox/go-h3geo-dist) - Distribution of Uber H3geo cells by virtual nodes. Stars:`1`.

**[⬆ back to top](#contents)**

## Go Compilers

_Tools for compiling Go to other languages._

- [gopherjs](https://github.com/gopherjs/gopherjs) - Compiler from Go to JavaScript. Stars:`12.0K`.
- [tardisgo](https://github.com/tardisgo/tardisgo) - Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. Stars:`426`.
- [c4go](https://github.com/Konstantin8105/c4go) - Transpile C code to Go code. Stars:`337`.
- [c2go](https://github.com/goplus/c2go) - Convert C code to Go code. Stars:`268`.
- [esp32](https://github.com/andygeiss/esp32-transpiler) - Transpile Go into Arduino code. Stars:`70`.
- [f4go](https://github.com/Konstantin8105/f4go) - Transpile FORTRAN 77 code to Go code. Stars:`36`.

**[⬆ back to top](#contents)**

## Goroutines

_Tools for managing and working with Goroutines._

- [ants](https://github.com/panjf2000/ants) - A high-performance and low-cost goroutine pool in Go. Stars:`10.8K`.
- [conc](https://github.com/sourcegraph/conc) - `conc` is your toolbelt for structured concurrency in go, making common tasks easier and safer. Stars:`7.1K`.
- [tunny](https://github.com/Jeffail/tunny) - Goroutine pool for golang. Stars:`3.7K`.
- [goworker](https://github.com/benmanns/goworker) - goworker is a Go-based background worker. Stars:`2.7K`.
- [workerpool](https://github.com/gammazero/workerpool) - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. Stars:`1.1K`.
- [pond](https://github.com/alitto/pond) - Minimalistic and High-performance goroutine worker pool written in Go. Stars:`939`.
- [grpool](https://github.com/ivpusic/grpool) - Lightweight Goroutine pool. Stars:`731`.
- [pool](https://github.com/go-playground/pool) - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. Stars:`720`.
- [gowp](https://github.com/xxjwxc/gowp) - gowp is concurrency limiting goroutine pool. Stars:`471`.
- [go-floc](https://github.com/workanator/go-floc) - Orchestrate goroutines with ease. Stars:`262`.
- [go-flow](https://github.com/kamildrazkiewicz/go-flow) - Control goroutines execution order. Stars:`214`.
- [artifex](https://github.com/borderstech/artifex) - Simple in-memory job queue for Golang using worker-based dispatching. Stars:`175`.
- [go-workers](https://github.com/catmullet/go-workers) - Easily and safely run workers for large data processing pipelines. Stars:`160`.
- [semaphore](https://github.com/marusama/semaphore) - Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). Stars:`160`.
- [neilotoole/errgroup](https://github.com/neilotoole/errgroup) - Drop-in alternative to `sync/errgroup`, limited to a pool of N worker goroutines. Stars:`148`.
- [routine](https://github.com/timandy/routine) - `routine` is a `ThreadLocal` for go library. It encapsulates and provides some easy-to-use, non-competitive, high-performance `goroutine` context access interfaces, which can help you access coroutine context information more gracefully. Stars:`144`.
- [flowmatic](https://github.com/carlmjohnson/flowmatic) - Structured concurrency made easy. Stars:`142`.
- [async](https://github.com/reugn/async) - An alternative sync library for Go (Future, Promise, Locks). Stars:`130`.
- [cyclicbarrier](https://github.com/marusama/cyclicbarrier) - CyclicBarrier for golang. Stars:`129`.
- [async](https://github.com/studiosol/async) - A safe way to execute functions asynchronously, recovering them in case of panic. Stars:`122`.
- [gollback](https://github.com/vardius/gollback) - asynchronous simple function utilities, for managing execution of closures and callbacks. Stars:`110`.
- [semaphore](https://github.com/kamilsk/semaphore) - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. Stars:`96`.
- [Hunch](https://github.com/AaronJan/Hunch) - Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. Stars:`95`.
- [threadpool](https://github.com/shettyh/threadpool) - Golang threadpool implementation. Stars:`94`.
- [worker-pool](https://github.com/vardius/worker-pool) - goworker is a Go simple async worker pool. Stars:`87`.
- [gpool](https://github.com/Sherifabdlnaby/gpool) - manages a resizeable pool of context-aware goroutines to bound concurrency. Stars:`86`.
- [go-actor](https://github.com/vladopajic/go-actor) - A tiny library for writing concurrent programs using actor model. Stars:`68`.
- [goccm](https://github.com/zenthangplus/goccm) - Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently. Stars:`66`.
- [nursery](https://github.com/arunsworld/nursery) - Structured concurrency in Go. Stars:`64`.
- [gowl](https://github.com/hamed-yousefi/gowl) - Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status. Stars:`59`.
- [routine](https://github.com/x-mod/routine) - go routine control with context, support: Main, Go, Pool and some useful Executors. Stars:`57`.
- [gohive](https://github.com/loveleshsharma/gohive) - A highly performant and easy to use Goroutine pool for Go. Stars:`44`.
- [kyoo](https://github.com/dirkaholic/kyoo) - Provides an unlimited job queue and concurrent worker pools. Stars:`42`.
- [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) - Like `sync.WaitGroup` with error handling and concurrency control. Stars:`41`.
- [parallel-fn](https://github.com/rafaeljesus/parallel-fn) - Run functions in parallel. Stars:`35`.
- [go-trylock](https://github.com/subchen/go-trylock) - TryLock support on read-write lock for Golang. Stars:`33`.
- [channelify](https://github.com/ddelizia/channelify) - Transform your function to return channels for easy and powerful parallel processing. Stars:`27`.
- [stl](https://github.com/ssgreg/stl) - Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. Stars:`27`.
- [execpool](https://github.com/hexdigest/execpool) - A pool built around exec.Cmd that spins up a given number of processes in advance and attaches stdin and stdout to them when needed. Very similar to FastCGI or Apache Prefork MPM but works for any command. Stars:`16`.
- [breaker](https://github.com/kamilsk/breaker) - Flexible mechanism to make execution flow interruptible. Stars:`15`.
- [concurrency-limiter](https://github.com/vivek-ng/concurrency-limiter) - Concurrency limiter with support for timeouts , dynamic priority and context cancellation of goroutines. Stars:`15`.
- [queue](https://github.com/AnikHasibul/queue) - Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more. Stars:`14`.
- [conexec](https://github.com/ITcathyh/conexec) - A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency. Stars:`13`.
- [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) - Manage a pool of goroutines using this lightweight library with a simple API. Stars:`12`.
- [hands](https://github.com/duanckham/hands) - A process controller used to control the execution and return strategies of multiple goroutines. Stars:`10`.
- [go-workerpool](https://github.com/zenthangplus/go-workerpool) - Inspired from Java Thread Pool, Go WorkerPool aims to control heavy Go Routines. Stars:`7`.
- [async-job](https://github.com/lab210-dev/async-job) - AsyncJob is an asynchronous queue job manager with light code, clear and speed. Stars:`6`.
- [oversight](https://cirello.io/oversight) - Oversight is a complete implementation of the Erlang supervision trees.

**[⬆ back to top](#contents)**

## GUI

_Libraries for building GUI Applications._

_Toolkits_

- [fyne](https://github.com/fyne-io/fyne) - Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android. Stars:`21.0K`.
- [webview](https://github.com/zserge/webview) - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). Stars:`11.4K`.
- [qt](https://github.com/therecipe/qt) - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). Stars:`9.9K`.
- [ui](https://github.com/andlabs/ui) - Platform-native GUI library for Go. Cross platform. Stars:`8.3K`.
- [app](https://github.com/murlokswarm/app) - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. Stars:`7.2K`.
- [walk](https://github.com/lxn/walk) - Windows application library kit for Go. Stars:`6.5K`.
- [go-astilectron](https://github.com/asticode/go-astilectron) - Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). Stars:`4.8K`.
- [go-sciter](https://github.com/sciter-sdk/go-sciter) - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. Stars:`2.5K`.
- [gotk3](https://github.com/gotk3/gotk3) - Go bindings for GTK3. Stars:`1.9K`.
- [gowd](https://github.com/dtylman/gowd) - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. Stars:`401`.
- [goradd/html5tag](https://github.com/goradd/html5tag) - Library for outputting HTML5 tags. Stars:`4`.
- [Wails](https://wails.io) - Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
- [go-gtk](https://mattn.github.io/go-gtk/) - Go bindings for GTK.
- [gio](https://gioui.org) - Gio is a library for writing cross-platform immediate mode GUI-s in Go. Gio supports all the major platforms: Linux, macOS, Windows, Android, iOS, FreeBSD, OpenBSD and WebAssembly.

_Interaction_

- [robotgo](https://github.com/go-vgo/robotgo) - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. Stars:`8.6K`.
- [systray](https://github.com/getlantern/systray) - Cross platform Go library to place an icon and menu in the notification area. Stars:`2.9K`.
- [gosx-notifier](https://github.com/deckarep/gosx-notifier) - OSX Desktop Notifications library for Go. Stars:`578`.
- [zenity](https://github.com/ncruces/zenity) - Cross-platform Go library and CLI to create simple dialogs that interact graphically with the user. Stars:`524`.
- [trayhost](https://github.com/shurcooL/trayhost) - Cross-platform Go library to place an icon in the host operating system's taskbar. Stars:`241`.
- [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) - OSX Sleep/Wake notifications in golang. Stars:`29`.
- [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) - OSX library to notify about any (pluggable) activity on your machine. Stars:`25`.
- [AppIndicator Go](https://github.com/gopherlibs/appindicator) - Go bindings for libappindicator3 C library. Stars:`3`.

**[⬆ back to top](#contents)**

## Hardware

_Libraries, tools, and tutorials for interacting with hardware._

- [arduino-cli](https://github.com/arduino/arduino-cli) - Official Arduino CLI and library. Can run standalone, or be incorporated into larger Go projects. Stars:`4.0K`.
- [go-rpio](https://github.com/stianeikeland/go-rpio) - GPIO for Go, doesn't require cgo. Stars:`2.0K`.
- [ghw](https://github.com/jaypipes/ghw) - Golang hardware discovery/inspection library. Stars:`1.5K`.
- [emgo](https://github.com/ziutek/emgo) - Go-like language for programming embedded systems (e.g. STM32 MCU). Stars:`1.0K`.
- [sysinfo](https://github.com/zcalusic/sysinfo) - A pure Go library providing Linux OS / kernel / hardware system information. Stars:`442`.
- [goroslib](https://github.com/aler9/goroslib) - Robot Operating System (ROS) library for Go. Stars:`269`.
- [go-osc](https://github.com/hypebeast/go-osc) - Open Sound Control (OSC) bindings for Go. Stars:`174`.
- [joystick](https://github.com/0xcafed00d/joystick) - a polled API to read the state of an attached joystick. Stars:`45`.

**[⬆ back to top](#contents)**

## Images

_Libraries for manipulating images._

- [gocv](https://github.com/hybridgroup/gocv) - Go package for computer vision using OpenCV 3.3+. Stars:`5.6K`.
- [imaginary](https://github.com/h2non/imaginary) - Fast and simple HTTP microservice for image resizing. Stars:`5.0K`.
- [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package. Stars:`4.8K`.
- [gg](https://github.com/fogleman/gg) - 2D rendering in pure Go. Stars:`4.0K`.
- [bild](https://github.com/anthonynsimon/bild) - Collection of image processing algorithms in pure Go. Stars:`3.8K`.
- [ln](https://github.com/fogleman/ln) - 3D line art rendering in Go. Stars:`3.2K`.
- [resize](https://github.com/nfnt/resize) - Image resizing for Go with common interpolation methods. Stars:`3.0K`.
- [bimg](https://github.com/h2non/bimg) - Small package for fast and efficient image processing using libvips. Stars:`2.3K`.
- [gowitness](https://github.com/sensepost/gowitness) - Screenshoting webpages using go and headless chrome on command line. Stars:`2.3K`.
- [pt](https://github.com/fogleman/pt) - Path tracing engine written in Go. Stars:`2.1K`.
- [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation. Stars:`2.0K`.
- [picfit](https://github.com/thoas/picfit) - An image resizing server written in Go. Stars:`1.9K`.
- [smartcrop](https://github.com/muesli/smartcrop) - Finds good crops for arbitrary images and crop sizes. Stars:`1.7K`.
- [gift](https://github.com/disintegration/gift) - Package of image processing filters. Stars:`1.7K`.
- [imagick](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API. Stars:`1.6K`.
- [canvas](https://github.com/tdewolff/canvas) - Vector graphics to PDF, SVG or rasterized image. Stars:`1.3K`.
- [geopattern](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string. Stars:`1.2K`.
- [stegify](https://github.com/DimitarPetrov/stegify) - Go tool for LSB steganography, capable of hiding any file within an image. Stars:`1.1K`.
- [govips](https://github.com/davidbyttow/govips) - A lightning fast image processing and resizing library for Go. Stars:`948`.
- [image2ascii](https://github.com/qeesung/image2ascii) - Convert image to ASCII. Stars:`784`.
- [goimagehash](https://github.com/corona10/goimagehash) - Go Perceptual image hashing package. Stars:`630`.
- [draft](https://github.com/lucasepe/draft) - Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax. Stars:`570`.
- [govatar](https://github.com/o1egl/govatar) - Library and CMD tool for generating funny avatars. Stars:`568`.
- [mort](https://github.com/aldor007/mort) - Storage and image processing server written in Go. Stars:`483`.
- [go-nude](https://github.com/koyachi/go-nude) - Nudity detection with Go. Stars:`377`.
- [steganography](https://github.com/auyer/steganography) - Pure Go Library for LSB steganography. Stars:`250`.
- [darkroom](https://github.com/gojek/darkroom) - An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. Stars:`215`.
- [rez](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD. Stars:`210`.
- [mergi](https://github.com/noelyahan/mergi) - Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). Stars:`205`.
- [gltf](https://github.com/qmuntal/gltf) - Efficient and robust glTF 2.0 reader, writer and validator. Stars:`204`.
- [img](https://github.com/hawx/img) - Selection of image manipulation tools. Stars:`150`.
- [go-webp](https://github.com/kolesa-team/go-webp) - Library for encode and decode webp pictures, using libwebp. Stars:`144`.
- [go-cairo](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library. Stars:`135`.
- [cameron](https://github.com/aofei/cameron) - An avatar generator for Go. Stars:`102`.
- [gridder](https://github.com/shomali11/gridder) - A Grid based 2D Graphics library. Stars:`64`.
- [webp-server](https://github.com/mehdipourfar/webp-server) - Simple and minimal image server capable of storing, resizing, converting and caching images. Stars:`62`.
- [color-extractor](https://github.com/marekm4/color-extractor) - Dominant color extractor with no external dependencies. Stars:`62`.
- [go-gd](https://github.com/bolknote/go-gd) - Go binding for GD library. Stars:`59`.
- [goimghdr](https://github.com/corona10/goimghdr) - The imghdr module determines the type of image contained in a file for Go. Stars:`41`.
- [tga](https://github.com/ftrvxmtrx/tga) - Package tga is a TARGA image format decoder/encoder. Stars:`33`.
- [go-webcolors](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go. Stars:`27`.
- [mpo](https://github.com/donatj/mpo) - Decoder and conversion tool for MPO 3D Photos. Stars:`11`.
- [scout](https://github.com/jonoton/scout) - Scout is a standalone open source software solution for DIY video security. Stars:`10`.

**[⬆ back to top](#contents)**

## IoT (Internet of Things)

_Libraries for programming devices of the IoT._

- [flogo](https://github.com/tibcosoftware/flogo) - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. Stars:`2.3K`.
- [mainflux](https://github.com/Mainflux/mainflux) - Industrial IoT Messaging and Device Management Server. Stars:`2.1K`.
- [ekuiper](https://github.com/lf-edge/ekuiper) - Lightweight data stream processing engine for IoT edge. Stars:`1.1K`.
- [gatt](https://github.com/paypal/gatt) - Gatt is a Go package for building Bluetooth Low Energy peripherals. Stars:`1.1K`.
- [connectordb](https://github.com/connectordb/connectordb) - Open-Source Platform for Quantified Self & IoT. Stars:`370`.
- [devices](https://github.com/goiot/devices) - Suite of libraries for IoT devices, experimental for x/exp/io. Stars:`260`.
- [huego](https://github.com/amimof/huego) - An extensive Philips Hue client library for Go. Stars:`235`.
- [sensorbee](https://github.com/sensorbee/sensorbee) - Lightweight stream processing engine for IoT. Stars:`221`.
- [eywa](https://github.com/xcodersun/eywa) - Project Eywa is essentially a connection manager that keeps track of connected devices. Stars:`60`.
- [gobot](https://github.com/hybridgroup/gobot/) - Gobot is a framework for robotics, physical computing, and the Internet of Things.
- [iot](https://github.com/vaelen/iot/) - IoT is a simple framework for implementing a Google IoT Core device.
- [periph](https://periph.io/) - Peripherals I/O to interface with low-level board facilities.

**[⬆ back to top](#contents)**

## Job Scheduler

_Libraries for scheduling jobs._

- [gocron](https://github.com/go-co-op/gocron) - Easy and fluent Go job scheduling. This is an actively maintained fork of [jasonlvhit/gocron](https://github.com/jasonlvhit/gocron). Stars:`3.7K`.
- [go-quartz](https://github.com/reugn/go-quartz) - Simple, zero-dependency scheduling library for Go. Stars:`1.3K`.
- [gron](https://github.com/roylee0704/gron) - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. Stars:`983`.
- [JobRunner](https://github.com/bamzi/jobrunner) - Smart and featureful cron job scheduler with job queuing and live monitoring built in. Stars:`974`.
- [Dagu](https://github.com/dagu-go/dagu) - No-code workflow executor. it executes DAGs defined in a simple YAML format. Stars:`840`.
- [jobs](https://github.com/albrow/jobs) - Persistent and flexible background jobs library. Stars:`495`.
- [scheduler](https://github.com/carlescere/scheduler) - Cronjobs scheduling made easy. Stars:`429`.
- [gronx](https://github.com/adhocore/gronx) - Cron expression parser, task runner and daemon consuming crontab like task list. Stars:`288`.
- [go-cron](https://github.com/rk/go-cron) - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. Stars:`218`.
- [tasks](https://github.com/madflojo/tasks) - An easy to use in-process scheduler for recurring tasks in Go. Stars:`167`.
- [goflow](https://github.com/fieldryand/goflow) - A simple but powerful DAG scheduler and dashboard. Stars:`151`.
- [clockwerk](https://github.com/onatm/clockwerk) - Go package to schedule periodic jobs using a simple, fluent syntax. Stars:`139`.
- [leprechaun](https://github.com/kilgaloon/leprechaun) - Job scheduler that supports webhooks, crons and classic scheduling. Stars:`99`.
- [cheek](https://github.com/datarootsio/cheek) - A simple crontab like scheduler that aims to offer a KISS approach to job scheduling. Stars:`89`.
- [cdule](https://github.com/deepaksinghvi/cdule) - Job scheduler library with database support Stars:`38`.
- [sched](https://github.com/romshark/sched) - A job scheduler with the ability to fast-forward time. Stars:`26`.
- [cronticker](https://github.com/krayzpipes/cronticker) - A ticker implementation to support cron schedules. Stars:`10`.

**[⬆ back to top](#contents)**

## JSON

_Libraries for working with JSON._

- [GJSON](https://github.com/tidwall/gjson) - Get a JSON value with one line of code. Stars:`12.6K`.
- [gojson](https://github.com/ChimeraCoder/gojson) - Automatically generate Go (golang) struct definitions from example JSON. Stars:`2.6K`.
- [fastjson](https://github.com/valyala/fastjson) - Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection. Stars:`2.0K`.
- [OjG](https://github.com/ohler55/ojg) - Optimized JSON for Go is a high performance parser with a variety of additional JSON tools including JSONPath. Stars:`628`.
- [marshmallow](https://github.com/PerimeterX/marshmallow) - Performant JSON unmarshaling for flexible use cases. Stars:`313`.
- [jsondiff](https://github.com/wI2L/jsondiff) - JSON diff library for Go based on RFC6902 (JSON Patch). Stars:`262`.
- [kazaam](https://github.com/Qntfy/kazaam) - API for arbitrary transformation of JSON documents. Stars:`259`.
- [gojq](https://github.com/elgs/gojq) - JSON query in Golang. Stars:`187`.
- [ajson](https://github.com/spyzhov/ajson) - Abstract JSON for golang with JSONPath support. Stars:`180`.
- [jsonvalue](https://github.com/Andrew-M-C/go.jsonvalue) - A fast and convinient library for unstructured JSON data, replacing `encoding/json`. Stars:`156`.
- [jettison](https://github.com/wI2L/jettison) - Fast and flexible JSON encoder for Go. Stars:`154`.
- [json2go](https://github.com/m-zajac/json2go) - Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. Stars:`119`.
- [gjo](https://github.com/skanehira/gjo) - Small utility to create JSON objects. Stars:`119`.
- [jsongo](https://github.com/ricardolonga/jsongo) - Fluent API to make it easier to create Json objects. Stars:`109`.
- [JayDiff](https://github.com/yazgazan/jaydiff) - JSON diff utility written in Go. Stars:`99`.
- [ujson](https://github.com/olvrng/ujson) - Fast and minimal JSON parser and transformer that works on unstructured JSON. Stars:`68`.
- [jsonf](https://github.com/miolini/jsonf) - Console tool for highlighted formatting and struct query fetching JSON. Stars:`64`.
- [jscan](https://github.com/romshark/jscan) - High performance zero-allocation JSON iterator. Stars:`57`.
- [go-respond](https://github.com/nicklaw5/go-respond) - Go package for handling common HTTP JSON responses. Stars:`51`.
- [mp](https://github.com/sanbornm/mp) - Simple cli email parser. It currently takes stdin and outputs JSON. Stars:`46`.
- [vjson](https://github.com/miladibra10/vjson) - Go package for validating JSON objects with declaring a JSON schema with fluent API. Stars:`34`.
- [jsoncolor](https://github.com/neilotoole/jsoncolor) - Drop-in replacement for `encoding/json` that outputs colorized JSON. Stars:`33`.
- [ask](https://github.com/simonnilsson/ask) - Easy access to nested values in maps and slices. Works in combination with encoding/json and other packages that "Unmarshal" arbitrary data into Go data-types. Stars:`25`.
- [gojmapr](https://github.com/limiu82214/gojmapr) - Get simple struct from complex json by json path. Stars:`19`.
- [mapslice-json](https://github.com/mickep76/mapslice-json) - Go MapSlice for ordered marshal/ unmarshal of maps in JSON. Stars:`15`.
- [dynjson](https://github.com/cocoonspace/dynjson) - Client-customizable JSON formats for dynamic APIs. Stars:`15`.
- [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) - Go bindings based on the JSON API errors reference. Stars:`13`.
- [go-jsonerror](https://github.com/ddymko/go-jsonerror) - Go-JsonError is meant to allow us to easily create json response errors that follow the JsonApi spec. Stars:`13`.
- [jsonhal](https://github.com/RichardKnop/jsonhal) - Simple Go package to make custom structs marshal into HAL compatible JSON responses. Stars:`13`.
- [epoch](https://github.com/vtopc/epoch) - Contains primitives for marshaling/unmarshaling Unix timestamp/epoch to/from build-in time.Time type in JSON. Stars:`12`.
- [jzon](https://github.com/zerosnake0/jzon) - JSON library with standard compatible API/behavior. Stars:`11`.
- [jsonic](https://github.com/sinhashubham95/jsonic) - Utilities to handle and query JSON without defining structs in a type safe manner. Stars:`9`.
- [ej](https://github.com/lucassscaravelli/ej) - Write and read JSON from different sources succinctly. Stars:`9`.
- [omg.jsonparser](https://github.com/dedalqq/omg.jsonparser) - Simple JSON parser with validation by condition via golang struct fields tags. Stars:`4`.
- [htmljson](https://github.com/nikolaydubina/htmljson) - Rich rendering of JSON as HTML in Go. Stars:`3`.
- [jsonhandlers](https://github.com/abusomani/jsonhandlers) - JSON library to expose simple handlers that lets you easily read and write json from various sources. Stars:`1`.
- [JSON-to-Proto](https://json-to-proto.github.io/) - Convert JSON to Protobuf online.
- [JSON-to-Go](https://mholt.github.io/json-to-go/) - Convert JSON to Go struct.

**[⬆ back to top](#contents)**

## Logging

_Libraries for generating and working with log files._

- [logrus](https://github.com/Sirupsen/logrus) - Structured logger for Go. Stars:`23.0K`.
- [zap](https://github.com/uber-go/zap) - Fast, structured, leveled logging in Go. Stars:`19.2K`.
- [zerolog](https://github.com/rs/zerolog) - Zero-allocation JSON logger. Stars:`8.6K`.
- [spew](https://github.com/davecgh/go-spew) - Implements a deep pretty printer for Go data structures to aid in debugging. Stars:`5.7K`.
- [lumberjack](https://github.com/natefinch/lumberjack) - Simple rolling logger, implements io.WriteCloser. Stars:`4.2K`.
- [glog](https://github.com/golang/glog) - Leveled execution logs for Go. Stars:`3.4K`.
- [tail](https://github.com/hpcloud/tail) - Go package striving to emulate the features of the BSD tail program. Stars:`2.6K`.
- [pp](https://github.com/k0kubun/pp) - Colored pretty printer for Go language. Stars:`1.7K`.
- [seelog](https://github.com/cihub/seelog) - Logging functionality with flexible dispatching, filtering, and formatting. Stars:`1.6K`.
- [log](https://github.com/apex/log) - Structured logging package for Go. Stars:`1.3K`.
- [log15](https://github.com/inconshreveable/log15) - Simple, powerful logging for Go. Stars:`1.1K`.
- [phuslu/log](https://github.com/phuslu/log) - High performance structured logging. Stars:`501`.
- [onelog](https://github.com/francoispqt/onelog) - Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenarios. Also, it is one of the logger with the lowest allocation. Stars:`414`.
- [logutils](https://github.com/hashicorp/logutils) - Utilities for slightly better logging in Go (Golang) extending the standard logger. Stars:`355`.
- [logxi](https://github.com/mgutz/logxi) - 12-factor app logger that is fast and makes you happy. Stars:`351`.
- [sqldb-logger](https://github.com/simukti/sqldb-logger) - A logger for Go SQL database driver without modify existing \*sql.DB stdlib usage. Stars:`312`.
- [httpretty](https://github.com/henvic/httpretty) - Pretty-prints your regular HTTP requests on your terminal for debugging (similar to http.DumpRequest). Stars:`302`.
- [log](https://github.com/go-playground/log) - Simple, configurable and scalable Structured Logging for Go. Stars:`289`.
- [go-logger](https://github.com/apsdehal/go-logger) - Simple logger of Go Programs, with level handlers. Stars:`286`.
- [rollingwriter](https://github.com/arthurkiller/rollingWriter) - RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. Stars:`281`.
- [slog](https://github.com/gookit/slog) - Lightweight, configurable, extensible logger for Go. Stars:`269`.
- [tint](https://github.com/lmittmann/tint) - A slog.Handler that writes tinted logs. Stars:`214`.
- [logur](https://github.com/logur/logur) - An opinionated logger interface and collection of logging best practices with adapters and integrations for well-known libraries ([logrus](https://github.com/sirupsen/logrus), [go-kit log](https://github.com/go-kit/kit/tree/master/log), [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog), etc). Stars:`197`.
- [glg](https://github.com/kpango/glg) - glg is simple and fast leveled logging library for Go. Stars:`188`.
- [logger](https://github.com/azer/logger) - Minimalistic logging library for Go. Stars:`157`.
- [xlog](https://github.com/rs/xlog) - Structured logger for `net/context` aware HTTP handlers with flexible dispatching. Stars:`137`.
- [ozzo-log](https://github.com/go-ozzo/ozzo-log) - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). Stars:`122`.
- [log-voyage](https://github.com/firstrow/logvoyage) - Full-featured logging saas written in golang. Stars:`94`.
- [slog-multi](https://github.com/samber/slog-multi) - Chain of slog.Handler (pipeline, fanout...). Stars:`68`.
- [go-cronowriter](https://github.com/utahta/go-cronowriter) - Simple writer that rotate log files automatically based on current date and time, like cronolog. Stars:`56`.
- [stdlog](https://github.com/alexcesaro/log) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. Stars:`47`.
- [noodlog](https://github.com/gyozatech/noodlog) - Parametrized JSON logging library which lets you obfuscate sensitive data and marshal any kind of content. No more printed pointers instead of values, nor escape chars for the JSON strings. Stars:`44`.
- [go-log](https://github.com/ian-kent/go-log) - Log4j implementation in Go. Stars:`42`.
- [logex](https://github.com/chzyer/logex) - Golang log lib, supports tracking and level, wrap by standard log lib. Stars:`41`.
- [gologger](https://github.com/sadlil/gologger) - Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. Stars:`41`.
- [journald](https://github.com/ssgreg/journald) - Go implementation of systemd Journal's native API for logging. Stars:`37`.
- [go-log](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers. Stars:`33`.
- [mlog](https://github.com/jbrodriguez/mlog) - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. Stars:`33`.
- [distillog](https://github.com/amoghe/distillog) - distilled levelled logging (think of it as stdlib + log levels). Stars:`32`.
- [logrusly](https://github.com/sebest/logrusly) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). Stars:`28`.
- [log](https://github.com/teris-io/log) - Structured log interface for Go cleanly separates logging facade from its implementation. Stars:`26`.
- [slog-formatter](https://github.com/samber/slog-formatter) - Common formatters for slog and helpers to build your own. Stars:`24`.
- [zkits-logger](https://github.com/edoger/zkits-logger) - A powerful zero-dependency JSON logger. Stars:`24`.
- [gomol](https://github.com/aphistic/gomol) - Multiple-output, structured logging for Go with extensible logging outputs. Stars:`19`.
- [logmatic](https://github.com/borderstech/logmatic) - Colorized logger for Golang with dynamic log level configuration. Stars:`16`.
- [xylog](https://github.com/xybor-x/xylog) - Leveled and structured logging, dynamic fields, high performance, zone management, simple configuration, and readable syntax. Stars:`15`.
- [glo](https://github.com/lajosbencz/glo) - PHP Monolog inspired logging facility with identical severity levels. Stars:`15`.
- [logrusiowriter](https://github.com/cabify/logrusiowriter) - `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. Stars:`14`.
- [log](https://github.com/heartwilltell/log) - Simple leveled logging wrapper around standard log package. Stars:`14`.
- [go-log](https://github.com/subchen/go-log) - Simple and configurable Logging in Go, with level, formatters and writers. Stars:`14`.
- [logo](https://github.com/mbndr/logo) - Golang logger to different configurable writers. Stars:`11`.
- [logdump](https://github.com/ewwwwwqm/logdump) - Package for multi-level logging. Stars:`11`.
- [kemba](https://github.com/clok/kemba) - A tiny debug logging tool inspired by [debug](https://github.com/visionmedia/debug), great for CLI tools and applications. Stars:`10`.
- [log](https://github.com/aerogo/log) - An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). Stars:`10`.
- [go-log](https://github.com/pieterclaerhout/go-log) - A logging library with stack traces, object dumping and optional timestamps. Stars:`10`.
- [xlog](https://github.com/xfxdev/xlog) - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. Stars:`8`.
- [zax](https://github.com/yuseferi/zax) - Integrate Context with Zap logger, which leads to more flexibility in Go logging. Stars:`6`.
- [structy/log](https://github.com/structy/log) - A simple to use log system, minimalist but with features for debugging and differentiation of messages. Stars:`5`.
- [slf4g](https://github.com/echocat/slf4g) - Simple Logging Facade for Golang: Simple structured logging; but powerful, extendable and customizable, with huge amount of learnings from decades of past logging frameworks. Stars:`3`.
- [log](https://github.com/no-src/log) - A simple logging framework out of the box. Stars:`2`.
- [yell](https://github.com/jfcg/yell) - Yet another minimalistic logging library. Stars:`1`.
- [gone/log](https://github.com/One-com/gone/tree/master/log) - Fast, extendable, full-featured, std-lib source compatible log library.

**[⬆ back to top](#contents)**

## Machine Learning

_Libraries for Machine Learning._

- [GoLearn](https://github.com/sjwhitworth/golearn) - General Machine Learning library for Go. Stars:`9.0K`.
- [gorse](https://github.com/zhenghaoz/gorse) - An offline recommender system backend based on collaborative filtering written in Go. Stars:`7.5K`.
- [gorgonia](https://github.com/gorgonia/gorgonia) - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. Stars:`5.1K`.
- [m2cgen](https://github.com/BayesWitnesses/m2cgen) - A CLI tool to transpile trained classic ML models into a native Go code with zero dependencies, written in Python with Go language support. Stars:`2.6K`.
- [tfgo](https://github.com/galeone/tfgo) - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. Stars:`2.3K`.
- [gosseract](https://github.com/otiai10/gosseract) - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. Stars:`2.2K`.
- [goml](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go. Stars:`1.5K`.
- [eaopt](https://github.com/MaxHalford/eaopt) - An evolutionary optimization library. Stars:`845`.
- [bayesian](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang. Stars:`769`.
- [CloudForest](https://github.com/ryanbressler/CloudForest) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. Stars:`728`.
- [ocrserver](https://github.com/otiai10/ocrserver) - A simple OCR API server, seriously easy to be deployed by Docker and Heroku. Stars:`587`.
- [gobrain](https://github.com/goml/gobrain) - Neural Networks written in go. Stars:`544`.
- [onnx-go](https://github.com/owulveryck/onnx-go) - Go Interface to Open Neural Network Exchange (ONNX). Stars:`521`.
- [go-deep](https://github.com/patrikeh/go-deep) - A feature-rich neural network library in Go. Stars:`463`.
- [regommend](https://github.com/muesli/regommend) - Recommendation & collaborative filtering engine. Stars:`302`.
- [Goptuna](https://github.com/c-bata/goptuna) - Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. Stars:`238`.
- [go-galib](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang. Stars:`195`.
- [goga](https://github.com/tomcraven/goga) - Genetic algorithm library for Go. Stars:`195`.
- [goRecommend](https://github.com/timkaye11/goRecommend) - Recommendation Algorithms library written in Go. Stars:`194`.
- [shield](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go. Stars:`154`.
- [go-fann](https://github.com/white-pony/go-fann) - Go bindings for Fast Artificial Neural Networks(FANN) library. Stars:`114`.
- [go-featureprocessing](https://github.com/nikolaydubina/go-featureprocessing) - Fast and convenient feature processing for low latency machine learning in Go. Stars:`100`.
- [goscore](https://github.com/asafschers/goscore) - Go Scoring API for PMML. Stars:`89`.
- [gonet](https://github.com/dathoangnd/gonet) - Neural Network for Go. Stars:`79`.
- [fonet](https://github.com/Fontinalis/fonet) - A Deep Neural Network library written in Go. Stars:`77`.
- [libsvm](https://github.com/datastream/libsvm) - libsvm golang version derived work based on LIBSVM 3.14. Stars:`73`.
- [neat](https://github.com/jinyeom/neat) - Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). Stars:`67`.
- [neural-go](https://github.com/schuyler/neural-go) - Multilayer perceptron network implemented in Go, with training via backpropagation. Stars:`65`.
- [go-pr](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang. Stars:`62`.
- [GoMind](https://github.com/surenderthakran/gomind) - A simplistic Neural Network Library in Go. Stars:`59`.
- [Varis](https://github.com/Xamber/Varis) - Golang Neural Network. Stars:`52`.
- [golinear](https://github.com/danieldk/golinear) - liblinear bindings for Go. Stars:`45`.
- [go-cluster](https://github.com/e-XpertSolutions/go-cluster) - Go implementation of the k-modes and k-prototypes clustering algorithms. Stars:`39`.
- [godist](https://github.com/e-dard/godist) - Various probability distributions, and associated methods. Stars:`35`.
- [randomforest](https://github.com/malaschitz/randomForest) - Easy to use Random Forest library for Go. Stars:`33`.
- [ddt](https://github.com/sgrodriguez/ddt) - Dynamic decision tree, create trees defining customizable rules. Stars:`29`.
- [evoli](https://github.com/khezen/evoli) - Genetic Algorithm and Particle Swarm Optimization library. Stars:`26`.
- [probab](https://github.com/ThePaw/probab) - Probability distribution functions. Bayesian inference. Written in pure Go. Stars:`18`.

**[⬆ back to top](#contents)**

## Messaging

_Libraries that implement messaging systems._

- [sarama](https://github.com/Shopify/sarama) - Go library for Apache Kafka. Stars:`10.2K`.
- [Centrifugo](https://github.com/centrifugal/centrifugo) - Real-time messaging (Websockets or SockJS) server in Go. Stars:`7.3K`.
- [gorush](https://github.com/appleboy/gorush) - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). Stars:`7.2K`.
- [machinery](https://github.com/RichardKnop/machinery) - Asynchronous task queue/job queue based on distributed message passing. Stars:`6.9K`.
- [Asynq](https://github.com/hibiken/asynq) - A simple, reliable, and efficient distributed task queue for Go built on top of Redis. Stars:`6.5K`.
- [Benthos](https://github.com/Jeffail/benthos) - A message streaming bridge between a range of protocols. Stars:`6.4K`.
- [Watermill](https://github.com/ThreeDotsLabs/watermill) - Working efficiently with message streams. Building event driven applications, enabling event sourcing, RPC over messages, sagas. Can use conventional pub/sub implementations like Kafka or RabbitMQ, but also HTTP or MySQL binlog. Stars:`6.0K`.
- [go-socket.io](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework. Stars:`5.3K`.
- [NATS Go Client](https://github.com/nats-io/nats) - Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. Stars:`4.7K`.
- [Confluent Kafka Golang Client](https://github.com/confluentinc/confluent-kafka-go) - confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform. Stars:`4.1K`.
- [Mercure](https://github.com/dunglas/mercure) - Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). Stars:`3.4K`.
- [melody](https://github.com/olahol/melody) - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. Stars:`3.2K`.
- [APNs2](https://github.com/sideshow/apns2) - HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. Stars:`2.8K`.
- [go-nsq](https://github.com/nsqio/go-nsq) - the official Go package for NSQ. Stars:`2.4K`.
- [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) - gopush-cluster is a go push server cluster. Stars:`2.1K`.
- [Uniqush-Push](https://github.com/uniqush/uniqush-push) - Redis backed unified push service for server-side notifications to mobile devices. Stars:`1.5K`.
- [EventBus](https://github.com/asaskevich/EventBus) - The lightweight event bus with async compatibility. Stars:`1.5K`.
- [Beaver](https://github.com/Clivern/Beaver) - A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. Stars:`1.5K`.
- [Chanify](https://github.com/chanify/chanify) - A push notification server send message to your iOS devices. Stars:`1.1K`.
- [zmq4](https://github.com/pebbe/zmq4) - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). Stars:`1.1K`.
- [amqp](https://github.com/rabbitmq/amqp091-go) - Go RabbitMQ Client Library. Stars:`941`.
- [Gollum](https://github.com/trivago/gollum) - A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. Stars:`930`.
- [dbus](https://github.com/godbus/dbus) - Native Go bindings for D-Bus. Stars:`873`.
- [golongpoll](https://github.com/jcuga/golongpoll) - HTTP longpoll server library that makes web pub-sub simple. Stars:`636`.
- [mangos](https://github.com/nanomsg/mangos) - Pure go implementation of the Nanomsg ("Scalability Protocols") with transport interoperability. Stars:`601`.
- [emitter](https://github.com/olebedev/emitter) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. Stars:`471`.
- [Glue](https://github.com/desertbit/glue) - Robust Go and Javascript Socket Library (Alternative to Socket.io). Stars:`409`.
- [pubsub](https://github.com/tuxychandru/pubsub) - Simple pubsub package for go. Stars:`391`.
- [Quamina](https://github.com/timbray/quamina) - Fast pattern-matching for filtering messages and events. Stars:`348`.
- [Bus](https://github.com/mustafaturan/bus) - Minimalist message bus implementation for internal communication. Stars:`305`.
- [messagebus](https://github.com/vardius/message-bus) - messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. Stars:`252`.
- [rabtap](https://github.com/jandelgado/rabtap) - RabbitMQ swiss army knife cli app. Stars:`242`.
- [guble](https://github.com/smancke/guble) - Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. Stars:`155`.
- [hub](https://github.com/leandro-lugaresi/hub) - A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. Stars:`132`.
- [redisqueue](https://github.com/robinjoseph08/redisqueue) - redisqueue provides a producer and consumer of a queue that uses Redis streams. Stars:`114`.
- [oplog](https://github.com/dailymotion/oplog) - Generic oplog/replication system for REST APIs. Stars:`114`.
- [rabbus](https://github.com/rafaeljesus/rabbus) - A tiny wrapper over amqp exchanges and queues. Stars:`97`.
- [go-mq](https://github.com/cheshir/go-mq) - RabbitMQ client with declarative configuration. Stars:`87`.
- [Ratus](https://github.com/hyperonym/ratus) - Ratus is a RESTful asynchronous task queue server. Stars:`87`.
- [drone-line](https://github.com/appleboy/drone-line) - Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. Stars:`79`.
- [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) - A tiny wrapper around NSQ topic and channel. Stars:`77`.
- [go-notify](https://github.com/TheCreeper/go-notify) - Native implementation of the freedesktop notification spec. Stars:`67`.
- [RapidMQ](https://github.com/sybrexsys/RapidMQ) - RapidMQ is a lightweight and reliable library for managing of the local messages queue. Stars:`65`.
- [Commander](https://github.com/jeroenrinzema/commander) - A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. Stars:`65`.
- [go-res](https://github.com/jirenius/go-res) - Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate. Stars:`63`.
- [event](https://github.com/agoalofalife/event) - Implementation of the pattern observer. Stars:`54`.
- [hare](https://github.com/leozz37/hare) - A user friendly library for sending messages and listening to TCP sockets. Stars:`52`.
- [ami](https://github.com/kak-tus/ami) - Go client to reliable queues based on Redis Cluster Streams. Stars:`26`.
- [go-vitotrol](https://github.com/maxatome/go-vitotrol) - Client library to Viessmann Vitotrol web service. Stars:`22`.
- [gosd](https://github.com/alexsniffin/gosd) - A library for scheduling when to dispatch a message to a channel. Stars:`22`.
- [rmqconn](https://github.com/sbabiv/rmqconn) - RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. Stars:`20`.
- [jazz](https://github.com/socifi/jazz) - A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. Stars:`18`.
- [gaurun-client](https://github.com/osamingo/gaurun-client) - Gaurun Client written in Go. Stars:`11`.
- [mob](https://github.com/erni27/mob) - mob is a generic-based, simple mediator / event aggregator library. It supports in-process requests / events processing.

**[⬆ back to top](#contents)**

## Microsoft Office

- [unioffice](https://github.com/unidoc/unioffice) - Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. Stars:`3.9K`.

### Microsoft Excel

_Libraries for working with Microsoft Excel._

- [excelize](https://github.com/xuri/excelize) - Golang library for reading and writing Microsoft Excel&trade; (XLSX) files. Stars:`15.6K`.
- [xlsx](https://github.com/tealeg/xlsx) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. Stars:`5.5K`.
- [go-excel](https://github.com/szyhf/go-excel) - A simple and light reader to read a relate-db-like excel as a table. Stars:`175`.
- [xlsx](https://github.com/plandem/xlsx) - Fast and safe way to read/update your existing Microsoft Excel files in Go programs. Stars:`165`.
- [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) - Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. Stars:`21`.
- [exl](https://github.com/go-the-way/exl) - Excel binding to struct written in Go.(Only supports Go1.18+) Stars:`17`.

**[⬆ back to top](#contents)**

## Miscellaneous

### Dependency Injection

_Libraries for working with dependency injection._

- [google/wire](https://github.com/google/wire) - Automated Initialization in Go. Stars:`11.0K`.
- [fx](https://github.com/uber-go/fx) - A dependency injection based application framework for Go (built on top of dig). Stars:`4.2K`.
- [dig](https://github.com/uber-go/dig) - A reflection based dependency injection toolkit for Go. Stars:`3.3K`.
- [do](https://github.com/samber/do) - A dependency injection framework based on Generics. Stars:`1.2K`.
- [GoLobby/Container](https://github.com/golobby/container) - GoLobby Container is a lightweight yet powerful IoC dependency injection container for the Go programming language. Stars:`477`.
- [goioc/di](https://github.com/goioc/di) - Spring-inspired Dependency Injection Container. Stars:`288`.
- [di](https://github.com/goava/di) - A dependency injection container for go programming language. Stars:`197`.
- [dingo](https://github.com/i-love-flamingo/dingo) - A dependency injection toolkit for Go, based on Guice. Stars:`164`.
- [alice](https://github.com/magic003/alice) - Additive dependency injection container for Golang. Stars:`51`.
- [wire](https://github.com/Fs02/wire) - Strict Runtime Dependency Injection for Golang. Stars:`38`.
- [linker](https://github.com/logrange/linker) - A reflection based dependency injection and inversion of control library with components lifecycle support. Stars:`35`.
- [nject](https://github.com/muir/nject) - A type safe, reflective framework for libraries, tests, http endpoints, and service startup. Stars:`27`.
- [gocontainer](https://github.com/vardius/gocontainer) - Simple Dependency Injection Container. Stars:`18`.
- [kinit](https://github.com/go-kata/kinit) - Customizable dependency injection container with the global mode, cascade initialization and panic-safe finalization. Stars:`8`.
- [HnH/di](https://github.com/HnH/di) - DI container library that is focused on clean API and flexibility. Stars:`6`.

**[⬆ back to top](#contents)**

### Project Layout

_**Unofficial** set of patterns for structuring projects._

- [golang-standards/project-layout](https://github.com/golang-standards/project-layout) - Set of common historical and emerging project layout patterns in the Go ecosystem. Note: despite the org-name they do not represent official golang standards, see [this issue](https://github.com/golang-standards/project-layout/issues/117) for more information. Nonetheless, some may find the layout useful. Stars:`41.0K`.
- [ardanlabs/service](https://github.com/ardanlabs/service) - A [starter kit](https://github.com/ardanlabs/service/wiki) for building production grade scalable web service applications. Stars:`2.8K`.
- [modern-go-application](https://github.com/sagikazarmark/modern-go-application) - Go application boilerplate and example applying modern practices. Stars:`1.6K`.
- [pagoda](https://github.com/mikestefanello/pagoda) - Rapid, easy full-stack web development starter kit built in Go. Stars:`866`.
- [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) - A Go application boilerplate template for quick starting projects following production best practices. Stars:`622`.
- [golang-templates/seed](https://github.com/golang-templates/seed) - Go application GitHub repository template. Stars:`390`.
- [go-starter](https://github.com/allaboutapps/go-starter) - An opinionated production-ready RESTful JSON backend template, highly integrated with VSCode DevContainers. Stars:`361`.
- [go-todo-backend](https://github.com/Fs02/go-todo-backend) - Go Todo Backend example using modular project layout for product microservice. Stars:`240`.
- [scaffold](https://github.com/catchplay/scaffold) - Scaffold generates a starter Go project layout. Lets you focus on business logic implemented. Stars:`142`.
- [go-sample](https://github.com/zitryss/go-sample) - A sample layout for Go application projects with the real code. Stars:`125`.
- [gobase](https://github.com/wajox/gobase) - A simple skeleton for golang application with basic setup for real golang application. Stars:`48`.
- [wangyoucao577/go-project-layout](https://github.com/wangyoucao577/go-project-layout) - Set of practices and discussions on how to structure Go project layout. Stars:`22`.
- [insidieux/inizio](https://github.com/insidieux/inizio) - Golang project layout generator with plugins. Stars:`16`.
- [go-module](https://github.com/octomation/go-module) - Template for a typical module written on Go. Stars:`13`.

**[⬆ back to top](#contents)**

### Strings

_Libraries for working with strings._

- [xstrings](https://github.com/huandu/xstrings) - Collection of useful string functions ported from other languages. Stars:`1.2K`.
- [sttr](https://github.com/abhimanyu003/sttr) - cross-platform, cli app to perform various operations on string. Stars:`735`.
- [strutil](https://github.com/ozgio/strutil) - String utilities. Stars:`196`.
- [gobeam/Stringy](https://github.com/gobeam/Stringy) - String manipulation library to convert string to camel case, snake case, kebab case / slugify etc. Stars:`188`.
- [caps](https://github.com/chanced/caps) - A case conversion library. Stars:`47`.
- [bexp](https://github.com/mkungla/bexp) - Go implementation of Brace Expansion mechanism to generate arbitrary strings.
- [go-formatter](https://gitlab.com/tymonx/go-formatter) - Implements **replacement fields** surrounded by curly braces `{}` format strings.

**[⬆ back to top](#contents)**

### Uncategorized

_These libraries were placed here because none of the other categories seemed to fit._

- [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). Stars:`9.3K`.
- [archiver](https://github.com/mholt/archiver) - Library and command for making and extracting .zip and .tar.gz archives. Stars:`4.0K`.
- [gatus](https://github.com/TwinProduction/gatus) - Automated service health dashboard. Stars:`3.9K`.
- [gofakeit](https://github.com/brianvoe/gofakeit) - Random data generator written in go. Stars:`3.2K`.
- [go-resiliency](https://github.com/eapache/go-resiliency) - Resiliency patterns for golang. Stars:`1.8K`.
- [base64Captcha](https://github.com/mojocn/base64Captcha) - Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. Stars:`1.8K`.
- [gosms](https://github.com/haxpax/gosms) - Your own local SMS gateway in Go that can be used to send SMS. Stars:`1.4K`.
- [go-commons-pool](https://github.com/jolestar/go-commons-pool) - Generic object pool for Golang. Stars:`1.2K`.
- [llvm](https://github.com/llir/llvm) - Library for interacting with LLVM IR in pure Go. Stars:`1.1K`.
- [shortid](https://github.com/teris-io/shortid) - Distributed generation of super short, unique, non-sequential, URL friendly IDs. Stars:`876`.
- [health](https://github.com/alexliesenfeld/health) - A simple and flexible health check library for Go. Stars:`686`.
- [shoutrrr](https://github.com/containrrr/shoutrrr) - Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others. Stars:`678`.
- [stateless](https://github.com/qmuntal/stateless) - A fluent library for creating state machines. Stars:`654`.
- [health](https://github.com/dimiro1/health) - Easy to use, extensible health check library. Stars:`444`.
- [banner](https://github.com/dimiro1/banner) - Add beautiful banners into your Go applications. Stars:`437`.
- [xz](https://github.com/ulikunitz/xz) - Pure golang package for reading and writing xz-compressed files. Stars:`425`.
- [gountries](https://github.com/pariz/gountries) - Package that exposes country and subdivision data. Stars:`392`.
- [conv](https://github.com/cstockton/go-conv) - Package conv provides fast and intuitive conversions across Go types. Stars:`386`.
- [lk](https://github.com/hyperboloide/lk) - A simple licensing library for golang. Stars:`306`.
- [ffmt](https://github.com/go-ffmt/ffmt) - Beautify data display for Humans. Stars:`294`.
- [healthcheck](https://github.com/etherlabsio/healthcheck) - An opinionated and concurrent health-check HTTP handler for RESTful services. Stars:`256`.
- [antch](https://github.com/antchfx/antch) - A fast, powerful and extensible web crawling & scraping framework. Stars:`252`.
- [go-unarr](https://github.com/gen2brain/go-unarr) - Decompression library for RAR, TAR, ZIP and 7z archives. Stars:`245`.
- [battery](https://github.com/distatus/battery) - Cross-platform, normalized battery information library. Stars:`225`.
- [bitio](https://github.com/icza/bitio) - Highly optimized bit-level Reader and Writer for Go. Stars:`221`.
- [stats](https://github.com/go-playground/stats) - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... Stars:`166`.
- [turtle](https://github.com/hackebrot/turtle) - Emojis for Go. Stars:`152`.
- [gtree](https://github.com/ddddddO/gtree) - Provide CLI, Package and Web for tree output and directories creation from Markdown or programmatically. Stars:`137`.
- [captcha](https://github.com/steambap/captcha) - Package captcha provides an easy to use, unopinionated API for captcha generation. Stars:`135`.
- [gotoprom](https://github.com/cabify/gotoprom) - Type-safe metrics builder wrapper library for the official Prometheus client. Stars:`106`.
- [indigo](https://github.com/osamingo/indigo) - Distributed unique ID generator of using Sonyflake and encoded by Base58. Stars:`104`.
- [gommit](https://github.com/antham/gommit) - Analyze git commit messages to ensure they follow defined patterns. Stars:`103`.
- [morse](https://github.com/alwindoss/morse) - Library to convert to and from morse code. Stars:`78`.
- [faker](https://github.com/pioz/faker) - Random fake data and struct generator for Go. Stars:`76`.
- [persian](https://github.com/mavihq/persian) - Some utilities for Persian language in go. Stars:`75`.
- [pdfgen](https://github.com/hyperboloide/pdfgen) - HTTP service to generate PDF from Json requests. Stars:`61`.
- [xkg](https://github.com/go-xkg/xkg) - X Keyboard Grabber. Stars:`56`.
- [datacounter](https://github.com/miolini/datacounter) - Go counters for readers/writer/http.ResponseWriter. Stars:`45`.
- [browscap_go](https://github.com/digitalcrab/browscap_go) - GoLang Library for [Browser Capabilities Project](https://browscap.org/). Stars:`44`.
- [url-shortener](https://github.com/pantrif/url-shortener) - A modern, powerful, and robust URL shortener microservice with mysql support. Stars:`42`.
- [sandid](https://github.com/aofei/sandid) - Every grain of sand on earth has its own ID. Stars:`40`.
- [autoflags](https://github.com/artyom/autoflags) - Go package to automatically define command line flags from struct fields. Stars:`39`.
- [gosh](https://github.com/osamingo/gosh) - Provide Go Statistics Handler, Struct, Measure Method. Stars:`34`.
- [xdg](https://github.com/rkoesters/xdg) - FreeDesktop.org (xdg) Specs implemented in Go. Stars:`33`.
- [metrics](https://github.com/pascaldekloe/metrics) - Library for metrics instrumentation and Prometheus exposition. Stars:`26`.
- [shellwords](https://github.com/Wing924/shellwords) - A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. Stars:`21`.
- [numa](https://github.com/lrita/numa) - NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. Stars:`16`.
- [anagent](https://github.com/mudler/anagent) - Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. Stars:`14`.
- [avgRating](https://github.com/kirillDanshin/avgRating) - Calculate average score and rating based on Wilson Score Equation. Stars:`14`.
- [hostutils](https://github.com/Wing924/hostutils) - A golang library for packing and unpacking FQDNs list. Stars:`12`.
- [go-commandbus](https://github.com/lana/go-commandbus) - A slight and pluggable command-bus for Go. Stars:`10`.
- [faker](https://github.com/neotoolkit/faker) - Fake data generator. Stars:`9`.
- [openapi](https://github.com/neotoolkit/openapi) - OpenAPI 3.x parser. Stars:`9`.
- [varint](https://github.com/chmike/varint) - A faster varying length integer encoder/decoder than the one provided in the standard library. Stars:`7`.
- [common](https://github.com/kubeservice-stack/common) - A library for server framework. Stars:`4`.
- [basexx](https://github.com/bobg/basexx) - Convert to, from, and between digit strings in various number bases. Stars:`4`.
- [sitemap-format](https://github.com/mingard/sitemap-format) - A simple sitemap generator, with a little syntactic sugar. Stars:`3`.
- [VarHandler](https://github.com/azr/generators/tree/master/varhandler) - Generate boilerplate http input and output handling.
- [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.

**[⬆ back to top](#contents)**

## Natural Language Processing

_Libraries for working with human languages._

See also [Text Processing](#text-processing) and [Text Analysis](#text-analysis).

### Language Detection

- [whatlanggo](https://github.com/abadojack/whatlanggo) - Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). Stars:`593`.
- [getlang](https://github.com/rylans/getlang) - Fast natural language detection package. Stars:`155`.
- [guesslanguage](https://github.com/endeveit/guesslanguage) - Functions to determine the natural language of a unicode text. Stars:`57`.
- [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) - Language Detection API Go Client. Supports batch requests, short phrase or single word language detection. Stars:`22`.

### Morphological Analyzers

- [spaGO](https://github.com/nlpodyssey/spago) - Self-contained Machine Learning and Natural Language Processing library in Go. Stars:`1.5K`.
- [kagome](https://github.com/ikawaha/kagome) - JP morphological analyzer written in pure Go. Stars:`742`.
- [nlp](https://github.com/james-bowman/nlp) - Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). Stars:`416`.
- [nlp](https://github.com/Shixzie/nlp) - Extract values from strings and fill your structs with nlp. Stars:`383`.
- [RAKE.go](https://github.com/afjoseph/RAKE.Go) - Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). Stars:`100`.
- [go-stem](https://github.com/agonopol/go-stem) - Implementation of the porter stemming algorithm. Stars:`75`.
- [go2vec](https://github.com/danieldk/go2vec) - Reader and utility functions for word2vec embeddings. Stars:`52`.
- [porter2](https://github.com/zhenjl/porter2) - Really fast Porter 2 stemmer. Stars:`46`.
- [snowball](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). Stars:`36`.
- [govader](https://github.com/jonreiter/govader) - Go implementation of [VADER Sentiment Analysis](https://github.com/cjhutto/vaderSentiment). Stars:`35`.
- [paicehusk](https://github.com/rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm. Stars:`29`.
- [golibstemmer](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2. Stars:`20`.
- [libtextcat](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. Stars:`13`.
- [porter](https://github.com/a2800276/porter) - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. Stars:`12`.
- [gosentiwordnet](https://github.com/dinopuguh/gosentiwordnet) - Sentiment analyzer using sentiwordnet lexicon in Go. Stars:`10`.
- [govader-backend](https://github.com/PIMPfiction/govader_backend) - Microservice implementation of [GoVader](https://github.com/jonreiter/govader). Stars:`5`.
- [spelling-corrector](https://github.com/jorelosorio/spellingcorrector) - A spelling corrector for the Spanish language or create your own. Stars:`2`.

### Slugifiers

- [slug](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support. Stars:`983`.
- [go-slugify](https://github.com/mozillazg/go-slugify) - Make pretty slug with multiple languages support. Stars:`87`.
- [Slugify](https://github.com/avelino/slugify) - Go slugify application that handles string. Stars:`32`.

### Tokenizers

- [prose](https://github.com/jdkato/prose) - Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. Stars:`3.0K`.
- [gse](https://github.com/go-ego/gse) - Go efficient text segmentation; support english, chinese, japanese and other. Stars:`2.3K`.
- [gojieba](https://github.com/yanyiwu/gojieba) - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. Stars:`2.2K`.
- [sentences](https://github.com/neurosnap/sentences) - Sentence tokenizer: converts text into a list of sentences. Stars:`374`.
- [segment](https://github.com/blevesearch/segment) - Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](https://www.unicode.org/reports/tr29/) Stars:`78`.
- [textcat](https://github.com/pebbe/textcat) - Go package for n-gram based text categorization, with support for utf-8 and raw text. Stars:`70`.
- [MMSEGO](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. Stars:`62`.
- [stemmer](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers. Stars:`51`.
- [gotokenizer](https://github.com/xujiajun/gotokenizer) - A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation) Stars:`18`.
- [shamoji](https://github.com/osamingo/shamoji) - The shamoji is word filtering package written in Go. Stars:`13`.

### Translation

- [go-pinyin](https://github.com/mozillazg/go-pinyin) - CN Hanzi to Hanyu Pinyin converter. Stars:`1.4K`.
- [gotext](https://github.com/leonelquinteros/gotext) - GNU gettext utilities for Go. Stars:`388`.
- [go-localize](https://github.com/m1/go-localize) - Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings. Stars:`54`.
- [iuliia-go](https://github.com/mehanizm/iuliia-go) - Transliterate Cyrillic → Latin in every possible way. Stars:`38`.
- [go-mystem](https://github.com/dveselov/mystem) - CGo bindings to Yandex.Mystem - russian morphology analyzer. Stars:`30`.
- [spreak](https://github.com/vorlif/spreak) - Flexible translation and humanization library for Go, based on the concepts behind gettext. Stars:`24`.
- [icu](https://github.com/goodsign/icu) - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. Stars:`21`.
- [t](https://github.com/youthlin/t) - Another i18n pkg for golang, which follows GNU gettext style and supports .po/.mo files: `t.T (gettext)`, `t.N (ngettext)`, etc. And it contains a cmd tool [xtemplate](https://github.com/youthlin/t/blob/main/cmd/xtemplate), which can extract messages as a pot file from text/html template. Stars:`15`.
- [go-i18n](https://github.com/nicksnyder/go-i18n/) - Package and an accompanying tool to work with localized text.

### Transliteration

- [go-unidecode](https://github.com/mozillazg/go-unidecode) - ASCII transliterations of Unicode text. Stars:`109`.
- [gounidecode](https://github.com/fiam/gounidecode) - Unicode transliterator (also known as unidecode) for Go. Stars:`79`.
- [transliterator](https://github.com/alexsergivan/transliterator) - Provides one-way string transliteration with supporting of language-specific transliteration rules. Stars:`35`.
- [enca](https://github.com/endeveit/enca) - Minimal cgo bindings for [libenca](https://cihar.com/software/enca/), which detects character encodings. Stars:`15`.

**[⬆ back to top](#contents)**

## Networking

_Libraries for working with various layers of the network._

- [fasthttp](https://github.com/valyala/fasthttp) - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. Stars:`19.9K`.
- [kcptun](https://github.com/xtaci/kcptun) - Extremely simple & fast udp tunnel based on KCP protocol. Stars:`13.4K`.
- [webrtc](https://github.com/pions/webrtc) - A pure Go implementation of the WebRTC API. Stars:`11.5K`.
- [quic-go](https://github.com/lucas-clemente/quic-go) - An implementation of the QUIC protocol in pure Go. Stars:`8.5K`.
- [gnet](https://github.com/panjf2000/gnet) - `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. Stars:`7.9K`.
- [dns](https://github.com/miekg/dns) - Go library for working with DNS. Stars:`7.1K`.
- [gopacket](https://github.com/google/gopacket) - Go library for packet processing with libpcap bindings. Stars:`5.7K`.
- [HTTPLab](https://github.com/gchaincl/httplab) - HTTPLabs let you inspect HTTP requests and forge responses. Stars:`3.9K`.
- [kcp-go](https://github.com/xtaci/kcp-go) - KCP - Fast and Reliable ARQ Protocol. Stars:`3.7K`.
- [netpoll](https://github.com/cloudwego/netpoll) - A high-performance non-blocking I/O networking framework, which focused on RPC scenarios, developed by ByteDance. Stars:`3.6K`.
- [gobgp](https://github.com/osrg/gobgp) - BGP implemented in the Go Programming Language. Stars:`3.3K`.
- [ssh](https://github.com/gliderlabs/ssh) - Higher-level API for building SSH servers (wraps crypto/ssh). Stars:`3.1K`.
- [fortio](https://github.com/fortio/fortio) - Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. Stars:`3.0K`.
- [water](https://github.com/songgao/water) - Simple TUN/TAP library. Stars:`1.7K`.
- [nbio](https://github.com/lesismal/nbio) - Pure Go 1000k+ connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use. Stars:`1.7K`.
- [gev](https://github.com/Allenxuxu/gev) - gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. Stars:`1.7K`.
- [go-getter](https://github.com/hashicorp/go-getter) - Go library for downloading files or directories from various sources using a URL. Stars:`1.5K`.
- [sftp](https://github.com/pkg/sftp) - Package sftp implements the SSH File Transfer Protocol as described in <https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt>. Stars:`1.4K`.
- [NFF-Go](https://github.com/intel-go/nff-go) - Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). Stars:`1.3K`.
- [grab](https://github.com/cavaliercoder/grab) - Go package for managing file downloads. Stars:`1.3K`.
- [ftp](https://github.com/jlaffaye/ftp) - Package ftp implements a FTP client as described in [RFC 959](https://tools.ietf.org/html/rfc959). Stars:`1.1K`.
- [mdns](https://github.com/hashicorp/mdns) - Simple mDNS (Multicast DNS) client/server library in Golang. Stars:`998`.
- [gosnmp](https://github.com/soniah/gosnmp) - Native Go library for performing SNMP actions. Stars:`977`.
- [vssh](https://github.com/yahoo/vssh) - Go library for building network and server automation over SSH protocol. Stars:`910`.
- [gmqtt](https://github.com/DrmagicE/gmqtt) - Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. Stars:`891`.
- [cidranger](https://github.com/yl2chen/cidranger) - Fast IP to CIDR lookup for Go. Stars:`849`.
- [easytcp](https://github.com/DarthPestilane/easytcp) - A light-weight TCP framework written in Go (Golang), built with message router. EasyTCP helps you build a TCP server easily fast and less painful. Stars:`704`.
- [lhttp](https://github.com/fanux/lhttp) - Powerful websocket framework, build your IM server more easily. Stars:`691`.
- [peerdiscovery](https://github.com/schollz/peerdiscovery) - Pure Go library for cross-platform local peer discovery using UDP multicast. Stars:`598`.
- [go-stun](https://github.com/ccding/go-stun) - Go implementation of the STUN client (RFC 3489 and RFC 5389). Stars:`592`.
- [gotcp](https://github.com/gansidui/gotcp) - Go package for quickly writing tcp applications. Stars:`508`.
- [gaio](https://github.com/xtaci/gaio) - High performance async-io networking for Golang in proactor mode. Stars:`506`.
- [stun](https://github.com/go-rtc/stun) - Go implementation of RFC 5389 STUN protocol. Stars:`487`.
- [gopcap](https://github.com/akrennmair/gopcap) - Go wrapper for libpcap. Stars:`473`.
- [tcp_server](https://github.com/firstrow/tcp_server) - Go library for building tcp servers faster. Stars:`423`.
- [raw](https://github.com/mdlayher/raw) - Package raw enables reading and writing data at the device driver level for a network interface. Stars:`421`.
- [winrm](https://github.com/masterzen/winrm) - Go WinRM client to remotely execute commands on Windows machines. Stars:`405`.
- [gws](https://github.com/lxzan/gws) - High-Performance WebSocket Server & Client With AsyncIO Supporting . Stars:`394`.
- [ftpserverlib](https://github.com/fclairamb/ftpserverlib) - Fully featured FTP server library. Stars:`371`.
- [arp](https://github.com/mdlayher/arp) - Package arp implements the ARP protocol, as described in RFC 826. Stars:`328`.
- [dnsmonster](https://github.com/mosajjal/dnsmonster) - Passive DNS Capture/Monitoring Framework. Stars:`273`.
- [ethernet](https://github.com/mdlayher/ethernet) - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. Stars:`269`.
- [buffstreams](https://github.com/stabbycutyou/buffstreams) - Streaming protocolbuffer data over TCP made easy. Stars:`253`.
- [gNxI](https://github.com/google/gnxi) - A collection of tools for Network Management that use the gNMI and gNOI protocols. Stars:`242`.
- [jazigo](https://github.com/udhos/jazigo) - Jazigo is a tool written in Go for retrieving configuration for multiple network devices. Stars:`195`.
- [utp](https://github.com/anacrolix/utp) - Go uTP micro transport protocol implementation. Stars:`170`.
- [canopus](https://github.com/zubairhamed/canopus) - CoAP Client/Server implementation (RFC 7252). Stars:`152`.
- [xtcp](https://github.com/xfxdev/xtcp) - TCP Server Framework with simultaneous full duplex communication, graceful shutdown, and custom protocol. Stars:`144`.
- [sslb](https://github.com/eduardonunesp/sslb) - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. Stars:`143`.
- [iplib](https://github.com/c-robinson/iplib) - Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html) Stars:`116`.
- [gldap](https://github.com/jimlambrt/gldap) - gldap provides an ldap server implementation and you provide handlers for its ldap operations. Stars:`95`.
- [ether](https://github.com/songgao/ether) - Cross-platform Go package for sending and receiving ethernet frames. Stars:`79`.
- [dhcp6](https://github.com/mdlayher/dhcp6) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. Stars:`76`.
- [packet](https://github.com/aerogo/packet) - Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. Stars:`74`.
- [go-powerdns](https://github.com/joeig/go-powerdns) - PowerDNS API bindings for Golang. Stars:`71`.
- [fullproxy](https://github.com/shoriwe/fullproxy) - A fully featured scriptable and daemon configurable proxy and pivoting toolkit with SOCKS5, HTTP, raw ports and reverse proxy protocols. Stars:`59`.
- [portproxy](https://github.com/aybabtme/portproxy) - Simple TCP proxy which adds CORS support to API's which don't support it. Stars:`53`.
- [linkio](https://github.com/ian-kent/linkio) - Network link speed simulation for Reader/Writer interfaces. Stars:`52`.
- [panoptes-stream](https://github.com/yahoo/panoptes-stream) - A cloud native distributed streaming network telemetry (gNMI, Juniper JTI and Cisco MDT). Stars:`39`.
- [golibwireshark](https://github.com/sunwxg/golibwireshark) - Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. Stars:`29`.
- [graval](https://github.com/koofr/graval) - Experimental FTP server framework. Stars:`28`.
- [publicip](https://github.com/polera/publicip) - Package publicip returns your public facing IPv4 address (internet egress). Stars:`26`.
- [httpproxy](https://github.com/wzshiming/httpproxy) - HTTP proxy handler and dialer. Stars:`21`.
- [goshark](https://github.com/sunwxg/goshark) - Package goshark use tshark to decode IP packet and create data struct to analyse packet. Stars:`16`.
- [tspool](https://github.com/two/tspool) - A TCP Library use worker pool to improve performance and protect your server. Stars:`14`.
- [llb](https://github.com/kirillDanshin/llb) - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. Stars:`14`.
- [gnet](https://github.com/fish-tennis/gnet) - `gnet` is a high-performance networking framework,especially for game servers. Stars:`9`.
- [go-sse](https://github.com/lampctl/go-sse) - Go client and server implementation of HTML server-sent events. Stars:`2`.
- [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.

**[⬆ back to top](#contents)**

### HTTP Clients

_Libraries for making HTTP requests._

- [resty](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client. Stars:`8.1K`.
- [req](https://github.com/imroc/req) - Simple Go HTTP client with Black Magic (Less code and More efficiency). Stars:`3.5K`.
- [heimdall](https://github.com/gojektech/heimdall) - An enhanced http client with retry and hystrix capabilities. Stars:`2.5K`.
- [grequests](https://github.com/levigross/grequests) - A Go "clone" of the great and famous Requests library. Stars:`2.0K`.
- [go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) - Retryable HTTP client in Go. Stars:`1.6K`.
- [sling](https://github.com/dghubble/sling) - Sling is a Go HTTP client library for creating and sending API requests. Stars:`1.6K`.
- [requests](https://github.com/carlmjohnson/requests) - HTTP requests for Gophers. Uses context.Context and doesn't hide the underlying net/http.Client, making it compatible with standard Go APIs. Also includes testing tools. Stars:`1.1K`.
- [gentleman](https://github.com/h2non/gentleman) - Full-featured plugin-driven HTTP client library. Stars:`1.0K`.
- [pester](https://github.com/sethgrid/pester) - Go HTTP client calls with retries, backoff, and concurrency. Stars:`621`.
- [go-cleanhttp](https://github.com/hashicorp/go-cleanhttp) - Get easily stdlib HTTP client, which does not share any state with other clients. Stars:`306`.
- [request](https://github.com/monaco-io/request) - HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency. Stars:`245`.
- [go-otelroundtripper](https://github.com/NdoleStudio/go-otelroundtripper) - Go http.RoundTripper that emits open telemetry metrics for HTTP requests. Stars:`73`.
- [go-http-client](https://github.com/bozd4g/go-http-client) - Make http calls simply and easily. Stars:`61`.
- [go-zoox/fetch](https://github.com/go-zoox/fetch) - A Powerful, Lightweight, Easy Http Client, inspired by Web Fetch API. Stars:`49`.
- [rq](https://github.com/ddo/rq) - A nicer interface for golang stdlib HTTP client. Stars:`49`.
- [httpretry](https://github.com/ybbus/httpretry) - Enriches the default go HTTP client with retry functionality. Stars:`39`.
- [go-req](https://github.com/wenerme/go-req) - Declarative golang HTTP client. Stars:`20`.

**[⬆ back to top](#contents)**

## OpenGL

_Libraries for using OpenGL in Go._

- [glfw](https://github.com/go-gl/glfw) - Go bindings for GLFW 3. Stars:`1.4K`.
- [gl](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow). Stars:`981`.
- [mathgl](https://github.com/go-gl/mathgl) - Pure Go math package specialized for 3D math, with inspiration from GLM. Stars:`499`.
- [goxjs/gl](https://github.com/goxjs/gl) - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). Stars:`165`.
- [goxjs/glfw](https://github.com/goxjs/glfw) - Go cross-platform glfw library for creating an OpenGL context and receiving events. Stars:`78`.
- [go-glmatrix](https://github.com/technohippy/go-glmatrix) - Go port of [glMatrix](https://glmatrix.net/) library. Stars:`7`.

**[⬆ back to top](#contents)**

## ORM

_Libraries that implement Object-Relational Mapping or datamapping techniques._

- [GORM](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang, aims to be developer friendly. Stars:`33.2K`.
- [ent](https://github.com/facebook/ent) - An entity framework for Go. Simple, yet powerful ORM for modeling and querying data. Stars:`13.9K`.
- [SQLBoiler](https://github.com/volatiletech/sqlboiler) - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. Stars:`5.9K`.
- [gorp](https://github.com/go-gorp/gorp) - Go Relational Persistence, ORM-ish library for Go. Stars:`3.7K`.
- [upper.io/db](https://github.com/upper/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. Stars:`3.4K`.
- [bun](https://github.com/uptrace/bun) - SQL-first Golang ORM. Successor of go-pg. Stars:`2.3K`.
- [gormt](https://github.com/xxjwxc/gormt) - Mysql database to golang gorm struct. Stars:`2.2K`.
- [Prisma](https://github.com/prisma/prisma-client-go) - Prisma Client Go, Typesafe database access for Go. Stars:`1.6K`.
- [reform](https://github.com/go-reform/reform) - Better ORM for Go, based on non-empty interfaces and code generation. Stars:`1.4K`.
- [pop/soda](https://github.com/gobuffalo/pop) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. Stars:`1.4K`.
- [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) - A flexible and powerful SQL string builder library plus a zero-config ORM. Stars:`1.1K`.
- [rel](https://github.com/go-rel/rel) - Modern Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API. Stars:`656`.
- [Zoom](https://github.com/albrow/zoom) - Blazing-fast datastore and querying engine built on Redis. Stars:`300`.
- [go-sql](https://github.com/rushteam/gosql) - A easy ORM for mysql. Stars:`175`.
- [grimoire](https://github.com/Fs02/grimoire) - Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). Stars:`159`.
- [golobby/orm](https://github.com/golobby/orm) - Simple, fast, type-safe, generic orm for developer happiness. Stars:`136`.
- [go-store](https://github.com/gosuri/go-store) - Simple and fast Redis backed key-value store library for Go. Stars:`111`.
- [go-firestorm](https://github.com/jschoedt/go-firestorm) - A simple ORM for Google/Firebase Cloud Firestore. Stars:`46`.
- [cacheme](https://github.com/Yiling-J/cacheme-go) - Schema based, typed Redis caching/memoize framework for Go. Stars:`22`.
- [marlow](https://github.com/marlow/marlow) - Generated ORM from project structs for compile time safety assurances. Stars:`13`.
- [lore](https://github.com/abrahambotros/lore) - Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. Stars:`12`.
- [XORM](https://gitea.com/xorm/xorm) - Simple and powerful ORM for Go. (Support: MySQL, MyMysql, PostgreSQL, Tidb, SQLite3, MsSql and Oracle).

**[⬆ back to top](#contents)**

## Package Management

_Official tooling for dependency and package management_

- [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) - Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

_Official experimental tooling for package management_

- [dep](https://github.com/golang/dep) - Go dependency tool. Stars:`13.0K`.
- [vgo](https://go.googlesource.com/vgo/) - Versioned Go.

_Unofficial libraries for package and dependency management._

- [glide](https://github.com/Masterminds/glide) - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. Stars:`8.2K`.
- [godep](https://github.com/tools/godep) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. Stars:`5.6K`.
- [govendor](https://github.com/kardianos/govendor) - Go Package Manager. Go vendor tool that works with the standard vendor file. Stars:`4.9K`.
- [gopm](https://github.com/gpmgo/gopm) - Go Package Manager. Stars:`2.5K`.
- [gom](https://github.com/mattn/gom) - Go Manager - bundle for go. Stars:`1.4K`.
- [gpm](https://github.com/pote/gpm) - Barebones dependency manager for Go. Stars:`1.2K`.
- [goop](https://github.com/nitrous-io/goop) - Simple dependency manager for Go (golang), inspired by Bundler. Stars:`780`.
- [modgv](https://github.com/lucasepe/modgv) - Converts 'go mod graph' output into Graphviz's DOT language. Stars:`464`.
- [nut](https://github.com/jingweno/nut) - Vendor Go dependencies. Stars:`235`.
- [johnny-deps](https://github.com/VividCortex/johnny-deps) - Minimal dependency version using Git. Stars:`214`.
- [gup](https://github.com/nao1215/gup) - Update binaries installed by "go install". Stars:`201`.
- [mvn-golang](https://github.com/raydac/mvn-golang) - plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. Stars:`157`.
- [VenGO](https://github.com/DamnWidget/VenGO) - create and manage exportable isolated go virtual environments. Stars:`122`.
- [gop](https://github.com/lunny/gop) - Build and manage your Go applications out of GOPATH. Stars:`50`.

**[⬆ back to top](#contents)**

## Performance

- [jaeger](https://github.com/jaegertracing/jaeger) - A distributed tracing system. Stars:`18.0K`.
- [pixie](https://github.com/pixie-labs/pixie) - No instrumentation tracing for Golang applications via eBPF. Stars:`4.8K`.
- [statsviz](https://github.com/arl/statsviz) - Live visualization of your Go application runtime statistics. Stars:`2.9K`.
- [profile](https://github.com/pkg/profile) - Simple profiling support package for Go. Stars:`1.9K`.
- [go-instrument](https://github.com/nikolaydubina/go-instrument) - Automatically add spans to all methods and functions. Stars:`93`.
- [tracer](https://github.com/kamilsk/tracer) - Simple, lightweight tracing. Stars:`79`.

**[⬆ back to top](#contents)**

## Query Language

- [graphql-go](https://github.com/graphql-go/graphql) - Implementation of GraphQL for Go. Stars:`9.3K`.
- [gqlgen](https://github.com/99designs/gqlgen) - go generate based graphql server library. Stars:`9.1K`.
- [graphql](https://github.com/neelance/graphql-go) - GraphQL server with a focus on ease of use. Stars:`4.5K`.
- [dasel](https://github.com/tomwright/dasel) - Query and update data structures using selectors from the command line. Comparable to jq/yq but supports JSON, YAML, TOML and XML with zero runtime dependencies. Stars:`4.5K`.
- [gojsonq](https://github.com/thedevsaddam/gojsonq) - A simple Go package to Query over JSON Data. Stars:`2.1K`.
- [rql](https://github.com/a8m/rql) - Resource Query Language for REST API. Stars:`294`.
- [jsonql](https://github.com/elgs/jsonql) - JSON query expression library in Golang. Stars:`268`.
- [jsonslice](https://github.com/bhmj/jsonslice) - Jsonpath queries with advanced filters. Stars:`79`.
- [graphql](https://github.com/tmc/graphql) - graphql parser + utilities. Stars:`57`.
- [rqp](https://github.com/timsolov/rest-query-parser) - Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query. Stars:`56`.
- [goven](https://github.com/SeldonIO/goven) - A drop-in query language for any database schema. Stars:`55`.
- [api-fu](https://github.com/ccbrown/api-fu) - Comprehensive GraphQL implementation. Stars:`50`.
- [straf](https://github.com/SonicRoshan/straf) - Easily Convert Golang structs to GraphQL objects. Stars:`34`.
- [jsonpath](https://github.com/AsaiYusuke/jsonpath) - A query library for retrieving part of JSON based on JSONPath syntax. Stars:`18`.
- [gws](https://github.com/Zaba505/gws) - Apollos' "GraphQL over Websocket" client and server implementation. Stars:`7`.
- [grapher](https://github.com/reaganiwadha/grapher) - A GraphQL field builder utilizing Go generics with extra utilities and features. Stars:`1`.

**[⬆ back to top](#contents)**

## Resource Embedding

- [statik](https://github.com/rakyll/statik) - Embeds static files into a Go executable. Stars:`3.6K`.
- [packr](https://github.com/gobuffalo/packr) - The simple and easy way to embed static files into Go binaries. Stars:`3.4K`.
- [go.rice](https://github.com/GeertJohan/go.rice) - go.rice is a Go package that makes working with resources such as HTML, JS, CSS, images, and templates very easy. Stars:`2.4K`.
- [vfsgen](https://github.com/shurcooL/vfsgen) - Generates a vfsdata.go file that statically implements the given virtual filesystem. Stars:`977`.
- [esc](https://github.com/mjibson/esc) - Embeds files into Go programs and provides http.FileSystem interfaces to them. Stars:`636`.
- [fileb0x](https://github.com/UnnoTed/fileb0x) - Simple tool to embed files in go with focus on "customization" and ease to use. Stars:`630`.
- [go-resources](https://github.com/omeid/go-resources) - Unfancy resources embedding with Go. Stars:`176`.
- [statics](https://github.com/go-playground/statics) - Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. Stars:`66`.
- [templify](https://github.com/wlbr/templify) - Embed external template files into Go code to create single file binaries. Stars:`30`.
- [rebed](https://github.com/soypat/rebed) - Recreate folder structures and files from Go 1.16's `embed.FS` type Stars:`27`.
- [debme](https://github.com/leaanthony/debme) - Create an `embed.FS` from an existing `embed.FS` subdirectory. Stars:`25`.
- [mule](https://github.com/wlbr/mule) - Embed external resources like images, movies ... into Go source code to create single file binaries using `go generate`. Focused on simplicity. Stars:`12`.

**[⬆ back to top](#contents)**

## Science and Data Analysis

_Libraries for scientific computing and data analyzing._

- [gonum](https://github.com/gonum/gonum) - Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. Stars:`6.7K`.
- [stats](https://github.com/montanaflynn/stats) - Statistics package with common functions missing from the Golang standard library. Stars:`2.8K`.
- [gonum/plot](https://github.com/gonum/plot) - gonum/plot provides an API for building and drawing plots in Go. Stars:`2.5K`.
- [gosl](https://github.com/cpmech/gosl) - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. Stars:`1.8K`.
- [streamtools](https://github.com/nytlabs/streamtools) - general purpose, graphical tool for dealing with streams of data. Stars:`1.3K`.
- [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Dataframes for machine-learning and statistics (similar to pandas). Stars:`1.0K`.
- [go-dsp](https://github.com/mjibson/go-dsp) - Digital Signal Processing for Go. Stars:`825`.
- [chart](https://github.com/vdobler/chart) - Simple Chart Plotting library for Go. Supports many graphs types. Stars:`755`.
- [goraph](https://github.com/gyuho/goraph) - Pure Go graph theory library(data structure, algorithm visualization). Stars:`720`.
- [orb](https://github.com/paulmach/orb) - 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. Stars:`701`.
- [graph](https://github.com/yourbasic/graph) - Library of basic graph algorithms. Stars:`650`.
- [ewma](https://github.com/VividCortex/ewma) - Exponentially-weighted moving averages. Stars:`411`.
- [calendarheatmap](https://github.com/nikolaydubina/calendarheatmap) - Calendar heatmap in plain Go inspired by Github contribution activity. Stars:`376`.
- [TextRank](https://github.com/DavidBelicza/TextRank) - TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. Stars:`183`.
- [gohistogram](https://github.com/VividCortex/gohistogram) - Approximate histograms for data streams. Stars:`172`.
- [sparse](https://github.com/james-bowman/sparse) - Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. Stars:`146`.
- [go-estimate](https://github.com/milosgajdos/go-estimate) - State estimation and filtering algorithms in Go. Stars:`104`.
- [pagerank](https://github.com/alixaxel/pagerank) - Weighted PageRank algorithm implemented in Go. Stars:`77`.
- [jsonl-graph](https://github.com/nikolaydubina/jsonl-graph) - Tool to manipulate JSONL graphs with graphviz support. Stars:`68`.
- [geom](https://github.com/skelterjohn/geom) - 2D geometry for golang. Stars:`55`.
- [evaler](https://github.com/soniah/evaler) - Simple floating point arithmetic expression evaluator. Stars:`52`.
- [decimal](https://github.com/db47h/decimal) - Package decimal implements arbitrary-precision decimal floating-point arithmetic. Stars:`37`.
- [triangolatte](https://github.com/tchayen/triangolatte) - 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. Stars:`35`.
- [goent](https://github.com/kzahedi/goent) - GO Implementation of Entropy Measures. Stars:`33`.
- [gograph](https://github.com/hmdsefi/gograph) -  A golang generic graph library that provides mathematical graph-theory and algorithms. Stars:`32`.
- [piecewiselinear](https://github.com/sgreben/piecewiselinear) - Tiny linear interpolation library. Stars:`26`.
- [godesim](https://github.com/soypat/godesim) - Extended/multivariable ODE solver framework for event-based simulations with simple API. Stars:`21`.
- [ode](https://github.com/ChristopherRabotin/ode) - Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. Stars:`21`.
- [PiHex](https://github.com/claygod/PiHex) - Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. Stars:`20`.
- [GoStats](https://github.com/OGFris/GoStats) - GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. Stars:`19`.
- [assocentity](https://github.com/ndabAP/assocentity) - Package assocentity returns the average distance from words to a given entity. Stars:`12`.
- [rootfinding](https://github.com/khezen/rootfinding) - root-finding algorithms library for finding roots of quadratic functions. Stars:`10`.
- [go-gt](https://github.com/ThePaw/go-gt) - Graph theory algorithms written in "Go" language. Stars:`10`.
- [bradleyterry](https://github.com/seanhagen/bradleyterry) - Provides a Bradley-Terry Model for pairwise comparisons. Stars:`8`.

**[⬆ back to top](#contents)**

## Security

_Libraries that are used to help make your application more secure._

- [age](https://github.com/FiloSottile/age) - A simple, modern and secure encryption tool (and Go library) with small explicit keys, no config options, and UNIX-style composability. Stars:`13.9K`.
- [lego](https://github.com/go-acme/lego) - Pure Go ACME client library and CLI tool (for use with Let's Encrypt). Stars:`6.4K`.
- [CertMagic](https://github.com/caddyserver/certmagic) - Mature, robust, and powerful ACME client integration for fully-managed TLS certificate issuance and renewal. Stars:`4.6K`.
- [Cameradar](https://github.com/Ullaakut/cameradar) - Tool and library to remotely hack RTSP streams from surveillance cameras. Stars:`3.4K`.
- [memguard](https://github.com/awnumar/memguard) - A pure Go library for handling sensitive values in memory. Stars:`2.4K`.
- [secure](https://github.com/unrolled/secure) - HTTP middleware for Go that facilitates some quick security wins. Stars:`2.1K`.
- [acmetool](https://github.com/hlandau/acme) - ACME (Let's Encrypt) client tool with automatic renewal. Stars:`2.0K`.
- [themis](https://github.com/cossacklabs/themis) - high-level cryptographic library for solving typical data security tasks (secure data storage, secure messaging, zero-knowledge proof authentication), available for 14 languages, best fit for multi-platform apps. Stars:`1.7K`.
- [Coraza](https://github.com/corazawaf/coraza) - Enterprise-ready, modsecurity and OWASP CRS compatible WAF library. Stars:`1.3K`.
- [acra](https://github.com/cossacklabs/acra) - Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. Stars:`1.2K`.
- [dongle](https://github.com/golang-module/dongle) - A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption. Stars:`735`.
- [nacl](https://github.com/kevinburke/nacl) - Go implementation of the NaCL set of API's. Stars:`535`.
- [go-password-validator](https://github.com/lane-c-wagner/go-password-validator) - Password validator based on raw cryptographic entropy values. Stars:`418`.
- [ssh-vault](https://github.com/ssh-vault/ssh-vault) - encrypt/decrypt using ssh keys. Stars:`384`.
- [booster](https://github.com/anatol/booster) - Fast initramfs generator with full-disk encryption support. Stars:`365`.
- [optimus-go](https://github.com/pjebs/optimus-go) - ID hashing and Obfuscation using Knuth's Algorithm. Stars:`349`.
- [firewalld-rest](https://github.com/prashantgupta24/firewalld-rest) - A rest application to dynamically update firewalld rules on a linux server. Stars:`330`.
- [BadActor](https://github.com/jaredfolkins/badactor) - In-memory, application-driven jailer built in the spirit of fail2ban. Stars:`315`.
- [go-yara](https://github.com/hillu/go-yara) - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". Stars:`310`.
- [passlib](https://github.com/hlandau/passlib) - Futureproof password hashing library. Stars:`285`.
- [simple-scrypt](https://github.com/elithrar/simple-scrypt) - Scrypt package with a simple, obvious API and automatic cost calibration built-in. Stars:`189`.
- [teler-waf](https://github.com/kitabisa/teler-waf) - teler-waf is a Go HTTP middleware that provide teler IDS functionality to protect against web-based attacks and improve the security of Go-based web applications. It is highly configurable and easy to integrate into existing Go applications. Stars:`178`.
- [argon2pw](https://github.com/raja/argon2pw) - Argon2 password hash generation with constant-time password comparison. Stars:`89`.
- [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) - A probably paranoid package for securely hashing and encrypting passwords. Stars:`55`.
- [go-generate-password](https://github.com/m1/go-generate-password) - Password generator that can be used on the cli or as a library. Stars:`48`.
- [certificates](https://github.com/mvmaasakkers/certificates) - An opinionated tool for generating tls certificates. Stars:`37`.
- [go-htpasswd](https://github.com/tg123/go-htpasswd) - Apache htpasswd Parser for Go. Stars:`32`.
- [secureio](https://github.com/xaionaro-go/secureio) - An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519. Stars:`30`.
- [secret](https://github.com/rsjethani/secret) - Prevent your secrets from leaking into logs, std\* etc. Stars:`21`.
- [argon2-hashing](https://github.com/andskur/argon2-hashing) - light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package. Stars:`19`.
- [sslmgr](https://github.com/adrianosela/sslmgr) - SSL certificates made easy with a high level wrapper around acme/autocert. Stars:`18`.
- [goArgonPass](https://github.com/dwin/goArgonPass) - Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. Stars:`16`.
- [Interpol](https://github.com/avahidi/interpol) - Rule-based data generator for fuzzing and penetration testing. Stars:`2`.
- [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) - Auto provision Let's Encrypt certificates and start a TLS server.

**[⬆ back to top](#contents)**

## Serialization

_Libraries and tools for binary serialization._

- [jsoniter](https://github.com/json-iterator/go) - High-performance 100% compatible drop-in replacement of "encoding/json". Stars:`12.5K`.
- [goprotobuf](https://github.com/golang/protobuf) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. Stars:`9.2K`.
- [mapstructure](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures. Stars:`7.0K`.
- [gogoprotobuf](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets. Stars:`5.6K`.
- [go-codec](https://github.com/ugorji/go) - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. Stars:`1.8K`.
- [csvutil](https://github.com/jszwec/csvutil) - High Performance, idiomatic CSV record encoding and decoding to native Go structures. Stars:`819`.
- [colfer](https://github.com/pascaldekloe/colfer) - Code generation for the Colfer binary format. Stars:`712`.
- [cbor](https://github.com/fxamacker/cbor) - Small, safe, and easy CBOR encoding and decoding library. Stars:`592`.
- [go-capnproto](https://github.com/glycerine/go-capnproto) - Cap'n Proto library and parser for go. Stars:`289`.
- [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. Stars:`159`.
- [structomap](https://github.com/tuvistavie/structomap) - Library to easily and dynamically generate maps from static structures. Stars:`140`.
- [binstruct](https://github.com/ghostiam/binstruct) - Golang binary decoder for mapping data into the structure. Stars:`75`.
- [bambam](https://github.com/glycerine/bambam) - generator for Cap'n Proto schemas from go. Stars:`67`.
- [asn1](https://github.com/PromonLogicalis/asn1) - Asn.1 BER and DER encoding library for golang. Stars:`53`.
- [bel](https://github.com/32leaves/bel) - Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. Stars:`35`.
- [fwencoder](https://github.com/o1egl/fwencoder) - Fixed width file parser (encoding and decoding library) for Go. Stars:`25`.
- [elastic](https://github.com/epiclabs-io/elastic) - Convert slices, maps or any other unknown value across different types at run-time, no matter what. Stars:`24`.
- [pletter](https://github.com/vimeda/pletter) - A standard way to wrap a proto message for message brokers. Stars:`20`.
- [fixedwidth](https://github.com/huydang284/fixedwidth) - Fixed-width text formatting (UTF-8 supported). Stars:`8`.
- [unitpacking](https://github.com/recolude/unitpacking) - Library to pack unit vectors into as fewest bytes as possible. Stars:`6`.
- [go-lctree](https://github.com/sbourlon/go-lctree) - Provides a CLI and primitives to serialize and deserialize [LeetCode binary trees](https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation). Stars:`4`.

**[⬆ back to top](#contents)**

## Server Applications

- [Caddy](https://github.com/caddyserver/caddy) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use. Stars:`48.4K`.
- [etcd](https://github.com/coreos/etcd) - Highly-available key value store for shared configuration and service discovery. Stars:`44.1K`.
- [minio](https://github.com/minio/minio) - Minio is a distributed object storage server. Stars:`40.2K`.
- [RoadRunner](https://github.com/spiral/roadrunner) - High-performance PHP application server, load-balancer and process manager. Stars:`7.3K`.
- [SFTPGo](https://github.com/drakkan/sftpgo) - Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support. It can serve local filesystem and Cloud Storage backends such as S3 and Google Cloud Storage. Stars:`6.6K`.
- [Easegress](https://github.com/megaease/easegress) - A cloud native high availability/performance traffic orchestration system with observability and extensibility. Stars:`5.5K`.
- [devd](https://github.com/cortesi/devd) - Local webserver for developers. Stars:`3.3K`.
- [flipt](https://github.com/markphelps/flipt) - A self contained feature flag solution written in Go and Vue.js Stars:`2.6K`.
- [Fider](https://github.com/getfider/fider) - Fider is an open platform to collect and organize customer feedback. Stars:`2.4K`.
- [algernon](https://github.com/xyproto/algernon) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. Stars:`2.3K`.
- [Wish](https://github.com/charmbracelet/wish) - Make SSH apps, just like that! Stars:`2.2K`.
- [Flagr](https://github.com/checkr/flagr) - Flagr is an open-source feature flagging and A/B testing service. Stars:`2.2K`.
- [Trickster](https://github.com/tricksterproxy/trickster) - HTTP reverse proxy cache and time series accelerator. Stars:`1.9K`.
- [discovery](https://github.com/Bilibili/discovery) - A registry for resilient mid-tier load balancing and failover. Stars:`1.7K`.
- [jackal](https://github.com/ortuman/jackal) - An XMPP server written in Go. Stars:`1.4K`.
- [go-feature-flag](https://github.com/thomaspoignant/go-feature-flag) - A feature flag solution, with only a YAML file in the backend (S3, GitHub, HTTP, local file ...), no server to install, just add a file in a central system and refer to it. Stars:`766`.
- [Euterpe](https://github.com/ironsmile/euterpe) - Self-hosted music streaming server with built-in web UI and REST API. Stars:`489`.
- [dummy](https://github.com/neotoolkit/dummy) - Run mock server based off an API contract with one command. Stars:`174`.
- [dudeldu](https://github.com/krotik/dudeldu) - A simple SHOUTcast server. Stars:`140`.
- [go-proxy-cache](https://github.com/fabiocicerchia/go-proxy-cache) - Simple Reverse Proxy with Caching, written in Go, using Redis. Stars:`94`.
- [lets-proxy2](https://github.com/rekby/lets-proxy2) - Reverse proxy for handle https with issue certificates in fly from lets-encrypt. Stars:`80`.
- [cortex-tenant](https://github.com/blind-oracle/cortex-tenant) - Prometheus remote write proxy that adds add Cortex tenant ID header based on metric labels. Stars:`73`.
- [psql-streamer](https://github.com/blind-oracle/psql-streamer) - Stream database events from PostgreSQL to Kafka. Stars:`52`.
- [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) - Nginx log parser and exporter to Prometheus. Stars:`37`.
- [simple-jwt-provider](https://github.com/leberKleber/simple-jwt-provider) - Simple and lightweight provider which exhibits JWTs, supports login, password-reset (via mail) and user management. Stars:`33`.
- [protoxy](https://github.com/camgraff/protoxy) - A proxy server that converts JSON request bodies to Protocol Buffers. Stars:`31`.
- [Moxy](https://github.com/sinhashubham95/moxy) - Moxy is a simple mocker and proxy application server, you can create mock endpoints as well as proxy requests in case no mock exists for the endpoint. Stars:`9`.
- [riemann-relay](https://github.com/blind-oracle/riemann-relay) - Relay to load-balance Riemann events and/or convert them to Carbon. Stars:`2`.
- [nsq](https://nsq.io/) - A realtime distributed messaging platform.
- [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.

**[⬆ back to top](#contents)**

## Stream Processing

_Libraries and tools for stream processing and reactive programming._

- [go-streams](https://github.com/reugn/go-streams) - Go stream processing library. Stars:`1.5K`.
- [machine](https://github.com/whitaker-io/machine) - Go library for writing and generating stream workers with built in metrics and traceability. Stars:`133`.
- [stream](https://github.com/youthlin/stream) - Go Stream, like Java 8 Stream: Filter/Map/FlatMap/Peek/Sorted/ForEach/Reduce... Stars:`80`.
- [goio](https://github.com/primetalk/goio) - An implementation of IO, Stream, Fiber for Golang, inspired by awesome Scala libraries cats and fs2. Stars:`71`.

**[⬆ back to top](#contents)**

## Template Engines

_Libraries and tools for templating and lexing._

- [sprig](https://github.com/Masterminds/sprig) - Useful template functions for Go templates. Stars:`3.7K`.
- [quicktemplate](https://github.com/valyala/quicktemplate) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. Stars:`2.8K`.
- [pongo2](https://github.com/flosch/pongo2) - Django-like template-engine for Go. Stars:`2.6K`.
- [jet](https://github.com/CloudyKit/jet) - Jet template engine. Stars:`1.1K`.
- [maroto](https://github.com/johnfercher/maroto) - A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. Stars:`1.0K`.
- [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang. Stars:`824`.
- [fasttemplate](https://github.com/valyala/fasttemplate) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](https://golang.org/pkg/text/template/). Stars:`733`.
- [ego](https://github.com/benbjohnson/ego) - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. Stars:`561`.
- [raymond](https://github.com/aymerick/raymond) - Complete handlebars implementation in Go. Stars:`543`.
- [goview](https://github.com/foolin/goview) - Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. Stars:`357`.
- [liquid](https://github.com/osteele/liquid) - Go implementation of Shopify Liquid templates. Stars:`221`.
- [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). Stars:`169`.
- [extemplate](https://github.com/dannyvankooten/extemplate) - Tiny wrapper around html/template to allow for easy file-based template inheritance. Stars:`53`.
- [gospin](https://github.com/m1/gospin) - Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations. Stars:`46`.
- [tbd](https://github.com/lucasepe/tbd) - A really simple way to create text templates with placeholders - exposes extra builtin Git repo metadata. Stars:`24`.
- [got](https://github.com/goradd/got) - A Go code generator inspired by Hero and Fasttemplate. Has include files, custom tag definitions, injected Go code, language translation, and more. Stars:`14`.

**[⬆ back to top](#contents)**

## Testing

_Libraries for testing codebases and generating test data._

- Testing Frameworks

  - [Testify](https://github.com/stretchr/testify) - Sacred extension to the standard go testing package. Stars:`20.3K`.
  - [go-cmp](https://github.com/google/go-cmp) - Package for comparing Go values in tests. Stars:`3.6K`.
  - [testcontainers-go](https://github.com/testcontainers/testcontainers-go) - A Go package that makes it simple to create and clean up container-based dependencies for automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers that should be run as part of a test and clean up those resources when the test is done. Stars:`2.3K`.
  - [httpexpect](https://github.com/gavv/httpexpect) - Concise, declarative, and easy to use end-to-end HTTP and REST API testing. Stars:`2.3K`.
  - [godog](https://github.com/cucumber/godog) - Cucumber BDD framework for Go. Stars:`2.0K`.
  - [is](https://github.com/matryer/is) - Professional lightweight testing mini-framework for Go. Stars:`1.6K`.
  - [gnomock](https://github.com/orlangure/gnomock) - integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks. Stars:`1.2K`.
  - [go-vcr](https://github.com/dnaeon/go-vcr) - Record and replay your HTTP interactions for fast, deterministic and accurate tests. Stars:`1.0K`.
  - [testfixtures](https://github.com/go-testfixtures/testfixtures) - A helper for Rails' like test fixtures to test database applications. Stars:`968`.
  - [goblin](https://github.com/franela/goblin) - Mocha like testing framework fo Go. Stars:`889`.
  - [baloo](https://github.com/h2non/baloo) - Expressive and versatile end-to-end HTTP API testing made easy. Stars:`762`.
  - [goc](https://github.com/qiniu/goc) - Goc is a comprehensive coverage testing system for The Go Programming Language. Stars:`699`.
  - [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) - Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. Stars:`622`.
  - [go-mutesting](https://github.com/zimmski/go-mutesting) - Mutation testing for Go source code. Stars:`593`.
  - [gotest.tools](https://github.com/gotestyourself/gotest.tools) - A collection of packages to augment the go testing package and support common patterns. Stars:`447`.
  - [gofight](https://github.com/appleboy/gofight) - API Handler Testing for Golang Router framework. Stars:`436`.
  - [testza](https://github.com/MarvinJWendt/testza) - Full-featured test framework with nice colorized output. Stars:`411`.
  - [go-testdeep](https://github.com/maxatome/go-testdeep) - Extremely flexible golang deep comparison, extends the go testing package. Stars:`396`.
  - [frisby](https://github.com/verdverm/frisby) - REST API testing framework. Stars:`274`.
  - [cupaloy](https://github.com/bradleyjkemp/cupaloy) - Simple snapshot testing addon for your test framework. Stars:`272`.
  - [got](https://github.com/ysmood/got) - An enjoyable golang test framework. Stars:`263`.
  - [endly](https://github.com/viant/endly) - Declarative end to end functional testing. Stars:`244`.
  - [go-carpet](https://github.com/msoap/go-carpet) - Tool for viewing test coverage in terminal. Stars:`240`.
  - [go-hit](https://github.com/Eun/go-hit) - Hit is an http integration test framework written in golang. Stars:`230`.
  - [commander](https://github.com/SimonBaeumer/commander) - Tool for testing cli applications on windows, linux and osx. Stars:`213`.
  - [charlatan](https://github.com/percolate/charlatan) - Tool to generate fake interface implementations for tests. Stars:`199`.
  - [dbcleaner](https://github.com/khaiql/dbcleaner) - Clean database for testing purpose, inspired by `database_cleaner` in Ruby. Stars:`156`.
  - [jsonassert](https://github.com/kinbiko/jsonassert) - Package for verifying that your JSON payloads are serialized correctly. Stars:`113`.
  - [GoSpec](https://github.com/orfjackal/gospec) - BDD-style testing framework for the Go programming language. Stars:`113`.
  - [testcase](https://github.com/adamluzsi/testcase) - Idiomatic testing framework for Behavior Driven Development. Stars:`111`.
  - [wstest](https://github.com/posener/wstest) - Websocket client for unit-testing a websocket http.Handler. Stars:`100`.
  - [gocrest](https://github.com/corbym/gocrest) - Composable hamcrest-like matchers for Go assertions. Stars:`94`.
  - [be](https://github.com/carlmjohnson/be) - The minimalist generic test assertion library. Stars:`74`.
  - [gherkingen](https://github.com/hedhyw/gherkingen) - BDD boilerplate generator and framework. Stars:`60`.
  - [Gont](https://github.com/stv0g/gont) - Go network testing toolkit for testing building complex network topologies using Linux namespaces. Stars:`59`.
  - [assert](https://github.com/go-playground/assert) - Basic Assertion Library used along side native go testing, with building blocks for custom assertions. Stars:`59`.
  - [covergates](https://github.com/covergates/covergates) - Self-hosted code coverage report review and management service. Stars:`57`.
  - [restit](https://github.com/yookoala/restit) - Go micro framework to help writing RESTful API integration test. Stars:`56`.
  - [gospecify](https://github.com/stesla/gospecify) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. Stars:`53`.
  - [gomatch](https://github.com/jfilipczyk/gomatch) - library created for testing JSON against patterns. Stars:`46`.
  - [dsunit](https://github.com/viant/dsunit) - Datastore testing for SQL, NoSQL, structured files. Stars:`43`.
  - [fluentassert](https://github.com/fluentassert/verify) - Extensible, type-safe, fluent assertion Go library. Stars:`34`.
  - [Hamcrest](https://github.com/rdrdr/hamcrest) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. Stars:`28`.
  - [schema](https://github.com/jgroeneveld/schema) - Quick and easy expression matching for JSON schemas used in requests and responses. Stars:`20`.
  - [fixenv](https://github.com/rekby/fixenv) - Fixture manage engine, inspired by pytest fixtures. Stars:`19`.
  - [flute](https://github.com/suzuki-shunsuke/flute) - HTTP client testing framework. Stars:`17`.
  - [testsql](https://github.com/zhulongcheng/testsql) - Generate test data from SQL files before testing and clear it after finished. Stars:`16`.
  - [gogiven](https://github.com/corbym/gogiven) - YATSPEC-like BDD testing framework for Go. Stars:`16`.
  - [gosuite](https://github.com/pavlo/gosuite) - Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. Stars:`12`.
  - [biff](https://github.com/fulldump/biff) - Bifurcation testing framework, BDD compatible. Stars:`12`.
  - [badio](https://github.com/cavaliercoder/badio) - Extensions to Go's `testing/iotest` package. Stars:`10`.
  - [stop-and-go](https://github.com/elgohr/stop-and-go) - Testing helper for concurrency. Stars:`9`.
  - [Tt](https://github.com/vcaesar/tt) - Simple and colorful test tools. Stars:`6`.
  - [trial](https://github.com/jgroeneveld/trial) - Quick and easy extendable assertions without introducing much boilerplate. Stars:`6`.
  - [go-testpredicate](https://github.com/maargenton/go-testpredicate) - Test predicate style assertions library with extensive diagnostics output. Stars:`5`.
  - [go-mysql-test-container](https://github.com/arikama/go-mysql-test-container) - Golang MySQL testcontainer to help with MySQL integration testing. Stars:`2`.
  - [omg.testingtools](https://github.com/dedalqq/omg.testingtools) - The simple library for change a values of private fields for testing. Stars:`1`.
  - [go-snaps](http://github.com/gkampitakis/go-snaps) - Jest-like snapshot testing in Golang.
  - [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd) - Convert markdown snippets into testable go code.
  - [gomega](https://onsi.github.io/gomega/) - Rspec like matcher/assertion library.
  - [apitest](https://apitest.dev) - Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
  - [ginkgo](https://onsi.github.io/ginkgo/) - BDD Testing Framework for Go.
  - [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload.
  - [gocheck](https://labix.org/gocheck) - More advanced testing framework alternative to gotest.

- Mock

  - [gomock](https://github.com/golang/mock) - Mocking framework for the Go programming language. Stars:`9.0K`.
  - [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database interactions. Stars:`5.4K`.
  - [mockery](https://github.com/vektra/mockery) - Tool to generate Go interfaces. Stars:`4.9K`.
  - [hoverfly](https://github.com/SpectoLabs/hoverfly) - HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. Stars:`2.2K`.
  - [gock](https://github.com/h2non/gock) - Versatile HTTP mocking made easy. Stars:`1.9K`.
  - [httpmock](https://github.com/jarcoal/httpmock) - Easy mocking of HTTP responses from external resources. Stars:`1.7K`.
  - [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects. Stars:`855`.
  - [go-txdb](https://github.com/DATA-DOG/go-txdb) - Single transaction based database driver mainly for testing purposes. Stars:`555`.
  - [minimock](https://github.com/gojuno/minimock) - Mock generator for Go interfaces. Stars:`506`.
  - [govcr](https://github.com/seborama/govcr) - HTTP mock for Golang: record and replay HTTP interactions for offline testing. Stars:`154`.
  - [go-localstack](https://github.com/elgohr/go-localstack) - Tool for using localstack in AWS testing. Stars:`71`.
  - [timex](https://github.com/cabify/timex) - A test-friendly replacement for the native `time` package. Stars:`69`.
  - [mockhttp](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter. Stars:`22`.
  - [mooncake](https://github.com/GuilhermeCaruso/mooncake) - A simple way to generate mocks for multiple purposes Stars:`17`.
  - [mockit](https://github.com/pasdam/mockit) - Allows functions and method easy mocking, without defining new types; it's similar to Mockito for Java. Stars:`16`.
  - [genmock](https://gitlab.com/so_literate/genmock) - Go mocking system with code generator for building calls of the interface methods.
  - [monkey](https://github.com/awterman/monkey) - One line to mock functions/methods/variables in place without dependency injection or code generation.

- Fuzzing and delta-debugging/reducing/shrinking.

  - [go-fuzz](https://github.com/dvyukov/go-fuzz) - Randomized testing system. Stars:`4.6K`.
  - [gofuzz](https://github.com/google/gofuzz) - Library for populating go objects with random values. Stars:`1.4K`.
  - [Tavor](https://github.com/zimmski/tavor) - Generic fuzzing and delta-debugging framework. Stars:`242`.

- Selenium and browser control tools.

  - [chromedp](https://github.com/knq/chromedp) - a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. Stars:`9.4K`.
  - [rod](https://github.com/go-rod/rod) - A Devtools driver to make web automation and scraping easy. Stars:`4.0K`.
  - [selenoid](https://github.com/aerokube/selenoid) - alternative Selenium hub server that launches browsers within containers. Stars:`2.4K`.
  - [playwright-go](https://github.com/mxschmitt/playwright-go) - browser automation library to control Chromium, Firefox and WebKit with a single API. Stars:`1.4K`.
  - [cdp](https://github.com/mafredri/cdp) - Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. Stars:`679`.
  - [ggr](https://github.com/aerokube/ggr) - a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. Stars:`302`.

- Fail injection
  - [failpoint](https://github.com/pingcap/failpoint) - An implementation of [failpoints](https://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. Stars:`782`.

**[⬆ back to top](#contents)**

## Text Processing

_Libraries for parsing and manipulating texts._

See also [Natural Language Processing](#natural-language-processing) and [Text Analysis](#text-analysis).

### Formatters

- [go-humanize](https://github.com/dustin/go-humanize) - Formatters for time, numbers, and memory size to human readable format. Stars:`3.7K`.
- [gotabulate](https://github.com/bndr/gotabulate) - Easily pretty-print your tabular data with Go. Stars:`298`.
- [align](https://github.com/Guitarbum722/align) - A general purpose application that aligns text. Stars:`80`.
- [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) - Fixed-width text formatting (encoder/decoder with reflection). Stars:`75`.
- [address](https://github.com/bojanz/address) - Handles address representation, validation and formatting. Stars:`67`.
- [textwrap](https://github.com/isbm/textwrap) - Wraps text at end of lines. Implementation of `textwrap` module from Python. Stars:`4`.
- [bytes](https://github.com/labstack/gommon/tree/master/bytes) - Formats and parses numeric byte values (10K, 2M, 3G, etc.).

### Markup Languages

- [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go. Stars:`5.2K`.
- [toml](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection). Stars:`4.2K`.
- [goldmark](https://github.com/yuin/goldmark) - A Markdown parser written in Go. Easy to extend, standard (CommonMark) compliant, well structured. Stars:`2.8K`.
- [go-toml](https://github.com/pelletier/go-toml) - Go library for the TOML format with query support and handy cli tools. Stars:`1.5K`.
- [htmlquery](https://github.com/antchfx/htmlquery) - An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. Stars:`622`.
- [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) - Convert HTML to Markdown. Even works with entire websites and can be extended through rules. Stars:`579`.
- [mxj](https://github.com/clbanning/mxj) - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. Stars:`568`.
- [goq](https://github.com/andrewstuart/goq) - Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). Stars:`243`.
- [bafi](https://github.com/mmalcek/bafi) - Universal JSON, BSON, YAML, XML translator to ANY format using templates. Stars:`78`.
- [go-output-format](https://github.com/drewstinnett/go-output-format) - Output go structures into multiple formats (YAML/JSON/etc) in your command line app. Stars:`11`.
- [bbConvert](https://github.com/CalebQ42/bbConvert) - Converts bbCode to HTML that allows you to add support for custom bbCode tags. Stars:`7`.
- [htmlyaml](https://github.com/nikolaydubina/htmlyaml) -  Rich rendering of YAML as HTML in Go Stars:`2`.
- [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.

### Parsers/Encoders/Decoders

- [sh](https://github.com/mvdan/sh) - Shell parser and formatter. Stars:`6.1K`.
- [gofeed](https://github.com/mmcdole/gofeed) - Parse RSS and Atom feeds in Go. Stars:`2.2K`.
- [when](https://github.com/olebedev/when) - Natural EN and RU language date/time parser with pluggable rules. Stars:`1.3K`.
- [commonregex](https://github.com/mingrammer/commonregex) - A collection of common regular expressions for Go. Stars:`864`.
- [omniparser](https://github.com/jf-tech/omniparser) - A versatile ETL library that parses text input (CSV/txt/JSON/XML/EDI/X12/EDIFACT/etc) in streaming fashion and transforms data into JSON output using data-driven schema. Stars:`588`.
- [gographviz](https://github.com/awalterschulze/gographviz) - Parses the Graphviz DOT language. Stars:`526`.
- [go-nmea](https://github.com/adrianmo/go-nmea) - NMEA parser library for the Go language. Stars:`197`.
- [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) - Editorconfig file parser and manipulator for Go. Stars:`120`.
- [sdp](https://github.com/gortc/sdp) - SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. Stars:`112`.
- [go-vcard](https://github.com/emersion/go-vcard) - Parse and format vCard. Stars:`89`.
- [did](https://github.com/ockam-network/did) - DID (Decentralized Identifiers) Parser and Stringer in Go. Stars:`77`.
- [allot](https://github.com/sbstjn/allot) - Placeholder and wildcard text parsing for CLI tools and bots. Stars:`57`.
- [tokenizer](https://github.com/bzick/tokenizer) - Parse any string, slice or infinite buffer to any tokens. Stars:`48`.
- [parth](https://github.com/codemodus/parth) - URL path segmentation parsing. Stars:`43`.
- [gonameparts](https://github.com/polera/gonameparts) - Parses human names into individual name parts. Stars:`38`.
- [normalize](https://github.com/avito-tech/normalize) - Sanitize, normalize and compare fuzzy text. Stars:`36`.
- [xj2go](https://github.com/stackerzzq/xj2go) - Convert xml or json to go struct. Stars:`31`.
- [codetree](https://github.com/aerogo/codetree) - Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. Stars:`22`.
- [go-fasttld](https://github.com/elliotwutingfeng/go-fasttld) - High performance effective top level domains (eTLD) extraction module. Stars:`18`.
- [parseargs-go](https://github.com/nproc/parseargs-go) - string argument parser that understands quotes and backslashes. Stars:`9`.
- [ltsv](https://github.com/Wing924/ltsv) - High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go. Stars:`8`.
- [doi](https://github.com/hscells/doi) - Document object identifier (doi) parser in Go. Stars:`8`.
- [encdec](https://github.com/mickep76/encdec) - Package provides a generic interface to encoders and decoders. Stars:`8`.

### Regular Expressions

- [rex](https://github.com/hedhyw/rex) - Regular expressions builder. Stars:`157`.
- [regroup](https://github.com/oriser/regroup) - Match regex expression named groups into go struct using struct tags and automatic parsing. Stars:`132`.
- [goregen](https://github.com/zach-klippenstein/goregen) - Library for generating random strings from regular expressions. Stars:`84`.
- [genex](https://github.com/alixaxel/genex) - Count and expand Regular Expressions into all matching Strings. Stars:`75`.
- [go-wildcard](https://github.com/IGLOU-EU/go-wildcard) - Simple and lightweight wildcard pattern matching. Stars:`48`.

### Sanitation

- [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer. Stars:`2.8K`.
- [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) - A sanitization-based swear filter for Go. Stars:`57`.

### Scrapers

- [colly](https://github.com/asciimoo/colly) - Fast and Elegant Scraping Framework for Gophers. Stars:`20.1K`.
- [GoQuery](https://github.com/PuerkitoBio/goquery) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language. Stars:`12.8K`.
- [xurls](https://github.com/mvdan/xurls) - Extract urls from text. Stars:`1.1K`.
- [dataflowkit](https://github.com/slotix/dataflowkit) - Web scraping Framework to turn websites into structured data. Stars:`603`.
- [gospider](https://github.com/zhshch2002/gospider) - A simple golang spider/scraping framework,build a spider in 3 lines. migrated from [goribot](https://github.com/zhshch2002/goribot) Stars:`196`.
- [pagser](https://github.com/foolin/pagser) - Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler. Stars:`85`.
- [Tagify](https://github.com/zoomio/tagify) - Produces a set of tags from given source. Stars:`33`.
- [go-recipe](https://github.com/kkyr/go-recipe) - A package for scraping recipes from websites. Stars:`22`.
- [walker](https://github.com/cyucelen/walker) - Seamlessly fetch paginated data from any source. Simple and high performance API scraping included. Stars:`5`.

### RSS

- [podcast](https://github.com/eduncan911/podcast) - iTunes Compliant and RSS 2.0 Podcast Generator in Golang Stars:`125`.
- [syndfeed](https://github.com/zhengchun/syndfeed) - A syndication feed for Atom 1.0 and RSS 2.0. Stars:`11`.

### Utility/Miscellaneous

- [go-runewidth](https://github.com/mattn/go-runewidth) - Functions to get fixed width of the character or string. Stars:`538`.
- [radix](https://github.com/yourbasic/radix) - Fast string sorting algorithm. Stars:`186`.
- [go-zero-width](https://github.com/trubitsyn/go-zero-width) - Zero-width character detection and removal for Go. Stars:`110`.
- [petrovich](https://github.com/striker2000/petrovich) - Petrovich is the library which inflects Russian names to given grammatical case. Stars:`44`.
- [kace](https://github.com/codemodus/kace) - Common case conversions covering common initialisms. Stars:`20`.
- [TySug](https://github.com/Dynom/TySug) - Alternative suggestions with respect to keyboard layouts. Stars:`17`.

**[⬆ back to top](#contents)**

## Third-party APIs

_Libraries for accessing third party APIs._

- [github](https://github.com/google/go-github) - Go library for accessing the GitHub REST API v3. Stars:`9.6K`.
- [aws-sdk-go](https://github.com/aws/aws-sdk-go) - The official AWS SDK for the Go programming language. Stars:`8.4K`.
- [go-openai](https://github.com/sashabaranov/go-openai) - OpenAI ChatGPT, DALL·E, Whisper API library for Go. Stars:`6.1K`.
- [slack](https://github.com/slack-go/slack) - Slack API in Go. Stars:`4.4K`.
- [discordgo](https://github.com/bwmarrin/discordgo) - Go bindings for the Discord Chat API. Stars:`4.2K`.
- [google](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go. Stars:`3.5K`.
- [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library. Stars:`3.3K`.
- [minio-go](https://github.com/minio/minio-go) - Minio Go Library for Amazon S3 compatible cloud storage. Stars:`2.0K`.
- [stripe](https://github.com/stripe/stripe-go) - Go client for the Stripe API. Stars:`1.8K`.
- [go-twitter](https://github.com/dghubble/go-twitter) - Go client library for the Twitter v1.1 APIs. Stars:`1.6K`.
- [go-jira](https://github.com/andygrunwald/go-jira) - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira) Stars:`1.3K`.
- [facebook](https://github.com/huandu/facebook) - Go Library that supports the Facebook Graph API. Stars:`1.2K`.
- [anaconda](https://github.com/ChimeraCoder/anaconda) - Go client library for the Twitter 1.1 API. Stars:`1.1K`.
- [githubql](https://github.com/shurcooL/githubql) - Go library for accessing the GitHub GraphQL API v4. Stars:`1.0K`.
- [webhooks](https://github.com/go-playground/webhooks) - Webhook receiver for GitHub and Bitbucket. Stars:`853`.
- [twitter-scraper](https://github.com/n0madic/twitter-scraper) - Scrape the Twitter Frontend API without authentication and limits. Stars:`639`.
- [paypal](https://github.com/logpacker/PayPal-Go-SDK) - Wrapper for PayPal payment API. Stars:`588`.
- [geo-golang](https://github.com/codingsince1985/geo-golang) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](https://developer.mapquest.com/documentation/geocoding-api/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](https://opencagedata.com/api), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. Stars:`473`.
- [lark](https://github.com/chyroc/lark) - [Feishu](https://open.feishu.cn/)/[Lark](https://open.larksuite.com/) Open API Go SDK, Support ALL Open API and Event Callback. Stars:`343`.
- [openaigo](https://github.com/otiai10/openaigo) - OpenAI GPT3/GPT3.5 ChatGPT API client library for Go. Stars:`267`.
- [ethrpc](https://github.com/onrik/ethrpc) - Go bindings for Ethereum JSON RPC API. Stars:`256`.
- [Trello](https://github.com/adlio/trello) - Go wrapper for the Trello API. Stars:`210`.
- [go-marathon](https://github.com/gambol99/go-marathon) - Go library for interacting with Mesosphere's Marathon PAAS. Stars:`201`.
- [go-lark](https://github.com/go-lark/lark) - An easy-to-use unofficial SDK for [Feishu](https://open.feishu.cn/) and [Lark](https://open.larksuite.com/) Open Platform. Stars:`159`.
- [wit-go](https://github.com/wit-ai/wit-go) - Go client for wit.ai HTTP API. Stars:`144`.
- [Medium](https://github.com/Medium/medium-sdk-go) - Golang SDK for Medium's OAuth2 API. Stars:`138`.
- [go-trending](https://github.com/andygrunwald/go-trending) - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. Stars:`135`.
- [pushover](https://github.com/gregdel/pushover) - Go wrapper for the Pushover API. Stars:`135`.
- [gostorm](https://github.com/jsgilmore/gostorm) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. Stars:`130`.
- [simples3](https://github.com/rhnvrm/simples3) - Simple no frills AWS S3 Library using REST with V4 Signing written in Go. Stars:`115`.
- [gosip](https://github.com/koltyakov/gosip) - Client library for SharePoint. Stars:`113`.
- [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP. Stars:`110`.
- [hipchat](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API. Stars:`104`.
- [cachet](https://github.com/andygrunwald/cachet) - Go client library for [Cachet (open source status page system)](https://cachethq.io/). Stars:`91`.
- [go-atlassian](https://github.com/ctreminiom/go-atlassian) - Go library for accessing the [Atlassian Cloud](https://www.atlassian.com/enterprise/cloud) services (Jira, Jira Service Management, Jira Agile, Confluence, Admin Cloud) Stars:`83`.
- [golang-tmdb](https://github.com/cyruzin/golang-tmdb) - Golang wrapper for The Movie Database API v3. Stars:`81`.
- [igdb](https://github.com/Henry-Sarabia/igdb) - Go client for the [Internet Game Database API](https://api.igdb.com/). Stars:`77`.
- [gogtrends](https://github.com/groovili/gogtrends) - Google Trends Unofficial API. Stars:`72`.
- [go-unsplash](https://github.com/hbagdi/go-unsplash) - Go client library for the [Unsplash.com](https://unsplash.com) API. Stars:`72`.
- [go-postman-collection](https://github.com/rbretecher/go-postman-collection) - Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia). Stars:`70`.
- [google-play-scraper](https://github.com/n0madic/google-play-scraper) - Get data from Google Play Store. Stars:`67`.
- [circleci](https://github.com/jszwedko/go-circleci) - Go client library for interacting with CircleCI's API. Stars:`64`.
- [airtable](https://github.com/mehanizm/airtable) - Go client library for the [Airtable API](https://airtable.com/api). Stars:`63`.
- [mixpanel](https://github.com/dukex/mixpanel) - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. Stars:`60`.
- [ynab](https://github.com/brunomvsouza/ynab.go) - Go wrapper for the YNAB API. Stars:`59`.
- [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) - Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). Stars:`57`.
- [clarifai](https://github.com/samuelcouch/clarifai) - Go client library for interfacing with the Clarifai API. Stars:`56`.
- [uptimerobot](https://github.com/bitfield/uptimerobot) - Go wrapper and command-line client for the Uptime Robot v2 API. Stars:`56`.
- [megos](https://github.com/andygrunwald/megos) - Client library for accessing an [Apache Mesos](https://mesos.apache.org/) cluster. Stars:`54`.
- [go-xkcd](https://github.com/nishanths/go-xkcd) - Go client for the xkcd API. Stars:`51`.
- [fcm](https://github.com/maddevsio/fcm) - Go library for Firebase Cloud Messaging. Stars:`50`.
- [gads](https://github.com/emiddleton/gads) - Google Adwords Unofficial API. Stars:`50`.
- [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) - Go MusicBrainz WS2 client library. Stars:`50`.
- [spotify](https://github.com/rapito/go-spotify) - Go Library to access Spotify WEB API. Stars:`47`.
- [golyrics](https://github.com/mamal72/golyrics) - Golyrics is a Go library to fetch music lyrics data from the Wikia website. Stars:`40`.
- [swag](https://github.com/zc2638/swag) - No comments, simple go wrapper to create swagger 2.0 compatible APIs. Support most routing frameworks, such as built-in, gin, chi, mux, echo, httprouter, fasthttp and more. Stars:`38`.
- [patreon-go](https://github.com/mxpv/patreon-go) - Go library for Patreon API. Stars:`36`.
- [go-myanimelist](https://github.com/nstratos/go-myanimelist) - Go client library for accessing the [MyAnimeList API](https://myanimelist.net/apiconfig/references/api/v2). Stars:`36`.
- [translate](https://github.com/poorny/translate) - Go online translation package. Stars:`33`.
- [lastpass-go](https://github.com/ansd/lastpass-go) - Go client library for the [LastPass](https://www.lastpass.com/) API. Stars:`31`.
- [gcm](https://github.com/Aorioli/gcm) - Go library for Google Cloud Messaging. Stars:`31`.
- [gami](https://github.com/bit4bit/gami) - Go library for Asterisk Manager Interface. Stars:`31`.
- [steam](https://github.com/sostronk/go-steam) - Go Library to interact with Steam game servers. Stars:`30`.
- [go-imgur](https://github.com/koffeinsource/go-imgur) - Go client library for [imgur](https://imgur.com) Stars:`25`.
- [GoFreeDB](https://github.com/FreeLeh/GoFreeDB) - Golang library providing common and simple database abstractions on top of Google Sheets. Stars:`25`.
- [shopify](https://github.com/rapito/go-shopify) - Go Library to make CRUD request to the Shopify API. Stars:`25`.
- [device-check-go](https://github.com/rinchsan/device-check-go) - Go client library for interacting with [iOS DeviceCheck API](https://developer.apple.com/documentation/devicecheck) v1. Stars:`21`.
- [jokeapi-go](https://github.com/icelain/jokeapi) - Go client for [JokeAPI](https://sv443.net/jokeapi/v2/). Stars:`21`.
- [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) - Go client library for interacting with Coinpaprika's API. Stars:`20`.
- [brewerydb](https://github.com/naegelejd/brewerydb) - Go library for accessing the BreweryDB API. Stars:`19`.
- [codeship-go](https://github.com/codeship/codeship-go) - Go client library for interacting with Codeship's API v2. Stars:`19`.
- [textbelt](https://github.com/dietsche/textbelt) - Go client for the textbelt.com txt messaging API. Stars:`19`.
- [go-aws-news](https://github.com/circa10a/go-aws-news) - Go application and library to fetch what's new from AWS. Stars:`17`.
- [go-hacknews](https://github.com/PaulRosset/go-hacknews) - Tiny Go client for HackerNews API. Stars:`16`.
- [google-analytics](https://github.com/chonthu/go-google-analytics) - Simple wrapper for easy google analytics reporting. Stars:`15`.
- [gopaapi5](https://github.com/utekaravinash/gopaapi5) - Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/). Stars:`14`.
- [go-sophos](https://github.com/esurdam/go-sophos) - Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. Stars:`13`.
- [go-openproject](https://github.com/manuelbcd/go-openproject) - Go client library for interacting with [OpenProject](https://docs.openproject.org/api/) API. Stars:`13`.
- [bqwriter](https://github.com/OTA-Insight/bqwriter) - High Level Go Library to write data into [Google BigQuery](https://cloud.google.com/bigquery) at a high throughout. Stars:`12`.
- [go-here](https://github.com/abdullahselek/go-here) - Go client library around the HERE location based APIs. Stars:`12`.
- [goami2](https://github.com/staskobzar/goami2) - AMI v2 library for Asterisk PBX. Stars:`12`.
- [go-swagger-ui](https://github.com/esurdam/go-swagger-ui) - Go library containing precompiled [Swagger UI](https://swagger.io/tools/swagger-ui/) for serving swagger json. Stars:`11`.
- [gomalshare](https://github.com/MonaxGT/gomalshare) - Go library MalShare API [malshare.com](https://www.malshare.com/) Stars:`11`.
- [smite](https://github.com/sergiotapia/smitego) - Go package to wraps access to the Smite game API. Stars:`11`.
- [tumblr](https://github.com/mattcunningham/gumblr) - Go wrapper for the Tumblr v2 API. Stars:`9`.
- [rrdaclient](https://github.com/Omie/rrdaclient) - Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. Stars:`9`.
- [goagi](https://github.com/staskobzar/goagi) - Go library to build Asterisk PBX agi/fastagi applications. Stars:`9`.
- [libgoffi](https://github.com/clevabit/libgoffi) - Library adapter toolbox for native [libffi](https://sourceware.org/libffi/) integration Stars:`9`.
- [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) - Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). Stars:`8`.
- [go-sptrans](https://github.com/sergioaugrod/go-sptrans) - Go client library for the SPTrans Olho Vivo API. Stars:`8`.
- [rawg-sdk-go](https://github.com/dimuska139/rawg-sdk-go) - Go library for the [RAWG Video Games Database](https://rawg.io/) API Stars:`8`.
- [go-chronos](https://github.com/axelspringer/go-chronos) - Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler Stars:`8`.
- [zooz](https://github.com/gojuno/go-zooz) - Go client for the Zooz API. Stars:`7`.
- [go-hibp](https://github.com/wneessen/go-hibp) - Simple Go binding to the "Have I Been Pwned" APIs. Stars:`7`.
- [newsapi-go](https://github.com/jellydator/newsapi-go) - Go client for [NewsAPI](https://newsapi.org/). Stars:`6`.
- [go-restcountries](https://github.com/chriscross0/go-restcountries) - Go library for the [REST Countries API](https://countrylayer.com/). Stars:`3`.
- [dusupay-sdk-go](https://github.com/Kachit/dusupay-sdk-go) - Unofficial Dusupay payment gateway API Client for Go Stars:`3`.
- [appstore-sdk-go](https://github.com/Kachit/appstore-sdk-go) - Unofficial Golang SDK for AppStore Connect API. Stars:`3`.
- [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) - The Playlyfe Rest API Go SDK. Stars:`2`.
- [vl-go](https://github.com/verifid/vl-go) - Go client library around the VerifID identity verification layer API. Stars:`2`.
- [fasapay-sdk-go](https://github.com/Kachit/fasapay-sdk-go) - Unofficial Fasapay payment gateway XML API Client for Golang. Stars:`2`.
- [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) - Go wrapper for the TripAdvisor API. Stars:`2`.
- [go-yapla](https://git.iglou.eu/Production/go-yapla) - Go client library for the Yapla v2.0 API.
- [go-telegraph](https://gitlab.com/toby3d/telegraph) - Telegraph publishing platform API client.

**[⬆ back to top](#contents)**

## Utilities

_General utilities and tools to make your life easier._

- [fzf](https://github.com/junegunn/fzf) - Command-line fuzzy finder written in Go. Stars:`53.6K`.
- [dive](https://github.com/wagoodman/dive) - A tool for exploring each layer in a Docker image. Stars:`37.7K`.
- [hub](https://github.com/github/hub) - wrap git commands with additional functionality to interact with github from the terminal. Stars:`22.5K`.
- [ctop](https://github.com/bcicen/ctop) - [Top-like](https://ctop.sh) interface (e.g. htop) for container metrics. Stars:`14.5K`.
- [sqlx](https://github.com/jmoiron/sqlx) - provides a set of extensions on top of the excellent built-in database/sql package. Stars:`14.0K`.
- [lo](https://github.com/samber/lo) - A Lodash like Go library based on Go 1.18+ Generics (map, filter, contains, find...) Stars:`12.7K`.
- [goreleaser](https://github.com/goreleaser/goreleaser) - Deliver Go binaries as fast and easily as possible. Stars:`11.9K`.
- [wuzz](https://github.com/asciimoo/wuzz) - Interactive cli tool for HTTP inspection. Stars:`10.3K`.
- [usql](https://github.com/knq/usql) - usql is a universal command-line interface for SQL databases. Stars:`8.2K`.
- [peco](https://github.com/peco/peco) - Simplistic interactive filtering tool. Stars:`7.4K`.
- [go-funk](https://github.com/thoas/go-funk) - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). Stars:`4.3K`.
- [godropbox](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications from Dropbox. Stars:`4.1K`.
- [hystrix-go](https://github.com/afex/hystrix-go) - Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. Stars:`4.0K`.
- [panicparse](https://github.com/maruel/panicparse) - Groups similar goroutines and colorizes stack dump. Stars:`3.4K`.
- [minify](https://github.com/tdewolff/minify) - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. Stars:`3.4K`.
- [goreporter](https://github.com/wgliang/goreporter) - Golang tool that does static analysis, unit testing, code review and generate code quality report. Stars:`3.1K`.
- [lancet](https://github.com/duke-git/lancet) - A comprehensive, efficient, and reusable util function library of go. Stars:`2.7K`.
- [mc](https://github.com/minio/mc) - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. Stars:`2.5K`.
- [create-go-app](https://github.com/create-go-app/cli) - A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command. Stars:`2.1K`.
- [Storm](https://github.com/asdine/storm) - Simple and powerful toolkit for BoltDB. Stars:`2.0K`.
- [EaseProbe](https://github.com/megaease/easeprobe) - A simple, standalone, and lightWeight tool that can do health/status checking daemon, support HTTP/TCP/SSH/Shell/Client/... probes, and Slack/Discord/Telegram/SMS... notification. Stars:`1.9K`.
- [filetype](https://github.com/h2non/filetype) - Small package to infer the file type checking the magic numbers signature. Stars:`1.9K`.
- [retry-go](https://github.com/avast/retry-go) - Simple library for retry mechanism. Stars:`1.8K`.
- [mole](https://github.com/davrodpin/mole) - cli app to easily create ssh tunnels. Stars:`1.7K`.
- [jump](https://github.com/gsamokovarov/jump) - Jump helps you navigate faster by learning your habits. Stars:`1.7K`.
- [boilr](https://github.com/tmrts/boilr) - Blazingly fast CLI tool for creating projects from boilerplate templates. Stars:`1.6K`.
- [gitbatch](https://github.com/isacikgoz/gitbatch) - manage your git repositories in one place. Stars:`1.5K`.
- [mimetype](https://github.com/gabriel-vasile/mimetype) - Package for MIME type detection based on magic numbers. Stars:`1.2K`.
- [circuitbreaker](https://github.com/rubyist/circuitbreaker) - Circuit Breakers in Go. Stars:`1.1K`.
- [scany](https://github.com/georgysavva/scany) - Library for scanning data from a database into Go structs and more. Stars:`964`.
- [git-time-metric](https://github.com/git-time-metric/gtm) - Simple, seamless, lightweight time tracking for Git. Stars:`946`.
- [hostctl](https://github.com/guumaster/hostctl) - A CLI tool to manage /etc/hosts with easy commands. Stars:`918`.
- [immortal](https://github.com/immortal/immortal) - \*nix cross-platform (OS agnostic) supervisor. Stars:`766`.
- [circuit](https://github.com/cep21/circuit) - An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. Stars:`708`.
- [delve](https://github.com/derekparker/delve) - Go debugger. Stars:`597`.
- [ergo](https://github.com/cristianoliveira/ergo) - The management of multiple local services running over different ports made easy. Stars:`574`.
- [htcat](https://github.com/htcat/htcat) - Parallel and Pipelined HTTP GET Utility. Stars:`552`.
- [koazee](https://github.com/wesovilabs/koazee) - Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. Stars:`521`.
- [clockwork](https://github.com/jonboulle/clockwork) - A simple fake clock for golang. Stars:`507`.
- [godaemon](https://github.com/VividCortex/godaemon) - Utility to write daemons. Stars:`490`.
- [changie](https://github.com/miniscruff/changie) - Automated changelog tool for preparing releases with lots of customization options. Stars:`490`.
- [go-dry](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go. Stars:`487`.
- [gopencils](https://github.com/bndr/gopencils) - Small and simple package to easily consume REST APIs. Stars:`446`.
- [gubrak](https://github.com/novalagung/gubrak) - Golang utility library with syntactic sugar. It's like lodash, but for golang. Stars:`441`.
- [request](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™. Stars:`428`.
- [clipboard](https://github.com/golang-design/clipboard) - 📋 cross-platform clipboard package in Go. Stars:`428`.
- [scan](https://github.com/blockloop/scan) - Scan golang `sql.Rows` directly to structs, slices, or primitive types. Stars:`422`.
- [Deepcopier](https://github.com/ulule/deepcopier) - Simple struct copying for Go. Stars:`416`.
- [remote-touchpad](https://github.com/Unrud/remote-touchpad) - Control mouse and keyboard from a smartphone. Stars:`404`.
- [go-rate](https://github.com/beefsack/go-rate) - Timed rate limiter for Go. Stars:`389`.
- [mani](https://github.com/alajmo/mani) - CLI tool to help you manage multiple repositories. Stars:`377`.
- [retry](https://github.com/kamilsk/retry) - The most advanced functional mechanism to perform actions repetitively until successful. Stars:`334`.
- [serve](https://github.com/syntaqx/serve) - A static http server anywhere you need. Stars:`305`.
- [util](https://github.com/shomali11/util) - Collection of useful utility functions. (strings, concurrency, manipulations, ...). Stars:`278`.
- [grofer](https://github.com/pesos/grofer) - A system and resource monitoring tool written in Golang! Stars:`278`.
- [countries](https://github.com/biter777/countries) - Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standards. Stars:`269`.
- [gotenv](https://github.com/subosito/gotenv) - Load environment variables from `.env` or any `io.Reader` in Go. Stars:`266`.
- [gohper](https://github.com/cosiner/gohper) - Various tools/modules help for development. Stars:`255`.
- [wifiqr](https://github.com/reugn/wifiqr) - Wi-Fi QR Code Generator. Stars:`236`.
- [go-trigger](https://github.com/sadlil/go-trigger) - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. Stars:`234`.
- [rospo](https://github.com/ferama/rospo) - Simple and reliable ssh tunnels with embedded ssh server in Golang. Stars:`232`.
- [pattern-match](https://github.com/alexpantyukhin/go-pattern-match) - Pattern matching library. Stars:`228`.
- [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) - XML Sitemap generator written in Go. Stars:`204`.
- [toolbox](https://github.com/viant/toolbox) - Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. Stars:`192`.
- [Death](https://github.com/vrecan/death) - Managing go application shutdown with signals. Stars:`190`.
- [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) - go:generate tool for wrapping symbols exported by golang plugins (1.8 only). Stars:`183`.
- [limiters](https://github.com/mennanov/limiters) - Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. Stars:`179`.
- [moldova](https://github.com/StabbyCutyou/moldova) - Utility for generating random data based on an input template. Stars:`166`.
- [rerun](https://github.com/ivpusic/rerun) - Recompiling and rerunning go apps when source changes. Stars:`165`.
- [apm](https://github.com/topfreegames/apm) - Process manager for Golang applications with an HTTP API. Stars:`163`.
- [robustly](https://github.com/VividCortex/robustly) - Runs functions resiliently, catching and restarting panics. Stars:`156`.
- [chyle](https://github.com/antham/chyle) - Changelog generator using a git repository with multiple configuration possibilities. Stars:`150`.
- [go-bsdiff](https://github.com/gabstv/go-bsdiff) - Pure Go bsdiff and bspatch libraries and CLI tools. Stars:`145`.
- [cryptgo](https://github.com/Gituser143/cryptgo) - Crytpgo is a TUI based application written purely in Go to monitor and observe cryptocurrency prices in real time! Stars:`136`.
- [nostromo](https://github.com/pokanop/nostromo) - CLI for building powerful aliases. Stars:`135`.
- [cmd](https://github.com/SimonBaeumer/cmd) - Library for executing shell commands on osx, windows and linux. Stars:`134`.
- [filter](https://github.com/gookit/filter) - provide filtering, sanitizing, and conversion of Go data. Stars:`132`.
- [onecache](https://github.com/adelowo/onecache) - Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). Stars:`132`.
- [lrserver](https://github.com/jaschaephraim/lrserver) - LiveReload server for Go. Stars:`124`.
- [goval](https://github.com/maja42/goval) - Evaluate arbitrary expressions in Go. Stars:`123`.
- [mongo-go-pagination](https://github.com/gobeam/mongo-go-pagination) - Mongodb Pagination for official mongodb/mongo-go-driver package which supports both normal queries and Aggregation pipelines. Stars:`123`.
- [sorty](https://github.com/jfcg/sorty) - Fast Concurrent / Parallel Sorting. Stars:`122`.
- [mssqlx](https://github.com/linxGnu/mssqlx) - Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. Stars:`101`.
- [goseaweedfs](https://github.com/linxGnu/goseaweedfs) - SeaweedFS client library with almost full features. Stars:`101`.
- [xferspdy](https://github.com/monmohan/xferspdy) - Xferspdy provides binary diff and patch library in golang. Stars:`99`.
- [go-lock](https://github.com/viney-shih/go-lock) - go-lock is a lock library implementing read-write mutex and read-write trylock without starvation. Stars:`93`.
- [go-health](https://github.com/Talento90/go-health) - Health package simplifies the way you add health check to your services. Stars:`92`.
- [mimemagic](https://github.com/zRedShift/mimemagic) - Pure Go ultra performant MIME sniffing library/utility. Stars:`91`.
- [pgo](https://github.com/arthurkushman/pgo) - Convenient functions for PHP community. Stars:`82`.
- [repeat](https://github.com/ssgreg/repeat) - Go implementation of different backoff strategies useful for retrying operations and heartbeating. Stars:`81`.
- [pm](https://github.com/VividCortex/pm) - Process (i.e. goroutine) manager with an HTTP API. Stars:`78`.
- [countries](https://github.com/pioz/countries) - All you need when you are working with countries in Go. Stars:`76`.
- [handy](https://github.com/miguelpragier/handy) - Many utilities and helpers like string handlers/formatters and validators. Stars:`74`.
- [netbug](https://github.com/e-dard/netbug) - Easy remote profiling of your services. Stars:`72`.
- [UNIS](https://github.com/esemplastic/unis) - Common Architecture™ for String Utilities in Go. Stars:`70`.
- [multitick](https://github.com/VividCortex/multitick) - Multiplexor for aligned tickers. Stars:`67`.
- [goreadability](https://github.com/philipjkim/goreadability) - Webpage summary extractor using Facebook Open Graph and arc90's readability. Stars:`66`.
- [retry](https://github.com/thedevsaddam/retry) - Simple and easy retry mechanism package for Go. Stars:`63`.
- [go-astitodo](https://github.com/asticode/go-astitodo) - Parse TODOs in your GO code. Stars:`60`.
- [minquery](https://github.com/icza/minquery) - MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). Stars:`60`.
- [golog](https://github.com/mlimaloureiro/golog) - Easy and lightweight CLI tool to time track your tasks. Stars:`59`.
- [dbt](https://github.com/nikogura/dbt) - A framework for running self-updating signed binaries from a central, trusted repository. Stars:`57`.
- [beyond](https://github.com/wesovilabs/beyond) - The Go tool that will drive you to the AOP world! Stars:`53`.
- [scan](https://github.com/wroge/scan) - Scan sql rows into any type powered by generics. Stars:`52`.
- [shutdown](https://github.com/ztrue/shutdown) - App shutdown hooks for `os.Signal` handling. Stars:`52`.
- [slice](https://github.com/psampaz/slice) - Type-safe functions for common Go slice operations. Stars:`51`.
- [copy-pasta](https://github.com/jutkko/copy-pasta) - Universal multi-workstation clipboard that uses S3 like backend for the storage. Stars:`50`.
- [golarm](https://github.com/msempere/golarm) - Fire alarms with system events. Stars:`50`.
- [goback](https://github.com/carlescere/goback) - Go simple exponential backoff package. Stars:`48`.
- [retry-go](https://github.com/rafaeljesus/retry-go) - Retrying made simple and easy for golang. Stars:`47`.
- [intrinsic](https://github.com/mengzhuo/intrinsic) - Use x86 SIMD without writing any assembly code. Stars:`46`.
- [backscanner](https://github.com/icza/backscanner) - A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. Stars:`46`.
- [set](https://github.com/nofeaturesonlybugs/set) - Performant and flexible struct mapping and loose type conversion. Stars:`44`.
- [sshman](https://github.com/shoobyban/sshman) - SSH Manager for authorized_keys files on multiple remote servers. Stars:`44`.
- [equalizer](https://github.com/reugn/equalizer) - Quota manager and rate limiter collection for Go. Stars:`43`.
- [slicer](https://github.com/leaanthony/slicer) - Makes working with slices easier. Stars:`43`.
- [go-httpheader](https://github.com/mozillazg/go-httpheader) - Go library for encoding structs into Header fields. Stars:`43`.
- [gostrutils](https://github.com/ik5/gostrutils) - Collections of string manipulation and conversion functions. Stars:`41`.
- [gpath](https://github.com/tenntenn/gpath) - Library to simplify access struct fields with Go's expression in reflection. Stars:`40`.
- [cvt](https://github.com/shockerli/cvt) - Easy and safe convert any value to another type. Stars:`38`.
- [evaluator](https://github.com/nullne/evaluator) - Evaluate an expression dynamically based on s-expression. It's simple and easy to extend. Stars:`37`.
- [pointer](https://github.com/xorcare/pointer) - Package pointer contains helper routines for simplifying the creation of optional fields of basic type. Stars:`37`.
- [ghokin](https://github.com/antham/ghokin) - Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). Stars:`35`.
- [myhttp](https://github.com/inancgumus/myhttp) - Simple API to make HTTP GET requests with timeout support. Stars:`35`.
- [rclient](https://github.com/zpatrick/rclient) - Readable, flexible, simple-to-use client for REST APIs. Stars:`35`.
- [tome](https://github.com/cyruzin/tome) - Tome was designed to paginate simple RESTful APIs. Stars:`34`.
- [copy](https://github.com/gotidy/copy) - Package for fast copying structs of different types. Stars:`33`.
- [throttle](https://github.com/yudppp/throttle) - Throttle is an object that will perform exactly one action per duration. Stars:`33`.
- [generate](https://github.com/go-playground/generate) - runs go generate recursively on a specified path or environment variable and can filter by regex. Stars:`29`.
- [mimesniffer](https://github.com/aofei/mimesniffer) - A MIME type sniffer for Go. Stars:`28`.
- [ugo](https://github.com/alxrm/ugo) - ugo is slice toolbox with concise syntax for Go. Stars:`27`.
- [goplaceholder](https://github.com/michiwend/goplaceholder) - a small golang lib to generate placeholder images. Stars:`26`.
- [watchhttp](https://github.com/nikolaydubina/watchhttp) - Run command periodically and expose latest STDOUT or its rich delta as HTTP endpoint. Stars:`25`.
- [gofn](https://github.com/tiendc/gofn) - High performance utitlity functions written using Generics for Go 1.18+. Stars:`24`.
- [graterm](https://github.com/skovtunenko/graterm) - Provides primitives to perform ordered (sequential/concurrent) GRAceful TERMination (aka shutdown) in Go application. Stars:`24`.
- [ctxutil](https://github.com/posener/ctxutil) - A collection of utility functions for contexts. Stars:`23`.
- [structs](https://github.com/PumpkinSeed/structs) - Implement simple functions to manipulate structs. Stars:`22`.
- [ptr](https://github.com/gotidy/ptr) - Package that provide functions for simplified creation of pointers from constants of basic types. Stars:`21`.
- [go-convert](https://github.com/Eun/go-convert) - Package go-convert enables you to convert a value into another type. Stars:`20`.
- [just](https://github.com/kazhuravlev/just) - Just a collection of useful functions for working with generic data structures. Stars:`19`.
- [jsend](https://github.com/clevergo/jsend) - JSend's implementation written in Go. Stars:`19`.
- [filler](https://github.com/yaronsumel/filler) - small utility to fill structs using "fill" tag. Stars:`17`.
- [contextplus](https://github.com/contextplus/contextplus) - Package contextplus provide more easy to use functions for contexts. Stars:`17`.
- [dlog](https://github.com/kirillDanshin/dlog) - Compile-time controlled logger to make your release smaller without removing debug calls. Stars:`17`.
- [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) - Go package for working with Problem Details. Stars:`17`.
- [okrun](https://github.com/xta/okrun) - go run error steamroller. Stars:`16`.
- [rest-go](https://github.com/edermanoel94/rest-go) - A package that provide many helpful methods for working with rest api. Stars:`16`.
- [go-type](https://github.com/mikekonan/go-types) - Library providing Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types. Stars:`15`.
- [go-actuator](https://github.com/sinhashubham95/go-actuator) - Production ready features for Go based web frameworks. Stars:`14`.
- [command](https://github.com/txgruppi/command) - Command pattern for Go with thread safe serial and parallel dispatcher. Stars:`14`.
- [silk](https://github.com/chrispassas/silk) - Read silk netflow files. Stars:`12`.
- [go-clip](https://github.com/prashantgupta24/go-clip) - A minimalistic clipboard manager for Mac. Stars:`12`.
- [retry](https://github.com/shafreeck/retry) - A pretty simple library to ensure your work to be done. Stars:`11`.
- [go-countries](https://github.com/mikekonan/go-countries) - Lightweight lookup over ISO-3166 codes. Stars:`10`.
- [blank](https://github.com/Henry-Sarabia/blank) - Verify or remove blanks and whitespace from strings. Stars:`10`.
- [statiks](https://github.com/janiltonmaciel/statiks) - Fast, zero-configuration, static HTTP filer server. Stars:`10`.
- [bleep](https://github.com/sinhashubham95/bleep) - Perform any number of actions on any set of OS signals in Go. Stars:`9`.
- [retry](https://github.com/percolate/retry) - A simple but highly configurable retry package for Go. Stars:`9`.
- [sliceconv](https://github.com/Henry-Sarabia/sliceconv) - Slice conversion between primitive types. Stars:`8`.
- [nfdump](https://github.com/chrispassas/nfdump) - Read nfdump netflow files. Stars:`8`.
- [loncha](https://github.com/kazu/loncha) - A high-performance slice Utilities. Stars:`7`.
- [reflectutils](https://github.com/muir/reflectutils) - Helpers for working with reflection: struct tag parsing; recursive walking; fill value from string. Stars:`6`.
- [go-pkg](https://github.com/chenquan/go-pkg) - A go toolkit. Stars:`6`.
- [lets-go](https://github.com/aplescia-chwy/lets-go) - Go module that provides common utilities for Cloud Native REST API development. Also contains AWS Specific utilities. Stars:`6`.
- [tik](https://github.com/andy2046/tik) - Simple and easy timing wheel package for Go. Stars:`5`.
- [olaf](https://github.com/btnguyen2k/olaf) - Twitter Snowflake implemented in Go. Stars:`5`.
- [goctx](https://github.com/zerosnake0/goctx) - Get your context value with high performance. Stars:`4`.
- [objwalker](https://github.com/rekby/objwalker) - Walk by go objects with reflection. Stars:`2`.
- [mergo](https://github.com/imdario/mergo) - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.

**[⬆ back to top](#contents)**

## UUID

_Libraries for working with UUIDs._

- [uuid](https://github.com/google/uuid) - Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services. Stars:`4.5K`.
- [ulid](https://github.com/oklog/ulid) - Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). Stars:`3.6K`.
- [xid](https://github.com/rs/xid) - Xid is a globally unique id generator library, ready to be safely used directly in your server code. Stars:`3.5K`.
- [uuid](https://github.com/gofrs/uuid) - Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. Stars:`1.4K`.
- [wuid](https://github.com/edwingeng/wuid) - An extremely fast globally unique number generator. Stars:`500`.
- [sno](https://github.com/muyo/sno) - Compact, sortable and fast unique IDs with embedded metadata. Stars:`86`.
- [nanoid](https://github.com/aidarkhanov/nanoid) - A tiny and efficient Go unique string ID generator. Stars:`58`.
- [goid](https://github.com/jakehl/goid) - Generate and Parse RFC4122 compliant V4 UUIDs. Stars:`38`.
- [gouid](https://github.com/twharmon/gouid) - Generate cryptographically secure random string IDs with just one allocation. Stars:`21`.
- [uuid](https://github.com/agext/uuid) - Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. Stars:`15`.
- [uniq](https://gitlab.com/skilstak/code/go/uniq) - No hassle safe, fast unique identifiers with commands.

**[⬆ back to top](#contents)**

## Validation

_Libraries for validation._

- [validator](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. Stars:`13.6K`.
- [govalidator](https://github.com/asaskevich/govalidator) - Validators and sanitizers for strings, numerics, slices and structs. Stars:`5.7K`.
- [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. Stars:`3.3K`.
- [govalidator](https://github.com/thedevsaddam/govalidator) - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. Stars:`1.2K`.
- [validate](https://github.com/gookit/validate) - Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. Stars:`921`.
- [checkdigit](https://github.com/osamingo/checkdigit) - Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). Stars:`101`.
- [jio](https://github.com/faceair/jio) - jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). Stars:`86`.
- [validate](https://github.com/gobuffalo/validate) - This package provides a framework for writing validations for Go applications. Stars:`86`.
- [gody](https://github.com/guiferpa/gody) - :balloon: A lightweight struct validator for Go. Stars:`64`.
- [govalid](https://github.com/twharmon/govalid) - Fast, tag-based validation for structs. Stars:`34`.
- [valix](https://github.com/marrow16/valix) Go package for validating requests Stars:`14`.
- [Validator](https://github.com/go-the-way/validator) - A lightweight model validator written in Go.Contains VFs:Min, Max, MinLength, MaxLength, Length, Enum, Regex. Stars:`5`.

**[⬆ back to top](#contents)**

## Version Control

_Libraries for version control._

- [go-git](https://github.com/go-git/go-git) - highly extensible Git implementation in pure Go. Stars:`4.8K`.
- [hercules](https://github.com/src-d/hercules) - gaining advanced insights from Git repository history. Stars:`1.9K`.
- [git2go](https://github.com/libgit2/git2go) - Go bindings for libgit2. Stars:`1.9K`.
- [gh](https://github.com/rjeczalik/gh) - Scriptable server and net/http middleware for GitHub Webhooks. Stars:`79`.
- [githooks](https://github.com/gabyx/githooks) - Per-repo and shared Git hooks with version control and auto update. Stars:`76`.
- [go-vcs](https://github.com/sourcegraph/go-vcs) - manipulate and inspect VCS repositories in Go. Stars:`75`.
- [froggit-go](https://github.com/jfrog/froggit-go) - Froggit-Go is a Go library, allowing to perform actions on VCS providers. Stars:`39`.
- [hgo](https://github.com/beyang/hgo) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories. Stars:`15`.
- [cli](https://gitlab.com/gitlab-org/cli) - An open-source GitLab command line tool bringing GitLab's cool features to your command line.

**[⬆ back to top](#contents)**

## Video

_Libraries for manipulating video._

- [goav](https://github.com/giorgisio/goav) - Comprehensive Go bindings for FFmpeg. Stars:`2.0K`.
- [m3u8](https://github.com/grafov/m3u8) - Parser and generator library of M3U8 playlists for Apple HLS. Stars:`1.1K`.
- [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries. Stars:`855`.
- [go-astits](https://github.com/asticode/go-astits) - Parse and demux MPEG Transport Streams (.ts) natively in GO. Stars:`492`.
- [go-astisub](https://github.com/asticode/go-astisub) - Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). Stars:`482`.
- [gortsplib](https://github.com/aler9/gortsplib) - Pure Go RTSP server and client library. Stars:`458`.
- [libvlc-go](https://github.com/adrg/libvlc-go) - Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). Stars:`359`.
- [gst](https://github.com/ziutek/gst) - Go bindings for GStreamer. Stars:`166`.
- [v4l](https://github.com/korandiz/v4l) - Video capture library for Linux, written in Go. Stars:`75`.
- [libgosubs](https://github.com/wargarblgarbl/libgosubs) - Subtitle format support for go. Supports .srt, .ttml, and .ass. Stars:`23`.
- [go-m3u8](https://github.com/etherlabsio/go-m3u8) - Parser and generator library for Apple m3u8 playlists. Actively maintained version of quangngotan95/go-m3u8 with improvements and latest HLS playlist parsing compatibility. Stars:`18`.
- [go-mpd](https://github.com/unki2aut/go-mpd) - Parser and generator library for MPEG-DASH manifest files. Stars:`16`.

**[⬆ back to top](#contents)**

## Web Frameworks

_Full stack web frameworks._

- [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. Stars:`70.6K`.
- [Beego](https://github.com/beego/beego) - beego is an open-source, high-performance web framework for the Go programming language. Stars:`30.0K`.
- [Fiber](https://github.com/gofiber/fiber) - An Express.js inspired web framework build on Fasthttp. Stars:`27.6K`.
- [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework. Stars:`26.2K`.
- [Revel](https://github.com/revel/revel) - High-productivity web framework for the Go language. Stars:`13.0K`.
- [GoFrame](https://github.com/gogf/gf) - GoFrame is a modular, full-featured and production-ready application development framework of golang. Stars:`9.8K`.
- [Goa](https://github.com/goadesign/goa) - Goa provides a holistic approach for developing remote APIs and microservices in Go. Stars:`5.3K`.
- [Hertz](https://github.com/cloudwego/hertz) - A high-performance and strong-extensibility Go HTTP framework that helps developers build microservices. Stars:`4.0K`.
- [Gizmo](https://github.com/NYTimes/gizmo) - Microservice toolkit used by the New York Times. Stars:`3.7K`.
- [go-json-rest](https://github.com/ant0ine/go-json-rest) - Quick and easy way to setup a RESTful JSON API. Stars:`3.5K`.
- [Macaron](https://github.com/go-macaron/macaron) - Macaron is a high productive and modular design web framework in Go. Stars:`3.4K`.
- [utron](https://github.com/gernest/utron) - Lightweight MVC framework for Go(Golang). Stars:`2.2K`.
- [Goyave](https://github.com/go-goyave/goyave) - Feature-complete REST API framework aimed at clean code and fast development, with powerful built-in functionalities. Stars:`1.3K`.
- [REST Layer](https://github.com/rs/rest-layer) - Framework to build REST/GraphQL API on top of databases with mostly configuration over code. Stars:`1.2K`.
- [Atreugo](https://github.com/savsgio/atreugo) - High performance and extensible micro web framework with zero memory allocations in hot paths. Stars:`1.1K`.
- [tigertonic](https://github.com/rcrowley/go-tigertonic) - Go framework for building JSON web services inspired by Dropwizard. Stars:`997`.
- [tango](https://github.com/lunny/tango) - Micro & pluggable web framework for Go. Stars:`834`.
- [Gearbox](https://github.com/abahmed/gearbox) - A web framework written in Go with a focus on high performance and memory optimization. Stars:`717`.
- [Aero](https://github.com/aerogo/aero) - High-performance web framework for Go, reaches top scores in Lighthouse. Stars:`531`.
- [gongular](https://github.com/mustafaakin/gongular) - Fast Go web framework with input mapping/validation and (DI) Dependency Injection. Stars:`502`.
- [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) - Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications. Stars:`446`.
- [Air](https://github.com/aofei/air) - An ideally refined web framework for Go. Stars:`431`.
- [rk-boot](https://github.com/rookie-ninja/rk-boot) - A bootstrapper library for building enterprise go microservice with Gin and gRPC quickly and easily. Stars:`425`.
- [neo](https://github.com/ivpusic/neo) - Neo is minimal and fast Go Web Framework with extremely simple API. Stars:`418`.
- [Flamingo](https://github.com/i-love-flamingo/flamingo) - Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc. Stars:`393`.
- [mango](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. Stars:`372`.
- [Gondola](https://github.com/rainycape/gondola) - The web framework for writing faster sites, faster. Stars:`310`.
- [uAdmin](https://github.com/uadmin/uadmin) - Fully featured web framework for Golang, inspired by Django. Stars:`286`.
- [WebGo](https://github.com/bnkamalesh/webgo) - A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). Stars:`283`.
- [Ginrpc](https://github.com/xxjwxc/ginrpc) - Gin parameter automatic binding tool,gin rpc tools. Stars:`276`.
- [Golf](https://github.com/dinever/golf) - Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. Stars:`266`.
- [hiboot](https://github.com/hidevopsio/hiboot) - hiboot is a high performance web application framework with auto configuration and dependency injection support. Stars:`179`.
- [appy](https://github.com/appist/appy) - An opinionated productive web framework that helps scaling business easier. Stars:`129`.
- [go-rest](https://github.com/ungerik/go-rest) - Small and evil REST framework for Go. Stars:`127`.
- [patron](https://github.com/beatlabs/patron) - Patron is a microservice framework following best cloud practices with a focus on productivity. Stars:`124`.
- [Microservice](https://github.com/claygod/microservice) - The framework for the creation of microservices, written in Golang. Stars:`111`.
- [rux](https://github.com/gookit/rux) - Simple and fast web framework for build golang HTTP applications. Stars:`94`.
- [vox](https://github.com/aisk/vox) - A golang web framework for humans, inspired by Koa heavily. Stars:`80`.
- [Golax](https://github.com/fulldump/golax) - A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. Stars:`76`.
- [YARF](https://github.com/yarf-framework/yarf) - Fast micro-framework designed to build REST APIs and web services in a fast and simple way. Stars:`65`.
- [Fireball](https://github.com/zpatrick/fireball) - More "natural" feeling web framework. Stars:`59`.
- [goa](https://github.com/goa-go/goa) - goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware. Stars:`49`.
- [GoTuna](https://github.com/gotuna/gotuna) - Minimalistic web framework for Go with mux router, middlewares, sessions, templates, embedded views and static files. Stars:`43`.
- [Pulse](https://github.com/gopulse/pulse) - Pulse is an HTTP web framework written in Go (Golang) Stars:`36`.
- [goweb](https://github.com/twharmon/goweb) - Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS. Stars:`35`.
- [rex](https://github.com/goanywhere/rex) - Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. Stars:`33`.
- [Resoursea](https://github.com/resoursea/api) - REST framework for quickly writing resource based services. Stars:`33`.
- [Don](https://github.com/abemedia/go-don) - A highly performant and simple to use API framework. Stars:`27`.
- [Banjo](https://github.com/nsheremet/banjo) - Very simple and fast web framework for Go. Stars:`21`.
- [anoweb](https://github.com/go-the-way/anoweb) - The lightweight and powerful web framework using the new way for Go.Another go the way. Stars:`5`.
- [golamb](https://github.com/twharmon/golamb) - Golamb makes it easier to write API endpoints for use with AWS Lambda and API Gateway. Stars:`5`.
- [Pnutmux](https://gitlab.com/fruitygo/pnutmux) - Pnutmux is a powerful Go web framework that uses regex for matching and handling HTTP requests. It offers features such as CORS handling, structured logging, URL parameters extraction, middlewares, and concurrency limiting.
- [Buffalo](https://gobuffalo.io) - Bringing the productivity of Rails to Go!
- [Confetti Framework](https://confetti-framework.github.io/docs/) - Confetti is a Go web application framework with an expressive, elegant syntax. Confetti combines the elegance of Laravel and the simplicity of Go.
- [aah](https://aahframework.org) - Scalable, performant, rapid development Web framework for Go.
- [Huma](https://github.com/danielgtaylor/huma/) - Framework for modern REST/GraphQL APIs with built-in OpenAPI 3, generated documentation, and a CLI.

**[⬆ back to top](#contents)**

### Middlewares

#### Actual middlewares

- [Tollbooth](https://github.com/didip/tollbooth) - Rate limit HTTP request handler. Stars:`2.5K`.
- [CORS](https://github.com/rs/cors) - Easily add CORS capabilities to your API. Stars:`2.4K`.
- [Limiter](https://github.com/ulule/limiter) - Dead simple rate limit middleware for Go. Stars:`1.8K`.
- [go-server-timing](https://github.com/mitchellh/go-server-timing) - Add/parse Server-Timing header. Stars:`851`.
- [go-fault](https://github.com/github/go-fault) - Fault injection middleware for Go. Stars:`479`.
- [ln-paywall](https://github.com/philippgille/ln-paywall) - Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). Stars:`136`.
- [XFF](https://github.com/sebest/xff) - Handle `X-Forwarded-For` header and friends. Stars:`97`.
- [rk-grpc](https://github.com/rookie-ninja/rk-grpc) - Middleware for gRPC with logging, metrics, auth, tracing etc. Stars:`64`.
- [rk-gin](https://github.com/rookie-ninja/rk-gin) - Middleware for Gin framework with logging, metrics, auth, tracing etc. Stars:`39`.
- [formjson](https://github.com/rs/formjson) - Transparently handle JSON input as a standard form POST. Stars:`38`.
- [client-timing](https://github.com/posener/client-timing) - An HTTP client for Server-Timing header. Stars:`24`.
- [echo-middleware](https://github.com/faabiosr/echo-middleware) - Middleware for Echo framework with logging and metrics. Stars:`14`.
- [mid](https://github.com/bobg/mid) - Miscellaneous HTTP middleware features: idiomatic error return from handlers; receive/respond with JSON data; request tracing; and more. Stars:`7`.

#### Libraries for creating HTTP middlewares

- [negroni](https://github.com/urfave/negroni) - Idiomatic HTTP middleware for Golang. Stars:`7.3K`.
- [alice](https://github.com/justinas/alice) - Painless middleware chaining for Go. Stars:`2.8K`.
- [render](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, and HTML template responses. Stars:`1.8K`.
- [stats](https://github.com/thoas/stats) - Go middleware that stores various information about your web application. Stars:`593`.
- [interpose](https://github.com/carbocation/interpose) - Minimalist net/http middleware for golang. Stars:`294`.
- [renderer](https://github.com/thedevsaddam/renderer) - Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. Stars:`255`.
- [muxchain](https://github.com/stephens2424/muxchain) - Lightweight middleware for net/http. Stars:`208`.
- [rye](https://github.com/InVisionApp/rye) - Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. Stars:`102`.
- [gores](https://github.com/alioygur/gores) - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. Stars:`100`.
- [mediary](https://github.com/HereMobilityDevelopers/mediary) - add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses. Stars:`86`.
- [chain](https://github.com/codemodus/chain) - Handler wrapper chaining with scoped data (net/context-based "middleware"). Stars:`64`.
- [catena](https://github.com/codemodus/catena) - http.Handler wrapper catenation (same API as "chain"). Stars:`9`.

**[⬆ back to top](#contents)**

### Routers

- [mux](https://github.com/gorilla/mux) - Powerful URL router and dispatcher for golang. Stars:`18.9K`.
- [httprouter](https://github.com/julienschmidt/httprouter) - High performance router. Use this and the standard http handlers to form a very high performance web framework. Stars:`15.6K`.
- [chi](https://github.com/go-chi/chi) - Small, fast and expressive HTTP router built on net/context. Stars:`14.9K`.
- [gocraft/web](https://github.com/gocraft/web) - Mux and middleware package in Go. Stars:`1.5K`.
- [Bone](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer. Stars:`1.3K`.
- [Goji](https://github.com/goji/goji) - Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. Stars:`944`.
- [fasthttprouter](https://github.com/buaazp/fasthttprouter) - High performance router forked from `httprouter`. The first router fit for `fasthttp`. Stars:`875`.
- [httptreemux](https://github.com/dimfeld/httptreemux) - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. Stars:`602`.
- [xujiajun/gorouter](https://github.com/xujiajun/gorouter) - A simple and fast HTTP router for Go. Stars:`532`.
- [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. Stars:`447`.
- [lars](https://github.com/go-playground/lars) - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. Stars:`390`.
- [Siesta](https://github.com/VividCortex/siesta) - Composable framework to write middleware and handlers. Stars:`351`.
- [vestigo](https://github.com/husobee/vestigo) - Performant, stand-alone, HTTP compliant URL Router for go web applications. Stars:`269`.
- [gowww/router](https://github.com/gowww/router) - Lightning fast HTTP router fully compatible with the net/http.Handler interface. Stars:`166`.
- [GoRouter](https://github.com/vardius/gorouter) - GoRouter is a Server/API micro framework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. Stars:`149`.
- [pure](https://github.com/go-playground/pure) - Is a lightweight HTTP router that sticks to the std "net/http" implementation. Stars:`140`.
- [alien](https://github.com/gernest/alien) - Lightweight and fast http router from outer space. Stars:`127`.
- [violetear](https://github.com/nbari/violetear) - Go HTTP router. Stars:`106`.
- [Bxog](https://github.com/claygod/Bxog) - Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. Stars:`104`.
- [xmux](https://github.com/rs/xmux) - High performance muxer based on `httprouter` with `net/context` support. Stars:`96`.
- [goblin](https://github.com/bmf-san/goblin) - A golang http router based on trie tree. Stars:`63`.
- [ngamux](https://github.com/ngamux/ngamux) - Simple HTTP router for Go. Stars:`59`.
- [bellt](https://github.com/GuilhermeCaruso/bellt) - A simple Go HTTP router. Stars:`54`.
- [FastRouter](https://github.com/razonyang/fastrouter) - a fast, flexible HTTP router written in Go. Stars:`22`.
- [GoLobby/Router](https://github.com/golobby/router) - GoLobby Router is a lightweight yet powerful HTTP router for the Go programming language. Stars:`21`.
- [nchi](https://github.com/muir/nchi) - chi-like router built on httprouter with dependency injection based middleware wrappers Stars:`14`.
- [goroute](https://github.com/goroute/route) - Simple yet powerful HTTP request multiplexer. Stars:`8`.

**[⬆ back to top](#contents)**

## WebAssembly

- [tinygo](https://github.com/tinygo-org/tinygo) - Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM. Stars:`13.3K`.
- [dom](https://github.com/dennwc/dom) - DOM library. Stars:`479`.
- [go-canvas](https://github.com/markfarnan/go-canvas) - Library to use HTML5 Canvas, with all drawing within go code. Stars:`212`.
- [wasmbrowsertest](https://github.com/agnivade/wasmbrowsertest) - Run Go WASM tests in your browser. Stars:`149`.
- [webapi](https://github.com/gowebapi/webapi) - Bindings for DOM and HTML generated from WebIDL. Stars:`142`.
- [vert](https://github.com/norunners/vert) - Interop between Go and JS values. Stars:`89`.

**[⬆ back to top](#contents)**

## Windows

- [go-ole](https://github.com/go-ole/go-ole) - Win32 OLE implementation for golang. Stars:`1.0K`.
- [d3d9](https://github.com/gonutz/d3d9) - Go bindings for Direct3D9. Stars:`146`.
- [gosddl](https://github.com/MonaxGT/gosddl) - Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. Stars:`10`.

**[⬆ back to top](#contents)**

## XML

_Libraries and tools for manipulating XML._

- [zek](https://github.com/miku/zek) - Generate a Go struct from XML. Stars:`678`.
- [xpath](https://github.com/antchfx/xpath) - XPath package for Go. Stars:`593`.
- [xquery](https://github.com/antchfx/xquery) - XQuery lets you extract data from HTML/XML documents using XPath expression. Stars:`156`.
- [xml2map](https://github.com/sbabiv/xml2map) - XML to MAP converter written Golang. Stars:`52`.
- [xmlwriter](https://github.com/shabbyrobe/xmlwriter) - Procedural XML generation API based on libxml2's xmlwriter module. Stars:`24`.
- [XML-Comp](https://github.com/xml-comp/xml-comp) - Simple command line XML comparer that generates diffs of folders, files and tags. Stars:`20`.

## Zero Trust

_Libraries and tools to implement Zero Trust architectures._

- [Cosign](https://github.com/sigstore/cosign) - Container Signing, Verification and Storage in an OCI registry. Stars:`3.5K`.
- [Spire](https://github.com/spiffe/spire) - SPIRE (the SPIFFE Runtime Environment) is a toolchain of APIs for establishing trust between software systems across a wide variety of hosting platforms. Stars:`1.5K`.
- [in-toto](https://github.com/in-toto/in-toto-golang) - Go implementation of the in-toto (provides a framework to protect the integrity of the software supply chain) python reference implementation. Stars:`98`.
- [Spiffe-Vault](https://github.com/philips-labs/spiffe-vault) - Utilizes Spiffe JWT authentication with Hashicorp Vault for secretless authentication. Stars:`61`.

## Code Analysis

_Source code analysis tools, also known as Static Application Security Testing (SAST) Tools._

- [GoLint](https://github.com/golang/lint) - Golint is a linter for Go source code. Stars:`4.0K`.
- [errcheck](https://github.com/kisielk/errcheck) - Errcheck is a program for checking for unchecked errors in Go programs. Stars:`2.1K`.
- [go-critic](https://github.com/go-critic/go-critic) - source code linter that brings checks that are currently not implemented in other linters. Stars:`1.7K`.
- [GoPlantUML](https://github.com/jfeliu007/goplantuml) - Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. Stars:`1.5K`.
- [gcvis](https://github.com/davecheney/gcvis) - Visualise Go program GC trace data in real time. Stars:`1.1K`.
- [php-parser](https://github.com/z7zmey/php-parser) - A Parser for PHP written in Go. Stars:`929`.
- [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) - go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. Stars:`752`.
- [goast-viewer](https://github.com/yuroyoro/goast-viewer) - Web based Golang AST visualizer. Stars:`719`.
- [golines](https://github.com/segmentio/golines) - Formatter that automatically shortens long lines in Go code. Stars:`658`.
- [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) - An easy way to find outdated dependencies of your Go projects. Stars:`640`.
- [Chronos](https://github.com/amit-davidson/Chronos) - Detects race conditions statically Stars:`405`.
- [todocheck](https://github.com/preslavmihaylov/todocheck) - Static code analyser which links TODO comments in code with issues in your issue tracker. Stars:`398`.
- [unconvert](https://github.com/mdempsky/unconvert) - Remove unnecessary type conversions from Go source. Stars:`359`.
- [dupl](https://github.com/mibk/dupl) - Tool for code clone detection. Stars:`316`.
- [tickgit](https://github.com/augmentable-dev/tickgit) - CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author. Stars:`307`.
- [gostatus](https://github.com/shurcooL/gostatus) - Command line tool, shows the status of repositories that contain Go packages. Stars:`243`.
- [vaccum](https://github.com/daveshanley/vacuum) - An ultra-super-fast, lightweight OpenAPI linter and quality checking tool. Stars:`240`.
- [apicompat](https://github.com/bradleyfalzon/apicompat) - Checks recent changes to a Go project for backwards incompatible changes. Stars:`178`.
- [go-checkstyle](https://github.com/qiniu/checkstyle) - checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. Stars:`128`.
- [lint](https://github.com/surullabs/lint) - Run linters as part of go test. Stars:`67`.
- [validate](https://github.com/mccoyst/validate) - Automatically validates struct fields with tags. Stars:`61`.
- [asty](https://github.com/asty-org/asty) - Converts golang AST to JSON and JSON to AST. Stars:`59`.
- [go-outdated](https://github.com/firstrow/go-outdated) - Console application that displays outdated packages. Stars:`45`.
- [ChainJacking](https://github.com/Checkmarx/chainjacking) - Find which of your Go lang direct GitHub dependencies is susceptible to ChainJacking attack. Stars:`43`.
- [usestdlibvars](https://github.com/sashamelentyev/usestdlibvars) - A linter that detect the possibility to use variables/constants from the Go standard library. Stars:`40`.
- [golang-ifood-sdk](https://github.com/arxdsilva/golang-ifood-sdk) - iFood API SDK. Stars:`10`.
- [goreturns](https://sourcegraph.com/github.com/sqs/goreturns) - Adds zero-value return statements to match the func return types.
- [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
- [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
- [blanket](https://gitlab.com/verygoodsoftwarenotvirus/blanket) - blanket is a tool that helps you catch functions which don't have direct unit tests in your Go packages.

**[⬆ back to top](#contents)**

## Editor Plugins

_Plugin for text editors and IDEs._

- [vim-go](https://github.com/fatih/vim-go) - Go development plugin for Vim. Stars:`15.5K`.
- [gocode](https://github.com/nsf/gocode) - Autocompletion daemon for the Go programming language. Stars:`5.0K`.
- [vscode-go](https://github.com/golang/vscode-go) - Extension for Visual Studio Code (VS Code) which provides support for the Go language. Stars:`3.5K`.
- [GoSublime](https://github.com/DisposaBoy/GoSublime) - Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. Stars:`3.4K`.
- [go-plus](https://github.com/joefitzgerald/go-plus) - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. Stars:`1.5K`.
- [go-mode](https://github.com/dominikh/go-mode.el) - Go mode for GNU/Emacs. Stars:`1.3K`.
- [coc-go language server extension for Vim/Neovim](https://github.com/josa42/coc-go) - This plugin adds [gopls](https://github.com/golang/tools/blob/master/gopls/README.md) features to Vim/Neovim. Stars:`543`.
- [goimports-reviser](https://github.com/incu6us/goimports-reviser) - Formatting tool for imports. Stars:`387`.
- [Watch](https://github.com/eaburns/Watch) - Runs a command in an acme win on file changes. Stars:`200`.
- [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) - Vim plugin to highlight syntax errors on save. Stars:`87`.
- [go-language-server](https://github.com/theia-ide/go-language-server) - A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. Stars:`31`.
- [gounit-vim](https://github.com/hexdigest/gounit-vim) - Vim plugin for generating Go tests based on the function's or method's signature. Stars:`23`.
- [theia-go-extension](https://github.com/theia-ide/theia-go-extension) - Go language support for the Theia IDE. Stars:`15`.
- [Go Doc](https://github.com/msyrus/vscode-go-doc) - A Visual Studio Code extension for showing definition in output and generating go doc. Stars:`6`.
- [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof) - This extension adds benchmark profiling support for the Go language to VS Code.
- [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) - Go plugin for JetBrains IDEs.

**[⬆ back to top](#contents)**

## Go Generate Tools

- [gotests](https://github.com/cweill/gotests) - Generate Go tests from your source code. Stars:`4.6K`.
- [genny](https://github.com/cheekybits/genny) - Elegant generics for Go. Stars:`1.7K`.
- [xgen](https://github.com/xuri/xgen) - XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator. Stars:`260`.
- [re2dfa](https://github.com/opennota/re2dfa) - Transform regular expressions into finite state machines and output Go source code. Stars:`193`.
- [hasgo](https://github.com/DylanMeeus/hasgo) - Generate Haskell inspired functions for your slices. Stars:`129`.
- [gonerics](https://github.com/bouk/gonerics) - Idiomatic Generics in Go. Stars:`114`.
- [gocontracts](https://github.com/Parquery/gocontracts) - brings design-by-contract to Go by synchronizing the code with the documentation. Stars:`98`.
- [gounit](https://github.com/hexdigest/gounit) - Generate Go tests using your own templates. Stars:`72`.
- [sqlgen](https://github.com/anqiansong/sqlgen) - Generate gorm, xorm, sqlx, bun, sql code from SQL file or DSN. Stars:`66`.
- [generic](https://github.com/usk81/generic) - flexible data type for Go. Stars:`48`.
- [options-gen](https://github.com/kazhuravlev/options-gen) - Functional options described by Dave Cheney's post "Functional options for friendly APIs". Stars:`43`.
- [godal](https://github.com/mafulong/godal) - Generate orm models corresponding to golang by specifying sql ddl file, which can be used by gorm. Stars:`17`.
- [TOML-to-Go](https://xuri.me/toml-to-go) - Translates TOML into a Go type in the browser instantly.

**[⬆ back to top](#contents)**

## Go Tools

- [go-swagger](https://github.com/go-swagger/go-swagger) - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. Stars:`8.8K`.
- [go-callvis](https://github.com/TrueFurby/go-callvis) - Visualize call graph of your Go program using dot format. Stars:`5.3K`.
- [OctoLinker](https://github.com/OctoLinker/browser-extension) - Navigate through go files efficiently with the OctoLinker browser extension for GitHub. Stars:`5.2K`.
- [depth](https://github.com/KyleBanks/depth) - Visualize dependency trees of any package by analyzing imports. Stars:`843`.
- [richgo](https://github.com/kyoh86/richgo) - Enrich `go test` outputs with text decorations. Stars:`802`.
- [rts](https://github.com/galeone/rts) - RTS: response to struct. Generates Go structs from server responses. Stars:`245`.
- [godbg](https://github.com/tylerwince/godbg) - Implementation of Rusts `dbg!` macro for quick and easy debugging during development. Stars:`196`.
- [typex](https://github.com/dtgorski/typex) - Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration. Stars:`180`.
- [roumon](https://github.com/becheran/roumon) - Monitor current state of all active goroutines via a command line interface. Stars:`146`.
- [gothanks](https://github.com/psampaz/gothanks) - GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers. Stars:`118`.
- [colorgo](https://github.com/songgao/colorgo) - Wrapper around `go` command for colorized `go build` output. Stars:`112`.
- [go-james](https://github.com/pieterclaerhout/go-james) - Go project skeleton creator, builds and tests your projects without the manual setup. Stars:`59`.
- [igo](https://github.com/rocketlaunchr/igo) - An igo to go transpiler (new language features for Go language!) Stars:`59`.
- [gotestdox](https://github.com/bitfield/gotestdox) - Show Go test results as readable sentences. Stars:`52`.
- [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) - Bash completion for go and wgo. Stars:`42`.
- [docs](https://github.com/go-oas/docs) - Automatically generate RESTful API documentation for GO projects - aligned with Open API Specification standard. Stars:`23`.
- [modver](https://github.com/bobg/modver) - Compare two versions of a Go module to check the version-number change required (major, minor, or patchlevel), according to [semver](https://semver.org/) rules. Stars:`8`.
- [textra](https://github.com/ravsii/textra) - Extract Go struct field names, types and tags for filtering and exporting. Stars:`3`.
- [gotemplate.io](https://gotemplate.io/) - Online tool to preview `text/template` templates live.
- [gomodrun](https://github.com/dustinblackman/gomodrun/) - Go tool that executes and caches binaries included in go.mod files.

**[⬆ back to top](#contents)**

## Software Packages

_Software written in Go._

**[⬆ back to top](#contents)**

### DevOps Tools

- [kubernetes](https://github.com/kubernetes/kubernetes) - Container Cluster Manager from Google. Stars:`100.6K`.
- [Moby](https://github.com/moby/moby) - Collaborative project for the container ecosystem to assemble container-based systems. Stars:`66.4K`.
- [traefik](https://github.com/containous/traefik) - Reverse proxy and load balancer with support for multiple backends. Stars:`44.1K`.
- [Gitea](https://github.com/go-gitea/gitea) - Fork of Gogs, entirely community driven. Stars:`37.8K`.
- [Vegeta](https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000! Stars:`21.6K`.
- [Hey](https://github.com/rakyll/hey) - Hey is a tiny program that sends some load to a web application. Stars:`16.0K`.
- [Packer](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. Stars:`14.6K`.
- [Mizu](https://github.com/up9inc/mizu) - API traffic viewer for Kubernetes enabling you to view all API communication between microservices, multiprotocol support: HTTP1.1, HTTP/2, AMQP, Kafka, Redis. Stars:`9.4K`.
- [webhook](https://github.com/adnanh/webhook) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. Stars:`9.2K`.
- [GVM](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions. Stars:`8.7K`.
- [Ddosify](https://github.com/ddosify/ddosify) - High-performance load testing tool, written in Golang. Stars:`7.7K`.
- [ko](https://github.com/google/ko) - Command line tool for building and deploying Go applications on Kubernetes Stars:`6.2K`.
- [KubeVela](https://github.com/kubevela/kubevela) - Cloud native application delivery. Stars:`5.5K`.
- [gaia](https://github.com/gaia-pipeline/gaia) - Build powerful pipelines in any programming language. Stars:`5.1K`.
- [gox](https://github.com/mitchellh/gox) - Dead simple, no frills Go cross compile tool. Stars:`4.6K`.
- [bombardier](https://github.com/codesenberg/bombardier) - Fast cross-platform HTTP benchmarking tool. Stars:`4.5K`.
- [script](https://github.com/bitfield/script) - Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. Stars:`4.2K`.
- [Pomerium](https://github.com/pomerium/pomerium) - Pomerium is an identity-aware access proxy. Stars:`3.6K`.
- [bosun](https://github.com/bosun-monitor/bosun) - Time Series Alerting Framework. Stars:`3.4K`.
- [kala](https://github.com/ajvb/kala) - Simplistic, modern, and performant job scheduler. Stars:`2.0K`.
- [fac](https://github.com/mkchoi212/fac) - Command-line user interface to fix git merge conflicts. Stars:`1.8K`.
- [s5cmd](https://github.com/peak/s5cmd) - Blazing fast S3 and local filesystem execution tool. Stars:`1.8K`.
- [goxc](https://github.com/laher/goxc) - build tool for Go, with a focus on cross-compiling and packaging. Stars:`1.7K`.
- [StatusOK](https://github.com/sanathp/statusok) - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. Stars:`1.6K`.
- [Fleet device management](https://github.com/fleetdm/fleet) - Lightweight, programmable telemetry for servers and workstations. Stars:`1.5K`.
- [ghorg](https://github.com/gabrie30/ghorg) - Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Gitea, and Bitbucket. Stars:`1.3K`.
- [go-selfupdate](https://github.com/sanbornm/go-selfupdate) - Enable your Go applications to self update. Stars:`1.2K`.
- [s3gof3r](https://github.com/rlmcpherson/s3gof3r) - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. Stars:`1.1K`.
- [uTask](https://github.com/ovh/utask) - Automation engine that models and executes business processes declared in yaml. Stars:`943`.
- [skm](https://github.com/TimothyYe/skm) - SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! Stars:`883`.
- [Scaleway-cli](https://github.com/scaleway/scaleway-cli) - Manage BareMetal Servers from Command Line (as easily as with Docker). Stars:`836`.
- [kwatch](https://github.com/abahmed/kwatch) - Monitor & detect crashes in your Kubernetes(K8s) cluster instantly. Stars:`776`.
- [cassowary](https://github.com/rogerwelin/cassowary) - Modern cross-platform HTTP load-testing tool written in Go. Stars:`676`.
- [kool](https://github.com/kool-dev/kool) - Command line tool for managing Docker environments as an easy way. Stars:`629`.
- [aurora](https://github.com/xuri/aurora) - Cross-platform web-based Beanstalkd queue server console. Stars:`580`.
- [govvv](https://github.com/ahmetalpbalkan/govvv) - “go build” wrapper to easily add version information into Go binaries. Stars:`535`.
- [Pewpew](https://github.com/bengadbois/pewpew) - Flexible HTTP command line stress tester. Stars:`384`.
- [jcli](https://github.com/jenkins-zh/jenkins-cli) - Jenkins CLI allows you manage your Jenkins as an easy way. Stars:`365`.
- [gonative](https://github.com/inconshreveable/gonative) - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. Stars:`334`.
- [trubka](https://github.com/xitonix/trubka) - A CLI tool to manage and troubleshoot Apache Kafka clusters with the ability of generically publishing/consuming protocol buffer and plain text events to/from Kafka. Stars:`330`.
- [Mora](https://github.com/emicklei/mora) - REST server for accessing MongoDB documents and meta data. Stars:`312`.
- [lstags](https://github.com/ivanilves/lstags) - Tool and API to sync Docker images across different registries. Stars:`307`.
- [Balerter](https://github.com/balerter/balerter) - A self-hosted script-based alerting manager. Stars:`289`.
- [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) - Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. Stars:`287`.
- [manssh](https://github.com/xwjdsh/manssh) - manssh is a command line tool for managing your ssh alias config easily. Stars:`285`.
- [gobrew](https://github.com/kevincobain2000/gobrew) - Go version manager. Super simple tool to install and manage Go versions. Install go without root. Gobrew doesn't require shell rehash. Stars:`264`.
- [dogo](https://github.com/liudng/dogo) - Monitoring changes in the source file and automatically compile and run (restart). Stars:`257`.
- [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) - Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. Stars:`252`.
- [godbg](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application. Stars:`227`.
- [Blast](https://github.com/dave/blast) - A simple tool for API load testing and batch jobs. Stars:`211`.
- [s3-proxy](https://github.com/oxyno-zeta/s3-proxy) - S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth). Stars:`204`.
- [abbreviate](https://github.com/dnnrly/abbreviate) - abbreviate is a tool turning long strings in to shorter ones with configurable separators, for example to embed branch names in to deployment stack IDs. Stars:`203`.
- [kcli](https://github.com/cswank/kcli) - Command line tool for inspecting kafka topics/partitions/messages. Stars:`196`.
- [gobrew](https://github.com/cryptojuice/gobrew) - gobrew lets you easily switch between multiple versions of go. Stars:`194`.
- [ostent](https://github.com/ostrost/ostent) - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. Stars:`177`.
- [grapes](https://github.com/yaronsumel/grapes) - Lightweight tool designed to distribute commands over ssh with ease. Stars:`163`.
- [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) - A go library and an executable that produces valid Dockerfiles using various input channels. Stars:`154`.
- [winrm-cli](https://github.com/masterzen/winrm-cli) - Cli tool to remotely execute commands on Windows machines. Stars:`154`.
- [drone-scp](https://github.com/appleboy/drone-scp) - Copy files and artifacts via SSH using a binary, docker or Drone CI. Stars:`123`.
- [tf-profile](https://github.com/datarootsio/tf-profile) - Profiler for Terraform runs. Generate global stats, resource-level stats or visualizations. Stars:`118`.
- [Mantil](https://github.com/mantil-io/mantil) - Go specific framework for building serverless applications on AWS that enables you to focus on pure Go code while Mantil takes care of the infrastructure. Stars:`104`.
- [go-furnace](https://github.com/go-furnace/go-furnace) - Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. Stars:`97`.
- [go-rocket-update](https://github.com/mouuff/go-rocket-update) - A simple way to make self updating Go applications - Supports Github and Gitlab. Stars:`83`.
- [Dropship](https://github.com/chrismckenzie/dropship) - Tool for deploying code via cdn. Stars:`64`.
- [drone-jenkins](https://github.com/appleboy/drone-jenkins) - Trigger downstream Jenkins jobs using a binary, docker or Drone CI. Stars:`37`.
- [docker-go-mingw](https://github.com/x1unix/docker-go-mingw) - Docker image for building Go binaries for Windows with MinGW toolchain. Stars:`37`.
- [Rodent](https://github.com/alouche/rodent) - Rodent helps you manage Go versions, projects and track dependencies. Stars:`34`.
- [awsenv](https://github.com/soniah/awsenv) - Small binary that loads Amazon (AWS) environment variables for a profile. Stars:`33`.
- [httpref](https://github.com/dnnrly/httpref) - httpref is a handy CLI reference for HTTP methods, status codes, headers, and TCP and UDP ports. Stars:`31`.
- [lwc](https://github.com/timdp/lwc) - A live-updating version of the UNIX wc command. Stars:`29`.
- [DepCharge](https://github.com/centerorbit/depcharge) - Helps orchestrating the execution of commands across the many dependencies in larger projects. Stars:`23`.
- [wait-for](https://github.com/dnnrly/wait-for) - Wait for something to happen (from the command line) before continuing. Easy orchestration of Docker services and other things. Stars:`14`.
- [aptly](https://github.com/smira/aptly) - aptly is a Debian repository management tool. Stars:`9`.
- [sg](https://github.com/ChristopherRabotin/sg) - Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. Stars:`8`.
- [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator) - Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
- [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
- [Wide](https://wide.b3log.org/login) - Web-based IDE for Teams using Golang.

**[⬆ back to top](#contents)**

### Other Software

- [croc](https://github.com/schollz/croc) - Easily and securely send files or folders from one computer to another. Stars:`23.2K`.
- [restic](https://github.com/restic/restic) - De-duplicating backup program. Stars:`20.8K`.
- [Seaweed File System](https://github.com/chrislusf/seaweedfs) - Fast, Simple and Scalable Distributed File System with O(1) disk seek. Stars:`18.0K`.
- [Gor](https://github.com/buger/gor) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. Stars:`17.7K`.
- [Comcast](https://github.com/tylertreat/Comcast) - Simulate bad network connections. Stars:`10.0K`.
- [toxiproxy](https://github.com/shopify/toxiproxy) - Proxy to simulate network and system conditions for automated tests. Stars:`9.3K`.
- [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul. Stars:`8.1K`.
- [LiteIDE](https://github.com/visualfc/liteide) - LiteIDE is a simple, open source, cross-platform Go IDE. Stars:`7.3K`.
- [drive](https://github.com/odeke-em/drive) - Google Drive client for the commandline. Stars:`6.6K`.
- [nes](https://github.com/fogleman/nes) - Nintendo Entertainment System (NES) emulator written in Go. Stars:`5.3K`.
- [scc](https://github.com/boyter/scc) - Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. Stars:`4.9K`.
- [Duplicacy](https://github.com/gilbertchen/duplicacy) - A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. Stars:`4.7K`.
- [blocky](https://github.com/0xERR0R/blocky) - Fast and lightweight DNS proxy as ad-blocker for local network with many features. Stars:`2.8K`.
- [myLG](https://github.com/mehrdadrad/mylg) - Command Line Network Diagnostic tool written in Go. Stars:`2.7K`.
- [GoBoy](https://github.com/Humpheh/goboy) - Nintendo Game Boy Color emulator written in Go. Stars:`2.5K`.
- [Stack Up](https://github.com/pressly/sup) - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. Stars:`2.4K`.
- [lgo](https://github.com/yunabe/lgo) - Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. Stars:`2.3K`.
- [Circuit](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. Stars:`2.0K`.
- [Documize](https://github.com/documize/community) - Modern wiki software that integrates data from SaaS tools. Stars:`1.9K`.
- [snap](https://github.com/intelsdi-x/snap) - Powerful telemetry framework. Stars:`1.8K`.
- [borg](https://github.com/crufter/borg) - Terminal based search engine for bash snippets. Stars:`1.6K`.
- [Plik](https://github.com/root-gg/plik) - Plik is a temporary file upload system (Wetransfer like) in Go. Stars:`1.3K`.
- [shell2http](https://github.com/msoap/shell2http) - Executing shell commands via http server (for prototyping or remote control). Stars:`1.1K`.
- [vFlow](https://github.com/VerizonDigital/vflow) - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. Stars:`1.0K`.
- [portal](https://github.com/SpatiumPortae/portal) - Portal is a quick and easy command-line file transfer utility from any computer to another. Stars:`922`.
- [peg](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. Stars:`920`.
- [Go Package Store](https://github.com/shurcooL/Go-Package-Store) - App that displays updates for the Go packages in your GOPATH. Stars:`895`.
- [Gokapi](https://github.com/Forceu/gokapi) - Lightweight server to share files, which expire after a set amount of downloads or days. Similar to Firefox Send, but without public upload. Stars:`755`.
- [Leaps](https://github.com/jeffail/leaps) - Pair programming service using Operational Transforms. Stars:`737`.
- [gfile](https://github.com/Antonito/gfile) - Securely transfer files between two computers, without any third party, over WebRTC. Stars:`705`.
- [Guora](https://github.com/meloalright/guora) - A self-hosted Quora like web application written in Go. Stars:`652`.
- [Gebug](https://github.com/moshebe/gebug) - A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly. Stars:`621`.
- [sake](https://github.com/alajmo/sake) - sake is a command runner for local and remote hosts. Stars:`608`.
- [gocc](https://github.com/goccmack/gocc) - Gocc is a compiler kit for Go written in Go. Stars:`579`.
- [mockingjay](https://github.com/quii/mockingjay-server) - Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. Stars:`543`.
- [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) - Video streaming torrent client. Stars:`455`.
- [woke](https://github.com/get-woke/woke) - Detect non-inclusive language in your source code. Stars:`387`.
- [ipe](https://github.com/dimiro1/ipe) - Open source Pusher server implementation compatible with Pusher client libraries written in GO. Stars:`367`.
- [ide](https://github.com/thestrukture/ide) - Browser accessible IDE. Designed for Go with Go. Stars:`354`.
- [tcpprobe](https://github.com/mehrdadrad/tcpprobe) - TCP tool for network performance and path monitoring, including socket statistics. Stars:`338`.
- [wellington](https://github.com/wellington/wellington) - Sass project management tool, extends the language with sprite functions (like Compass). Stars:`303`.
- [Cherry](https://github.com/rafael-santiago/cherry) - Tiny webchat server in Go. Stars:`289`.
- [GoNB](https://github.com/janpfeifer/gonb) - Interactive Go programming with Jupyter Notebooks (also works in VSCode, Binder and Google's Colab). Stars:`267`.
- [Neo-cowsay](https://github.com/Code-Hex/Neo-cowsay) - 🐮 cowsay is reborn. for a New Era. Stars:`258`.
- [hotswap](https://github.com/edwingeng/hotswap) - A complete solution to reload your go code without restarting your server, interrupting or blocking any ongoing procedure. Stars:`241`.
- [tcpdog](https://github.com/mehrdadrad/tcpdog) - eBPF based TCP observability. Stars:`235`.
- [joincap](https://github.com/assafmo/joincap) - Command-line utility for merging multiple pcap files together. Stars:`192`.
- [Orbit](https://github.com/gulien/orbit) - A simple tool for running commands and generating files from templates. Stars:`177`.
- [crawley](https://github.com/s0rg/crawley) - Web scraper/crawler for cli. Stars:`153`.
- [vaku](https://github.com/lingrino/vaku) - CLI & API for folder-based functions in Vault like copy, move, and search. Stars:`150`.
- [stew](https://github.com/marwanhawari/stew) - An independent package manager for compiled binaries. Stars:`115`.
- [dp](https://github.com/scryinfo/dp) - Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. Stars:`84`.
- [boxed](https://github.com/tejo/boxed) - Dropbox based blog engine. Stars:`78`.
- [term-quiz](https://github.com/crazcalm/term-quiz) - Quizzes for your terminal. Stars:`24`.
- [naclpipe](https://github.com/unix4fun/naclpipe) - Simple NaCL EC25519 based crypto pipe tool written in Go. Stars:`23`.
- [Snitch](https://github.com/lucasgomide/snitch) - Simple way to notify your team and many tools when someone has deployed any application via Tsuru. Stars:`16`.
- [GoDocTooltip](https://github.com/diankong/GoDocTooltip) - Chrome extension for Go Doc sites, which shows function description as tooltip at function list. Stars:`13`.
- [hoofli](https://github.com/dnnrly/hoofli) - Generate PlantUML diagrams from Chrome or Firefox network inspections. Stars:`6`.
- [protoncheck](https://github.com/servusdei2018/protoncheck) - ProtonMail module for waybar/polybar/yabar/i3blocks. Stars:`4`.
- [hugo](https://gohugo.io/) - Fast and Modern Static Website Engine.
- [syncthing](https://syncthing.net/) - Open, decentralized file synchronization tool and protocol.
- [Better Go Playground](https://goplay.tools) - Go playground with syntax highlight, code completion and other features.
- [GoLand](https://jetbrains.com/go) - Full featured cross-platform Go IDE.
- [goblin](https://goblin.reaper.im) - Golang binaries in a curl, built by goblins.
- [Docker](https://www.docker.com/) - Open platform for distributed applications for developers and sysadmins.
- [tsuru](https://tsuru.io/) - Extensible and open source Platform as a Service software.
- [limetext](https://limetext.github.io) - Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
- [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
- [zs](https://git.mills.io/prologic/zs) - an extremely minimal static site generator.

**[⬆ back to top](#contents)**

# Resources

_Where to discover new Go libraries._

**[⬆ back to top](#contents)**

## Benchmarks

- [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) - Go web framework benchmark. Stars:`1.9K`.
- [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router benchmark and comparison. Stars:`1.6K`.
- [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods. Stars:`1.5K`.
- [skynet](https://github.com/atemerev/skynet) - Skynet 1M threads microbenchmark. Stars:`1.0K`.
- [speedtest-resize](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language. Stars:`236`.
- [go-benchmarks](https://github.com/tylertreat/go-benchmarks) - Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. Stars:`146`.
- [gospeed](https://github.com/feyeleanor/GoSpeed) - Go micro-benchmarks for calculating the speed of language constructs. Stars:`111`.
- [autobench](https://github.com/davecheney/autobench) - Framework to compare the performance between different Go versions. Stars:`96`.
- [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) - Collection of benchmarks for popular Go database/SQL utilities. Stars:`65`.
- [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) - Benchmarks of common basic operations for the Go language. Stars:`60`.
- [go-ml-benchmarks](https://github.com/nikolaydubina/go-ml-benchmarks) - benchmarks for machine learning inference in Go. Stars:`28`.
- [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) - Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. Stars:`26`.
- [kvbench](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark. Stars:`25`.
- [go-json-benchmark](https://github.com/zerosnake0/go-json-benchmark) - Go JSON benchmark. Stars:`8`.

**[⬆ back to top](#contents)**

## Conferences

- [GoCon](https://gocon.connpass.com/) - Tokyo, Japan.
- [GoDays](https://www.godays.io/) - Berlin, Germany.
- [GoLab](https://golab.io/) - Florence, Italy.
- [GopherChina](https://gopherchina.org) - Shanghai, China.
- [GopherCon](https://www.gophercon.com/) - Denver, USA.
- [GopherCon Australia](https://gophercon.com.au/) - Sydney, Australia.
- [GopherCon Brazil](https://gopherconbr.org) - Florianópolis, Brazil.
- [GopherCon Europe](https://gophercon.eu/) - Berlin, Germany.
- [GopherCon India](https://www.gophercon.in/) - Pune, India.
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
- [Go Faster](https://leanpub.com/gofaster) - This book seeks to shorten your learning curve and help you become a proficient Go programmer, faster.
- [Know Go: Generics](https://bitfieldconsulting.com/books/generics) - A guide to understanding and using generics in Go.
- [Lets-Go](https://lets-go.alexedwards.net) - A step-by-step guide to creating fast, secure and maintanable web applications with Go.
- [Lets-Go-Further](https://lets-go-further.alexedwards.net) - Advanced patterns for building APIs and web applications in Go.
- [The Power of Go: Tests](https://bitfieldconsulting.com/books/tests) - A guide to testing in Go.
- [The Power of Go: Tools](https://bitfieldconsulting.com/books/tools) - A guide to writing command-line tools in Go.
- [Writing A Compiler In Go](https://compilerbook.com)
- [Writing An Interpreter In Go](https://interpreterbook.com) - Book that introduces dozens of techniques for writing idiomatic, expressive, and efficient Go code that avoids common pitfalls.

### Free e-books

- [GoBooks](https://github.com/dariubs/GoBooks) - A curated list of Go books. Stars:`14.5K`.
- [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) Stars:`9.3K`.
- [Go AST Book (Chinese)](https://github.com/chai2010/go-ast-book) - A book focusing on Go `go/*` packages. Stars:`5.1K`.
- [Go Succinctly](https://github.com/thedevsir/gosuccinctly) - in Persian. Stars:`23`.
- [How To Code in Go eBook](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook) - A 600 page introduction to Go aimed at first time developers.
- [Go 101](https://go101.org) - A book focusing on Go syntax/semantics and all kinds of details.
- [Building Web Apps With Go](https://codegangsta.gitbooks.io/building-web-apps-with-go/content/)
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/)
- [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
- [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
- [Network Programming With Go](https://jan.newmarch.name/golang/)
- [Practical Go Lessons](https://www.practical-go-lessons.com/)
- [Spaceship Go A Journey to the Standard Library](https://blasrodri.github.io/spaceship-go-gh-pages/)
- [The Go Programming Language](https://www.gopl.io/)
- [An Introduction to Programming in Go](http://www.golang-book.com/)
- [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)

**[⬆ back to top](#contents)**

## Gophers

- [gophers](https://github.com/egonelbre/gophers) - Free gophers. Stars:`3.2K`.
- [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) - Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. Stars:`3.0K`.
- [gophers](https://github.com/ashleymcnamara/gophers) - Gopher artworks by Ashley McNamara. Stars:`2.8K`.
- [gopherize.me](https://github.com/matryer/gopherize.me) - Gopherize yourself. Stars:`651`.
- [gophericons](https://github.com/shalakhin/gophericons) Stars:`610`.
- [gopher-stickers](https://github.com/tenntenn/gopher-stickers) Stars:`562`.
- [gophers](https://github.com/sillecelik/go-gopher) - Gopher amigurumi toy pattern. Stars:`120`.
- [gopher-logos](https://github.com/GolangUA/gopher-logos) - adorable gopher logos. Stars:`114`.
- [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) - Go gopher Vector Data [.ai, .svg]. Stars:`65`.
- [gophers](https://github.com/rogeralsing/gophers) - random gopher graphics. Stars:`58`.
- [gophers](https://github.com/scraly/gophers) - Gophers by Aurélie Vache. Stars:`24`.

**[⬆ back to top](#contents)**

## Meetups

- [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
- [Belfast Gophers](https://www.meetup.com/Belfast-Gophers/)
- [Belgrade Golang Meetup](https://www.meetup.com/golang-serbia/)
- [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
- [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
- [Canberra Gophers](https://www.meetup.com/Canberra-Gophers/)
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
- [Golang Korea](https://www.meetup.com/GDG-Golang-Korea/)
- [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
- [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
- [Golang Marseille](https://www.meetup.com/fr-FR/Golang-Marseille/)
- [Golang Melbourne](https://www.meetup.com/golang-mel/)
- [Golang Mountain View](https://www.meetup.com/Golang-Mountain-View/)
- [Golang North East](https://www.meetup.com/en-AU/Golang-North-East/)
- [Golang Paris](https://www.meetup.com/Golang-Paris/)
- [Golang Poland](https://www.meetup.com/Golang-Poland/)
- [Golang Pune](https://www.meetup.com/Golang-Pune/)
- [Golang Rotterdam](https://www.meetup.com/golang-rotterdam/)
- [Golang Singapore](https://www.meetup.com/golangsg/)
- [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
- [Golang Sydney, AU](https://www.meetup.com/golang-syd/)
- [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
- [Golang Taipei](https://www.meetup.com/golang-taipei-meetup/)
- [Golang Thessaloniki](https://www.meetup.com/thessaloniki-golang-meetup/)
- [Golang Turkey](https://kommunity.com/goturkiye)
- [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
- [Golang Vienna, Austria](https://www.meetup.com/viennago/)
- [Golang Москва](https://www.meetup.com/Golang-Moscow/)
- [GoSF - San Francisco, CA](https://www.meetup.com/golangsf)
- [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
- [Seattle Go Programmers](https://www.meetup.com/golang/)
- [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
- [Utah Go User Group](https://www.meetup.com/utahgophers/)
- [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)

_Add the group of your city/country here (send **PR**)_

**[⬆ back to top](#contents)**

## Style Guides

- [bahlo/go-styleguide](https://github.com/bahlo/go-styleguide) Stars:`1.4K`.
- [CockroachDB](https://github.com/cockroachdb/cockroach/blob/master/docs/style.md)
- [GitLab](https://docs.gitlab.com/ee/development/go_guide/)
- [Google](https://google.github.io/styleguide/go/)
- [Hyperledger](https://github.com/hyperledger/fabric/blob/release-1.4/docs/source/style-guides/go-style.rst)
- [Magnetico](https://github.com/boramalper/magnetico/wiki/magnetico-Design-Specification)
- [Sourcegraph](https://docs.sourcegraph.com/dev/background-information/languages/go)
- [Thanos](https://thanos.io/tip/contributing/coding-style-guide.md/)
- [Trybe](https://github.com/betrybe/playbook-go/blob/main/README_EN.md)
- [Uber](https://github.com/uber-go/guide/blob/master/style.md)

**[⬆ back to top](#contents)**

## Social Media

### Twitter

- [@GoDiscussions](https://twitter.com/GoDiscussions)
- [@golang](https://twitter.com/golang)
- [@golang_news](https://twitter.com/golang_news)
- [@golangch](https://twitter.com/golangch)
- [@golangflow](https://twitter.com/golangflow)
- [@golangweekly](https://twitter.com/golangweekly)

**[⬆ back to top](#contents)**

### Reddit

- [r/golang](https://www.reddit.com/r/golang/)

**[⬆ back to top](#contents)**

## Websites

- [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) - List of other amazingly awesome lists. Stars:`30.4K`.
- [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) - Curated list of awesome remote jobs. A lot of them are looking for Go hackers. Stars:`26.1K`.
- [Awesome Golang Workshops](https://github.com/amit-davidson/awesome-golang-workshops) - A curated list of awesome golang workshops. Stars:`486`.
- [golang-graphics](https://github.com/mholt/golang-graphics) - Collection of Go images, graphics, and art. Stars:`138`.
- [gocryforhelp](https://github.com/ninedraft/gocryforhelp) - Collection of Go projects that needs help. Good place to start your open-source way in Go. Stars:`41`.
- [awesome-go-extra](https://github.com/xwjdsh/awesome-go-extra) - Parse awesome-go README file and generate a new README file with repo info. Stars:`21`.
- [godoc.org](https://godoc.org/) - Documentation for open source Go packages.
- [Golang News](https://golangnews.com) - Links and news about Go programming.
- [Explore Go Libraries & Projects](https://kandi.openweaver.com/explore/go) - Discover & find a curated list of popular & new Go libraries, top authors, trending project kits, discussions, tutorials & learning resources on kandi.
- [Go Blog](https://blog.golang.org) - The official Go blog.
- [Go Code Club](https://www.youtube.com/watch?v=nvoIPQYdx9g&list=PLEcwzBXTPUE_YQR7R0BRtHBYJ0LN3Y0i3) - A group of Gophers read and discuss a different Go project every week.
- [Go Community on Hashnode](https://hashnode.com/n/go) - Community of Gophers on Hashnode.
- [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
- [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki.
- [Go Proverbs](https://go-proverbs.github.io/) - Go Proverbs by Rob Pike.
- [Go Report Card](https://goreportcard.com) - A report card for your Go package.
- [go.dev](https://go.dev/) - A hub for Go developers.
- [Coding Mystery](https://codingmystery.com) - Solve exciting escape-room-inspired programming challenges using Go.
- [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
- [Golang Developer Jobs](https://golangjob.xyz) - Developer Jobs exclusively for Golang related Roles.
- [Golang Flow](https://golangflow.io) - Post Updates, News, Packages and more.
- [CodinGame](https://www.codingame.com/) - Learn Go by solving interactive tasks using small games as practical examples.
- [Golang Resources](https://golangresources.com) - A curation of the best articles, exercises, talks and videos to learn Go.
- [Golang Weekly](https://discu.eu/weekly/golang/) - Each monday projects, tutorials and articles about Go.
- [Code with Mukesh](https://codewithmukesh.com/blog/category/golang) - Software Engineer and Blogs @ codewithmukesh.com.
- [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
- [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for #golang enthusiasts.
- [Gopher Community Chat](https://invite.slack.golangbridge.org) - Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
- [Gophercises](https://gophercises.com/) - Free coding exercises for budding gophers.
- [gowalker.org](https://gowalker.org) - Go Project API documentation.
- [json2go](https://m-zajac.github.io/json2go) - Advanced JSON to Go struct conversion - online tool.
- [justforfunc](https://www.youtube.com/c/justforfunc) - Youtube channel dedicated to Go programming language tips and tricks, hosted by Francesc Campoy [@francesc](https://twitter.com/francesc).
- [Learn Go Programming](https://blog.learngoprogramming.com) - Learn Go concepts with illustrations.
- [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
- [r/Golang](https://www.reddit.com/r/golang) - News about Go.
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
- [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) - Golang ebook intro how to build a web app with golang. Stars:`42.2K`.
- [go-patterns](https://github.com/tmrts/go-patterns) - Curated list of Go design patterns, recipes and idioms. Stars:`22.6K`.
- [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) - Learn Go with test-driven development. Stars:`19.9K`.
- [Learn Go with 1000+ Exercises](https://github.com/inancgumus/learngo) - Learn Go with thousands of examples, exercises, and quizzes. Stars:`17.6K`.
- [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) - Go's reference card. Stars:`7.7K`.
- [go-clean-template](https://github.com/evrone/go-clean-template) - Clean Architecture template for Golang services. Stars:`5.3K`.
- [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) - Examples of Golang compared to Node.js for learning. Stars:`4.0K`.
- [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) - A little e-book on Ethereum Development with Go. Stars:`1.6K`.
- [Working with Go](https://github.com/mkaz/working-with-go) - Intro to go for experienced programmers. Stars:`1.2K`.
- [goapp](https://github.com/bnkamalesh/goapp) - An opinionated guideline to structure & develop a Go web application/service. Stars:`639`.
- [Go in 7 days](https://github.com/harrytran103/7_days_of_go) - Learn everything about Go in 7 days (from a Nodejs developer). Stars:`124`.
- [Design Patterns in Go](https://github.com/shubhamzanwar/design-patterns) - Collection of programming design patterns implemented in Go. Stars:`104`.
- [Hex Monscape](https://github.com/Haraj-backend/hex-monscape) - Getting started guidelines in writing maintainable code using Hexagonal Architecture. Stars:`42`.
- [Debugged.it Go patterns](https://github.com/haveyoudebuggedit/go-patterns) - Advanced Go patterns with ready-to-run examples. Stars:`11`.
- [Go By Example](https://gobyexample.com/) - Hands-on introduction to Go using annotated example programs.
- [How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker) - Learn how to use Docker for Go development and how to build production Docker images.
- [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
- [Go Language Tutorial](https://www.javatpoint.com/go-tutorial) - Learn Go language Tutorial.
- [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
- [Games With Go](https://www.youtube.com/watch?v=9D4yH7e_ea8&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x) - A video series teaching programming and game development.
- [CodeCrafters Golang Track](https://app.codecrafters.io/tracks/go) - Achieve mastery in advanced Go by building your own Redis, Docker, Git, and SQLite. Featuring goroutines, systems programming, file I/O, and more.
- [Golang Tutorial Guide](https://www.freecodecamp.org/news/golang-tutorial-list-free-courses-learn-go-programming-language/) - A List of Free Courses to Learn the Go Programming Language.
- [Golangbot](https://golangbot.com/learn-golang-series/) - Tutorials to get started with programming in Go.
- [GopherCoding](https://gophercoding.com/) - Collection of code snippets and tutorials to help tackle every day issues.
- [GopherSnippets](https://gophersnippets.com/) - Code snippets with tests and testable examples for the Go programming language.
- [Gosamples](https://gosamples.dev/) - Collection of code snippets that let you solve everyday code problems.
- [Hackr.io](https://hackr.io/tutorials/learn-golang) - Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
- [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30) - How to cancel MySQL queries.
- [How to Benchmark: dbq vs sqlx vs GORM](https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5) - Learn how to benchmark in Go. As a case-study, we will benchmark dbq, sqlx and GORM.
- [Go Tutorial](https://www.tutorialspoint.com/go/index.htm) - Learn Go programming.
- [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
- [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9) - How to cache slow database queries.
- [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
- [Learning Go by examples](https://dev.to/aurelievache/learning-go-by-examples-introduction-448n) - Series of articles in order to learn Golang language by concrete applications as example.
- [Microservices with Go](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_) - Dive deep into building microservices using Go, including gRPC.
- [package main](https://www.youtube.com/packagemain) - YouTube channel about Programming in Go.
- [Programming with Google Go](https://www.coursera.org/specializations/google-golang) - Coursera Specialization to learn about Go from scratch.
- [Saving a Third of Our Memory by Re-ordering Go Struct Fields](https://qvault.io/golang/struct-field-ordering-memory/) - How inefficient field ordering in Go structs.
- [Scaling Go Applications](https://betterstack.com/community/guides/scaling-go/) - Everything about building, deploying and scaling Go applications in production.
- [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
- [W3basic Go Tutorials](https://www.w3basic.com/golang/) - W3Basic provides an in-depth tutorial and well-organized content to learn Golang programming.
- [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql) - We’ll write an API with the help of the powerful Gorilla Mux.
- [Your basic Go](https://yourbasic.org/golang) - Huge collection of tutorials and how to's.

**[⬆ back to top](#contents)**

### Guided Learning

- [The Go Developer Roadmap](https://roadmap.sh/golang) - A visual roadmap that new Go developers can follow through to help them learn Go.
- [The Go Learning Path](https://tutorialedge.net/paths/golang/) - A guided learning path containing a mix of free and premium resources.

**[⬆ back to top](#contents)**
