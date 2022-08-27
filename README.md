# Rich
rich: a simple, streak-oriented, daily habit tracker written in Go

## Installation
Clone the repo, run ``go build``, and then copy the resulting ``rich`` binary
into your $PATH.

## Overview
rich stores each habit in a file, where the name of the file is the same name
that you gave the habit. The layout is as follows:
```
<date of the last time the habit was performed>
<the current streak>
```
These habit files will be stored in ``$RICH_HOME``, or ``$HOME/.local/share/rich``
if ``$RICH_HOME`` is unset. **rich does not create these directories automatically**.
``mkdir`` your directory manually before your first use.

## Commands
- ``rich new <habit name> [streak]`` - create a new habit.
The streak will be set to "0" by default (no streak yet). The last date will be
set to yesterday; the habit is not marked as complete for today. If your
streak includes today, submit (streak - 1), and then run ``rich mark <habit name>``.
- ``rich delete <habit name>`` - delete a habit. There is no confirmation prompt.
- ``rich mark <habit name> ...`` - mark one more habits as complete for today.
- ``rich set <habit name> <streak>`` - manually set the streak of a habit.
- ``rich streak <habit name>`` - print the streak length of a habit.
- ``rich todo [habit name]`` - by default, list habits that have yet to be
completed today. If you provide a habit as an additional argument, instead of
listing anything, it will output "todo" (with exit code 0) or "done" (with exit
code 1) depending on whether the habit has been marked today.
- ``rich [list]`` - list all existing habits with padded spacing, sorted by
streak length.

## $RICH_HOOK
When a day passes for which a habit is not marked as completed, the habit
becomes "missed". By default (i.e. no hook is set), when this happens, the
habit's streak is reset to zero. However, you can override (or add) to this
behaviour through a hook.

To do so, set ``$RICH_HOOK`` to an executable (can be any programming
language). The following arguments will be passed in this order:
1. full path to habit file
2. last completion date in YYYY-MM-DD format
3. streak length before reset
4. current date in YYYY-MM-DD format

The purpose of this is extensibility. One example is making a graph with
gnuplot representing your streak lengths for a habit over time. Another example
is making a "forgot to mark" system (perhaps you completed a habit yesterday
but forgot to mark it so now it has been reset).

## wealth
I've provided an example wrapper called (with its own README) in the
[wealth/](https://github.com/michaelskyba/rich/tree/master/wealth) directory.
