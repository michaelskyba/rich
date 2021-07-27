package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	var home_dir string
	if os.Getenv("RICH_HOME") == "" {
		home_dir = fmt.Sprintf("%v/.local/share/rich", os.Getenv("HOME"))
	} else {
		home_dir = os.Getenv("RICH_HOME")
	}

	// list habits and streaks
	if len(os.Args) < 3 {
		// Get list of habits
		habit_files, _ := ioutil.ReadDir(home_dir)

		// Print each name
		for _, habit_file := range habit_files {
			fmt.Println(habit_file.Name())
		}
	}
}
