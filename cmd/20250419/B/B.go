package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var Q int
	fmt.Fscan(br, &Q)

	queue := list.New()
	for i := 0; i < Q; i++ {
		var z int
		fmt.Fscan(br, &z)

		switch z {
		case 1:
			var X int
			fmt.Fscan(br, &X)

			queue.PushBack(X)
		case 2:
			x := queue.Front()
			queue.Remove(x)
			fmt.Println(x.Value)
		}
	}
}
