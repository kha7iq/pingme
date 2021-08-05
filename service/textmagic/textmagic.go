package textmagic

import (
	"context"
	"fmt"
	"strings"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/textmagic"
	"github.com/urfave/cli/v2"
)

//TextMagic struct describes required data needed to integrate with TextMagic
type TextMagic struct {
	Token    string
	User     string
	Subject  string
	Message  string
	Receiver string
}

//Send method sends a message via TextMagic service
func Send() *cli.Command {
	var textMagicOpts TextMagic
	return &cli.Command{
		Name:  "textmagic",
		Usage: "Send message via TextMagic",
		UsageText: "pingme textmagic --token 'tokenabc' --user 'sid123' " +
			"--subject 'foo' --receiver '+140001442' --msg 'some message'",
		Description: `textmagic provides ability to send sms to multiple numbers.
You can specify multiple receivers by separating the value with a comma.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &textMagicOpts.Token,
				Name:        "token",
				Usage:       "TextMagic token",
				Aliases:     []string{"t"},
				Required:    true,
				EnvVars:     []string{"TEXTMAGIC_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &textMagicOpts.User,
				Name:        "user",
				Usage:       "TextMagic user",
				Aliases:     []string{"u"},
				Required:    true,
				EnvVars:     []string{"TEXTMAGIC_USER"},
			},
			&cli.StringFlag{
				Destination: &textMagicOpts.Subject,
				Name:        "title",
				Usage:       "Title of the message",
				EnvVars:     []string{"TEXTMAGIC_TITLE"},
			},
			&cli.StringFlag{
				Destination: &textMagicOpts.Receiver,
				Name:        "receiver",
				Usage:       "Receiver(s) of the message",
				Aliases:     []string{"r"},
				Required:    true,
				EnvVars:     []string{"TEXTMAGIC_RECEIVER"},
			},
			&cli.StringFlag{
				Destination: &textMagicOpts.Message,
				Name:        "msg",
				Usage:       "Message to send",
				Aliases:     []string{"m"},
				Required:    true,
				EnvVars:     []string{"TEXTMAGIC_MESSAGE"},
			},
		},
		Action: func(c *cli.Context) error {
			textMagicService := textmagic.New(textMagicOpts.User, textMagicOpts.Token)
			receivers, err := getReceivers(textMagicOpts.Receiver)
			if err != nil {
				return fmt.Errorf("invalid receivers provided, %w", err)
			}
			textMagicService.AddReceivers(receivers...)

			notifier := notify.New()
			notifier.UseServices(textMagicService)

			err = notifier.Send(context.Background(), textMagicOpts.Subject, textMagicOpts.Message)
			if err != nil {
				return fmt.Errorf("could not send textMagic message, %w", err)
			}
			return nil
		},
	}
}

func getReceivers(receivers string) ([]string, error) {
	if len(receivers) == 0 {
		return nil, fmt.Errorf("no receivers found")
	}
	r := strings.Split(receivers, ",")
	return r, nil
}
