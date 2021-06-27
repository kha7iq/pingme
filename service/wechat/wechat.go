package wechat

import (
	"context"
	"log"
	"strings"

	"github.com/kha7iq/pingme/service/helpers"
	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/wechat"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/urfave/cli/v2"
)

// Wechat struct holds data parsed via flags for the service.
type Wechat struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	Title          string
	Message        string
	Receivers      string
}

// Send parse values from *cli.context and return *cli.Command.
// Values include wechat official account id, secret, server token, encoding AES key,
// Message, Title, and Receivers.
// If multiple receivers are provided then the string is split with "," separator and
// each receiverID is added to receiver.
func Send() *cli.Command {
	var wechatOpts Wechat
	return &cli.Command{
		Name:  "wechat",
		Usage: "Send message to wechat official account",
		Description: `Wechat sends message to Wechat Official Account using appid, appsecrete 
and server token to authenticate 
AND then send messages to defined account. 
Multiple receiverss can be used separated by comma.`,
		UsageText: "pingme wechat --appid '123' --appsecret '123' --token '123' --aes '123' --msg 'some message'  --receivers 'aaa,bbb,ccc'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &wechatOpts.AppID,
				Name:        "appid",
				Required:    true,
				Usage:       "AppID of wechat official account.",
				EnvVars:     []string{"WECHAT_APPID"},
			},
			&cli.StringFlag{
				Destination: &wechatOpts.AppSecret,
				Name:        "appsecret",
				Required:    true,
				Usage:       "AppSecret of wechat official account.",
				EnvVars:     []string{"WECHAT_APPSECRET"},
			},
			&cli.StringFlag{
				Destination: &wechatOpts.Token,
				Name:        "token",
				Required:    true,
				Usage:       "Token of server used for sending message.",
				EnvVars:     []string{"WECHAT_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &wechatOpts.EncodingAESKey,
				Name:        "aes",
				Required:    true,
				Usage:       "Encoding AES Key of server used for sending message.",
				EnvVars:     []string{"WECHAT_ENCODING_AES_KEY"},
			},
			&cli.StringFlag{
				Destination: &wechatOpts.Receivers,
				Name:        "receivers",
				Required:    true,
				Usage:       "Comma-separated list of receiver IDs.",
				EnvVars:     []string{"WECHAT_RECEIVERS"},
			},
			&cli.StringFlag{
				Destination: &wechatOpts.Message,
				Name:        "msg",
				Required:    true,
				Usage:       "Message content.",
				EnvVars:     []string{"WECHAT_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &wechatOpts.Title,
				Name:        "title",
				Value:       helpers.TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"WECHAT_TITLE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			wechatSvc := wechat.New(&wechat.Config{
				AppID:          wechatOpts.AppID,
				AppSecret:      wechatOpts.AppSecret,
				Token:          wechatOpts.Token,
				EncodingAESKey: wechatOpts.EncodingAESKey,
				Cache:          cache.NewMemory(),
			})

			// Add receiver IDs
			recv := strings.Split(wechatOpts.Receivers, ",")
			for _, r := range recv {
				wechatSvc.AddReceivers(r)
			}

			notifier := notify.New()
			notifier.UseServices(wechatSvc)

			err := notifier.Send(context.Background(), wechatOpts.Title, wechatOpts.Message)
			if err != nil {
				log.Fatalf("notifier.Send() failed: %s", err.Error())
			}

			log.Println("Successfully sent!")

			return nil
		},
	}
}
