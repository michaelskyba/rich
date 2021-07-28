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

## Commands
- ``rich new <habit name> [streak]`` - create a new habit.
Streak will be set to "0" by default (no streak yet). The last date will be set 
to yesterday - i.e. the habit is not marked as complete for today. If your streak
includes today, just submit (streak - 1), and then run rich mark <habit name>.
- ``rich mark <habit name>`` - mark a habit as complete for today.
- ``rich todo`` - list habits that have yet to be completed today.
- ``rich`` - list all existing habits. The justification for this feature's existence
is that it will also give you the streak number of each habit and sort habits accordingly.
