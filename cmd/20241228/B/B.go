package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(len(strings.ReplaceAll(s, "00", "0")))
}
