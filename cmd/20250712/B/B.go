package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	str := strings.Builder{}

	for i := 0; i < N; i++ {
		var c string
		var l int

		fmt.Scan(&c, &l)

		if l > 100 {
			fmt.Println("Too Long")
			return
		}

		for j := 0; j < l; j++ {
			str.WriteString(c)
		}

		if str.Len() > 100 {
			fmt.Println("Too Long")
			return
		}
	}

	fmt.Println(str.String())
}
