package hduhelp

import (
	"net/http"

	"github.com/hduLib/hdu/client"
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
