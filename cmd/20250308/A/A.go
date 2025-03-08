package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	var last = -1
	var count = 0
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)

		if A != last {
			last = A
			count = 1
		} else {
			count++
			if count == 3 {
				fmt.Println("Yes")
				return
			}
		}
	}

	fmt.Println("No")
}
