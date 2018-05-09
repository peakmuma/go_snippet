package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main2() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}
