package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ab7z/advent/y15"
)

func main() {
	path := "y15/input_d01.txt"
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Could not read the input file:", err)
	}

	fmt.Println("r", y15.CalcFloor(bytes))
	fmt.Println("r", y15.CalcFloor2(bytes))
}
