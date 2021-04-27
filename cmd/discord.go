package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

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

// SendToDiscord parse values from *cli.context and return *cli.Command.
// Values include discord bot token, userID, channelIDs, Message and Title.
// If multiple channels are provided then the string is split with "," separator and
// each channelID is added to receiver.
func SendToDiscord() *cli.Command {
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
				Value:       TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"DISCORD_MSG_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			notifier := notify.New()
			discordSvc := discord.New()

			if err := discordSvc.AuthenticateWithBotToken(discordOpts.Token); err != nil {
				return fmt.Errorf("unable to authenticate %v\n", err)
			}

			chn := strings.Split(discordOpts.Channel, ",")
			for _, v := range chn {
				if len(v) <= 0 {
					return fmt.Errorf(EmptyChannel)
				}

				discordSvc.AddReceivers(v)
			}

			notifier.UseServices(discordSvc)

			if err := notifier.Send(
				context.Background(),
				discordOpts.Title,
				discordOpts.Message,
			); err != nil {
				return err
			}
			log.Println("Successfully sent!")
			return nil
		},
	}
}
