package whatsapp

import (
	"context"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/whatsapp"
	"github.com/urfave/cli/v2"
)

// whatsapp struct holds data parsed via flags for whatsapp service.
type whatsApp struct {
	Contact string
	Message string
	Title   string
}

// Send parse values from *cli.context and return *cli.Command.
// Values include whatsapp message, title, contacts.
// If multiple contacts are provided then the string is split with "," separator and
// each contact is added to receiver.
func Send() *cli.Command {
	var whatsAppOpts whatsApp
	return &cli.Command{
		Name:  "whatsapp",
		Usage: "Send message to whatsapp",
		Description: `WhatsApp uses authentication using QR code on terminal and sends messages to defined contacts.
Multiple contacts can be used separated by comma ','.
All configuration options are also available via environment variables.`,
		UsageText: "pingme whatsapp --contact 'contact1' --msg 'some message'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &whatsAppOpts.Contact,
				Name:        "contact",
				Required:    true,
				Aliases:     []string{"c"},
				Usage:       "WhatsApp contacts, if sending to multiple contacts separate with ','.",
				EnvVars:     []string{"WHATSAPP_CONTACT"},
			},
			&cli.StringFlag{
				Destination: &whatsAppOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content.",
				EnvVars:     []string{"WHATSAPP_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &whatsAppOpts.Title,
				Name:        "title",
				Value:       helpers.TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"WHATSAPP_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			whatsappSvc, err := whatsapp.New()
			if err != nil {
				return err
			}

			err = whatsappSvc.LoginWithQRCode()
			if err != nil {
				return err
			}

			contacts := strings.Split(whatsAppOpts.Contact, ",")
			for _, v := range contacts {
				if len(v) <= 0 {
					return helpers.ErrChannel
				}

				whatsappSvc.AddReceivers(v)
			}

			notifier := notify.New()
			notifier.UseServices(whatsappSvc)

			err = notifier.Send(context.Background(), whatsAppOpts.Title, whatsAppOpts.Message)
			if err != nil {
				return err
			}
			log.Println("Successfully sent!")
			return nil
		},
	}
}
