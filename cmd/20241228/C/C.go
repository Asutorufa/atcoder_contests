package main

import (
	"fmt"
)

func main() {
	var K int
	fmt.Scanln(&K)

	var S string
	fmt.Scanln(&S)

	var T string
	fmt.Scanln(&T)

	if len(S) < len(T) {
		S, T = T, S
	}

	if len(S) > len(T) && len(S)-len(T) > 1 {
		fmt.Println("No")
		return
	}

	if len(S) == len(T) {
		var diff int
		for i := range S {
			if S[i] != T[i] {
				diff++

				if diff > K {
					fmt.Println("No")
					return
				}
			}
		}

		fmt.Println("Yes")
		return
	}

	first := true
	var i, j int
	add := func() {
		i++
		j++
	}

	for ; i < len(S); add() {
		if j >= len(T) {
			break
		}

		if S[i] != T[j] {
			if !first {
				fmt.Println("No")
				return
			}

			i++
			first = false
		}
	}

	fmt.Println("Yes")
}

/*
1
abbc
ac

1
abc
ac

1
abc
ab

1
adc
abc
*/
