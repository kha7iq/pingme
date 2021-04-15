package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/slack"
	"github.com/urfave/cli/v2"
)

// slackPingMe struct holds data parsed via flags for slack service.
type slackPingMe struct {
	Token   string
	Message string
	Channel string
	Title   string
}

// SendToSlack parse values from *cli.context and return *cli.Command.
// Values include slack token, channelIDs, Message and Title.
// If multiple channels are provided then the string is split with "," separator and
// each channelID is added to receiver.
func SendToSlack() *cli.Command {
	var slackOpts slackPingMe
	return &cli.Command{
		Name:      "slack",
		Usage:     "Send message to slack",
		UsageText: "pingme slack --token '123' --channel '12345,67890' --message 'some message'",
		Description: `Slack uses token to authenticate and send messages to defined channels.
Multiple channel ids can be used separated by comma ','.
All configuration options are also available via environment variables.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &slackOpts.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "Token of slack bot used for sending message.",
				EnvVars:     []string{"SLACK_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &slackOpts.Channel,
				Name:        "channel",
				Required:    true,
				Aliases:     []string{"c"},
				Usage:       "Channel ids for slack,if sending to multiple channels separate with ','.",
				EnvVars:     []string{"SLACK_CHANNELS"},
			},
			&cli.StringFlag{
				Destination: &slackOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content.",
				EnvVars:     []string{"SLACK_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &slackOpts.Title,
				Name:        "title",
				Value:       TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"SLACK_MSG_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			notifier := notify.New()
			slackSvc := slack.New(slackOpts.Token)
			chn := strings.Split(slackOpts.Channel, ",")
			for _, v := range chn {
				if len(v) <= 0 {
					return fmt.Errorf(EmptyChannel)
				}

				slackSvc.AddReceivers(v)
			}

			notifier.UseServices(slackSvc)

			if err := notifier.Send(
				context.Background(),
				slackOpts.Title,
				slackOpts.Message,
			); err != nil {
				return err
			}
			log.Println("Successfully sent!")
			return nil
		},
	}
}
