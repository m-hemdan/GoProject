package main

import (
	"log"
	"testing"
)

func TestParadise(t *testing.T) {
	got := Paradise("xxx")
	want := "this idea is xxx"

	if got != want {
		log.Fatalf("error - want is %s and %s", got, want)
	}
}
