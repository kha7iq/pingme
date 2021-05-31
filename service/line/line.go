package line

import (
	"context"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/line"
	"github.com/urfave/cli/v2"
)

// Line struct holds data parsed via flags for the service
type Line struct {
	ChanSecret      string
	ChanAccessToken string
	Message         string
	Receivers       string
	Title           string
}

// Send parses values from *cli.context and returns a *cli.Command.
// Values include channel secret, channel access token, receiver IDs (group or user), Message and Title.
// If multiple receiver IDs are provided, then the string is split with "," separator and each receiver ID is added to the receiver.
func Send() *cli.Command {
	var lineOpts Line
	return &cli.Command{
		Name:        "line",
		Usage:       "Send message to line messenger",
		Description: `Line messenger uses a channel secret and a channel access token to authenticate & send messages through line to various receivers. Multiple receiver IDs can be used separated by comma ','. All configuration options are also available via environment variables.`,
		UsageText:   "pingme line --channelsecret '123' --channelaccesstoken '123' --msg 'some message' --receivers '123,456,789'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &lineOpts.ChanSecret,
				Name:        "channelsecret",
				Required:    true,
				Usage:       "Channel secret.",
				EnvVars:     []string{"LINE_CHANNELSECRET"},
			},
			&cli.StringFlag{
				Destination: &lineOpts.ChanAccessToken,
				Name:        "channelaccesstoken",
				Required:    true,
				Usage:       "Channel access token.",
				EnvVars:     []string{"LINE_CHANNELACCESSTOKEN"},
			},
			&cli.StringFlag{
				Destination: &lineOpts.Message,
				Name:        "msg",
				Required:    true,
				Usage:       "Message content.",
				EnvVars:     []string{"LINE_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &lineOpts.Title,
				Name:        "title",
				Value:       helpers.TimeValue,
				Usage:       "Message title.",
				EnvVars:     []string{"LINE_MSG_TITLE"},
			},
			&cli.StringFlag{
				Destination: &lineOpts.Receivers,
				Name:        "receivers",
				Required:    true,
				Usage:       "Comma-separated list of user or group receiver IDs.",
				EnvVars:     []string{"LINE_RECEIVER_IDS"},
			},
		},
		Action: func(ctx *cli.Context) error {
			notifier := notify.New()
			lineSvc, err := line.New(lineOpts.ChanSecret, lineOpts.ChanAccessToken)
			if err != nil {
				return err
			}

			// Add receiver IDs
			recv := strings.Split(lineOpts.Receivers, ",")
			for _, r := range recv {
				lineSvc.AddReceivers(r)
			}

			notifier.UseServices(lineSvc)

			if err := notifier.Send(
				context.Background(),
				lineOpts.Title,
				lineOpts.Message,
			); err != nil {
				return err
			}

			log.Println("Successfully sent!")

			return nil
		},
	}
}
