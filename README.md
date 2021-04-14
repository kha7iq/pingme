<h2 align="center">
  <br>
  <p align="center"><img width=30% src="https://github.com/kha7iq/pingme/blob/master/.github/img/logo.png"></p>
</h2>

<h4 align="center">PingMe CLI</h4>

<p align="center">
    <a href="https://github.com/kha7iq/pingme/blob/master/LICENSE.md">
    <img alt="License" src="https://img.shields.io/github/license/kha7iq/pingme?style=flat-square&logo=github&logoColor=white">
    <a href="https://github.com/kha7iq/pingme/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/kha7iq/pingme?style=flat-square&logo=github&logoColor=white">
</p>

<p align="center">
  <a href="#about">About</a> •
  <a href="#about">Documentation</a> •
  <a href="#supported-services">Supported Services</a> •
  <a href="#install">Install</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#show-your-support">Show Your Support</a> •
</p>

---

## About

**PingMe** is a personal project to satisfy my needs of having alerts, most of major platforms have integration to send alerts

but its not always useful, either you are stuck with one particlaur platform or you have to do alot of integrations. I needed a small app

which i can just call from my backup scripts, cron jobs, CI/CD pipelines or from anywhere to send a message with particular information.

And i can ship it everywhere with ease.

Everything should be configurable via enviornment variables and i can simply export the logs or messages to a variable which will be sent

as message. And most of all this should serve as a swiss army knife sort of tool which supports multiple platforms.

Hence the birth of PingMe.


## Supported services
- *Discord*
- *Email*
- *Microsoft Teams*
- *RocketChat*
- *Slack*
- *Telegram*


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
Alternativly you can head over to [release pages](https://github.com/kha7iq/pingme/releases) and download the binary for windows & all other supported platforms.

## Usage

```bash
❯ pingme

NAME:
   PingMe - Send message to multiple platforms

USAGE:
   pingme [global options] command [command options] [arguments...]

DESCRIPTION:
   PingMe is a CLI tool which provides the ability to send messages or alerts to multiple
   messaging platforms and also email, everything is configurable via environment
   variables and command line switches.Currently supported platforms include Slack, Telegram,
   RocketChat, Discord, Microsoft Teams and email address.

COMMANDS:
   telegram    Send message to telegram
   rocketchat  Send message to rocketchat
   slack       Send message to slack
   discord     Send message to discord
   teams       Send message to microsoft teams
   email       Send an email
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```


### Telegram
```bash
pingme  telegram  --token "0125:AAFHvnYf_ABC"  --msg "This is a new message ✈" --channel="-1001001001,-1002002001"
```


* ![Demo](https://github.com/kha7iq/pingme/blob/master/.github/img/pingme.gif)


## Configuration

All the flags have crosponding enviornment variables assosiated with it. You can either provide the value with flags

or export to a variable. View the [Documentation Page](https://github.com/) for more details


## Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/kha7iq/pingme/issues). You can also take a look at the [contributing guide](https://github.com/kha7iq/pingme/blob/master/CONTRIBUTING.md).



## Show your support

Give a ⭐️  if you like this project!
