// average2 calculates the average of several numbers.
// 메모장 사용 안하고 go run 시 뒤에 값 붙여서 실행하는 버전
// p.195 ~ 200
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 { ///... : 가변 매개변수(입력 받을 인자의 개수를 가변적으로 만들 수 있음)
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("Average: %0.2f\n", average(numbers...))
}

//go run book2.go 90 80 70
// Average: 80.00
