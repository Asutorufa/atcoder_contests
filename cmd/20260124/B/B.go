package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	q := readInt()

	play := false
	volume := 0
	for i := 0; i < q; i++ {
		cmd := readInt()
		switch cmd {
		case 1:
			volume++
		case 2:
			if volume > 0 {
				volume--
			}
		case 3:
			play = !play
		}

		if play && volume >= 3 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

var br = bufio.NewReader(os.Stdin)

func readInt() int {
	var x int
	fmt.Fscan(br, &x)
	return x
}

func readString() string {
	var x string
	fmt.Fscan(br, &x)
	return x
}

func readX[T any]() T {
	var x T
	fmt.Fscan(br, &x)
	return x
}
