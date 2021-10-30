package cas

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/dop251/goja"
	request "github.com/parnurzeal/gorequest"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

var ltRegexp = regexp.MustCompile("<input type=\"hidden\" id=\"lt\" name=\"lt\" value=\"(.*)\" />")
var executionRegexp = regexp.MustCompile("<input type=\"hidden\" name=\"execution\" value=\"(.*)\" />")

//go:embed des.js
var rawJs []byte

func Login(URL, user, passwd string) (*http.Response, error) {
	var lt, execution string
	//搞到lt和execution
	req := request.New()
	req.AppendHeader("User-Agent", "Mozilla/5.0 (Linux; U; Android 4.1.2; zh-cn; GT-I9300 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 MicroMessenger/5.2.380 Edg/94.0.4606.71")
	res, body, errs := req.Get(URL).End()
	if errs != nil || res.StatusCode != 200 {
		return nil, errs[0]
	}
	tmp := ltRegexp.FindStringSubmatch(body)
	if len(tmp) != 2 {
		return nil, errors.New("提取lt错误")
	}
	lt = tmp[1]
	tmp = executionRegexp.FindStringSubmatch(body)
	if len(tmp) != 2 {
		return nil, errors.New("提取execution错误")
	}
	execution = tmp[1]
	//获取rsa
	rsa, err := getRsa(user + passwd + lt)
	if err != nil {
		return nil, err
	}

	postData := url.Values{}
	postData.Add("rsa", rsa)
	postData.Add("ul", strconv.Itoa(len(user)))
	postData.Add("pl", strconv.Itoa(len(passwd)))
	postData.Add("lt", lt)
	postData.Add("execution", execution)
	postData.Add("_eventId", "submit")

	res, _, errs = req.Post(res.Request.URL.String()).Type("form").Send(map[string]string{
		"rsa":       rsa,
		"ul":        strconv.Itoa(len(user)),
		"pl":        strconv.Itoa(len(passwd)),
		"lt":        lt,
		"execution": execution,
		"_eventId":  "submit"}).End()

	if errs != nil {
		return nil, errs[0]
	}

	if res.StatusCode != 200 {
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("http code isn't 200:%s", string(content))
	}
	return res, nil
}

func getRsa(data string) (string, error) {
	vm := goja.New()
	_, err := vm.RunString(string(rawJs))
	if err != nil {
		panic(err)
	}
	strEnc, valid := goja.AssertFunction(vm.Get("strEnc"))
	if !valid {
		return "", errors.New("invalid js")
	}
	value, err := strEnc(nil, vm.ToValue(data), vm.ToValue("1"), vm.ToValue("2"), vm.ToValue("3"))
	if err != nil {
		panic(err)
	}
	var result = value.String()
	return result, nil
}
