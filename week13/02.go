// p.192 ~ 194
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args 값을 출력함, 다른 인자를 넣어 실행하면 다른 값이 출력 됨
	fmt.Println(os.Args) //Args == Arguments
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[2])

}

// go run 02.go 인하공업전문대학 컴퓨터정보과 오픈소스프로그래밍
// 출력 값 :
//[인하공업전문대학 컴퓨터정보과 오픈소스프로그래밍]
//컴퓨터정보과
