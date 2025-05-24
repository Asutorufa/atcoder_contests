package main

import (
	"fmt"
)

// 暴力解。。。。
func main() {
	var H, W int
	fmt.Scan(&H, &W)

	A := make([][]uint64, H)
	for i := range A {
		A[i] = make([]uint64, W)
		for j := range A[i] {
			fmt.Scan(&A[i][j])
		}
	}

	type Board = [][]bool

	// 初始化：空板
	initBoard := make(Board, H)
	for i := range initBoard {
		initBoard[i] = make([]bool, W)
	}

	var copyBoard = func(b Board) Board {
		newB := make(Board, H)
		for i := range b {
			newB[i] = make([]bool, W)
			copy(newB[i], b[i])
		}
		return newB
	}

	// 保存所有合法覆盖方式
	var boards []Board
	boards = append(boards, initBoard)

	// 枚举每个格子，在当前位置尝试放置骨牌
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			var next []Board
			for _, b := range boards {
				// 如果当前位置已经被覆盖，则跳过
				if b[i][j] {
					continue
				}
				// 尝试横放（右边）
				if j+1 < W && !b[i][j+1] {
					nb := copyBoard(b)
					nb[i][j], nb[i][j+1] = true, true
					next = append(next, nb)
				}
				// 尝试竖放（下边）
				if i+1 < H && !b[i+1][j] {
					nb := copyBoard(b)
					nb[i][j], nb[i+1][j] = true, true
					next = append(next, nb)
				}
			}
			boards = append(boards, next...)
		}
	}

	// 遍历所有覆盖状态，计算未覆盖位置的异或和
	var ans uint64
	for _, b := range boards {
		var xorSum uint64
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				if !b[i][j] {
					xorSum ^= A[i][j]
				}
			}
		}
		if xorSum > ans {
			ans = xorSum
		}
	}

	fmt.Println(ans)
}
