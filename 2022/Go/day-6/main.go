package main

import (
	"Go/pkg/cast"
	"Go/pkg/styles"
	"Go/pkg/utils"
	"fmt"
	"strings"
)

var file = "input.txt"
var signal string
var marker string

func main() {
	PartOne()
	PartTwo()
}

func PartOne() (string, int) {
	var counter = 5
	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		signal = scanner.Text()
	}

	marker = signal[0:4]
	signal = signal[4:]

	if ValidateMarkerPartOne(marker) {
		fmt.Printf("marker: %v \n", marker)
		fmt.Print(styles.PurpleLabel("Counter: "))
		fmt.Println(styles.GreenText(cast.ToString(counter)))
	}

	for _, value := range signal {
		marker = marker[1:] + string(value)

		if ValidateMarkerPartOne(marker) {
			fmt.Print(styles.PurpleLabel("Marker: "))
			fmt.Println(styles.GreenText(marker))
			fmt.Print(styles.PurpleLabel("Counter: "))
			fmt.Println(styles.GreenText(cast.ToString(counter)))
			break
		}
		counter++
	}
	return marker, counter
}

func PartTwo() int {
	var counter = 15
	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		signal = scanner.Text()
	}

	marker = signal[0:14]
	signal = signal[14:]

	if ValidateMarkerPartTwo(marker) {
		fmt.Print(styles.PurpleLabel("Counter: "))
		fmt.Println(styles.GreenText(cast.ToString(counter)))
	}

	for _, value := range signal {
		marker = marker[1:] + string(value)

		if ValidateMarkerPartTwo(marker) {
			fmt.Print(styles.PurpleLabel("Counter: "))
			fmt.Println(styles.GreenText(cast.ToString(counter)))
			break
		}

		counter++
	}
	return counter
}

func ValidateMarkerPartTwo(s string) bool {

	for i := 0; i < len(s); i++ {
		if strings.Contains(
			strings.Replace(s, string(s[i]), "", 1),
			string(s[i])) {
			return false
		}
	}
	return true
}

func ValidateMarkerPartOne(s string) bool {

	if strings.Contains(string(s[1:4]), string(s[0])) ||
		strings.Contains(string(s[0])+s[2:4], string(s[1])) ||
		strings.Contains(string(s[0:2])+s[3:4], string(s[2])) {
		return false
	}

	return true
}
