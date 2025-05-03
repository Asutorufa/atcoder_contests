package main

import "fmt"

func main() {
	var S string
	fmt.Scanln(&S)

	for i := 'a'; i <= 'z'; i++ {
		if !contains(S, byte(i)) {
			fmt.Println(string(i))
			return
		}
	}
}

func contains(s string, c byte) bool {
	for i := range s {
		if s[i] == c {
			return true
		}
	}

	return false
}
