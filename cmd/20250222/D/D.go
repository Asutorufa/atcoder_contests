package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scanln(&S)

	current := []byte(S)
	for len(current) > 0 {
		// fmt.Println("current", string(current))

		i1 := strings.Index(string(current), "()")
		if i1 != -1 {
			current = replaceAll(current, i1)
		}

		i2 := strings.Index(string(current), "[]")
		if i2 != -1 {
			current = replaceAll(current, i2)
		}

		i3 := strings.Index(string(current), "<>")
		if i3 != -1 {
			current = replaceAll(current, i3)
		}

		if i1 == -1 && i2 == -1 && i3 == -1 {
			break
		}
	}

	if len(current) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

var maps = map[byte]byte{
	'(': ')',
	'[': ']',
	'<': '>',
}

func replaceAll(b []byte, index int) []byte {
	start := index
	end := index + 1

	// fmt.Println("current", string(b[start:end+1]))

	var last int
	for i := 1; ; i++ {
		if end+i >= len(b) || start-i < 0 {
			break
		}

		// fmt.Println(string(b[start-i]), string(b[end+i]), maps[b[start-i]] == b[end+i])

		if maps[b[start-i]] == b[end+i] {
			last = i
		} else {
			break
		}
	}

	// fmt.Println(last, string(b[start-last:end+last+1]))
	return append(b[0:start-last], b[end+last+1:]...)
}
