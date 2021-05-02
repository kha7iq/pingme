package twillio

import (
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"
	"github.com/sfreiberg/gotwilio"
	"github.com/urfave/cli/v2"
)

// Twillio struct holds data parsed via flags for the service
type Twillio struct {
	Title      string
	Token      string
	AccountSid string
	Sender     string
	Receiver   string
	Message    string
}

// Send parse values from *cli.context and return *cli.Command
// and send messages to target numbers.
// If multiple receivers are provided then the string is split with "," separator and
// message is sent to each number.
func Send() *cli.Command {
	var twillioOpts Twillio
	return &cli.Command{
		Name:  "twillio",
		Usage: "Send sms via twillio",
		UsageText: "pingme twillio --token 'tokenabc' --account 'sid123' " +
			"--sender '+140001442' --receiver '+140001442'' --msg 'some message'",
		Description: `Twillio provides ability to send sms to multiple numbers.
You can specify multiple receivers by separating the value with a comma.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &twillioOpts.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "Auth token for twillio account.",
				EnvVars:     []string{"TWILLIO_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &twillioOpts.AccountSid,
				Name:        "account",
				Required:    true,
				Aliases:     []string{"a"},
				Usage:       "Twillio account sid",
				EnvVars:     []string{"TWILLIO_ACCOUNT_SID"},
			},
			&cli.StringFlag{
				Destination: &twillioOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content.",
				EnvVars:     []string{"TWILLIO_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &twillioOpts.Title,
				Name:        "title",
				Usage:       "Title of the message.",
				EnvVars:     []string{"TWILLIO_TITLE"},
			},
			&cli.StringFlag{
				Destination: &twillioOpts.Sender,
				Name:        "sender",
				Aliases:     []string{"s"},
				Usage:       "Sender's phone number",
				EnvVars:     []string{"TWILLIO_SENDER"},
			},
			&cli.StringFlag{
				Destination: &twillioOpts.Receiver,
				Name:        "receiver",
				Aliases:     []string{"r"},
				Usage:       "Receiver's phone number",
				EnvVars:     []string{"TWILLIO_RECEIVER"},
			},
		},
		Action: func(ctx *cli.Context) error {
			client := gotwilio.NewTwilioClient(twillioOpts.AccountSid, twillioOpts.Token)
			fullMessage := twillioOpts.Title + "\n" + twillioOpts.Message

			numbers := strings.Split(twillioOpts.Receiver, ",")
			for _, v := range numbers {
				if len(v) == 0 {
					return helpers.ErrChannel
				}

				_, exception, err := client.SendSMS(twillioOpts.Sender, twillioOpts.Receiver, fullMessage, "", "")
				if err != nil {
					return err
				}
				if exception != nil {
					return exception
				}
			}
			log.Println("Successfully sent!")
			return nil
		},
	}
}
