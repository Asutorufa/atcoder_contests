package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	for i := 0; i < 10-2; i++ {
		a := X + Y
		X = Y
		Y = reverseInt(a)
	}

	fmt.Println(Y)
}

func reverseInt(n int) int {
	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	}
	res := 0
	for n > 0 {
		res = res*10 + n%10
		n /= 10
	}
	return res * sign
}
