package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}

func httpDo() {
	client := &http.Client{}
	url := "http://pm.shangdejigou.cn/user-login-L3VzZXItZWZmb3J0Y2FsZW5kYXItbGljaGVuZ2xvbmcuaHRtbA==.html"
	req, err := http.NewRequest("GET", url, nil) //最后一个参数是body
	check(err)
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Printf(string(body))
}
