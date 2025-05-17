package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	// fmt.Println(big.NewInt(0).Mul(
	// 	big.NewInt(int64(math.Pow(10, 18))), big.NewInt(int64(math.Pow(10, 18)))))

	br := bufio.NewReader(os.Stdin)

	var N, K int
	fmt.Fscan(br, &N, &K)

	now := big.NewInt(1)

	for i := 0; i < N; i++ {
		var A int
		fmt.Fscan(br, &A)

		now = big.NewInt(0).Mul(now, big.NewInt(int64(A)))

		if len(now.String()) > K {
			now = big.NewInt(1)
		}
	}

	fmt.Println(now)
}
