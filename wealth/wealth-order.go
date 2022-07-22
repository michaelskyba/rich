package main

// Internal usage: ./wealth-order <habit list file> <order file>

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func hdl(err error) {
	if err != nil {
		panic(err)
	}
}

func getLines(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()
	hdl(err)

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

/*
With a hash table, we can check for duplicates in O(1) time instead of O(r)
time(where r is the size of the ordered input). This means that the unique
extraction process will take O(h) time instead of O(hr) time (where h is the
size of the full habits list file). Since we're already going to be using O(r)
space with the list of lines, the space would be O(2r) which is shortened to
O(r). Basically, the use of a hash table makes the program significantly faster
at high inputs without using much extra memory.
*/
func createHashTable(lines []string) map[string]bool {
	seen := map[string]bool{}

	for _, line := range lines {
		seen[line] = true
	}

	return seen
}

func getUnique(filename string, seen map[string]bool) []string {
	file, err := os.Open(filename)
	defer file.Close()
	hdl(err)

	scanner := bufio.NewScanner(file)
	unique := []string{}

	for scanner.Scan() {
		text := scanner.Text()
		if !seen[text] {
			unique = append(unique, text)
		}
	}

	return unique
}

func main() {
	filename1 := os.Args[1]
	filename2 := os.Args[2]

	ordered := getLines(filename2)
	orderedHash := createHashTable(ordered)
	unordered := getUnique(filename1, orderedHash)

	fmt.Println(strings.Join(ordered, "\n"))
	fmt.Println(strings.Join(unordered, "\n"))
}
