package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
	"strconv"
	"sort"
)

type Streak struct {
	Name string
	Length int
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
		var digits int
		if max == 0 {
			digits = 1
		} else {
			digits = 0
			for max > 0 {
				max /= 10
				digits++
			}
		}

		fmt.Println(habits)
		fmt.Println(digits)
	}
}
