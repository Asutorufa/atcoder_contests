package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	count := 0

	exceptI := true
	for i := 0; i < len(S); {
		if (exceptI && S[i] == 'i') || (!exceptI && S[i] == 'o') {
			i++
		} else {
			count++
		}

		exceptI = !exceptI
	}

	if !exceptI {
		count++
	}

	fmt.Println(count)
}
