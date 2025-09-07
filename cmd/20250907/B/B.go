package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(br, &H, &W)

	mas := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscan(br, &mas[i])
	}

	for i := range mas {
		for j := range mas[i] {
			if mas[i][j] == '#' {
				if !ok(i, j, mas) {
					fmt.Println("No")
					return
				}
			}
		}
	}

	fmt.Println("Yes")
}

func ok(i, j int, mas []string) bool {
	w := len(mas[0])

	count := 0
	if i > 0 {
		if mas[i-1][j] == '#' {
			count++
		}
	}

	if i < len(mas)-1 {
		if mas[i+1][j] == '#' {
			count++
		}
	}

	if j > 0 {
		if mas[i][j-1] == '#' {
			count++
		}
	}

	if j < w-1 {
		if mas[i][j+1] == '#' {
			count++
		}
	}

	return count == 2 || count == 4
}
