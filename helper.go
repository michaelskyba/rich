package main

import (
	"bufio"
	"fmt"
	"os"
)

func hdl(err error, message string) {
	if err != nil {
		printError(message)
	}
}

func printError(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func userError() {
	printError(`Commands:
rich new <habit name> [streak]
rich delete <habit name>
rich mark <habit name> ...
rich todo
rich [list]
See the README for more information.`)
}

func getDigits(num int) int {
	var digits int

	if num == 0 {
		digits = 1
	} else {
		digits = 0
		for num > 0 {
			num /= 10
			digits++
		}
	}

	return digits
}

func getLine(filename string, i int) string {
	file, err := os.Open(filename)
	defer file.Close()

	hdl(err, "Error: Couldn't read habit file")

	line := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line == i {
			return scanner.Text()
		}
		line++
	}

	return ""
}
