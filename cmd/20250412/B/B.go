package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	login := false
	count := 0
	for i := 0; i < N; i++ {
		var S string
		fmt.Scan(&S)

		switch S {
		case "login":
			login = true
		case "logout":
			login = false
		case "private":
			if !login {
				count++
			}
		}
	}

	fmt.Println(count)
}
