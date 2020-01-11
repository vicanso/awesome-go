# Awesome Go

[![Build Status](https://travis-ci.org/avelino/awesome-go.svg?branch=master)](https://travis-ci.org/avelino/awesome-go) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](http://gophers.slack.com/messages/awesome) [![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)

[![patreon avelino](https://c5.patreon.com/external/logo/become_a_patron_button@2x.png)](https://www.patreon.com/avelinosource) financial support to Awesome Go

A curated list of awesome Go frameworks, libraries and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

### Contributing

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/master/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*

### Contents

- [Awesome Go](#awesome-go)
    - [Audio and Music](#audio-and-music)
    - [Authentication and OAuth](#authentication-and-oauth)
    - [Bot Building](#bot-building)
    - [Command Line](#command-line)
    - [Configuration](#configuration)
    - [Continuous Integration](#continuous-integration)
    - [CSS Preprocessors](#css-preprocessors)
    - [Data Structures](#data-structures)
    - [Database](#database)
    - [Database Drivers](#database-drivers)
    - [Date and Time](#date-and-time)
    - [Distributed Systems](#distributed-systems)
    - [Dynamic DNS](#dynamic-dns)
    - [Email](#email)
    - [Embeddable Scripting Languages](#embeddable-scripting-languages)
    - [Error Handling](#error-handling)
    - [Files](#files)
    - [Financial](#financial)
    - [Forms](#forms)
    - [Functional](#functional)
    - [Game Development](#game-development)
    - [Generation and Generics](#generation-and-generics)
    - [Geographic](#geographic)
    - [Go Compilers](#go-compilers)
    - [Goroutines](#goroutines)
    - [GUI](#gui)
    - [Hardware](#hardware)
    - [Images](#images)
    - [IoT](#iot-internet-of-things)
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
    - [Natural Language Processing](#natural-language-processing)
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
    - [Windows](#windows)
    - [XML](#xml)

- [Tools](#tools)
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
    - [Gophers](#gophers)
    - [Meetups](#meetups)
    - [Twitter](#twitter)
    - [Websites](#websites)
        - [Tutorials](#tutorials)

## Audio and Music

*Libraries for manipulating audio.*

* [Oto](https://github.com/hajimehoshi/oto) - A low-level library to play sound on multiple platforms. Stars:`505`.
* [PortAudio](https://github.com/gordonklaus/portaudio) - Go bindings for the PortAudio audio I/O library. Stars:`330`.
* [music-theory](https://github.com/go-music-theory/music-theory) - Music theory models in Go. Stars:`274`.
* [waveform](https://github.com/mdlayher/waveform) - Go package capable of generating waveform images from audio streams. Stars:`269`.
* [portmidi](https://github.com/rakyll/portmidi) - Go bindings for PortMidi. Stars:`221`.
* [id3v2](https://github.com/bogem/id3v2) - Fast and stable ID3 parsing and writing library for Go. Stars:`129`.
* [flac](https://github.com/mewkiz/flac) - Native Go FLAC encoder/decoder with support for FLAC streams. Stars:`112`.
* [mix](https://github.com/go-mix/mix) - Sequence-based Go-native audio mixer for music apps. Stars:`106`.
* [mp3](https://github.com/tcolgate/mp3) - Native Go MP3 decoder. Stars:`104`.
* [go-sox](https://github.com/krig/go-sox) - libsox bindings for go. Stars:`103`.
* [malgo](https://github.com/gen2brain/malgo) - Mini audio library. Stars:`90`.
* [taglib](https://github.com/wtolson/go-taglib) - Go bindings for taglib. Stars:`72`.
* [gaad](https://github.com/Comcast/gaad) - Native Go AAC bitstream parser. Stars:`64`.
* [minimp3](https://github.com/tosone/minimp3) - Lightweight MP3 decoder library. Stars:`34`.
* [go_mediainfo](https://github.com/zhulik/go_mediainfo) - libmediainfo bindings for go. Stars:`28`.
* [EasyMIDI](https://github.com/algoGuy/EasyMIDI) - EasyMidi is a simple and reliable library for working with standard midi file (SMF). Stars:`26`.
* [vorbis](https://github.com/mccoyst/vorbis) - "Native" Go Vorbis decoder (uses CGO, but has no dependencies). Stars:`24`.
* [gosamplerate](https://github.com/dh1tw/gosamplerate) - libsamplerate bindings for go. Stars:`9`.

## Authentication and OAuth

*Libraries for implementing authentications schemes.*

* [jwt-go](https://github.com/dgrijalva/jwt-go) - Golang implementation of JSON Web Tokens (JWT). Stars:`6.9K`.
* [casbin](https://github.com/hsluoyz/casbin) - Authorization library that supports access control models like ACL, RBAC, ABAC. Stars:`5.9K`.
* [oauth2](https://github.com/golang/oauth2) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine and App Engine support. Stars:`2.7K`.
* [goth](https://github.com/markbates/goth) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box. Stars:`2.5K`.
* [authboss](https://github.com/volatiletech/authboss) - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure, and start building your app without having to build an authentication system each time. Stars:`2.1K`.
* [osin](https://github.com/openshift/osin) - Golang OAuth2 server library. Stars:`1.6K`.
* [go-jose](https://github.com/square/go-jose) - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs. Stars:`1.4K`.
* [go-oauth2-server](https://github.com/RichardKnop/go-oauth2-server) - Standalone, specification-compliant,  OAuth2 server written in Golang. Stars:`1.4K`.
* [gologin](https://github.com/dghubble/gologin) - chainable handlers for login with OAuth1 and OAuth2 authentication providers. Stars:`1.1K`.
* [gorbac](https://github.com/mikespook/gorbac) - provides a lightweight role-based access control (RBAC) implementation in Golang. Stars:`970`.
* [loginsrv](https://github.com/tarent/loginsrv) - JWT login microservice with plugable backends such as OAuth2 (Github), htpasswd, osiam. Stars:`884`.
* [scs](https://github.com/alexedwards/scs) - Session Manager for HTTP servers. Stars:`655`.
* [permissions2](https://github.com/xyproto/permissions2) - Library for keeping track of users, login states and permissions. Uses secure cookies and bcrypt. Stars:`382`.
* [paseto](https://github.com/o1egl/paseto) - Golang implementation of Platform-Agnostic Security Tokens (PASETO). Stars:`286`.
* [httpauth](https://github.com/goji/httpauth) - HTTP Authentication middleware. Stars:`190`.
* [jeff](https://github.com/abraithwaite/jeff) - Simple, flexible, secure and idiomatic web session management with pluggable backends. Stars:`184`.
* [jwt-auth](https://github.com/adam-hanna/jwt-auth) - JWT middleware for Golang http servers with many configuration options. Stars:`164`.
* [jwt](https://github.com/pascaldekloe/jwt) - Lightweight JSON Web Token (JWT) library. Stars:`138`.
* [session](https://github.com/icza/session) - Go session management for web servers (including support for Google App Engine - GAE). Stars:`95`.
* [branca](https://github.com/hako/branca) - Golang implementation of Branca Tokens. Stars:`91`.
* [jwt](https://github.com/robbert229/jwt) - Clean and easy to use implementation of JSON Web Tokens (JWT). Stars:`78`.
* [sessionup](https://github.com/swithek/sessionup) - Simple, yet effective HTTP session management and identification package. Stars:`76`.
* [sessions](https://github.com/adam-hanna/sessions) - Dead simple, highly performant, highly customizable sessions service for go http servers. Stars:`49`.
* [sjwt](https://github.com/brianvoe/sjwt) - Simple jwt generator and parser. Stars:`49`.
* [rbac](https://github.com/zpatrick/rbac) - Minimalistic RBAC package for Go applications. Stars:`43`.
* [securecookie](https://github.com/chmike/securecookie) - Efficient secure cookie encoding/decoding. Stars:`34`.
* [sessiongate-go](https://github.com/f0rmiga/sessiongate-go) - Go session management using the SessionGate Redis module. Stars:`8`.
* [signedvalue](https://github.com/sashka/signedvalue) - Signed and timestamped strings compatible with [Tornado's](https://github.com/tornadoweb/tornado) `create_signed_value`, `decode_signed_value`, and therefore `set_secure_cookie` and `get_secure_cookie`. Stars:`8`.
* [scope](https://github.com/SonicRoshan/scope) - Easily Manage OAuth2 Scopes In Go. Stars:`3`.
* [cookiestxt](https://github.com/mengzhuo/cookiestxt) - provides parser of cookies.txt file format. Stars:`2`.

## Bot Building

*Libraries for building and working with bots.*

* [telegram-bot-api](https://github.com/Syfaro/telegram-bot-api) - Simple and clean Telegram bot client. Stars:`1.9K`.
* [telebot](https://github.com/tucnak/telebot) - Telegram bot framework written in Go. Stars:`1.1K`.
* [go-chat-bot](https://github.com/go-chat-bot/bot) - IRC, Slack & Telegram bot written in Go. Stars:`535`.
* [slacker](https://github.com/shomali11/slacker) - Easy to use framework to create Slack bots. Stars:`345`.
* [Golang CryptoTrading Bot](https://github.com/saniales/golang-crypto-trading-bot) - A golang implementation of a console-based trading bot for cryptocurrency exchanges. Stars:`272`.
* [Kelp](https://github.com/stellar/kelp) - official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies. Stars:`249`.
* [tbot](https://github.com/yanzay/tbot) - Telegram bot server with API similar to net/http. Stars:`243`.
* [Tenyks](https://github.com/kyleterry/tenyks) - Service oriented IRC bot using Redis and JSON for messaging. Stars:`168`.
* [go-sarah](https://github.com/oklahomer/go-sarah) - Framework to build bot for desired chat services including LINE, Slack, Gitter and more. Stars:`151`.
* [hanu](https://github.com/sbstjn/hanu) - Framework for writing Slack bots. Stars:`116`.
* [go-tgbot](https://github.com/olebedev/go-tgbot) - Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router and middleware. Stars:`91`.
* [margelet](https://github.com/zhulik/margelet) - Framework for building Telegram bots. Stars:`60`.
* [govkbot](https://github.com/nikepan/govkbot) - Simple Go [VK](https://vk.com) bot library. Stars:`29`.
* [slackscot](https://github.com/alexandre-normand/slackscot) - Another framework for building Slack bots. Stars:`20`.
* [micha](https://github.com/onrik/micha) - Go Library for Telegram bot api. Stars:`12`.
* [go-joe](https://joe-bot.net) - A general-purpose bot library inspired by Hubot but written in Go.

## Command Line

### Standard CLI

*Libraries for building standard or basic Command Line applications.*

* [cobra](https://github.com/spf13/cobra) - Commander for modern Go CLI interactions. Stars:`15.4K`.
* [urfave/cli](https://github.com/urfave/cli) - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli). Stars:`12.9K`.
* [kingpin](https://github.com/alecthomas/kingpin) - Command line and flag parser supporting sub commands. Stars:`2.8K`.
* [go-flags](https://github.com/jessevdk/go-flags) - go command line option parser. Stars:`1.6K`.
* [readline](https://github.com/chzyer/readline) - Pure golang implementation that provides most features in GNU-Readline under MIT license. Stars:`1.4K`.
* [Dnote](https://github.com/dnote/dnote) - A simple and end-to-end encrypted notebook for developers. Stars:`1.4K`.
* [docopt.go](https://github.com/docopt/docopt.go) - Command-line arguments parser that will make you smile. Stars:`1.2K`.
* [mitchellh/cli](https://github.com/mitchellh/cli) - Go library for implementing command-line interfaces. Stars:`1.1K`.
* [pflag](https://github.com/spf13/pflag) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags. Stars:`958`.
* [cli-init](https://github.com/tcnksm/gcli) - The easy way to start building Golang command line applications. Stars:`888`.
* [go-arg](https://github.com/alexflint/go-arg) - Struct-based argument parsing in Go. Stars:`812`.
* [complete](https://github.com/posener/complete) - Write bash completions in Go + Go command bash completion. Stars:`662`.
* [mow.cli](https://github.com/jawher/mow.cli) - Go library for building CLI applications with sophisticated flag and argument parsing and validation. Stars:`650`.
* [liner](https://github.com/peterh/liner) - Go readline-like library for command-line interfaces. Stars:`632`.
* [flaggy](https://github.com/integrii/flaggy) - A robust and idiomatic flags package with excellent subcommand support. Stars:`529`.
* [cli](https://github.com/mkideal/cli) - Feature-rich and easy to use command-line package based on golang struct tags. Stars:`507`.
* [ops](https://github.com/nanovms/ops) - Unikernel Builder/Orchestrator. Stars:`332`.
* [argparse](https://github.com/akamensky/argparse) - Command line argument parser inspired by Python's argparse module. Stars:`149`.
* [commandeer](https://github.com/jaffee/commandeer) - Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags. Stars:`104`.
* [flag](https://github.com/cosiner/flag) - Simple but powerful command line option parsing library for Go supporting subcommand. Stars:`103`.
* [sflags](https://github.com/octago/sflags) - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin and other libraries. Stars:`103`.
* [ukautz/clif](https://github.com/ukautz/clif) - Small command line interface framework. Stars:`98`.
* [wmenu](https://github.com/dixonwille/wmenu) - Easy to use menu structure for cli applications that prompts users to make choices. Stars:`95`.
* [job](https://github.com/liujianping/job) - JOB, make your short-term command as a long-term job. Stars:`67`.
* [cli](https://github.com/teris-io/cli) - Simple and complete API for building command line interfaces in Go. Stars:`59`.
* [env](https://github.com/codingconcepts/env) - Tag-based environment configuration for structs. Stars:`46`.
* [1build](https://github.com/gopinath-langote/1build) - Command line tool to frictionlessly manage project-specific commands. Stars:`45`.
* [wlog](https://github.com/dixonwille/wlog) - Simple logging interface that supports cross-platform color and concurrency. Stars:`40`.
* [gocmd](https://github.com/devfacet/gocmd) - Go library for building command line applications. Stars:`37`.
* [flagvar](https://github.com/sgreben/flagvar) - A collection of flag argument types for Go's standard `flag` package. Stars:`33`.
* [strumt](https://github.com/antham/strumt) - Library to create prompt chain. Stars:`32`.
* [cmdr](https://github.com/hedzr/cmdr) - A POSIX/GNU style, getopt-like command-line UI Go library. Stars:`27`.
* [go-getoptions](https://github.com/DavidGamba/go-getoptions) - Go option parser inspired on the flexibility of Perl’s GetOpt::Long. Stars:`19`.
* [argv](https://github.com/cosiner/argv) - Go library to split command line string as arguments array using the bash syntax. Stars:`19`.
* [go-commander](https://github.com/yitsushi/go-commander) - Go library to simplify CLI workflow. Stars:`15`.
* [sand](https://github.com/Zaba505/sand) - Simple API for creating interpreters and so much more. Stars:`9`.
* [ts](https://github.com/liujianping/ts) - Timestamp convert & compare tool. Stars:`7`.
* [clîr](https://github.com/leaanthony/clir) - A Simple and Clear CLI library. Dependency free. Stars:`3`.
* [cmd](https://github.com/posener/cmd) - Extends the standard `flag` package to support sub commands and more in idomatic way. Stars:`3`.
* [climax](http://github.com/tucnak/climax) - Alternative CLI with "human face", in spirit of Go command.
* [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli) - cli application framework with auto configuration and dependency injection.

### Advanced Console UIs

*Libraries for building Console Applications and Console User Interfaces.*

* [termui](https://github.com/gizak/termui) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib). Stars:`9.4K`.
* [gocui](https://github.com/jroimartin/gocui) - Minimalist Go library aimed at creating Console User Interfaces. Stars:`5.9K`.
* [termbox-go](https://github.com/nsf/termbox-go) - Termbox is a library for creating cross-platform text-based interfaces. Stars:`3.6K`.
* [go-prompt](https://github.com/c-bata/go-prompt) - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit). Stars:`2.7K`.
* [uiprogress](https://github.com/gosuri/uiprogress) - Flexible library to render progress bars in terminal applications. Stars:`1.6K`.
* [asciigraph](https://github.com/guptarohit/asciigraph) - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies. Stars:`1.2K`.
* [uilive](https://github.com/gosuri/uilive) - Library for updating terminal output in realtime. Stars:`1.1K`.
* [termdash](https://github.com/mum4k/termdash) - Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui). Stars:`908`.
* [progressbar](https://github.com/schollz/progressbar) - Basic thread-safe progress bar that works in every OS. Stars:`829`.
* [mpb](https://github.com/vbauerster/mpb) - Multi progress bar for terminal applications. Stars:`818`.
* [aurora](https://github.com/logrusorgru/aurora) - ANSI terminal colors that supports fmt.Printf/Sprintf. Stars:`758`.
* [uitable](https://github.com/gosuri/uitable) - Library to improve readability in terminal apps using tabular data. Stars:`537`.
* [go-colorable](https://github.com/mattn/go-colorable) - Colorable writer for windows. Stars:`400`.
* [go-isatty](https://github.com/mattn/go-isatty) - isatty for golang. Stars:`387`.
* [chalk](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output. Stars:`325`.
* [gookit/color](https://github.com/gookit/color) - Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows. Stars:`297`.
* [tabby](https://github.com/cheynewallace/tabby) - A tiny library for super simple Golang tables. Stars:`258`.
* [go-colortext](https://github.com/daviddengcn/go-colortext) - Go library for color output in terminals. Stars:`200`.
* [simpletable](https://github.com/alexeyco/simpletable) - Simple tables in terminal with Go. Stars:`183`.
* [cfmt](https://github.com/mingrammer/cfmt) - Contextual fmt inspired by bootstrap color classes. Stars:`70`.
* [tabular](https://github.com/InVisionApp/tabular) - Print ASCII tables from command line utilities without the need to pass large sets of data to the API. Stars:`35`.
* [colourize](https://github.com/TreyBastian/colourize) - Go library for ANSI colour text in terminals. Stars:`17`.
* [ctc](https://github.com/wzshiming/ctc) - The non-invasive cross-platform terminal color library does not need to modify the Print method. Stars:`13`.
* [go-ataman](https://github.com/workanator/go-ataman) - Go library for rendering ANSI colored text templates in terminals. Stars:`8`.
* [gommon/color](https://github.com/labstack/gommon/tree/master/color) - Style terminal text.

## Configuration

*Libraries for configuration parsing.*

* [viper](https://github.com/spf13/viper) - Go configuration with fangs. Stars:`10.9K`.
* [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) - Go library for managing configuration data from environment variables. Stars:`2.7K`.
* [godotenv](https://github.com/joho/godotenv) - Go port of Ruby's dotenv library (Loads environment variables from `.env`). Stars:`2.5K`.
* [ini](https://github.com/go-ini/ini) - Go package to read and write INI files. Stars:`1.8K`.
* [env](https://github.com/caarlos0/env) - Parse environment variables to Go structs (with defaults). Stars:`1.4K`.
* [konfig](https://github.com/lalamove/konfig) - Composable, observable and performant config handling for Go for the distributed processing era. Stars:`545`.
* [confita](https://github.com/heetch/confita) - Load configuration in cascade from multiple backends into a struct. Stars:`286`.
* [store](https://github.com/tucnak/store) - Lightweight configuration manager for Go. Stars:`248`.
* [config](https://github.com/olebedev/config) - JSON or YAML configuration wrapper with environment variables and flags parsing. Stars:`222`.
* [config](https://github.com/JeremyLoy/config) - Cloud native application configuration. Bind ENV to structs in only two lines. Stars:`221`.
* [joshbetz/config](https://github.com/joshbetz/config) - Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP. Stars:`198`.
* [hjson](https://github.com/hjson/hjson-go) - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments. Stars:`184`.
* [envconfig](https://github.com/vrischmann/envconfig) - Read your configuration from environment variables. Stars:`171`.
* [koanf](https://github.com/knadh/koanf) - Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line. Stars:`147`.
* [gookit/config](https://github.com/gookit/config) - application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge. Stars:`129`.
* [gcfg](https://github.com/go-gcfg/gcfg) - read INI-style configuration files into Go structs; supports user-defined types and subsections. Stars:`123`.
* [goConfig](https://github.com/crgimenes/goConfig) - Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file. Stars:`120`.
* [envh](https://github.com/antham/envh) - Helpers to manage environment variables. Stars:`92`.
* [envcfg](https://github.com/tomazk/envcfg) - Un-marshaling environment variables to Go structs. Stars:`90`.
* [gofigure](https://github.com/ian-kent/gofigure) - Go application configuration made easy. Stars:`58`.
* [configure](https://github.com/paked/configure) - Provides configuration through multiple sources, including JSON, flags and environment variables. Stars:`53`.
* [harvester](https://github.com/beatlabs/harvester) - Harvester, a easy to use static and dynamic configuration package supportig seeding, env vars and Consul integration. Stars:`49`.
* [xdg](https://github.com/OpenPeeDeeP/xdg) - Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html). Stars:`44`.
* [config](https://github.com/golobby/config) - A lightweight yet powerful config package for Go projects. Stars:`38`.
* [cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Minimalistic configuration reader (from files, ENV, and wherever you want). Stars:`30`.
* [go-up](https://github.com/ufoscout/go-up) - A simple configuration library with recursive placeholders resolution and no magic. Stars:`26`.
* [ingo](https://github.com/schachmat/ingo) - Flags persisted in an ini-like config file. Stars:`26`.
* [mini](https://github.com/sasbury/mini) - Golang package for parsing ini-style configuration files. Stars:`20`.
* [conflate](https://github.com/the4thamigo-uk/conflate) - Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema. Stars:`11`.
* [envconf](https://github.com/ian-kent/envconf) - Configuration from environment. Stars:`8`.
* [genv](https://github.com/sakirsensoy/genv) - Read environment variables easily with dotenv support. Stars:`8`.
* [sprbox](https://github.com/oblq/sprbox) - Build-environment aware toolbox factory and agnostic config parser (YAML, TOML, JSON and Environment vars). Stars:`5`.
* [go-ssm-config](https://github.com/ianlopshire/go-ssm-config) - Go utility for loading configuration parameters from AWS SSM (Parameter Store). Stars:`3`.
* [nasermirzaei89/env](https://github.com/nasermirzaei89/env) - Simple useful package for read environment variables. Stars:`2`.
* [onion](http://github.com/goraz/onion) - Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP.
* [gone/jconf](https://github.com/One-com/gone/tree/master/jconf) - Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.

## Continuous Integration

*Tools for help with continuous integration.*

* [drone](https://github.com/drone/drone) - Drone is a Continuous Integration platform built on Docker, written in Go. Stars:`20.3K`.
* [CDS](https://github.com/ovh/cds) - Enterprise-Grade CI/CD and DevOps Automation Open Source Platform. Stars:`2.6K`.
* [goveralls](https://github.com/mattn/goveralls) - Go integration for Coveralls.io continuous code coverage tracking system. Stars:`622`.
* [overalls](https://github.com/go-playground/overalls) - Multi-Package go project coverprofile for tools like goveralls. Stars:`99`.
* [duci](https://github.com/duck8823/duci) - A simple ci server no needs domain specific languages. Stars:`49`.
* [gomason](https://github.com/nikogura/gomason) - Test, Build, Sign, and Publish your go binaries from a clean workspace. Stars:`40`.
* [roveralls](https://github.com/LawrenceWoodman/roveralls) - Recursive coverage testing tool. Stars:`12`.

## CSS Preprocessors

*Libraries for preprocessing CSS files.*

* [gcss](https://github.com/yosssi/gcss) - Pure Go CSS Preprocessor. Stars:`427`.
* [go-libsass](https://github.com/wellington/go-libsass) - Go wrapper to the 100% Sass compatible libsass project. Stars:`147`.

## Data Structures

*Generic datastructures and algorithms in Go.*

* [gods](https://github.com/emirpasic/gods) - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc. Stars:`7.6K`.
* [go-datastructures](https://github.com/Workiva/go-datastructures) - Collection of useful, performant, and thread-safe data structures. Stars:`5.4K`.
* [golang-set](https://github.com/deckarep/golang-set) - Thread-Safe and Non-Thread-Safe high-performance sets for Go. Stars:`1.4K`.
* [boomfilters](https://github.com/tylertreat/BoomFilters) - Probabilistic data structures for processing continuous, unbounded streams. Stars:`1.2K`.
* [gota](https://github.com/kniren/gota) - Implementation of dataframes, series, and data wrangling methods for Go. Stars:`1.0K`.
* [roaring](https://github.com/RoaringBitmap/roaring) - Go package implementing compressed bitsets. Stars:`768`.
* [willf/bloom](https://github.com/willf/bloom) - Go package implementing Bloom filters. Stars:`734`.
* [hyperloglog](https://github.com/axiomhq/hyperloglog) - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction. Stars:`680`.
* [cuckoofilter](https://github.com/seiflotfy/cuckoofilter) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go. Stars:`570`.
* [bitset](https://github.com/willf/bitset) - Go package implementing bitsets. Stars:`531`.
* [trie](https://github.com/derekparker/trie) - Trie implementation in Go. Stars:`450`.
* [algorithms](https://github.com/shady831213/algorithms) - Algorithms and data structures.CLRS study. Stars:`339`.
* [gocache](https://github.com/eko/gocache) - A complete Go cache library with mutiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more. Stars:`338`.
* [go-geoindex](https://github.com/hailocab/go-geoindex) - In-memory geo index. Stars:`317`.
* [mafsa](https://github.com/smartystreets/mafsa) - MA-FSA implementation with Minimal Perfect Hashing. Stars:`280`.
* [goskiplist](https://github.com/ryszard/goskiplist) - Skip list implementation in Go. Stars:`215`.
* [hilbert](https://github.com/google/hilbert) - Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves. Stars:`192`.
* [merkletree](https://github.com/cbergoon/merkletree) - Implementation of a merkle tree providing an efficient and secure verification of the contents of data structures. Stars:`175`.
* [binpacker](https://github.com/zhuangsirui/binpacker) - Binary packer and unpacker helps user build custom binary stream. Stars:`132`.
* [bloom](https://github.com/zhenjl/bloom) - Bloom filters implemented in Go. Stars:`131`.
* [go-adaptive-radix-tree](https://github.com/plar/go-adaptive-radix-tree) - Go implementation of Adaptive Radix Tree. Stars:`125`.
* [ttlcache](https://github.com/diegobernardes/ttlcache) - In-memory LRU string-interface{} map with expiration for golang. Stars:`122`.
* [skiplist](https://github.com/MauriceGit/skiplist) - Very fast Go Skiplist implementation. Stars:`111`.
* [iter](https://github.com/disksing/iter) - Go implementation of C++ STL iterators and algorithms. Stars:`105`.
* [go-rquad](https://github.com/aurelien-rainone/go-rquad) - Region quadtrees with efficient point location and neighbour finding. Stars:`101`.
* [encoding](https://github.com/zhenjl/encoding) - Integer Compression Libraries for Go. Stars:`97`.
* [ring](https://github.com/TheTannerRyan/ring) - Go implementation of a high performance, thread safe bloom filter. Stars:`96`.
* [deque](https://github.com/gammazero/deque) - Fast ring-buffer deque (double-ended queue). Stars:`95`.
* [conjungo](https://github.com/InVisionApp/conjungo) - A small, powerful and flexible merge library. Stars:`84`.
* [bit](https://github.com/yourbasic/bit) - Golang set data structure with bonus bit-twiddling functions. Stars:`75`.
* [skiplist](https://github.com/gansidui/skiplist) - Skiplist implementation in Go. Stars:`67`.
* [levenshtein](https://github.com/agnivade/levenshtein) - Implementation to calculate levenshtein distance in Go. Stars:`62`.
* [goconcurrentqueue](https://github.com/enriquebris/goconcurrentqueue) - Concurrent FIFO queue. Stars:`55`.
* [bloom](https://github.com/yourbasic/bloom) - Golang Bloom filter implementation. Stars:`47`.
* [count-min-log](https://github.com/seiflotfy/count-min-log) - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory). Stars:`44`.
* [go-mcache](https://github.com/OrlovEvgeny/go-mcache) - Fast in-memory key:value store/cache library. Pointer caches. Stars:`43`.
* [levenshtein](https://github.com/agext/levenshtein) - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix. Stars:`37`.
* [remember-go](https://github.com/rocketlaunchr/remember-go) - A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory). Stars:`33`.
* [concurrent-writer](https://github.com/free/concurrent-writer) - Highly concurrent drop-in replacement for `bufio.Writer`. Stars:`26`.
* [crunch](https://github.com/superwhiskers/crunch) - Go package implementing buffers for handling various datatypes easily. Stars:`23`.
* [goset](https://github.com/zoumo/goset) - A useful Set collection implementation for Go. Stars:`19`.
* [pipeline](https://github.com/hyfather/pipeline) - An implementation of pipelines with fan-in and fan-out. Stars:`19`.
* [typ](https://github.com/gurukami/typ) - Null Types, Safe primitive type conversion and fetching value from complex structures. Stars:`15`.
* [deque](https://github.com/edwingeng/deque) - A highly optimized double-ended queue. Stars:`14`.
* [dict](https://github.com/srfrog/dict) - Python-like dictionaries (dict) for Go. Stars:`12`.
* [go-ef](https://github.com/amallia/go-ef) - A Go implementation of the Elias-Fano encoding. Stars:`11`.
* [hide](https://github.com/emvi/hide) - ID type with marshalling to/from hash to prevent sending IDs to clients. Stars:`11`.
* [null](https://github.com/emvi/null) - Nullable Go types that can be marshalled/unmarshalled to/from JSON. Stars:`8`.
* [mspm](https://github.com/BlackRabbitt/mspm) - Multi-String Pattern Matching Algorithm for information retrieval. Stars:`8`.
* [set](https://github.com/StudioSol/set) - Simple set data structure implementation in Go using LinkedHashMap. Stars:`7`.
* [treap](https://github.com/perdata/treap) - Persistent, fast ordered map using tree heaps. Stars:`7`.
* [timedmap](https://github.com/zekroTJA/timedmap) - Map with expiring key-value pairs. Stars:`6`.
* [gofal](https://github.com/xxjwxc/gofal) - fractional api for Go. Stars:`5`.
* [ptrie](https://github.com/viant/ptrie) - An implementation of prefix tree. Stars:`3`.
* [parsefields](https://github.com/MonaxGT/parsefields) - Tools for parse JSON-like logs for collecting unique fields and events. Stars:`3`.

## Database

*Databases implemented in Go.*

* [prometheus](https://github.com/prometheus/prometheus) - Monitoring system and time series database. Stars:`28.6K`.
* [tidb](https://github.com/pingcap/tidb) - TiDB is a distributed SQL database. Inspired by the design of Google F1. Stars:`22.2K`.
* [influxdb](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics. Stars:`18.1K`.
* [cockroach](https://github.com/cockroachdb/cockroach) - Scalable, Geo-Replicated, Transactional Datastore. Stars:`17.6K`.
* [dgraph](https://github.com/dgraph-io/dgraph) - Scalable, Distributed, Low Latency, High Throughput Graph Database. Stars:`12.2K`.
* [bolt](https://github.com/boltdb/bolt) - Low-level key/value database for Go. Stars:`10.4K`.
* [groupcache](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases. Stars:`8.1K`.
* [badger](https://github.com/dgraph-io/badger) - Fast key-value store in Go. Stars:`7.2K`.
* [rqlite](https://github.com/rqlite/rqlite) - The lightweight, distributed, relational database built on SQLite. Stars:`5.3K`.
* [BigCache](https://github.com/allegro/bigcache) - Efficient key/value cache for gigabytes of data. Stars:`3.5K`.
* [goleveldb](https://github.com/syndtr/goleveldb) - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go. Stars:`3.5K`.
* [go-cache](https://github.com/pmylund/go-cache) - In-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications. Stars:`3.4K`.
* [ledisdb](https://github.com/siddontang/ledisdb) - Ledisdb is a high performance NoSQL like Redis based on LevelDB. Stars:`3.2K`.
* [buntdb](https://github.com/tidwall/buntdb) - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support. Stars:`2.6K`.
* [tiedot](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang. Stars:`2.4K`.
* [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics) - fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL. Stars:`1.6K`.
* [cache2go](https://github.com/muesli/cache2go) - In-memory key:value cache which supports automatic invalidation based on timeouts. Stars:`1.2K`.
* [nutsdb](https://github.com/xujiajun/nutsdb) - Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as  list, set, sorted set. Stars:`1.1K`.
* [GCache](https://github.com/bluele/gcache) - Cache library with support for expirable Cache, LFU, LRU and ARC. Stars:`1.0K`.
* [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL) - CovenantSQL is a SQL database on blockchain. Stars:`972`.
* [diskv](https://github.com/peterbourgon/diskv) - Home-grown disk-backed key-value store. Stars:`842`.
* [moss](https://github.com/couchbase/moss) - Moss is a simple LSM key-value storage engine written in 100% Go. Stars:`745`.
* [fastcache](https://github.com/VictoriaMetrics/fastcache) - fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead. Stars:`664`.
* [eliasdb](https://github.com/krotik/eliasdb) - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language. Stars:`556`.
* [levigo](https://github.com/jmhodges/levigo) - Levigo is a Go wrapper for LevelDB. Stars:`377`.
* [Bitcask](https://github.com/prologic/bitcask) - Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL). Stars:`255`.
* [pudge](https://github.com/recoilme/pudge) - Fast and simple  key/value store written using Go's standard library. Stars:`238`.
* [piladb](https://github.com/fern4lvarez/piladb) - Lightweight RESTful database engine based on stack data structures. Stars:`172`.
* [Vasto](https://github.com/chrislusf/vasto) - A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption. Stars:`166`.
* [Kivik](https://github.com/go-kivik/kivik) - Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases. Stars:`139`.
* [slowpoke](https://github.com/recoilme/slowpoke) - Key-value store with persistence. Stars:`94`.
* [Scribble](https://github.com/nanobox-io/golang-scribble) - Tiny flat file JSON store. Stars:`82`.
* [couchcache](https://github.com/codingsince1985/couchcache) - RESTful caching micro-service backed by Couchbase server. Stars:`46`.
* [bcache](https://github.com/iwanbk/bcache) - Eventually consistent distributed in-memory  cache Go library. Stars:`41`.
* [clusteredBigCache](https://github.com/oaStuff/clusteredBigCache) - BigCache with clustering support and individual item expiration. Stars:`32`.
* [cache](https://github.com/akyoto/cache) - In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage. Stars:`29`.
* [Coffer](https://github.com/claygod/coffer) - Simple ACID key-value database that supports transactions. Stars:`17`.
* [tempdb](https://github.com/rafaeljesus/tempdb) - Key-value store for temporary items. Stars:`14`.
* [gorocksdb](https://github.com/kapitan-k/gorocksdb) - Gorocksdb is a wrapper for [RocksDB](https://rocksdb.org) written in Go. Stars:`12`.
* [tracedb](https://github.com/unit-io/tracedb) - Fast timeseries database for IoT, realtime messaging  applications. Access tracedb with pubsub over tcp or websocket using github.com/unit-io/trace application. Stars:`5`.

*Database schema migration.*

* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library. Stars:`3.5K`.
* [sql-migrate](https://github.com/rubenv/sql-migrate) - Database migration tool. Allows embedding migrations into the application using go-bindata. Stars:`1.6K`.
* [skeema](https://github.com/skeema/skeema) - Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools. Stars:`500`.
* [gormigrate](https://github.com/go-gormigrate/gormigrate) - Database schema migration helper for Gorm ORM. Stars:`393`.
* [goose](https://github.com/steinbacher/goose) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts. Stars:`128`.
* [darwin](https://github.com/GuiaBolso/darwin) - Database schema evolution library for Go. Stars:`96`.
* [migrator](https://github.com/lopezator/migrator) - Dead simple Go database migration library. Stars:`83`.
* [go-pg-migrations](https://github.com/robinjoseph08/go-pg-migrations) - A Go package to help write migrations with go-pg/pg. Stars:`38`.
* [gondolier](https://github.com/emvi/gondolier) - Database migration library using struct decorators. Stars:`26`.
* [pravasan](https://github.com/pravasan/pravasan) - Simple Migration tool - currently for MySQL but planning to soon support Postgres, SQLite, MongoDB, etc. Stars:`24`.
* [go-fixtures](https://github.com/RichardKnop/go-fixtures) - Django style fixtures for Golang's excellent built-in database/sql library. Stars:`22`.
* [avro](https://github.com/khezen/avro) - Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes. Stars:`10`.
* [schema](https://github.com/adlio/schema) - Library to embed schema migrations for database/sql-compatible databases inside your Go binaries. Stars:`4`.
* [soda](https://github.com/gobuffalo/pop/tree/master/soda) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.

*Database tools.*

* [vitess](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services. Stars:`9.3K`.
* [pgweb](https://github.com/sosedoff/pgweb) - Web-based PostgreSQL database browser. Stars:`6.2K`.
* [kingshard](https://github.com/flike/kingshard) - kingshard is a high performance proxy for MySQL powered by Golang. Stars:`5.0K`.
* [orchestrator](https://github.com/github/orchestrator) - MySQL replication topology manager & visualizer. Stars:`3.3K`.
* [go-mysql-elasticsearch](https://github.com/siddontang/go-mysql-elasticsearch) - Sync your MySQL data into Elasticsearch automatically. Stars:`2.7K`.
* [pREST](https://github.com/nuveo/prest) - Serve a RESTful API from any PostgreSQL database. Stars:`2.2K`.
* [go-mysql](https://github.com/siddontang/go-mysql) - Go toolset to handle MySQL protocol and replication. Stars:`2.2K`.
* [chproxy](https://github.com/Vertamedia/chproxy) - HTTP proxy for ClickHouse database. Stars:`354`.
* [clickhouse-bulk](https://github.com/nikepan/clickhouse-bulk) - Collects small insterts and sends big requests to ClickHouse servers. Stars:`165`.
* [myreplication](https://github.com/2tvenom/myreplication) - MySql binary log replication listener. Supports statement and row based replication. Stars:`155`.
* [octillery](https://github.com/knocknote/octillery) - Go package for sharding databases ( Supports every ORM or raw SQL ). Stars:`71`.
* [dbbench](https://github.com/sj14/dbbench) - Database benchmarking tool with support for several databases and scripts. Stars:`31`.
* [prep](https://github.com/hexdigest/prep) - Use prepared SQL statements without changing your code. Stars:`25`.
* [datagen](https://github.com/codingconcepts/datagen) - A fast data generator that's multi-table aware and supports multi-row DML. Stars:`13`.
* [rwdb](https://github.com/andizzle/rwdb) - rwdb provides read replica capability for multiple database servers setup. Stars:`10`.
* [bucket](https://github.com/PumpkinSeed/bucket) - Optimized data structure framework for Couchbase specialized on one bucket usage. Stars:`6`.

*SQL query builder, libraries for building and using SQL.*

* [Squirrel](https://github.com/Masterminds/squirrel) - Go library that helps you build SQL queries. Stars:`2.7K`.
* [xo](https://github.com/knq/xo) - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server. Stars:`2.3K`.
* [gendry](https://github.com/didi/gendry) - Non-invasive SQL builder and powerful data binder. Stars:`928`.
* [goqu](https://github.com/doug-martin/goqu) - Idiomatic SQL builder and query library. Stars:`712`.
* [Dotsql](https://github.com/gchaincl/dotsql) - Go library that helps you keep sql files in one place and use them with ease. Stars:`489`.
* [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx) - Powerful data retrieval methods as well as DB-agnostic query building capabilities. Stars:`453`.
* [sqrl](https://github.com/elgris/sqrl) - SQL query builder, fork of Squirrel with improved performance. Stars:`197`.
* [jet](https://github.com/go-jet/jet) - Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure. Stars:`189`.
* [dbq](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go. Stars:`84`.
* [igor](https://github.com/galeone/igor) - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax. Stars:`78`.
* [godbal](https://github.com/xujiajun/godbal) - Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily. Stars:`50`.
* [sqlf](https://github.com/leporo/sqlf) - Fast SQL query builder. Stars:`7`.
* [qry](https://github.com/HnH/qry) - Tool that generates constants from files with raw SQL queries. Stars:`4`.
* [ormlite](https://github.com/pupizoid/ormlite) - Lightweight package containing some ORM-like features and helpers for sqlite databases.
* [scaneo](https://github.com/variadico/scaneo) - Generate Go code to convert database rows into arbitrary structs.
* [Squalus](https://gitlab.com/qosenergy/squalus) - Thin layer over the Go SQL package that makes it easier to perform queries.

## Database Drivers

*Libraries for connecting and operating databases.*

* Relational Databases
    * [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go. Stars:`8.9K`.
    * [pq](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql. Stars:`5.6K`.
    * [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that uses database/sql. Stars:`3.7K`.
    * [pgx](https://github.com/jackc/pgx) - PostgreSQL driver supporting features beyond those exposed by database/sql. Stars:`2.3K`.
    * [go-mssqldb](https://github.com/denisenkom/go-mssqldb) - Microsoft MSSQL driver for Go. Stars:`1.1K`.
    * [go-oci8](https://github.com/mattn/go-oci8) - Oracle driver for go that uses database/sql. Stars:`429`.
    * [goracle](https://github.com/go-goracle/goracle) - Oracle driver for Go, using the ODPI-C driver. Stars:`280`.
    * [firebirdsql](https://github.com/nakagami/firebirdsql) - Firebird RDBMS SQL driver for Go. Stars:`114`.
    * [go-adodb](https://github.com/mattn/go-adodb) - Microsoft ActiveX Object DataBase driver for go that uses database/sql. Stars:`96`.
    * [gofreetds](https://github.com/minus5/gofreetds) - Microsoft MSSQL driver. Go wrapper over [FreeTDS](http://www.freetds.org). Stars:`94`.
    * [avatica](https://github.com/apache/calcite-avatica-go) - Apache Avatica/Phoenix SQL driver for database/sql. Stars:`46`.
    * [bgc](https://github.com/viant/bgc) - Datastore Connectivity for BigQuery for go. Stars:`12`.

* NoSQL Databases
    * [redis](https://github.com/go-redis/redis) - Redis client for Golang. Stars:`7.8K`.
    * [redigo](https://github.com/gomodule/redigo) - Redigo is a Go client for the Redis database. Stars:`6.9K`.
    * [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver for the Go language. Stars:`3.9K`.
    * [mgo](https://github.com/globalsign/mgo) - (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms. Stars:`1.7K`.
    * [gorethink](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB. Stars:`1.5K`.
    * [neoism](https://github.com/jmcvetta/neoism) - Neo4j client for Golang. Stars:`362`.
    * [redeo](https://github.com/bsm/redeo) - Redis-protocol compatible TCP servers/services. Stars:`361`.
    * [aerospike-client-go](https://github.com/aerospike/aerospike-client-go) - Aerospike client in Go language. Stars:`315`.
    * [gocb](https://github.com/couchbase/gocb) - Official Couchbase Go SDK. Stars:`307`.
    * [go-couchbase](https://github.com/couchbase/go-couchbase) - Couchbase client in Go. Stars:`297`.
    * [go-rejson](https://github.com/nitishm/go-rejson) - Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease. Stars:`109`.
    * [Neo4j-GO](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang. Stars:`74`.
    * [godis](https://github.com/piaohao/godis) - redis client implement by golang, inspired by jedis. Stars:`67`.
    * [dynago](https://github.com/underarmour/dynago) - Dynago is a principle of least surprise client for DynamoDB. Stars:`66`.
    * [arangolite](https://github.com/solher/arangolite) - Lightweight golang driver for ArangoDB. Stars:`66`.
    * [mgm](https://github.com/kamva/mgm) - MongoDB model-based ODM for Go (based on official MongoDB driver). Stars:`38`.
    * [go-pilosa](https://github.com/pilosa/go-pilosa) - Go client library for Pilosa. Stars:`33`.
    * [forestdb](https://github.com/couchbase/goforestdb) - Go bindings for ForestDB. Stars:`28`.
    * [neo4j](https://github.com/cihangir/neo4j) - Neo4j Rest API Bindings for Golang. Stars:`25`.
    * [goriak](https://github.com/zegl/goriak) - Go language driver for Riak KV. Stars:`24`.
    * [xredis](https://github.com/shomali11/xredis) - Typesafe, customizable, clean & easy to use Redis client. Stars:`10`.
    * [godscache](https://github.com/defcronyke/godscache) - A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached. Stars:`6`.
    * [asc](https://github.com/viant/asc) - Datastore Connectivity for Aerospike for go. Stars:`4`.
    * [gocql](http://gocql.github.io) - Go language driver for Apache Cassandra.
    * [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language.

* Search and Analytic Databases.
    * [bleve](https://github.com/blevesearch/bleve) - Modern text indexing library for go. Stars:`6.2K`.
    * [riot](https://github.com/go-ego/riot) - Go Open Source, Distributed, Simple and efficient Search Engine. Stars:`5.0K`.
    * [elastic](https://github.com/olivere/elastic) - Elasticsearch client for Go. Stars:`4.7K`.
    * [go-elasticsearch](https://github.com/elastic/go-elasticsearch) - Official Elasticsearch client for Go. Stars:`2.1K`.
    * [elastigo](https://github.com/mattbaird/elastigo) - Elasticsearch client library. Stars:`954`.
    * [elasticsql](https://github.com/cch123/elasticsql) - Convert sql to elasticsearch dsl in Go. Stars:`478`.
    * [skizze](https://github.com/seiflotfy/skizze) - probabilistic data-structures service and storage. Stars:`69`.
    * [goes](https://github.com/OwnLocal/goes) - Library to interact with Elasticsearch. Stars:`24`.

* Multiple Backends.
    * [cayley](https://github.com/google/cayley) - Graph database with support for multiple backends. Stars:`13.1K`.
    * [gokv](https://github.com/philippgille/gokv) - Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more). Stars:`155`.
    * [cachego](https://github.com/fabiorphp/cachego) - Golang Cache component for multiple drivers. Stars:`113`.
    * [dsc](https://github.com/viant/dsc) - Datastore connectivity for SQL, NoSQL, structured files. Stars:`14`.

## Date and Time

*Libraries for working with dates and times.*

* [now](https://github.com/jinzhu/now) - Now is a time toolkit for golang. Stars:`2.4K`.
* [dateparse](https://github.com/araddon/dateparse) - Parse date's without knowing format in advance. Stars:`962`.
* [carbon](https://github.com/uniplaces/carbon) - Simple Time extension with a lot of util methods, ported from PHP Carbon library. Stars:`383`.
* [durafmt](https://github.com/hako/durafmt) - Time duration formatting library for Go. Stars:`272`.
* [timeutil](https://github.com/leekchan/timeutil) - Useful extensions (Timedelta, Strftime, ...) to the golang's time package. Stars:`173`.
* [iso8601](https://github.com/relvacode/iso8601) - Efficiently parse ISO8601 date-times without regex. Stars:`69`.
* [timespan](https://github.com/SaidinWoT/timespan) - For interacting with intervals of time, defined as a start time and a duration. Stars:`67`.
* [go-persian-calendar](https://github.com/yaa110/go-persian-calendar) - The implementation of the Persian (Solar Hijri) Calendar in Go (golang). Stars:`66`.
* [date](https://github.com/rickb777/date) - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day. Stars:`34`.
* [feiertage](https://github.com/wlbr/feiertage) - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving... Stars:`24`.
* [go-sunrise](https://github.com/nathan-osman/go-sunrise) - Calculate the sunrise and sunset times for a given location. Stars:`17`.
* [kair](https://github.com/GuilhermeCaruso/kair) - Date and Time - Golang Formatting Library. Stars:`12`.
* [NullTime](https://github.com/kirillDanshin/nulltime) - Nullable `time.Time`. Stars:`9`.
* [tuesday](https://github.com/osteele/tuesday) - Ruby-compatible Strftime function. Stars:`7`.
* [cronrange](https://github.com/1set/cronrange) - Parses Cron-style time range expressions, checks if the given time is within any ranges. Stars:`6`.
* [strftime](https://github.com/awoodbeck/strftime) - C99-compatible strftime formatter. Stars:`4`.
* [go-week](https://github.com/stoewer/go-week) - An efficient package to work with ISO8601 week dates. Stars:`2`.

## Distributed Systems

*Packages that help with building Distributed Systems.*

* [go-kit](https://github.com/go-kit/kit) - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc. Stars:`16.0K`.
* [grpc-go](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC. Stars:`10.5K`.
* [micro](https://github.com/micro/micro) - Pluggable microservice toolkit and distributed systems platform. Stars:`7.4K`.
* [NATS](https://github.com/nats-io/gnatsd) - Lightweight, high performance messaging system for microservices, IoT, and cloud native systems. Stars:`7.1K`.
* [rpcx](https://github.com/smallnest/rpcx) - Distributed pluggable RPC service framework like alibaba Dubbo. Stars:`4.3K`.
* [tendermint](https://github.com/tendermint/tendermint) - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols. Stars:`3.5K`.
* [torrent](https://github.com/anacrolix/torrent) - BitTorrent client package. Stars:`3.3K`.
* [raft](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol, by HashiCorp. Stars:`3.2K`.
* [dragonboat](https://github.com/lni/dragonboat) - A feature complete and high performance multi-group Raft library in Go. Stars:`2.8K`.
* [glow](https://github.com/chrislusf/glow) - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go. Stars:`2.7K`.
* [gleam](https://github.com/chrislusf/gleam) - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed. Stars:`2.4K`.
* [KrakenD](https://github.com/devopsfaith/krakend) - Ultra performant API Gateway framework with middlewares. Stars:`2.3K`.
* [emitter-io](https://github.com/emitter-io/emitter) - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love. Stars:`2.2K`.
* [liftbridge](https://github.com/liftbridge-io/liftbridge) - Lightweight, fault-tolerant message streams for NATS. Stars:`1.4K`.
* [hprose](https://github.com/hprose/hprose-golang) - Very newbility RPC Library, support 25+ languages now. Stars:`1.1K`.
* [ringpop-go](https://github.com/uber/ringpop-go) - Scalable, fault-tolerant application-layer sharding for Go applications. Stars:`599`.
* [gorpc](https://github.com/valyala/gorpc) - Simple, fast and scalable RPC library for high load. Stars:`575`.
* [go-health](https://github.com/InVisionApp/go-health) - Library for enabling asynchronous dependency health checks in your service. Stars:`519`.
* [rain](https://github.com/cenkalti/rain) - BitTorrent client and library. Stars:`387`.
* [digota](https://github.com/digota/digota) - grpc ecommerce microservice. Stars:`322`.
* [sleuth](https://github.com/ursiform/sleuth) - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)). Stars:`312`.
* [go-sundheit](https://github.com/AppsFlyer/go-sundheit) - A library built to provide support for defining async service health checks for golang services. Stars:`279`.
* [go-jump](https://github.com/dgryski/go-jump) - Port of Google's "Jump" Consistent Hash function. Stars:`276`.
* [consistent](https://github.com/buraksezer/consistent) - Consistent hashing with bounded loads. Stars:`258`.
* [dht](https://github.com/anacrolix/dht) - BitTorrent Kademlia DHT implementation. Stars:`144`.
* [jsonrpc](https://github.com/osamingo/jsonrpc) - The jsonrpc package helps implement of JSON-RPC 2.0. Stars:`118`.
* [jsonrpc](https://github.com/ybbus/jsonrpc) - JSON-RPC 2.0 HTTP client implementation. Stars:`111`.
* [redis-lock](https://github.com/bsm/redislock) - Simplified distributed locking implementation using Redis. Stars:`82`.
* [celeriac](https://github.com/svcavallar/celeriac.v1) - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go. Stars:`58`.
* [doublejump](https://github.com/edwingeng/doublejump) - A revamped Google's jump consistent hash. Stars:`47`.
* [drmaa](https://github.com/dgruber/drmaa) - Job submission library for cluster schedulers based on the DRMAA standard. Stars:`28`.
* [flowgraph](https://github.com/vectaport/flowgraph) - flow-based programming package. Stars:`25`.
* [outboxer](https://github.com/italolelis/outboxer) - Outboxer is a go library that implements the outbox pattern. Stars:`23`.
* [dynatomic](https://github.com/tylfin/dynatomic) - A library for using DynamoDB as an atomic counter. Stars:`10`.
* [resgate](https://resgate.io/) - Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
* [pglock](https://cirello.io/pglock) - PostgreSQL-backed distributed locking implementation.
* [raft](https://github.com/coreos/etcd/tree/master/raft) - Go implementation of the Raft consensus protocol, by CoreOS.
* [dynamolock](https://cirello.io/dynamolock) - DynamoDB-backed distributed locking implementation.
* [dot](https://github.com/dotchain/dot/) - distributed sync using operational transformation/OT.

## Dynamic DNS

*Tools for updating dynamic DNS records.*

* [GoDNS](https://github.com/timothyye/godns) - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go. Stars:`519`.
* [DDNS](https://github.com/skibish/ddns) - Personal DDNS client with Digital Ocean Networking DNS as backend. Stars:`122`.
* [dyndns](https://gitlab.com/alcastle/dyndns) - Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.

## Email

*Libraries and tools that implement email creation and sending.*

* [MailHog](https://github.com/mailhog/MailHog) - Email and SMTP testing with web and API interface. Stars:`5.8K`.
* [hermes](https://github.com/matcornic/hermes) - Golang package that generates clean, responsive HTML e-mails. Stars:`1.7K`.
* [email](https://github.com/jordan-wright/email) - A robust and flexible email library for Go. Stars:`1.2K`.
* [go-imap](https://github.com/emersion/go-imap) - IMAP library for clients and servers. Stars:`858`.
* [SendGrid](https://github.com/sendgrid/sendgrid-go) - SendGrid's Go library for sending email. Stars:`554`.
* [mailgun-go](https://github.com/mailgun/mailgun-go) - Go library for sending mail with the Mailgun API. Stars:`425`.
* [Hectane](https://github.com/hectane/hectane) - Lightweight SMTP client providing an HTTP API. Stars:`173`.
* [douceur](https://github.com/aymerick/douceur) - CSS inliner for your HTML emails. Stars:`167`.
* [go-message](https://github.com/emersion/go-message) - Streaming library for the Internet Message Format and mail messages. Stars:`131`.
* [smtp](https://github.com/mailhog/smtp) - SMTP server protocol state machine. Stars:`53`.
* [go-dkim](https://github.com/toorop/go-dkim) - DKIM library, to sign & verify email. Stars:`52`.
* [go-premailer](https://github.com/vanng822/go-premailer) - Inline styling for HTML mail in Go. Stars:`43`.
* [mailchain](https://github.com/mailchain/mailchain) - Send encrypted emails to blockchain addresses written in Go. Stars:`31`.
* [go-simple-mail](https://github.com/xhit/go-simple-mail) - Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send. Stars:`12`.
* [chasquid](https://blitiri.com.ar/p/chasquid) - SMTP server written in Go.

## Embeddable Scripting Languages

*Embedding other languages inside your go code.*

* [otto](https://github.com/robertkrimen/otto) - JavaScript interpreter written in Go. Stars:`5.1K`.
* [gopher-lua](https://github.com/yuin/gopher-lua) - Lua 5.1 VM and compiler written in Go. Stars:`3.2K`.
* [go-lua](https://github.com/Shopify/go-lua) - Port of the Lua 5.2 VM to pure Go. Stars:`1.8K`.
* [tengo](https://github.com/d5/tengo) - Bytecode compiled script language for Go. Stars:`1.5K`.
* [expr](https://github.com/antonmedv/expr) - an engine that can evaluate expressions. Stars:`1.0K`.
* [go-python](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API. Stars:`982`.
* [anko](https://github.com/mattn/anko) - Scriptable interpreter written in Go. Stars:`977`.
* [go-php](https://github.com/deuill/go-php) - PHP bindings for Go. Stars:`723`.
* [go-duktape](https://github.com/olebedev/go-duktape) - Duktape JavaScript engine bindings for Go. Stars:`676`.
* [golua](https://github.com/aarzilli/golua) - Go bindings for Lua C API. Stars:`457`.
* [gisp](https://github.com/jcla1/gisp) - Simple LISP in Go. Stars:`432`.
* [cel-go](https://github.com/google/cel-go) - Fast, portable, non-Turing complete expression evaluation with gradual typing. Stars:`348`.
* [gval](https://github.com/PaesslerAG/gval) - A highly customizable expression language written in Go. Stars:`175`.
* [gentee](https://github.com/gentee/gentee) - Embeddable scripting programming language. Stars:`41`.
* [binder](https://github.com/alexeyco/binder) - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua). Stars:`34`.
* [purl](https://github.com/ian-kent/purl) - Perl 5.18.2 embedded in Go. Stars:`29`.
* [ngaro](https://github.com/db47h/ngaro) - Embeddable Ngaro VM implementation enabling scripting in Retro. Stars:`19`.

## Error Handling

*Libraries for handling errors.*

* [errors](https://github.com/pkg/errors) - Package that provides simple error handling primitives. Stars:`5.4K`.
* [go-multierror](https://github.com/hashicorp/go-multierror) - Go (golang) package for representing a list of errors as a single error. Stars:`798`.
* [errorx](https://github.com/joomcode/errorx) - A feature rich error package with stack traces, composition of errors and more. Stars:`600`.
* [tracerr](https://github.com/ztrue/tracerr) - Golang errors with stack trace and source fragments. Stars:`576`.
* [eris](https://github.com/rotisserie/eris) - A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors. Stars:`426`.
* [errlog](https://github.com/snwfdhmp/errlog) - Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place. Stars:`275`.
* [emperror](https://github.com/emperror/emperror) - Error handling tools and best practices for Go libraries and applications. Stars:`93`.
* [errors](https://github.com/emperror/errors) - Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives. Stars:`37`.
* [werr](https://github.com/txgruppi/werr) - Error Wrapper creates an wrapper for the error type in Go which captures the File, Line and Stack of where it was called. Stars:`12`.
* [errors](https://github.com/neuronlabs/errors) - Simple golang error handling with classification primitives. Stars:`3`.
* [Falcon](https://github.com/SonicRoshan/falcon) - A Simple Yet Highly Powerful Package For Error Handling. Stars:`2`.

## Files

*Libraries for handling files and file systems.*

* [afero](https://github.com/spf13/afero) - FileSystem Abstraction System for Go. Stars:`2.5K`.
* [pdfcpu](https://github.com/hhrutter/pdfcpu) - PDF processor. Stars:`1.3K`.
* [notify](https://github.com/rjeczalik/notify) - File system event notification library with simple API, similar to os/signal. Stars:`519`.
* [bigfile](https://github.com/bigfile/bigfile) - A file transfer system, support to manage files with http api, rpc call and ftp client. Stars:`96`.
* [opc](https://github.com/qmuntal/opc) - Load Open Packaging Conventions (OPC) files for Go. Stars:`63`.
* [go-csv-tag](https://github.com/artonge/go-csv-tag) - Load csv file using tag. Stars:`61`.
* [skywalker](https://github.com/dixonwille/skywalker) - Package to allow one to concurrently go through a filesystem with ease. Stars:`53`.
* [vfs](https://github.com/C2FO/vfs) - A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS. Stars:`40`.
* [tarfs](https://github.com/posener/tarfs) - Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files. Stars:`39`.
* [afs](https://github.com/viant/afs) - Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go. Stars:`27`.
* [go-exiftool](https://github.com/barasher/go-exiftool) - Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...). Stars:`20`.
* [go-gtfs](https://github.com/artonge/go-gtfs) - Load gtfs files in go. Stars:`18`.
* [checksum](https://github.com/codingsince1985/checksum) - Compute message digest, like MD5 and SHA256, for large files. Stars:`14`.
* [go-decent-copy](https://github.com/hugocarreira/go-decent-copy) - Copy files for humans. Stars:`14`.
* [flop](https://github.com/homedepot/flop) - File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html). Stars:`13`.
* [parquet](https://github.com/parsyl/parquet) - Read and write [parquet](https://parquet.apache.org) files. Stars:`5`.
* [stl](https://gitlab.com/russoj88/stl) - Modules to read and write STL (stereolithography) files.  Concurrent algorithm for reading.

## Financial

*Packages for accounting and finance.*

* [decimal](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers. Stars:`1.9K`.
* [go-money](https://github.com/rhymond/go-money) - Implementation of Fowler's Money pattern. Stars:`680`.
* [go-finance](https://github.com/FlashBoys/go-finance) - Comprehensive financial markets data in Go. Stars:`537`.
* [accounting](https://github.com/leekchan/accounting) - money and currency formatting for golang. Stars:`527`.
* [techan](https://github.com/sdcoffey/techan) - Technical analysis library with advanced market analysis and trading strategies. Stars:`212`.
* [orderbook](https://github.com/i25959341/orderbook) - Matching Engine for Limit Order Book in Golang. Stars:`116`.
* [ofxgo](https://github.com/aclindsa/ofxgo) - Query OFX servers and/or parse the responses (with example command-line client). Stars:`71`.
* [vat](https://github.com/dannyvankooten/vat) - VAT number validation & EU VAT rates. Stars:`64`.
* [transaction](https://github.com/claygod/transaction) - Embedded transactional database of accounts, running in multithreaded mode. Stars:`62`.
* [go-finance](https://github.com/alpeb/go-finance) - Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations. Stars:`47`.
* [currency](https://github.com/bnkamalesh/currency) - High performant & accurate currency computation package. Stars:`18`.
* [go-finance](https://github.com/pieterclaerhout/go-finance) - Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers. Stars:`2`.

## Forms

*Libraries for working with forms.*

* [nosurf](https://github.com/justinas/nosurf) - CSRF protection middleware for Go. Stars:`1.0K`.
* [binding](https://github.com/mholt/binding) - Binds form and JSON data from net/http Request to struct. Stars:`761`.
* [gorilla/csrf](https://github.com/gorilla/csrf) - CSRF protection for Go web applications & services. Stars:`488`.
* [form](https://github.com/go-playground/form) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support. Stars:`383`.
* [conform](https://github.com/leebenson/conform) - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags. Stars:`182`.
* [formam](https://github.com/monoculum/formam) - decode form's values into a struct. Stars:`135`.
* [forms](https://github.com/albrow/forms) - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files. Stars:`107`.
* [bind](https://github.com/robfig/bind) - Bind form data to any Go values. Stars:`24`.
* [queryparam](https://github.com/tomwright/queryparam) - Decode `url.Values` into usable struct values of standard or custom types. Stars:`1`.

## Functional

*Packages to support functional programming in Go.*

* [go-underscore](https://github.com/tobyhede/go-underscore) - Useful collection of helpfully functional Go collection utilities. Stars:`1.1K`.
* [fpGo](https://github.com/TeaEntityLab/fpGo) - Monad, Functional Programming features for Golang. Stars:`133`.
* [fuego](https://github.com/seborama/fuego) - Functional Experiment in Go. Stars:`59`.

## Game Development

*Awesome game development libraries.*

* [Leaf](https://github.com/name5566/leaf) - Lightweight game server framework. Stars:`3.3K`.
* [Pixel](https://github.com/faiface/pixel) - Hand-crafted 2D game library in Go. Stars:`2.7K`.
* [Ebiten](https://github.com/hajimehoshi/ebiten) - dead simple 2D game library in Go. Stars:`2.4K`.
* [goworld](https://github.com/xiaonanln/goworld) - Scalable game server engine, featuring space-entity framework and hot-swapping. Stars:`1.4K`.
* [go-sdl2](https://github.com/veandco/go-sdl2) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/). Stars:`1.3K`.
* [nano](https://github.com/lonng/nano) - Lightweight, facility, high performance golang based game server framework. Stars:`1.2K`.
* [engo](https://github.com/EngoEngine/engo) - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm. Stars:`1.1K`.
* [gonet](https://github.com/xtaci/gonet) - Game server skeleton implemented with golang. Stars:`1.1K`.
* [termloop](https://github.com/JoelOtter/termloop) - Terminal-based game engine for Go, built on top of Termbox. Stars:`1.1K`.
* [g3n](https://github.com/g3n/engine) - Go 3D Game Engine. Stars:`904`.
* [Oak](https://github.com/oakmound/oak) - Pure Go game engine. Stars:`697`.
* [Azul3D](https://github.com/azul3d/engine) - 3D game engine written in Go. Stars:`441`.
* [Pitaya](https://github.com/topfreegames/pitaya) - Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK. Stars:`433`.
* [raylib-go](https://github.com/gen2brain/raylib-go) - Go bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to learn videogames programming. Stars:`423`.
* [go-astar](https://github.com/beefsack/go-astar) - Go implementation of the A\* path finding algorithm. Stars:`346`.
* [GarageEngine](https://github.com/vova616/GarageEngine) - 2d game engine written in Go working on OpenGL. Stars:`318`.
* [go3d](https://github.com/ungerik/go3d) - Performance oriented 2D/3D math package for Go. Stars:`172`.
* [glop](https://github.com/runningwild/glop) - Glop (Game Library Of Power) is a fairly simple cross-platform game library. Stars:`76`.
* [go-collada](https://github.com/GlenKelley/go-collada) - Go package for working with the Collada file format. Stars:`15`.

## Generation and Generics

*Tools to enhance the language with features like generics via code generation.*

* [go-linq](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go. Stars:`1.9K`.
* [jennifer](https://github.com/dave/jennifer) - Generate arbitrary Go code without templates. Stars:`1.4K`.
* [gen](https://github.com/clipperhouse/gen) - Code generation tool for ‘generics’-like functionality. Stars:`1.1K`.
* [goderive](https://github.com/awalterschulze/goderive) - Derives functions from input types. Stars:`779`.
* [GoWrap](https://github.com/hexdigest/gowrap) - Generate decorators for Go interfaces using simple templates. Stars:`307`.
* [interfaces](https://github.com/rjeczalik/interfaces) - Command line tool for generating interface definitions. Stars:`198`.
* [go-enum](https://github.com/abice/go-enum) - Code generation for enums from code comments. Stars:`94`.
* [pkgreflect](https://github.com/ungerik/pkgreflect) - Go preprocessor for package scoped reflection. Stars:`91`.
* [efaceconv](https://github.com/t0pep0/efaceconv) - Code generation tool for high performance conversion from interface{} to immutable type without allocations. Stars:`44`.
* [gotype](https://github.com/wzshiming/gotype) - Golang source code parsing, usage like reflect package. Stars:`28`.
* [generis](https://github.com/senselogic/GENERIS) - Code generation tool providing generics, free-form macros, conditional compilation and HTML templating. Stars:`20`.
* [go-xray](https://github.com/pieterclaerhout/go-xray) - Helpers for making the use of reflection easier. Stars:`5`.

## Geographic

*Geographic tools and servers*

* [Tile38](https://github.com/tidwall/tile38) - Geolocation DB with spatial index and realtime geofencing. Stars:`6.6K`.
* [S2 geometry](https://github.com/golang/geo) - S2 geometry library in Go. Stars:`983`.
* [geocache](https://github.com/melihmucuk/geocache) - In-memory cache that is suitable for geolocation based applications. Stars:`117`.
* [osm](https://github.com/paulmach/osm) - Library for reading, writing and working with OpenStreetMap data and APIs. Stars:`90`.
* [WGS84](https://github.com/wroge/wgs84) - Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM). Stars:`43`.
* [geoserver](https://github.com/hishamkaram/geoserver) - geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API. Stars:`27`.
* [gismanager](https://github.com/hishamkaram/gismanager) - Publish Your GIS Data(Vector Data) to PostGIS and Geoserver. Stars:`22`.
* [pbf](https://github.com/maguro/pbf) - OpenStreetMap PBF golang encoder/decoder. Stars:`19`.

## Go Compilers

*Tools for compiling Go to other languages.*

* [gopherjs](https://github.com/gopherjs/gopherjs) - Compiler from Go to JavaScript. Stars:`9.0K`.
* [llgo](https://github.com/go-llvm/llgo) - LLVM-based compiler for Go. Stars:`1.0K`.
* [tardisgo](https://github.com/tardisgo/tardisgo) - Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler. Stars:`394`.
* [c4go](https://github.com/Konstantin8105/c4go) - Transpile C code to Go code. Stars:`188`.
* [f4go](https://github.com/Konstantin8105/f4go) - Transpile FORTRAN 77 code to Go code. Stars:`23`.

## Goroutines

*Tools for managing and working with Goroutines.*

* [ants](https://github.com/panjf2000/ants) - A high-performance and low-cost goroutine pool in Go. Stars:`2.9K`.
* [goworker](https://github.com/benmanns/goworker) - goworker is a Go-based background worker. Stars:`2.3K`.
* [tunny](https://github.com/Jeffail/tunny) - Goroutine pool for golang. Stars:`1.5K`.
* [grpool](https://github.com/ivpusic/grpool) - Lightweight Goroutine pool. Stars:`535`.
* [pool](https://github.com/go-playground/pool) - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation. Stars:`534`.
* [workerpool](https://github.com/gammazero/workerpool) - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued. Stars:`234`.
* [go-floc](https://github.com/workanator/go-floc) - Orchestrate goroutines with ease. Stars:`175`.
* [go-flow](https://github.com/kamildrazkiewicz/go-flow) - Control goroutines execution order. Stars:`121`.
* [GoSlaves](https://github.com/themester/GoSlaves) - Simple and Asynchronous Goroutine pool library. Stars:`88`.
* [gowp](https://github.com/xxjwxc/gowp) - gowp is concurrency limiting goroutine pool. Stars:`88`.
* [semaphore](https://github.com/kamilsk/semaphore) - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context. Stars:`82`.
* [semaphore](https://github.com/marusama/semaphore) - Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations). Stars:`80`.
* [gpool](https://github.com/Sherifabdlnaby/gpool) - manages a resizeable pool of context-aware goroutines to bound concurrency. Stars:`61`.
* [worker-pool](https://github.com/vardius/worker-pool) - goworker is a Go simple async worker pool. Stars:`51`.
* [breaker](https://github.com/kamilsk/breaker) - Flexible mechanism to make execution flow interruptible. Stars:`48`.
* [cyclicbarrier](https://github.com/marusama/cyclicbarrier) - CyclicBarrier for golang. Stars:`43`.
* [gollback](https://github.com/vardius/gollback) - asynchronous simple function utilities, for managing execution of closures and callbacks. Stars:`30`.
* [async](https://github.com/studiosol/async) - A safe way to execute functions asynchronously, recovering them in case of panic. Stars:`29`.
* [parallel-fn](https://github.com/rafaeljesus/parallel-fn) - Run functions in parallel. Stars:`26`.
* [threadpool](https://github.com/shettyh/threadpool) - Golang threadpool implementation. Stars:`24`.
* [artifex](https://github.com/borderstech/artifex) - Simple in-memory job queue for Golang using worker-based dispatching. Stars:`21`.
* [Hunch](https://github.com/AaronJan/Hunch) - Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive. Stars:`19`.
* [stl](https://github.com/ssgreg/stl) - Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism. Stars:`11`.
* [gohive](https://github.com/loveleshsharma/gohive) - A highly performant and easy to use Goroutine pool for Go. Stars:`9`.
* [routine](https://github.com/x-mod/routine) - go routine control with context, support: Main, Go, Pool and some useful Executors. Stars:`5`.
* [go-trylock](https://github.com/subchen/go-trylock) - TryLock support on read-write lock for Golang. Stars:`5`.
* [go-tools/multithreading](https://github.com/nikhilsaraf/go-tools) - Manage a pool of goroutines using this lightweight library with a simple API. Stars:`5`.
* [queue](https://github.com/AnikHasibul/queue) - Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more. Stars:`3`.
* [go-waitgroup](https://github.com/pieterclaerhout/go-waitgroup) - Like `sync.WaitGroup` with error handling and concurrency control. Stars:`3`.
* [oversight](https://cirello.io/oversight) - Oversight is a complete implementation of the Erlang supervision trees.

## GUI

*Libraries for building GUI Applications.*

*Toolkits*

* [fyne](https://github.com/fyne-io/fyne) - Cross platform native GUIs designed for Go, rendered using EFL. Supports: Linux, macOS, Windows. Stars:`7.5K`.
* [ui](https://github.com/andlabs/ui) - Platform-native GUI library for Go. Cross platform. Stars:`7.3K`.
* [qt](https://github.com/therecipe/qt) - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi). Stars:`6.8K`.
* [webview](https://github.com/zserge/webview) - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux). Stars:`5.2K`.
* [walk](https://github.com/lxn/walk) - Windows application library kit for Go. Stars:`4.2K`.
* [app](https://github.com/murlokswarm/app) - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress. Stars:`3.2K`.
* [go-astilectron](https://github.com/asticode/go-astilectron) - Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron). Stars:`3.0K`.
* [go-sciter](https://github.com/sciter-sdk/go-sciter) - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform. Stars:`1.6K`.
* [gotk3](https://github.com/gotk3/gotk3) - Go bindings for GTK3. Stars:`883`.
* [gowd](https://github.com/dtylman/gowd) - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform. Stars:`235`.
* [go-gtk](http://mattn.github.io/go-gtk/) - Go bindings for GTK.
* [Wails](https://wails.app) - Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.

*Interaction*

* [robotgo](https://github.com/go-vgo/robotgo) - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other. Stars:`4.8K`.
* [systray](https://github.com/getlantern/systray) - Cross platform Go library to place an icon and menu in the notification area. Stars:`942`.
* [gosx-notifier](https://github.com/deckarep/gosx-notifier) - OSX Desktop Notifications library for Go. Stars:`503`.
* [trayhost](https://github.com/shurcooL/trayhost) - Cross-platform Go library to place an icon in the host operating system's taskbar. Stars:`168`.
* [go-appindicator](https://github.com/dawidd6/go-appindicator) - Go bindings for libappindicator3 C library. Stars:`4`.
* [mac-activity-tracker](https://github.com/prashantgupta24/activity-tracker) - OSX library to notify about any (pluggable) activity on your machine. Stars:`2`.
* [mac-sleep-notifier](https://github.com/prashantgupta24/mac-sleep-notifier) - OSX Sleep/Wake notifications in golang. Stars:`1`.


## Hardware

*Libraries, tools, and tutorials for interacting with hardware.*

See [go-hardware](https://github.com/rakyll/go-hardware) for a comprehensive list.

## Images

*Libraries for manipulating images.*

* [gocv](https://github.com/hybridgroup/gocv) - Go package for computer vision using OpenCV 3.3+. Stars:`2.9K`.
* [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package. Stars:`2.9K`.
* [imaginary](https://github.com/h2non/imaginary) - Fast and simple HTTP microservice for image resizing. Stars:`2.8K`.
* [bild](https://github.com/anthonynsimon/bild) - Collection of image processing algorithms in pure Go. Stars:`2.7K`.
* [ln](https://github.com/fogleman/ln) - 3D line art rendering in Go. Stars:`2.6K`.
* [resize](https://github.com/nfnt/resize) - Image resizing for Go with common interpolation methods. Stars:`2.2K`.
* [gg](https://github.com/fogleman/gg) - 2D rendering in pure Go. Stars:`2.1K`.
* [pt](https://github.com/fogleman/pt) - Path tracing engine written in Go. Stars:`1.8K`.
* [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation. Stars:`1.4K`.
* [smartcrop](https://github.com/muesli/smartcrop) - Finds good crops for arbitrary images and crop sizes. Stars:`1.3K`.
* [gift](https://github.com/disintegration/gift) - Package of image processing filters. Stars:`1.3K`.
* [picfit](https://github.com/thoas/picfit) - An image resizing server written in Go. Stars:`1.2K`.
* [go-opencv](https://github.com/lazywei/go-opencv) - Go bindings for OpenCV. Stars:`1.1K`.
* [imagick](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API. Stars:`1.1K`.
* [geopattern](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string. Stars:`1.0K`.
* [bimg](https://github.com/h2non/bimg) - Small package for fast and efficient image processing using libvips. Stars:`867`.
* [stegify](https://github.com/DimitarPetrov/stegify) - Go tool for LSB steganography, capable of hiding any file within an image. Stars:`590`.
* [canvas](https://github.com/tdewolff/canvas) - Vector graphics to PDF, SVG or rasterized image. Stars:`382`.
* [mort](https://github.com/aldor007/mort) - Storage and image processing server written in Go. Stars:`381`.
* [image2ascii](https://github.com/qeesung/image2ascii) - Convert image to ASCII. Stars:`349`.
* [govatar](https://github.com/o1egl/govatar) - Library and CMD tool for generating funny avatars. Stars:`332`.
* [go-nude](https://github.com/koyachi/go-nude) - Nudity detection with Go. Stars:`296`.
* [goimagehash](https://github.com/corona10/goimagehash) - Go Perceptual image hashing package. Stars:`260`.
* [rez](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD. Stars:`194`.
* [img](https://github.com/hawx/img) - Selection of image manipulation tools. Stars:`132`.
* [darkroom](https://github.com/gojek/darkroom) - An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency. Stars:`109`.
* [go-cairo](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library. Stars:`90`.
* [mergi](https://github.com/noelyahan/mergi) - Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate). Stars:`85`.
* [gltf](https://github.com/qmuntal/gltf) - Efficient and robust glTF 2.0 reader, writer and validator. Stars:`57`.
* [go-gd](https://github.com/bolknote/go-gd) - Go binding for GD library. Stars:`54`.
* [steganography](https://github.com/auyer/steganography) - Pure Go Library for LSB steganography. Stars:`36`.
* [cameron](https://github.com/aofei/cameron) - An avatar generator for Go. Stars:`36`.
* [goimghdr](https://github.com/corona10/goimghdr) - The imghdr module determines the type of image contained in a file for Go. Stars:`31`.
* [tga](https://github.com/ftrvxmtrx/tga) - Package tga is a TARGA image format decoder/encoder. Stars:`25`.
* [go-webcolors](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go. Stars:`24`.
* [mpo](https://github.com/donatj/mpo) - Decoder and conversion tool for MPO 3D Photos. Stars:`6`.

## IoT (Internet of Things)

*Libraries for programming devices of the IoT.*

* [flogo](https://github.com/tibcosoftware/flogo) - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration. Stars:`1.3K`.
* [gatt](https://github.com/paypal/gatt) - Gatt is a Go package for building Bluetooth Low Energy peripherals. Stars:`865`.
* [mainflux](https://github.com/Mainflux/mainflux) - Industrial IoT Messaging and Device Management Server. Stars:`751`.
* [devices](https://github.com/goiot/devices) - Suite of libraries for IoT devices, experimental for x/exp/io. Stars:`229`.
* [connectordb](https://github.com/connectordb/connectordb) - Open-Source Platform for Quantified Self & IoT. Stars:`190`.
* [sensorbee](https://github.com/sensorbee/sensorbee) - Lightweight stream processing engine for IoT. Stars:`187`.
* [huego](https://github.com/amimof/huego) - An extensive Philips Hue client library for Go. Stars:`127`.
* [eywa](https://github.com/xcodersun/eywa) - Project Eywa is essentially a connection manager that keeps track of connected devices. Stars:`43`.
* [gobot](https://github.com/hybridgroup/gobot/) - Gobot is a framework for robotics, physical computing, and the Internet of Things.
* [iot](https://github.com/vaelen/iot/) - IoT is a simple framework for implementing a Google IoT Core device.
* [periph](https://periph.io/) - Peripherals I/O to interface with low-level board facilities.

## Job Scheduler

*Libraries for scheduling jobs.*

* [gron](https://github.com/roylee0704/gron) - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly. Stars:`678`.
* [JobRunner](https://github.com/bamzi/jobrunner) - Smart and featureful cron job scheduler with job queuing and live monitoring built in. Stars:`638`.
* [jobs](https://github.com/albrow/jobs) - Persistent and flexible background jobs library. Stars:`461`.
* [scheduler](https://github.com/carlescere/scheduler) - Cronjobs scheduling made easy. Stars:`313`.
* [go-cron](https://github.com/rk/go-cron) - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons. Stars:`178`.
* [clockwork](https://github.com/whiteShtef/clockwork) - Simple and intuitive job scheduling library in Go. Stars:`95`.
* [leprechaun](https://github.com/kilgaloon/leprechaun) - Job scheduler that supports webhooks, crons and classic scheduling. Stars:`50`.
* [clockwerk](http://github.com/onatm/clockwerk) - Go package to schedule periodic jobs using a simple, fluent syntax.

## JSON

*Libraries for working with JSON.*

* [GJSON](https://github.com/tidwall/gjson) - Get a JSON value with one line of code. Stars:`5.7K`.
* [gojson](https://github.com/ChimeraCoder/gojson) - Automatically generate Go (golang) struct definitions from example JSON. Stars:`2.1K`.
* [kazaam](https://github.com/Qntfy/kazaam) - API for arbitrary transformation of JSON documents. Stars:`143`.
* [gojq](https://github.com/elgs/gojq) - JSON query in Golang. Stars:`142`.
* [jsongo](https://github.com/ricardolonga/jsongo) - Fluent API to make it easier to create Json objects. Stars:`94`.
* [gjo](https://github.com/skanehira/gjo) - Small utility to create JSON objects. Stars:`70`.
* [JayDiff](https://github.com/yazgazan/jaydiff) - JSON diff utility written in Go. Stars:`59`.
* [jettison](https://github.com/wI2L/jettison) - High performance, reflection-less JSON encoder for Go. Stars:`57`.
* [jsonf](https://github.com/miolini/jsonf) - Console tool for highlighted formatting and struct query fetching JSON. Stars:`54`.
* [mp](https://github.com/sanbornm/mp) - Simple cli email parser. It currently takes stdin and outputs JSON. Stars:`37`.
* [go-respond](https://github.com/nicklaw5/go-respond) - Go package for handling common HTTP JSON responses. Stars:`29`.
* [json2go](https://github.com/m-zajac/json2go) - Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all. Stars:`27`.
* [ajson](https://github.com/spyzhov/ajson) - Abstract JSON for golang with JSONPath support. Stars:`19`.
* [jsonhal](https://github.com/RichardKnop/jsonhal) - Simple Go package to make custom structs marshal into HAL compatible JSON responses. Stars:`9`.
* [go-jsonerror](https://github.com/ddymko/go-jsonerror) - Go-JsonError is ment to allow us to easily create json response errors that follow the JsonApi spec. Stars:`9`.
* [jsonapi-errors](https://github.com/AmuzaTkts/jsonapi-errors) - Go bindings based on the JSON API errors reference. Stars:`6`.
* [JSON-to-Go](https://mholt.github.io/json-to-go/) - Convert JSON to Go struct.

## Logging

*Libraries for generating and working with log files.*

* [logrus](https://github.com/Sirupsen/logrus) - Structured logger for Go. Stars:`13.6K`.
* [zap](https://github.com/uber-go/zap) - Fast, structured, leveled logging in Go. Stars:`8.8K`.
* [spew](https://github.com/davecgh/go-spew) - Implements a deep pretty printer for Go data structures to aid in debugging. Stars:`3.6K`.
* [zerolog](https://github.com/rs/zerolog) - Zero-allocation JSON logger. Stars:`2.8K`.
* [glog](https://github.com/golang/glog) - Leveled execution logs for Go. Stars:`2.5K`.
* [lumberjack](https://github.com/natefinch/lumberjack) - Simple rolling logger, implements io.WriteCloser. Stars:`1.7K`.
* [tail](https://github.com/hpcloud/tail) - Go package striving to emulate the features of the BSD tail program. Stars:`1.7K`.
* [seelog](https://github.com/cihub/seelog) - Logging functionality with flexible dispatching, filtering, and formatting. Stars:`1.4K`.
* [log15](https://github.com/inconshreveable/log15) - Simple, powerful logging for Go. Stars:`941`.
* [log](https://github.com/apex/log) - Structured logging package for Go. Stars:`775`.
* [onelog](https://github.com/francoispqt/onelog) - Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenario. Also, it is one of the logger with the lowest allocation. Stars:`358`.
* [logxi](https://github.com/mgutz/logxi) - 12-factor app logger that is fast and makes you happy. Stars:`340`.
* [log](https://github.com/go-playground/log) - Simple, configurable and scalable Structured Logging for Go. Stars:`270`.
* [logutils](https://github.com/hashicorp/logutils) - Utilities for slightly better logging in Go (Golang) extending the standard logger. Stars:`264`.
* [go-logger](https://github.com/apsdehal/go-logger) - Simple logger of Go Programs, with level handlers. Stars:`246`.
* [logger](https://github.com/azer/logger) - Minimalistic logging library for Go. Stars:`136`.
* [xlog](https://github.com/rs/xlog) - Structured logger for `net/context` aware HTTP handlers with flexible dispatching. Stars:`131`.
* [rollingwriter](https://github.com/arthurkiller/rollingWriter) - RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation. Stars:`128`.
* [ozzo-log](https://github.com/go-ozzo/ozzo-log) - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail). Stars:`110`.
* [log-voyage](https://github.com/firstrow/logvoyage) - Full-featured logging saas written in golang. Stars:`84`.
* [glg](https://github.com/kpango/glg) - glg is simple and fast leveled logging library for Go. Stars:`60`.
* [stdlog](https://github.com/alexcesaro/log) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs. Stars:`42`.
* [gologger](https://github.com/sadlil/gologger) - Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch. Stars:`38`.
* [go-log](https://github.com/ian-kent/go-log) - Log4j implementation in Go. Stars:`34`.
* [logex](https://github.com/chzyer/logex) - Golang log lib, supports tracking and level, wrap by standard log lib. Stars:`33`.
* [go-log](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers. Stars:`26`.
* [logrusly](https://github.com/sebest/logrusly) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/). Stars:`25`.
* [go-cronowriter](https://github.com/utahta/go-cronowriter) - Simple writer that rotate log files automatically based on current date and time, like cronolog. Stars:`24`.
* [log](https://github.com/teris-io/log) - Structured log interface for Go cleanly separates logging facade from its implementation. Stars:`23`.
* [journald](https://github.com/ssgreg/journald) - Go implementation of systemd Journal's native API for logging. Stars:`22`.
* [distillog](https://github.com/amoghe/distillog) - distilled levelled logging (think of it as stdlib + log levels). Stars:`21`.
* [mlog](https://github.com/jbrodriguez/mlog) - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output. Stars:`19`.
* [gomol](https://github.com/aphistic/gomol) - Multiple-output, structured logging for Go with extensible logging outputs. Stars:`15`.
* [go-log](https://github.com/subchen/go-log) - Simple and configurable Logging in Go, with level, formatters and writers. Stars:`9`.
* [logdump](https://github.com/ewwwwwqm/logdump) - Package for multi-level logging. Stars:`9`.
* [glo](https://github.com/lajosbencz/glo) - PHP Monolog inspired logging facility with identical severity levels. Stars:`8`.
* [logrusiowriter](https://github.com/cabify/logrusiowriter) - `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger. Stars:`7`.
* [log](https://github.com/aerogo/log) - An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection). Stars:`6`.
* [logmatic](https://github.com/borderstech/logmatic) - Colorized logger for Golang with dynamic log level configuration. Stars:`6`.
* [xlog](https://github.com/xfxdev/xlog) - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format. Stars:`6`.
* [logo](https://github.com/mbndr/logo) - Golang logger to different configurable writers. Stars:`5`.
* [go-log](https://github.com/pieterclaerhout/go-log) - A logging library with strack traces, object dumping and optional timestamps. Stars:`1`.
* [gone/log](https://github.com/One-com/gone/tree/master/log) - Fast, extendable, full-featured, std-lib source compatible log library.

## Machine Learning

*Libraries for Machine Learning.*

* [GoLearn](https://github.com/sjwhitworth/golearn) - General Machine Learning library for Go. Stars:`6.9K`.
* [gorgonia](https://github.com/gorgonia/gorgonia) - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms. Stars:`3.1K`.
* [tfgo](https://github.com/galeone/tfgo) - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python. Stars:`1.3K`.
* [goml](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go. Stars:`1.1K`.
* [gosseract](https://github.com/otiai10/gosseract) - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library. Stars:`1.0K`.
* [gorse](https://github.com/zhenghaoz/gorse) - An offline recommender system backend based on collaborative filtering written in Go. Stars:`876`.
* [CloudForest](https://github.com/ryanbressler/CloudForest) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go. Stars:`660`.
* [eaopt](https://github.com/MaxHalford/eaopt) - An evolutionary optimization library. Stars:`656`.
* [bayesian](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang. Stars:`648`.
* [gobrain](https://github.com/goml/gobrain) - Neural Networks written in go. Stars:`418`.
* [regommend](https://github.com/muesli/regommend) - Recommendation & collaborative filtering engine. Stars:`263`.
* [ocrserver](https://github.com/otiai10/ocrserver) - A simple OCR API server, seriously easy to be deployed by Docker and Heroku. Stars:`261`.
* [go-deep](https://github.com/patrikeh/go-deep) - A feature-rich neural network library in Go. Stars:`246`.
* [onnx-go](https://github.com/owulveryck/onnx-go) - Go Interface to Open Neural Network Exchange (ONNX). Stars:`233`.
* [go-galib](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang. Stars:`176`.
* [goRecommend](https://github.com/timkaye11/goRecommend) - Recommendation Algorithms library written in Go. Stars:`151`.
* [shield](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go. Stars:`129`.
* [go-fann](https://github.com/white-pony/go-fann) - Go bindings for Fast Artificial Neural Networks(FANN) library. Stars:`103`.
* [Goptuna](https://github.com/c-bata/goptuna) - Bayesian optimization framework for black-box functions written in Go. Everything will be optimized. Stars:`101`.
* [goga](https://github.com/tomcraven/goga) - Genetic algorithm library for Go. Stars:`81`.
* [libsvm](https://github.com/datastream/libsvm) - libsvm golang version derived work based on LIBSVM 3.14. Stars:`64`.
* [neural-go](https://github.com/schuyler/neural-go) - Multilayer perceptron network implemented in Go, with training via backpropagation. Stars:`61`.
* [go-pr](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang. Stars:`57`.
* [neat](https://github.com/jinyeom/neat) - Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT). Stars:`56`.
* [goscore](https://github.com/asafschers/goscore) - Go Scoring API for PMML. Stars:`43`.
* [golinear](https://github.com/danieldk/golinear) - liblinear bindings for Go. Stars:`39`.
* [fonet](https://github.com/Fontinalis/fonet) - A Deep Neural Network library written in Go. Stars:`35`.
* [Varis](https://github.com/Xamber/Varis) - Golang Neural Network. Stars:`26`.
* [godist](https://github.com/e-dard/godist) - Various probability distributions, and associated methods. Stars:`25`.
* [go-cluster](https://github.com/e-XpertSolutions/go-cluster) - Go implementation of the k-modes and k-prototypes clustering algorithms. Stars:`21`.
* [probab](https://github.com/ThePaw/probab) - Probability distribution functions. Bayesian inference. Written in pure Go. Stars:`11`.
* [evoli](https://github.com/khezen/evoli) - Genetic Algorithm and Particle Swarm Optimization library. Stars:`9`.
* [GoMind](https://github.com/surenderthakran/gomind) - A simplistic Neural Network Library in Go. Stars:`7`.
* [randomforest](https://github.com/malaschitz/randomForest) - Easy to use Random Forest library for Go. Stars:`2`.

## Messaging

*Libraries that implement messaging systems.*

* [sarama](https://github.com/Shopify/sarama) - Go library for Apache Kafka. Stars:`5.3K`.
* [gorush](https://github.com/appleboy/gorush) - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm). Stars:`4.1K`.
* [Centrifugo](https://github.com/centrifugal/centrifugo) - Real-time messaging (Websockets or SockJS) server in Go. Stars:`4.0K`.
* [machinery](https://github.com/RichardKnop/machinery) - Asynchronous task queue/job queue based on distributed message passing. Stars:`3.8K`.
* [go-socket.io](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework. Stars:`3.2K`.
* [NATS Go Client](https://github.com/nats-io/nats) - Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library. Stars:`2.6K`.
* [APNs2](https://github.com/sideshow/apns2) - HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps. Stars:`2.2K`.
* [Benthos](https://github.com/Jeffail/benthos) - A message streaming bridge between a range of protocols. Stars:`2.2K`.
* [gopush-cluster](https://github.com/Terry-Mao/gopush-cluster) - gopush-cluster is a go push server cluster. Stars:`1.9K`.
* [Mercure](https://github.com/dunglas/mercure) - Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events). Stars:`1.8K`.
* [melody](https://github.com/olahol/melody) - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling. Stars:`1.7K`.
* [go-nsq](https://github.com/nsqio/go-nsq) - the official Go package for NSQ. Stars:`1.6K`.
* [mangos](https://github.com/go-mangos/mangos) - Pure go implementation of the Nanomsg ("Scalable Protocols") with transport interoperability. Stars:`1.5K`.
* [Uniqush-Push](https://github.com/uniqush/uniqush-push) - Redis backed unified push service for server-side notifications to mobile devices. Stars:`1.1K`.
* [zmq4](https://github.com/pebbe/zmq4) - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2). Stars:`820`.
* [Gollum](https://github.com/trivago/gollum) - A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations. Stars:`799`.
* [Beaver](https://github.com/Clivern/Beaver) - A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps. Stars:`777`.
* [EventBus](https://github.com/asaskevich/EventBus) - The lightweight event bus with async compatibility. Stars:`630`.
* [golongpoll](https://github.com/jcuga/golongpoll) - HTTP longpoll server library that makes web pub-sub simple. Stars:`439`.
* [dbus](https://github.com/godbus/dbus) - Native Go bindings for D-Bus. Stars:`404`.
* [emitter](https://github.com/olebedev/emitter) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins. Stars:`331`.
* [Glue](https://github.com/desertbit/glue) - Robust Go and Javascript Socket Library (Alternative to Socket.io). Stars:`324`.
* [pubsub](https://github.com/tuxychandru/pubsub) - Simple pubsub package for go. Stars:`305`.
* [guble](https://github.com/smancke/guble) - Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence. Stars:`142`.
* [Bus](https://github.com/mustafaturan/bus) - Minimalist message bus implementation for internal communication. Stars:`131`.
* [oplog](https://github.com/dailymotion/oplog) - Generic oplog/replication system for REST APIs. Stars:`100`.
* [rabtap](https://github.com/jandelgado/rabtap) - RabbitMQ swiss army knife cli app. Stars:`97`.
* [messagebus](https://github.com/vardius/message-bus) - messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD. Stars:`78`.
* [rabbus](https://github.com/rafaeljesus/rabbus) - A tiny wrapper over amqp exchanges and queues. Stars:`68`.
* [drone-line](https://github.com/appleboy/drone-line) - Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI. Stars:`65`.
* [nsq-event-bus](https://github.com/rafaeljesus/nsq-event-bus) - A tiny wrapper around NSQ topic and channel. Stars:`57`.
* [RapidMQ](https://github.com/sybrexsys/RapidMQ) - RapidMQ is a lightweight and reliable library for managing of the local messages queue. Stars:`57`.
* [go-notify](https://github.com/TheCreeper/go-notify) - Native implementation of the freedesktop notification spec. Stars:`47`.
* [Commander](https://github.com/jeroenrinzema/commander) - A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka. Stars:`37`.
* [hub](https://github.com/leandro-lugaresi/hub) - A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges. Stars:`32`.
* [event](https://github.com/agoalofalife/event) - Implementation of the pattern observer. Stars:`32`.
* [go-vitotrol](https://github.com/maxatome/go-vitotrol) - Client library to Viessmann Vitotrol web service. Stars:`12`.
* [redisqueue](https://github.com/robinjoseph08/redisqueue) - redisqueue provides a producer and consumer of a queue that uses Redis streams. Stars:`8`.
* [jazz](https://github.com/socifi/jazz) - A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages. Stars:`8`.
* [gaurun-client](https://github.com/osamingo/gaurun-client) - Gaurun Client written in Go. Stars:`8`.
* [rmqconn](https://github.com/sbabiv/rmqconn) - RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed. Stars:`2`.

## Microsoft Office

* [unioffice](https://github.com/unidoc/unioffice) - Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents. Stars:`2.0K`.

### Microsoft Excel

*Libraries for working with Microsoft Excel.*

* [excelize](https://github.com/360EntSecGroup-Skylar/excelize) - Golang library for reading and writing Microsoft Excel™ (XLSX) files. Stars:`5.3K`.
* [xlsx](https://github.com/tealeg/xlsx) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs. Stars:`3.8K`.
* [xlsx](https://github.com/plandem/xlsx) - Fast and safe way to read/update your existing Microsoft Excel files in Go programs. Stars:`101`.
* [go-excel](https://github.com/szyhf/go-excel) - A simple and light reader to read a relate-db-like excel as a table. Stars:`59`.
* [goxlsxwriter](https://github.com/fterrag/goxlsxwriter) - Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files. Stars:`14`.

## Miscellaneous

### Dependency Injection

*Libraries for working with dependency injection.*

* [dig](https://github.com/uber-go/dig) - A reflection based dependency injection toolkit for Go. Stars:`1.2K`.
* [fx](https://github.com/uber-go/fx) - A dependency injection based application framework for Go (built on top of dig). Stars:`1.1K`.
* [container](https://github.com/golobby/container) - A powerful IoC Container with fluent and easy-to-use interface. Stars:`62`.
* [inject](https://github.com/defval/inject) - A reflection based dependency injection container with simple interface. Stars:`48`.
* [dingo](https://github.com/i-love-flamingo/dingo) - A dependency injection toolkit for Go, based on Guice. Stars:`36`.
* [alice](https://github.com/magic003/alice) - Additive dependency injection container for Golang. Stars:`36`.
* [wire](https://github.com/Fs02/wire) - Strict Runtime Dependency Injection for Golang. Stars:`22`.
* [linker](https://github.com/logrange/linker) - A reflection based dependency injection and inversion of control library with components lifecycle support. Stars:`18`.
* [gocontainer](https://github.com/vardius/gocontainer) - Simple Dependency Injection Container. Stars:`11`.

### Project Layout

*Unofficial set of patterns for structuring projects.*

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout) - Set of common historical and emerging project layout patterns in the Go ecosystem. Stars:`12.5K`.
* [go-restful-api](https://github.com/qiangxue/go-restful-api) - An idiomatic Go RESTful API starter kit following SOLID principles and Clean Architecture with a common project layout. Stars:`990`.
* [modern-go-application](https://github.com/sagikazarmark/modern-go-application) - Go application boilerplate and example applying modern practices. Stars:`472`.
* [cookiecutter-golang](https://github.com/lacion/cookiecutter-golang) - A Go application boilerplate template for quick starting projects following production best practices. Stars:`309`.
* [scaffold](https://github.com/catchplay/scaffold) - Scaffold generates starter Go project layout. Lets you focus on business logic implemeted. Stars:`49`.
* [go-sample](https://github.com/zitryss/go-sample) - A sample layout for Go application projects with the real code. Stars:`41`.

### Strings

*Libraries for working with strings.*

* [xstrings](https://github.com/huandu/xstrings) - Collection of useful string functions ported from other languages. Stars:`693`.
* [strutil](https://github.com/ozgio/strutil) - String utilities. Stars:`82`.

*These libraries were placed here because none of the other categories seemed to fit.*

* [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc). Stars:`4.4K`.
* [archiver](https://github.com/mholt/archiver) - Library and command for making and extracting .zip and .tar.gz archives. Stars:`2.7K`.
* [gosms](https://github.com/haxpax/gosms) - Your own local SMS gateway in Go that can be used to send SMS. Stars:`1.3K`.
* [go-resiliency](https://github.com/eapache/go-resiliency) - Resiliency patterns for golang. Stars:`947`.
* [gofakeit](https://github.com/brianvoe/gofakeit) - Random data generator written in go. Stars:`791`.
* [go-commons-pool](https://github.com/jolestar/go-commons-pool) - Generic object pool for Golang. Stars:`754`.
* [base64Captcha](https://github.com/mojocn/base64Captcha) - Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha. Stars:`739`.
* [shortid](https://github.com/teris-io/shortid) - Distributed generation of super short, unique, non-sequential, URL friendly IDs. Stars:`507`.
* [llvm](https://github.com/llir/llvm) - Library for interacting with LLVM IR in pure Go. Stars:`472`.
* [health](https://github.com/dimiro1/health) - Easy to use, extensible health check library. Stars:`374`.
* [conv](https://github.com/cstockton/go-conv) - Package conv provides fast and intuitive conversions across Go types. Stars:`348`.
* [banner](https://github.com/dimiro1/banner) - Add beautiful banners into your Go applications. Stars:`245`.
* [gountries](https://github.com/pariz/gountries) - Package that exposes country and subdivision data. Stars:`228`.
* [antch](https://github.com/antchfx/antch) - A fast, powerful and extensible web crawling & scraping framework. Stars:`165`.
* [ffmt](https://github.com/go-ffmt/ffmt) - Beautify data display for Humans. Stars:`145`.
* [lk](https://github.com/hyperboloide/lk) - A simple licensing library for golang. Stars:`142`.
* [stateless](https://github.com/qmuntal/stateless) - A fluent library for creating state machines. Stars:`140`.
* [battery](https://github.com/distatus/battery) - Cross-platform, normalized battery information library. Stars:`138`.
* [stats](https://github.com/go-playground/stats) - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc... Stars:`129`.
* [bitio](https://github.com/icza/bitio) - Highly optimized bit-level Reader and Writer for Go. Stars:`110`.
* [healthcheck](https://github.com/etherlabsio/healthcheck) - An opinionated and concurrent health-check HTTP handler for RESTful services. Stars:`94`.
* [turtle](https://github.com/hackebrot/turtle) - Emojis for Go. Stars:`91`.
* [ghorg](https://github.com/gabrie30/ghorg) - Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, and Bitbucket. Stars:`90`.
* [gommit](https://github.com/antham/gommit) - Analyze git commit messages to ensure they follow defined patterns. Stars:`83`.
* [go-unarr](https://github.com/gen2brain/go-unarr) - Decompression library for RAR, TAR, ZIP and 7z archives. Stars:`80`.
* [indigo](https://github.com/osamingo/indigo) - Distributed unique ID generator of using Sonyflake and encoded by Base58. Stars:`58`.
* [morse](https://github.com/alwindoss/morse) - Library to convert to and from morse code. Stars:`54`.
* [captcha](https://github.com/steambap/captcha) - Package captcha provides an easy to use, unopinionated API for captcha generation. Stars:`48`.
* [xkg](https://github.com/go-xkg/xkg) - X Keyboard Grabber. Stars:`45`.
* [persian](https://github.com/mavihq/persian) - Some utilities for Persian language in go. Stars:`37`.
* [pdfgen](https://github.com/hyperboloide/pdfgen) - HTTP service to generate PDF from Json requests. Stars:`36`.
* [browscap_go](https://github.com/digitalcrab/browscap_go) - GoLang Library for [Browser Capabilities Project](http://browscap.org/). Stars:`31`.
* [datacounter](https://github.com/miolini/datacounter) - Go counters for readers/writer/http.ResponseWriter. Stars:`30`.
* [gotoprom](https://github.com/cabify/gotoprom) - Type-safe metrics builder wrapper library for the official Prometheus client. Stars:`30`.
* [autoflags](https://github.com/artyom/autoflags) - Go package to automatically define command line flags from struct fields. Stars:`25`.
* [xdg](https://github.com/rkoesters/xdg) - FreeDesktop.org (xdg) Specs implemented in Go. Stars:`21`.
* [gosh](https://github.com/osamingo/gosh) - Provide Go Statistics Handler, Struct, Measure Method. Stars:`19`.
* [url-shortener](https://github.com/pantrif/url-shortener) - A modern, powerful, and robust URL shortener microservice with mysql support. Stars:`18`.
* [sandid](https://github.com/aofei/sandid) - Every grain of sand on earth has its own ID. Stars:`14`.
* [anagent](https://github.com/mudler/anagent) - Minimalistic, pluggable Golang evloop/timer handler with dependency-injection. Stars:`11`.
* [avgRating](https://github.com/kirillDanshin/avgRating) - Calculate average score and rating based on Wilson Score Equation. Stars:`10`.
* [hostutils](https://github.com/Wing924/hostutils) - A golang library for packing and unpacking FQDNs list. Stars:`8`.
* [shellwords](https://github.com/Wing924/shellwords) - A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell. Stars:`7`.
* [metrics](https://github.com/pascaldekloe/metrics) - Library for metrics instrumentation and Prometheus exposition. Stars:`6`.
* [numa](https://github.com/lrita/numa) - NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code. Stars:`2`.
* [VarHandler](https://github.com/azr/generators/tree/master/varhandler) - Generate boilerplate http input and output handling.
* [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.

## Natural Language Processing

*Libraries for working with human languages.*

* [prose](https://github.com/jdkato/prose) - Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only. Stars:`2.4K`.
* [gse](https://github.com/go-ego/gse) - Go efficient text segmentation; support english, chinese, japanese and other. Stars:`1.2K`.
* [when](https://github.com/olebedev/when) - Natural EN and RU language date/time parser with pluggable rules. Stars:`1.0K`.
* [gojieba](https://github.com/yanyiwu/gojieba) - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm. Stars:`946`.
* [go-pinyin](https://github.com/mozillazg/go-pinyin) - CN Hanzi to Hanyu Pinyin converter. Stars:`632`.
* [kagome](https://github.com/ikawaha/kagome) - JP morphological analyzer written in pure Go. Stars:`470`.
* [whatlanggo](https://github.com/abadojack/whatlanggo) - Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc). Stars:`388`.
* [nlp](https://github.com/Shixzie/nlp) - Extract values from strings and fill your structs with nlp. Stars:`358`.
* [sentences](https://github.com/neurosnap/sentences) - Sentence tokenizer:  converts text into a list of sentences. Stars:`264`.
* [nlp](https://github.com/james-bowman/nlp) - Go Natural Language Processing library supporting LSA (Latent Semantic Analysis). Stars:`236`.
* [go-nlp](https://github.com/nuance/go-nlp) - Utilities for working with discrete probability distributions and other tools useful for doing NLP work. Stars:`82`.
* [getlang](https://github.com/rylans/getlang) - Fast natural language detection package. Stars:`78`.
* [gounidecode](https://github.com/fiam/gounidecode) - Unicode transliterator (also known as unidecode) for Go. Stars:`68`.
* [go-unidecode](https://github.com/mozillazg/go-unidecode) - ASCII transliterations of Unicode text. Stars:`64`.
* [textcat](https://github.com/pebbe/textcat) - Go package for n-gram based text categorization, with support for utf-8 and raw text. Stars:`62`.
* [MMSEGO](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm. Stars:`58`.
* [go-stem](https://github.com/agonopol/go-stem) - Implementation of the porter stemming algorithm. Stars:`54`.
* [RAKE.go](https://github.com/Obaied/RAKE.go) - Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE). Stars:`52`.
* [stemmer](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers. Stars:`48`.
* [segment](https://github.com/blevesearch/segment) - Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](http://www.unicode.org/reports/tr29/) Stars:`46`.
* [porter2](https://github.com/zhenjl/porter2) - Really fast Porter 2 stemmer. Stars:`35`.
* [go2vec](https://github.com/danieldk/go2vec) - Reader and utility functions for word2vec embeddings. Stars:`32`.
* [petrovich](https://github.com/striker2000/petrovich) - Petrovich is the library which inflects Russian names to given grammatical case. Stars:`27`.
* [paicehusk](https://github.com/rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm. Stars:`26`.
* [snowball](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/). Stars:`24`.
* [go-mystem](https://github.com/dveselov/mystem) - CGo bindings to Yandex.Mystem - russian morphology analyzer. Stars:`23`.
* [icu](https://github.com/goodsign/icu) - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1. Stars:`19`.
* [golibstemmer](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2. Stars:`16`.
* [shamoji](https://github.com/osamingo/shamoji) - The shamoji is word filtering package written in Go. Stars:`11`.
* [libtextcat](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2. Stars:`10`.
* [porter](https://github.com/a2800276/porter) - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm. Stars:`8`.
* [gotokenizer](https://github.com/xujiajun/gotokenizer) - A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation) Stars:`6`.
* [go-localize](https://github.com/m1/go-localize) - Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings. Stars:`2`.
* [detectlanguage](https://github.com/detectlanguage/detectlanguage-go) - Language Detection API Go Client. Supports batch requests, short phrase or single word language detection. Stars:`1`.
* [go-i18n](https://github.com/nicksnyder/go-i18n/) - Package and an accompanying tool to work with localized text.

## Networking

*Libraries for working with various layers of the network.*

* [kcptun](https://github.com/xtaci/kcptun) - Extremely simple & fast udp tunnel based on KCP protocol. Stars:`11.4K`.
* [fasthttp](https://github.com/valyala/fasthttp) - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http. Stars:`11.2K`.
* [dns](https://github.com/miekg/dns) - Go library for working with DNS. Stars:`4.3K`.
* [quic-go](https://github.com/lucas-clemente/quic-go) - An implementation of the QUIC protocol in pure Go. Stars:`3.7K`.
* [HTTPLab](https://github.com/gchaincl/httplab) - HTTPLabs let you inspect HTTP requests and forge responses. Stars:`3.5K`.
* [gopacket](https://github.com/google/gopacket) - Go library for packet processing with libpcap bindings. Stars:`3.2K`.
* [webrtc](https://github.com/pions/webrtc) - A pure Go implementation of the WebRTC API. Stars:`2.9K`.
* [kcp-go](https://github.com/xtaci/kcp-go) - KCP - Fast and Reliable ARQ Protocol. Stars:`2.4K`.
* [gobgp](https://github.com/osrg/gobgp) - BGP implemented in the Go Programming Language. Stars:`1.8K`.
* [gnet](https://github.com/panjf2000/gnet) - `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go. Stars:`1.6K`.
* [ssh](https://github.com/gliderlabs/ssh) - Higher-level API for building SSH servers (wraps crypto/ssh). Stars:`1.3K`.
* [fortio](https://github.com/fortio/fortio) - Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC. Stars:`1.1K`.
* [water](https://github.com/songgao/water) - Simple TUN/TAP library. Stars:`942`.
* [go-getter](https://github.com/hashicorp/go-getter) - Go library for downloading files or directories from various sources using a URL. Stars:`832`.
* [sftp](https://github.com/pkg/sftp) - Package sftp implements the SSH File Transfer Protocol as described in https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt. Stars:`819`.
* [gev](https://github.com/Allenxuxu/gev) - gev is a lightweight, fast non-blocking TCP network library based on Reactor mode. Stars:`763`.
* [NFF-Go](https://github.com/intel-go/nff-go) - Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF). Stars:`753`.
* [grab](https://github.com/cavaliercoder/grab) - Go package for managing file downloads. Stars:`623`.
* [mdns](https://github.com/hashicorp/mdns) - Simple mDNS (Multicast DNS) client/server library in Golang. Stars:`606`.
* [ftp](https://github.com/jlaffaye/ftp) - Package ftp implements a FTP client as described in [RFC 959](http://tools.ietf.org/html/rfc959). Stars:`601`.
* [lhttp](https://github.com/fanux/lhttp) - Powerful websocket framework, build your IM server more easily. Stars:`544`.
* [gosnmp](https://github.com/soniah/gosnmp) - Native Go library for performing SNMP actions. Stars:`483`.
* [gotcp](https://github.com/gansidui/gotcp) - Go package for quickly writing tcp applications. Stars:`447`.
* [cidranger](https://github.com/yl2chen/cidranger) - Fast IP to CIDR lookup for Go. Stars:`431`.
* [peerdiscovery](https://github.com/schollz/peerdiscovery) - Pure Go library for cross-platform local peer discovery using UDP multicast. Stars:`394`.
* [gopcap](https://github.com/akrennmair/gopcap) - Go wrapper for libpcap. Stars:`373`.
* [go-stun](https://github.com/ccding/go-stun) - Go implementation of the STUN client (RFC 3489 and RFC 5389). Stars:`339`.
* [raw](https://github.com/mdlayher/raw) - Package raw enables reading and writing data at the device driver level for a network interface. Stars:`336`.
* [stun](https://github.com/go-rtc/stun) - Go implementation of RFC 5389 STUN protocol. Stars:`327`.
* [tcp_server](https://github.com/firstrow/tcp_server) - Go library for building tcp servers faster. Stars:`308`.
* [buffstreams](https://github.com/stabbycutyou/buffstreams) - Streaming protocolbuffer data over TCP made easy. Stars:`235`.
* [winrm](https://github.com/masterzen/winrm) - Go WinRM client to remotely execute commands on Windows machines. Stars:`233`.
* [arp](https://github.com/mdlayher/arp) - Package arp implements the ARP protocol, as described in RFC 826. Stars:`209`.
* [ethernet](https://github.com/mdlayher/ethernet) - Package ethernet implements marshaling and unmarshaling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags. Stars:`192`.
* [utp](https://github.com/anacrolix/utp) - Go uTP micro transport protocol implementation. Stars:`153`.
* [jazigo](https://github.com/udhos/jazigo) - Jazigo is a tool written in Go for retrieving configuration for multiple network devices. Stars:`139`.
* [canopus](https://github.com/zubairhamed/canopus) - CoAP Client/Server implementation (RFC 7252). Stars:`137`.
* [sslb](https://github.com/eduardonunesp/sslb) - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance. Stars:`122`.
* [gmqtt](https://github.com/DrmagicE/gmqtt) - Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1. Stars:`119`.
* [gNxI](https://github.com/google/gnxi) - A collection of tools for Network Management that use the gNMI and gNOI protocols. Stars:`114`.
* [xtcp](https://github.com/xfxdev/xtcp) - TCP Server Framework with simultaneous full duplex communication,graceful shutdown,custom protocol. Stars:`92`.
* [dhcp6](https://github.com/mdlayher/dhcp6) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315. Stars:`64`.
* [ether](https://github.com/songgao/ether) - Cross-platform Go package for sending and receiving ethernet frames. Stars:`64`.
* [linkio](https://github.com/ian-kent/linkio) - Network link speed simulation for Reader/Writer interfaces. Stars:`47`.
* [portproxy](https://github.com/aybabtme/portproxy) - Simple TCP proxy which adds CORS support to API's which don't support it. Stars:`42`.
* [packet](https://github.com/aerogo/packet) - Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed. Stars:`42`.
* [iplib](https://github.com/c-robinson/iplib) - Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html) Stars:`27`.
* [graval](https://github.com/koofr/graval) - Experimental FTP server framework. Stars:`25`.
* [publicip](https://github.com/polera/publicip) - Package publicip returns your public facing IPv4 address (internet egress). Stars:`18`.
* [golibwireshark](https://github.com/sunwxg/golibwireshark) - Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data. Stars:`17`.
* [llb](https://github.com/kirillDanshin/llb) - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response. Stars:`10`.
* [go-powerdns](https://github.com/joeig/go-powerdns) - PowerDNS API bindings for Golang. Stars:`10`.
* [goshark](https://github.com/sunwxg/goshark) - Package goshark use tshark to decode IP packet and create data struct to analyse packet. Stars:`9`.
* [tspool](https://github.com/two/tspool) - A TCP Library use worker pool to improve performance and protect your server. Stars:`6`.
* [gosocsvr](https://github.com/rakeki/gosocsvr) - Socket server made simple. Stars:`4`.
* [httpproxy](https://github.com/wzshiming/httpproxy) - HTTP proxy handler and dialer. Stars:`1`.
* [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.

### HTTP Clients

*Libraries for making HTTP requests.*

* [resty](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client. Stars:`2.5K`.
* [grequests](https://github.com/levigross/grequests) - A Go "clone" of the great and famous Requests library. Stars:`1.5K`.
* [heimdall](https://github.com/gojektech/heimdall) - An enchanced http client with retry and hystrix capabilities. Stars:`1.2K`.
* [sling](https://github.com/dghubble/sling) - Sling is a Go HTTP client library for creating and sending API requests. Stars:`1.1K`.
* [gentleman](https://github.com/h2non/gentleman) - Full-featured plugin-driven HTTP client library. Stars:`717`.
* [pester](https://github.com/sethgrid/pester) - Go HTTP client calls with retries, backoff, and concurrency. Stars:`348`.
* [sreq](https://github.com/winterssy/sreq) - A simple, user-friendly and concurrent safe HTTP request library for Go. Stars:`45`.
* [rq](https://github.com/ddo/rq) - A nicer interface for golang stdlib HTTP client. Stars:`33`.

## OpenGL

*Libraries for using OpenGL in Go.*

* [glfw](https://github.com/go-gl/glfw) - Go bindings for GLFW 3. Stars:`832`.
* [gl](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow). Stars:`678`.
* [mathgl](https://github.com/go-gl/mathgl) - Pure Go math package specialized for 3D math, with inspiration from GLM. Stars:`313`.
* [goxjs/gl](https://github.com/goxjs/gl) - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android). Stars:`130`.
* [goxjs/glfw](https://github.com/goxjs/glfw) - Go cross-platform glfw library for creating an OpenGL context and receiving events. Stars:`59`.

## ORM

*Libraries that implement Object-Relational Mapping or datamapping techniques.*

* [GORM](https://github.com/jinzhu/gorm) - The fantastic ORM library for Golang, aims to be developer friendly. Stars:`16.8K`.
* [Xorm](https://github.com/go-xorm/xorm) - Simple and powerful ORM for Go. Stars:`5.7K`.
* [go-pg](https://github.com/go-pg/pg) - PostgreSQL ORM with focus on PostgreSQL specific features and performance. Stars:`3.4K`.
* [gorp](https://github.com/go-gorp/gorp) - Go Relational Persistence, ORM-ish library for Go. Stars:`3.3K`.
* [SQLBoiler](https://github.com/volatiletech/sqlboiler) - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema. Stars:`2.6K`.
* [upper.io/db](https://github.com/upper/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers. Stars:`2.0K`.
* [reform](https://github.com/go-reform/reform) - Better ORM for Go, based on non-empty interfaces and code generation. Stars:`836`.
* [pop/soda](https://github.com/gobuffalo/pop) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite. Stars:`786`.
* [QBS](https://github.com/coocood/qbs) - Stands for Query By Struct. A Go ORM. Stars:`543`.
* [go-queryset](https://github.com/jirfag/go-queryset) - 100% type-safe ORM with code generation and MySQL, PostgreSQL, Sqlite3, SQL Server support based on GORM. Stars:`481`.
* [go-sqlbuilder](https://github.com/huandu/go-sqlbuilder) - A flexible and powerful SQL string builder library plus a zero-config ORM. Stars:`340`.
* [gormt](https://github.com/xxjwxc/gormt) - Mysql database to golang gorm struct. Stars:`261`.
* [Zoom](https://github.com/albrow/zoom) - Blazing-fast datastore and querying engine built on Redis. Stars:`247`.
* [grimoire](https://github.com/Fs02/grimoire) - Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3). Stars:`123`.
* [go-store](https://github.com/gosuri/go-store) - Simple and fast Redis backed key-value store library for Go. Stars:`100`.
* [Marlow](https://github.com/dadleyy/marlow) - Generated ORM from project structs for compile time safety assurances. Stars:`72`.
* [rel](https://github.com/Fs02/rel) - Golang SQL Repository Layer for Clean (Onion) Architecture. Stars:`40`.
* [go-firestorm](https://github.com/jschoedt/go-firestorm) - A simple ORM for Google/Firebase Cloud Firestore. Stars:`6`.
* [lore](https://github.com/abrahambotros/lore) - Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go. Stars:`5`.
* [beego orm](https://github.com/astaxie/beego/tree/master/orm) - Powerful orm framework for go. Support: pq/mysql/sqlite3.

## Package Management

*Official tooling for dependency and package management*

* [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) - Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

*Official experimental tooling for package management*

* [dep](https://github.com/golang/dep) - Go dependency tool. Stars:`13.0K`.
* [vgo](https://go.googlesource.com/vgo/) - Versioned Go.

*Unofficial libraries for package and dependency management.*

* [glide](https://github.com/Masterminds/glide) - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip. Stars:`7.9K`.
* [godep](https://github.com/tools/godep) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies. Stars:`5.7K`.
* [govendor](https://github.com/kardianos/govendor) - Go Package Manager. Go vendor tool that works with the standard vendor file. Stars:`5.0K`.
* [gopm](https://github.com/gpmgo/gopm) - Go Package Manager. Stars:`2.5K`.
* [gom](https://github.com/mattn/gom) - Go Manager - bundle for go. Stars:`1.4K`.
* [gpm](https://github.com/pote/gpm) - Barebones dependency manager for Go. Stars:`1.2K`.
* [goop](https://github.com/nitrous-io/goop) - Simple dependency manager for Go (golang), inspired by Bundler. Stars:`779`.
* [nut](https://github.com/jingweno/nut) - Vendor Go dependencies. Stars:`244`.
* [johnny-deps](https://github.com/VividCortex/johnny-deps) - Minimal dependency version using Git. Stars:`214`.
* [gigo](https://github.com/LyricalSecurity/gigo) - PIP-like dependency tool for golang, with support for private repositories and hashes. Stars:`198`.
* [VenGO](https://github.com/DamnWidget/VenGO) - create and manage exportable isolated go virtual environments. Stars:`116`.
* [mvn-golang](https://github.com/raydac/mvn-golang) - plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure. Stars:`95`.
* [gop](https://github.com/lunny/gop) - Build and manage your Go applications out of GOPATH. Stars:`50`.

## Performance

* [jaeger](https://github.com/jaegertracing/jaeger) - A distributed tracing system. Stars:`10.0K`.
* [profile](https://github.com/pkg/profile) - Simple profiling support package for Go. Stars:`1.2K`.
* [tracer](https://github.com/kamilsk/tracer) - Simple, lightweight tracing. Stars:`17`.

## Query Language

* [graphql-go](https://github.com/graphql-go/graphql) - Implementation of GraphQL for Go. Stars:`5.9K`.
* [graphql](https://github.com/neelance/graphql-go) - GraphQL server with a focus on ease of use. Stars:`3.1K`.
* [gojsonq](https://github.com/thedevsaddam/gojsonq) - A simple Go package to Query over JSON Data. Stars:`1.1K`.
* [jsonql](https://github.com/elgs/jsonql) - JSON query expression library in Golang. Stars:`216`.
* [rql](https://github.com/a8m/rql) - Resource Query Language for REST API. Stars:`130`.
* [graphql](https://github.com/tmc/graphql) - graphql parser + utilities. Stars:`52`.
* [jsonslice](https://github.com/bhmj/jsonslice) - Jsonpath queries with advanced filters. Stars:`32`.
* [straf](https://github.com/SonicRoshan/straf) - Easily Convert Golang structs to GraphQL objects. Stars:`7`.

## Resource Embedding

* [packr](https://github.com/gobuffalo/packr) - The simple and easy way to embed static files into Go binaries. Stars:`2.6K`.
* [statik](https://github.com/rakyll/statik) - Embeds static files into a Go executable. Stars:`2.4K`.
* [go.rice](https://github.com/GeertJohan/go.rice) - go.rice is a Go package that makes working with resources such as html,js,css,images and templates very easy. Stars:`1.8K`.
* [vfsgen](https://github.com/shurcooL/vfsgen) - Generates a vfsdata.go file that statically implements the given virtual filesystem. Stars:`776`.
* [esc](https://github.com/mjibson/esc) - Embeds files into Go programs and provides http.FileSystem interfaces to them. Stars:`515`.
* [fileb0x](https://github.com/UnnoTed/fileb0x) - Simple tool to embed files in go with focus on "customization" and ease to use. Stars:`494`.
* [go-resources](https://github.com/omeid/go-resources) - Unfancy resources embedding with Go. Stars:`160`.
* [statics](https://github.com/go-playground/statics) - Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks. Stars:`54`.
* [templify](https://github.com/wlbr/templify) - Embed external template files into Go code to create single file binaries. Stars:`22`.
* [go-embed](https://github.com/pyros2097/go-embed) - Generates go code to embed resource files into your library or executable. Stars:`18`.

## Science and Data Analysis

*Libraries for scientific computing and data analyzing.*

* [gonum](https://github.com/gonum/gonum) - Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more. Stars:`3.4K`.
* [stats](https://github.com/montanaflynn/stats) - Statistics package with common functions missing from the Golang standard library. Stars:`1.4K`.
* [gosl](https://github.com/cpmech/gosl) - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more. Stars:`1.4K`.
* [gonum/plot](https://github.com/gonum/plot) - gonum/plot provides an API for building and drawing plots in Go. Stars:`1.3K`.
* [streamtools](https://github.com/nytlabs/streamtools) - general purpose, graphical tool for dealing with streams of data. Stars:`1.3K`.
* [go-dsp](https://github.com/mjibson/go-dsp) - Digital Signal Processing for Go. Stars:`657`.
* [goraph](https://github.com/gyuho/goraph) - Pure Go graph theory library(data structure, algorith visualization). Stars:`618`.
* [chart](https://github.com/vdobler/chart) - Simple Chart Plotting library for Go. Supports many graphs types. Stars:`615`.
* [graph](https://github.com/yourbasic/graph) - Library of basic graph algorithms. Stars:`285`.
* [ewma](https://github.com/VividCortex/ewma) - Exponentially-weighted moving averages. Stars:`283`.
* [orb](https://github.com/paulmach/orb) - 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support. Stars:`243`.
* [gohistogram](https://github.com/VividCortex/gohistogram) - Approximate histograms for data streams. Stars:`136`.
* [dataframe-go](https://github.com/rocketlaunchr/dataframe-go) - Dataframes for machine-learning and statistics (similar to pandas). Stars:`132`.
* [TextRank](https://github.com/DavidBelicza/TextRank) - TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support. Stars:`89`.
* [sparse](https://github.com/james-bowman/sparse) - Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries. Stars:`81`.
* [pagerank](https://github.com/alixaxel/pagerank) - Weighted PageRank algorithm implemented in Go. Stars:`53`.
* [geom](https://github.com/skelterjohn/geom) - 2D geometry for golang. Stars:`43`.
* [evaler](https://github.com/soniah/evaler) - Simple floating point arithmetic expression evaluator. Stars:`40`.
* [goent](https://github.com/kzahedi/goent) - GO Implementation of Entropy Measures. Stars:`15`.
* [triangolatte](https://github.com/tchayen/triangolatte) - 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs. Stars:`12`.
* [ode](https://github.com/ChristopherRabotin/ode) - Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions. Stars:`11`.
* [GoStats](https://github.com/OGFris/GoStats) - GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions. Stars:`11`.
* [piecewiselinear](https://github.com/sgreben/piecewiselinear) - Tiny linear interpolation library. Stars:`10`.
* [PiHex](https://github.com/claygod/PiHex) - Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi. Stars:`8`.
* [assocentity](https://github.com/ndabAP/assocentity) - Package assocentity returns the average distance from words to a given entity. Stars:`6`.
* [go-gt](https://github.com/ThePaw/go-gt) - Graph theory algorithms written in "Go" language. Stars:`5`.
* [rootfinding](https://github.com/khezen/rootfinding) - root-finding algorithms library for finding roots of quadratic functions. Stars:`3`.
* [bradleyterry](https://github.com/seanhagen/bradleyterry) - Provides a Bradley-Terry Model for pairwise comparisons. Stars:`2`.

## Security

*Libraries that are used to help make your application more secure.*

* [lego](https://github.com/xenolf/lego) - Pure Go ACME client library and CLI tool (for use with Let's Encrypt). Stars:`3.8K`.
* [Cameradar](https://github.com/Ullaakut/cameradar) - Tool and library to remotely hack RTSP streams from surveillance cameras. Stars:`2.0K`.
* [acmetool](https://github.com/hlandau/acme) - ACME (Let's Encrypt) client tool with automatic renewal. Stars:`1.7K`.
* [memguard](https://github.com/awnumar/memguard) - A pure Go library for handling sensitive values in memory. Stars:`1.7K`.
* [secure](https://github.com/unrolled/secure) - HTTP middleware for Go that facilitates some quick security wins. Stars:`1.4K`.
* [acra](https://github.com/cossacklabs/acra) - Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system. Stars:`521`.
* [nacl](https://github.com/kevinburke/nacl) - Go implementation of the NaCL set of API's. Stars:`464`.
* [BadActor](https://github.com/jaredfolkins/badactor) - In-memory, application-driven jailer built in the spirit of fail2ban. Stars:`260`.
* [passlib](https://github.com/hlandau/passlib) - Futureproof password hashing library. Stars:`232`.
* [ssh-vault](https://github.com/ssh-vault/ssh-vault) - encrypt/decrypt using ssh keys. Stars:`215`.
* [simple-scrypt](https://github.com/elithrar/simple-scrypt) - Scrypt package with a simple, obvious API and automatic cost calibration built-in. Stars:`162`.
* [go-yara](https://github.com/hillu/go-yara) - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)". Stars:`146`.
* [argon2pw](https://github.com/raja/argon2pw) - Argon2 password hash generation with constant-time password comparison. Stars:`80`.
* [goSecretBoxPassword](https://github.com/dwin/goSecretBoxPassword) - A probably paranoid package for securely hashing and encrypting passwords. Stars:`34`.
* [goArgonPass](https://github.com/dwin/goArgonPass) - Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations. Stars:`12`.
* [certificates](https://github.com/mvmaasakkers/certificates) - An opinionated tool for generating tls certificates. Stars:`10`.
* [sslmgr](https://github.com/adrianosela/sslmgr) - SSL certificates made easy with a high level wrapper around acme/autocert. Stars:`8`.
* [go-generate-password](https://github.com/m1/go-generate-password) - Password generator that can be used on the cli or as a library. Stars:`3`.
* [Interpol](https://bitbucket.org/vahidi/interpol) - Rule-based data generator for fuzzing and penetration testing.
* [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) - Auto provision Let's Encrypt certificates and start a TLS server.

## Serialization

*Libraries and tools for binary serialization.*

* [jsoniter](https://github.com/json-iterator/go) - High-performance 100% compatible drop-in replacement of "encoding/json". Stars:`7.0K`.
* [goprotobuf](https://github.com/golang/protobuf) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers. Stars:`6.0K`.
* [gogoprotobuf](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets. Stars:`3.4K`.
* [mapstructure](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures. Stars:`3.0K`.
* [go-codec](https://github.com/ugorji/go) - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support. Stars:`1.3K`.
* [colfer](https://github.com/pascaldekloe/colfer) - Code generation for the Colfer binary format. Stars:`507`.
* [csvutil](https://github.com/jszwec/csvutil) - High Performance, idiomatic CSV record encoding and decoding to native Go structures. Stars:`339`.
* [go-capnproto](https://github.com/glycerine/go-capnproto) - Cap'n Proto library and parser for go. Stars:`276`.
* [php_session_decoder](https://github.com/yvasiyarov/php_session_decoder) - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions. Stars:`133`.
* [structomap](https://github.com/tuvistavie/structomap) - Library to easily and dynamically generate maps from static structures. Stars:`107`.
* [cbor](https://github.com/fxamacker/cbor) - Small, safe, and easy CBOR encoding and decoding library. Stars:`100`.
* [bambam](https://github.com/glycerine/bambam) - generator for Cap'n Proto schemas from go. Stars:`60`.
* [asn1](https://github.com/PromonLogicalis/asn1) - Asn.1 BER and DER encoding library for golang. Stars:`43`.
* [binstruct](https://github.com/ghostiam/binstruct) - Golang binary decoder for mapping data into the structure. Stars:`14`.
* [fwencoder](https://github.com/o1egl/fwencoder) - Fixed width file parser (encoding and decoding library) for Go. Stars:`10`.
* [bel](https://github.com/32leaves/bel) - Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC. Stars:`8`.
* [pletter](https://github.com/vimeda/pletter) - A standard way to wrap a proto message for message brokers. Stars:`7`.
* [fixedwidth](https://github.com/huydang284/fixedwidth) - Fixed-width text formatting (UTF-8 supported). Stars:`4`.

## Server Applications

* [etcd](https://github.com/coreos/etcd) - Highly-available key value store for shared configuration and service discovery. Stars:`29.1K`.
* [Caddy](https://github.com/mholt/caddy) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use. Stars:`25.8K`.
* [minio](https://github.com/minio/minio) - Minio is a distributed object storage server. Stars:`19.8K`.
* [RoadRunner](https://github.com/spiral/roadrunner) - High-performance PHP application server, load-balancer and process manager. Stars:`3.8K`.
* [devd](https://github.com/cortesi/devd) - Local webserver for developers. Stars:`2.9K`.
* [algernon](https://github.com/xyproto/algernon) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber. Stars:`1.7K`.
* [SFTPGo](https://github.com/drakkan/sftpgo) - Full featured and highly configurable SFTP server software. Stars:`1.2K`.
* [flipt](https://github.com/markphelps/flipt) - A self contained feature flag solution written in Go and Vue.js Stars:`1.1K`.
* [Flagr](https://github.com/checkr/flagr) - Flagr is an open-source feature flagging and A/B testing service. Stars:`967`.
* [Fider](https://github.com/getfider/fider) - Fider is an open platform to collect and organize customer feedback. Stars:`927`.
* [discovery](https://github.com/Bilibili/discovery) - A registry for resilient mid-tier load balancing and failover. Stars:`820`.
* [jackal](https://github.com/ortuman/jackal) - An XMPP server written in Go. Stars:`766`.
* [dudeldu](https://github.com/krotik/dudeldu) - A simple SHOUTcast server. Stars:`107`.
* [psql-streamer](https://github.com/blind-oracle/psql-streamer) - Stream database events from PostgreSQL to Kafka. Stars:`12`.
* [nginx-prometheus](https://github.com/blind-oracle/nginx-prometheus) - Nginx log parser and exporter to Prometheus. Stars:`9`.
* [nsq](http://nsq.io/) - A realtime distributed messaging platform.
* [riemann-relay](https://github.com/blind-oracle/riemann-relay) - Relay to load-balance Riemann events and/or convert them to Carbon.
* [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.
* [yakvs](https://git.sci4me.com/sci4me/yakvs) - Small, networked, in-memory key-value store.


## Stream Processing

*Libraries and tools for stream processing and reactive programming.*

* [go-streams](https://github.com/reugn/go-streams) - Go stream processing library. Stars:`262`.

## Template Engines

*Libraries and tools for templating and lexing.*

* [gofpdf](https://github.com/jung-kurt/gofpdf) - PDF document generator with high level support for text, drawing and images. Stars:`3.5K`.
* [quicktemplate](https://github.com/valyala/quicktemplate) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it. Stars:`1.6K`.
* [pongo2](https://github.com/flosch/pongo2) - Django-like template-engine for Go. Stars:`1.6K`.
* [hero](https://github.com/shiyanhui/hero) - Hero is a handy, fast and powerful go template engine. Stars:`1.3K`.
* [mustache](https://github.com/hoisie/mustache) - Go implementation of the Mustache template language. Stars:`976`.
* [amber](https://github.com/eknkc/amber) - Amber is an elegant templating engine for Go Programming Language It is inspired from HAML and Jade. Stars:`836`.
* [ace](https://github.com/yosssi/ace) - Ace is an HTML template engine for Go, inspired by Slim and Jade. Ace is a refinement of Gold. Stars:`775`.
* [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang. Stars:`726`.
* [jet](https://github.com/CloudyKit/jet) - Jet template engine. Stars:`610`.
* [ego](https://github.com/benbjohnson/ego) - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled. Stars:`423`.
* [raymond](https://github.com/aymerick/raymond) - Complete handlebars implementation in Go. Stars:`359`.
* [fasttemplate](https://github.com/valyala/fasttemplate) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](http://golang.org/pkg/text/template/). Stars:`334`.
* [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/). Stars:`146`.
* [liquid](https://github.com/osteele/liquid) - Go implementation of Shopify Liquid templates. Stars:`92`.
* [maroto](https://github.com/johnfercher/maroto) - A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple. Stars:`92`.
* [goview](https://github.com/foolin/goview) - Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application. Stars:`89`.
* [kasia.go](https://github.com/ziutek/kasia.go) - Templating system for HTML and other text documents - go implementation. Stars:`71`.
* [velvet](https://github.com/gobuffalo/velvet) - Complete handlebars implementation in Go. Stars:`70`.
* [damsel](https://github.com/dskinner/damsel) - Markup language featuring html outlining via css-selectors, extensible via pkg html/template and others. Stars:`21`.
* [extemplate](https://github.com/dannyvankooten/extemplate) - Tiny wrapper around html/template to allow for easy file-based template inheritance. Stars:`16`.
* [gospin](https://github.com/m1/gospin) - Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations. Stars:`5`.

## Testing

*Libraries for testing codebases and generating test data.*

* Testing Frameworks
    * [Testify](https://github.com/stretchr/testify) - Sacred extension to the standard go testing package. Stars:`9.4K`.
    * [go-cmp](https://github.com/google/go-cmp) - Package for comparing Go values in tests. Stars:`1.5K`.
    * [httpexpect](https://github.com/gavv/httpexpect) - Concise, declarative, and easy to use end-to-end HTTP and REST API testing. Stars:`1.2K`.
    * [godog](https://github.com/DATA-DOG/godog) - Cucumber or Behat like BDD framework for Go. Stars:`882`.
    * [baloo](https://github.com/h2non/baloo) - Expressive and versatile end-to-end HTTP API testing made easy. Stars:`666`.
    * [goblin](https://github.com/franela/goblin) - Mocha like testing framework fo Go. Stars:`649`.
    * [testfixtures](https://github.com/go-testfixtures/testfixtures) - A helper for Rails' like test fixtures to test database applications. Stars:`423`.
    * [go-vcr](https://github.com/dnaeon/go-vcr) - Record and replay your HTTP interactions for fast, deterministic and accurate tests. Stars:`364`.
    * [go-mutesting](https://github.com/zimmski/go-mutesting) - Mutation testing for Go source code. Stars:`323`.
    * [gofight](https://github.com/appleboy/gofight) - API Handler Testing for Golang Router framework. Stars:`289`.
    * [frisby](https://github.com/verdverm/frisby) - REST API testing framework. Stars:`255`.
    * [go-carpet](https://github.com/msoap/go-carpet) - Tool for viewing test coverage in terminal. Stars:`201`.
    * [charlatan](https://github.com/percolate/charlatan) - Tool to generate fake interface implementations for tests. Stars:`192`.
    * [gotest.tools](https://github.com/gotestyourself/gotest.tools) - A collection of packages to augment the go testing package and support common patterns. Stars:`141`.
    * [endly](https://github.com/viant/endly) - Declarative end to end functional testing. Stars:`120`.
    * [GoSpec](https://github.com/orfjackal/gospec) - BDD-style testing framework for the Go programming language. Stars:`112`.
    * [cupaloy](https://github.com/bradleyjkemp/cupaloy) - Simple snapshot testing addon for your test framework. Stars:`97`.
    * [dbcleaner](https://github.com/khaiql/dbcleaner) - Clean database for testing purpose, inspired by `database_cleaner` in Ruby. Stars:`97`.
    * [go-testdeep](https://github.com/maxatome/go-testdeep) - Extremely flexible golang deep comparison, extends the go testing package. Stars:`75`.
    * [wstest](https://github.com/posener/wstest) - Websocket client for unit-testing a websocket http.Handler. Stars:`71`.
    * [gospecify](https://github.com/stesla/gospecify) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec. Stars:`53`.
    * [restit](https://github.com/yookoala/restit) - Go micro framework to help writing RESTful API integration test. Stars:`50`.
    * [commander](https://github.com/SimonBaeumer/commander) - Tool for testing cli applications on windows, linux and osx. Stars:`42`.
    * [gomatch](https://github.com/jfilipczyk/gomatch) - library created for testing JSON against patterns. Stars:`32`.
    * [jsonassert](https://github.com/kinbiko/jsonassert) - Package for verifying that your JSON payloads are serialized correctly. Stars:`30`.
    * [dsunit](https://github.com/viant/dsunit) - Datastore testing for SQL, NoSQL, structured files. Stars:`27`.
    * [Hamcrest](https://github.com/rdrdr/hamcrest) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results. Stars:`27`.
    * [assert](https://github.com/go-playground/assert) - Basic Assertion Library used along side native go testing, with building blocks for custom assertions. Stars:`17`.
    * [testcase](https://github.com/adamluzsi/testcase) - Idiomatic testing framework for Behavior Driven Development. Stars:`12`.
    * [flute](https://github.com/suzuki-shunsuke/flute) - HTTP client testing framework. Stars:`11`.
    * [badio](https://github.com/cavaliercoder/badio) - Extensions to Go's `testing/iotest` package. Stars:`9`.
    * [gosuite](https://github.com/pavlo/gosuite) - Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests. Stars:`9`.
    * [gocrest](https://github.com/corbym/gocrest) - Composable hamcrest-like matchers for Go assertions. Stars:`9`.
    * [embedded-postgres](https://github.com/fergusstrange/embedded-postgres) - Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test. Stars:`8`.
    * [gogiven](https://github.com/corbym/gogiven) - YATSPEC-like BDD testing framework for Go. Stars:`8`.
    * [schema](https://github.com/jgroeneveld/schema) - Quick and easy expression matching for JSON schemas used in requests and responses. Stars:`8`.
    * [testsql](https://github.com/zhulongcheng/testsql) - Generate test data from SQL files before testing and clear it after finished. Stars:`7`.
    * [biff](https://github.com/fulldump/biff) - Bifurcation testing framework, BDD compatible. Stars:`7`.
    * [Tt](https://github.com/vcaesar/tt) - Simple and colorful test tools. Stars:`6`.
    * [trial](https://github.com/jgroeneveld/trial) - Quick and easy extendable assertions without introducing much boilerplate. Stars:`4`.
    * [apitest](https://apitest.dev) - Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
    * [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd) - Convert markdown snippets into testable go code.
    * [ginkgo](http://onsi.github.io/ginkgo/) - BDD Testing Framework for Go.
    * [gocheck](http://labix.org/gocheck) - More advanced testing framework alternative to gotest.
    * [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload.
    * [gomega](http://onsi.github.io/gomega/) - Rspec like matcher/assertion library.

* Mock
    * [gomock](https://github.com/golang/mock) - Mocking framework for the Go programming language. Stars:`3.5K`.
    * [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database interactions. Stars:`2.2K`.
    * [hoverfly](https://github.com/SpectoLabs/hoverfly) - HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI. Stars:`1.5K`.
    * [gock](https://github.com/h2non/gock) - Versatile HTTP mocking made easy. Stars:`940`.
    * [httpmock](https://github.com/jarcoal/httpmock) - Easy mocking of HTTP responses from external resources. Stars:`687`.
    * [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects. Stars:`389`.
    * [minimock](https://github.com/gojuno/minimock) - Mock generator for Go interfaces. Stars:`288`.
    * [go-txdb](https://github.com/DATA-DOG/go-txdb) - Single transaction based database driver mainly for testing purposes. Stars:`216`.
    * [govcr](https://github.com/seborama/govcr) - HTTP mock for Golang: record and replay HTTP interactions for offline testing. Stars:`87`.
    * [mockhttp](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter. Stars:`22`.

* Fuzzing and delta-debugging/reducing/shrinking.
    * [go-fuzz](https://github.com/dvyukov/go-fuzz) - Randomized testing system. Stars:`3.3K`.
    * [gofuzz](https://github.com/google/gofuzz) - Library for populating go objects with random values. Stars:`690`.
    * [Tavor](https://github.com/zimmski/tavor) - Generic fuzzing and delta-debugging framework. Stars:`218`.

* Selenium and browser control tools.
    * [chromedp](https://github.com/knq/chromedp) - a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol. Stars:`4.2K`.
    * [selenoid](https://github.com/aerokube/selenoid) - alternative Selenium hub server that launches browsers within containers. Stars:`1.4K`.
    * [cdp](https://github.com/mafredri/cdp) - Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it. Stars:`398`.
    * [ggr](https://github.com/aerokube/ggr) - a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs. Stars:`233`.

* Fail injection
    * [failpoint](https://github.com/pingcap/failpoint) - An implementation of [failpoints](http://www.freebsd.org/cgi/man.cgi?query=fail) for Golang. Stars:`471`.

## Text Processing

*Libraries for parsing and manipulating texts.*

* Specific Formats
    * [colly](https://github.com/asciimoo/colly) - Fast and Elegant Scraping Framework for Gophers. Stars:`9.8K`.
    * [GoQuery](https://github.com/PuerkitoBio/goquery) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language. Stars:`8.2K`.
    * [blackfriday](https://github.com/russross/blackfriday) - Markdown processor in Go. Stars:`4.2K`.
    * [toml](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection). Stars:`3.0K`.
    * [sh](https://github.com/mvdan/sh) - Shell parser and formatter. Stars:`2.4K`.
    * [go-humanize](https://github.com/dustin/go-humanize) - Formatters for time, numbers, and memory size to human readable format. Stars:`2.1K`.
    * [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer. Stars:`1.4K`.
    * [inject](https://github.com/facebookgo/inject) - Package inject provides a reflect based injector. Stars:`1.2K`.
    * [gofeed](https://github.com/mmcdole/gofeed) - Parse RSS and Atom feeds in Go. Stars:`1.2K`.
    * [go-toml](https://github.com/pelletier/go-toml) - Go library for the TOML format with query support and handy cli tools. Stars:`686`.
    * [commonregex](https://github.com/mingrammer/commonregex) - A collection of common regular expressions for Go. Stars:`572`.
    * [slug](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support. Stars:`489`.
    * [mxj](https://github.com/clbanning/mxj) - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages. Stars:`353`.
    * [dataflowkit](https://github.com/slotix/dataflowkit) - Web scraping Framework to turn websites into structured data. Stars:`341`.
    * [gographviz](https://github.com/awalterschulze/gographviz) - Parses the Graphviz DOT language. Stars:`338`.
    * [gotext](https://github.com/leonelquinteros/gotext) - GNU gettext utilities for Go. Stars:`247`.
    * [go-runewidth](https://github.com/mattn/go-runewidth) - Functions to get fixed width of the character or string. Stars:`239`.
    * [htmlquery](https://github.com/antchfx/htmlquery) - An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression. Stars:`182`.
    * [goq](https://github.com/andrewstuart/goq) - Declarative unmarshaling of HTML using struct tags with jQuery syntax (uses GoQuery). Stars:`158`.
    * [go-nmea](https://github.com/adrianmo/go-nmea) - NMEA parser library for the Go language. Stars:`108`.
    * [sdp](https://github.com/gortc/sdp) - SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)]. Stars:`80`.
    * [align](https://github.com/Guitarbum722/align) - A general purpose application that aligns text. Stars:`61`.
    * [genex](https://github.com/alixaxel/genex) - Count and expand Regular Expressions into all matching Strings. Stars:`57`.
    * [go-slugify](https://github.com/mozillazg/go-slugify) - Make pretty slug with multiple languages support. Stars:`57`.
    * [goribot](https://github.com/zhshch2002/goribot) - A simple golang spider/scraping framework,build a spider in 3 lines. Stars:`55`.
    * [editorconfig-core-go](https://github.com/editorconfig/editorconfig-core-go) - Editorconfig file parser and manipulator for Go. Stars:`53`.
    * [go-zero-width](https://github.com/trubitsyn/go-zero-width) - Zero-width character detection and removal for Go. Stars:`45`.
    * [guesslanguage](https://github.com/endeveit/guesslanguage) - Functions to determine the natural language of a unicode text. Stars:`45`.
    * [allot](https://github.com/sbstjn/allot) - Placeholder and wildcard text parsing for CLI tools and bots. Stars:`38`.
    * [goregen](https://github.com/zach-klippenstein/goregen) - Library for generating random strings from regular expressions. Stars:`38`.
    * [go-fixedwidth](https://github.com/ianlopshire/go-fixedwidth) - Fixed-width text formatting (encoder/decoder with reflection). Stars:`33`.
    * [go-vcard](https://github.com/emersion/go-vcard) - Parse and format vCard. Stars:`32`.
    * [gonameparts](https://github.com/polera/gonameparts) - Parses human names into individual name parts. Stars:`30`.
    * [Slugify](https://github.com/avelino/slugify) - Go slugify application that handles string. Stars:`27`.
    * [did](https://github.com/ockam-network/did) - DID (Decentralized Identifiers) Parser and Stringer in Go. Stars:`26`.
    * [codetree](https://github.com/aerogo/codetree) - Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure. Stars:`13`.
    * [enca](https://github.com/endeveit/enca) - Minimal cgo bindings for [libenca](http://cihar.com/software/enca/). Stars:`8`.
    * [syndfeed](https://github.com/zhengchun/syndfeed) - A syndication feed for Atom 1.0 and RSS 2.0. Stars:`7`.
    * [bbConvert](https://github.com/CalebQ42/bbConvert) - Converts bbCode to HTML that allows you to add support for custom bbCode tags. Stars:`5`.
    * [doi](https://github.com/hscells/doi) - Document object identifier (doi) parser in Go. Stars:`4`.
    * [encdec](https://github.com/mickep76/encdec) - Package provides a generic interface to encoders and decodersa. Stars:`3`.
    * [ltsv](https://github.com/Wing924/ltsv) - High performance [LTSV (Labeled Tab Separeted Value)](http://ltsv.org/) reader for Go. Stars:`3`.
    * [gommon/bytes](https://github.com/labstack/gommon/tree/master/bytes) - Format bytes to string.
    * [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
* Utility
    * [xurls](https://github.com/mvdan/xurls) - Extract urls from text. Stars:`566`.
    * [gotabulate](https://github.com/bndr/gotabulate) - Easily pretty-print your tabular data with Go. Stars:`222`.
    * [radix](https://github.com/yourbasic/radix) - fast string sorting algorithm. Stars:`152`.
    * [parth](https://github.com/codemodus/parth) - URL path segmentation parsing. Stars:`34`.
    * [gofuckyourself](https://github.com/JoshuaDoes/gofuckyourself) - A sanitization-based swear filter for Go. Stars:`19`.
    * [xj2go](https://github.com/stackerzzq/xj2go) - Convert xml or json to go struct. Stars:`18`.
    * [kace](https://github.com/codemodus/kace) - Common case conversions covering common initialisms. Stars:`12`.
    * [parseargs-go](https://github.com/nproc/parseargs-go) - string argument parser that understands quotes and backslashes. Stars:`7`.
    * [Tagify](https://github.com/zoomio/tagify) - Produces a set of tags from given source. Stars:`3`.
    * [TySug](https://github.com/Dynom/TySug) - Alternative suggestions with respect to keyboard layouts. Stars:`3`.
	* [textwrap](https://github.com/isbm/textwrap) - Implementation of `textwrap` module from Python. Stars:`1`.

## Third-party APIs

*Libraries for accessing third party APIs.*

* [aws-sdk-go](https://github.com/aws/aws-sdk-go) - The official AWS SDK for the Go programming language. Stars:`5.5K`.
* [github](https://github.com/google/go-github) - Go library for accessing the GitHub REST API v3. Stars:`5.4K`.
* [slack](https://github.com/nlopes/slack) - Slack API in Go. Stars:`2.7K`.
* [google](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go. Stars:`2.1K`.
* [google-cloud](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library. Stars:`2.0K`.
* [discordgo](https://github.com/bwmarrin/discordgo) - Go bindings for the Discord Chat API. Stars:`1.1K`.
* [stripe](https://github.com/stripe/stripe-go) - Go client for the Stripe API. Stars:`1.1K`.
* [anaconda](https://github.com/ChimeraCoder/anaconda) - Go client library for the Twitter 1.1 API. Stars:`1.0K`.
* [go-twitter](https://github.com/dghubble/go-twitter) - Go client library for the Twitter v1.1 APIs. Stars:`885`.
* [minio-go](https://github.com/minio/minio-go) - Minio Go Library for Amazon S3 compatible cloud storage. Stars:`844`.
* [facebook](https://github.com/huandu/facebook) - Go Library that supports the Facebook Graph API. Stars:`823`.
* [go-jira](https://github.com/andygrunwald/go-jira) - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira) Stars:`663`.
* [githubql](https://github.com/shurcooL/githubql) - Go library for accessing the GitHub GraphQL API v4. Stars:`583`.
* [webhooks](https://github.com/go-playground/webhooks) - Webhook receiver for GitHub and Bitbucket. Stars:`433`.
* [geo-golang](https://github.com/codingsince1985/geo-golang) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](http://open.mapquestapi.com/geocoding/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](http://geocoder.opencagedata.com/api.html), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs. Stars:`335`.
* [paypal](https://github.com/logpacker/PayPal-Go-SDK) - Wrapper for PayPal payment API. Stars:`332`.
* [go-marathon](https://github.com/gambol99/go-marathon) - Go library for interacting with Mesosphere's Marathon PAAS. Stars:`192`.
* [ethrpc](https://github.com/onrik/ethrpc) - Go bindings for Ethereum JSON RPC API. Stars:`181`.
* [gostorm](https://github.com/jsgilmore/gostorm) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells. Stars:`125`.
* [Trello](https://github.com/adlio/trello) - Go wrapper for the Trello API. Stars:`125`.
* [Medium](https://github.com/Medium/medium-sdk-go) - Golang SDK for Medium's OAuth2 API. Stars:`122`.
* [hipchat (xmpp)](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP. Stars:`112`.
* [hipchat](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API. Stars:`108`.
* [go-trending](https://github.com/andygrunwald/go-trending) - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github. Stars:`105`.
* [cachet](https://github.com/andygrunwald/cachet) - Go client library for [Cachet (open source status page system)](https://cachethq.io/). Stars:`78`.
* [pushover](https://github.com/gregdel/pushover) - Go wrapper for the Pushover API. Stars:`66`.
* [wit-go](https://github.com/wit-ai/wit-go) - Go client for wit.ai HTTP API. Stars:`59`.
* [megos](https://github.com/andygrunwald/megos) - Client library for accessing an [Apache Mesos](http://mesos.apache.org/) cluster. Stars:`57`.
* [clarifai](https://github.com/samuelcouch/clarifai) - Go client library for interfacing with the Clarifai API. Stars:`57`.
* [igdb](https://github.com/Henry-Sarabia/igdb) - Go client for the [Internet Game Database API](https://api.igdb.com/). Stars:`55`.
* [gads](https://github.com/emiddleton/gads) - Google Adwords Unofficial API. Stars:`47`.
* [circleci](https://github.com/jszwedko/go-circleci) - Go client library for interacting with CircleCI's API. Stars:`46`.
* [go-xkcd](https://github.com/nishanths/go-xkcd) - Go client for the xkcd API. Stars:`40`.
* [amazon-product-advertising-api](https://github.com/ngs/go-amazon-product-advertising-api) - Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html). Stars:`38`.
* [GoMusicBrainz](https://github.com/michiwend/gomusicbrainz) - Go MusicBrainz WS2 client library. Stars:`37`.
* [simples3](https://github.com/rhnvrm/simples3) - Simple no frills AWS S3 Library using REST with V4 Signing written in Go. Stars:`36`.
* [uptimerobot](https://github.com/bitfield/uptimerobot) - Go wrapper and command-line client for the Uptime Robot v2 API. Stars:`35`.
* [fcm](https://github.com/maddevsio/fcm) - Go library for Firebase Cloud Messaging. Stars:`34`.
* [golyrics](https://github.com/mamal72/golyrics) - Golyrics is a Go library to fetch music lyrics data from the Wikia website. Stars:`33`.
* [mixpanel](https://github.com/dukex/mixpanel) - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications. Stars:`30`.
* [translate](https://github.com/poorny/translate) - Go online translation package. Stars:`29`.
* [gcm](https://github.com/Aorioli/gcm) - Go library for Google Cloud Messaging. Stars:`29`.
* [gami](https://github.com/bit4bit/gami) - Go library for Asterisk Manager Interface. Stars:`27`.
* [go-unsplash](https://github.com/hbagdi/go-unsplash) - Go client library for the [Unsplash.com](https://unsplash.com) API. Stars:`25`.
* [spotify](https://github.com/rapito/go-spotify) - Go Library to access Spotify WEB API. Stars:`21`.
* [ynab](https://github.com/brunomvsouza/ynab.go) - Go wrapper for the YNAB API. Stars:`20`.
* [patreon-go](https://github.com/mxpv/patreon-go) - Go library for Patreon API. Stars:`19`.
* [steam](https://github.com/sostronk/go-steam) - Go Library to interact with Steam game servers. Stars:`19`.
* [shopify](https://github.com/rapito/go-shopify) - Go Library to make CRUD request to the Shopify API. Stars:`19`.
* [go-twitch](https://github.com/knspriggs/go-twitch) - Go client for interacting with the Twitch v3 API. Stars:`18`.
* [textbelt](https://github.com/dietsche/textbelt) - Go client for the textbelt.com txt messaging API. Stars:`16`.
* [brewerydb](https://github.com/naegelejd/brewerydb) - Go library for accessing the BreweryDB API. Stars:`16`.
* [codeship-go](https://github.com/codeship/codeship-go) - Go client library for interacting with Codeship's API v2. Stars:`16`.
* [go-myanimelist](https://github.com/nstratos/go-myanimelist) - Go client library for accessing the [MyAnimeList API](http://myanimelist.net/modules.php?go=api). Stars:`15`.
* [go-imgur](https://github.com/koffeinsource/go-imgur) - Go client library for [imgur](https://imgur.com) Stars:`14`.
* [coinpaprika-go](https://github.com/coinpaprika/coinpaprika-api-go-client) - Go client library for interacting with Coinpaprika's API. Stars:`12`.
* [google-analytics](https://github.com/chonthu/go-google-analytics) - Simple wrapper for easy google analytics reporting. Stars:`11`.
* [smite](https://github.com/sergiotapia/smitego) - Go package to wraps access to the Smite game API. Stars:`10`.
* [go-hacknews](https://github.com/PaulRosset/go-hacknews) - Tiny Go client for HackerNews API. Stars:`10`.
* [lastpass-go](https://github.com/ansd/lastpass-go) - Go client library for the [LastPass](https://www.lastpass.com/) API. Stars:`8`.
* [rrdaclient](https://github.com/Omie/rrdaclient) - Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP. Stars:`8`.
* [go-sptrans](https://github.com/sergioaugrod/go-sptrans) - Go client library for the SPTrans Olho Vivo API. Stars:`7`.
* [tumblr](https://github.com/mattcunningham/gumblr) - Go wrapper for the Tumblr v2 API. Stars:`6`.
* [go-here](https://github.com/abdullahselek/go-here) - Go client library around the HERE location based APIs. Stars:`6`.
* [google-email-audit-api](https://github.com/ngs/go-google-email-audit-api) - Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/). Stars:`6`.
* [go-sophos](https://github.com/esurdam/go-sophos) - Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies. Stars:`5`.
* [zooz](https://github.com/gojuno/go-zooz) - Go client for the Zooz API. Stars:`5`.
* [gomalshare](https://github.com/MonaxGT/gomalshare) - Go library MalShare API [malshare.com](http://www.malshare.com/) Stars:`4`.
* [go-chronos](https://github.com/axelspringer/go-chronos) - Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler Stars:`3`.
* [TripAdvisor](https://github.com/mrbenosborne/tripadvisor-golang) - Go wrapper for the TripAdvisor API. Stars:`1`.
* [vl-go](https://github.com/verifid/vl-go) - Go client library around the VerifID identity verification layer API. Stars:`1`.
* [playlyfe](https://github.com/playlyfe/playlyfe-go-sdk) - The Playlyfe Rest API Go SDK. Stars:`1`.
* [libgoffi](https://github.com/clevabit/libgoffi) - Library adapter toolbox for native [libffi](http://sourceware.org/libffi/) integration Stars:`1`.
* [go-telegraph](https://gitlab.com/toby3d/telegraph) - Telegraph publishing platform API client.

## Utilities

*General utilities and tools to make your life easier.*

* [fzf](https://github.com/junegunn/fzf) - Command-line fuzzy finder written in Go. Stars:`26.3K`.
* [hub](https://github.com/github/hub) - wrap git commands with additional functionality to interact with github from the terminal. Stars:`18.3K`.
* [delve](https://github.com/derekparker/delve) - Go debugger. Stars:`13.1K`.
* [ctop](https://github.com/bcicen/ctop) - [Top-like](http://ctop.sh) interface (e.g. htop) for container metrics. Stars:`9.3K`.
* [wuzz](https://github.com/asciimoo/wuzz) - Interactive cli tool for HTTP inspection. Stars:`8.5K`.
* [sqlx](https://github.com/jmoiron/sqlx) - provides a set of extensions on top of the excellent built-in database/sql package. Stars:`7.6K`.
* [peco](https://github.com/peco/peco) - Simplistic interactive filtering tool. Stars:`5.7K`.
* [usql](https://github.com/knq/usql) - usql is a universal command-line interface for SQL databases. Stars:`5.5K`.
* [goreleaser](https://github.com/goreleaser/goreleaser) - Deliver Go binaries as fast and easily as possible. Stars:`5.0K`.
* [godropbox](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications from Dropbox. Stars:`3.8K`.
* [realize](https://github.com/tockins/realize) - Go build system with file watchers and live reload. Run, build and watch file changes with custom paths. Stars:`3.5K`.
* [goreporter](https://github.com/wgliang/goreporter) - Golang tool that does static analysis, unit testing, code review and generate code quality report. Stars:`2.6K`.
* [hystrix-go](https://github.com/afex/hystrix-go) - Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker. Stars:`2.3K`.
* [panicparse](https://github.com/maruel/panicparse) - Groups similar goroutines and colorizes stack dump. Stars:`2.2K`.
* [Task](https://github.com/go-task/task) - simple "Make" alternative. Stars:`2.2K`.
* [minify](https://github.com/tdewolff/minify) - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats. Stars:`2.0K`.
* [go-funk](https://github.com/thoas/go-funk) - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...). Stars:`1.5K`.
* [mmake](https://github.com/tj/mmake) - Modern Make. Stars:`1.5K`.
* [Storm](https://github.com/asdine/storm) - Simple and powerful toolkit for BoltDB. Stars:`1.5K`.
* [mole](https://github.com/davrodpin/mole) - cli app to easily create ssh tunnels. Stars:`1.3K`.
* [mc](https://github.com/minio/mc) - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems. Stars:`1.3K`.
* [boilr](https://github.com/tmrts/boilr) - Blazingly fast CLI tool for creating projects from boilerplate templates. Stars:`1.0K`.
* [filetype](https://github.com/h2non/filetype) - Small package to infer the file type checking the magic numbers signature. Stars:`1.0K`.
* [mergo](https://github.com/imdario/mergo) - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements. Stars:`1.0K`.
* [spinner](https://github.com/briandowns/spinner) - Go package to easily provide a terminal spinner with options. Stars:`892`.
* [circuitbreaker](https://github.com/rubyist/circuitbreaker) - Circuit Breakers in Go. Stars:`839`.
* [git-time-metric](https://github.com/git-time-metric/gtm) - Simple, seamless, lightweight time tracking for Git. Stars:`763`.
* [jump](https://github.com/gsamokovarov/jump) - Jump helps you navigate faster by learning your habits. Stars:`760`.
* [immortal](https://github.com/immortal/immortal) - \*nix cross-platform (OS agnostic) supervisor. Stars:`632`.
* [htcat](https://github.com/htcat/htcat) - Parallel and Pipelined HTTP GET Utility. Stars:`503`.
* [go-dry](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go. Stars:`443`.
* [gopencils](https://github.com/bndr/gopencils) - Small and simple package to easily consume REST APIs. Stars:`427`.
* [godaemon](https://github.com/VividCortex/godaemon) - Utility to write daemons. Stars:`421`.
* [circuit](https://github.com/cep21/circuit) - An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern. Stars:`414`.
* [request](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™. Stars:`369`.
* [ergo](https://github.com/cristianoliveira/ergo) - The management of multiple local services running over different ports made easy. Stars:`351`.
* [koazee](https://github.com/wesovilabs/koazee) - Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays. Stars:`349`.
* [go-rate](https://github.com/beefsack/go-rate) - Timed rate limiter for Go. Stars:`295`.
* [gohper](https://github.com/cosiner/gohper) - Various tools/modules help for development. Stars:`248`.
* [clockwork](https://github.com/jonboulle/clockwork) - A simple fake clock for golang. Stars:`243`.
* [Deepcopier](https://github.com/ulule/deepcopier) - Simple struct copying for Go. Stars:`233`.
* [serve](https://github.com/syntaqx/serve) - A static http server anywhere you need. Stars:`199`.
* [gubrak](https://github.com/novalagung/gubrak) - Golang utility library with syntactic sugar. It's like lodash, but for golang. Stars:`198`.
* [go-trigger](https://github.com/sadlil/go-trigger) - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project. Stars:`189`.
* [retry](https://github.com/kamilsk/retry) - The most advanced functional mechanism to perform actions repetitively until successful. Stars:`186`.
* [mimetype](https://github.com/gabriel-vasile/mimetype) - Package for MIME type detection based on magic numbers. Stars:`180`.
* [go-bind-plugin](https://github.com/wendigo/go-bind-plugin) - go:generate tool for wrapping symbols exported by golang plugins (1.8 only). Stars:`164`.
* [gotenv](https://github.com/subosito/gotenv) - Load environment variables from `.env` or any `io.Reader` in Go. Stars:`161`.
* [rerun](https://github.com/ivpusic/rerun) - Recompiling and rerunning go apps when source changes. Stars:`155`.
* [util](https://github.com/shomali11/util) - Collection of useful utility functions. (strings, concurrency, manipulations, ...). Stars:`149`.
* [moldova](https://github.com/StabbyCutyou/moldova) - Utility for generating random data based on an input template. Stars:`148`.
* [Death](https://github.com/vrecan/death) - Managing go application shutdown with signals. Stars:`147`.
* [robustly](https://github.com/VividCortex/robustly) - Runs functions resiliently, catching and restarting panics. Stars:`140`.
* [apm](https://github.com/topfreegames/apm) - Process manager for Golang applications with an HTTP API. Stars:`133`.
* [toolbox](https://github.com/viant/toolbox) - Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer. Stars:`115`.
* [chyle](https://github.com/antham/chyle) - Changelog generator using a git repository with multiple configuration possibilities. Stars:`115`.
* [go-sitemap-generator](https://github.com/ikeikeikeike/go-sitemap-generator) - XML Sitemap generator written in Go. Stars:`114`.
* [onecache](https://github.com/adelowo/onecache) - Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc). Stars:`106`.
* [lrserver](https://github.com/jaschaephraim/lrserver) - LiveReload server for Go. Stars:`105`.
* [go-bsdiff](https://github.com/gabstv/go-bsdiff) - Pure Go bsdiff and bspatch libraries and CLI tools. Stars:`90`.
* [pm](https://github.com/VividCortex/pm) - Process (i.e. goroutine) manager with an HTTP API. Stars:`75`.
* [xferspdy](https://github.com/monmohan/xferspdy) - Xferspdy provides binary diff and patch library in golang. Stars:`70`.
* [UNIS](https://github.com/esemplastic/unis) - Common Architecture™ for String Utilities in Go. Stars:`69`.
* [go-health](https://github.com/Talento90/go-health) - Health package simplifies the way you add health check to your services. Stars:`68`.
* [mssqlx](https://github.com/linxGnu/mssqlx) - Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind. Stars:`67`.
* [multitick](https://github.com/VividCortex/multitick) - Multiplexor for aligned tickers. Stars:`61`.
* [repeat](https://github.com/ssgreg/repeat) - Go implementation of different backoff strategies useful for retrying operations and heartbeating. Stars:`61`.
* [minquery](https://github.com/icza/minquery) - MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off). Stars:`51`.
* [mimemagic](https://github.com/zRedShift/mimemagic) - Pure Go ultra performant MIME sniffing library/utility. Stars:`50`.
* [handy](https://github.com/miguelpragier/handy) - Many utilities and helpers like string handlers/formatters and validators. Stars:`49`.
* [golog](https://github.com/mlimaloureiro/golog) - Easy and lightweight CLI tool to time track your tasks. Stars:`47`.
* [go-astitodo](https://github.com/asticode/go-astitodo) - Parse TODOs in your GO code. Stars:`46`.
* [goseaweedfs](https://github.com/linxGnu/goseaweedfs) - SeaweedFS client library with almost full features. Stars:`43`.
* [goreadability](https://github.com/philipjkim/goreadability) - Webpage summary extractor using Facebook Open Graph and arc90's readability. Stars:`43`.
* [goback](https://github.com/carlescere/goback) - Go simple exponential backoff package. Stars:`42`.
* [gaper](https://github.com/maxcnunes/gaper) - Builds and restarts a Go project when it crashes or some watched file changes. Stars:`41`.
* [intrinsic](https://github.com/mengzhuo/intrinsic) - Use x86 SIMD without writing any assembly code. Stars:`41`.
* [copy-pasta](https://github.com/jutkko/copy-pasta) - Universal multi-workstation clipboard that uses S3 like backend for the storage. Stars:`38`.
* [golarm](https://github.com/msempere/golarm) - Fire alarms with system events. Stars:`37`.
* [pgo](https://github.com/arthurkushman/pgo) - Convenient functions for PHP community. Stars:`34`.
* [retry](https://github.com/thedevsaddam/retry) - Simple and easy retry mechanism package for Go. Stars:`34`.
* [myhttp](https://github.com/inancgumus/myhttp) - Simple API to make HTTP GET requests with timeout support. Stars:`33`.
* [gpath](https://github.com/tenntenn/gpath) - Library to simplify access struct fields with Go's expression in reflection. Stars:`32`.
* [retry-go](https://github.com/rafaeljesus/retry-go) - Retrying made simple and easy for golang. Stars:`31`.
* [rclient](https://github.com/zpatrick/rclient) - Readable, flexible, simple-to-use client for REST APIs. Stars:`31`.
* [beyond](https://github.com/wesovilabs/beyond) - The Go tool that will drive you to the AOP world! Stars:`30`.
* [dbt](https://github.com/nikogura/dbt) - A framework for running self-updating signed binaries from a central, trusted repository. Stars:`23`.
* [gostrutils](https://github.com/ik5/gostrutils) - Collections of string manipulation and conversion functions. Stars:`22`.
* [scan](https://github.com/blockloop/scan) - Scan golang `sql.Rows` directly to structs, slices, or primitive types. Stars:`22`.
* [ugo](https://github.com/alxrm/ugo) - ugo is slice toolbox with concise syntax for Go. Stars:`22`.
* [generate](https://github.com/go-playground/generate) - runs go generate recursively on a specified path or environment variable and can filter by regex. Stars:`21`.
* [filter](https://github.com/gookit/filter) - provide filtering, sanitizing, and conversion of Go data. Stars:`21`.
* [goplaceholder](https://github.com/michiwend/goplaceholder) - a small golang lib to generate placeholder images. Stars:`21`.
* [evaluator](https://github.com/nullne/evaluator) - Evaluate an expression dynamicly based on s-expression. It's simple and easy to extend. Stars:`20`.
* [slice](https://github.com/psampaz/slice) - Type-safe functions for common Go slice operations. Stars:`16`.
* [dlog](https://github.com/kirillDanshin/dlog) - Compile-time controlled logger to make your release smaller without removing debug calls. Stars:`15`.
* [go-httpheader](https://github.com/mozillazg/go-httpheader) - Go library for encoding structs into Header fields. Stars:`15`.
* [structs](https://github.com/PumpkinSeed/structs) - Implement simple functions to manipulate structs. Stars:`14`.
* [okrun](https://github.com/xta/okrun) - go run error steamroller. Stars:`14`.
* [filler](https://github.com/yaronsumel/filler) - small utility to fill structs using "fill" tag. Stars:`14`.
* [cmd](https://github.com/SimonBaeumer/cmd) - Library for executing shell commands on osx, windows and linux. Stars:`14`.
* [rerate](https://github.com/abo/rerate) - Redis-based rate counter and rate limiter for Go. Stars:`14`.
* [ghokin](https://github.com/antham/ghokin) - Parallelized formatter with no external dependencies for gherkin (cucumber, behat...). Stars:`13`.
* [slicer](https://github.com/leaanthony/slicer) - Makes working with slices easier. Stars:`12`.
* [rest-go](https://github.com/edermanoel94/rest-go) - A package that provide many helpful methods for working with rest api. Stars:`12`.
* [mimesniffer](https://github.com/aofei/mimesniffer) - A MIME type sniffer for Go. Stars:`11`.
* [ctxutil](https://github.com/posener/ctxutil) - A collection of utility functions for contexts. Stars:`11`.
* [retry](https://github.com/shafreeck/retry) - A pretty simple library to ensure your work to be done. Stars:`10`.
* [backscanner](https://github.com/icza/backscanner) - A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward. Stars:`9`.
* [command](https://github.com/txgruppi/command) - Command pattern for Go with thread safe serial and parallel dispatcher. Stars:`9`.
* [tome](https://github.com/cyruzin/tome) - Tome was designed to paginate simple RESTful APIs. Stars:`9`.
* [limiters](https://github.com/mennanov/limiters) - Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks. Stars:`8`.
* [shutdown](https://github.com/ztrue/shutdown) - App shutdown hooks for `os.Signal` handling. Stars:`7`.
* [retry](https://github.com/percolate/retry) - A simple but highly configurable retry package for Go. Stars:`5`.
* [silk](https://github.com/chrispassas/silk) - Read silk netflow files. Stars:`4`.
* [go-problemdetails](https://github.com/mvmaasakkers/go-problemdetails) - Go package for working with Problem Details. Stars:`4`.
* [blank](https://github.com/Henry-Sarabia/blank) - Verify or remove blanks and whitespace from strings. Stars:`4`.
* [sliceconv](https://github.com/Henry-Sarabia/sliceconv) - Slice conversion between primitive types. Stars:`4`.
* [ptr](https://github.com/gotidy/ptr) - Package that provide functions for simplified creation of pointers from constants of basic types. Stars:`3`.
* [olaf](https://github.com/btnguyen2k/olaf) - Twitter Snowflake implemented in Go. Stars:`1`.
* [sslice](https://github.com/yaa110/sslice) - Create a slice which is always sorted.

## UUID

*Libraries for working with UUIDs.*

* [ulid](https://github.com/oklog/ulid) - Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier). Stars:`1.8K`.
* [uuid](https://github.com/gofrs/uuid) - Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid. Stars:`641`.
* [wuid](https://github.com/edwingeng/wuid) - An extremely fast unique number generator, 10-135 times faster than UUID. Stars:`322`.
* [goid](https://github.com/jakehl/goid) - Generate and Parse RFC4122 compliant V4 UUIDs. Stars:`26`.
* [sno](https://github.com/muyo/sno) - Compact, sortable and fast unique IDs with embedded metadata. Stars:`23`.
* [nanoid](https://github.com/aidarkhanov/nanoid) - A tiny and efficient Go unique string ID generator. Stars:`16`.
* [uuid](https://github.com/agext/uuid) - Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier. Stars:`10`.
* [uniq](https://gitlab.com/skilstak/code/go/uniq) - No hassle safe, fast unique identifiers with commands.

## Validation

*Libraries for validation.*

* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving. Stars:`4.6K`.
* [govalidator](https://github.com/asaskevich/govalidator) - Validators and sanitizers for strings, numerics, slices and structs. Stars:`3.9K`.
* [ozzo-validation](https://github.com/go-ozzo/ozzo-validation) - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags. Stars:`1.2K`.
* [govalidator](https://github.com/thedevsaddam/govalidator) - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation. Stars:`812`.
* [validate](https://github.com/gookit/validate) - Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features. Stars:`192`.
* [checkdigit](https://github.com/osamingo/checkdigit) - Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.). Stars:`48`.
* [jio](https://github.com/faceair/jio) - jio is a json schema validator similar to [joi](https://github.com/hapijs/joi). Stars:`30`.
* [validate](https://github.com/gobuffalo/validate) - This package provides a framework for writing validations for Go applications. Stars:`26`.
* [terraform-validator](https://github.com/thazelart/terraform-validator) - A norms and conventions validator for Terraform. Stars:`25`.

## Version Control

*Libraries for version control.*

* [go-git](https://github.com/src-d/go-git) - highly extensible Git implementation in pure Go. Stars:`4.8K`.
* [git2go](https://github.com/libgit2/git2go) - Go bindings for libgit2. Stars:`1.4K`.
* [hercules](https://github.com/src-d/hercules) - gaining advanced insights from Git repository history. Stars:`808`.
* [gh](https://github.com/rjeczalik/gh) - Scriptable server and net/http middleware for GitHub Webhooks. Stars:`72`.
* [go-vcs](https://github.com/sourcegraph/go-vcs) - manipulate and inspect VCS repositories in Go. Stars:`72`.
* [hgo](https://github.com/beyang/hgo) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories. Stars:`12`.

## Video

*Libraries for manipulating video.*

* [goav](https://github.com/giorgisio/goav) - Comphrensive Go bindings for FFmpeg. Stars:`979`.
* [m3u8](https://github.com/grafov/m3u8) - Parser and generator library of M3U8 playlists for Apple HLS. Stars:`648`.
* [gmf](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries. Stars:`577`.
* [go-astits](https://github.com/asticode/go-astits) - Parse and demux MPEG Transport Streams (.ts) natively in GO. Stars:`284`.
* [go-astisub](https://github.com/asticode/go-astisub) - Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.). Stars:`219`.
* [gst](https://github.com/ziutek/gst) - Go bindings for GStreamer. Stars:`155`.
* [libvlc-go](https://github.com/adrg/libvlc-go) - Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player). Stars:`81`.
* [go-m3u8](https://github.com/quangngotan95/go-m3u8) - Parser and generator library for Apple m3u8 playlists. Stars:`44`.
* [v4l](https://github.com/korandiz/v4l) - Video capture library for Linux, written in Go. Stars:`36`.
* [libgosubs](https://github.com/wargarblgarbl/libgosubs) - Subtitle format support for go. Supports .srt, .ttml, and .ass. Stars:`12`.

## Web Frameworks

*Full stack web frameworks.*

* [Gin](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity. Stars:`34.6K`.
* [Beego](https://github.com/astaxie/beego) - beego is an open-source, high-performance web framework for the Go programming language. Stars:`23.0K`.
* [Echo](https://github.com/labstack/echo) - High performance, minimalist Go web framework. Stars:`16.2K`.
* [Revel](https://github.com/revel/revel) - High-productivity web framework for the Go language. Stars:`11.5K`.
* [Goa](https://github.com/goadesign/goa) - Goa provides a holistic approach for developing remote APIs and microservices in Go. Stars:`3.7K`.
* [go-json-rest](https://github.com/ant0ine/go-json-rest) - Quick and easy way to setup a RESTful JSON API. Stars:`3.4K`.
* [Gizmo](https://github.com/NYTimes/gizmo) - Microservice toolkit used by the New York Times. Stars:`3.0K`.
* [Macaron](https://github.com/go-macaron/macaron) - Macaron is a high productive and modular design web framework in Go. Stars:`2.9K`.
* [utron](https://github.com/gernest/utron) - Lightweight MVC framework for Go(Golang). Stars:`2.2K`.
* [tigertonic](https://github.com/rcrowley/go-tigertonic) - Go framework for building JSON web services inspired by Dropwizard. Stars:`999`.
* [tango](https://github.com/lunny/tango) - Micro & pluggable web framework for Go. Stars:`832`.
* [gongular](https://github.com/mustafaakin/gongular) - Fast Go web framework with input mapping/validation and (DI) Dependency Injection. Stars:`421`.
* [neo](https://github.com/ivpusic/neo) - Neo is minimal and fast Go Web Framework with extremely simple API. Stars:`401`.
* [Air](https://github.com/aofei/air) - An ideally refined web framework for Go. Stars:`362`.
* [mango](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333. Stars:`343`.
* [Gondola](https://github.com/rainycape/gondola) - The web framework for writing faster sites, faster. Stars:`314`.
* [Golf](https://github.com/dinever/golf) - Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library. Stars:`239`.
* [Aero](https://github.com/aerogo/aero) - High-performance web framework for Go, reaches top scores in Lighthouse. Stars:`226`.
* [go-rest](https://github.com/ungerik/go-rest) - Small and evil REST framework for Go. Stars:`116`.
* [hiboot](https://github.com/hidevopsio/hiboot) - hiboot is a high performance web application framework with auto configuration and dependency injection support. Stars:`109`.
* [WebGo](https://github.com/bnkamalesh/webgo) - A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc). Stars:`84`.
* [Flamingo](https://github.com/i-love-flamingo/flamingo) - Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc. Stars:`75`.
* [Golax](https://github.com/fulldump/golax) - A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more. Stars:`71`.
* [uAdmin](https://github.com/uadmin/uadmin) - Fully featured web framework for Golang, inspired by Django. Stars:`69`.
* [Microservice](https://github.com/claygod/microservice) - The framework for the creation of microservices, written in Golang. Stars:`69`.
* [YARF](https://github.com/yarf-framework/yarf) - Fast micro-framework designed to build REST APIs and web services in a fast and simple way. Stars:`54`.
* [Fireball](https://github.com/zpatrick/fireball) - More "natural" feeling web framework. Stars:`50`.
* [patron](https://github.com/beatlabs/patron) - Patron is a microservice framework following best cloud practices with a focus on productivity. Stars:`46`.
* [vox](https://github.com/aisk/vox) - A golang web framework for humans, inspired by Koa heavily. Stars:`46`.
* [Ginrpc](https://github.com/xxjwxc/ginrpc) - Gin parameter automatic binding tool,gin rpc tools. Stars:`45`.
* [Flamingo Commerce](https://github.com/i-love-flamingo/flamingo-commerce) - Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications. Stars:`39`.
* [Resoursea](https://github.com/resoursea/api) - REST framework for quickly writing resource based services. Stars:`30`.
* [rex](https://github.com/goanywhere/rex) - Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`. Stars:`29`.
* [goa](https://github.com/goa-go/goa) - goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware. Stars:`27`.
* [rux](https://github.com/gookit/rux) - Simple and fast web framework for build golang HTTP applications. Stars:`18`.
* [Banjo](https://github.com/nsheremet/banjo) - Very simple and fast web framework for Go. Stars:`11`.
* [REST Layer](http://rest-layer.io) - Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
* [Buffalo](http://gobuffalo.io) - Bringing the productivity of Rails to Go!
* [aah](https://aahframework.org) - Scalable, performant, rapid development Web framework for Go.

### Middlewares

#### Actual middlewares

* [Tollbooth](https://github.com/didip/tollbooth) - Rate limit HTTP request handler. Stars:`1.4K`.
* [CORS](https://github.com/rs/cors) - Easily add CORS capabilities to your API. Stars:`1.4K`.
* [Limiter](https://github.com/ulule/limiter) - Dead simple rate limit middleware for Go. Stars:`860`.
* [go-server-timing](https://github.com/mitchellh/go-server-timing) - Add/parse Server-Timing header. Stars:`767`.
* [ln-paywall](https://github.com/philippgille/ln-paywall) - Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin). Stars:`98`.
* [XFF](https://github.com/sebest/xff) - Handle `X-Forwarded-For` header and friends. Stars:`72`.
* [formjson](https://github.com/rs/formjson) - Transparently handle JSON input as a standard form POST. Stars:`34`.
* [client-timing](https://github.com/posener/client-timing) - An HTTP client for Server-Timing header. Stars:`16`.

#### Libraries for creating HTTP middlewares

* [negroni](https://github.com/urfave/negroni) - Idiomatic HTTP middleware for Golang. Stars:`6.5K`.
* [alice](https://github.com/justinas/alice) - Painless middleware chaining for Go. Stars:`1.9K`.
* [render](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, and HTML template responses. Stars:`1.3K`.
* [stats](https://github.com/thoas/stats) - Go middleware that stores various information about your web application. Stars:`550`.
* [interpose](https://github.com/carbocation/interpose) - Minimalist net/http middleware for golang. Stars:`289`.
* [muxchain](https://github.com/stephens2424/muxchain) - Lightweight middleware for net/http. Stars:`209`.
* [renderer](https://github.com/thedevsaddam/renderer) - Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go. Stars:`178`.
* [rye](https://github.com/InVisionApp/rye) - Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context. Stars:`93`.
* [gores](https://github.com/alioygur/gores) - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs. Stars:`88`.
* [chain](https://github.com/codemodus/chain) - Handler wrapper chaining with scoped data (net/context-based "middleware"). Stars:`64`.
* [go-wrap](https://github.com/go-on/wrap) - Small middlewares package for net/http. Stars:`59`.
* [catena](https://github.com/codemodus/catena) - http.Handler wrapper catenation (same API as "chain"). Stars:`7`.

### Routers

* [mux](https://github.com/gorilla/mux) - Powerful URL router and dispatcher for golang. Stars:`10.9K`.
* [httprouter](https://github.com/julienschmidt/httprouter) - High performance router. Use this and the standard http handlers to form a very high performance web framework. Stars:`10.6K`.
* [chi](https://github.com/go-chi/chi) - Small, fast and expressive HTTP router built on net/context. Stars:`6.8K`.
* [gocraft/web](https://github.com/gocraft/web) - Mux and middleware package in Go. Stars:`1.4K`.
* [Bone](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer. Stars:`1.2K`.
* [fasthttprouter](https://github.com/buaazp/fasthttprouter) - High performance router forked from `httprouter`. The first router fit for `fasthttp`. Stars:`805`.
* [Goji](https://github.com/goji/goji) - Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`. Stars:`790`.
* [xujiajun/gorouter](https://github.com/xujiajun/gorouter) - A simple and fast HTTP router for Go. Stars:`473`.
* [httptreemux](https://github.com/dimfeld/httptreemux) - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter. Stars:`404`.
* [lars](https://github.com/go-playground/lars) - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks. Stars:`379`.
* [ozzo-routing](https://github.com/go-ozzo/ozzo-routing) - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs. Stars:`368`.
* [Siesta](https://github.com/VividCortex/siesta) - Composable framework to write middleware and handlers. Stars:`352`.
* [vestigo](https://github.com/husobee/vestigo) - Performant, stand-alone, HTTP compliant URL Router for go web applications. Stars:`258`.
* [gowww/router](https://github.com/gowww/router) - Lightning fast HTTP router fully compatible with the net/http.Handler interface. Stars:`159`.
* [alien](https://github.com/gernest/alien) - Lightweight and fast http router from outer space. Stars:`109`.
* [violetear](https://github.com/nbari/violetear) - Go HTTP router. Stars:`96`.
* [Bxog](https://github.com/claygod/Bxog) - Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters. Stars:`92`.
* [pure](https://github.com/go-playground/pure) - Is a lightweight HTTP router that sticks to the std "net/http" implementation. Stars:`91`.
* [xmux](https://github.com/rs/xmux) - High performance muxer based on `httprouter` with `net/context` support. Stars:`87`.
* [GoRouter](https://github.com/vardius/gorouter) - GoRouter is a Server/API micro framwework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`. Stars:`58`.
* [bellt](https://github.com/GuilhermeCaruso/bellt) - A simple Go HTTP router. Stars:`40`.
* [FastRouter](https://github.com/razonyang/fastrouter) - a fast, flexible HTTP router written in Go. Stars:`19`.
* [goroute](https://github.com/goroute/route) - Simple yet powerful HTTP request multiplexer. Stars:`6`.

## Windows

* [go-ole](https://github.com/go-ole/go-ole) - Win32 OLE implementation for golang. Stars:`597`.
* [d3d9](https://github.com/gonutz/d3d9) - Go bindings for Direct3D9. Stars:`96`.
* [gosddl](https://github.com/MonaxGT/gosddl) - Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL. Stars:`1`.

## XML

*Libraries and tools for manipulating XML.*

* [zek](https://github.com/miku/zek) - Generate a Go struct from XML. Stars:`314`.
* [xpath](https://github.com/antchfx/xpath) - XPath package for Go. Stars:`212`.
* [xquery](https://github.com/antchfx/xquery) - XQuery lets you extract data from HTML/XML documents using XPath expression. Stars:`152`.
* [xml2map](https://github.com/sbabiv/xml2map) - XML to MAP converter written Golang. Stars:`19`.
* [XML-Comp](https://github.com/xml-comp/xml-comp) - Simple command line XML comparer that generates diffs of folders, files and tags. Stars:`16`.
* [xmlwriter](https://github.com/shabbyrobe/xmlwriter) - Procedural XML generation API based on libxml2's xmlwriter module. Stars:`7`.

# Tools

*Go software and plugins.*

## Code Analysis

* [GoLint](https://github.com/golang/lint) - Golint is a linter for Go source code. Stars:`3.3K`.
* [errcheck](https://github.com/kisielk/errcheck) - Errcheck is a program for checking for unchecked errors in Go programs. Stars:`1.4K`.
* [gcvis](https://github.com/davecheney/gcvis) - Visualise Go program GC trace data in real time. Stars:`953`.
* [php-parser](https://github.com/z7zmey/php-parser) - A Parser for PHP written in Go. Stars:`687`.
* [go-critic](https://github.com/go-critic/go-critic) - source code linter that brings checks that are currently not implemented in other linters. Stars:`652`.
* [goast-viewer](https://github.com/yuroyoro/goast-viewer) - Web based Golang AST visualizer. Stars:`416`.
* [go-cleanarch](https://github.com/roblaszczak/go-cleanarch) - go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects. Stars:`302`.
* [unconvert](https://github.com/mdempsky/unconvert) - Remove unnecessary type conversions from Go source. Stars:`274`.
* [go-mod-outdated](https://github.com/psampaz/go-mod-outdated) - An easy way to find outdated dependencies of your Go projects. Stars:`251`.
* [gostatus](https://github.com/shurcooL/gostatus) - Command line tool, shows the status of repositories that contain Go packages. Stars:`242`.
* [dupl](https://github.com/mibk/dupl) - Tool for code clone detection. Stars:`186`.
* [apicompat](https://github.com/bradleyfalzon/apicompat) - Checks recent changes to a Go project for backwards incompatible changes. Stars:`170`.
* [go-checkstyle](https://github.com/qiniu/checkstyle) - checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments. Stars:`103`.
* [tickgit](https://github.com/augmentable-dev/tickgit) - CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author. Stars:`98`.
* [GoPlantUML](https://github.com/jfeliu007/goplantuml) - Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them. Stars:`83`.
* [lint](https://github.com/surullabs/lint) - Run linters as part of go test. Stars:`64`.
* [validate](https://github.com/mccoyst/validate) - Automatically validates struct fields with tags. Stars:`62`.
* [go-outdated](https://github.com/firstrow/go-outdated) - Console application that displays outdated packages. Stars:`45`.
* [golines](https://github.com/segmentio/golines) - Formatter that automatically shortens long lines in Go code. Stars:`28`.
* [tarp](https://github.com/verygoodsoftwarenotvirus/tarp) - tarp finds functions and methods without direct unit tests in Go source code. Stars:`14`.
* [gosimple](https://github.com/dominikh/go-tools/tree/master/cmd/gosimple) - gosimple is a linter for Go source code that specialises on simplifying code.
* [Golint online](http://go-lint.appspot.com/) - Lints online Go source files on GitHub, Bitbucket and Google Project Hosting using the golint package.
* [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
* [goreturns](https://sourcegraph.com/github.com/sqs/goreturns) - Adds zero-value return statements to match the func return types.
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
* [GoCover.io](http://gocover.io/) - GoCover.io offers the code coverage of any golang package as a service.
* [unused](https://github.com/dominikh/go-tools/tree/master/cmd/unused) - unused checks Go code for unused constants, variables, functions and types.
* [GolangCI](https://golangci.com/) - GolangCI is an automated Golang code review service for GitHub pull requests. Service is open source and it's free for open source projects.

## Editor Plugins

* [vim-go](https://github.com/fatih/vim-go) - Go development plugin for Vim. Stars:`11.5K`.
* [vscode-go](https://github.com/Microsoft/vscode-go) - Extension for Visual Studio Code (VS Code) which provides support for the Go language. Stars:`5.6K`.
* [gocode](https://github.com/nsf/gocode) - Autocompletion daemon for the Go programming language. Stars:`4.8K`.
* [GoSublime](https://github.com/DisposaBoy/GoSublime) - Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features. Stars:`3.3K`.
* [go-plus](https://github.com/joefitzgerald/go-plus) - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting. Stars:`1.5K`.
* [go-mode](https://github.com/dominikh/go-mode.el) - Go mode for GNU/Emacs. Stars:`1.0K`.
* [Watch](https://github.com/eaburns/Watch) - Runs a command in an acme win on file changes. Stars:`171`.
* [vim-compiler-go](https://github.com/rjohnsondev/vim-compiler-go) - Vim plugin to highlight syntax errors on save. Stars:`81`.
* [go-language-server](https://github.com/theia-ide/go-language-server) - A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol. Stars:`33`.
* [gounit-vim](https://github.com/hexdigest/gounit-vim) - Vim plugin for generating Go tests based on the function's or method's signature. Stars:`17`.
* [theia-go-extension](https://github.com/theia-ide/theia-go-extension) - Go language support for the Theia IDE. Stars:`13`.
* [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof) - This extension adds benchmark profiling support for the Go language to VS Code.
* [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) - Go plugin for JetBrains IDEs.

## Go Generate Tools

* [gotests](https://github.com/cweill/gotests) - Generate Go tests from your source code. Stars:`2.5K`.
* [genny](https://github.com/cheekybits/genny) - Elegant generics for Go. Stars:`1.1K`.
* [re2dfa](https://github.com/opennota/re2dfa) - Transform regular expressions into finite state machines and output Go source code. Stars:`176`.
* [gocontracts](https://github.com/Parquery/gocontracts) - brings design-by-contract to Go by synchronizing the code with the documentation. Stars:`55`.
* [gounit](https://github.com/hexdigest/gounit) - Generate Go tests using your own templates. Stars:`35`.
* [generic](https://github.com/usk81/generic) - flexible data type for Go. Stars:`32`.
* [hasgo](https://github.com/DylanMeeus/hasgo) - Generate Haskell inspired functions for your slices. Stars:`31`.
* [gonerics](http://github.com/bouk/gonerics) - Idiomatic Generics in Go.
* [TOML-to-Go](https://xuri.me/toml-to-go) - Translates TOML into a Go type in the browser instantly.

## Go Tools

* [go-swagger](https://github.com/go-swagger/go-swagger) - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API. Stars:`4.7K`.
* [OctoLinker](https://github.com/OctoLinker/browser-extension) - Navigate through go files efficiently with the OctoLinker browser extension for GitHub. Stars:`4.1K`.
* [go-callvis](https://github.com/TrueFurby/go-callvis) - Visualize call graph of your Go program using dot format. Stars:`2.3K`.
* [depth](https://github.com/KyleBanks/depth) - Visualize dependency trees of any package by analyzing imports. Stars:`431`.
* [richgo](https://github.com/kyoh86/richgo) - Enrich `go test` outputs with text decorations. Stars:`427`.
* [rts](https://github.com/galeone/rts) - RTS: response to struct. Generates Go structs from server responses. Stars:`193`.
* [godbg](https://github.com/tylerwince/godbg) - Implementation of Rusts `dbg!` macro for quick and easy debugging during development. Stars:`163`.
* [colorgo](https://github.com/songgao/colorgo) - Wrapper around `go` command for colorized `go build` output. Stars:`101`.
* [gothanks](https://github.com/psampaz/gothanks) - GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers. Stars:`83`.
* [go-pkg-complete](https://github.com/skelterjohn/go-pkg-complete) - Bash completion for go and wgo. Stars:`37`.
* [generator-go-lang](https://github.com/axelspringer/generator-go-lang) - A [Yeoman](http://yeoman.io) generator to get new Go projects started. Stars:`17`.
* [go-james](https://github.com/pieterclaerhout/go-james) - Go project skeleton creator, builds and tests your projects without the manual setup. Stars:`7`.
* [gilbert](https://go-gilbert.github.io) - Build system and task runner for Go projects.
* [gb](https://getgb.io/) - An easy to use project based build tool for the Go programming language.

## Software Packages

*Software written in Go.*

### DevOps Tools

* [kubernetes](https://github.com/kubernetes/kubernetes) - Container Cluster Manager from Google. Stars:`62.2K`.
* [Moby](https://github.com/moby/moby) - Collaborative project for the container ecosystem to assemble container-based systems. Stars:`56.1K`.
* [traefik](https://github.com/containous/traefik) - Reverse proxy and load balancer with support for multiple backends. Stars:`26.9K`.
* [Gitea](https://github.com/go-gitea/gitea) - Fork of Gogs, entirely community driven. Stars:`17.8K`.
* [Vegeta](https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000! Stars:`13.6K`.
* [Packer](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration. Stars:`9.7K`.
* [Hey](https://github.com/rakyll/hey) - Hey is a tiny program that sends some load to a web application. Stars:`7.4K`.
* [GVM](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions. Stars:`4.9K`.
* [webhook](https://github.com/adnanh/webhook) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server. Stars:`4.8K`.
* [gaia](https://github.com/gaia-pipeline/gaia) - Build powerful pipelines in any programming language. Stars:`4.0K`.
* [gox](https://github.com/mitchellh/gox) - Dead simple, no frills Go cross compile tool. Stars:`3.6K`.
* [bosun](https://github.com/bosun-monitor/bosun) - Time Series Alerting Framework. Stars:`2.9K`.
* [bombardier](https://github.com/codesenberg/bombardier) - Fast cross-platform HTTP benchmarking tool. Stars:`1.9K`.
* [goxc](https://github.com/laher/goxc) - build tool for Go, with a focus on cross-compiling and packaging. Stars:`1.6K`.
* [fac](https://github.com/mkchoi212/fac) - Command-line user interface to fix git merge conflicts. Stars:`1.6K`.
* [kala](https://github.com/ajvb/kala) - Simplistic, modern, and performant job scheduler. Stars:`1.4K`.
* [StatusOK](https://github.com/sanathp/statusok) - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected. Stars:`1.3K`.
* [script](https://github.com/bitfield/script) - Making it easy to write shell-like scripts in Go for DevOps and system administration tasks. Stars:`1.3K`.
* [s3gof3r](https://github.com/rlmcpherson/s3gof3r) - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3. Stars:`1.0K`.
* [Pomerium](https://github.com/pomerium/pomerium) - Pomerium is an identity-aware access proxy. Stars:`925`.
* [go-selfupdate](https://github.com/sanbornm/go-selfupdate) - Enable your Go applications to self update. Stars:`704`.
* [skm](https://github.com/TimothyYe/skm) - SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily! Stars:`575`.
* [Scaleway-cli](https://github.com/scaleway/scaleway-cli) - Manage BareMetal Servers from Command Line (as easily as with Docker). Stars:`559`.
* [aurora](https://github.com/xuri/aurora) - Cross-platform web-based Beanstalkd queue server console. Stars:`437`.
* [govvv](https://github.com/ahmetalpbalkan/govvv) - “go build” wrapper to easily add version information into Go binaries. Stars:`421`.
* [gonative](https://github.com/inconshreveable/gonative) - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages. Stars:`314`.
* [Mora](https://github.com/emicklei/mora) - REST server for accessing MongoDB documents and meta data. Stars:`268`.
* [lstags](https://github.com/ivanilves/lstags) - Tool and API to sync Docker images across different registries. Stars:`241`.
* [dogo](https://github.com/liudng/dogo) - Monitoring changes in the source file and automatically compile and run (restart). Stars:`222`.
* [godbg](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application. Stars:`219`.
* [Pewpew](https://github.com/bengadbois/pewpew) - Flexible HTTP command line stress tester. Stars:`216`.
* [manssh](https://github.com/xwjdsh/manssh) - manssh is a command line tool for managing your ssh alias config easily. Stars:`209`.
* [gobrew](https://github.com/cryptojuice/gobrew) - gobrew lets you easily switch between multiple versions of go. Stars:`176`.
* [Blast](https://github.com/dave/blast) - A simple tool for API load testing and batch jobs. Stars:`174`.
* [ostent](https://github.com/ostrost/ostent) - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB. Stars:`164`.
* [grapes](https://github.com/yaronsumel/grapes) - Lightweight tool designed to distribute commands over ssh with ease. Stars:`143`.
* [uTask](https://github.com/ovh/utask) - Automation engine that models and executes business processes declared in yaml. Stars:`122`.
* [easyssh-proxy](https://github.com/appleboy/easyssh-proxy) - Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`. Stars:`119`.
* [jcli](https://github.com/jenkins-zh/jenkins-cli) - Jenkins CLI allows you manage your Jenkins as an easy way. Stars:`113`.
* [kcli](https://github.com/cswank/kcli) - Command line tool for inspecting kafka topics/partitions/messages. Stars:`102`.
* [terraform-provider-openapi](https://github.com/dikhan/terraform-provider-openapi) - Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed. Stars:`90`.
* [winrm-cli](https://github.com/masterzen/winrm-cli) - Cli tool to remotely execute commands on Windows machines. Stars:`82`.
* [go-furnace](https://github.com/go-furnace/go-furnace) - Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean. Stars:`73`.
* [drone-scp](https://github.com/appleboy/drone-scp) - Copy files and artifacts via SSH using a binary, docker or Drone CI. Stars:`66`.
* [Dockerfile-Generator](https://github.com/ozankasikci/dockerfile-generator) - A go library and an executable that produces valid Dockerfiles using various input channels. Stars:`60`.
* [Dropship](https://github.com/chrismckenzie/dropship) - Tool for deploying code via cdn. Stars:`47`.
* [Rodent](https://github.com/alouche/rodent) - Rodent helps you manage Go versions, projects and track dependencies. Stars:`30`.
* [drone-jenkins](https://github.com/appleboy/drone-jenkins) - Trigger downstream Jenkins jobs using a binary, docker or Drone CI. Stars:`26`.
* [awsenv](https://github.com/soniah/awsenv) - Small binary that loads Amazon (AWS) environment variables for a profile. Stars:`22`.
* [lwc](https://github.com/timdp/lwc) - A live-updating version of the UNIX wc command. Stars:`10`.
* [DepCharge](https://github.com/centerorbit/depcharge) - Helps orchestrating the execution of commands across the many dependencies in larger projects. Stars:`9`.
* [sg](https://github.com/ChristopherRabotin/sg) - Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response. Stars:`5`.
* [aptly](https://github.com/smira/aptly) - aptly is a Debian repository management tool. Stars:`1`.
* [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
* [Wide](https://wide.b3log.org/login) - Web-based IDE for Teams using Golang.
* [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator) - Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.

### Other Software

* [Gor](https://github.com/buger/gor) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time. Stars:`12.1K`.
* [restic](https://github.com/restic/restic) - De-duplicating backup program. Stars:`9.1K`.
* [Seaweed File System](https://github.com/chrislusf/seaweedfs) - Fast, Simple and Scalable Distributed File System with O(1) disk seek. Stars:`9.0K`.
* [confd](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul. Stars:`6.8K`.
* [Comcast](https://github.com/tylertreat/Comcast) - Simulate bad network connections. Stars:`6.4K`.
* [LiteIDE](https://github.com/visualfc/liteide) - LiteIDE is a simple, open source, cross-platform Go IDE. Stars:`5.8K`.
* [drive](https://github.com/odeke-em/drive) - Google Drive client for the commandline. Stars:`5.2K`.
* [nes](https://github.com/fogleman/nes) - Nintendo Entertainment System (NES) emulator written in Go. Stars:`4.3K`.
* [toxiproxy](https://github.com/shopify/toxiproxy) - Proxy to simulate network and system conditions for automated tests. Stars:`4.3K`.
* [Duplicacy](https://github.com/gilbertchen/duplicacy) - A cross-platform network and cloud backup tool based on the idea of lock-free deduplication. Stars:`3.1K`.
* [croc](https://github.com/schollz/croc) - Easily and securely send files or folders from one computer to another. Stars:`2.8K`.
* [myLG](https://github.com/mehrdadrad/mylg) - Command Line Network Diagnostic tool written in Go. Stars:`2.3K`.
* [GoBoy](https://github.com/Humpheh/goboy) - Nintendo Game Boy Color emulator written in Go. Stars:`2.2K`.
* [Stack Up](https://github.com/pressly/sup) - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers. Stars:`2.1K`.
* [lgo](https://github.com/yunabe/lgo) - Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility. Stars:`1.9K`.
* [Circuit](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications. Stars:`1.8K`.
* [snap](https://github.com/intelsdi-x/snap) - Powerful telemetry framework. Stars:`1.8K`.
* [borg](https://github.com/crufter/borg) - Terminal based search engine for bash snippets. Stars:`1.5K`.
* [scc](https://github.com/boyter/scc) - Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates. Stars:`1.2K`.
* [Documize](https://github.com/documize/community) - Modern wiki software that integrates data from SaaS tools. Stars:`965`.
* [Go Package Store](https://github.com/shurcooL/Go-Package-Store) - App that displays updates for the Go packages in your GOPATH. Stars:`883`.
* [peg](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator. Stars:`663`.
* [Leaps](https://github.com/jeffail/leaps) - Pair programming service using Operational Transforms. Stars:`659`.
* [vFlow](https://github.com/VerizonDigital/vflow) - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector. Stars:`640`.
* [gfile](https://github.com/Antonito/gfile) - Securely transfer files between two computers, without any third party, over WebRTC. Stars:`523`.
* [shell2http](https://github.com/msoap/shell2http) - Executing shell commands via http server (for prototyping or remote control). Stars:`490`.
* [mockingjay](https://github.com/quii/mockingjay-server) - Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests. Stars:`437`.
* [go-peerflix](https://github.com/Sioro-Neoku/go-peerflix) - Video streaming torrent client. Stars:`390`.
* [gocc](https://github.com/goccmack/gocc) - Gocc is a compiler kit for Go written in Go. Stars:`373`.
* [wellington](https://github.com/wellington/wellington) - Sass project management tool, extends the language with sprite functions (like Compass). Stars:`292`.
* [ipe](https://github.com/dimiro1/ipe) - Open source Pusher server implementation compatible with Pusher client libraries written in GO. Stars:`289`.
* [ide](https://github.com/thestrukture/ide) - Browser accessible IDE. Designed for Go with Go. Stars:`255`.
* [Cherry](https://github.com/rafael-santiago/cherry) - Tiny webchat server in Go. Stars:`205`.
* [orange-cat](https://github.com/noraesae/orange-cat) - Markdown previewer written in Go. Stars:`182`.
* [Orbit](https://github.com/gulien/orbit) - A simple tool for running commands and generating files from templates. Stars:`133`.
* [joincap](https://github.com/assafmo/joincap) - Command-line utility for merging multiple pcap files together. Stars:`128`.
* [boxed](https://github.com/tejo/boxed) - Dropbox based blog engine. Stars:`72`.
* [dp](https://github.com/scryinfo/dp) - Through SDK for data exchange with blockchain, developers can get easy access to DAPP development. Stars:`57`.
* [naclpipe](https://github.com/unix4fun/naclpipe) - Simple NaCL EC25519 based crypto pipe tool written in Go. Stars:`20`.
* [term-quiz](https://github.com/crazcalm/term-quiz) - Quizzes for your terminal. Stars:`17`.
* [Snitch](https://github.com/lucasgomide/snitch) - Simple way to notify your team and many tools when someone has deployed any application via Tsuru. Stars:`15`.
* [GoDocTooltip](https://github.com/diankong/GoDocTooltip) - Chrome extension for Go Doc sites, which shows function description as tooltip at function list. Stars:`11`.
* [Docker](http://www.docker.com/) - Open platform for distributed applications for developers and sysadmins.
* [syncthing](https://syncthing.net/) - Open, decentralized file synchronization tool and protocol.
* [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
* [limetext](http://limetext.org/) - Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
* [tsuru](https://tsuru.io/) - Extensible and open source Platform as a Service software.
* [hugo](http://gohugo.io/) - Fast and Modern Static Website Engine.
* [GoLand](https://jetbrains.com/go) - Full featured cross-platform Go IDE.

# Resources

*Where to discover new Go libraries.*

## Benchmarks

* [go-http-routing-benchmark](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router benchmark and comparison. Stars:`1.3K`.
* [go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark) - Go web framework benchmark. Stars:`1.1K`.
* [skynet](https://github.com/atemerev/skynet) - Skynet 1M threads microbenchmark. Stars:`946`.
* [go_serialization_benchmarks](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods. Stars:`942`.
* [speedtest-resize](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language. Stars:`186`.
* [go-benchmarks](https://github.com/tylertreat/go-benchmarks) - Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches. Stars:`130`.
* [gospeed](https://github.com/feyeleanor/GoSpeed) - Go micro-benchmarks for calculating the speed of language constructs. Stars:`99`.
* [autobench](https://github.com/davecheney/autobench) - Framework to compare the performance between different Go versions. Stars:`90`.
* [gocostmodel](https://github.com/PuerkitoBio/gocostmodel) - Benchmarks of common basic operations for the Go language. Stars:`55`.
* [golang-sql-benchmark](https://github.com/tyler-smith/golang-sql-benchmark) - Collection of benchmarks for popular Go database/SQL utilities. Stars:`50`.
* [go-benchmark-app](https://github.com/mrLSD/go-benchmark-app) - Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results. Stars:`20`.
* [kvbench](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark. Stars:`16`.

## Conferences

* [Capital Go](http://www.capitalgolang.com) - Washington, D.C., USA.
* [dotGo](http://www.dotgo.eu) - Paris, France.
* [GoCon](http://gocon.connpass.com/) - Tokyo, Japan.
* [GoDays](https://www.godays.io/) - Berlin, Germany.
* [GoLab](http://golab.io/) - Florence, Italy.
* [GolangUK](http://golanguk.com/) - London, UK.
* [GopherChina](http://gopherchina.org) - Shanghai, China.
* [GopherCon](http://www.gophercon.com/) - Denver, USA.
* [GopherCon Australia](https://gophercon.com.au/) - Sydney, Australia.
* [GopherCon Brazil](https://gopherconbr.org) - Florianópolis, BR.
* [GopherCon Europe](https://gophercon.is/) - Berlin, Germany.
* [GopherCon India](https://www.gophercon.in/) - Pune, India.
* [GopherCon Israel](https://www.gophercon.org.il/) - Tel Aviv, Israel.
* [GopherCon Russia](https://www.gophercon-russia.ru) - Moscow, Russia.
* [GopherCon Singapore](https://gophercon.sg) - Mapletree Business City, Singapore.
* [GopherCon Vietnam](https://gophercon.vn/) - Ho Chi Minh City, Vietnam.
* [GothamGo](http://gothamgo.com/) - New York City, USA.
* [GoWayFest](https://goway.io/) - Minsk, Belarus.

## E-Books

* [GoBooks](https://github.com/dariubs/GoBooks) - A curated list of Go books. Stars:`7.4K`.
* [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example) Stars:`6.3K`.
* [Go Succinctly](https://github.com/thedevsir/gosuccinctly) - in Persian. Stars:`13`.
* [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
* [Go 101](https://go101.org) - A book focusing on Go syntax/semantics and all kinds of details.
* [Go Bootcamp](http://golangbootcamp.com)
* [Building Web Apps With Go](https://www.gitbook.com/book/codegangsta/building-web-apps-with-go/details)
* [Build Web Application with Golang](https://www.gitbook.com/book/astaxie/build-web-application-with-golang/details)
* [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
* [Network Programming With Go](https://jan.newmarch.name/go/)
* [The Go Programming Language](http://www.gopl.io/)
* [An Introduction to Programming in Go](http://www.golang-book.com/)
* [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)
* [Writing A Compiler In Go](https://compilerbook.com)
* [Writing An Interpreter In Go](https://interpreterbook.com)

## Gophers

* [gophers](https://github.com/ashleymcnamara/gophers) - Gopher artworks by Ashley McNamara. Stars:`2.0K`.
* [gophers](https://github.com/egonelbre/gophers) - Free gophers. Stars:`1.8K`.
* [Free Gophers Pack](https://github.com/MariaLetta/free-gophers-pack) - Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster. Stars:`1.8K`.
* [gophericons](https://github.com/shalakhin/gophericons) Stars:`562`.
* [gopher-stickers](https://github.com/tenntenn/gopher-stickers) Stars:`462`.
* [gopherize.me](https://github.com/matryer/gopherize.me) - Gopherize yourself. Stars:`356`.
* [gopher-vector](https://github.com/golang-samples/gopher-vector) Stars:`350`.
* [gopher-logos](https://github.com/GolangUA/gopher-logos) - adorable gopher logos. Stars:`69`.
* [gophers](https://github.com/rogeralsing/gophers) - random gopher graphics. Stars:`50`.
* [gophers](https://github.com/sillecelik/go-gopher) - Gopher amigurumi toy pattern. Stars:`42`.
* [Go-gopher-Vector](https://github.com/keygx/Go-gopher-Vector) - Go gopher Vector Data [.ai, .svg]. Stars:`30`.

## Meetups

* [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
* [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
* [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
* [Canberra Gophers](https://www.meetup.com/Canberra-Gophers/)
* [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
* [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
* [Go Toronto](https://www.meetup.com/go-toronto/)
* [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
* [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
* [GoJakarta](https://www.meetup.com/GoJakarta/)
* [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
* [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
* [Golang Baltimore, MD](https://www.meetup.com/BaltimoreGolang/)
* [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
* [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
* [Golang Boston](https://www.meetup.com/bostongo/)
* [Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)
* [Golang Cardiff, UK](https://www.meetup.com/Cardiff-Go-Meetup/)
* [Golang Copenhagen](https://www.meetup.com/Go-Cph/)
* [Golang Curitiba - Brazil](https://www.meetup.com/GolangCWB/)
* [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
* [Golang Dorset, UK](https://www.meetup.com/golang-dorset/)
* [Golang Gurgaon, India](https://www.meetup.com/Gurgaon-Go-Meetup/)
* [Golang Hamburg - Germany](https://www.meetup.com/Go-User-Group-Hamburg/)
* [Golang Israel](https://www.meetup.com/Go-Israel/)
* [Golang Joinville - Brazil](https://www.meetup.com/Joinville-Go-Meetup/)
* [Golang Korea](https://www.meetup.com/GDG-Golang-Korea/)
* [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
* [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
* [Golang Marseille](https://www.meetup.com/fr-FR/Golang-Marseille/)
* [Golang Melbourne](https://www.meetup.com/golang-mel/)
* [Golang Mountain View](https://www.meetup.com/Golang-Mountain-View/)
* [Golang New York](https://www.meetup.com/nycgolang/)
* [Golang Paris](https://www.meetup.com/Golang-Paris/)
* [Golang Pune](https://www.meetup.com/Golang-Pune/)
* [Golang Singapore](https://www.meetup.com/golangsg/)
* [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
* [Golang Sydney, AU](https://www.meetup.com/golang-syd/)
* [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
* [Golang Taipei](https://www.meetup.com/golang-taipei-meetup/)
* [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
* [Golang Казань](https://www.meetup.com/GolangKazan/)
* [Golang Москва](https://www.meetup.com/Golang-Moscow/)
* [Golang Питер](https://www.meetup.com/Golang-Peter/)
* [GoSF - San Francisco, CA](https://www.meetup.com/golangsf)
* [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
* [Seattle Go Programmers](https://www.meetup.com/golang/)
* [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
* [Utah Go User Group](https://www.meetup.com/utahgophers/)
* [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)

*Add the group of your city/country here (send **PR**)*

## Twitter

* [@golang](https://twitter.com/golang)
* [@golang_news](https://twitter.com/golang_news)
* [@golangch](https://twitter.com/golangch)
* [@golangflow](https://twitter.com/golangflow)
* [@golangweekly](https://twitter.com/golangweekly)

## Websites

* [awesome-awesomeness](https://github.com/bayandin/awesome-awesomeness) - List of other amazingly awesome lists. Stars:`25.5K`.
* [Awesome Remote Job](https://github.com/lukasz-madon/awesome-remote-job) - Curated list of awesome remote jobs. A lot of them are looking for Go hackers. Stars:`15.5K`.
* [golang-graphics](https://github.com/mholt/golang-graphics) - Collection of Go images, graphics, and art. Stars:`142`.
* [gocryforhelp](https://github.com/ninedraft/gocryforhelp) - Collection of Go projects that needs help. Good place to start your open-source way in Go. Stars:`38`.
* [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
* [Go Challenge](http://golang-challenge.org/) - Learn Go by solving problems and getting feedback from Go experts.
* [Go Community on Hashnode](https://hashnode.com/n/go) - Community of Gophers on Hashnode.
* [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
* [Go In 5 Minutes](https://www.goin5minutes.com/) - 5 minute screencasts focused on getting one thing done.
* [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki.
* [Go Report Card](https://goreportcard.com) - A report card for your Go package.
* [go.dev](https://go.dev/) - A hub for Go developers.
* [Go Blog](http://blog.golang.org) - The official Go blog.
* [godoc.org](https://godoc.org/) - Documentation for open source Go packages.
* [Golang Developer Jobs](https://golangjob.xyz) - Developer Jobs exclusivly for Golang related Roles.
* [Golang Flow](https://golangflow.io) - Post Updates, News, Packages and more.
* [Golang News](https://golangnews.com) - Links and news about Go programming.
* [CodinGame](https://www.codingame.com/) - Learn Go by solving interactive tasks using small games as practical examples.
* [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
* [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for #golang enthusiasts.
* [Gopher Community Chat](https://invite.slack.golangbridge.org) - Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
* [Gophercises](https://gophercises.com/) - Free coding exercises for budding gophers.
* [gowalker.org](https://gowalker.org) - Go Project API documentation.
* [json2go](https://m-zajac.github.io/json2go) - Advanced JSON to Go struct conversion - online tool.
* [justforfunc](https://www.youtube.com/c/justforfunc) - Youtube channel dedicated to Go programming language tips and tricks, hosted by  Francesc Campoy [@francesc](https://twitter.com/francesc).
* [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
* [r/Golang](https://www.reddit.com/r/golang) - News about Go.
* [studygolang](https://studygolang.com) - The community of studygolang in China.
* [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.
* [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

### Tutorials

* [Build web application with Golang](https://github.com/astaxie/build-web-application-with-golang) - Golang ebook intro how to build a web app with golang. Stars:`33.5K`.
* [go-patterns](https://github.com/tmrts/go-patterns) - Curated list of Go design patterns, recipes and idioms. Stars:`11.8K`.
* [Learn Go with TDD](https://github.com/quii/learn-go-with-tests) - Learn Go with test-driven development. Stars:`9.8K`.
* [Go Cheat Sheet](https://github.com/a8m/go-lang-cheat-sheet) - Go's reference card. Stars:`4.3K`.
* [Working with Go](https://github.com/mkaz/working-with-go) - Intro to go for experienced programmers. Stars:`1.1K`.
* [Golang for Node.js Developers](https://github.com/miguelmota/golang-for-nodejs-developers) - Examples of Golang compared to Node.js for learning. Stars:`1.1K`.
* [Ethereum Development with Go](https://github.com/miguelmota/ethereum-development-with-go-book) - A little e-book on Ethereum Development with Go. Stars:`536`.
* [50 Shades of Go](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/) - Traps, Gotchas, and Common Mistakes for New Golang Devs.
* [Games With Go](http://gameswithgo.org/) - A video series teaching programming and game development.
* [Go By Example](https://gobyexample.com/) - Hands-on introduction to Go using annotated example programs.
* [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30) - How to cancel MySQL queries.
* [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
* [Go Playground for iOS](https://itunes.apple.com/us/app/go-playground/id1437518275?ls=1&mt=8) - Interactively edit & play Go snippets on your mobile device.
* [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9) - How to cache slow database queries.
* [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
* [Golangbot](https://golangbot.com/learn-golang-series/) - Tutorials to get started with programming in Go.
* [GolangCode](https://golangcode.com/) - Collection of code snippets and tutorials to help tackle every day issues.
* [Hackr.io](https://hackr.io/tutorials/learn-golang) - Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
* [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
* [A Tour of Go](http://tour.golang.org/) - Interactive tour of Go.
* [Learning Golang - From zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/) - Getting started with golang for beginner.
* [package main](https://www.youtube.com/packagemain) - YouTube channel about Programming in Go.
* [Programming with Google Go](https://www.coursera.org/specializations/google-golang) - Coursera Specialization to learn about Go from scratch.
* [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
* [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo) - Building a Golang site for e-commerce (demo included).
* [Your basic Go](http://yourbasic.org/golang) - Huge collection of tutorials and how to's.
