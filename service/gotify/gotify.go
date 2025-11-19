package gotify

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/kha7iq/pingme/service/helpers"

	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/client/message"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
	"github.com/urfave/cli/v2"
)

// Gotify struct holds data parsed via flags for the service
type Gotify struct {
	URL      string
	Token    string
	Priority int
	Title    string
	Message  string
}

// SendMessage sends a message to gotify server.
func SendMessage(serverURL, token, title, msg string, priority int) error {
	if serverURL == "" {
		return fmt.Errorf("gotify server URL is required")
	}
	if token == "" {
		return fmt.Errorf("gotify token is required")
	}
	if msg == "" {
		return fmt.Errorf("message is required")
	}

	parsedURL, err := url.Parse(serverURL)
	if err != nil {
		return fmt.Errorf("invalid gotify URL: %w", err)
	}

	client := gotify.NewClient(parsedURL, &http.Client{})
	params := message.NewCreateMessageParams()
	params.Body = &models.MessageExternal{
		Title:    title,
		Message:  msg,
		Priority: priority,
	}

	_, err = client.Message.CreateMessage(params, auth.TokenAuth(token))
	if err != nil {
		return fmt.Errorf("failed to send gotify message: %w", err)
	}

	log.Println("Successfully sent!")
	return nil
}

// Send parse values from *cli.context and return *cli.Command
func Send() *cli.Command {
	var gotifyOpts Gotify
	return &cli.Command{
		Name:  "gotify",
		Usage: "Send push notification to gotify server",
		UsageText: "pingme gotify  --url 'https://example.com' --token 'tokenabc' --title 'some title' " +
			" --msg 'some message' --priority 5",
		Description: `With gotify you can send messages to any Gotify server`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &gotifyOpts.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "Application token of gotify server",
				EnvVars:     []string{"GOTIFY_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &gotifyOpts.URL,
				Name:        "url",
				Required:    true,
				Aliases:     []string{"u"},
				Usage:       "Gotify server Endpoint",
				EnvVars:     []string{"GOTIFY_URL"},
			},
			&cli.StringFlag{
				Destination: &gotifyOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content",
				EnvVars:     []string{"GOTIFY_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &gotifyOpts.Title,
				Name:        "title",
				Usage:       "Title of the message.",
				Value:       helpers.TimeValue,
				EnvVars:     []string{"GOTIFY_TITLE"},
			},
			&cli.IntFlag{
				Destination: &gotifyOpts.Priority,
				Name:        "priority",
				Aliases:     []string{"p"},
				Usage:       "Message priority i.e 1-7",
				Value:       5,
				EnvVars:     []string{"GOTIFY_PRIORITY"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return SendMessage(
				gotifyOpts.URL,
				gotifyOpts.Token,
				gotifyOpts.Title,
				gotifyOpts.Message,
				gotifyOpts.Priority,
			)
		},
	}
}
