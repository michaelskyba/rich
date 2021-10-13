# Rich
rich: a simple, streak-oriented, daily habit tracker written in Go

## Installation
Clone the repo, run ``go build rich.go``, and then copy the resulting ``rich``
binary into your $PATH.

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
Streak will be set to "0" by default (no streak yet). The last date will be set 
to yesterday - i.e. the habit is not marked as complete for today. If your streak
includes today, just submit (streak - 1), and then run ``rich mark <habit name>``.
- ``rich mark <habit name>`` - mark a habit as complete for today.
- ``rich todo`` - list habits that have yet to be completed today.
- ``rich`` - list all existing habits. The justification for this feature's existence
is that it will also give you the streak number of each habit and sort habits accordingly.

## $RICH_HOOK
rich provides a hook which will run when a habit is being reset. Set ``$RICH_HOOK`` to a
an executable file (the language doesn't matter). The following arguments will be passed
in the order I have indicated:
``<filename of habit> <last completion date> <streak length before reset> <current date>``.

The purpose of this is extensibility. One example is making a graph with gnuplot representing
your streak lengths for a habit over time. Another example is making a "forgot to mark" system
(perhaps you completed a habit yesterday but forgot to mark it so now it has been reset).
Without a hook, neither of these could not be implemented elegantly: you would have to reimplement
rich's late mark detection.
