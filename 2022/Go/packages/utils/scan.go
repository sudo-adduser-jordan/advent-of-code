package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFile(s string) (*bufio.Scanner, *os.File) {
	file, error := os.Open(s)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	return scanner, file
}
