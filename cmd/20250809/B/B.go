package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	tarray := []int{}

	for i, v := range S {
		if v == 't' {
			tarray = append(tarray, i)
		}
	}

	if len(tarray) <= 2 {
		fmt.Println(float64(0))
		return
	}

	max := float64(0)

	for i, vi := range tarray {
		for j := len(tarray) - 1; j >= 0; j-- {
			if i >= j {
				break
			}

			vj := tarray[j]

			num := j - i + 1

			l := len(S[vi : vj+1])
			if l == 2 {
				continue
			}

			ju := float64(num-2) / float64(l-2)
			if ju > max {
				max = ju
			}
		}
	}

	fmt.Println(max)
}

//cotxxxxxxteecup
