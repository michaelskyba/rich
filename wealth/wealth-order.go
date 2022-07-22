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

func main() {
	filename1 := os.Args[1]
	filename2 := os.Args[2]

	fmt.Println(filename1, filename2)

	ordered := getLines(filename2)
	orderedHash := createHashTable(ordered)

	fmt.Println(ordered)
	fmt.Println(orderedHash)
}
