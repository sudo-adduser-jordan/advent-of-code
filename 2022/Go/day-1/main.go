package main

import (
	"Go/pkg/cast"
	"Go/pkg/styles"
	"Go/pkg/utils"
	"fmt"
	"sort"
	"strconv"
)

var file = "input.txt"
var total int

func main() {
	PartOne()
	PartTwo()
}

func PartOne() int {
	calories := make(map[int]int)
	counter := 1

	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			calories[counter] = total
			total = 0
			counter += 1
		} else {
			total += number
		}
	}

	var v int
	for _, value := range calories {
		if value > v {
			v = value
		}
	}

	fmt.Print(styles.PurpleLabel("The elf with the greatest amount of calories has a total of: "))
	fmt.Println(styles.GreenText(cast.ToString(v)))
	return v
}

func PartTwo() int {
	calories := make(map[int]int)
	counter := 1

	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			calories[counter] = total
			total = 0
			counter += 1
		} else {
			total += number
		}
	}

	keys := make([]int, 0, len(calories))
	for k := range calories {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return calories[keys[i]] < calories[keys[j]]
	})

	var amount int
	for i := len(keys) - 3; i < len(keys); i++ {
		amount += calories[keys[i]]
	}
	fmt.Print(styles.PurpleLabel("The amount of calories that the top three elves are carrying is: "))
	fmt.Println(styles.GreenText(cast.ToString(amount)))
	return amount
}
