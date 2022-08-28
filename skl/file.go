package skl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type OSSFile struct {
	// "hdu-checkin"
	Bucket string `json:"bucket"`
	// 文件名(不含路径)
	FileName string `json:"fileName"`
	// 疑似为文件路径，如"leave-file/2022-08-28/[随机大小写数字混合19位字符串].png"
	Key string `json:"key"`
}

type signeResp struct {
	Key         string `json:"key"`
	ContentType string `json:"contentType"`
	Bucket      string `json:"bucket"`
	Host        string `json:"host"`
	Url         string `json:"url"`
}

const signeURL = "https://skl.hdu.edu.cn/api/oss/generateSigne?fileName=oss."

func (user *User) Upload(file string) (*OSSFile, error) {
	// check file
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	// get signe url
	req, err := http.NewRequest(http.MethodGet, signeURL+filepath.Ext(file), nil)
	if err != nil {
		return nil, err
	}
	user.addHeaderToReq(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("statuscode is not ok: %s", string(body))
	}
	res := new(signeResp)
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	// upload
	req, err = http.NewRequest(http.MethodPut, res.Url, f)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", res.ContentType)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("statuscode is not ok: %s", string(body))
	}
	return &OSSFile{
		Bucket:   res.Bucket,
		FileName: filepath.Base(file),
		Key:      res.Key,
	}, nil
}
