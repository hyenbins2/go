package main

import "fmt"

func main() {
	primes := [3]int{2, 3, 5} // initialize by array literal
	//primes[2] = 6
	for i := 0; i < 3; i++ {
		fmt.Println(i, primes[i])
	}
	// 초기화 하지 않은 원소의 제로 값은 해당 원소 타입의 제로값으로 결정된다
	test := [5]bool{true, true, true}
	fmt.Println(test[3]) // zero value
	fmt.Println(test)
	//fmt.Println(test[5]) // 컴파일에러, invalid argument: index 5 out of bounds [0:5]

	i := 0
	//for i < 6 { // 런타임에러, panic: runtime error: index out of range [5] with length 5
	for i < len(test) {
		fmt.Println(test[i])
		i++
	}

	// for prime := range primes {  // index만 출력
	// 	fmt.Println(prime)
	// }

	// for idx, prime := range primes {  // idx를 사용하지 않아 컴파일 에러
	// 	fmt.Println(prime)
	// }

	for _, prime := range primes {
		fmt.Println(prime)
	}

	for idx, _ := range primes {
		fmt.Println(idx)
	}

	var sum int = 0
	for _, prime := range primes {
		//fmt.Println(prime)
		sum = sum + prime
	}
	fmt.Println(sum)
	//fmt.Println(float64(sum) / float64(len(primes)))
	fmt.Printf("%.2f\n", float64(sum)/float64(len(primes)))

	//primes[3] = 7 // index 범위 초과
	//fmt.Println("Arrays")
	//fmt.Println(primes[2])
}
