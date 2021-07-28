# Rich
rich: a simple, streak-oriented habit tracker written in Go

## Installation
Clone the repo, run ``go build rich.go``, and then copy the resulting ``rich``
binary into your $PATH.

## Overview
rich stores each habit in a file, where the name of the file is the same name
that you gave the habit. The layout is as follows:
```
<date of the last time the habit was performed>
<the interval of days by which the habit should be performed (e.g. "7" for a weekly habit)>
<the current streak>
```
With this data, each of rich's functionality is able to be implemented.
By staying aware of this, you now understand how to extend rich, adding your own,
personal features.

## Commands
- ``rich new <habit name> [interval] [streak]`` - create a new habit. Interval will be set 
to "1" by default (daily) and streak will be set to "0" by default (no streak yet).
- ``rich mark <habit name>`` - mark a habit as performed for this interval.
- `` rich todo`` - list habits that have yet to be marked for their current interval.
- ``rich`` - list all existing habits. The justification for this feature's existence
is that it will also give you the streak number of each habit and sort habits accordingly.
