package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var esURL string
var totalSuccess int

func deleteESFromFile() {
	esURL = "http://172.16.100.209:9200/stuinfotable/table/"
	totalSuccess = 0
	ids := getIdsFromFile()
	for _,v := range ids {
		deleteESData(v)
	}
	fmt.Println("##########total", len(ids))
	fmt.Println("##########totalSuccess", totalSuccess)
}

func getIdsFromFile() []string{
	b, err := ioutil.ReadFile("f:/ids.txt")
	check(err)
	str := string(b)
	return strings.Split(str, "\n")
}

func deleteESData(id string) {
	client := &http.Client{}
	url := esURL + id
	req, err := http.NewRequest("DELETE", url, nil)
	check(err)
	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	check(err)
	if resp.StatusCode == 200 {
		totalSuccess++
	}
	fmt.Printf("%s \n response status %s \n response body %s \n",url, resp.Status, string(respBody))
	fmt.Println("------------------------------")
}
