package main

import (
	"fmt"
	"os"
	"os/exec"
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

func catch_error(err error, message string) {
	if err != nil {
		fmt.Println(message)
		os.Exit(1)
	}
}

func user_error() {
	fmt.Println(`Commands:
rich new <habit name> [streak]
rich mark <habit name>
rich todo
rich
See the README for more information.`)

	os.Exit(1)
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
	habit_time, err := time.Parse("2006-01-02 MST", get_line(filename, 0))
	catch_error(err, "Error: Invalid date in habit file")

	// Has due date passed? (is current time > habit_time + 2 days?)
	current_time := time.Now()
	if current_time.After(habit_time.AddDate(0, 0, 2)) {

		habit_file, err := ioutil.ReadFile(filename)
		catch_error(err, "Error: Couldn't open habit file")

		lines := strings.Split(string(habit_file), "\n")

		// RICH_HOOK: filename last_completion old_streak cdate
		// Don't run it if the streak is already 0, that would be useless
		hook := os.Getenv("RICH_HOOK")
		if hook != "" && lines[1] != "0" {
			habit_time := habit_time.Format("2006-01-02")
			current_time := current_time.Format("2006-01-02")

			exec.Command(hook, filename, habit_time, lines[1], current_time).Output()
		}

		// Reset streak
		lines[1] = "0"
		err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
		catch_error(err, "Error: Couldn't write to habit file")
	}
}

// Has a habit been marked today?
func is_marked(filename string) bool {
	cdate := time.Now().Format("2006-01-02 MST")
	habit_date := get_line(filename, 0)

	return cdate == habit_date
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
		habit_files, err := ioutil.ReadDir(home_dir)
		catch_error(err, "Error: Invalid home directory")

		var habit_filenames []string
		for _, habit_file := range habit_files {
			habit_filenames = append(habit_filenames, habit_file.Name())
		}

		// Iterate over habit names, printing if not marked
		for _, habit_filename := range habit_filenames {
			full_path := fmt.Sprintf("%v/%v", home_dir, habit_filename)

			if !is_marked(full_path) {
				fmt.Println(habit_filename)
			}
		}

		os.Exit(0)

	} else if len(os.Args) == 2 {
		// e.g. "rich foo"
		user_error()
	}

	full_path := fmt.Sprintf("%v/%v", home_dir, os.Args[2])
	switch os.Args[1] {

	case "new":
		var err error

		// Default streak
		streak := 0
		if len(os.Args) > 3 {
			streak, err = strconv.Atoi(os.Args[3])
			catch_error(err, "Error: Invalid streak")

			if streak < 0 {
				fmt.Println("Error: Invalid streak")
				os.Exit(1)
			}
		}

		// We need yesterday - otherwise, if you set a streak, rich mark will reset it
		// It will think the streak was in the past

		time_string := time.Now().AddDate(0, 0, -1).Format("2006-01-02 MST")
		content := []byte(fmt.Sprintf("%v\n%v\n", time_string, streak))
		err = ioutil.WriteFile(full_path, content, 0644)

		catch_error(err, "Error: Couldn't create habit file")

		os.Exit(0)

	case "mark":
		// Iterate over every habit listed to mark
		for i, habit := range os.Args {

			// The 0th and 1st arguments are not habits
			if i < 2 {
				continue
			}

			// We can't use "full_path" because that's hardcoded to the first habit
			mark_path := fmt.Sprintf("%v/%v", home_dir, habit)

			if is_marked(mark_path) {
				fmt.Println("That habit has already been completed today.")
				os.Exit(1)

			} else {
				update_streak(mark_path)

				// Open habit file
				habit_file, err := ioutil.ReadFile(mark_path)
				catch_error(err, "Error: Couldn't read habit file")

				lines := strings.Split(string(habit_file), "\n")

				// Update date
				lines[0] = time.Now().Format("2006-01-02 MST")

				// Increment streak
				streak, err := strconv.Atoi(lines[1])
				catch_error(err, "Error: Invalid streak in habit file")

				lines[1] = strconv.Itoa(streak + 1)

				err = ioutil.WriteFile(mark_path, []byte(strings.Join(lines, "\n")), 0644)
				catch_error(err, "Error: Couldn't write to habit file")
			}
		}

	default:
		// e.g. "rich foo bar"
		user_error()
	}
}
