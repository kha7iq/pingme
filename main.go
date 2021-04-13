package main

import (
	"log"
	"os"

	"github.com/kha7iq/pingme/cmd"

	"github.com/urfave/cli/v2"
)

var version = "v1.0"

func main() {
	app := cli.NewApp()
	app.Name = "PingMe"
	app.Version = version
	app.Usage = "Send message to multiple platforms"
	app.Description = `PingMe is a CLI tool which provides the ability to send messages or alerts to multiple 
messaging platforms and also email, everything is configurable via environment
variables and command line switches.Currently supported platforms include Slack, Telegram,
RocketChat, Discord, Microsoft Teams and email address.`

	app.Commands = []*cli.Command{
		cmd.SendToTelegram(),
		cmd.SendToRocketChat(),
		cmd.SendToSlack(),
		cmd.SendToDiscord(),
		cmd.SendToTeams(),
		cmd.SendToEmail(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
