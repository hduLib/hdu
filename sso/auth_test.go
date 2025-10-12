package sso

import (
	"net/url"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	// 从环境变量中读取用户名和密码，避免硬编码在代码中
	// 请在测试前设置好环境变量 HDU_USER 和 HDU_PASS
	user := os.Getenv("HDU_USER")
	pass := os.Getenv("HDU_PASS")

	if user == "" || pass == "" {
		t.Skip("请设置环境变量 HDU_USER 和 HDU_PASS 以进行登录测试")
	}

	// 目标 service URL，这里使用信息门户作为测试目标
	serviceURL := "https://i.hdu.edu.cn/tp_up"
	loginURL := "https://sso.hdu.edu.cn/login?service=" + serviceURL

	t.Logf("正在尝试为用户 %s 登录...", user)

	// 调用 GenLoginReq 函数
	req, err := GenLoginReq(loginURL, user, pass)
	if err != nil {
		t.Fatalf("GenLoginReq 失败: %v", err)
	}

	// 验证结果
	if req == nil {
		t.Fatal("GenLoginReq 返回了一个 nil 请求")
	}

	finalURL := req.URL.String()
	t.Logf("成功跳转到 URL: %s", finalURL)

	// 检查最终的 URL 域名是否是我们期望的 service URL 的域名
	finalURLObj := req.URL
	serviceURLObj, err := url.Parse(serviceURL)
	if err != nil {
		t.Fatalf("无法解析 serviceURL: %v", err)
	}
	if finalURLObj.Host != serviceURLObj.Host {
		t.Errorf("期望跳转到域名 %s，但实际跳转到了 %s", serviceURLObj.Host, finalURLObj.Host)
	}

	// 检查是否获取到了 Cookies
	cookies := req.Cookies()
	if len(cookies) == 0 {
		t.Error("未能获取到任何 Cookies")
	}

	t.Logf("成功获取到 %d 个 Cookies", len(cookies))
	var foundCASTGC bool
	for _, c := range cookies {
		t.Logf("  - [名称: %s, 域: %s]", c.Name, c.Domain)
		if c.Name == "CASTGC" {
			foundCASTGC = true
		}
	}

	if !foundCASTGC {
		t.Log("警告：在最终请求的 Cookies 中未找到 'CASTGC'，但这通常是正常的，因为该 Cookie 属于 sso.hdu.edu.cn 域。")
	}
}
