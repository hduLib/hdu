package chaoxing

import (
	"encoding/json"
	"fmt"
	"github.com/hduLib/hdu/chaoxing/utils"
	"github.com/hduLib/hdu/net"
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

func LoginByPhoneAndPwd(phone string, passwd string) (*Cx, error) {
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

	resp, err := net.DefaultClient.Do(req)
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