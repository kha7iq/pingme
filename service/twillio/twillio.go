package twillio

import (
	"fmt"
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

// SendMessage sends SMS via twillio to multiple receivers.
// receivers can be comma-separated string of phone numbers.
func SendMessage(accountSID, token, sender, receivers, title, message string) error {
	if accountSID == "" {
		return fmt.Errorf("twillio account SID is required")
	}
	if token == "" {
		return fmt.Errorf("twillio token is required")
	}
	if sender == "" {
		return fmt.Errorf("twillio sender phone number is required")
	}
	if receivers == "" {
		return fmt.Errorf("twillio receiver phone number is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	client := gotwilio.NewTwilioClient(accountSID, token)
	fullMessage := title + "\n" + message

	numbers := strings.Split(receivers, ",")
	for _, phoneNumber := range numbers {
		phoneNumber = strings.TrimSpace(phoneNumber)
		if len(phoneNumber) == 0 {
			return helpers.ErrChannel
		}

		_, exception, err := client.SendSMS(sender, phoneNumber, fullMessage, "", "")
		if err != nil {
			return fmt.Errorf("failed to send SMS to %s: %w", phoneNumber, err)
		}
		if exception != nil {
			return fmt.Errorf("twillio exception for %s: %v", phoneNumber, exception)
		}
	}

	log.Println("Successfully sent!")
	return nil
}

// Send parse values from *cli.context and return *cli.Command
// and send messages to target numbers.
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
			return SendMessage(
				twillioOpts.AccountSid,
				twillioOpts.Token,
				twillioOpts.Sender,
				twillioOpts.Receiver,
				twillioOpts.Title,
				twillioOpts.Message,
			)
		},
	}
}
