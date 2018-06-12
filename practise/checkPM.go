package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"strconv"
)

func checkPMLog() {
	today := getTodayUnix()
	var names = []string{"wangzhuzhu", "qianlei", "baizhoujun", "xujian01", "liwenbin", "lichenglong", "yuanguopeng", "xiexiangpeng", "gaolei"}
	// var names = []string{"gaolei"}
	keyStr := "\"start\":" + strconv.FormatInt(today, 10)
	for _,name := range names {
		bodystr := getWebContent(name)
		if len(bodystr) < 10000 {
			if strings.Contains(bodystr, "user-login") {
				loginPM()
			} else {
				fmt.Println(bodystr)
				fmt.Println(name, "get web content failed")
				continue
			}
		}
		if strings.Contains(bodystr, keyStr) {
			fmt.Println(name, "write")
		} else {
			fmt.Println(name, "not write")
		}
	}
	fmt.Println()
	fmt.Println("按任意键退出...")
	fmt.Scanln()
}

func loginPM() {
	pmURL := "http://pm.shangdejigou.cn/user-login.html" 
	client := &http.Client{}
	data := url.Values{}
	data.Add("account", "gaolei")
	data.Add("password", "gaolei123456")
	data.Add("keepLogin[]", "on")
	req, _ := http.NewRequest("POST", pmURL, strings.NewReader(data.Encode())) //最后一个参数是body
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	client.Do(req)
	// resp,_ := client.Do(req)
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// check(err)
	// fmt.Println(string(body))
}

func getWebContent(name string) string {
	pmURL := "http://pm.shangdejigou.cn/user-effortcalendar-" + name + ".html" 
	client := &http.Client{}
	req, err := http.NewRequest("GET", pmURL, nil) //最后一个参数是body
	req.Header.Set("Cookie", "keepLogin=on; zentaosid=qgg91d648j80teffm06hlfivk4; downloading=1; lang=zh-cn; device=desktop; theme=default; za=gaolei; ___rl__test__cookies=1517986754570; OUTFOX_SEARCH_USER_ID_NCOO=2112521259.954677")
	check(err)
	resp, err := client.Do(req)
	check(err)
	if resp.StatusCode != 200 {
		fmt.Println(resp)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	bodystr := string(body)
	return bodystr
}

func getTodayUnix() int64 {
	t := time.Now()
	tm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return tm.Unix()
}