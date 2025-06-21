package main

import (
	"fmt"
	"math/big"
)

var Nstr, Kstr string

var ZZZ = new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)

func main() {
	solve2()
	return

	fmt.Scan(&Nstr, &Kstr)

	OK, _ := new(big.Int).SetString(Kstr, 10)
	N, _ := new(big.Int).SetString(Nstr, 10)

	var K *big.Int

	if new(big.Int).Mod(OK, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		K = OK
	} else {
		K = new(big.Int).Sub(OK, big.NewInt(1))
	}

	if N.Cmp(K) == -1 {
		fmt.Println(1)
		return
	}

	if N.Cmp(K) == 0 {
		fmt.Println(K)
		return
	}

	// K * (2 ^ (N - K)) - (2 ^ (N - K) - 1)

	current := new(big.Int).Exp(big.NewInt(2),
		new(big.Int).Sub(N, K), nil)
	sss := new(big.Int).Sub(
		new(big.Int).Mul(K, current),
		new(big.Int).Sub(current, big.NewInt(1)),
	)

	if new(big.Int).Mod(OK, big.NewInt(2)).Cmp(big.NewInt(0)) == 1 {
		sss = sss.Add(sss, big.NewInt(1))
	}

	fmt.Println(new(big.Int).Mod(sss, ZZZ))
}

// 1000000 500001

func solve2() {
	var n, k int
	fmt.Scan(&n, &k)

	a := make([]int, n+1)
	for i := 0; i <= n; i++ {
		a[i] = 1
	}

	s := k // 滑动窗口和：一开始是前k个1相加

	for i := k; i <= n; i++ {
		a[i] = s
		s = (s - a[i-k] + a[i]) % 1000000000
		// 确保非负
		if s < 0 {
			s += 1000000000
		}
	}

	fmt.Println(a[n])
}
