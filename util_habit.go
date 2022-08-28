package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Check if a habit has been missed and execute a hook (or reset) accordingly
func checkMissed(habitPath string) {
	habitTime, err := time.Parse("2006-01-02 MST", getLine(habitPath, 0))
	hdl(err, "Error: Invalid date in habit file")

	// Has due date passed? (i.e. Is current time > habitTime + 2 days?)
	currentTime := time.Now()
	if currentTime.Before(habitTime.AddDate(0, 0, 2)) {
		return
	}

	hook := os.Getenv("RICH_HOOK")
	if hook != "" {
		habitTime := habitTime.Format("2006-01-02")
		currentTime := currentTime.Format("2006-01-02")

		// RICH_HOOK: filepath last_completion old_streak cdate
		exec.Command(hook, habitPath, habitTime, lines[1], currentTime).Output()

	} else {
		// Don't run reset if the streak is already 0: that would be useless
		if lines[1] != "0" {
			return
		}

		habitFile, err := ioutil.ReadFile(habitPath)
		hdl(err, "Error: Couldn't open habit file")
		lines := strings.Split(string(habitFile), "\n")

		// Reset streak
		lines[1] = "0"
		writeLines(habitPath, lines)
	}
}

// Has a habit been marked as completed for today?
func isMarked(habitPath string) bool {
	cdate := time.Now().Format("2006-01-02 MST")
	habitDate := getLine(habitPath, 0)

	return cdate == habitDate
}

// Decide where habits will be read/stored
func getHomeDir() string {
	if os.Getenv("RICH_HOME") == "" {
		return fmt.Sprintf("%v/.local/share/rich", os.Getenv("HOME"))
	}
	return os.Getenv("RICH_HOME")
}

func getStreak(habitPath string) int {
	streak, err := strconv.Atoi(getLine(habitPath, 1))
	hdl(err, "Error: Invalid streak in habit file")
	return streak
}
