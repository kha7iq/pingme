package rocketchat

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/rocketchat"
	"github.com/urfave/cli/v2"
)

type rocketChat struct {
	Token     string
	UserID    string
	Message   string
	Channel   string
	Title     string
	ServerURL string
	Scheme    string
}

// SendMessage sends a message to rocketchat channels.
// channels can be comma-separated string of channel names.
func SendMessage(serverURL, scheme, userID, token, channels, title, message string) error {
	if serverURL == "" {
		return fmt.Errorf("rocketchat server URL is required")
	}
	if userID == "" {
		return fmt.Errorf("rocketchat user ID is required")
	}
	if token == "" {
		return fmt.Errorf("rocketchat token is required")
	}
	if channels == "" {
		return fmt.Errorf("rocketchat channel is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	notifier := notify.New()

	rocketChatSvc, err := rocketchat.New(serverURL, scheme, userID, token)
	if err != nil {
		return fmt.Errorf("failed to create rocketchat service: %w", err)
	}

	chn := strings.Split(channels, ",")
	for _, v := range chn {
		v = strings.TrimSpace(v)
		if len(v) <= 0 {
			return helpers.ErrChannel
		}
		rocketChatSvc.AddReceivers(v)
	}

	notifier.UseServices(rocketChatSvc)

	if err = notifier.Send(context.Background(), title, message); err != nil {
		return fmt.Errorf("failed to send rocketchat message: %w", err)
	}

	log.Println("Successfully sent!")
	return nil
}

// Send parse values from *cli.context and return *cli.Command.
func Send() *cli.Command {
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
				Destination: &rocketChatOpts.UserID,
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
				Value:       helpers.TimeValue,
				Usage:       "Title of the message",
				EnvVars:     []string{"ROCKETCHAT_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return SendMessage(
				rocketChatOpts.ServerURL,
				rocketChatOpts.Scheme,
				rocketChatOpts.UserID,
				rocketChatOpts.Token,
				rocketChatOpts.Channel,
				rocketChatOpts.Title,
				rocketChatOpts.Message,
			)
		},
	}
}
