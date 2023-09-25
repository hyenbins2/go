package main

import (
	"fmt"
	"math/rand"
)

func main() {
	dice := rand.Intn(6) + 1 // 1 ~ 6 사이의 정수를 만들기 위해 + 1하기
	fmt.Println(dice)

}
