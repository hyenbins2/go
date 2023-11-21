package main

import "fmt"

func double(number int) {
	number = number * 2
}
func double2(number2 *int) {
	*number2 = *number2 * 2
}

func main() {
	var amount int = 5
	double(amount)
	fmt.Printf("%d\n", amount)
}
func main() {
	var amount2 int = 5
	double(&amount2)
	fmt.Printf("%d\n", amount2)
}
