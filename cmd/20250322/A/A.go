package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	buf := make([]byte, N)

	for i := range buf {
		buf[i] = '-'
	}

	if N%2 == 0 {
		buf[N/2-1] = '='
		buf[N/2] = '='
	} else {
		buf[N/2] = '='
	}

	fmt.Println(string(buf))
}
