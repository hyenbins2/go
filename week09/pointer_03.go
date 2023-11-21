package main

import "fmt"

//pass by pointer (call by pointer)
func swap() {
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func main() {

	a := 10
	b := 20

	c := &a
	// intger  타입의 변수 주소는 intger 타입의 포인터 변수로만 받을수있다
	fmt.Printf("%T %T\n", c) //*int

	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

}
