package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestGreeting(t *testing.T) {
	var buf bytes.Buffer
	Greeting(&buf, "Chethan", "Welcome")
	expected := "Hello Chethan!! Welcome\n"
	got := buf.String()
	if expected != got {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
	//assert.Equal(t, "Hello Chethan!! Welcome\n", buf.String())
}

func TestGreeting1(t *testing.T) {
	originalStdOutput := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	Greeting1("Chethan")
	w.Close()
	os.Stdout = originalStdOutput
	expected := "Hi Chethan\n"
	bs, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	got := string(bs)
	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
	println("Done")
}
