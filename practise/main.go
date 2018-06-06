package main
import (
	// "time"
	// "fmt"
)

func main() {
	checkPMLog()
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}