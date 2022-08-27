package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Reset streak if past due date
func updateStreak(filename string) {
	habitTime, err := time.Parse("2006-01-02 MST", getLine(filename, 0))
	hdl(err, "Error: Invalid date in habit file")

	// Has due date passed? (i.e. Is current time > habitTime + 2 days?)
	currentTime := time.Now()
	if currentTime.After(habitTime.AddDate(0, 0, 2)) {
		habitFile, err := ioutil.ReadFile(filename)
		hdl(err, "Error: Couldn't open habit file")

		lines := strings.Split(string(habitFile), "\n")

		// RICH_HOOK: filename last_completion old_streak cdate
		// Don't run it if the streak is already 0: that would be useless
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

// Decide where habits will be read/stored
func getHomeDir() string {
	if os.Getenv("RICH_HOME") == "" {
		return fmt.Sprintf("%v/.local/share/rich", os.Getenv("HOME"))
	}
	return os.Getenv("RICH_HOME")
}
