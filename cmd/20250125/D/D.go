package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.ListenAndServe(":8080", nil)
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscanln(br, &N)

	var z []int
	readIntFunc(N, br, func(x, _ int) {
		z = append(z, x)
	})

	resp := map[int]struct{}{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j == i {
				continue
			}

			zz := -1
			sum := z[i] + z[j]

			for in, v := range z {
				if zz == -1 {
					zz = v
					continue
				}
				if in == i {
					zz ^= 0
					continue
				}

				if in == j {
					zz ^= sum
					continue
				}

				zz ^= v
			}

			fmt.Println(zz)

			resp[zz] = struct{}{}
		}
	}

	fmt.Println(len(resp))
}

func readIntFunc(n int, r *bufio.Reader, f func(_ int, index int)) {
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(r, &x)
		f(x, i)
	}
}
