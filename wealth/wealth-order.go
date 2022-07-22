package main

// Internal usage: ./wealth-order <habit list file> <order file>

import (
	"fmt"
	"os"
)

func main() {
	filename1 := os.Args[1]
	filename2 := os.Args[2]

	fmt.Println(filename1, filename2)
}
