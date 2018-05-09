package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

type kv struct {
	Key string
	Value int
}

func main() {
	b, err := ioutil.ReadFile("e:/test.txt")
	if err != nil {
		fmt.Print(err)
	}
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
	for _, kvObj := range kvArray {
		fmt.Println(kvObj.Key, kvObj.Value)
	}
	// for k,v := range wordMap {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(str)
}

func wordCount(text string) map[string]int {
	token := []rune{' ', ',', ';', '.', '(', ')', '{', '}' , '[' , ']' , '\t', '\r', '\n', '"', ':', '+', '/', '\\', '=', '$'};
	var wordMap map[string]int = make(map[string]int, 0)
	startIndex := 0
	letterStart := false
	for i,v := range text {
		if !runeInArray(v, token) {
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

func runeInArray(a rune, list []rune) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
