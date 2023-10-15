package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "cs r?cks~"
	replaceWords := strings.NewReplacer("?", "o")  //모든 "?"를 "o"로 치환하도록 반환
	fixedWord := replaceWords.Replace(brokenWords) // strings.Replacer의Replace메서드 호출하여 치환할 문자열 전달
	fmt.Println(fixedWord)                         //반환된 문자열 출력

}
