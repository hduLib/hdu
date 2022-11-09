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
	// UA:在浙学/2 CFNetwork/1390 Darwin/22.0.0
	req.Header.Set("User-Agent", "%E5%9C%A8%E6%B5%99%E5%AD%A6/2 CFNetwork/1390 Darwin/22.0.0")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
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
	user.openid = gjson.Get(str, "data.loginResult.openid").String()
	return user, nil
}
