package main

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
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil) //最后一个参数是body
	check(err)
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	fmt.Printf(string(body))
}
