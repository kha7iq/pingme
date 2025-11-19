package telegram

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
	"github.com/urfave/cli/v2"
)

// teleGram struct holds data parsed via flags for telegram service.
type teleGram struct {
	Token   string
	Message string
	Channel string
	Title   string
}

// SendMessage sends a message to telegram channels.
// This is the core logic extracted for reuse by both CLI and webhook.
// channels can be comma-separated string of channel IDs.
func SendMessage(token, channels, title, message string) error {
	if token == "" {
		return fmt.Errorf("telegram token is required")
	}
	if channels == "" {
		return fmt.Errorf("telegram channel is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	notifier := notify.New()

	telegramSvc, err := telegram.New(token)
	if err != nil {
		return fmt.Errorf("failed to create telegram service: %w", err)
	}

	chn := strings.Split(channels, ",")
	for _, v := range chn {
		v = strings.TrimSpace(v)
		if len(v) <= 0 {
			return helpers.ErrChannel
		}
		chatID, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid channel ID '%s': %w", v, err)
		}
		telegramSvc.AddReceivers(chatID)
	}

	notifier.UseServices(telegramSvc)

	if err = notifier.Send(context.Background(), title, message); err != nil {
		return fmt.Errorf("failed to send telegram message: %w", err)
	}

	log.Println("Successfully sent!")
	return nil
}

// Send parse values from *cli.context and return *cli.Command.
// Values include telegram token, channelIDs, Message and Title.
// If multiple channels are provided they the string is split with "," separator and
// each channelID is added to receiver.
func Send() *cli.Command {
	var telegramOpts teleGram
	return &cli.Command{
		Name:  "telegram",
		Usage: "Send message to telegram",
		Description: `Telegram uses bot token to authenticate & send messages to defined channels.
Multiple channel ids can be used separated by comma ','.
All configuration options are also available via environment variables.`,
		UsageText: "pingme telegram --token '123' --channel '-123456' --msg 'some message'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &telegramOpts.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "Token of telegram bot used for sending message.",
				EnvVars:     []string{"TELEGRAM_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &telegramOpts.Channel,
				Name:        "channel",
				Aliases:     []string{"c"},
				Required:    true,
				Usage:       "Channel ids of telegram.",
				EnvVars:     []string{"TELEGRAM_CHANNELS"},
			},
			&cli.StringFlag{
				Destination: &telegramOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content.",
				EnvVars:     []string{"TELEGRAM_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &telegramOpts.Title,
				Name:        "title",
				Value:       helpers.TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"TELEGRAM_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			// Now just call the extracted function
			return SendMessage(
				telegramOpts.Token,
				telegramOpts.Channel,
				telegramOpts.Title,
				telegramOpts.Message,
			)
		},
	}
}
