# Wtime


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
wtime-srv -work=25 -break=5
```

It is all happening on the same host, so there are no ports here, just Unix
Domain Sockets (UDS). UDS-based implementation is faster and unambiguous.
You grab a socket file and channel all the communication through this one
file.

Then you can grab its output with Netcat without any hassle or whatever:

```sh
nc -U /tmp/wtime.sock
```

And then you can put it on your Tmux status line, for example.


## Development setup


### Development environment


### Testing
