package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bufio"
	"strconv"
	"sort"
	"strings"
	"time"
)

type Streak struct {
	Name string
	Length int
}

func get_digits(num int) int {
	var digits int

	if num == 0 {
		digits = 1
	} else {
		digits = 0
		for num > 0 {
			num  /= 10
			digits++
		}
	}

	return digits
}

func get_line(filename string, i int) string {
	file, _ := os.Open(filename)
	defer file.Close()

	line := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if line == i {
			return scanner.Text()
		}
		line++
	}

	return ""
}

func list(home_dir string) {
	// Get habit names
	habit_files, _ := ioutil.ReadDir(home_dir)
	var habit_filenames []string
	for _, habit_file := range habit_files {
		habit_filenames = append(habit_filenames, habit_file.Name())
	}

	// Get habit streaks
	var habits []Streak
	for _, habit_filename := range habit_filenames {
		full_path := fmt.Sprintf("%v/%v", home_dir, habit_filename)

		streak, _ := strconv.Atoi(get_line(full_path, 2))
		habits = append(habits, Streak{habit_filename, streak})
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
	max_digits := get_digits(max)

	// List info
	for _, habit := range habits {
		// Use digit information to decide on trailing spaces
		trailing := strings.Repeat(" ", max_digits - get_digits(habit.Length))

		fmt.Printf("%v%v - %v\n", trailing, habit.Length, habit.Name)
	}
}

// Reset streak if past due date
func update_streak(filename string) {

	// Get time from habit file
	habit_time, _ := time.Parse("2006-01-02 MST", get_line(filename, 0))

	// Has due date passed? (is current time > habit_time + 2 days?)
	if time.Now().After(habit_time.AddDate(0, 0, 2)) {

		habit_file, _ := ioutil.ReadFile(filename)
		lines := strings.Split(string(habit_file), "\n")

		// Reset streak
		lines[1] = "0"
		_ = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
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
	if len(os.Args) == 1 {
		list(home_dir)
		os.Exit(0)

	} else if os.Args[1] == "todo" {

		// Get habit names
		habit_files, _ := ioutil.ReadDir(home_dir)
		var habit_filenames []string
		for _, habit_file := range habit_files {
			habit_filenames = append(habit_filenames, habit_file.Name())
		}

		// Iterate over habit names, comparing dates
		for _, habit_filename := range habit_filenames {
			full_path := fmt.Sprintf("%v/%v", home_dir, habit_filename)
			fmt.Println(get_line(full_path, 0))
		}

		os.Exit(0)
	}

	switch os.Args[1] {

	case "new":
		// Defaults
		interval := 1
		if len(os.Args) > 3 {
			interval, _ = strconv.Atoi(os.Args[3])
		}
		streak := 0
		if len(os.Args) > 4 {
			streak, _ = strconv.Atoi(os.Args[4])
		}

		filename := fmt.Sprintf("%v/%v", home_dir, os.Args[2])

		// April Fool's Day is arbitrary, it just needs a day in the past
		content := []byte(fmt.Sprintf("2021-04-01\n%v\n%v", interval, streak))
		_ = ioutil.WriteFile(filename, content, 0644)

		os.Exit(0)

	case "mark":
		update_streak(fmt.Sprintf("%v/%v", home_dir, os.Args[2]))
		os.Exit(0)

	default:
		fmt.Println("See the README for usage.")
		os.Exit(1)
	}
}
