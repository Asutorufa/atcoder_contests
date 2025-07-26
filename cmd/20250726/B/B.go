package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	bytes := []byte(S)

	if strings.IndexByte(S, '#') == -1 {
		bytes[0] = 'o'
		fmt.Println(string(bytes))
		return
	}

	for i, v := range bytes {
		if v != '#' {
			continue
		}

		if i != 0 && bytes[i-1] == '.' && !findBeforeO(bytes, i-1) {
			bytes[i-1] = 'o'
		}

		if i < len(bytes)-1 && bytes[i+1] == '.' && !findAfterO(bytes, i+1) {
			bytes[i+1] = 'o'
		}
	}

	fmt.Println(string(bytes))
}

func findBeforeO(s []byte, i int) bool {
	for ; i >= 0; i-- {
		if s[i] == 'o' {
			return true
		}

		if s[i] == '#' {
			return false
		}
	}

	return false
}

func findAfterO(s []byte, i int) bool {
	for ; i < len(s); i++ {
		if s[i] == 'o' {
			return true
		}

		if s[i] == '#' {
			return false
		}
	}

	return false
}

/*
.#.#.#.#.
....#...
########
........
....####
####....
.######.
*/
