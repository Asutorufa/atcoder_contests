package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	As := make([]int, N)
	Bs := make([]int, M)
	for i := range As {
		fmt.Fscan(br, &As[i])
	}
	for i := range Bs {
		fmt.Fscan(br, &Bs[i])
	}

	sort.Ints(Bs)

	bprefix := make([]int64, M+1)
	for i := 0; i < M; i++ {
		bprefix[i+1] = bprefix[i] + int64(Bs[i])
	}

	var ret int64 = 0

	for _, A := range As {
		k := sort.SearchInts(Bs, A+1)

		left := int64(A)*int64(k) - bprefix[k]
		right := (bprefix[M] - bprefix[k]) - int64(A)*int64(M-k)

		ret += left + right
		ret %= 998244353
	}

	fmt.Println(ret)
}
