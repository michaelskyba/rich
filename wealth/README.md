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
go build wealth-date.go
su -c "cp wealth wealth-* /usr/local/bin/"

# Initial configuration setup
mkdir -p $XDG_CONFIG_HOME/wealth
ls $XDG_CONFIG_HOME/rich > $XDG_CONFIG_HOME/wealth/order
```

## Status
wealth is not in a usable state but is (probably) in development.

## Dependencies
- rich
- Go
- fzf (used as a menu)
- Basics like coreutils and a shell that is at least somewhat POSIX-compliant

## TODO
- Make sure you're always using wealth-order instead of the raw order file
- Make sure you use grep unambiguously
	e.g. if you have a habit called "exercise", it shouldn't be matched for
	"exercise-run"
