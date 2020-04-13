package client

type DB interface {
	Save(c NewOauth2Client) (*Oauth2Client, error)
}

type Oauth2ClientUsecaseConfig struct {
	db DB
}

type NewOauth2Client struct{}

type Oauth2Client struct{}

func (c *Oauth2ClientUsecaseConfig) NewClient(newClient NewOauth2Client) (*Oauth2Client, error) {

	client, err := c.db.Save(newClient)
	if err != nil {
		return nil, err
	}
	return client, nil
}
