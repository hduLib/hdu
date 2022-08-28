package net

import (
	"net/http"
)

type Client interface {
	Do(r *http.Request) (*http.Response, error)
}

// DefaultClient do all requests, you can change it as you implement the interface.
var DefaultClient Client = &http.Client{}
