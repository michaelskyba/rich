# Rich
rich: a simple, streak-oriented, daily habit tracker written in Go

## Installation
```sh
git clone https://github.com/michaelskyba/rich
cd rich
go build .
su -c "cp rich /usr/local/bin/"
mkdir -p ${RICH_HOME:-$HOME/.local/share/rich}
```

## Overview
rich stores each habit in a file, where the name of the file is the same name
that you gave the habit. The layout is as follows:
```
<date of the last time the habit was performed>
<the current streak>
```
These habit files will be stored in ``$RICH_HOME``, or ``$HOME/.local/share/rich``
if ``$RICH_HOME`` is unset. **rich does not create these directories
automatically**. ``mkdir`` your directory manually before your first use, as
shown in the installation section.

## Commands
### ``rich new <habit name> [streak]``
Create a new habit. The streak will be set to "0" by default (no streak yet).
The last completion date will be set to yesterday.

### ``rich delete <habit name>``
Delete a habit. There is no confirmation prompt.

### ``rich mark <habit name> ...``
Mark one more habits as complete for today.

### ``rich set <habit name> <streak>``
Manually set the streak of a habit. The last completion date will be set to
yesterday.

### ``rich streak <habit name>``
Print the streak length of a habit.

### ``rich todo [habit name]``
If no habit name is provided, print a list of habits that have yet to be
completed today.

If a habit is provided as an additional argument, instead of listing anything,
output "todo" (with exit code 0) or "done" (with exit code 1) depending on
whether the habit has been marked today.

### ``rich [list]``
List all existing habits with padded spacing, sorted by streak length.

## $RICH_HOOK
When a day passes for which a habit is not marked as completed, the habit
becomes "missed". Specifically, this state is reached when the current date is
two days past the habit's last completion date. For example, if it's 2022-08-23
now but the habit was last completed on 2022-08-21, you have forgotten to
complete the habit for 2022-08-22. The checking (only) happens when ``rich
[list]``, ``rich streak``, or ``rich mark`` are used.

By default (i.e. no hook is set), when a day is missed, the habit's streak is
reset to zero. However, you can override (or add) to this behaviour through a
hook. To do so, set ``$RICH_HOOK`` to an executable (can be any programming
language). The following arguments will be passed in this order:
1. full path to habit file
2. last completion date in YYYY-MM-DD format
3. streak length before reset
4. current date in YYYY-MM-DD format

The purpose of this is extensibility. One example is making a graph with gnuplot
representing your streak lengths for a habit over time. Another example is
making a "forgot to mark" system (perhaps you completed a habit yesterday but
forgot to mark it so now it has been reset).

### Simple implementation examples
#### Ignore missed days
/etc/profile:
```sh
export RICH_HOOK=:
```

#### Decrease the streak by five on a missed day
/etc/profile:
```sh
export RICH_HOOK=/home/michael/sync/shell/rich_hook
```
/home/michael/sync/shell/rich_hook:
```sh
#!/bin/sh
habit_name=${1%%*/}
prev_streak=$3
new_streak=$((prev_streak - 5))

rich set $habit_name $new_streak
```

## wealth
I've provided an example wrapper (with its own README) in the
[wealth/](https://github.com/michaelskyba/rich/tree/master/wealth) directory.
