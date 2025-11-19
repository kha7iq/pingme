package mattermost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kha7iq/pingme/service/helpers"

	"github.com/urfave/cli/v2"
)

// matterMost struct holds data parsed via flags for the service
type matterMost struct {
	Title     string
	Token     string
	ServerURL string
	Scheme    string
	APIURL    string
	Message   string
	ChanIDs   string
}

// matterMostResponse struct holds the server responses
type matterMostResponse struct {
	ID         string `json:"id"`
	CreateAt   int64  `json:"create_at"`
	UpdateAt   int64  `json:"update_at"`
	EditAt     int    `json:"edit_at"`
	DeleteAt   int    `json:"delete_at"`
	IsPinned   bool   `json:"is_pinned"`
	UserID     string `json:"user_id"`
	ChannelID  string `json:"channel_id"`
	RootID     string `json:"root_id"`
	ParentID   string `json:"parent_id"`
	OriginalID string `json:"original_id"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	Props      struct {
		FromBot string `json:"from_bot"`
	} `json:"props"`
	Hashtags      string   `json:"hashtags"`
	PendingPostID string   `json:"pending_post_id"`
	ReplyCount    int      `json:"reply_count"`
	Metadata      struct{} `json:"metadata"`
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

// SendMessage sends a message to mattermost channels.
// channels can be comma-separated string of channel IDs.
func SendMessage(token, serverURL, scheme, apiURL, channels, title, message string) error {
	if token == "" {
		return fmt.Errorf("mattermost token is required")
	}
	if serverURL == "" {
		return fmt.Errorf("mattermost server URL is required")
	}
	if channels == "" {
		return fmt.Errorf("mattermost channel is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	initialize()

	endPointURL := scheme + "://" + serverURL + apiURL
	bearer := "Bearer " + token
	fullMessage := title + "\n" + message

	ids := strings.Split(channels, ",")
	for _, channelID := range ids {
		channelID = strings.TrimSpace(channelID)
		if len(channelID) == 0 {
			return helpers.ErrChannel
		}

		jsonData, err := toJSON(channelID, fullMessage)
		if err != nil {
			return fmt.Errorf("error parsing json: %w", err)
		}

		if err := sendMattermost(endPointURL, bearer, jsonData); err != nil {
			return fmt.Errorf("failed to send message to channel %s: %w", channelID, err)
		}
	}
	return nil
}

// Send parse values from *cli.context and return *cli.Command
// and send messages to target channels.
func Send() *cli.Command {
	var mattermostOpts matterMost
	return &cli.Command{
		Name:  "mattermost",
		Usage: "Send message to mattermost",
		UsageText: "pingme mattermost --token '123' --channel '12345567' --url 'localhost' --scheme http " +
			"--msg 'some message'",
		Description: `Mattermost uses token to authenticate and channel ids for targets.
You can specify multiple channels by separating the value with ','.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &mattermostOpts.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "Personal access token or Bot token for authorization.",
				EnvVars:     []string{"MATTERMOST_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &mattermostOpts.ChanIDs,
				Name:        "channel",
				Required:    false,
				Aliases:     []string{"c"},
				Usage:       "Channel IDs, if sending to multiple channels separate with ','.",
				EnvVars:     []string{"MATTERMOST_CHANNELS"},
			},
			&cli.StringFlag{
				Destination: &mattermostOpts.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Usage:       "Message content.",
				EnvVars:     []string{"MATTERMOST_MESSAGE"},
			},
			&cli.StringFlag{
				Destination: &mattermostOpts.Title,
				Name:        "title",
				Value:       helpers.TimeValue,
				Usage:       "Title of the message.",
				EnvVars:     []string{"MATTERMOST_TITLE"},
			},
			&cli.StringFlag{
				Destination: &mattermostOpts.ServerURL,
				Name:        "url",
				Aliases:     []string{"u"},
				Value:       "localhost",
				Required:    true,
				Usage:       "URL of your mattermost server i.e example.com",
				EnvVars:     []string{"MATTERMOST_SERVER_URL"},
			},
			&cli.StringFlag{
				Destination: &mattermostOpts.Scheme,
				Name:        "scheme",
				Value:       "https",
				Usage:       "For server with no tls chose http, by default it uses https",
				EnvVars:     []string{"MATTERMOST_SCHEME"},
			},
			&cli.StringFlag{
				Destination: &mattermostOpts.APIURL,
				Name:        "api",
				Value:       "/api/v4/posts",
				Usage:       "Unless using older version of api default is fine.",
				EnvVars:     []string{"MATTERMOST_API_URL"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return SendMessage(
				mattermostOpts.Token,
				mattermostOpts.ServerURL,
				mattermostOpts.Scheme,
				mattermostOpts.APIURL,
				mattermostOpts.ChanIDs,
				mattermostOpts.Title,
				mattermostOpts.Message,
			)
		},
	}
}

// toJSON takes strings and convert them to json byte array
func toJSON(channel string, msg string) ([]byte, error) {
	m := make(map[string]string, 2)
	m["channel_id"] = channel
	m["message"] = msg
	js, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return js, nil
}

// sendMattermost function take the server url, authentication token
// message and channel id in the form of json byte array and sends
// message to mattermost.
func sendMattermost(url string, token string, jsonPayload []byte) error {
	var response matterMostResponse

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
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

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}
	log.Printf("Server Reply: %v\nMessage: %v\n", response.ID, response.Message)

	return nil
}
