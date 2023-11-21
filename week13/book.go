// github에서 가져옴
// 메모장 text 파일에쓴 값을 사용해서 출력하는 버전
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers, err := GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
func GetFloats(fileName string) ([]float64, error) { //check
	var numbers []float64 //check
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		//make 생략하고 append로 넣음 append는 기존에 있는 슬라이스에 make하는 기능이 있음
		// 배열의 개수가 고정되어있는 문제점을 해결함
		numbers = append(numbers, number) //update 된 곳 check
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
