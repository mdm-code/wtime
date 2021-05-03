/*
This tool lets you count your time spent at work. It lets you specify the
amount of time you spend sitting plus a break. For example, you can declare
that you want to spend 25 minutes working and then take a five-minute break.
*/
package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

const (
	USAGE = `usage:
   wtime
   wtime -work=25m -rest=5m
   wtime -work=3h30m43s

params:
`
	SOCK   = "/tmp/wtime.sock"
	TICK   = time.Second
	EMOJIS = "💣😴"
)

func main() {
	var (
		protocol           string
		addr               string
		work, rest, emojis string
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, USAGE)
		flag.PrintDefaults()
	}
	flag.StringVar(&protocol, "protocol", "unix", "server protocol")
	flag.StringVar(&addr, "addr", SOCK, "socket address")
	flag.StringVar(&work, "work", "25m", "time for work")
	flag.StringVar(&rest, "rest", "5m", "time for rest")
	flag.StringVar(&emojis, "emoji", EMOJIS, "emojis to loop through")
	flag.Parse()

	s, err := NewServer(protocol, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wtime %s", err)
		os.Exit(1)
	}
	defer s.Close()

	workDur, err := time.ParseDuration(work)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wtime: %s", err)
		os.Exit(2)
	}
	restDur, err := time.ParseDuration(rest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wtime: %s", err)
		os.Exit(2)
	}

	split := []rune(emojis)
	if l := len(split); l != 2 {
		fmt.Fprintf(os.Stderr, "wtime: two emojis needed but %d was given", l)
	}
	if err := s.Run(workDur, restDur, split); err != nil {
		fmt.Fprintf(os.Stderr, "wtime %s", err)
	}
}

// UDS server layout.
type Server struct {
	addr, protocol string
	serv           net.Listener
}

// Run launches the server that now listens for connections.
func (s *Server) Run(work, rest time.Duration, emojis []rune) (err error) {
	s.serv, err = net.Listen(s.protocol, s.addr)
	for {
		conn, err := s.serv.Accept()
		if err != nil {
			err = fmt.Errorf("unable to establish connection")
			break
		}
		if conn == nil {
			err = fmt.Errorf("unable to communicate")
			break
		}
		go handleConn(conn, work, rest, emojis)
	}
	return
}

// Close down the server.
func (s *Server) Close() {
	s.serv.Close()
}

// Transmit the message to the connected client.
func handleConn(conn net.Conn, work, rest time.Duration, emojis []rune) {
	defer conn.Close()
	countdown(conn, work, emojis[0])
	countdown(conn, rest, emojis[1])
}

// Count down
func countdown(conn net.Conn, d time.Duration, emoji rune) {
	timer := time.NewTimer(d)
	ticker := time.NewTicker(TICK)
loop:
	for {
		select {
		case <-ticker.C:
			d -= time.Duration(TICK)
			_, err := io.WriteString(conn, "\r"+string(emoji)+" "+d.String())
			if err != nil {
				break loop
			}
		case <-timer.C:
			break loop
		}
	}
}

// NewServer returns a fresh instance of a server on a Unix Domain Socket.
func NewServer(protocol, addr string) (Server, error) {
	if err := os.RemoveAll(addr); err != nil {
		return Server{}, err
	}
	if p := strings.ToLower(protocol); p != "unix" {
		return Server{}, fmt.Errorf("%s protocol not implemented", p)
	}
	return Server{
		addr:     addr,
		protocol: protocol,
	}, nil
}
