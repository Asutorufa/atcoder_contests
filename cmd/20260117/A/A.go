package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	p := readInt()
	q := readInt()
	x := readInt()
	y := readInt()

	if x >= p && x < p+100 && y >= q && y < q+100 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

var br = bufio.NewReader(os.Stdin)

func readInt() int {
	var x int
	fmt.Fscan(br, &x)
	return x
}

func readString() string {
	var x string
	fmt.Fscan(br, &x)
	return x
}
