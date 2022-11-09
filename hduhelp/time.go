package hduhelp

import (
	"github.com/hduLib/hdu/client"
	"net/http"
)

func Time() (*TimeResp, error) {
	req, err := http.NewRequest("GET", timeURL, nil)
	if err != nil {
		return nil, err
	}
	time := new(TimeResp)
	err = client.Get(req, time)
	if err != nil {
		return nil, err
	}
	return time, nil
}
