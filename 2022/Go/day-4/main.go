package main

import (
	"Go/pkg/cast"
	"Go/pkg/styles"
	"Go/pkg/utils"
	"fmt"
	"strconv"
	"strings"
)

var file = "input.txt"

func main() {
	PartOne()
	PartTwo()
}

func PartOne() int {
	var total int
	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()

		array := strings.Split(line, ",")

		stringOne := array[0]
		stringTwo := array[1]

		arrayOne := strings.Split(stringOne, "-")
		arrayTwo := strings.Split(stringTwo, "-")

		rangeOneStart, _ := strconv.Atoi(arrayTwo[0])
		rangeOneEnd, _ := strconv.Atoi(arrayTwo[1])
		rangeTwoStart, _ := strconv.Atoi(arrayOne[0])
		rangeTwoEnd, _ := strconv.Atoi(arrayOne[1])

		if rangeTwoStart >= rangeOneStart && rangeTwoEnd <= rangeOneEnd || rangeOneStart >= rangeTwoStart && rangeOneEnd <= rangeTwoEnd {
			total++
		}

	}

	fmt.Print(styles.PurpleLabel("The amount of assignment pairs that one range fully contains the other is: "))
	fmt.Println(styles.GreenText(cast.ToString(total)))
	return total
}

func PartTwo() int {
	var total int
	scanner, file := utils.ScanFile(file)
	defer file.Close()
	for scanner.Scan() {
		line := scanner.Text()

		array := strings.Split(line, ",")

		stringOne := array[0]
		stringTwo := array[1]

		arrayOne := strings.Split(stringOne, "-")
		arrayTwo := strings.Split(stringTwo, "-")

		rangeOneStart, _ := strconv.Atoi(arrayTwo[0])
		rangeOneEnd, _ := strconv.Atoi(arrayTwo[1])
		rangeTwoStart, _ := strconv.Atoi(arrayOne[0])
		rangeTwoEnd, _ := strconv.Atoi(arrayOne[1])

		if rangeOneStart <= rangeTwoEnd && rangeOneEnd >= rangeTwoStart || rangeTwoStart <= rangeOneEnd && rangeTwoEnd >= rangeOneStart {
			total++
		}
	}

	fmt.Print(styles.PurpleLabel("The amount of assignment pairs that the ranges overlap: "))
	fmt.Println(styles.GreenText(cast.ToString(total)))
	return total
}
