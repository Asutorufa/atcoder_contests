package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscanln(br, &H, &W)

	StartX, StartY := -1, -1
	EndX, EndY := -1, -1

	masu := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Fscanln(br, &masu[i])

		for j := 0; j < W; j++ {
			if masu[i][j] != '#' {
				continue
			}

			if StartX == -1 {
				StartX = j
			} else if StartX > j {
				StartX = j
			}

			if StartY == -1 {
				StartY = i
			}

			if EndX == -1 {
				EndX = j
			} else if EndX < j {
				EndX = j
			}

			if EndY == -1 {
				EndY = i
			} else if EndY < i {
				EndY = i
			}
		}
	}

	// fmt.Println(StartX, StartY, EndX, EndY)

	for _, v := range masu[StartY : EndY+1] {
		// fmt.Println(v[StartX : EndX+1])
		for _, z := range v[StartX : EndX+1] {
			if z == '.' {
				fmt.Println("No")
				return
			}
		}
	}

	fmt.Println("Yes")
}

/*
3 3
??#
???
???

*/
