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
- ``rich todo`` - list habits that have yet to be completed today.
- ``rich [list]`` - list all existing habits. The justification for this feature's
existence is that it will also give you the streak number of each habit and
sort habits accordingly.

## $RICH_HOOK
rich provides a hook which will run when a habit is being reset. Reset means 
being "reset to zero", so a check will not count as reset if the habit streak 
is already zero. Set ``$RICH_HOOK`` to an executable (could be any programming
language).

The following arguments will be passed in this order:
1. full path to habit file
2. last completion date in YYYY-MM-DD format
3. streak length before reset
4. current date in YYYY-MM-DD format

The purpose of this is extensibility. One example is making a graph with
gnuplot representing your streak lengths for a habit over time. Another example
is making a "forgot to mark" system (perhaps you completed a habit yesterday
but forgot to mark it so now it has been reset).

Without a hook, both of these would be annoying to carry out. You would have to
reimplement rich's late mark detection or keep backups of ``rich list`` output
and check for zeros in the current output.
