package goss

import (
	"context"
	"github.com/dghubble/sling"
	"golang.org/x/oauth2"
)

type Client struct {
	Instances *Instances
	Plans     *Plans
}

func NewClientFromToken(apiKey string) *Client {
	return NewClient("", apiKey)
}

func NewClient(baseUrl, apiKey string) *Client {
	c := sling.New().
		Client(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey}))).
		Base(baseUrl)
	return &Client{
		Instances: &Instances{sling: c},
		Plans:     &Plans{sling: c},
	}
}
