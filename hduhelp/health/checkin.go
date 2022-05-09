package health

import (
	"encoding/json"
	"errors"
	"fmt"
	request "github.com/parnurzeal/gorequest"
	"log"
	"time"
)

const (
	validateURL = "https://api.hduhelp.com/token/validate"
	checkInURL  = "https://api.hduhelp.com/salmon_base/health/checkin?sign="
	infoURL     = "https://api.hduhelp.com/salmon_base/person/info"
	dailyURL    = "https://api.hduhelp.com/salmon_base/health/checkin/today"
	phoneURl    = "https://api.hduhelp.com/salmon_base/health/phone"
	codeURL     = "https://api.hduhelp.com/salmon_base/health/code"
)

var ErrInvalid = errors.New("invalid token")

type ansPayload struct {
	Ques1  string      `json:"ques1"`
	Ques2  string      `json:"ques2"`
	Ques3  interface{} `json:"ques3"`
	Ques4  string      `json:"ques4"`
	Ques5  string      `json:"ques5"`
	Ques6  string      `json:"ques6"`
	Ques7  interface{} `json:"ques7"`
	Ques77 interface{} `json:"ques77"`
	Ques8  interface{} `json:"ques8"`
	Ques88 interface{} `json:"ques88"`
	Ques9  interface{} `json:"ques9"`
	Ques10 interface{} `json:"ques10"`
	Ques11 interface{} `json:"ques11"`
	Ques12 interface{} `json:"ques12"`
	Ques13 interface{} `json:"ques13"`
	Ques14 interface{} `json:"ques14"`
	Ques15 string      `json:"ques15"`
	Ques16 string      `json:"ques16"`
	Ques17 string      `json:"ques17"`
	Ques18 string      `json:"ques18"`
	Ques19 interface{} `json:"ques19"`
	Ques20 string      `json:"ques20"`
	Ques21 string      `json:"ques21"`
	Ques22 string      `json:"ques22"`
	Ques23 string      `json:"ques23"`
	Ques24 string      `json:"ques24"`
	CarTo  []string    `json:"carTo"`
}

