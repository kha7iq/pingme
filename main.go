package main

import (
	"log"
	"os"

	"github.com/kha7iq/pingme/service/gotify"
	"github.com/kha7iq/pingme/service/textmagic"

	"github.com/kha7iq/pingme/service/mastodon"
	"github.com/kha7iq/pingme/service/twillio"
	"github.com/kha7iq/pingme/service/wechat"
	"github.com/kha7iq/pingme/service/zulip"

	"github.com/kha7iq/pingme/service/discord"
	"github.com/kha7iq/pingme/service/email"
	"github.com/kha7iq/pingme/service/line"
	"github.com/kha7iq/pingme/service/mattermost"
	"github.com/kha7iq/pingme/service/msteams"
	"github.com/kha7iq/pingme/service/pushbullet"
	"github.com/kha7iq/pingme/service/pushover"
	"github.com/kha7iq/pingme/service/rocketchat"
	"github.com/kha7iq/pingme/service/slack"
	"github.com/kha7iq/pingme/service/telegram"

	"github.com/urfave/cli/v2"
)

// Version variable is used for semVer
var Version string

// main with combine all the function into commands
func main() {
	app := cli.NewApp()
	app.Name = "PingMe"
	app.Version = Version
	app.Usage = "Send message to multiple platforms"
	app.Description = `PingMe is a CLI tool which provides the ability to send messages or alerts to multiple 
messaging platforms and also email, everything is configurable via environment
variables and command line switches.Currently supported platforms include Slack, Telegram,
RocketChat, Discord, Pushover, Mattermost, Pushbullet, Microsoft Teams, Twillio, Mastodon,
email address, Line, Gotify and Wechat.`
	// app.Commands contains the subcommands as functions which return []*cli.Command.
	app.Commands = []*cli.Command{
		telegram.Send(),
		rocketchat.Send(),
		slack.Send(),
		discord.Send(),
		msteams.Send(),
		pushover.Send(),
		email.Send(),
		mattermost.Send(),
		pushbullet.Send(),
		twillio.Send(),
		zulip.Send(),
		mastodon.Send(),
		line.Send(),
		wechat.Send(),
		gotify.Send(),
		textmagic.Send(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
