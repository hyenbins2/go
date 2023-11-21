// p. 186
// slice와 zero value 값
package main

import "fmt"

func main() {
	var a []string //값을 할당하지 않았으니 출력 시 nil값(0) 나옴
	var b []bool
	//a = make([]string, 4, 5)

	b = append(b, true)           //bool값 = true, true를쓰게 되면 메모리 공간 할당 됨
	fmt.Printf("%#v %#v\n", a, b) // zero value of empty slice

	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))
}
