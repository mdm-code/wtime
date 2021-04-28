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
		protocol          string
		addr              string
		work, rest, emoji string
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, USAGE)
		flag.PrintDefaults()
	}
	flag.StringVar(&protocol, "protocol", "unix", "server protocol")
	flag.StringVar(&addr, "addr", SOCK, "socket address")
	flag.StringVar(&work, "work", "25m", "time for work")
	flag.StringVar(&rest, "rest", "5m", "time for work")
	flag.StringVar(&emoji, "emoji", EMOJIS, "rotate these emojis")
	flag.Parse()

	s, err := NewServer(protocol, addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wtime %s", err)
		os.Exit(1)
	}
	defer s.Close()

	wDur, err := time.ParseDuration(work)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wtime: %s", err)
		os.Exit(2)
	}
	rDur, err := time.ParseDuration(rest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wtime: %s", err)
		os.Exit(2)
	}

	if err := s.Run(wDur, rDur); err != nil {
		fmt.Fprintf(os.Stderr, "wtime %s", err)
	}
}

// UDS server layout.
type Server struct {
	addr, protocol string
	serv           net.Listener
}

// Run launches the server that now listens for connections.
func (s *Server) Run(w, r time.Duration) (err error) {
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
		go HandleConn(conn, w, r)
	}
	return
}

// Close the listener server.
func (s *Server) Close() {
	s.serv.Close()
}

// Transmit the message to connected client(s).
func HandleConn(conn net.Conn, w, r time.Duration) {
	defer conn.Close()
	emojis := []rune(EMOJIS)
	countdown(conn, w, emojis[0])
	countdown(conn, r, emojis[1])
}

func countdown(conn net.Conn, d time.Duration, e rune) {
	timer := time.NewTimer(d)
	ticker := time.NewTicker(TICK)

loop:
	for {
		select {
		case <-ticker.C:
			d -= time.Duration(TICK)
			_, err := io.WriteString(conn, string(e)+" "+d.String()+"\n")
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
	return Server{
		addr:     addr,
		protocol: protocol,
	}, nil
}
