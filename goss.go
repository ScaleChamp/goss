package goss

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client

	BaseURL *url.URL

	Instances InstancesServiceOp
	Plans     PlansServiceOp
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "goss")
	return req, nil
}

func (c *Client) Do(ctx context.Context, request *http.Request, v interface{}) error {
	response, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		return err
	}
	defer func() {
		if rerr := response.Body.Close(); rerr != nil {
			err = rerr
		}
	}()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("err: %s", response.Status)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func NewClientFromToken(apiKey string) *Client {
	return NewClient("https://api.scalechamp.com", apiKey)
}

func NewClient(baseUrl, apiKey string) *Client {
	sourceToken := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: apiKey})
	client := oauth2.NewClient(context.Background(), sourceToken)
	u, _ := url.Parse(baseUrl)
	c := &Client{client: client, BaseURL: u}
	c.Instances = &Instances{client: c}
	c.Plans = &Plans{client: c}
	return c
}
