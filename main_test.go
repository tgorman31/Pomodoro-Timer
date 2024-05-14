package main

import (
	tm "pomodoro-timer/time"
	"testing"
)

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello World"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestTime_Parse(t *testing.T) {
	got1, got2 := tm.Time_Parse("00:10")
	want1 := 0
	want2 := 10

	if got1 != want1 {
		t.Errorf("got %q want %q", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %q want %q", got2, want2)
	}
}
