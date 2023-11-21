// Package datafile allows reading data samples from files.
// 7장 sample code
// p.210 직접 쳐보기ㅁㅁㅁㅁㅁ
package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	lines, err := GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	//책 p. 210 before
	counts := make(map[string]int) //after, make : slice 만들때, map 만들 때 사용
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}

// GetStrings reads a string from each line of a file.
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
