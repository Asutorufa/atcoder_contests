package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)

	var H, W int
	fmt.Fscan(br, &H, &W)

	masu := make([]string, H)

	var S, G [2]int

	for i := range masu {
		fmt.Fscan(br, &masu[i])

		if S == [2]int{} {
			if j := strings.IndexByte(masu[i], 'S'); j != -1 {
				S = [2]int{i, j}
			}
		}

		if G == [2]int{} {
			if j := strings.IndexByte(masu[i], 'G'); j != -1 {
				G = [2]int{i, j}
			}
		}
	}

	// 缓存需要加上switch的状态
	used := map[CacheKey]int{}

	canAccess := func(x, y int, sw bool, count int) bool {
		if x < 0 || x >= H || y < 0 || y >= W {
			return false
		}

		if fu, ok := used[CacheKey{pos: [2]int{x, y}, Switch: sw}]; ok && fu <= count {
			return false
		}

		switch masu[x][y] {
		case '#':
			return false
		case 'o':
			return !sw
		case 'x':
			return sw
		}

		return true
	}

	queue := []Status{}
	queue = append(queue, Status{pos: S})
	used[Status{pos: S}.Key()] = 0

	for len(queue) > 0 {
		v := queue[0]
		// fmt.Println("---", v)
		queue = queue[1:]

		if v.pos == G {
			fmt.Println(v.Count)
			return
		}

		sw := v.Switch
		if masu[v.pos[0]][v.pos[1]] == '?' {
			sw = !sw
		}

		x := v.pos[0]
		y := v.pos[1]

		if canAccess(x-1, y, sw, v.Count+1) {
			s := Status{pos: [2]int{x - 1, y}, Switch: sw, Count: v.Count + 1}
			queue = append(queue, s)
			used[s.Key()] = s.Count
		}

		if canAccess(x+1, y, sw, v.Count+1) {
			s := Status{pos: [2]int{x + 1, y}, Switch: sw, Count: v.Count + 1}
			queue = append(queue, s)
			used[s.Key()] = s.Count
		}

		if canAccess(x, y-1, sw, v.Count+1) {
			s := Status{pos: [2]int{x, y - 1}, Switch: sw, Count: v.Count + 1}
			queue = append(queue, s)
			used[s.Key()] = s.Count
		}

		if canAccess(x, y+1, sw, v.Count+1) {
			s := Status{pos: [2]int{x, y + 1}, Switch: sw, Count: v.Count + 1}
			queue = append(queue, s)
			used[s.Key()] = s.Count
		}
	}

	fmt.Println(-1)
}

type Status struct {
	pos    [2]int
	Switch bool
	Count  int
}

func (s Status) Key() CacheKey {
	return CacheKey{
		pos:    s.pos,
		Switch: s.Switch,
	}
}

type CacheKey struct {
	pos    [2]int
	Switch bool
}
