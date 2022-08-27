package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var userError = `Commands:
rich new <habit name> [streak]
rich delete <habit name>
rich mark <habit name> ...
rich todo
rich [list]
See the README for more information.`

func hdl(err error, message string) {
	if err != nil {
		printError(message)
	}
}

func printError(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
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

func writeLines(habitPath string, lines []string) {
	err := ioutil.WriteFile(habitPath, []byte(strings.Join(lines, "\n")), 0644)
	hdl(err, "Error: Couldn't write to habit file")
}

func getLine(habitPath string, i int) string {
	file, err := os.Open(habitPath)
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
