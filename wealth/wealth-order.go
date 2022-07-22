package main

// Internal usage: ./wealth-order <habit list file> <order file>

import (
	"fmt"
	"os"
	"bufio"
)

func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

func getLines(filename string) []string {
	file, err := os.Open(filename)
	hdl(err)

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	filename1 := os.Args[1]
	filename2 := os.Args[2]

	fmt.Println(filename1, filename2)

	ordered := getLines(filename2)
	fmt.Println(ordered)
}
