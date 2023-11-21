package main

import "fmt"

func main() {
	//포인터 : 메모리 주소를 값으로 가지는 타입, &주소 연산자 *포인터 기호
	var a int = 10
	var pa *int
	pa = &a

	fmt.Println(&a, a) //
	fmt.Println(pa, *pa)
	fmt.Println(&pa)

	fmt.Printf("%T %T\n", a, pa)

}
