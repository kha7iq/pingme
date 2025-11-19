package pushbullet

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"

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

// SendMessage sends a message via pushbullet to devices.
// devices can be comma-separated string of device nicknames.
func SendMessage(token, devices, title, message string) error {
	if token == "" {
		return fmt.Errorf("pushbullet token is required")
	}
	if devices == "" {
		return fmt.Errorf("pushbullet device is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	notifier := notify.New()
	pushBulletSvc := pushbullet.New(token)

	deviceList := strings.Split(devices, ",")
	for _, v := range deviceList {
		v = strings.TrimSpace(v)
		if len(v) <= 0 {
			return helpers.ErrChannel
		}
		pushBulletSvc.AddReceivers(v)
	}

	notifier.UseServices(pushBulletSvc)

	if err := notifier.Send(context.Background(), title, message); err != nil {
		return fmt.Errorf("failed to send pushbullet message: %w", err)
	}

	log.Println("Successfully sent!")
	return nil
}

// SendSMS sends an SMS via pushbullet.
func SendSMS(token, device, phoneNumber, title, message string) error {
	if token == "" {
		return fmt.Errorf("pushbullet token is required")
	}
	if device == "" {
		return fmt.Errorf("pushbullet device is required")
	}
	if phoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	notifier := notify.New()

	pushBulletSmsSvc, err := pushbullet.NewSMS(token, device)
	if err != nil {
		return fmt.Errorf("failed to create pushbullet SMS service: %w", err)
	}

	numbers := strings.Split(phoneNumber, ",")
	for _, v := range numbers {
		v = strings.TrimSpace(v)
		if len(v) <= 0 {
			return helpers.ErrChannel
		}
		pushBulletSmsSvc.AddReceivers(v)

		notifier.UseServices(pushBulletSmsSvc)

		if err := notifier.Send(context.Background(), title, message); err != nil {
			return fmt.Errorf("failed to send SMS to %s: %w", v, err)
		}
	}

	log.Println("Successfully sent!")
	return nil
}

// Send parse values from *cli.context and return *cli.Command.
func Send() *cli.Command {
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
				Value:       helpers.TimeValue,
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
			if pushBulletOpts.SMS {
				return SendSMS(
					pushBulletOpts.Token,
					pushBulletOpts.Device,
					pushBulletOpts.PhoneNumber,
					pushBulletOpts.Title,
					pushBulletOpts.Message,
				)
			}
			return SendMessage(
				pushBulletOpts.Token,
				pushBulletOpts.Device,
				pushBulletOpts.Title,
				pushBulletOpts.Message,
			)
		},
	}
}
