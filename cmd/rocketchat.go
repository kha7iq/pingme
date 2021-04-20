package cmd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/rocketchat"
	"github.com/urfave/cli/v2"
)

type rocketChat struct {
	Token     string
	UserId    string
	Message   string
	Channel   string
	Title     string
	ServerURL string
	Scheme    string
}

var (
	// EmptyChannel variable holds default error message if no channel is provided.
	EmptyChannel = "channel name or id can not be empty"
	TimeValue    = "⏰ " + time.Now().Format(time.UnixDate)
)

// SendToRocketChat parse values from *cli.context and return *cli.Command.
// Values include rocketchat token, , UserId, channelIDs, ServerURL, Scheme, Message and Title.
// If multiple channels are provided then the string is split with "," separator and
// each channelID is added to receiver.
func SendToRocketChat() *cli.Command {
	var rocketChatOpts rocketChat
	return &cli.Command{
		Name:  "rocketchat",
		Usage: "Send message to rocketchat",
		Description: `RocketChat uses token & userID to authenticate and send messages to defined channels.
Multiple channel ids can be used separated by comma ','.
All configuration options are also available via environment variables.`,
		UsageText: "pingme rocketchat --userid '123' --token 'xyz'  --url 'localhost' --scheme 'http'" +
			" --channel 'alert' --msg 'some message'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &rocketChatOpts.UserId,
				Name:        "userid",
				Aliases:     []string{"id"},
				Required:    true,
				Usage:       "User ID",
				EnvVars:     []string{"ROCKETCHAT_USERID"},
			},
			&cli.StringFlag{
				Destination: &rocketChatOpts.Token,
				Name:        "token",
				Required:    true,
				Aliases:     []string{"t"},
				Usage:       "Auth token for sending message, can also be set as environment variable",
				EnvVars:     []string{"ROCKETCHAT_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &rocketChatOpts.ServerURL,
				Name:        "url",
				Aliases:     []string{"u"},
				Required:    true,
				Usage:       "Rocketchat server url",
				EnvVars:     []string{"ROCKETCHAT_SERVER_URL"},
			},
			&cli.StringFlag{
				Destination: &rocketChatOpts.Scheme,
				Name:        "scheme",
				Usage:       "URL scheme http/https",
				Value:       "https",
				Aliases:     []string{"s"},
				EnvVars:     []string{"ROCKETCHAT_URL_SCHEME"},
			},
			&cli.StringFlag{
				Destination: &rocketChatOpts.Channel,
				Name:        "channel",
				Aliases:     []string{"c"},
				Usage:       "Channel names separated by comma ','",
				EnvVars:     []string{"ROCKETCHAT_CHANNELS"},
			},
			&cli.StringFlag{
				Destination: &rocketChatOpts.Message,
				Name:        "msg",
				Usage:       "Message content",
				EnvVars:     []string{"ROCKETCHAT_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &rocketChatOpts.Title,
				Name:        "title",
				Value:       TimeValue,
				Usage:       "Title of the message",
				EnvVars:     []string{"ROCKETCHAT_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			notifier := notify.New()

			rocketChatSvc, err := rocketchat.New(rocketChatOpts.ServerURL, rocketChatOpts.Scheme,
				rocketChatOpts.UserId, rocketChatOpts.Token)
			if err != nil {
				return err
			}
			chn := strings.Split(rocketChatOpts.Channel, ",")
			for _, v := range chn {
				if len(v) <= 0 {
					return fmt.Errorf(EmptyChannel)
				}

				rocketChatSvc.AddReceivers(v)
			}

			notifier.UseServices(rocketChatSvc)

			if err = notifier.Send(
				context.Background(),
				rocketChatOpts.Title,
				rocketChatOpts.Message,
			); err != nil {
				return err
			}
			return nil
		},
	}
}
