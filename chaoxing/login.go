package chaoxing

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hduLib/hdu/cas"
	"github.com/hduLib/hdu/chaoxing/utils"
	"github.com/hduLib/hdu/client"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type loginResp struct {
	Msg    string `json:"msg2"`
	Status bool   `json:"status"`
	Url    string `json:"url"`
}

func LoginWithPhoneAndPwd(phone string, passwd string) (*Cx, error) {
	payload := url.Values{}
	payload.Set("fid", "1001")
	payload.Set("uname", utils.EncryptByAES(phone))
	payload.Set("password", utils.EncryptByAES(passwd))
	payload.Set("refer", "http://i.mooc.chaoxing.com")
	payload.Set("t", "true")
	payload.Set("doubleFactorLogin", "0")
	payload.Set("independentId", "0")
	payload.Set("validate", "")

	req, err := http.NewRequest(http.MethodPost, fanyaLoginURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Referer", "http://passport2.chaoxing.com/login?loginType=4&newversion=true&fid=1001&refer=http://i.mooc.chaoxing.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 Edg/106.0.1370.52")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request fail:%v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request fail:http code is %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	lres := &loginResp{}
	if err := json.Unmarshal(body, lres); err != nil {
		return nil, err
	}
	if !lres.Status {
		return nil, fmt.Errorf("login fail:%s", lres.Msg)
	}
	return newUser(resp.Cookies()), nil
}

func LoginWithCas(user, passwd string) (*Cx, error) {
	req, err := cas.GenLoginReq(ssoLoginURL, user, passwd)
	if err != nil {
		return nil, fmt.Errorf("fail to gen login request:%v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, &client.ErrNotOk{StatusCode: resp.StatusCode, Body: string(body)}
	}
	if resp.Request.URL.String() != ssoSuccessURL {
		return nil, errors.New("invalid id or password")
	}
	return newUser(resp.Request.Cookies()), nil
}
