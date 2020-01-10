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

## Images

*Libraries for manipulating images.*

* [gocv](https://github.com/hybridgroup/gocv) - Go package for computer vision using OpenCV 3.3+.Star:`2867`.
* [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package.Star:`2857`.
* [imaginary](https://github.com/h2non/imaginary) - Fast and simple HTTP microservice for image resizing.Star:`2796`.
* [bild](https://github.com/anthonynsimon/bild) - Collection of image processing algorithms in pure Go.Star:`2736`.
* [ln](https://github.com/fogleman/ln) - 3D line art rendering in Go.Star:`2586`.
* [resize](https://github.com/nfnt/resize) - Image resizing for Go with common interpolation methods.Star:`2247`.
* [gg](https://github.com/fogleman/gg) - 2D rendering in pure Go.Star:`2087`.
* [pt](https://github.com/fogleman/pt) - Path tracing engine written in Go.Star:`1826`.
* [svgo](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation.Star:`1410`.
* [smartcrop](https://github.com/muesli/smartcrop) - Finds good crops for arbitrary images and crop sizes.Star:`1305`.
* [gift](https://github.com/disintegration/gift) - Package of image processing filters.Star:`1285`.
* [picfit](https://github.com/thoas/picfit) - An image resizing server written in Go.Star:`1165`.
* [go-opencv](https://github.com/lazywei/go-opencv) - Go bindings for OpenCV.Star:`1139`.
* [imagick](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API.Star:`1057`.
* [geopattern](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string.Star:`1040`.
* [bimg](https://github.com/h2non/bimg) - Small package for fast and efficient image processing using libvips.Star:`867`.
* [stegify](https://github.com/DimitarPetrov/stegify) - Go tool for LSB steganography, capable of hiding any file within an image.Star:`590`.
* [canvas](https://github.com/tdewolff/canvas) - Vector graphics to PDF, SVG or rasterized image.Star:`382`.
* [mort](https://github.com/aldor007/mort) - Storage and image processing server written in Go.Star:`381`.
* [image2ascii](https://github.com/qeesung/image2ascii) - Convert image to ASCII.Star:`349`.
* [govatar](https://github.com/o1egl/govatar) - Library and CMD tool for generating funny avatars.Star:`332`.
* [go-nude](https://github.com/koyachi/go-nude) - Nudity detection with Go.Star:`296`.
* [goimagehash](https://github.com/corona10/goimagehash) - Go Perceptual image hashing package.Star:`260`.
* [rez](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD.Star:`194`.
* [img](https://github.com/hawx/img) - Selection of image manipulation tools.Star:`132`.
* [darkroom](https://github.com/gojek/darkroom) - An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.Star:`109`.
* [go-cairo](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library.Star:`90`.
* [mergi](https://github.com/noelyahan/mergi) - Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).Star:`85`.
* [gltf](https://github.com/qmuntal/gltf) - Efficient and robust glTF 2.0 reader, writer and validator.Star:`57`.
* [go-gd](https://github.com/bolknote/go-gd) - Go binding for GD library.Star:`54`.
* [steganography](https://github.com/auyer/steganography) - Pure Go Library for LSB steganography.Star:`36`.
* [cameron](https://github.com/aofei/cameron) - An avatar generator for Go.Star:`36`.
* [goimghdr](https://github.com/corona10/goimghdr) - The imghdr module determines the type of image contained in a file for Go.Star:`31`.
* [tga](https://github.com/ftrvxmtrx/tga) - Package tga is a TARGA image format decoder/encoder.Star:`25`.
* [go-webcolors](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go.Star:`24`.
* [mpo](https://github.com/donatj/mpo) - Decoder and conversion tool for MPO 3D Photos.Star:`6`.
