package mymath

func MyPower(base int, exponent int) int {
	result := 1
	for i := 1; i <= exponent; i++ {
		result = result * base
	}
	return result
}

func MyAbs(number int) int { // 절대 값 함수
	if number < 0 { // 음수면
		return number * -1
	}
	return number
}
