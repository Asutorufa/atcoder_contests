package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var S string
	fmt.Fscan(br, &S)

	data := []byte(S)

	sum := 0
	count := 0

	for i := len(data) - 1; i >= 0; i-- {
		r := m[data[i]]

		cr := 0
		if r-sum >= 0 {
			cr = r - sum
		} else {
			cr = 10 + r - sum
		}

		count += cr

		sum += cr
		if sum >= 10 {
			sum = sum % 10
		}
	}

	fmt.Println(count + len(data))
}

var m = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}
