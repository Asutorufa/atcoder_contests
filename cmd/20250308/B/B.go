package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var Q int
	fmt.Fscan(br, &Q)

	arrays := make([]int, 0, 100)

	for i := 0; i < Q; i++ {
		var t int
		fmt.Fscan(br, &t)

		switch t {
		case 1:
			var x int
			fmt.Fscan(br, &x)

			arrays = append(arrays, x)
		case 2:
			if len(arrays) == 0 {
				fmt.Println(0)
			} else {
				x := arrays[len(arrays)-1]
				arrays = arrays[:len(arrays)-1]
				fmt.Println(x)
			}
		}
	}
}
