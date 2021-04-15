package cmd

import (
	"log"

	"github.com/gregdel/pushover"
	"github.com/urfave/cli/v2"
)

// slackPingMe struct holds data parsed via flags for slack service.
type pushoverPingMe struct {
	Token     string
	Recipient string
	Message   string
}

// SendToSlack parse values from *cli.context and return *cli.Command.
// Values include slack token, channelIDs, Message and Title.
// If multiple channels are provided then the string is split with "," separator and
// each channelID is added to receiver.
func SendToPushover() *cli.Command {
	var pushoverOpts pushoverPingMe
	return &cli.Command{
		Name:      "pushover",
		Usage:     "Send message to pushover",
		UsageText: "pingme pushover --token '123' --user '12345' --message 'some message'",
		Description: `Pushover uses token to authenticate application and user token to  send messages to the user.
All configuration options are also available via environment variables.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &pushoverOpts.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "Token of pushover application used for authenticate application.",
				EnvVars:     []string{"PUSHOVER_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &pushoverOpts.Recipient,
				Name:        "user",
				Required:    true,
				Aliases:     []string{"u"},
				Usage:       "User token used for sending message to user.",
				EnvVars:     []string{"PUSHOVER_USER"},
			},
			&cli.StringFlag{
				Destination: &pushoverOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content.",
				EnvVars:     []string{"PUSHOVER_MESSAGE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			app := pushover.New(pushoverOpts.Token)
			recipient := pushover.NewRecipient(pushoverOpts.Recipient)
			message := pushover.NewMessage(pushoverOpts.Message)
			_, err := app.SendMessage(message, recipient)
			if err != nil {
				return err
			}
			log.Println("Successfully sent!")
			return nil
		},
	}
}
