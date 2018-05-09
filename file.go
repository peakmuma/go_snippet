package main

import (
	// "fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"os"
)

type kv struct {
	Key string
	Value int
}

func main() {
	b, err := ioutil.ReadFile("e:/test.txt")
	check(err)
	str := string(b)
	wordMap := wordCount(str)
	//sort
	var kvArray []kv
	for k,v := range wordMap {
		kvArray = append(kvArray, kv{k, v})
	}
	sort.Slice(kvArray, func(i,j int) bool {
		return kvArray[i].Value > kvArray[j].Value
	})

	//write to file
	f, err := os.Create("e:/test_res.txt")
	check(err)
	defer f.Close()
	for _, kvObj := range kvArray {
		str = kvObj.Key + " " + strconv.Itoa(kvObj.Value) + "\n"
		// fmt.Print(str)
		f.WriteString(str)
	}
	f.Sync()
	// for k,v := range wordMap {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(str)
}

func wordCount(text string) map[string]int {
	var wordMap map[string]int = make(map[string]int, 0)
	startIndex := 0
	letterStart := false
	for i,v := range text {
		if isLetter(v) {
			if !letterStart {
				letterStart = true
				startIndex = i
			}
		} else {
			if letterStart {
				if i - startIndex > 2 {
					wordMap[text[startIndex:i]]++					
				}
				letterStart=false
			}
		}
	}
	return wordMap
}


func isLetter (a rune) bool{
	return (a >= 65 && a <= 90) || (a >= 97 && a <= 122) || a == '-' || a == '\''
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
