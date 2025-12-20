package main

import "fmt"

func main() {
	var H, W, N int
	fmt.Scan(&H, &W, &N)

	hmap := map[int]int{}

	for h := range H {
		for range W {
			var X int
			fmt.Scan(&X)
			hmap[X] = h
		}
	}

	hc := map[int]int{}
	ret := 0
	for range N {
		var B int
		fmt.Scan(&B)

		_, ok := hmap[B]
		if !ok {
			continue
		}

		hc[hmap[B]] += 1
		if hc[hmap[B]] > ret {
			ret = hc[hmap[B]]
		}
	}

	fmt.Println(ret)
}
