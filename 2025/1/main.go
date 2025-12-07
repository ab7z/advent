package main

import (
	"bufio"
	"embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var data embed.FS

const totalClicks = 100

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
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	var rotation string
	var direction uint8
	currentPosition := 50
	var clicks, password int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rotation = strings.TrimSpace(scanner.Text())
		direction = rotation[0]
		if clicks, err = strconv.Atoi(rotation[1:]); err != nil {
			fmt.Println(err.Error())
			return
		}

		switch direction {
		case 'L':
			currentPosition = (currentPosition - clicks) % totalClicks
		case 'R':
			currentPosition = (currentPosition + clicks) % totalClicks
		default:
			fmt.Println("not a valid direction")
			return
		}

		if currentPosition == 0 {
			password++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("part one: %v\n", password)
}

func partTwo() {
	file, err := data.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	var rotation string
	var direction uint8
	var clicks, password, totalRotation int

	currentPosition := 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rotation = strings.TrimSpace(scanner.Text())
		direction = rotation[0]
		if clicks, err = strconv.Atoi(rotation[1:]); err != nil {
			fmt.Println(err.Error())
			return
		}

		switch direction {
		case 'L':
			totalRotation = (((totalClicks - currentPosition) % totalClicks) + clicks) / 100
			currentPosition = (currentPosition - clicks) % totalClicks
			if currentPosition < 0 {
				currentPosition = totalClicks + currentPosition
			}
		case 'R':
			totalRotation = (currentPosition + clicks) / totalClicks
			currentPosition = (currentPosition + clicks) % totalClicks
		default:
			fmt.Println("not a valid direction")
			return
		}

		if currentPosition == 0 {
			password++
			if totalRotation > 1 {
				password += totalRotation - 1
			}
		} else if totalRotation > 0 {
			password += totalRotation
		}

		totalRotation = 0
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("part two: %v\n", math.Abs(float64(password)))
}
