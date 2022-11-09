package client

import (
	"errors"
	"net/http"
)

type Client interface {
	Do(r *http.Request) (*http.Response, error)
}

// DefaultClient do all requests, you can change it as you implement the interface.
var DefaultClient Client = CommonClient

// CommonClient is set as DefaultClient default
// you maybe puzzled about CheckRedirect is used instead of cookieJar.
// it's because cookieJar is global, a cookie from one request may affect in another way
// that never be considered.
var CommonClient = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		for _, v := range append(via[len(via)-1].Cookies(), req.Response.Cookies()...) {
			ck, err := req.Cookie(v.Name)
			if err != nil {
				if errors.Is(err, http.ErrNoCookie) {
					req.AddCookie(v)
				} else {
					return err
				}
				continue
			}
			ck.Value = v.Value
		}
		return nil
	},
}
