package main

import (
	"Go/pkg/cast"
	"Go/pkg/styles"
	"Go/pkg/utils"
	"fmt"
)

var file = "input.txt"

func main() {
	PartOne()
	PartTwo()
}

func PartOne() int {
	scanner, file := utils.ScanFile(file)
	defer file.Close()

	var score int
	for scanner.Scan() {
		line := scanner.Text()

		switch line {
		case "A X":
			score += 4
		case "A Y":
			score += 8
		case "A Z":
			score += 3

		case "B X":
			score += 1
		case "B Y":
			score += 5
		case "B Z":
			score += 9

		case "C X":
			score += 7
		case "C Y":
			score += 2
		case "C Z":
			score += 6
		default:
			fmt.Println("error")
		}

	}

	fmt.Print(styles.PurpleLabel("Score: "))
	fmt.Println(styles.GreenText(cast.ToString(score)))
	return score
}

func PartTwo() int {
	scanner, file := utils.ScanFile(file)
	defer file.Close()

	var score int
	for scanner.Scan() {
		line := scanner.Text()

		switch line {
		case "A X":
			score += 3
		case "A Y":
			score += 4
		case "A Z":
			score += 8

		case "B X":
			score += 1
		case "B Y":
			score += 5
		case "B Z":
			score += 9

		case "C X":
			score += 2
		case "C Y":
			score += 6
		case "C Z":
			score += 7
		default:
			fmt.Println("error")
		}

	}

	fmt.Print(styles.PurpleLabel("Score: "))
	fmt.Println(styles.GreenText(cast.ToString(score)))
	return score
}
