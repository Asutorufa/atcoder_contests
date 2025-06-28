package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Log(solve(getArrays("1 3 2 5")))
	t.Log(solve(getArrays("1 1 1 1 3 3 2 5 5 5 5 5")))
	t.Log(solve(getArrays("1 100")))
	t.Log(solve(getArrays("298077099 766294630 440423914 59187620 725560241 585990757 965580536 623321126 550925214 917827435")))

	t.Log(solve(getArrays("1 1 1 1 1")))
	t.Log(solve(getArrays("1 1 1 1 2 3")))
	t.Log(solve(getArrays("1 2 3 4 5 200")))
	t.Log(solve(getArrays("2 1 1 1 1 1 1")))
	t.Log(solve(getArrays("100 1 1 1 1 1 1")))

	s := []string{"z"}
	fmt.Println(s[0:0])
}

func getArrays(s string) []int {
	fields := strings.Split(s, " ")

	arrays := []int{}
	for _, v := range fields {
		i, _ := strconv.Atoi(v)
		arrays = append(arrays, i)
	}

	return arrays
}
