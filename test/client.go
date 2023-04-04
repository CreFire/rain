package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	reqURL      = "http://localhost.com:8080"
	loginURL    = reqURL + "/end/login"
	selectURL   = reqURL + "/end/Select"
	registerURL = reqURL + "/register"
	homeURL     = reqURL + "/home"
)

func main() {
	// 登录接口测试
	data := strings.NewReader(`{"username": "admin", "password": "123456"}`)
	resp, err := http.Post(loginURL, "application/json", data)
	if err != nil {
		fmt.Println("请求登录接口失败：", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取登录接口响应失败：", err)
		return
	}
	fmt.Println("登录接口响应：", string(body))

	// 查询接口测试
	url := fmt.Sprintf("%s/%s", selectURL, "1001")
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("请求查询接口失败：", err)
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取查询接口响应失败：", err)
		return
	}
	fmt.Println("查询接口响应：", string(body))

	// 注册接口测试
	data = strings.NewReader(`{"username": "admin", "password": "123456", "email": "admin@example.com"}`)
	resp, err = http.Post(registerURL, "application/json", data)
	if err != nil {
		fmt.Println("请求注册接口失败：", err)
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取注册接口响应失败：", err)
		return
	}
	fmt.Println("注册接口响应：", string(body))

	// 主页接口测试
	url = fmt.Sprintf("%s/%s", homeURL, "1001/actions")
	resp, err = http.Get(url)
	if err != nil {
		fmt.Println("请求主页接口失败：", err)
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取主页接口响应失败：", err)
		return
	}
	fmt.Println("主页接口响应：", string(body))
}
