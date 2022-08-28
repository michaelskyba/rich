package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Streak struct {
	Name   string
	Length int
}

// List habits and streaks
func list(homeDir string) {
	habitFiles, err := ioutil.ReadDir(homeDir)
	hdl(err, "Error: Invalid home directory")

	// Get habit names
	var habitFilenames []string
	for _, habitFile := range habitFiles {
		habitFilenames = append(habitFilenames, habitFile.Name())
	}

	// Reset lost streaks
	for _, habitFilename := range habitFilenames {
		habitPath := fmt.Sprintf("%v/%v", homeDir, habitFilename)
		checkMissed(habitPath)
	}

	// Get habit streaks
	var habits []Streak
	for _, habitFilename := range habitFilenames {
		habitPath := fmt.Sprintf("%v/%v", homeDir, habitFilename)
		habits = append(habits, Streak{habitFilename, getStreak(habitPath)})
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

	for _, habit := range habits {
		fmt.Printf("%*v - %v\n", maxDigits, habit.Length, habit.Name)
	}
}

func todoAll(homeDir string) {
	habitFiles, err := ioutil.ReadDir(homeDir)
	hdl(err, "Error: Invalid home directory")

	// Get habit names
	var habitFilenames []string
	for _, habitFile := range habitFiles {
		habitFilenames = append(habitFilenames, habitFile.Name())
	}

	// Iterate over habit names, printing if not marked
	for _, habitFilename := range habitFilenames {
		habitPath := fmt.Sprintf("%v/%v", homeDir, habitFilename)

		if !isMarked(habitPath) {
			fmt.Println(habitFilename)
		}
	}
}

func createHabit(habitPath string, argsLen int) {
	var err error

	// Default streak
	streak := 0
	if argsLen == 4 {
		streak, err = strconv.Atoi(os.Args[3])
		hdl(err, "Error: Invalid streak")

		if streak < 0 {
			printError("Error: Invalid streak")
		}
	}

	err = setStreak(habitPath, streak)
	hdl(err, "Error: Couldn't create habit file")
}

func markHabit(homeDir string) {
	for i, habit := range os.Args {
		// The 0th and 1st arguments are not habits
		if i < 2 {
			continue
		}

		habitPath := fmt.Sprintf("%v/%v", homeDir, habit)

		if isMarked(habitPath) {
			fmt.Printf("'%v' has already been completed today.\n", habit)

		} else {
			checkMissed(habitPath)

			habitFile, err := ioutil.ReadFile(habitPath)
			hdl(err, "Error: Couldn't read habit file")

			lines := strings.Split(string(habitFile), "\n")

			// Update date
			lines[0] = time.Now().Format("2006-01-02 MST")

			// Increment streak
			streak, err := strconv.Atoi(lines[1])
			hdl(err, "Error: Invalid streak in habit file")
			lines[1] = strconv.Itoa(streak + 1)

			writeLines(habitPath, lines)
		}
	}
}

func todoSingle(habitPath string) {
	if isMarked(habitPath) {
		fmt.Println("done")
		os.Exit(1)
	} else {
		fmt.Println("todo")
		os.Exit(0)
	}
}

// Manually set the streak of an individual habit
func setStreakCommand(habitPath string, streak string) {
	_, err := strconv.Atoi(streak)
	hdl(err, "Error: Invalid streak provided")

	err = setStreak(habitPath, streak)
	hdl(err, "Error: Couldn't write to habit file")
}