type checkInPayload struct {
	Name      string `json:"name,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Province  string `json:"province,omitempty"`
	City      string `json:"city,omitempty"`
	//这里的country是因为杭电助手那边拼错了，我没办法
	County        string `json:"country,omitempty"`
	AnswerJsonStr string `json:"answerJsonStr,omitempty"`
}

type info struct {
	StaffId    string `json:"staffId"`
	StaffName  string `json:"staffName"`
	StaffState string `json:"staffState"`
	StaffType  string `json:"staffType"`
	UnitCode   string `json:"unitCode"`
}

type infoRes struct {
	Cache bool   `json:"cache"`
	Data  info   `json:"data"`
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

type validRes struct {
	Data validate `json:"Data"`
}

type validate struct {
	AccessToken string `json:"accessToken"`
	ClientID    string `json:"clientID"`
	ExpiredTime int64  `json:"expiredTime"`
	GrantType   string `json:"grantType"`
	IsValid     int    `json:"isValid"`
	School      string `json:"school"`
	StaffId     string `json:"staffId"`
	TokenType   int    `json:"tokenType"`
	Uid         int    `json:"uid"`
}

type Health struct {
	token   string
	profile struct {
		name, id, province, city, county string
		atHome                           bool
	}
	req *request.SuperAgent
}

func New() *Health {
	req := request.New()
	req.Set("User-Agent", "Mozilla/5.0 (Linux; Android 10; MAR-AL00 Build/HUAWEIMAR-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/86.0.4240.99 XWEB/3135 MMWEBSDK/20210902 Mobile Safari/537.36 MMWEBID/871 MicroMessenger/8.0.14.2000(0x28000E37) Process/toolsmp WeChat/arm64 Weixin NetType/4G Language/zh_CN ABI/arm64")
	req.Set("Referer", "https://healthcheckin.hduhelp.com/")
	req.Set("Origin", "https://healthcheckin.hduhelp.com")
	req.Set("X-Requested-With", "com.tencent.mm")
	req.Set("Sec-Fetch-Site", "same-site")
	req.Set("Sec-Fetch-Mode", "cors")
	req.Set("Sec-Fetch-Dest", "empty")
	req.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	//	req.Proxy("http://127.0.0.1:8080")
	return &Health{
		profile: struct {
			name, id, province, city, county string
			atHome                           bool
		}{
			province: "330000",
			city:     "330100",
			county:   "330101",
			atHome:   false,
		},
		req: req,
	}
}

func (h *Health) Checkin() error {

	now := time.Now().Unix()

	//https://github.com/dreamer2q/hdu-health-checkin/blob/master/checkin/client.go#L87
	sign := h.getSign(now)

	req := h.getReq()

	req.Options(checkInURL + sign).End()

	req.Post(checkInURL + sign).Send(checkInPayload{
		Name:          h.profile.name,
		Timestamp:     now,
		Province:      h.profile.province,
		City:          h.profile.city,
		County:        h.profile.county,
		AnswerJsonStr: newAnsPayload(h.profile.province, h.profile.city, h.profile.county, h.profile.atHome),
	})

	req.Set("Host", "api.hduhelp.com")
	req.Set("Accept", "application/json, text/plain, */*")
	req.Set("Content-Type", "application/json;charset=UTF-8")

	res, body, errs := req.End()
	if errs != nil {
		return errs[0]
	}

	if res.StatusCode != 200 {
		return errors.New(body)
	}
	return nil
}

func (h *Health) getSign(timestamp int64) string {
	sign := fmt.Sprintf("%s%s%d%s%s%s", h.profile.name, b64(h.profile.id), timestamp, b64(h.profile.province), h.profile.city, h.profile.county)
	return s1(sign)
}

func (h *Health) SetToken(token string) error {
	h.token = token
	h.req.Set("Authorization", "token "+h.token)
	_, err := h.Validate()
	if err != nil {
		return err
	}
	info, err := h.Info()
	if err != nil {
		return err
	}
	h.profile.id = info.StaffId
	h.profile.name = info.StaffName
	return nil
}

// SetLocation is used to change location,and the default location is Hangzhou.
func (h *Health) SetLocation(code string) error {
	if len(code) != 6 {
		return errors.New("invalid code")
	}
	h.profile.province = code[0:2] + "0000"
	h.profile.city = code[0:4] + "00"
	h.profile.county = code
	if code != "330101" {
		h.profile.atHome = true
	}
	return nil
}

func (h *Health) AtHome() {
	h.profile.atHome = true
}

func (h *Health) AtSchool() {
	h.profile.atHome = false
	h.SetLocation("330101")
}

func (h *Health) Validate() (*validate, error) {
	res, body, err := h.getReq().Get(validateURL).EndBytes()
	if err != nil {
		return nil, err[0]
	}
	if res.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	var resForm = new(validRes)
	if err := json.Unmarshal(body, resForm); err != nil {
		return nil, err
	}
	if resForm.Data.IsValid != 1 {
		return nil, ErrInvalid
	}
	return &resForm.Data, nil
}

func (h *Health) getReq() *request.SuperAgent {
	return h.req.Clone()
}

func (h *Health) Info() (*info, error) {
	res, body, err := h.getReq().Get(infoURL).EndBytes()
	if err != nil && res.StatusCode != 200 {
		return nil, err[0]
	}
	var infoForm infoRes
	if err := json.Unmarshal(body, &infoForm); err != nil {
		return nil, err
	}
	if infoForm.Msg != "success" || infoForm.Error != 0 {
		return nil, errors.New(infoForm.Msg)
	}
	return &infoForm.Data, nil
}

func (h *Health) Daily() {
	//todo:daily
	h.getReq().Get(dailyURL).End()
}

func (h *Health) Phone() {
	//todo:phone
	h.getReq().Get(phoneURl).End()
}

func (h *Health) Code() {
	//todo:code
	h.getReq().Get(codeURL).End()
}

func newAnsPayload(province, city, county string, atHome bool) string {
	ans := ansPayload{
		Ques1:  "健康良好",
		Ques4:  "否",
		Ques5:  "否",
		Ques6:  "",
		Ques15: "否",
		Ques16: "否",
		Ques17: "无新冠肺炎确诊或疑似",
		Ques18: "37度以下",
		Ques20: "绿码",
		Ques21: "否",
		Ques22: "否",
		Ques23: "否",
		Ques24: "共三针 - 已完成第三针",
		CarTo:  []string{province, city, county},
	}
	if atHome {
		ans.Ques2 = "正常在家"
	} else {
		ans.Ques2 = "正常在校（未经学校审批，不得提前返校）"
	}
	buf, err := json.Marshal(ans)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(buf)
}
