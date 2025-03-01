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

	last := -1
	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)
		if last == -1 {
			last = A
			continue
		}

		if last < A {
			last = A
			continue
		}

		fmt.Println("No")
		return
	}

	fmt.Println("Yes")
}
