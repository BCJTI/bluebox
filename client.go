package bluebox

import (
	"fmt"
)

type Client struct {
	AccountID string
	username  string
	password  string
}

func NewClient(accountID, username, password string) *Client {

	return &Client{
		AccountID: accountID,
		username:  username,
		password:  password,
	}

}

func (c *Client) HealthCheck() error {
	return c.Get(fmt.Sprintf("%shealthcheck", baseUrl), nil, nil, nil)
}
