package main

import "fmt"

// p. 182
// 슬라이스는 자체적으로 데이터를 저장하지 않으며 내부 배열의 원소에 대한 주소만 가지고 있음
// 내부 배열을 변경하면 슬라이스도 변경됨(반대도 가능)
// 여러 슬라이스가 동일한 내부 배열을 가리키고 있을 때도 동일

func main() {
	a := []string{"a", "b", "c", "d"}
	as := a[0:2] //a,b
	as[1] = "z"
	c := append(a, "y")      //메모리 공간이 늘어나지 않는다
	d := append(a, "Y", "N") //메모리 공간이 늘어난다
	fmt.Println(as)

	b := [4]int{4, 3, 2, 1}
	bs := b[1:3]
	fmt.Println(bs)

	// len이 cap 이상이면 cap공간이 자동으로 늘어난다
	a = make([]string, 4, 5) // len = 4, cap = 5
	//a := []string{"a", "b", "c", "d"}
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as = a[0:2]                    //a,b
	fmt.Println(a, len(a), cap(a)) //len = 넣은 기름양, cap = 기름통
	as[1] = "Z"                    //["a", "z"]
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("%x %x %x\n", &a[0], &as[0], &as[1])
	fmt.Println(a, len(a), cap(a)) // 4 5
	fmt.Println(c, len(c), cap(c)) // 5 8
	fmt.Println(d, len(d), cap(d)) // 6 8
	fmt.Printf("%x %x %x %x  %x\n", &a[0], &as[0], &as[1], &c[0], &d[0])

	// append 함수 사용하여 추가하면 새로운 공간에 할당하게 되어 메모리 주소가 변경되며
	// 해당 주소에 append하게 되면 기존 값에서 변경되지 않음

}
