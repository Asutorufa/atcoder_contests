package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scanln(&S)

	var bytes bytes.Buffer

	for {
		i := strings.Index(S, "WA")
		if i == -1 {
			break
		}

		// fmt.Println(S)

		resp := replaceBeforeAll([]byte(S[:i+2]), i)
		bytes.Write(resp)
		S = S[i+2:]
	}

	bytes.Write([]byte(S[:]))

	fmt.Println(bytes.String())
}

func replaceBeforeAll(s []byte, index int) []byte {
	s[index] = 'A'
	s[index+1] = 'C'
	for i := index - 1; i >= 0; i-- {
		// fmt.Println(i, string(s[i]))
		if s[i] == 'W' {
			s[i] = 'A'
			s[i+1] = 'C'
		} else {
			break
		}
	}

	return s
}
