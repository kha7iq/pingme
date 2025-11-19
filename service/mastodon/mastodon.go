package mastodon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/urfave/cli/v2"
)

// Mastodon struct holds data parsed via flags for the service
type Mastodon struct {
	Title     string
	Token     string
	ServerURL string
	Message   string
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var Client HTTPClient

func initialize() {
	Client = &http.Client{
		Timeout: 10 * time.Second,
	}
}

// SendMessage sends a status message to mastodon.
func SendMessage(token, serverURL, title, message string) error {
	if token == "" {
		return fmt.Errorf("mastodon token is required")
	}
	if serverURL == "" {
		return fmt.Errorf("mastodon server URL is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	initialize()

	endPointURL := "https://" + serverURL + "/api/v1/statuses/"
	bearer := "Bearer " + token
	fullMessage := title + "\n" + message

	if err := sendMastodon(endPointURL, bearer, fullMessage); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

// Send parse values from *cli.context and return *cli.Command
// and sets a status message for mastodon.
func Send() *cli.Command {
	var mastodonOpts Mastodon
	return &cli.Command{
		Name:  "mastodon",
		Usage: "Set status message for mastodon",
		UsageText: "pingme mastodon --token '123' --url 'mastodon.social' --title 'PingMe'  " +
			"--msg 'some message'",
		Description: `Mastodon uses application token to authorize and sets a status message`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &mastodonOpts.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "Application token for authorization.",
				EnvVars:     []string{"MASTODON_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &mastodonOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content.",
				EnvVars:     []string{"MASTODON_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &mastodonOpts.Title,
				Name:        "title",
				Usage:       "Title of the message.",
				EnvVars:     []string{"MASTODON_TITLE"},
			},
			&cli.StringFlag{
				Destination: &mastodonOpts.ServerURL,
				Name:        "url",
				Aliases:     []string{"u"},
				Value:       "mastodon.social",
				Required:    true,
				Usage:       "URL of mastodon server i.e mastodon.social",
				EnvVars:     []string{"MASTODON_SERVER"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return SendMessage(
				mastodonOpts.Token,
				mastodonOpts.ServerURL,
				mastodonOpts.Title,
				mastodonOpts.Message,
			)
		},
	}
}

// sendMastodon function take the server url, authorization token
// and message string to set the status.
func sendMastodon(url string, token string, msg string) error {
	reqBody, err := json.Marshal(map[string]string{
		"status": msg,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	checkErr, ok := data["error"]
	if ok {
		return fmt.Errorf("%v", checkErr)
	}

	log.Printf("Success!!\nVisibility: %v\nURL: %v\n", data["visibility"], data["url"])
	return nil
}
