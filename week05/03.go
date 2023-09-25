package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore, _ := reader.ReadString('\n')   // _ 빈 식별자 사용하여 에러 무시하기
	inputScore, err := reader.ReadString('\n') // _ 에러 처리하기

	if err != nil { //err가 nil이 아니면 err 보고하고 프로그램 종료
		log.Fatal(err)
	}

	inputScore = strings.TrimSpace(inputScore)
	score, err := strconv.ParseFloat(inputScore, 32)

	var grade string
	if score >= 90 {
		grade = "A grade"
	} else {
		grade = "nuder A grade"
	}
	//if inputScore >= 90 { //mismatched types string and untyped int > 에러 타입 다름
	//	grade := "A grade"
	//} else {
	//	grade := "nuder A grade"
	//}

	fmt.Println(inputScore)
	fmt.Println("", grade)

}
