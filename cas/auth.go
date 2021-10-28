package cas

import (
	_ "embed"
	"errors"
	"github.com/dop251/goja"
	request "github.com/parnurzeal/gorequest"
	"net/url"
	"regexp"
	"strconv"
)

var ltRegexp = regexp.MustCompile("<input type=\"hidden\" id=\"lt\" name=\"lt\" value=\"(.*)\" />")
var executionRegexp = regexp.MustCompile("<input type=\"hidden\" name=\"execution\" value=\"(.*)\" />")
var tokenRegexp = regexp.MustCompile("https://healthcheckin.hduhelp.com/\\?auth=([0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12})")

const jumpURL = "https://api.hduhelp.com/login/direct/cas?clientID=healthcheckin&redirect=https://healthcheckin.hduhelp.com/#/auth"

//go:embed des.js
var rawJs []byte

func GetToken(user, passwd string) (string, error) {
	var lt, execution string
	//搞到lt和execution
	req := request.New()
	req.AppendHeader("User-Agent", "Mozilla/5.0 (Linux; U; Android 4.1.2; zh-cn; GT-I9300 Build/JZO54K) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30 MicroMessenger/5.2.380 Edg/94.0.4606.71")
	res, body, errs := req.Get(jumpURL).End()
	if errs != nil && res.StatusCode == 200 {
		return "", errs[0]
	}
	tmp := ltRegexp.FindStringSubmatch(body)
	if len(tmp) != 2 {
		return "", errors.New("提取lt错误")
	}
	lt = tmp[1]
	tmp = executionRegexp.FindStringSubmatch(body)
	if len(tmp) != 2 {
		return "", errors.New("提取execution错误")
	}
	execution = tmp[1]
	//获取rsa
	rsa, err := getRsa(user + passwd + lt)
	if err != nil {
		return "", err
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
		return "", errs[0]
	}
	tmp = tokenRegexp.FindStringSubmatch(res.Request.URL.String())
	return tmp[1], nil
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
