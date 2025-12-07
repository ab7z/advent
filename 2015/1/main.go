package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var data string

func main() {
	data = strings.TrimSpace(data)

	partOne()
	partTwo()
}

func partOne() {
	floor := 0
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case '(':
			floor++
		case ')':
			floor--
		default:
			fmt.Println("not valid char")
			return
		}
	}
	fmt.Println("one:", floor)
}

func partTwo() {
	floor := 0
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case '(':
			floor++
		case ')':
			floor--
		default:
			fmt.Println("not valid char")
			return
		}
		if floor < 0 {
			fmt.Println("two:", i+1)
			return
		}
	}
	// driven by the instruction: "The first character in the instructions has position 1”
	fmt.Println("two:", 1)
}
