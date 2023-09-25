package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "cs r?cks~"
	replaceWords := strings.NewReplacer("?", "o")
	fixedWord := replaceWords.Replace(brokenWords)
	fmt.Println(fixedWord)

}
