package zjooc

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hduLib/hdu/internal/client"
	"github.com/tidwall/gjson"
	"net/http"
)

func Login(account, password string) (*User, error) {
	payload := LoginReq{
		LoginName: account,
		Password:  password,
		Type:      1,
	}
	b, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, loginUrl, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	body, err := client.Post(req)
	str := string(body)
	success := gjson.Get(str, "success").Bool()
	if !success {
		return nil, errors.New(gjson.Get(str, "message").String())
	}
	user := new(User)
	user.openid = gjson.Get(str, "data").String()
	return user, nil
}
