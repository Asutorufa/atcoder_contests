package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	mas := make([][]int, N)
	for i := range mas {
		mas[i] = make([]int, N)
	}

	r, c := 0, (N-1)/2
	k := 1
	mas[r][c] = k

	for range (N * N) - 1 {
		mr := (((r - 1) % N) + N) % N
		mc := (((c + 1) % N) + N) % N

		if mas[mr][mc] == 0 {
			r, c = mr, mc
			mas[mr][mc] = k + 1
			k++
		} else {
			r = (((r + 1) % N) + N) % N
			mas[r][c] = k + 1
			k++
		}
	}

	for i := range mas {
		for j := range mas[i] {
			fmt.Printf("%d ", mas[i][j])
		}
		fmt.Println()
	}
}
