package main

import (
	tm "pomodoro-timer/time"

	in "github.com/tgorman31/Go-Modules/input"
)

func main() {

	wrk := in.User_Input("Please enter a work tm as such ##:##")
	// brk := in.User_Input("Please enter a break tm as such ##:##")

	wMin, wSec := tm.Time_Parse(wrk)
	// bMin, bSec := tm.Time_Parse(brk)

	// tm.Timer("Work", wMin, wSec)
	// tm.Timer("Break", bMin, bSec)
	tm.Timer_Tea("NewWork", wMin, wSec)
	tm.Launch_Timer()
}
