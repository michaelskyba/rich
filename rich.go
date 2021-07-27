package main

import (
	"fmt"
	"os"
)

func main() {
	var home_dir string
	if os.Getenv("RICH_HOME") == "" {
		home_dir = fmt.Sprintf("%v/.local/share/rich", os.Getenv("HOME"))
	} else {
		home_dir = os.Getenv("RICH_HOME")
	}

	fmt.Println(home_dir)
}
