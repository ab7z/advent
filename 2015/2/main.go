package main

import (
	"bufio"
	"embed"
	_ "embed"
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
	file, err := data.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	sizes := make([]string, 3)
	var l, w, h, s1, s2, s3, sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sizes = strings.Split(scanner.Text(), "x")
		l, _ = strconv.Atoi(sizes[0])
		w, _ = strconv.Atoi(sizes[1])
		h, _ = strconv.Atoi(sizes[2])
		s1 = l * w
		s2 = w * h
		s3 = h * l
		sum += 2*(s1+s2+s3) + min(s1, s2, s3)
	}
	fmt.Println("one:", sum)
}

func partTwo() {
	file, err := data.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	sizes := make([]string, 3)
	var l, w, h, sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sizes = strings.Split(scanner.Text(), "x")
		l, _ = strconv.Atoi(sizes[0])
		w, _ = strconv.Atoi(sizes[1])
		h, _ = strconv.Atoi(sizes[2])
		sum += (l * w * h) + 2*min(h+l, w+h, l+w)
	}
	fmt.Println("two:", sum)
}
