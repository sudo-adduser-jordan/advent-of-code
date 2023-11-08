package main

import (
	"Go/pkg/cast"
	"Go/pkg/styles"
	"Go/pkg/utils"

	"fmt"
	"strconv"
)

const file = "input.txt"

type Point struct {
	x int
	y int
}

func main() {
	// PartOne()
	PartTwo()
}

func PartOne() {

	head := Point{0, 0}
	tail := Point{0, 0}
	visited := map[Point]bool{}

	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		direction := rune(line[0])
		motions, _ := strconv.Atoi(line[2:])

		for motions > 0 {

			switch direction {
			case 'R':
				head.x++
			case 'L':
				head.x--
			case 'U':
				head.y++
			case 'D':
				head.y--
			}

			MoveTail(head, &tail)
			visited[tail] = true
			motions--
		}
	}
	fmt.Print(styles.PurpleLabel("The amount of postions the tail visited:"))
	fmt.Print(" ")
	fmt.Print(styles.GreenText(cast.ToString((len(visited)))))
}

func PartTwo() {

	visited := map[Point]bool{}
	knots := [10]Point{}

	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		direction := rune(line[0])
		motions, _ := strconv.Atoi(line[2:])

		for motions > 0 {

			switch direction {
			case 'R':
				knots[0].x++
			case 'L':
				knots[0].x--
			case 'U':
				knots[0].y++
			case 'D':
				knots[0].y--
			}

			for i := 0; i < 9; i++ {
				knots[i+1] = MoveTailPart2(knots[i], knots[i+1])
			}

			visited[knots[9]] = true
			motions--
		}
	}
	fmt.Print(styles.PurpleLabel("The amount of postions the tail visited:"))
	fmt.Print(" ")
	fmt.Print(styles.GreenText(cast.ToString((len(visited)))))
}

func MoveTail(head Point, tail *Point) {

	if head.x-tail.x > 1 {
		tail.x++
		tail.y = head.y
	}

	if head.x-tail.x < -1 {
		tail.x--
		tail.y = head.y

	}

	if head.y-tail.y > 1 {
		tail.y++
		tail.x = head.x
	}

	if head.y-tail.y < -1 {
		tail.y--
		tail.x = head.x
	}

}

func MoveTailPart2(head Point, tail Point) Point {

	point := Point{head.x - tail.x, head.y - tail.y}

	switch point {
	case
		Point{-2, 1},
		Point{-1, 2},
		Point{0, 2},
		Point{1, 2},
		Point{2, 1},
		Point{2, 2},
		Point{-2, 2}:
		tail.y++
	}

	switch point {
	case
		Point{1, 2},
		Point{2, 1},
		Point{2, 0},
		Point{2, -1},
		Point{1, -2},
		Point{2, 2},
		Point{2, -2}:
		tail.x++
	}

	switch point {
	case
		Point{1, -2},
		Point{-1, -2},
		Point{0, -2},
		Point{2, -1},
		Point{2, -2},
		Point{-2, -2},
		Point{-2, -1}:
		tail.y--
	}

	switch point {
	case
		Point{-1, -2},
		Point{-1, 2},
		Point{-2, -2},
		Point{-2, -1},
		Point{-2, 0},
		Point{-2, 1},
		Point{-2, 2}:
		tail.x--
	}

	return tail
}
