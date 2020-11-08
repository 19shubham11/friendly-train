package main

import (
	"testing"
	"bytes"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Shubham")

	got := buffer.String()
	want := "Hello, Shubham"

	if got != want{
		t.Errorf("got - %q, want- %q", got, want)
	}

}