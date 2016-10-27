package main

import (
	"github.com/mrsaints/go-cabot/cabot"
)

type Config struct {
	BaseURL  string
	Username string
	Password string
}

func (c *Config) Client() (interface{}, error) {
	base_url := cabot.WithBaseURL(c.BaseURL)
	auth := cabot.WithBasicAuth(c.Username, c.Password)
	client, err := cabot.NewClient(base_url, auth)
	if err != nil {
		return nil, err
	}
	return client, nil
}
