package health

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BaiMeow/hdu/cas"
	request "github.com/parnurzeal/gorequest"
	"time"
)

const (
	AnsStr = "{\"ques1\":\"健康良好\",\"ques2\":\"正常在校（未经学校审批，不得提前返校）\",\"ques3\":null,\"ques4\":\"否\",\"ques5\":\"否\",\"ques6\":\"\",\"ques7\":null,\"ques77\":null,\"ques8\":null,\"ques88\":null,\"ques9\":null,\"ques10\":null,\"ques11\":null,\"ques12\":null,\"ques13\":null,\"ques14\":null,\"ques15\":\"否\",\"ques16\":\"否\",\"ques17\":\"无新冠肺炎确诊或疑似\",\"ques18\":\"37度以下\",\"ques19\":null,\"ques20\":\"绿码\",\"ques21\":\"否\",\"ques22\":\"否\",\"ques23\":\"否\",\"ques24\":\"共二针 - 已完成第二针\",\"carTo\":[\"330000\",\"330100\",\"330101\"]}"

	checkInURL  = "https://api.hduhelp.com/base/healthcheckin?sign="
	validateURL = "https://api.hduhelp.com/token/validate"
	infoURL     = "https://api.hduhelp.com/salmon_base/person/info"
	dailyURL    = "https://api.hduhelp.com/base/healthcheckin/info/daily"
	phoneURl    = "https://api.hduhelp.com/base/healthcheckin/phone"
	codeURL     = "https://api.hduhelp.com/salmon_base/health/code"
)

var ErrInvalid = errors.New("invalid token")

type checkInPayload struct {
	Name          string `json:"name,omitempty"`
	Timestamp     int64  `json:"timestamp,omitempty"`
	Province      string `json:"province,omitempty"`
	City          string `json:"city,omitempty"`
	Country       string `json:"country,omitempty"`
	AnswerJsonStr string `json:"answerjsonstr,omitempty"`
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
		name, id, province, city, country string
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
			name, id, province, city, country string
		}{
			province: "330000",
			city:     "330100",
			country:  "330101",
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
		Country:       h.profile.country,
		AnswerJsonStr: AnsStr,
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
	sign := fmt.Sprintf("%s%s%d%s%s%s", h.profile.name, b64(h.profile.id), timestamp, b64(h.profile.province), h.profile.city, h.profile.country)
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
func (h *Health) SetLocation(province, city, country string) {
	h.profile.province = province
	h.profile.city = city
	h.profile.country = country
}

func (h *Health) GetTokenByCas(passwd string) error {
	token, err := cas.GetToken(h.profile.id, passwd)
	if err != nil {
		return err
	}
	if err := h.SetToken(token); err != nil {
		return err
	}
	return nil
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
