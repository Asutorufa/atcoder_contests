package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	mas := make([][]string, N)
	for i := 0; i < N; i++ {
		mas[i] = make([]string, N)
	}

	for i := 0; i < N; i++ {
		j := N - i

		// fmt.Println(i, j)

		for ii := i; ii < j; ii++ {
			for jj := i; jj < j; jj++ {
				// fmt.Println(ii, jj)
				if (i+1)%2 == 0 {
					mas[ii][jj] = "."
				} else {
					mas[ii][jj] = "#"
				}
			}
		}

		// fmt.Println()

		// for _, v := range mas {
		// 	fmt.Println(strings.Join(v, ""))
		// }
	}

	for _, v := range mas {
		fmt.Println(strings.Join(v, ""))
	}
}
