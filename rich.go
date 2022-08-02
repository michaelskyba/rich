package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

// Reset streak if past due date
func updateStreak(filename string) {

	// Get time from habit file
	habitTime, err := time.Parse("2006-01-02 MST", getLine(filename, 0))
	hdl(err, "Error: Invalid date in habit file")

	// Has due date passed? (is current time > habitTime + 2 days?)
	currentTime := time.Now()
	if currentTime.After(habitTime.AddDate(0, 0, 2)) {

		habitFile, err := ioutil.ReadFile(filename)
		hdl(err, "Error: Couldn't open habit file")

		lines := strings.Split(string(habitFile), "\n")

		// RICH_HOOK: filename last_completion old_streak cdate
		// Don't run it if the streak is already 0, that would be useless
		hook := os.Getenv("RICH_HOOK")
		if hook != "" && lines[1] != "0" {
			habitTime := habitTime.Format("2006-01-02")
			currentTime := currentTime.Format("2006-01-02")

			exec.Command(hook, filename, habitTime, lines[1], currentTime).Output()
		}

		// Reset streak
		lines[1] = "0"
		err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
		hdl(err, "Error: Couldn't write to habit file")
	}
}

// Has a habit been marked today?
func isMarked(filename string) bool {
	cdate := time.Now().Format("2006-01-02 MST")
	habitDate := getLine(filename, 0)

	return cdate == habitDate
}

func main() {
	// Decide where habits will be read/stored
	var homeDir string
	if os.Getenv("RICH_HOME") == "" {
		homeDir = fmt.Sprintf("%v/.local/share/rich", os.Getenv("HOME"))
	} else {
		homeDir = os.Getenv("RICH_HOME")
	}

	// list habits and streaks
	if len(os.Args) == 1 || os.Args[1] == "list" {
		list(homeDir)
		os.Exit(0)

	} else if os.Args[1] == "todo" {

		// Get habit names
		habitFiles, err := ioutil.ReadDir(homeDir)
		hdl(err, "Error: Invalid home directory")

		var habitFilenames []string
		for _, habitFile := range habitFiles {
			habitFilenames = append(habitFilenames, habitFile.Name())
		}

		// Iterate over habit names, printing if not marked
		for _, habitFilename := range habitFilenames {
			fullPath := fmt.Sprintf("%v/%v", homeDir, habitFilename)

			if !isMarked(fullPath) {
				fmt.Println(habitFilename)
			}
		}

		os.Exit(0)

	} else if len(os.Args) == 2 {
		// e.g. "rich foo"
		userError()
	}

	fullPath := fmt.Sprintf("%v/%v", homeDir, os.Args[2])
	switch os.Args[1] {

	case "new":
		var err error

		// Default streak
		streak := 0
		if len(os.Args) > 3 {
			streak, err = strconv.Atoi(os.Args[3])
			hdl(err, "Error: Invalid streak")

			if streak < 0 {
				printError("Error: Invalid streak")
			}
		}

		// We need yesterday - otherwise, if you set a streak, rich mark will reset it
		// It will think the streak was in the past

		timeString := time.Now().AddDate(0, 0, -1).Format("2006-01-02 MST")
		content := []byte(fmt.Sprintf("%v\n%v\n", timeString, streak))
		err = ioutil.WriteFile(fullPath, content, 0644)

		hdl(err, "Error: Couldn't create habit file")

	case "delete":
		err := os.Remove(fullPath)
		hdl(err, "Error: Couldn't delete habit file")

	case "mark":
		// Iterate over every habit listed to mark
		for i, habit := range os.Args {

			// The 0th and 1st arguments are not habits
			if i < 2 {
				continue
			}

			// We can't use "fullPath" because that's hardcoded to the first habit
			markPath := fmt.Sprintf("%v/%v", homeDir, habit)

			if isMarked(markPath) {
				fmt.Printf("'%v' has already been completed today.\n", habit)

			} else {
				updateStreak(markPath)

				// Open habit file
				habitFile, err := ioutil.ReadFile(markPath)
				hdl(err, "Error: Couldn't read habit file")

				lines := strings.Split(string(habitFile), "\n")

				// Update date
				lines[0] = time.Now().Format("2006-01-02 MST")

				// Increment streak
				streak, err := strconv.Atoi(lines[1])
				hdl(err, "Error: Invalid streak in habit file")

				lines[1] = strconv.Itoa(streak + 1)

				err = ioutil.WriteFile(markPath, []byte(strings.Join(lines, "\n")), 0644)
				hdl(err, "Error: Couldn't write to habit file")
			}
		}

	default:
		// e.g. "rich foo bar"
		userError()
	}
}
