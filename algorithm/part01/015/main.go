package main

import "fmt"

func main() {
	num := divide(-2147483648, 1)
	fmt.Println(num)
}

func divide(dividend int, divisor int) int {
	sign := 0
	answ := 1
	maxi := 1<<31 - 1
	mini := 0 - maxi - 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor
	}

	dividendOrigin := dividend
	divisorOrigin := divisor

	if dividend < divisor {
		return 0
	}

	dividend = dividend - divisor
	for dividend > divisor {
		dividend = dividend - divisor
		divisor = divisor + divisor
		if dividend < 0 {
			break
		}
		answ = answ + answ
	}

	answ = answ + divide(dividendOrigin-divisor, divisorOrigin)

	if sign < 0 {
		answ = 0 - answ
	}
	if answ > maxi || answ < mini {
		return maxi
	}
	return answ
}
