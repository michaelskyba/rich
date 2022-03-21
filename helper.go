package main

import (
	"fmt"
	"os"
	"bufio"
)

func catch_error(err error, message string) {
	if err != nil {
		fmt.Println(message)
		os.Exit(1)
	}
}

func user_error() {
	fmt.Println(`Commands:
rich new <habit name> [streak]
rich delete <habit name>
rich mark <habit name> ...
rich todo
rich [list]
See the README for more information.`)

	os.Exit(1)
}

func get_digits(num int) int {
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

func get_line(filename string, i int) string {
	file, err := os.Open(filename)
	defer file.Close()

	catch_error(err, "Error: Couldn't read habit file")

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
