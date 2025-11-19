package dispatcher

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/kha7iq/pingme/internal/types"

	"github.com/kha7iq/pingme/service/discord"
	"github.com/kha7iq/pingme/service/email"
	"github.com/kha7iq/pingme/service/gotify"
	"github.com/kha7iq/pingme/service/line"
	"github.com/kha7iq/pingme/service/mastodon"
	"github.com/kha7iq/pingme/service/matrix"
	"github.com/kha7iq/pingme/service/mattermost"
	"github.com/kha7iq/pingme/service/pushbullet"
	"github.com/kha7iq/pingme/service/pushover"
	"github.com/kha7iq/pingme/service/rocketchat"
	"github.com/kha7iq/pingme/service/slack"
	"github.com/kha7iq/pingme/service/telegram"
	"github.com/kha7iq/pingme/service/twillio"
	"github.com/kha7iq/pingme/service/wechat"
	"github.com/kha7iq/pingme/service/zulip"
)

// Dispatcher routes webhook requests to appropriate services
type Dispatcher struct{}

// New creates a new dispatcher
func New() *Dispatcher {
	return &Dispatcher{}
}

// Dispatch sends the message to the specified service
func (d *Dispatcher) Dispatch(ctx context.Context, req *types.WebhookRequest) error {
	switch req.Service {
	case "pushover":
		return d.sendPushover(req)
	case "telegram":
		return d.sendTelegram(req)
	case "slack":
		return d.sendSlack(req)
	case "discord":
		return d.sendDiscord(req)
	case "email":
		return d.sendEmail(req)
	case "mattermost":
		return d.sendMattermost(req)
	case "rocketchat":
		return d.sendRocketChat(req)
	case "pushbullet":
		return d.sendPushbullet(req)
	case "twillio":
		return d.sendTwillio(req)
	case "zulip":
		return d.sendZulip(req)
	case "mastodon":
		return d.sendMastodon(req)
	case "line":
		return d.sendLine(req)
	case "wechat":
		return d.sendWeChat(req)
	case "gotify":
		return d.sendGotify(req)
	case "matrix":
		return d.sendMatrix(req)
	default:
		return fmt.Errorf("unsupported service: %s", req.Service)
	}
}

// sendPushover sends message via Pushover
func (d *Dispatcher) sendPushover(req *types.WebhookRequest) error {
	token := os.Getenv("PUSHOVER_TOKEN")
	user := os.Getenv("PUSHOVER_USER")

	if token == "" || user == "" {
		return fmt.Errorf("PUSHOVER_TOKEN and PUSHOVER_USER environment variables required")
	}

	priority := 0
	if req.Priority != 0 {
		priority = req.Priority
	}

	return pushover.SendMessage(token, user, req.Title, req.Message, priority)
}

// sendTelegram sends message via Telegram
func (d *Dispatcher) sendTelegram(req *types.WebhookRequest) error {
	token := os.Getenv("TELEGRAM_TOKEN")
	channels := os.Getenv("TELEGRAM_CHANNELS")

	if token == "" || channels == "" {
		return fmt.Errorf("TELEGRAM_TOKEN and TELEGRAM_CHANNELS environment variables required")
	}

	return telegram.SendMessage(token, channels, req.Title, req.Message)
}

// sendSlack sends message via Slack
func (d *Dispatcher) sendSlack(req *types.WebhookRequest) error {
	token := os.Getenv("SLACK_TOKEN")
	channels := os.Getenv("SLACK_CHANNELS")

	if token == "" || channels == "" {
		return fmt.Errorf("SLACK_TOKEN and SLACK_CHANNELS environment variables required")
	}

	return slack.SendMessage(token, channels, req.Title, req.Message)
}

// sendDiscord sends message via Discord
func (d *Dispatcher) sendDiscord(req *types.WebhookRequest) error {
	token := os.Getenv("DISCORD_TOKEN")
	channels := os.Getenv("DISCORD_CHANNELS")

	if token == "" || channels == "" {
		return fmt.Errorf("DISCORD_TOKEN and DISCORD_CHANNELS environment variables required")
	}

	return discord.SendMessage(token, channels, req.Title, req.Message)
}

