package main

import (
	"Go/pkg/cast"
	"Go/pkg/mathy"
	"Go/pkg/styles"
	"Go/pkg/utils"
	"bufio"
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

	matrix := ParseGrid(scanner)
	points := cast.ToString(FindPoints(matrix))

	fmt.Print(styles.PurpleLabel("The amount of trees visiable:"))
	fmt.Println(styles.GreenText(" " + points))
	return cast.ToInt(points)
}

func PartTwo() int {
	scanner, file := utils.ScanFile(file)
	defer file.Close()

	matrix := ParseGrid(scanner)
	score := cast.ToString(FindPointsPartTwo(matrix))

	fmt.Print(styles.PurpleLabel("the highest scenic score possible:"))
	fmt.Println(styles.GreenText(" " + score))
	return cast.ToInt(score)
}

func ParseGrid(scanner *bufio.Scanner) [][]int {
	var matrix [][]int
	for scanner.Scan() {
		row := cast.SliceToInt(scanner.Bytes())
		matrix = append(matrix, row)
	}
	return matrix
}

func FindPoints(matrix [][]int) int {
	count := (len(matrix) * 4) - 4

	rows := len(matrix)
	for current_row := 1; current_row < rows-1; current_row++ {

		columns := len(matrix[current_row])
		for current_column := 1; current_column < columns-1; current_column++ {

			point := matrix[current_row][current_column]

			if CheckTop(matrix, current_row, current_column, point) ||
				CheckBottom(matrix, current_row, current_column, point) ||
				CheckRight(matrix, current_row, current_column, point) ||
				CheckLeft(matrix, current_row, current_column, point) {
				count++
			}
		}
	}
	return count
}

func CheckTop(matrix [][]int, current_row int, current_column int, point int) bool {

	var max int
	for row := 0; row < current_row; row++ {
		if max < matrix[row][current_column] {
			max = matrix[row][current_column]
		}
	}

	if point > max {
		return true
	}

	return false
}

func CheckBottom(matrix [][]int, current_row int, current_column int, point int) bool {

	var max int
	for row := len(matrix) - 1; row > current_row; row-- {
		if max < matrix[row][current_column] {
			max = matrix[row][current_column]
		}
	}

	if point > max {
		return true
	}

	return false
}

func CheckLeft(matrix [][]int, current_row int, current_column int, point int) bool {

	var max int
	for column := 0; column < current_column; column++ {
		if max < matrix[current_row][column] {
			max = matrix[current_row][column]
		}
	}

	if point > max {
		return true
	}

	return false
}

func CheckRight(matrix [][]int, current_row int, current_column int, point int) bool {

	var max int
	for column := len(matrix[current_row]) - 1; column > current_column; column-- {
		if max < matrix[current_row][column] {
			max = matrix[current_row][column]
		}
	}

	if point > max {
		return true

	}
	return false
}

func FindPointsPartTwo(matrix [][]int) int {
	var score_slice []int
	var score int

	rows := len(matrix)
	for current_row := 1; current_row < rows-1; current_row++ {

		columns := len(matrix[current_row])
		for current_column := 1; current_column < columns-1; current_column++ {

			point := matrix[current_row][current_column]

			top_score := CheckTopPartTwo(matrix, current_row, current_column, point)
			bottom_score := CheckBottomPartTwo(matrix, current_row, current_column, point)
			left_score := CheckLeftPartTwo(matrix, current_row, current_column, point)
			right_score := CheckRightPartTwo(matrix, current_row, current_column, point)

			score = top_score * bottom_score * left_score * right_score
			score_slice = append(score_slice, score)
		}
	}
	return mathy.MaxIntSlice(score_slice)
}

func CheckTopPartTwo(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0
	for row := current_row - 1; row >= 0; row-- {
		count++
		if matrix[row][current_column] >= point {
			break
		}
	}
	return count
}

func CheckBottomPartTwo(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0
	for row := current_row + 1; row <= len(matrix)-1; row++ {
		count++
		if matrix[row][current_column] >= point {
			break
		}
	}
	return count
}

func CheckLeftPartTwo(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0
	for column := current_column - 1; column >= 0; column-- {
		count++
		if matrix[current_row][column] >= point {
			break
		}
	}
	return count
}

func CheckRightPartTwo(matrix [][]int, current_row int, current_column int, point int) int {

	count := 0

	for column := current_column + 1; column <= len(matrix)-1; column++ {
		count++
		if matrix[current_row][column] >= point {
			break
		}
	}
	return count
}
