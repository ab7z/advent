package main

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var data embed.FS

func main() {
	partOne()
	partTwo()
}

func partOne() {
	content, err := data.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ranges := strings.Split(string(content), ",")
	invalids := 0

	for _, r := range ranges {
		ids := strings.Split(r, "-")
		firstID, _ := strconv.Atoi(ids[0])
		lastID, _ := strconv.Atoi(ids[1])

		for i := firstID; i <= lastID; i++ {
			id := strconv.Itoa(i)
			if len(id)%2 == 0 && id[:(len(id)/2)] == id[(len(id)/2):] {
				invalids += i
			}
		}
	}

	fmt.Println(invalids)
}

func partTwo() {
	content, err := data.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ranges := strings.Split(string(content), ",")
	invalids := 0

	for _, r := range ranges {
		ids := strings.Split(r, "-")
		firstID, _ := strconv.Atoi(ids[0])
		lastID, _ := strconv.Atoi(ids[1])

		for i := firstID; i <= lastID; i++ {
			if i <= 10 {
				continue
			}

			id := strconv.Itoa(i)

			if strings.Count(id, string(id[0])) == len(id) {
				invalids += i
				continue
			}

			for j := 2; j <= len(id)/2; j++ {
				c := strings.Count(id, id[:j])
				if j*c == len(id) {
					invalids += i
					break
				}
			}
		}
	}

	fmt.Println(invalids)
}
