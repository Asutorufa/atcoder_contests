package main

import "fmt"

func main() {
	var S, T string
	fmt.Scan(&S, &T)

	includes := []rune{}
	first := true
	for i, v := range S {
		if v >= 'A' && v <= 'Z' {
			if first {
				first = false
				continue
			}

			includes = append(includes, rune(S[i-1]))
		}
	}

	Tmaps := map[rune]bool{}
	for _, v := range T {
		Tmaps[v] = true
	}

	for _, v := range includes {
		if !Tmaps[v] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
