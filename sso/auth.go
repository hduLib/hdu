package sso

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/hduLib/hdu/client"
)

// 正则表达式用于从 HTML 中提取 execution 和 croypto
var executionRegexp = regexp.MustCompile(`id="login-page-flowkey"[^>]*>([^<]+)`)
var croyptoRegexp = regexp.MustCompile(`id="login-croypto"[^>]*>([^<]+)`)

func GenLoginReq(URL, user, passwd string) (*http.Request, error) {
	// 1. GET 登录页，获取 execution 和 croypto
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建GET请求失败: %v", err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("执行GET请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("获取登录页面失败，状态码：%d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %v", err)
	}

	execMatch := executionRegexp.FindSubmatch(body)
	if len(execMatch) < 2 {
		return nil, errors.New("未能从页面中提取 execution (flowkey)")
	}
	execution := string(execMatch[1])

	croyptoMatch := croyptoRegexp.FindSubmatch(body)
	if len(croyptoMatch) < 2 {
		return nil, errors.New("未能从页面中提取 croypto")
	}
	// croypto 现在总是在第一个捕获组
	croypto := string(croyptoMatch[1])

	// 2. 使用 AES 加密密码
	encryptedPasswd, err := AesEncrypt(croypto, passwd)
	if err != nil {
		return nil, fmt.Errorf("使用AES加密密码失败: %v", err)
	}

	// 3. 构造 POST 请求
	postData := url.Values{}
	postData.Set("username", user)
	postData.Set("password", encryptedPasswd)
	postData.Set("execution", execution)
	postData.Set("croypto", croypto)
	postData.Set("type", "UsernamePassword")
	postData.Set("_eventId", "submit")
	postData.Set("geolocation", "")

	postReq, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(postData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("创建POST请求失败: %v", err)
	}

	// 4. 设置请求头
	postReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	postReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.81 Safari/537.36")
	postReq.Header.Set("Referer", URL)
	for _, c := range resp.Cookies() {
		postReq.AddCookie(c)
	}

	// 5. 发送登录请求并处理重定向
	finalResp, err := client.Do(postReq)
	if err != nil {
		return nil, fmt.Errorf("执行POST请求失败: %v", err)
	}
	defer finalResp.Body.Close()

	// 6. 检查登录是否成功
	finalURL := finalResp.Request.URL.String()
	if strings.Contains(finalURL, "sso.hdu.edu.cn") {
		return nil, errors.New("登录失败，请检查用户名或密码")
	}

	// 登录成功，返回最终的请求对象，其中包含了所有 cookies
	return finalResp.Request, nil
}
