package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	ss := strings.Split(S, "-")

	i, _ := strconv.Atoi(ss[0])
	j, _ := strconv.Atoi(ss[1])

	if j < 8 {
		fmt.Printf("%d-%d\n", i, j+1)
	} else {
		fmt.Printf("%d-%d\n", i+1, 1)
	}

}
