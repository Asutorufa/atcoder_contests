package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDanger(S string) map[int]bool {
	danger := make(map[int]bool)
	for i := 0; i < len(S); i++ {
		if S[i] == '1' {
			danger[i+1] = true
		}
	}

	return danger

}

func Resolve(c int, S string) bool {
	danger := getDanger(S)

	yakuhinn := []int{}

	for i := 1; i <= c; i++ {
		yakuhinn = append(yakuhinn, i)
	}

	// クソーーー
	// memoが忘れた, memoがあれば　こんな事って。。。
	return dfs(yakuhinn, make([]bool, len(yakuhinn)), NewBitCalculator(), danger, map[int]bool{})
}

func dfs(nums []int, used []bool, path *BitCalculator, danger map[int]bool, memo map[int]bool) bool {
	if path.length == len(nums) {
		return true
	}

	if val, ok := memo[path.Value()]; ok {
		return val
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		num := nums[i] - 1

		if danger[path.Value()|(1<<num)] {
			continue
		}

		path.AddBit(num)
		used[i] = true

		if dfs(nums, used, path, danger, memo) {
			return true
		}

		path.Undo(num)
		used[i] = false
	}

	memo[path.Value()] = false

	return false
}

type BitCalculator struct {
	value  int
	length int
}

func NewBitCalculator() *BitCalculator { return &BitCalculator{value: 0} }
func (bc *BitCalculator) AddBit(pos int) {
	bc.value |= 1 << pos
	bc.length++
}
func (bc *BitCalculator) Undo(last int) {
	bc.value &^= 1 << last
	bc.length--
}
func (bc *BitCalculator) Value() int { return bc.value }

func main() {
	br := bufio.NewReader(os.Stdin)

	var T int
	fmt.Fscan(br, &T)

	for i := 0; i < T; i++ {
		var N int
		fmt.Fscan(br, &N)
		var S string
		fmt.Fscan(br, &S)
		if Resolve(N, S) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
