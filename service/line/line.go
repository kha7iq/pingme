package line

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/line"
	"github.com/urfave/cli/v2"
)

// Line struct holds data parsed via flags for the service
type Line struct {
	Secret    string
	Token     string
	Message   string
	Receivers string
	Title     string
}

// SendMessage sends a message to line messenger receivers.
// receivers can be comma-separated string of user or group IDs.
func SendMessage(secret, token, receivers, title, message string) error {
	if secret == "" {
		return fmt.Errorf("line channel secret is required")
	}
	if token == "" {
		return fmt.Errorf("line channel access token is required")
	}
	if receivers == "" {
		return fmt.Errorf("line receiver IDs are required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	notifier := notify.New()

	lineSvc, err := line.New(secret, token)
	if err != nil {
		return fmt.Errorf("failed to create line service: %w", err)
	}

	// Add receiver IDs
	recv := strings.Split(receivers, ",")
	for _, r := range recv {
		r = strings.TrimSpace(r)
		if r != "" {
			lineSvc.AddReceivers(r)
		}
	}

	notifier.UseServices(lineSvc)

	if err := notifier.Send(context.Background(), title, message); err != nil {
		return fmt.Errorf("failed to send line message: %w", err)
	}

	log.Println("Successfully sent!")
	return nil
}

// Send parses values from *cli.context and returns a *cli.Command.
func Send() *cli.Command {
	var lineOpts Line
	return &cli.Command{
		Name:  "line",
		Usage: "Send message to line messenger",
		Description: `Line messenger uses a channel secret and
a channel access token to authenticate & send messages
through line to various receivers.`,
		UsageText: "pingme line --secret '123' --token '123' --msg 'some message' --receivers '123,456,789'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &lineOpts.Secret,
				Name:        "secret",
				Required:    true,
				Usage:       "Channel secret.",
				EnvVars:     []string{"LINE_SECRET"},
			},
			&cli.StringFlag{
				Destination: &lineOpts.Token,
				Name:        "token",
				Required:    true,
				Usage:       "Channel access token.",
				EnvVars:     []string{"LINE_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &lineOpts.Message,
				Name:        "msg",
				Required:    true,
				Usage:       "Message content.",
				EnvVars:     []string{"LINE_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &lineOpts.Title,
				Name:        "title",
				Value:       helpers.TimeValue,
				Usage:       "Message title.",
				EnvVars:     []string{"LINE_MSG_TITLE"},
			},
			&cli.StringFlag{
				Destination: &lineOpts.Receivers,
				Name:        "receivers",
				Required:    true,
				Usage:       "Comma-separated list of user or group receiver IDs.",
				EnvVars:     []string{"LINE_RECEIVER_IDS"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return SendMessage(
				lineOpts.Secret,
				lineOpts.Token,
				lineOpts.Receivers,
				lineOpts.Title,
				lineOpts.Message,
			)
		},
	}
}
