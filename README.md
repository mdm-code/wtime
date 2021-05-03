[![Actions Status](https://github.com/mdm-code/wtime/workflows/CI/CD/badge.svg)](https://github.com/mdm-code/wtime/actions)

# Wtime

This program gives you better control over your time at work by splitting work
into manageable chunks of time. Conversely, you can use it to make sure your
stand up and exercise regularly.


## Introduction

Sitting for hours without a break in front of a computer is bad for your health
and general wellbeing. Every thirty minutes or so, you might want to stand up
and do a little bit of exercise so that you do not put too much strain on your
spine and your body in general.

Splitting your work time into manageable chunks can give a better control over
the time spent working so that you do not go out of focus or start doing
overtime because you don't really now how long you've worked.

The program let's you control your time at work by splitting it into productive
sessions followed by short breaks. Say, you might want to split your time into
thirty-minute-long intervals with twenty-five minutes of work followed by a
five-minute break. You sit back for twenty-five minutes, stay focused on a task
and then you stand up, do some work that requires you to move, or if there is
nothing else to do, then do a little bit of stretching, push-ups or whatever
rocks your boat. You won't regret it.


## How to use it?

You can use it and modify it however you like, but if I can suggest something,
then I would stick to a very basic and non-intrusive usage.

All options that you can pass at the startup can be accessed with the `--help`
flag.

First, you kick start the server:

```sh
wtime -work=25 -rest=5
```

You might want to pass your own alternating emoji to the `-emojis` parameter.
There should be no more no less but two of them.

It is all happening on the same host, so there are no ports here, just Unix
Domain Sockets (UDS). UDS-based implementation is faster and unambiguous.
You grab a socket file and channel all the communication through this one
file.

Then you can grab its output with `netcat` without any hassle or whatever:

```sh
nc -U /tmp/wtime.sock
```

Or you can put it on your `tmux` status line, for example, with this line:

```
set -g status-right "#(cat /tmp/wtime.sock) %A, %B %-e, %Y, %-l:%M:%S%p"
```

You can use some of that `cat` magic and dial in to the socket in loops to get
a counter for work time followed by a moment of respite.


## Development setup

There isn't much going on here for Go as all tools ship with the compiler.
All core commands are specified in the `Makefile`.


### Development environment

You might want to install the package to see how it works:

```sh
make install
```

Make sure your Go environment is properly set up, but this is something that
you should take care on your own. There are many ways, and I am not elaborate
on it here and now.


### Testing

To run the test suite key in:

```sh
make test

# or

go test -v ./...
```

To build the binary:

```sh
go build --race -o bin/wtime main.go
```

You can do both things at once by typing bare `make`.

