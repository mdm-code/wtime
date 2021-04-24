/*
This tool lets you count your time spent at work. It lets you specify the
amount of time you spend sitting plus a break. For example, you can declare
that you want to spend 25 minutes working and then take a five-minute break.
*/
package main

import (
	"fmt"
	"net"
	"strings"
)

// TODO: I want to have a server running on Unix Domain Sockets or UDS for short

const SOCK = "/tmp/wtime.sock"

func main() {
}

type Server interface {
	Run() error
	Close() error
}

type ServerUDS struct {
	addr string
	serv net.Listener
}

func (s *ServerUDS) Run() error   { return fmt.Errorf("") }
func (s *ServerUDS) Close() error { return fmt.Errorf("") }

// Create fresh server with the specified protocol and address.
// The addr attribute is loosely understood and here it corresponds
// to a file name on the localhost. This is the way UDS works.
func NewServer(prot, addr string) (Server, error) {
	switch strings.ToLower(prot) {
	case "unix":
		return &ServerUDS{
			addr: addr,
		}, nil
	case "tcp":
	}
	return nil, fmt.Errorf("unsupported protocal given")
}
