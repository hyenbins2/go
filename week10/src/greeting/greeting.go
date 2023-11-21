package greeting

import "fmt"

//func hello() {
func Hello() { // 함수의 이름 첫 글자를 대문자로 해야 외부 파일에서 참조 할 수 있다
	fmt.Println("안녕하세요!")
}

func Hi() {
	fmt.Println("안녕!")
}
