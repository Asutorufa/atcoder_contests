package main

import "fmt"

func main() {
	var S string
	fmt.Scanln(&S)

	bytes := make([]byte, 0, len(S))
	for _, v := range S {
		if v == '2' {
			bytes = append(bytes, '2')
		}
	}

	fmt.Println(string(bytes))
}
