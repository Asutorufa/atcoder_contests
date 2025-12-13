package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)

	var str strings.Builder
	for range N - len(S) {
		str.WriteByte('o')
	}
	str.WriteString(S)
	fmt.Println(str.String())
}
