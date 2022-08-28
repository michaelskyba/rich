# wealth
``wealth`` is an example of a user-friendly rich wrapper.

It's not expected to follow the UNIX philosophy. If you're one of the two people
who are familiar with my other projects, ``wealth``'s relationship to ``rich`` is
the same as ``shovel``'s to ``tunnel``.

## Installation
```sh
cd /path/to/rich
cd wealth

go build wealth-order.go
su -c "cp wealth wealth-* /usr/local/bin/"

# Initial configuration setup
mkdir -p $XDG_CONFIG_HOME/wealth
ls $XDG_CONFIG_HOME/rich > $XDG_CONFIG_HOME/wealth/order
```

## Order file syntax
The order file controls the habit order in menus. Each line should either be
blank, a comment, or the filename of a habit. Here is a valid example.

The habits:
```sh
$ ls $XDG_CONFIG_HOME/rich
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
~~wealth is mostly in a usable state but I haven't done much testing yet,
especially for abstract edge cases.~~
2022-08-28: wealth is currently completely unusable. Ignore everything else in
this README until I have made it compatible with the new rich design.

## Dependencies
- rich
- Go
- fzf (used as a menu and a prompt)
- Basics like coreutils and a shell that is at least somewhat POSIX-compliant

## Example screenshot
![screenshot](https://raw.githubusercontent.com/michaelskyba/rich/master/wealth/assets/screenshot.webp)

(Of the "Toggle habit completion." screen)

## TODO
- Add error message on forgotten habits screen if there are none
- Improve the UI consistency (e.g. right now there is an echo "no habits" screen
and a fzf "no habits" screen, which makes no sense)
