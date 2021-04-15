package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/pushbullet"
	"github.com/urfave/cli/v2"
)

// teleGram struct holds data parsed via flags for telegram service.
type pushBullet struct {
	Token     string
	Message   string
	Title     string
	Receivers string
}

// SendToTelegram parse values from *cli.context and return *cli.Command.
// Values include telegram token, channelIDs, Message and Title.
// If multiple channels are provided they the string is split with "," separator and
// each channelID is added to receiver.
func SendToPushBullet() *cli.Command {
	var pushBulletOpts pushBullet
	return &cli.Command{
		Name:  "pushbullet",
		Usage: "Send message to pushbullet",
		Description: `Pushbullet uses API token to authenticate & send messages to defined devices.
Multiple device nicknames can be used separated by comma ','.
All configuration options are also available via environment variables.`,
		UsageText: "pingme pushbullet --token '123' --devices 'Web123, myAndroid' --msg 'some message'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &pushBulletOpts.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "Token of pushbullet api used for sending message.",
				EnvVars:     []string{"PUSHBULLET_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &pushBulletOpts.Receivers,
				Name:        "devives",
				Aliases:     []string{"d"},
				Required:    true,
				Usage:       "DeviceNicknames of pushbullet.",
				EnvVars:     []string{"PUSHBULLET_DEVICES"},
			},
			&cli.StringFlag{
				Destination: &pushBulletOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content.",
				EnvVars:     []string{"PUSHBULLET_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &pushBulletOpts.Title,
				Name:        "title",
				Value:       TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"PUSHBULLET_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			notifier := notify.New()

			pushbulletSvc := pushbullet.New(pushBulletOpts.Token)
			devices := strings.Split(pushBulletOpts.Receivers, ",")
			for _, v := range devices {
				if len(v) <= 0 {
					return fmt.Errorf(EmptyChannel)
				}
				pushbulletSvc.AddReceivers(v)
			}

			notifier.UseServices(pushbulletSvc)

			if err := notifier.Send(
				context.Background(),
				pushBulletOpts.Title,
				pushBulletOpts.Message,
			); err != nil {
				return err
			}

			log.Println("Successfully sent!")
			return nil
		},
	}
}
