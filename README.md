<h2 align="center">
  <br>
  <p align="center"><img width=30% src="https://raw.githubusercontent.com/kha7iq/pingme/master/.github/img/logo.png"></p>
</h2>

<h4 align="center">PingMe CLI</h4>

<p align="center">
   <a href="https://github.com/kha7iq/pingme/releases">
   <img alt="Release" src="https://img.shields.io/github/v/release/kha7iq/pingme">
   <a href="https://goreportcard.com/report/github.com/kha7iq/pingme">
   <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/kha7iq/pingme">
   <a href="#">
   <img alt="Build" src="https://img.shields.io/github/workflow/status/kha7iq/pingme/goreleaser">
   <a href="https://github.com/kha7iq/pingme/issues">
   
   <img alt="GitHub issues" src="https://img.shields.io/github/issues/kha7iq/pingme?style=flat-square&logo=github&logoColor=white">
   <a href="https://github.com/kha7iq/pingme/blob/master/LICENSE.md">
   <img alt="License" src="https://img.shields.io/github/license/kha7iq/pingme">
   <a href="#">
   <img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/kha7iq/pingme">
   <a href="https://pkg.go.dev/github.com/kha7iq/pingme">
   <img alt="Go Dev Reference" src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat">
</p>

<p align="center">
  <a href="#about">About</a> •
  <a href="https://kha7iq.github.io/pingme">Documentation</a> •
  <a href="#supported-services">Supported Services</a> •
  <a href="#install">Install</a> •
  <a href="#github-action">Github Action</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#show-your-support">Show Your Support</a> •
</p>

---

## About

**PingMe** is a personal project to satisfy my needs of having alerts, most major platforms have integration to send alerts
but its not always useful, either you are stuck with one particular platform, or you have to do alot of integrations. I needed a small app
which i can just call from my backup scripts, cron jobs, CI/CD pipelines or from anywhere to send a message with particular information.
And i can ship it everywhere with ease.
Hence, the birth of PingMe.

Everything is configurable via environment variables, and you can simply export the logs or messages to a variable which will be sent
as message. And most of all this serves as a swiss army knife sort of tool which supports multiple platforms.



## Supported services
- *Discord*
- *Email*
- *Microsoft Teams*
- *RocketChat*
- *Slack*
- *Telegram*
- *Pushover*
- *Mattermost*


## Install

### Linux & MacOs
```bash
brew install kha7iq/tap/pingme
```

### Go Get
```bash
go get -u github.com/kha7iq/pingme
```

### Windows
```powershell
scoop bucket add pingme https://github.com/kha7iq/scoop-bucket.git
scoop install pingme
```

Alternatively you can head over to [release pages](https://github.com/kha7iq/pingme/releases) and download the binary for windows & all other supported platforms.


### Docker
Docker container is also available on both dockerhub and github container registry.

`latest` tage will always pull the latest version avaialbe, you can also download specific version.
Checkout [release](https://github.com/kha7iq/pingme/releases) page for available versions.

Docker Registry
```bash
docker pull khaliq/pingme:latest
```
Gighub Registry
```bash
docker pull ghcr.io/kha7iq/pingme:latest
```
Run
```bash
docker run ghcr.io/kha7iq/pingme:latest
```
## Github Action

A github action is also available now for this app, you can find it on [Github Market Place](https://github.com/marketplace/actions/pingme-action) or from this [repository](https://github.com/kha7iq/pingme-action) on github.

Usage examples for workflow are available in the repo.


## Usage

```bash
❯ pingme

NAME:
   PingMe - Send message to multiple platforms

USAGE:
   main [global options] command [command options] [arguments...]

DESCRIPTION:
   PingMe is a CLI tool which provides the ability to send messages or alerts to multiple
   messaging platforms and also email, everything is configurable via environment
   variables and command line switches.Currently supported platforms include Slack, Telegram,
   RocketChat, Discord, Pushover, Mattermost, Microsoft Teams and email address.

COMMANDS:
   telegram    Send message to telegram
   rocketchat  Send message to rocketchat
   slack       Send message to slack
   discord     Send message to discord
   teams       Send message to microsoft teams
   pushover    Send message to pushover
   email       Send an email
   mattermost  Send message to mattermost
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

Check [Documentation Page](https://kha7iq.github.io/pingme/#/) for more details.

## Configuration

All the flags have corresponding environment variables associated with it. You can either provide the value with flags
or export to a variable.

View the [Documentation Page](https://kha7iq.github.io/pingme/#/) for more details.


## Contributing

Contributions, issues and feature requests are welcome!<br/>Feel free to check [issues page](https://github.com/kha7iq/pingme/issues). You can also take a look at the [contributing guide](https://github.com/kha7iq/pingme/blob/master/CONTRIBUTING.md).


## Show your support

Give a ⭐️  if you like this project!
