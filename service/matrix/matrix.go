package matrix

import (
	"fmt"
	"log"
	"strings"

	"github.com/matrix-org/gomatrix"
	"github.com/urfave/cli/v2"
)

type matrixPingMe struct {
	Username   string
	Password   string
	Token      string
	Url        string
	ServerName string
	Room       string
	RoomID     string
	Domain     string
	Message    string
	AutoJoin   bool
}

// SendMessage sends a message to a matrix room.
func SendMessage(serverURL, username, password, token, room, roomID, domain, serverName, message string, autoJoin bool) error {
	if serverURL == "" {
		return fmt.Errorf("matrix server URL is required")
	}
	if message == "" {
		return fmt.Errorf("message is required")
	}

	m := &matrixPingMe{
		Username:   username,
		Password:   password,
		Token:      token,
		Url:        serverURL,
		ServerName: serverName,
		Room:       room,
		RoomID:     roomID,
		Domain:     domain,
		Message:    message,
		AutoJoin:   autoJoin,
	}

	// Login
	client, err := m.login()
	if err != nil {
		return fmt.Errorf("failed to login to matrix: %w", err)
	}

	// Parse and set variables
	err = m.setupVars()
	if err != nil {
		return err
	}

	// If necessary, join the given room
	err = m.joinRoomIfNecessary(client)
	if err != nil {
		return err
	}

	// Send the message
	_, err = client.SendText(m.Room, m.Message)
	if err != nil {
		return fmt.Errorf("failed to send matrix text: %w", err)
	}

	log.Println("Successfully sent!")
	return nil
}

func Send() *cli.Command {
	var matrix matrixPingMe
	return &cli.Command{
		Name:  "matrix",
		Usage: "Send message via matrix",
		UsageText: "pingme matrix --token 'syt_YW...E2qD' --room 'LRovrjPJaRChcTKgoK:matrix.org' " +
			"--url 'matrix-client.matrix.org' --autoJoin --msg 'Hello, Matrix!'",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &matrix.Username,
				Name:        "username",
				Aliases:     []string{"u"},
				Usage:       "Matrix username",
				EnvVars:     []string{"MATRIX_USER"},
			},
			&cli.StringFlag{
				Destination: &matrix.Password,
				Name:        "password",
				Aliases:     []string{"p"},
				Usage:       "Matrix password",
				EnvVars:     []string{"MATRIX_PASSWORD"},
			},
			&cli.StringFlag{
				Destination: &matrix.Token,
				Name:        "token",
				Aliases:     []string{"t"},
				Usage:       "Matrix access token. Can be used instead of username+password",
				EnvVars:     []string{"MATRIX_ACCESS_TOKEN"},
			},
			&cli.StringFlag{
				Destination: &matrix.Url,
				Name:        "url",
				Usage:       "Matrix server URL",
				EnvVars:     []string{"MATRIX_SERVER_URL"},
			},
			&cli.StringFlag{
				Destination: &matrix.ServerName,
				Name:        "serverName",
				Usage:       "Can be provided if requests should be routed via a particular server",
				EnvVars:     []string{"MATRIX_SERVER_NAME"},
			},
			&cli.StringFlag{
				Destination: &matrix.Room,
				Name:        "room",
				Usage:       "Matrix room to send the message to, in the format <roomId>:<domain>",
				EnvVars:     []string{"MATRIX_ROOM"},
			},
			&cli.StringFlag{
				Destination: &matrix.RoomID,
				Name:        "roomId",
				Usage:       "Matrix room ID to send the message to. The exclamation mark at the beginning can be excluded.",
				EnvVars:     []string{"MATRIX_ROOM_ID"},
			},
			&cli.StringFlag{
				Destination: &matrix.Domain,
				Name:        "domain",
				Usage:       "Used in conjunction with room ID to get the desired room",
				EnvVars:     []string{"MATRIX_DOMAIN"},
			},
			&cli.StringFlag{
				Destination: &matrix.Message,
				Name:        "msg",
				Aliases:     []string{"m"},
				Required:    true,
				Usage:       "Message to send to matrix",
				EnvVars:     []string{"MATRIX_MESSAGE"},
			},
			&cli.BoolFlag{
				Destination: &matrix.AutoJoin,
				Name:        "autoJoin",
				Usage:       "If enabled, will automatically join the specified room if not already joined",
				EnvVars:     []string{"MATRIX_AUTO_JOIN"},
			},
		},
		Action: func(ctx *cli.Context) error {
			return SendMessage(
				matrix.Url,
				matrix.Username,
				matrix.Password,
				matrix.Token,
				matrix.Room,
				matrix.RoomID,
				matrix.Domain,
				matrix.ServerName,
				matrix.Message,
				matrix.AutoJoin,
			)
		},
	}
}

func (m *matrixPingMe) setupVars() error {
	if !strings.HasPrefix(m.RoomID, "!") {
		m.RoomID = "!" + m.RoomID
	}

	if m.Room == "" {
		if m.RoomID == "" || m.Domain == "" {
			return fmt.Errorf("matrix room, or room ID and domain must be provided")
		}
		m.Room = fmt.Sprintf("%s:%s", m.RoomID, m.Domain)
	}
	return nil
}

func (m *matrixPingMe) joinRoomIfNecessary(client *gomatrix.Client) error {
	joined, err := client.JoinedRooms()
	if err != nil {
		return fmt.Errorf("failed to get joined rooms: %w", err)
	}

	foundRoom := false
	for _, room := range joined.JoinedRooms {
		if room == m.Room {
			foundRoom = true
			break
		}
	}

	if !foundRoom {
		if !m.AutoJoin {
			return fmt.Errorf("not joined room '%s' and --autoJoin is set to false", m.Room)
		}

		_, err = client.JoinRoom(m.Room, m.ServerName, nil)
		if err != nil {
			return fmt.Errorf("failed to auto join room '%s': %w", m.Room, err)
		}
	}
	return nil
}

func (m *matrixPingMe) login() (*gomatrix.Client, error) {
	client, err := gomatrix.NewClient(m.Url, "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to create matrix client: %w", err)
	}

	var resp *gomatrix.RespLogin
	if m.Token != "" {
		resp, err = client.Login(&gomatrix.ReqLogin{
			Type:  "m.login.token",
			Token: m.Token,
		})
		if err != nil {
			return nil, err
		}
	} else if m.Username != "" && m.Password != "" {
		resp, err = client.Login(&gomatrix.ReqLogin{
			Type:     "m.login.password",
			User:     m.Username,
			Password: m.Password,
		})
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no token, or username and password provided")
	}

	client.SetCredentials(resp.UserID, resp.AccessToken)
	m.Token = resp.AccessToken
	return client, nil
}
