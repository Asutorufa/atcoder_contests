package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	var S, T string
	fmt.Fscan(br, &S, &T)

	Sb := []byte(S)
	Tb := []byte(T)

	diff := make([]int, N+2)

	for i := 0; i < M; i++ {
		var L, R int
		fmt.Fscan(br, &L, &R)
		diff[L]++
		diff[R+1]--
	}

	swap := 0
	for i := 1; i <= N; i++ {
		swap += diff[i]
		if swap%2 == 1 {
			Sb[i-1], Tb[i-1] = Tb[i-1], Sb[i-1]
		}
	}

	fmt.Println(string(Sb))
}

// func main() {
// 	br := bufio.NewReader(os.Stdin)

// 	var N, M int
// 	fmt.Fscan(br, &N, &M)

// 	var S, T string
// 	fmt.Fscan(br, &S, &T)

// 	Sb := []byte(S)
// 	Tb := []byte(T)
// 	var Cache = make([]byte, 0, N)

// 	for i := 0; i < M; i++ {
// 		var L, R int
// 		fmt.Fscan(br, &L, &R)

// 		Cache = Cache[:R-L+1]
// 		copy(Cache, Sb[L-1:R])
// 		copy(Sb[L-1:R], Tb[L-1:R])
// 		copy(Tb[L-1:R], Cache)
// 	}

// 	fmt.Println(string(Sb))
// }
