// Package datafile allows reading data samples from files. 5장 배열
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// p.166 텍스트 파일과 배열
// GetFloats reads a float64 from each line of a file.
func GetFloats(fileName string) ([4]float64, error) {
	var numbers [4]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
