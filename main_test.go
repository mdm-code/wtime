package main

import (
	"log"
	"net"
	"testing"
	"time"
)

// Spin up the UDS server.
func init() {
	protocol, addr := "unix", "/tmp/wtime-test.sock"

	workDur, err := time.ParseDuration("25m")
	if err != nil {
		log.Println("couldn't parse work duration")
	}
	restDur, err := time.ParseDuration("5m")
	if err != nil {
		log.Println("couldn't parse rest duration")
	}

	s, err := NewServer(protocol, addr)
	if err != nil {
		log.Println("couldn't start UDS server")
	}

	// Start the server in a goroutine not to block the tests
	go func() {
		s.Run(workDur, restDur, []rune(EMOJIS))
	}()
}

// Test the communication with the server.
func TestSeverComm(t *testing.T) {
	// The baby needs a moment to get rolling. Hold
	// on for a second to make sure it's up and running.
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("unix", "/tmp/wtime-test.sock")
	if err != nil {
		t.Error("couldn't connect to the server: ", err)
	}
	defer conn.Close()

	outputs := []string{"\r💣 24m59s", "\r💣 24m58s"}
	for i := 0; i < 2; i++ {
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			panic(err)
		}
		has := string(buf[:n])
		want := outputs[i]
		if has != want {
			t.Errorf("has: %s != want %s", has, want)
		}
	}
}

// UDS is the only protocol implemented at this point.
func TestTcpFails(t *testing.T) {
	protocol, addr := "tcp", ":8888"
	_, err := NewServer(protocol, addr)
	if err == nil {
		t.Errorf("protocol %s should not be allowed to pass", protocol)
	}
}
