package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// p.166 텍스트 파일과 배열
func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) //file의 위치 : scanner
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())

	}
}
