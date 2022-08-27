package main

import (
	"fmt"
	"os"
)

func main() {
	ln := len(os.Args)
	homeDir := getHomeDir()

	var habitPath string
	if ln > 2 {
		habitPath = fmt.Sprintf("%v/%v", homeDir, os.Args[2])
	}

	var c string
	if ln > 1 {
		c = os.Args[1]
	}

	switch {
	case (c == "list" && ln == 2) || ln == 1:
		list(homeDir)

	case c == "todo" && ln == 2:
		todoAll(homeDir)

	case c == "new" && (ln == 3 || ln == 4):
		createHabit(habitPath, ln)

	case c == "delete" && ln == 3:
		err := os.Remove(habitPath)
		hdl(err, "Error: Couldn't delete habit file")

	case c == "mark" && ln > 2:
		markHabit(homeDir)

	default:
		printError(userError)
	}
}
