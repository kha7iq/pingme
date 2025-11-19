package zulip

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

// Zulip holds all the necessary options to use zulip
type Zulip struct {
	ZBot
	Type    string
	To      string
	Topic   string
	Content string
	Domain  string
}

type ZBot struct {
	EmailID string
	APIKey  string
}

type ZResponse struct {
	ID      int    `json:"id"`
	Message string `json:"msg"`
	Result  string `json:"result"`
	Code    string `json:"code"`
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

// SendMessage sends a message to zulip stream or private user.
func SendMessage(domain, botEmail, apiKey, msgType, to, topic, content string) error {
	if domain == "" {
		return fmt.Errorf("zulip domain is required")
	}
	if botEmail == "" {
		return fmt.Errorf("zulip bot email is required")
	}
	if apiKey == "" {
		return fmt.Errorf("zulip API key is required")
	}
	if to == "" {
		return fmt.Errorf("zulip 'to' field is required")
	}
	if content == "" {
		return fmt.Errorf("message content is required")
	}

	initialize()

	zulipOpts := Zulip{
		ZBot: ZBot{
			EmailID: botEmail,
			APIKey:  apiKey,
		},
		Type:    msgType,
		To:      to,
		Topic:   topic,
		Content: content,
		Domain:  domain,
	}

	resp, err := SendZulipMessage(domain, zulipOpts)
	if err != nil {
		return err
	}

	if resp.Result == "success" {
		log.Printf("Successfully sent! Server Reply ID: %v\nResult: %v\n", resp.ID, resp.Result)
		return nil
	}

	return fmt.Errorf("failed to send: %s", resp.Message)
}

func Send() *cli.Command {
	var zulipOpts Zulip
	return &cli.Command{
		Name:  "zulip",
		Usage: "Send message to zulip",
		UsageText: "pingme zulip --email 'john.doe@email.com' --api-key '12345567' --to 'london' --type 'stream' " +
			"--topic 'some topic' --content 'content of the message'",
		Description: `Zulip uses token and email to authenticate and ids for users or streams.
You can specify multiple userIds by separating the value with ','.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &zulipOpts.Domain,
				Name:        "domain",
				Aliases:     []string{},
				Required:    true,
				Usage:       "Your zulip domain",
				EnvVars:     []string{"ZULIP_DOMAIN"},
			},
			&cli.StringFlag{
				Destination: &zulipOpts.EmailID,
				Name:        "email",
				Aliases:     []string{},
				Required:    true,
				Usage:       "Email ID of the bot",
				EnvVars:     []string{"ZULIP_BOT_EMAIL_ADDRESS"},
			},
			&cli.StringFlag{
				Destination: &zulipOpts.APIKey,
				Name:        "api-key",
				Aliases:     []string{},
				Required:    true,
				Usage:       "API Key of the bot",
				EnvVars:     []string{"ZULIP_BOT_API_KEY"},
			},
			&cli.StringFlag{
				Destination: &zulipOpts.Type,
				Name:        "type",
				Aliases:     []string{},
				Usage:       "The type of message to be sent. private for a private message and stream for a stream message.",
				EnvVars:     []string{"ZULIP_MSG_TYPE"},
			},
			&cli.StringFlag{
				Destination: &zulipOpts.To,
				Name:        "to",
				Aliases:     []string{},
				Usage:       "For stream messages, the name of the stream. For private messages, csv of email addresses",
				EnvVars:     []string{"ZULIP_STREAM_NAME"},
			},
			&cli.StringFlag{
				Destination: &zulipOpts.Topic,
				Name:        "topic",
				Aliases:     []string{},
				Usage:       "The topic of the message. Only required for stream messages 'type=stream', ignored otherwise.",
				EnvVars:     []string{"ZULIP_TOPIC"},
			},
			&cli.StringFlag{
				Destination: &zulipOpts.Content,
				Name:        "msg",
				Aliases:     []string{},
				Required:    true,
				Usage:       "The content of the message.",
				EnvVars:     []string{"ZULIP_MESSAGE"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return SendMessage(
				zulipOpts.Domain,
				zulipOpts.EmailID,
				zulipOpts.APIKey,
				zulipOpts.Type,
				zulipOpts.To,
				zulipOpts.Topic,
				zulipOpts.Content,
			)
		},
	}
}

func getTo(messageType string, to string) string {
	if messageType == "stream" {
		return to
	}
	privateTo, _ := json.Marshal(strings.Split(to, ","))
	return string(privateTo)
}

// SendZulipMessage function takes the zulip domain and zulip bot
// type, to, topic and content in the form of json byte array and sends
// message to zulip.
func SendZulipMessage(zulipDomain string, zulipOpts Zulip) (*ZResponse, error) {
	data := url.Values{}
	data.Set("type", zulipOpts.Type)
	data.Set("to", getTo(zulipOpts.Type, zulipOpts.To))
	data.Set("topic", zulipOpts.Topic)
	data.Set("content", zulipOpts.Content)

	var response ZResponse

	endPointURL := "https://" + zulipDomain + "/api/v1/messages"
	req, err := http.NewRequest("POST", endPointURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	zulipBot := zulipOpts.ZBot
	req.SetBasicAuth(zulipBot.EmailID, zulipBot.APIKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
