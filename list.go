package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

type Streak struct {
	Name   string
	Length int
}

func list(homeDir string) {
	// Get habit names
	habitFiles, err := ioutil.ReadDir(homeDir)
	hdl(err, "Error: Invalid home directory")

	var habitFilenames []string
	for _, habitFile := range habitFiles {
		habitFilenames = append(habitFilenames, habitFile.Name())
	}

	// Reset lost streaks
	for _, habitFilename := range habitFilenames {
		fullPath := fmt.Sprintf("%v/%v", homeDir, habitFilename)
		updateStreak(fullPath)
	}

	// Get habit streaks
	var habits []Streak
	for _, habitFilename := range habitFilenames {
		fullPath := fmt.Sprintf("%v/%v", homeDir, habitFilename)

		streak, err := strconv.Atoi(getLine(fullPath, 1))
		hdl(err, "Error: Invalid streak in habit file")

		habits = append(habits, Streak{habitFilename, streak})
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
	maxDigits := getDigits(max)

	// List info
	for _, habit := range habits {
		fmt.Printf("%*v - %v\n", maxDigits, habit.Length, habit.Name)
	}
}
