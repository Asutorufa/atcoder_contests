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

	str := make([]int, N)
	for i := 0; i < N; i++ {
		str[i] = i + 1
	}

	offset := 0

	for i := 0; i < Q; i++ {
		var t, a1 int
		fmt.Fscan(br, &t, &a1)

		switch t {
		case 1:
			var a2 int
			fmt.Fscan(br, &a2)

			q := a1 - 1

			a1 = q + offset
			if a1 > N-1 {
				a1 %= N
			}

			str[a1] = a2
		case 2:
			q := a1 - 1

			a1 = q + offset
			if a1 > N-1 {
				a1 %= N
			}

			fmt.Println(str[a1])
		case 3:
			offset += a1
			offset = offset % N
		}
	}
}

/*
5 6
1 3 555
2 3
1 2 1000000
3 4
2 2
2 3

0 1       2 3 4
1,1000000,3,4,5
-4
5,1,1000000,3,4
0 1  2  3  4

4+-4 = 0
3+-4 =-1+5=4

1 + -4 = -3


0 1 2 3 4
0 -1 -2 -3 -4


*/
