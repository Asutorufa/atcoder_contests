package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	maps := map[rune]int{}
	for _, v := range S {
		maps[v]++
	}

	for z, v := range maps {
		if v == 1 {
			fmt.Println(string(z))
			return
		}
	}
	fmt.Println("")
}
