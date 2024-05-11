package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

func User_Input(msg string) (out string) {
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)

	out, err := reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Error reading input: %e", err)
	}

	out = strings.TrimSpace(out)

	return out
}

func Time_Parse(time string) (min int, sec int) {
	// var t []string

	t := strings.Split(time, ":")

	m, _ := strconv.Atoi(t[0])
	s, _ := strconv.Atoi(t[1])

	return m, s
}

func main() {
	fmt.Println(Hello())
	wrk := User_Input("Please enter a work time as such ##:##")
	brk := User_Input("Please enter a break time as such ##:##")

	wMin, wSec := Time_Parse(wrk)
	bMin, bSec := Time_Parse(brk)

	Timer("Work", wMin, wSec)
	Timer("Break", bMin, bSec)
}
