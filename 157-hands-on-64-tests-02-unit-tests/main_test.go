package main

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(7, 5)
	want := 14
	if got != want {
		log.Fatalf("Sum was incorrect, got: %d, want: %d.", got, want)
	}

}
