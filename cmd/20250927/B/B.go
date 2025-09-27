package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	Ps := map[int]struct{}{}

	for i := 1; i <= N; i++ {
		Ps[i] = struct{}{}
	}

	As := make([]int, 0, N)
	for i := 1; i <= N; i++ {
		var a int
		fmt.Scan(&a)
		As = append(As, a)
		if a != -1 {
			if _, ok := Ps[a]; ok {
				delete(Ps, a)
			} else {
				fmt.Println("No")
				return
			}
		}
	}

	fmt.Println("Yes")

	for _, v := range As {
		if v == -1 {
			for k := range Ps {
				fmt.Print(k, " ")
				delete(Ps, k)
				break
			}
		} else {
			fmt.Print(v, " ")
		}
	}

	fmt.Println()
}
