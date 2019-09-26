package goss

import (
	"github.com/dghubble/sling"
	"net/http"
)

type API struct {
	sling *sling.Sling
}

func New(baseUrl, apiKey string) *API {
	return &API{
		sling: sling.New().
			Client(http.DefaultClient).
			Base(baseUrl).
			SetBasicAuth("", apiKey),
	}
}
