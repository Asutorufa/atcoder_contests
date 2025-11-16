package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	ret := 0
	for i1, v1 := range []int{A, B, C} {
		for i2, v2 := range []int{A, B, C} {
			for i3, v3 := range []int{A, B, C} {
				if i1 != i2 && i2 != i3 && i1 != i3 {
					x := v1*100 + v2*10 + v3
					if x > ret {
						ret = x
					}
				}
			}
		}
	}

	fmt.Println(ret)
}
