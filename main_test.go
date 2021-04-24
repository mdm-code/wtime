package main

import "testing"

// Unit test server setup.
func init() {
	sock, prot := "/tmp/wtime.test.sock", "unix"
	serv, err := NewServer(sock, prot)
	if err != nil {
		return
	}
	// Spin up the server in a separate Goroutine to make sure test aren't
	// blocking.
	go func() {
		serv.Run()
	}()
}

// Verify server communicaton with a client.
func TestServer(t *testing.T) {
	_ = []struct {
		ttype string
		want  []byte
	}{
		{
			"Making request no. 1",
			[]byte("00:01\n"),
		},
		{
			"Making request no. 2",
			[]byte("00:02\n"),
		},
	}
}
