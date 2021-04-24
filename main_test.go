package main

import "testing"

// Check if the placeholder Greeter says hi.
func TestGreeter(t *testing.T) {
	want := "Hello, world!\n"
	has := Greeter()
	if want != has {
		t.Errorf("want %s; has %s", want, has)
	}
}
