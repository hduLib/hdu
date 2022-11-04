package net

import (
	"encoding/json"
	"io"
	"net/http"
)

func Get(req *http.Request, data interface{}) error {
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return &ErrNotOk{resp.StatusCode, string(body)}
	}
	if err := json.Unmarshal(body, data); err != nil {
		return err
	}
	return nil
}

func Post(req *http.Request) ([]byte, error) {
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, &ErrNotOk{resp.StatusCode, string(resBody)}
	}
	return resBody, err
}
