package main

import (
	"fmt"

	"main.go/src/greeting"
	"main.go/src/mymath"
)

func main() {
	fmt.Println(mymath.MyPower(2, 10))
	greeting.Hello()
	fmt.Println(mymath.MyAbs(99))
	fmt.Println(mymath.MyAbs(-7))
	greeting.Hi()
}
