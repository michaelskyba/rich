package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
	"strconv"
	"sort"
	"strings"
)

type Streak struct {
	Name string
	Length int
}

func List(home_dir string) {
	// Get habit names
	habit_files, _ := ioutil.ReadDir(home_dir)
	var habit_filenames []string
	for _, habit_file := range habit_files {
		habit_filenames = append(habit_filenames, habit_file.Name())
	}

	// Get habit streaks
	var habits []Streak
	for _, habit_filename := range habit_filenames {
		habit_file, _ := os.Open(fmt.Sprintf("%v/%v", home_dir, habit_filename))

		// Iterate over lines in habit file, find third line
		line := 0
		scanner := bufio.NewScanner(habit_file)
		for scanner.Scan() {
			if line == 2 {
				streak, _ := strconv.Atoi(scanner.Text())
				habits = append(habits, Streak{habit_filename, streak})
			}
			line++
		}

		habit_file.Close()
	}

	// Sort habits based on streak lengths
	sort.Slice(habits, func(i, j int) bool {
		return habits[i].Length > habits[j].Length
	})

	// Find highest total streak
	var max int
	for _, habit := range habits {
		if habit.Length > max {
			max = habit.Length
		}
	}

	// Get number of digits
	var max_digits int
	if max == 0 {
		max_digits = 1
	} else {
		max_digits = 0
		for max > 0 {
			max /= 10
			max_digits++
		}
	}

	// List info
	for _, habit := range habits {
		// Get number of digits of current streak (not max)
		streak := habit.Length
		var digits int
		if streak == 0 {
			digits = 1
		} else {
			digits = 0
			for streak > 0 {
				streak /= 10
				digits++
			}
		}

		// Use digit infomration to decide on trailing spaces
		trailing := strings.Repeat(" ", max_digits - digits)

		fmt.Printf("%v%v - %v\n", trailing, habit.Length, habit.Name)
	}
}

func main() {
	// Decide where habits will be read/stored
	var home_dir string
	if os.Getenv("RICH_HOME") == "" {
		home_dir = fmt.Sprintf("%v/.local/share/rich", os.Getenv("HOME"))
	} else {
		home_dir = os.Getenv("RICH_HOME")
	}

	// list habits and streaks
	if len(os.Args) < 3 {
		List(home_dir)
	}
}
