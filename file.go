package main

import (
	// "fmt"
	"io/ioutil"
	"os"
	// "sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func main() {
	b, err := ioutil.ReadFile("e:/test.txt")
	check(err)
	str := string(b)
	//count the word
	wordMap := wordCount(str)
	//get my words
	mywords := getWordsFromRepo();
	var newWords []string
	//write new words to my word repo
	for k := range wordMap {
		if !isExist(k, mywords) {
			newWords = append(newWords, k)
		}
	}
	addWordsToRepo(newWords)

	//sort
	// var kvArray []kv
	// for k, v := range wordMap {
	// 	kvArray = append(kvArray, kv{k, v})
	// }
	// sort.Slice(kvArray, func(i, j int) bool {
	// 	return kvArray[i].Value > kvArray[j].Value
	// })

	//write to file
	// f, err := os.Create("e:/test_res.txt")
	// check(err)
	// defer f.Close()
	// for _, kvObj := range kvArray {
	// 	str = kvObj.Key + " " + strconv.Itoa(kvObj.Value) + "\n"
	// 	// fmt.Print(str)
	// 	f.WriteString(str)
	// }
	// f.Sync()

	// for k,v := range wordMap {
	// 	fmt.Println(k, v)
	// }
	// fmt.Println(str)
}

func wordCount(text string) map[string]int {
	var wordMap map[string]int = make(map[string]int, 0)
	startIndex := 0
	letterStart := false
	for i, v := range text {
		if isLetter(v) {
			if !letterStart {
				letterStart = true
				startIndex = i
			}
		} else {
			if letterStart {
				if i-startIndex > 2 {
					// 如果首个字符是大写，转换成小写添加进去
					wordMap[text[startIndex:i]]++
				}
				letterStart = false
			}
		}
	}
	return wordMap
}

func isLetter(a rune) bool {
	return (a >= 65 && a <= 90) || (a >= 97 && a <= 122) || a == '-' || a == '\''
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getWordsFromRepo() []string{
	b, err := ioutil.ReadFile("myword.txt")
	check(err)
	str := string(b)
	return strings.Split(str, "\n")
}

func addWordsToRepo(words []string) {
	// open file and append, if not exist , create it
	f, err := os.OpenFile("myword.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	for _,word := range words {
		f.WriteString(word)
		f.WriteString("\n")
	}
}

func isExist(word string, words []string) bool {
	for _,v := range words {
		if word == v {
			return true
		}
	}
	return false
}