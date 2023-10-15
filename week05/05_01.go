package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n') // 에러 발생하는 부분에 빈문자열 문자 "_" 넣어서 에러 무시 가능

	input, err := reader.ReadString('\n')
	if err != nil { // err 변수 선언 후 err 값이 nil이 아닐 경우 err값 출력 후 프로그램 종료
		log.Fatal(err)
	}
	fmt.Println(input)
	// log.Fatal(err) //에러 반환 값이 err변수에 저장되고, 에러 반환 값을 출력

	// ReadString은 2개의 값을 반환하는데 값을 값을 하나의 변수에 할당하려고 했기에 오류 발생
	// input := reader.ReadString('\n') 에러 발생코드 : assignment mismatch: 1 variable but reader.ReadString returns 2 values

	// err 변수를 선언하고 사용하지 않아 오류 발생
	// input, err := reader.ReadString('\n')

	// 43p 에러 해결 문제 예제 풀기
	fileInfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())

}
