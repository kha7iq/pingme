package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/mail"
	"github.com/urfave/cli/v2"
)

// email struct holds data parsed via flags for email service.
type email struct {
	SenderAddress   string
	Password        string
	Host            string
	ReceiverAddress string
	Subject         string
	Message         string
	Port            string
	Identity        string
}

// SendToEmail parse values from *cli.context and return *cli.Command.
// SendAddress is used for authentication with smtp server, host and port is required
// the default value for port is set to "587" and host as "smtp.gmail.com"
// If multiple ReceiverAddress are provided then the string value is split with "," separator and
// each ReceiverAddress is added to receiver.
func SendToEmail() *cli.Command {
	var emailOpts email
	return &cli.Command{
		Name:  "email",
		Usage: "Send an email",
		Description: `Email uses username  & password to authenticate for sending emails.
SMTP hostname i.e smtp.gmail.com and port i.e (587) should be provided as well for the server.
Multiple email ids can be used separated by comma ',' as receiver email address.
All configuration options are also available via environment variables.`,
		UsageText: "pingme email --pass '123456' --sender 'abc@email.com' --rec 'xyz@example.com' " +
			"--msg 'some message' --sub 'email subject' --host 'smtp.gmail.com' --port '587'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &emailOpts.Password,
				Name:        "pass",
				Required:    true,
				Aliases:     []string{"p"},
				Usage:       "Password of email address, use environment variable",
				EnvVars:     []string{"EMAIL_PASSWORD"},
			},
			&cli.StringFlag{
				Destination: &emailOpts.ReceiverAddress,
				Name:        "rec",
				Aliases:     []string{"r"},
				Required:    true,
				Usage:       "Receiver email address, if multiple separate with ','",
				EnvVars:     []string{"EMAIL_RECEIVER"},
			},
			&cli.StringFlag{
				Destination: &emailOpts.SenderAddress,
				Name:        "sender",
				Aliases:     []string{"s"},
				Required:    true,
				Usage:       "Senders email address",
				EnvVars:     []string{"EMAIL_SENDER"},
			},
			&cli.StringFlag{
				Destination: &emailOpts.Identity,
				Name:        "identity",
				Aliases:     []string{"idn"},
				Value:       "",
				Usage:       "Senders email Identity if any",
				EnvVars:     []string{"EMAIL_IDENTITY"},
			},
			&cli.StringFlag{
				Destination: &emailOpts.Host,
				Name:        "host",
				Aliases:     []string{"ho"},
				Value:       "smtp.gmail.com",
				Usage:       "SMTP Host",
				EnvVars:     []string{"EMAIL_HOST"},
			},
			&cli.StringFlag{
				Destination: &emailOpts.Port,
				Name:        "port",
				Value:       "587",
				Aliases:     []string{"po"},
				Usage:       "SMTP Port",
				EnvVars:     []string{"EMAIL_PORT"},
			},
			&cli.StringFlag{
				Destination: &emailOpts.Message,
				Name:        "msg",
				Usage:       "Message content",
				Aliases:     []string{"m"},
				EnvVars:     []string{"EMAIL_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &emailOpts.Subject,
				Name:        "sub",
				Value:       TimeValue,
				Usage:       "Subject of the email",
				EnvVars:     []string{"EMAIL_SUBJECT"},
			},
		},
		Action: func(ctx *cli.Context) error {
			notifier := notify.New()
			emailSvc := mail.New(emailOpts.SenderAddress, emailOpts.Host+":"+emailOpts.Port)
			emailSvc.AuthenticateSMTP(emailOpts.Identity, emailOpts.SenderAddress, emailOpts.Password, emailOpts.Host)

			chn := strings.Split(emailOpts.ReceiverAddress, ",")
			for _, v := range chn {
				if len(v) <= 0 {
					return fmt.Errorf(EmptyChannel)
				}

				emailSvc.AddReceivers(v)
			}

			notifier.UseServices(emailSvc)

			if err := notifier.Send(
				context.Background(),
				emailOpts.Subject,
				emailOpts.Message,
			); err != nil {
				return err
			}
			log.Println("Successfully sent!")
			return nil
		},
	}
}
