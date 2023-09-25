package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	rand.Intn(Time.Now) + 1
	answer := rand.Intn(100) + 1 // 1 ~ 100

	fmt.Println("Guess number (1  ~ 100) : ")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; 1 < 10; i++ {
		fmt.Println("You chance :", 10-i)
		fmt.Println("Guess number: ")
		inputNumber, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		inputNumberString := strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString)

		if err != nil {
			log.Fatal(err)
		}
		if inputNumber < answer {
			fmt.Println("Ur guess number is lower than answer.")
		} else if inputNumber > answer {
			fmt.Println("Ur guess number is higther than answer.")
		}
	}
}
