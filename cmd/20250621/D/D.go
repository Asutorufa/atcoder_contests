package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Query struct {
	Opt, Pc int
	S       string
}

func main() {
	br := bufio.NewReader(os.Stdin)

	var N, Q int
	fmt.Fscan(br, &N, &Q)

	var queries []Query
	for i := 0; i < Q; i++ {
		var opt, pc int
		fmt.Fscan(br, &opt, &pc)

		switch opt {
		case 1, 3:
			queries = append(queries, Query{Opt: opt, Pc: pc})
		case 2:
			var s string
			fmt.Fscan(br, &s)

			queries = append(queries, Query{Opt: opt, Pc: pc, S: s})
		}
	}

	ans := []string{}
	i := 0

	for j := Q - 1; j >= 0; j-- {
		q := queries[j]
		switch q.Opt {
		// pc 的变为server的
		case 1:
			// 上面pc添加时不能--，会导致这里出问题
			if i == q.Pc {
				i = 0
			}
		case 2:
			if i == q.Pc {
				ans = append(ans, q.S)
			}
		case 3:
			// server 的变成pc的
			if i == 0 {
				i = q.Pc
			}
		}
	}

	reverse(ans)
	fmt.Println(strings.Join(ans, ""))
}

func reverse(strs []string) {
	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}
}
