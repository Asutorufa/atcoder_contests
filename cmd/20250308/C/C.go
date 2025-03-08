package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(br, &N, &M)

	var Bs []int = make([]int, 0, N)
	var Ws []int = make([]int, 0, M)

	for i := 0; i < N; i++ {
		var B int
		fmt.Fscan(br, &B)

		Bs = append(Bs, B)
	}

	for i := 0; i < M; i++ {
		var W int
		fmt.Fscan(br, &W)

		Ws = append(Ws, W)
	}

	sort.Slice(Bs, func(i, j int) bool { return Bs[i] > Bs[j] })
	sort.Slice(Ws, func(i, j int) bool { return Ws[i] > Ws[j] })

	sum := 0
	var i, w int
	for i, w = range Ws {
		if M > N && i >= N {
			break
		}

		if w < 0 {
			break
		}

		b := Bs[i]

		if sum < sum+w+b {
			sum = sum + w + b
		} else {
			break
		}
	}

	if w >= 0 {
		i++
	}

	for ; i < N && Bs[i] > 0; i++ {
		sum += Bs[i]
	}

	fmt.Println(sum)
}

/*
3 5
3 2 1
3 2 1 4 5

3+4+5+3+2+1 = 18


5 3
3 2 1 4 5
3 2 1
*/
