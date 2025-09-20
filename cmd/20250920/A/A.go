package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	maps := map[int]int{}

	maps[a]++
	maps[b]++
	maps[c]++

	for _, v := range maps {
		if v >= 2 {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
}
