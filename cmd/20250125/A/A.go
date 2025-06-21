package main

import "fmt"

func main() {
	var A1, A2, A3, A4, A5 int
	fmt.Scanln(&A1, &A2, &A3, &A4, &A5)

	arrs := []int{A1, A2, A3, A4, A5}

	count := 0
	last := -1
	for _, v := range arrs {
		if last == -1 {
			last = v
			continue
		}

		if last > v {
			count++
			if count > 1 {
				fmt.Println("No")
				return
			}
		} else {
			last = v
		}
	}

	if count == 1 {
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")
}