// sendEmail sends message via Email
func (d *Dispatcher) sendEmail(req *types.WebhookRequest) error {
	sender := os.Getenv("EMAIL_SENDER")
	password := os.Getenv("EMAIL_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	receiver := os.Getenv("EMAIL_RECEIVER")

	if sender == "" || password == "" || host == "" || port == "" || receiver == "" {
		return fmt.Errorf("EMAIL_SENDER, EMAIL_PASSWORD, EMAIL_HOST, EMAIL_PORT, and EMAIL_RECEIVER environment variables required")
	}

	return email.SendMessage(sender, password, host, port, "", receiver, req.Title, req.Message)
}

// sendMattermost sends message via Mattermost
func (d *Dispatcher) sendMattermost(req *types.WebhookRequest) error {
	token := os.Getenv("MATTERMOST_TOKEN")
	serverURL := os.Getenv("MATTERMOST_SERVER_URL")
	channels := os.Getenv("MATTERMOST_CHANNELS")
	scheme := os.Getenv("MATTERMOST_SCHEME")
	if scheme == "" {
		scheme = "https"
	}

	if token == "" || serverURL == "" || channels == "" {
		return fmt.Errorf("MATTERMOST_TOKEN, MATTERMOST_SERVER_URL, and MATTERMOST_CHANNELS environment variables required")
	}

	return mattermost.SendMessage(token, serverURL, scheme, "/api/v4/posts", channels, req.Title, req.Message)
}

// sendRocketChat sends message via RocketChat
func (d *Dispatcher) sendRocketChat(req *types.WebhookRequest) error {
	serverURL := os.Getenv("ROCKETCHAT_SERVER_URL")
	userID := os.Getenv("ROCKETCHAT_USERID")
	token := os.Getenv("ROCKETCHAT_TOKEN")
	channels := os.Getenv("ROCKETCHAT_CHANNELS")
	scheme := os.Getenv("ROCKETCHAT_URL_SCHEME")
	if scheme == "" {
		scheme = "https"
	}

	if serverURL == "" || userID == "" || token == "" || channels == "" {
		return fmt.Errorf("ROCKETCHAT_SERVER_URL, ROCKETCHAT_USERID, ROCKETCHAT_TOKEN, and ROCKETCHAT_CHANNELS environment variables required")
	}

	return rocketchat.SendMessage(serverURL, scheme, userID, token, channels, req.Title, req.Message)
}

// sendPushbullet sends message via Pushbullet
func (d *Dispatcher) sendPushbullet(req *types.WebhookRequest) error {
	token := os.Getenv("PUSHBULLET_TOKEN")
	device := os.Getenv("PUSHBULLET_DEVICE")

	if token == "" || device == "" {
		return fmt.Errorf("PUSHBULLET_TOKEN and PUSHBULLET_DEVICE environment variables required")
	}

	return pushbullet.SendMessage(token, device, req.Title, req.Message)
}

// sendTwillio sends SMS via Twilio
func (d *Dispatcher) sendTwillio(req *types.WebhookRequest) error {
	accountSID := os.Getenv("TWILLIO_ACCOUNT_SID")
	token := os.Getenv("TWILLIO_TOKEN")
	sender := os.Getenv("TWILLIO_SENDER")
	receiver := os.Getenv("TWILLIO_RECEIVER")

	if accountSID == "" || token == "" || sender == "" || receiver == "" {
		return fmt.Errorf("TWILLIO_ACCOUNT_SID, TWILLIO_TOKEN, TWILLIO_SENDER, and TWILLIO_RECEIVER environment variables required")
	}

	return twillio.SendMessage(accountSID, token, sender, receiver, req.Title, req.Message)
}

// sendZulip sends message via Zulip
func (d *Dispatcher) sendZulip(req *types.WebhookRequest) error {
	domain := os.Getenv("ZULIP_DOMAIN")
	botEmail := os.Getenv("ZULIP_BOT_EMAIL_ADDRESS")
	apiKey := os.Getenv("ZULIP_BOT_API_KEY")
	msgType := os.Getenv("ZULIP_MSG_TYPE")
	stream := os.Getenv("ZULIP_STREAM_NAME")

	if domain == "" || botEmail == "" || apiKey == "" || stream == "" {
		return fmt.Errorf("ZULIP_DOMAIN, ZULIP_BOT_EMAIL_ADDRESS, ZULIP_BOT_API_KEY, and ZULIP_STREAM_NAME environment variables required")
	}

	if msgType == "" {
		msgType = "stream"
	}

	return zulip.SendMessage(domain, botEmail, apiKey, msgType, stream, req.Title, req.Message)
}

// sendMastodon sends message via Mastodon
func (d *Dispatcher) sendMastodon(req *types.WebhookRequest) error {
	token := os.Getenv("MASTODON_TOKEN")
	serverURL := os.Getenv("MASTODON_SERVER")

	if token == "" || serverURL == "" {
		return fmt.Errorf("MASTODON_TOKEN and MASTODON_SERVER environment variables required")
	}

	return mastodon.SendMessage(token, serverURL, req.Title, req.Message)
}

// sendLine sends message via Line
func (d *Dispatcher) sendLine(req *types.WebhookRequest) error {
	secret := os.Getenv("LINE_SECRET")
	token := os.Getenv("LINE_TOKEN")
	receivers := os.Getenv("LINE_RECEIVER_IDS")

	if secret == "" || token == "" || receivers == "" {
		return fmt.Errorf("LINE_SECRET, LINE_TOKEN, and LINE_RECEIVER_IDS environment variables required")
	}

	return line.SendMessage(secret, token, receivers, req.Title, req.Message)
}

// sendWeChat sends message via WeChat
func (d *Dispatcher) sendWeChat(req *types.WebhookRequest) error {
	appID := os.Getenv("WECHAT_APPID")
	appSecret := os.Getenv("WECHAT_APPSECRET")
	token := os.Getenv("WECHAT_TOKEN")
	aesKey := os.Getenv("WECHAT_ENCODING_AES_KEY")
	receivers := os.Getenv("WECHAT_RECEIVERS")

	if appID == "" || appSecret == "" || token == "" || aesKey == "" || receivers == "" {
		return fmt.Errorf("WECHAT_APPID, WECHAT_APPSECRET, WECHAT_TOKEN, WECHAT_ENCODING_AES_KEY, and WECHAT_RECEIVERS environment variables required")
	}

	return wechat.SendMessage(appID, appSecret, token, aesKey, receivers, req.Title, req.Message)
}

// sendGotify sends message via Gotify
func (d *Dispatcher) sendGotify(req *types.WebhookRequest) error {
	url := os.Getenv("GOTIFY_URL")
	token := os.Getenv("GOTIFY_TOKEN")

	if url == "" || token == "" {
		return fmt.Errorf("GOTIFY_URL and GOTIFY_TOKEN environment variables required")
	}

	priority := 5
	if req.Priority != 0 {
		priority = req.Priority
	} else if priorityStr := os.Getenv("GOTIFY_PRIORITY"); priorityStr != "" {
		if p, err := strconv.Atoi(priorityStr); err == nil {
			priority = p
		}
	}

	return gotify.SendMessage(url, token, req.Title, req.Message, priority)
}

// sendMatrix sends message via Matrix
func (d *Dispatcher) sendMatrix(req *types.WebhookRequest) error {
	serverURL := os.Getenv("MATRIX_SERVER_URL")
	accessToken := os.Getenv("MATRIX_ACCESS_TOKEN")
	roomID := os.Getenv("MATRIX_ROOM_ID")
	domain := os.Getenv("MATRIX_DOMAIN")

	if serverURL == "" || accessToken == "" {
		return fmt.Errorf("MATRIX_SERVER_URL and MATRIX_ACCESS_TOKEN environment variables required")
	}

	room := os.Getenv("MATRIX_ROOM")
	if room == "" && (roomID == "" || domain == "") {
		return fmt.Errorf("MATRIX_ROOM or (MATRIX_ROOM_ID and MATRIX_DOMAIN) environment variables required")
	}

	return matrix.SendMessage(serverURL, "", "", accessToken, room, roomID, domain, "", req.Message, false)
}
