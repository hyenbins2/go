package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	month := now.Month()

	fmt.Println(year, month)
	//fmt.Println(now.Month())
	//fmt.Println(now.Hour(), now.Minute(), now.Second())

}
