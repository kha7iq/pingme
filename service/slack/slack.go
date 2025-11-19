package slack

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"

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

// SendMessage sends a message to slack channels.
// channels can be comma-separated string of channel IDs.
func SendMessage(token, channels, title, message string) error {
	if token == "" {
		return fmt.Errorf("slack token is required")
	}
	if channels == "" {
		return fmt.Errorf("slack channel is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	notifier := notify.New()
	slackSvc := slack.New(token)

	chn := strings.Split(channels, ",")
	for _, v := range chn {
		v = strings.TrimSpace(v)
		if len(v) <= 0 {
			return helpers.ErrChannel
		}
		slackSvc.AddReceivers(v)
	}

	notifier.UseServices(slackSvc)

	if err := notifier.Send(context.Background(), title, message); err != nil {
		return fmt.Errorf("failed to send slack message: %w", err)
	}

	log.Println("Successfully sent!")
	return nil
}

// Send parse values from *cli.context and return *cli.Command.
func Send() *cli.Command {
	var slackOpts slackPingMe
	return &cli.Command{
		Name:      "slack",
		Usage:     "Send message to slack",
		UsageText: "pingme slack --token '123' --channel '12345,67890' --msg 'some message'",
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
				Value:       helpers.TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"SLACK_MSG_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return SendMessage(
				slackOpts.Token,
				slackOpts.Channel,
				slackOpts.Title,
				slackOpts.Message,
			)
		},
	}
}
