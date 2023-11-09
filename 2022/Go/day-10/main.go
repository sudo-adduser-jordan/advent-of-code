package main

import (
	"Go/pkg/styles"
	"Go/pkg/utils"
	"fmt"
	"strconv"
)

const file = "input.txt"

func main() {
	// PartOne()
	PartTwo()
}

func PartOne() {
	cycles := 1
	registir := 1
	signal := map[int]int{}

	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		code := line[:4]

		switch code {
		case "noop":
			cycles++
			signal[cycles] = cycles * registir
		case "addx":
			value, _ := strconv.Atoi(line[5:])
			cycles++
			signal[cycles] = cycles * registir
			cycles++
			registir += value
			signal[cycles] = cycles * registir
		}
	}

	sum :=
		signal[20] +
			signal[60] +
			signal[100] +
			signal[140] +
			signal[180] +
			signal[220]

	fmt.Printf("The sum of the six signal strengths is: %v", sum)
}

func PartTwo() {
	cycle := 0
	registir := 1

	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		code := line[:4]

		switch code {
		case "noop":
			PrintCRT(registir, cycle)
			cycle++
		case "addx":
			value, _ := strconv.Atoi(line[5:])
			PrintCRT(registir, cycle)
			cycle++
			PrintCRT(registir, cycle)
			cycle++
			registir += value
		}
	}
}

func PrintCRT(spritePosition int, cycle int) {
	if cycle%40 == 0 {
		println()
	}

	if spritePosition-1 == cycle%40 ||
		spritePosition == cycle%40 ||
		spritePosition+1 == cycle%40 {
		fmt.Print(styles.RedText("#"))
	} else {
		print(".")
	}
}
