package main

//14p
import (
	"fmt"
)

/*
	func main() {
		// 오류 발생 code
		// fmt.Println(math.Floor("inha"))
		// fmt.Println(strings.Title("3.14"))

		fmt.Println('A')              //rune unicode
		fmt.Println('김', '\n')        //rune unicode
		fmt.Println(math.Floor(2.15)) //(내림함수)
		fmt.Println(math.Ceil(2.15))  //(올림함수)
		fmt.Println(strings.Title("오픈소스프로그래밍~\n\"Go\""))
		fmt.Println(strings.Title("Open Sorce Progamming Go!~"))

}
*/
/* func main() {

	var a int = 9
	var b float32 = 3.14
	var c, d string
	var f string
	var g bool // 빈 문자열 = false

	//a = 9
	//b = 2.7
	c = "inha~"
	d = "Go ----"
	h := '2'
	i := "문자열" //:= >> 단축 변수 선언
	J := "변수명이 대문자로 시작하면 다른 패키지에서 이 변수를 사용할 수 있음"
	korean := "이혜빈"

	fmt.Println(a, b, c, d, f, g, h, i, J, korean)
	fmt.Println(reflect.TypeOf(b)) //f, g는 zero value 출력
	fmt.Println(reflect.TypeOf(h)) //int32
	fmt.Println(reflect.TypeOf(i)) //string

}
*/
func main() {

	var taxRate float64 = 0.08
	var tax float64 = taxRate

	fmt.Println("Tax is", tax, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars avilable")
	fmt.Println("Width")
}
