package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W, X, Y int
	fmt.Fscanln(br, &H, &W, &X, &Y)

	mas := []string{}
	for i := 0; i < H; i++ {
		var s string
		fmt.Fscanln(br, &s)
		mas = append(mas, s)
	}

	X--
	Y--

	var T string
	fmt.Fscanln(br, &T)

	var count int
	cache := map[string]bool{}

	// fmt.Println("T---", T)
	for _, v := range T {

		// fmt.Println(X+1, Y+1)

		if mas[X][Y] == '@' && !cache[fmt.Sprintf("%d-%d", X, Y)] {
			count++
			cache[fmt.Sprintf("%d-%d", X, Y)] = true
		}

		switch v {
		case 'L':
			if Y > 0 {
				if mas[X][Y-1] == '#' {
					continue
				}

				Y--
			}
		case 'R':
			if Y < W-1 {
				if mas[X][Y+1] == '#' {
					continue
				}
				Y++
			}
		case 'U':
			if X > 0 {
				if mas[X-1][Y] == '#' {
					continue
				}
				X--
			}
		case 'D':
			if X < H-1 {
				if mas[X+1][Y] == '#' {
					continue
				}
				X++
			}
		}

		if mas[X][Y] == '@' && !cache[fmt.Sprintf("%d-%d", X, Y)] {
			count++
			cache[fmt.Sprintf("%d-%d", X, Y)] = true
		}
	}

	fmt.Println(X+1, Y+1, count)
}
