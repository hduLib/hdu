package ocr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"
)

const customUrl = "https://www.jfbym.com/api/YmServer/customApi"

var token string

// SetToken sets the token for later usage.
func SetToken(t string) {
	token = t
}

type (
	commonConfig = map[string]interface{}
)

var (
	commonConfigPool = &sync.Pool{
		New: func() any {
			return make(map[string]interface{}, 3)
		},
	}

	// better use obj-pool to release the pressure of GOGC
)

// Generated by https://quicktype.io

type CommonTypeResponse struct {
	Msg  string `json:"msg"`
	Code int64  `json:"code"`
	Data Data   `json:"data"`
}

type Data struct {
	Code       int64   `json:"code"`
	Data       string  `json:"data"`
	Time       float64 `json:"time"`
	UniqueCode string  `json:"unique_code"`
}

var (
	commonTypeRespPool = &sync.Pool{
		New: func() any {
			return new(CommonTypeResponse)
		},
	}
)

// commonVerify needs a base64 encoded image to send a request
// to yunma ocr to get the result. To find more about the common
// type of verification, see: <https://www.jfbym.com/demo/>.
func commonVerify(image string) (string, error) {
	if token == "" {
		return "", errors.New("token unset")
	}

	// construct common type config
	cfg := commonConfigPool.Get().(map[string]interface{})
	defer commonConfigPool.Put(cfg)
	cfg["image"] = image
	cfg["type"] = "10110"
	cfg["token"] = token

	// send request to yunma ocr
	rawcfg, _ := json.Marshal(cfg)
	body := bytes.NewReader(rawcfg)
	resp, err := http.Post(customUrl, "application/json;charset=utf-8", body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// recv from yunma ocr
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	commonTypeResp := commonTypeRespPool.Get().(*CommonTypeResponse)
	defer commonTypeRespPool.Put(commonTypeResp)
	err = json.Unmarshal(data, commonTypeResp)
	if err != nil {
		return "", err
	}
	return commonTypeResp.Data.Data, nil
}
