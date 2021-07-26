# Rich
rich: a simple, streak-oriented habit tracker written in Go

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
