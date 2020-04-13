package oauth2client

import (
	"crypto/rand"
	"fmt"
	password "github.com/ernstvorsteveld/go-password"
	"github.com/google/uuid"
)

type PasswordSettings struct {
	pwdLength int
}

type ClientUsecaseConfig struct {
	handler    NewClientHandler
	pwdSetting PasswordSettings
}

type NewClientHandler interface {
	HandlePassword(newClient *NewClient)
	HandleClientId(newClient *NewClient)
}

func (c *ClientUsecaseConfig) HandlePassword(n *NewClient) {
	l := c.pwdSetting.pwdLength
	b := make([]byte, l)
	pwd, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	bp := password.BcryptPassword{Cost: 10}
	bp.Hash(string(pwd))
	if err != nil {

	}
	n.secret = &bp
}

func (c *ClientUsecaseConfig) HandleClientId(n *NewClient) {
	if n.client_id == "" {
		n.client_id = uuid.UUID.New()
	}
}

type Usecase interface {
	NewClient(newClient NewClient) (*Client, error)
}

func (c *ClientUsecaseConfig) NewClient(newClient NewClient) (*Client, error) {
	c.handler.HandleClientId(&newClient)
	return client, nil
}
