package net

import (
	"net/http"
)

type Client interface {
	Do(r *http.Request) (*http.Response, error)
}

// DefaultClient do all requests, you can change it as you implement the interface.
var DefaultClient Client = &http.Client{
	// package chaoxing needs it for adding cookie to following request
	// during redirecting. If you write your client, plz pay attention to it.
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		for _, v := range via[0].Cookies() {
			req.AddCookie(v)
		}
		return nil
	},
}
