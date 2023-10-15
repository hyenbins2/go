package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 시드 설정

	answer := rand.Intn(100) + 1 // 1 ~ 100

	fmt.Println("Guess number (1 ~ 100) : ")
	// fmt.Println(answer) // 디버깅용으로 정답 숫자 출력

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ { // 루프 조건 수정
		fmt.Println("Your chance :", 10-i)
		fmt.Println("Guess number: ")
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNumberString = strings.TrimSpace(inputNumberString) // 입력 문자열 공백 제거
		inputNumber, err := strconv.Atoi(inputNumberString)      // 문자열을 정수로 변환

		if err != nil {
			log.Fatal(err)
		}
		if inputNumber < answer {
			fmt.Println("Your guess number is lower than the answer.")
		} else if inputNumber > answer {
			fmt.Println("Your guess number is higher than the answer.")
		} else {
			fmt.Println("Congratulations! You guessed the correct number.")
			break // 정답을 맞췄으므로 루프 종료
		}
	}
	fmt.Println("The correct answer was:", answer)
}
