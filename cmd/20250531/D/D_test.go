package main

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	fmt.Println(solve("10011"))
	fmt.Println(solve("1111111111"))
	fmt.Println(solve("0000000"))
	fmt.Println(solve("1111111100"))
	fmt.Println(solve("000000011"))
	fmt.Println(solve("11111111001"))
	fmt.Println()
	fmt.Println(solve("01"))
	fmt.Println(solve("1000010011"))
	fmt.Println(solve("111100010011"))
	fmt.Println(solve("111"))
	fmt.Println(solve("00010101"))
	fmt.Println()
	fmt.Println(solve("1000011"))
	fmt.Println(solve("110111"))
	fmt.Println(solve("010101"))
	fmt.Println(solve("0110110"))
	fmt.Println(solve("01010101"))
}
