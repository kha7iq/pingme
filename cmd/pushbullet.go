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

// pushBullet struct holds data parsed via flags for pushbullet service.
type pushBullet struct {
	Token       string
	Message     string
	Title       string
	Device      string
	PhoneNumber string
	SMS         bool
}

// SendToPushBullet parse values from *cli.context and return *cli.Command.
// Values include pushbullet token, Device, phone number, Message and Title.
// If multiple devices are provided they the string is split with "," separator and
// each device is added to receiver.
func SendToPushBullet() *cli.Command {
	var pushBulletOpts pushBullet
	return &cli.Command{
		Name:  "pushbullet",
		Usage: "Send message to pushbullet",
		Description: `Pushbullet uses API token to authenticate & send messages to defined devices.
Multiple device nicknames or numbers can be used separated by comma.`,
		UsageText: "pingme pushbullet --token '123' --device 'Web123, myAndroid' --msg 'some message'\n" +
			"pingme pushbullet --token '123' --sms true --device 'Web123' --msg 'some message' --number '00123456789'",
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
				Destination: &pushBulletOpts.Device,
				Name:        "device",
				Aliases:     []string{"d"},
				Required:    true,
				Usage:       "Device's nickname of pushbullet.",
				EnvVars:     []string{"PUSHBULLET_DEVICE"},
			},
			&cli.StringFlag{
				Destination: &pushBulletOpts.PhoneNumber,
				Name:        "number",
				Aliases:     []string{"n"},
				Usage:       "Target phone number",
				EnvVars:     []string{"PUSHBULLET_NUMBER"},
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
			&cli.BoolFlag{
				Destination: &pushBulletOpts.SMS,
				Name:        "sms",
				Value:       false,
				Usage:       "To send sms message set the value to 'true'",
				EnvVars:     []string{"PUSHBULLET_SMS"},
			},
		},
		Action: func(ctx *cli.Context) error {
			notifier := notify.New()

			switch pushBulletOpts.SMS {
			case true:
				pushBulletSmsSvc, err := pushbullet.NewSMS(pushBulletOpts.Token, pushBulletOpts.Device)
				if err != nil {
					return err
				}
				devices := strings.Split(pushBulletOpts.PhoneNumber, ",")
				for _, v := range devices {
					if len(v) <= 0 {
						return fmt.Errorf(EmptyChannel)
					}
					pushBulletSmsSvc.AddReceivers(v)

					notifier.UseServices(pushBulletSmsSvc)

					if err := notifier.Send(
						context.Background(),
						pushBulletOpts.Title,
						pushBulletOpts.Message,
					); err != nil {
						return err
					}
				}
			default:
				pushBulletSvc := pushbullet.New(pushBulletOpts.Token)

				devices := strings.Split(pushBulletOpts.Device, ",")
				for _, v := range devices {
					if len(v) <= 0 {
						return fmt.Errorf(EmptyChannel)
					}
					pushBulletSvc.AddReceivers(v)
				}

				notifier.UseServices(pushBulletSvc)

				if err := notifier.Send(
					context.Background(),
					pushBulletOpts.Title,
					pushBulletOpts.Message,
				); err != nil {
					return err
				}
			}

			log.Println("Successfully sent!")
			return nil
		},
	}
}
