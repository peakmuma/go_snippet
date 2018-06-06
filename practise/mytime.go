package main
import (
	"time"
	"fmt"
)

func testTime() {
	t := time.Now()
	tm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	fmt.Print(tm.Unix())
}