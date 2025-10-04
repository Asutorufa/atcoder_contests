package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type counter struct {
	id    int
	count int
}

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(br, &N, &Q)

	maps := map[int]*list.Element{}
	l := list.New()

	for i := 1; i <= N; i++ {
		item := &counter{id: i, count: 1}
		e := l.PushBack(item)
		maps[i] = e
	}

	for i := 0; i < Q; i++ {
		ret := 0
		var X, Y int
		fmt.Fscan(br, &X, &Y)

		now := maps[X]
		tar := maps[Y]
		if tar == nil || now == nil {
			fmt.Println(0)
			continue
		}
		tc := tar.Value.(*counter)

		for now != nil {
			cc := now.Value.(*counter)
			pre := now.Prev()
			l.Remove(now)

			delete(maps, cc.id)
			now = pre

			if count := cc.count; count > 0 {
				tc.count += count
				ret += count
			}
		}

		fmt.Println(ret)
	}
}

/*
8 5
2 6
3 5
1 7
5 7
7 8

8 3
4 6
2 4
6 7
*/
