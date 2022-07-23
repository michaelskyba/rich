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

## Status
wealth is mostly in a usable state but I haven't done much testing yet,
especially for abstract edge cases.

## Dependencies
- rich
- Go
- fzf (used as a menu and a prompt)
- Basics like coreutils and a shell that is at least somewhat POSIX-compliant

## TODO
- Add error message on forgotten habits screen if there are none
