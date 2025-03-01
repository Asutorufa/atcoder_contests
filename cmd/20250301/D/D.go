package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReaderSize(os.Stdin, 655350)

	var N, Q int
	fmt.Fscan(br, &N, &Q)

	// hato -> suid
	hato := make([]int, N+1)
	for i := 0; i < N+1; i++ {
		hato[i] = i
	}
	// su -> suid
	su := make([]int, N+1)
	copy(su, hato)
	// suid -> su
	suid := make([]int, N+1)
	copy(suid, hato)

	for i := 0; i < Q; i++ {
		var Op int
		fmt.Fscan(br, &Op)

		switch Op {
		case 1:
			var a, b int
			fmt.Fscan(br, &a, &b)

			hato[a] = su[b]

		case 2:
			var a, b int
			fmt.Fscan(br, &a, &b)

			suid[su[a]], suid[su[b]] = suid[su[b]], suid[su[a]]
			su[a], su[b] = su[b], su[a]

		case 3:
			var a int
			fmt.Fscan(br, &a)
			fmt.Println(suid[hato[a]])
		}
	}
}
