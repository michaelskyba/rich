# wealth
``wealth`` is an example of a user-friendly rich wrapper.

It's not expected to follow the UNIX philosophy. If you're one of the two people
who are familiar with my other projects, ``wealth``'s relationship to ``rich`` is
the same as ``shovel``'s to ``tunnel``.

## Installation
```sh
cd /path/to/rich
cd wealth
./install
```

## Order file syntax
The order file controls the habit order in menus. Each line should either be
blank, a comment, or the filename of a habit. Here is a valid example.

The habits:
```sh
$ ls ${RICH_HOME:-$HOME/.local/share/rich}
2500ml_water
30m_treadmill
no_homicide
```

The order file:
```
# Passive habits
no_homicide

# Active habits
2500ml_water
30m_treadmill
```
As you can see, comments are lines that start with ``#``. ``x # y`` is not a
valid comment.

## Status
I'm pretty sure that it should mostly work, but I have done zero testing.

## Dependencies
- rich
- Go
- fzf (used as a menu and a prompt)
- Basics like coreutils and a shell that is at least somewhat POSIX-compliant

## Example screenshot
![screenshot](https://raw.githubusercontent.com/michaelskyba/rich/master/wealth/assets/screenshot.webp)

(Of the "Toggle habit completion." screen)
