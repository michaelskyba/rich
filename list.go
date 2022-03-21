package main

import (
	"io/ioutil"
	"fmt"
	"strconv"
	"sort"
)

type Streak struct {
	Name string
	Length int
}

func list(home_dir string) {
	// Get habit names
	habit_files, err := ioutil.ReadDir(home_dir)
	catch_error(err, "Error: Invalid home directory")

	var habit_filenames []string
	for _, habit_file := range habit_files {
		habit_filenames = append(habit_filenames, habit_file.Name())
	}

	// Reset lost streaks
	for _, habit_filename := range habit_filenames {
		full_path := fmt.Sprintf("%v/%v", home_dir, habit_filename)
		update_streak(full_path)
	}

	// Get habit streaks
	var habits []Streak
	for _, habit_filename := range habit_filenames {
		full_path := fmt.Sprintf("%v/%v", home_dir, habit_filename)

		streak, err := strconv.Atoi(get_line(full_path, 1))
		catch_error(err, "Error: Invalid streak in habit file")

		habits = append(habits, Streak{habit_filename, streak})
	}

	// Sort habits based on streak lengths
	sort.Slice(habits, func(i, j int) bool {
		return habits[i].Length > habits[j].Length
	})

	// Find number of digits for padding with spaces
	var max int
	for _, habit := range habits {
		if habit.Length > max {
			max = habit.Length
		}
	}
	max_digits := get_digits(max)

	// List info
	for _, habit := range habits {
		fmt.Printf("%*v - %v\n", max_digits, habit.Length, habit.Name)
	}
}
