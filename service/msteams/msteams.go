package msteams

import (
	"context"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"

	"github.com/nikoksr/notify"
	msteams2 "github.com/nikoksr/notify/service/msteams"
	"github.com/urfave/cli/v2"
)

// msTeams struct holds data parsed via flags for microsoft teams service.
type msTeams struct {
	Webhook string
	Message string
	Title   string
}

// Send parse values from *cli.context and return *cli.Command.
// Values include Ms Teams Webhook, Message and Title.
// If multiple webhooks are provided then the string is split with "," separator and
// each webhook is added to receiver.
func Send() *cli.Command {
	var msTeamOpt msTeams
	return &cli.Command{
		Name:  "teams",
		Usage: "Send message to microsoft teams",
		Description: `Teams uses webhooks to send messages, you can add multiple webhooks separated by comma ',' or 
you can add permissions for multiple channels to single webhook.`,
		UsageText: "pingme teams --webhook 'https://example.webhook.office.com/xx' --msg 'some message'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &msTeamOpt.Webhook,
				Name:        "webhook",
				Aliases:     []string{"w"},
				Required:    true,
				Usage:       "Webhook associated with teams channel, if multiple separate with ','.",
				EnvVars:     []string{"TEAMS_WEBHOOK"},
			},
			&cli.StringFlag{
				Destination: &msTeamOpt.Message,
				Name:        "msg",
				Usage:       "Message content.",
				EnvVars:     []string{"TEAMS_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &msTeamOpt.Title,
				Name:        "title",
				Value:       helpers.TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"TEAMS_MSG_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			notifier := notify.New()
			teamsSvc := msteams2.New()

			chn := strings.Split(msTeamOpt.Webhook, ",")
			for _, v := range chn {
				if len(v) <= 0 {
					return helpers.ErrChannel
				}
				teamsSvc.AddReceivers(v)
			}

			notifier.UseServices(teamsSvc)

			if err := notifier.Send(
				context.Background(),
				msTeamOpt.Title,
				msTeamOpt.Message,
			); err != nil {
				return err
			}
			log.Println("Successfully sent!")
			return nil
		},
	}
}
