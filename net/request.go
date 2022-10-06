package net

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get(req *http.Request, data interface{}) error {
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code is %d:%s", resp.StatusCode, string(body))
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
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code is %d:%s", resp.StatusCode, string(resBody))
	}
	return resBody, err
}
