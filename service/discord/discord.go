package discord

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/discord"
	"github.com/urfave/cli/v2"
)

// discordPingMe struct holds data parsed via flags for discord service.
type discordPingMe struct {
	Token   string
	Message string
	Channel string
	Title   string
}

// SendMessage sends a message to discord channels.
// channels can be comma-separated string of channel IDs.
func SendMessage(token, channels, title, message string) error {
	if token == "" {
		return fmt.Errorf("discord token is required")
	}
	if channels == "" {
		return fmt.Errorf("discord channel is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	notifier := notify.New()
	discordSvc := discord.New()

	if err := discordSvc.AuthenticateWithBotToken(token); err != nil {
		return fmt.Errorf("unable to authenticate: %w", err)
	}

	chn := strings.Split(channels, ",")
	for _, v := range chn {
		v = strings.TrimSpace(v)
		if len(v) <= 0 {
			return helpers.ErrChannel
		}
		discordSvc.AddReceivers(v)
	}

	notifier.UseServices(discordSvc)

	if err := notifier.Send(context.Background(), title, message); err != nil {
		return fmt.Errorf("failed to send discord message: %w", err)
	}

	log.Println("Successfully sent!")
	return nil
}

// Send parse values from *cli.context and return *cli.Command.
func Send() *cli.Command {
	var discordOpts discordPingMe
	return &cli.Command{
		Name:  "discord",
		Usage: "Send message to discord",
		Description: `Discord uses bot token to authenticate & send messages to defined channels.
Multiple channel ids can be used separated by comma ','.
All configuration options are also available via environment variables.`,
		UsageText: "pingme discord --token '123' --channel '12345,67890' --msg 'some message'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &discordOpts.Token,
				Name:        "token",
				Required:    true,
				Usage:       "Token of discord bot used for sending message.",
				EnvVars:     []string{"DISCORD_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &discordOpts.Channel,
				Name:        "channel",
				Required:    true,
				Usage:       "Channel ids of discord.",
				EnvVars:     []string{"DISCORD_CHANNELS"},
			},
			&cli.StringFlag{
				Destination: &discordOpts.Message,
				Name:        "msg",
				Usage:       "Message content.",
				EnvVars:     []string{"DISCORD_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &discordOpts.Title,
				Name:        "title",
				Value:       helpers.TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"DISCORD_MSG_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return SendMessage(
				discordOpts.Token,
				discordOpts.Channel,
				discordOpts.Title,
				discordOpts.Message,
			)
		},
	}
}
