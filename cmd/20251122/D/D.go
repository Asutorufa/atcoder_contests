package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M uint64
	fmt.Fscan(br, &N, &M)

	var arrays = make([]uint64, 0, N)
	for range N {
		var S uint64
		fmt.Fscan(br, &S)
		arrays = append(arrays, S)
	}

	var ret = 0

	var remCounts [20]map[uint64]int
	for i := range 20 {
		remCounts[i] = make(map[uint64]int)
	}

	/*
		Target = (A*10^digits(B) + B) %M == 0

		==> (A*10^digits(B) + B) % M  = ((A*10^digits(B)) % M + B % M) % M == 0
		==>
		==> AA = (A*10^digits(B)) % M, BB = B % M
		==> (AA + BB) % M = 0
		==> AA + BB = 0 / AA + BB = M
		==> AA = M - BB
	*/

	pow10 := make([]uint64, 20)
	pow10[0] = 1
	for i := 1; i < 20; i++ {
		hi, lo := bits.Mul64(pow10[i-1], 10)
		_, r := bits.Div64(hi, lo, M)
		pow10[i] = r
	}

	for _, v := range arrays {
		for k := 1; k <= 19; k++ {
			// AA = (A*10^digits(B)) % M
			hi, lo := bits.Mul64(v, pow10[k])
			_, rem := bits.Div64(hi, lo, M)
			remCounts[k][rem]++
		}
	}

	for _, v2 := range arrays {
		length := countDigits(v2)
		// BB
		remV2 := v2 % M

		// M - BB
		targetRem := (M - remV2)
		if targetRem == M {
			targetRem = 0
		}

		// get is AA = M - BB
		if count, ok := remCounts[length][targetRem]; ok {
			ret += count
		}
	}

	fmt.Println(ret)
}

func countDigits(n uint64) int {
	switch {
	case n < 10:
		return 1
	case n < 100:
		return 2
	case n < 1000:
		return 3
	case n < 10000:
		return 4
	case n < 100000:
		return 5
	case n < 1000000:
		return 6
	case n < 10000000:
		return 7
	case n < 100000000:
		return 8
	case n < 1000000000:
		return 9
	case n < 10000000000:
		return 10
	case n < 100000000000:
		return 11
	case n < 1000000000000:
		return 12
	case n < 10000000000000:
		return 13
	case n < 100000000000000:
		return 14
	case n < 1000000000000000:
		return 15
	case n < 10000000000000000:
		return 16
	case n < 100000000000000000:
		return 17
	case n < 1000000000000000000:
		return 18
	default:
		return 19
	}
}
