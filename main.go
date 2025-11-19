package main

import (
	"log"
	"os"

	"github.com/kha7iq/pingme/internal/server"
	"github.com/kha7iq/pingme/service/discord"
	"github.com/kha7iq/pingme/service/email"
	"github.com/kha7iq/pingme/service/gotify"
	"github.com/kha7iq/pingme/service/line"
	"github.com/kha7iq/pingme/service/mastodon"
	"github.com/kha7iq/pingme/service/matrix"
	"github.com/kha7iq/pingme/service/mattermost"
	"github.com/kha7iq/pingme/service/msteams"
	"github.com/kha7iq/pingme/service/pushbullet"
	"github.com/kha7iq/pingme/service/pushover"
	"github.com/kha7iq/pingme/service/rocketchat"
	"github.com/kha7iq/pingme/service/slack"
	"github.com/kha7iq/pingme/service/telegram"
	"github.com/kha7iq/pingme/service/twillio"
	"github.com/kha7iq/pingme/service/wechat"
	"github.com/kha7iq/pingme/service/zulip"
	"github.com/urfave/cli/v2"
)

var (
	// Version variable is used for semVer
	Version string
)

func main() {
	app := cli.NewApp()
	app.Name = "PingMe"
	app.Version = Version
	app.Usage = "Send message to multiple platforms"
	app.Description = `PingMe is a CLI tool to send messages to multiple platforms.
It also supports running as a webhook server to receive and dispatch notifications.`

	app.Commands = []*cli.Command{
		// Webhook server command
		{
			Name:    "serve",
			Aliases: []string{"server", "webhook"},
			Usage:   "Start webhook server",
			Description: `Start a webhook server that receives POST requests and dispatches notifications.
Configuration is done via environment variables (same as CLI commands).

Authentication (optional):
  Set PINGME_AUTH_METHOD to: "apikey", "hmac", "basic", or "none"
  
  For apikey: Set PINGME_API_KEYS="key1,key2,key3"
  For hmac: Set PINGME_HMAC_SECRET="your-secret"
  For basic: Set PINGME_BASIC_USER="user" and PINGME_BASIC_PASS="pass"`,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "port",
					Aliases: []string{"p"},
					Usage:   "Server port",
					Value:   "8080",
					EnvVars: []string{"PINGME_PORT"},
				},
				&cli.StringFlag{
					Name:    "host",
					Aliases: []string{"H"},
					Usage:   "Server host",
					Value:   "0.0.0.0",
					EnvVars: []string{"PINGME_HOST"},
				},
			},
			Action: func(c *cli.Context) error {
				port := c.String("port")
				host := c.String("host")

				srv := server.New(host, port)
				return srv.Start()
			},
		},
		// service commands
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
		mastodon.Send(),
		wechat.Send(),
		zulip.Send(),
		gotify.Send(),
		line.Send(),
		matrix.Send(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
