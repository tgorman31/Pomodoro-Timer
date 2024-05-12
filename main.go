package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	in "github.com/tgorman31/Go-Modules/input"
)

func Hello() string {
	return "Hello World"
}

// Takes an input of a timer name and minutes and seconds and starts a timer
func Timer(name string, min int, sec int) {
	s := (min * 60) + sec
	timer := time.NewTimer(time.Duration(s) * time.Second)

	fmt.Printf("Running %s timer for %v seconds \n", name, s)
	<-timer.C
	fmt.Printf("%s timer completed \n", name)
}

func Time_Parse(time string) (min int, sec int) {

	t := strings.Split(time, ":")

	m, _ := strconv.Atoi(t[0])
	s, _ := strconv.Atoi(t[1])

	return m, s
}

func main() {
	fmt.Println(Hello())

	wrk := in.User_Input("Please enter a work time as such ##:##")
	brk := in.User_Input("Please enter a break time as such ##:##")

	wMin, wSec := Time_Parse(wrk)
	bMin, bSec := Time_Parse(brk)

	Timer("Work", wMin, wSec)
	Timer("Break", bMin, bSec)
}
