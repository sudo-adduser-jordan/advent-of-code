package main

import (
	"Go/pkg/cast"
	"Go/pkg/styles"
	"Go/pkg/utils"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var file = "input.txt"
var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var priority []int

func main() {
	PartOne()
	PartTwo()
}

func PartOne() int {
	scanner, file := utils.ScanFile(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		char := matchCharacterPartOne(line)
		x := numberValueOf(char)
		priority = append(priority, x)

	}
	sum := sumArray(priority)

	fmt.Print(styles.PurpleLabel("The sum of the matching items: "))
	fmt.Println(styles.GreenText(cast.ToString(sum)))
	return sum
}

func PartTwo() int {
	file, error := os.Open(file)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	defer file.Close()

	len, error := LineCounter(file)
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	x := (len + 1) / 3
	for i := 0; i < x; i++ {

		scanner.Scan()
		lineOne := scanner.Text()
		scanner.Scan()
		lineTwo := scanner.Text()
		scanner.Scan()
		lineThree := scanner.Text()

		c := matchCharacterPartTwo(lineOne, lineTwo, lineThree)

		n := numberValueOf(c)
		priority = append(priority, n)
	}

	sum := sumArray(priority)

	fmt.Print(styles.PurpleLabel("The sum of the matching items: "))
	fmt.Println(styles.GreenText(cast.ToString(sum)))
	return sum
}

func sumArray(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func numberValueOf(c rune) int {
	number := 0
	for index, value := range alphabet {
		if c == value {
			number = index + 1
			break
		}
	}
	return number
}

func matchCharacterPartOne(line string) rune {
	var char rune
	midpoint := len(line) / 2

outerloop:
	for i := 0; i < midpoint; i++ {
		charOne := rune(line[i])

		for j := midpoint; j < len(line); j++ {
			charTwo := rune(line[j])

			if charOne == charTwo {
				char = charOne
				break outerloop
			}
		}
	}
	return char
}

func matchCharacterPartTwo(lineOne string, lineTwo string, lineThree string) rune {
	var c rune
	for i := 0; i < len(lineOne); i++ {
		c = rune(lineOne[i])
		if strings.ContainsRune(lineTwo, c) && strings.ContainsRune(lineThree, c) {
			return c
		}
	}
	return c
}

func LineCounter(r io.Reader) (int, error) {

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}
