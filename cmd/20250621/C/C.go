package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(br, &N, &Q)

	masu := make([]int, N)

	count := 0

	for i := 0; i < Q; i++ {
		var A int
		fmt.Fscan(br, &A)

		A--

		// 0 is white
		// 1 is black
		if masu[A] == 0 {
			if N == 1 {
				count++
			} else {
				if A != 0 && A != N-1 {
					if masu[A-1] == 0 && masu[A+1] == 0 {
						count++
					} else if masu[A-1] == 1 && masu[A+1] == 1 {
						count--
					}
				} else {
					if A == 0 {
						if masu[A+1] == 0 {
							count++
						}
					} else {
						if masu[A-1] == 0 {
							count++
						}
					}
				}
			}

			masu[A] = 1
		} else {
			if N == 1 {
				count--
			} else {
				if A != 0 && A != N-1 {
					if masu[A-1] == 1 && masu[A+1] == 1 {
						count++
					} else if masu[A-1] == 0 && masu[A+1] == 0 {
						count--
					}
				} else {
					if A == 0 {
						if masu[A+1] == 0 {
							count--
						}
					} else {
						if masu[A-1] == 0 {
							count--
						}
					}
				}
			}
			masu[A] = 0
		}

		fmt.Println(count)
	}
}

/*
2 4
1 2 1 2


3 6
1 3 2 1 3 2
1
2
1
1
1
0
*/
