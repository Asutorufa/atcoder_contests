package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var N int
	fmt.Fscan(br, &N)

	wakata := make([]int, N)

	skillDep := make(map[int]map[int]struct{})

	ret := make(map[int]struct{}, N)

	for i := 1; i <= N; i++ {
		var a, b int
		fmt.Fscan(br, &a, &b)

		if a == 0 && b == 0 {
			wakata = append(wakata, i)
			ret[i] = struct{}{}
		} else {
			if skillDep[a] == nil {
				skillDep[a] = map[int]struct{}{}
			}
			if skillDep[b] == nil {
				skillDep[b] = map[int]struct{}{}
			}

			skillDep[a][i] = struct{}{}
			skillDep[b][i] = struct{}{}
		}
	}

	// fmt.Println(skillDep)

	for _, v := range wakata {
		learn(ret, skillDep, v)
	}

	fmt.Println(len(ret))
}

func learn(m map[int]struct{}, skillDep map[int]map[int]struct{}, v int) {
	for k := range skillDep[v] {
		if _, ok := m[k]; ok {
			continue
		}

		m[k] = struct{}{}

		learn(m, skillDep, k)
	}
}
