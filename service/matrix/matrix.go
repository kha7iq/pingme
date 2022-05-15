package matrix

import (
	"fmt"
	"github.com/matrix-org/gomatrix"
	"github.com/urfave/cli/v2"
	"strings"
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
			// Login
			client, err := matrix.login()
			if err != nil {
				return fmt.Errorf("failed to login to matrix: %v", err)
			}

			// Parse and set variables
			err = matrix.setupVars()
			if err != nil {
				return err
			}

			// If necessary, join the given room
			err = matrix.joinRoomIfNecessary(client)
			if err != nil {
				return err
			}

			// Send the message
			_, err = client.SendText(matrix.Room, matrix.Message)
			if err != nil {
				return fmt.Errorf("failed to send matrix text: %v", err)
			}
			return nil
		},
	}
}

/*
setupVars will ensure the room ID begins with an exclamation mark and set the room string if not
already set, using the room ID and domain. If the room string, room id and domain are not set,
an error will be thrown.
*/
func (m *matrixPingMe) setupVars() error {
	// Format the room ID
	if !strings.HasPrefix(m.RoomID, "!") {
		m.RoomID = "!" + m.RoomID
	}

	// Create the matrix room string if not already provided
	if m.Room == "" {
		if m.RoomID == "" || m.Domain == "" {
			return fmt.Errorf("matrix room, or room ID and domain must be provided")
		}

		m.Room = fmt.Sprintf("%s:%s", m.RoomID, m.Domain)
	}
	return nil
}

/*
joinRoomIfNecessary gets all the joined rooms and checks if the desired room is in the list.
If not, and autoJoin is set to true - will attempt to join the room. If autoJoin is set to
false, an error will be thrown
*/
func (m *matrixPingMe) joinRoomIfNecessary(client *gomatrix.Client) error {
	// Get already joined rooms
	joined, err := client.JoinedRooms()
	if err != nil {
		return fmt.Errorf("failed to get joined rooms: %v", err)
	}

	// Check if we've already joined the desired room
	foundRoom := false
	for _, room := range joined.JoinedRooms {
		if room == m.Room {
			foundRoom = true
			break
		}
	}

	// If not, try auto join the room
	if !foundRoom {
		if !m.AutoJoin {
			return fmt.Errorf("not joined room '%s' and --autoJoin is set to false", m.Room)
		}

		_, err = client.JoinRoom(m.Room, m.ServerName, nil)
		if err != nil {
			return fmt.Errorf("failed to auto join room '%s': %v", m.Room, err)
		}
	}
	return nil
}

/*
login creates a gomatrix.Client instance, connecting to the given URL, using the provided login details
*/
func (m *matrixPingMe) login() (*gomatrix.Client, error) {
	// Create a client instance
	client, err := gomatrix.NewClient(m.Url, "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to create matrix client: %v", err)
	}

	// Attempt to log in with whatever login details were provided.
	// Or, throw an error if no login details were given
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

	// Set the access token for this session
	client.SetCredentials(resp.UserID, resp.AccessToken)
	m.Token = resp.AccessToken
	return client, nil
}
