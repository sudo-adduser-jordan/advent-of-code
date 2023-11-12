package main

import (
	"Go/pkg/utils"
	"bufio"
	"fmt"
	"strings"
)

// const file = "input.txt"
const file = "test.txt"
const start = "S"
const end = "E"

func main() {
	PartOne()
	// PartTwo()
}

func PartOne() {
	scanner, file := utils.ScanFile(file)
	defer file.Close()

	grid := ScanGrid(scanner)
	steps := ShortestPath(&grid)

	for _, line := range grid {
		fmt.Println(line)
	}

	fmt.Printf("The fewest steps required to move from the current position to the location that should get the best signal is: %v", steps)
}

func PartTwo() {

}

func ScanGrid(scanner *bufio.Scanner) [][]string {
	var matrix [][]string
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}
	return matrix
}

func ShortestPath(grid *[][]string) int {

	return 0
}

func ShortestPathReversed(grid *[][]string) int {

	return 0
}
