package pushover

import (
	"fmt"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"

	"github.com/gregdel/pushover"
	"github.com/urfave/cli/v2"
)

// pushOver struct holds data parsed via flags for pushover service
type pushOver struct {
	Token     string
	Recipient string
	Message   string
	Title     string
	Priority  int
}

// SendMessage sends a message to pushover users.
// This is the core logic extracted for reuse by both CLI and webhook.
// recipients can be comma-separated string of user tokens.
func SendMessage(token, recipients, title, message string, priority int) error {
	if token == "" {
		return fmt.Errorf("pushover token is required")
	}
	if recipients == "" {
		return fmt.Errorf("pushover user token is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	app := pushover.New(token)

	msg := &pushover.Message{
		Title:    title,
		Message:  message,
		Priority: priority,
		Retry:    60,
		Expire:   3600,
	}

	users := strings.Split(recipients, ",")

	for _, userToken := range users {
		userToken = strings.TrimSpace(userToken)
		if len(userToken) == 0 {
			return helpers.ErrChannel
		}
		recipient := pushover.NewRecipient(userToken)
		responsePushOver, err := app.SendMessage(msg, recipient)
		if err != nil {
			return fmt.Errorf("failed to send to user %s: %w", userToken, err)
		}
		log.Printf("Successfully sent to %s!\n%v\n", userToken, responsePushOver)
	}
	return nil
}

// Send parse values from *cli.context and return *cli.Command.
// Values include token, users, Message and Title.
// If multiple users are provided then the string is split with "," separator and
// each user is added to receiver.
func Send() *cli.Command {
	var pushOverOpts pushOver
	return &cli.Command{
		Name:      "pushover",
		Usage:     "Send message to pushover",
		UsageText: "pingme pushover --token '123' --user '12345,567' --msg 'some message'",
		Description: `Pushover uses token to authenticate application and user token to send messages to the user.
All configuration options are also available via environment variables.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &pushOverOpts.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "Token of pushover application used for authenticate application.",
				EnvVars:     []string{"PUSHOVER_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &pushOverOpts.Recipient,
				Name:        "user",
				Required:    true,
				Aliases:     []string{"u"},
				Usage:       "User token used for sending message to user,if sending to multiple users separate with ','.",
				EnvVars:     []string{"PUSHOVER_USER"},
			},
			&cli.StringFlag{
				Destination: &pushOverOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content.",
				EnvVars:     []string{"PUSHOVER_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &pushOverOpts.Title,
				Name:        "title",
				Value:       helpers.TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"PUSHOVER_TITLE"},
			},
			&cli.IntFlag{
				Destination: &pushOverOpts.Priority,
				Name:        "priority",
				Aliases:     []string{"p"},
				Value:       0,
				Usage:       "Priority of the message.",
				EnvVars:     []string{"PUSHOVER_PRIORITY"},
			},
		},
		Action: func(ctx *cli.Context) error {
			// Now just call the extracted function
			return SendMessage(
				pushOverOpts.Token,
				pushOverOpts.Recipient,
				pushOverOpts.Title,
				pushOverOpts.Message,
				pushOverOpts.Priority,
			)
		},
	}
}
