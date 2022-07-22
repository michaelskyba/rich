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
su -c "cp wealth wealth-order wealth-create /usr/local/bin/"

# Initial configuration setup
config=${XDG_CONFIG_HOME:-$HOME/.config}
mkdir -p $config/wealth
ls $config/rich > $config/wealth/order
```

## Status
wealth is not in a usable state but is (probably) in development.

## Dependencies
- rich
- Go
- A shell that is at least somewhat POSIX-compliant
- fzf (used as a menu)
