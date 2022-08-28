package net

import (
	"net/http"
	"time"
)

// LowNetworkClient example
type LowNetworkClient struct {
	http.Client
	retry int
}

func NewLowNetworkClient(timeout time.Duration, retry int) *LowNetworkClient {
	return &LowNetworkClient{http.Client{Timeout: timeout}, retry}
}

// Do is rewritten to retry
func (lc *LowNetworkClient) Do(r *http.Request) (*http.Response, error) {
	var (
		err  error
		resp *http.Response
	)
	for i := 0; i < lc.retry; i++ {
		resp, err = lc.Client.Do(r)
		if err == nil {
			return resp, nil
		}
	}
	return nil, err
}
