package pushover

import (
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
}

// Send parse values from *cli.context and return *cli.Command.
// Values include  token, users, Message and Title.
// If multiple users are provided then the string is split with "," separator and
// each user is added to receiver.
func Send() *cli.Command {
	var pushOverOpts pushOver
	return &cli.Command{
		Name:      "pushover",
		Usage:     "Send message to pushover",
		UsageText: "pingme pushover --token '123' --user '12345,567' --msg 'some message'",
		Description: `Pushover uses token to authenticate application and user token to  send messages to the user.
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
				Usage:       "User token used for sending message to user,if sending to multiple userss separate with ','.",
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
		},
		Action: func(ctx *cli.Context) error {
			app := pushover.New(pushOverOpts.Token)
			message := pushover.NewMessageWithTitle(pushOverOpts.Message, pushOverOpts.Title)
			users := strings.Split(pushOverOpts.Recipient, ",")

			for _, v := range users {
				if len(v) == 0 {
					return helpers.ErrChannel
				}
				recipient := pushover.NewRecipient(v)
				responsePushOver, err := app.SendMessage(message, recipient)
				if err != nil {
					return err
				}
				log.Printf("Successfully sent!\n%v\n", responsePushOver)
			}
			return nil
		},
	}
}
