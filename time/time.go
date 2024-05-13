package time

import (
	"fmt"
	"pomodoro-timer/timer"
	tmr "pomodoro-timer/timer"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

// Takes an input of a timer name and minutes and seconds and starts a timer
func Timer(name string, min int, sec int) {
	s := (min * 60) + sec
	timer := time.NewTimer(time.Duration(s) * time.Second)

	fmt.Printf("Running %s timer for %v seconds \n", name, s)
	<-timer.C
	fmt.Printf("%s timer completed \n", name)
}

func Timer_Tea(name string, min int, sec int) {
	s := (min * 60) + sec
	// timer := time.NewTimer(time.Duration(s) * time.Second)
	m := model{timer: timer.New(time.Duration(s))}

	m.Init()
	// fmt.Printf("Running %s timer for %v seconds \n", name, s)
	// <-timer.C
	// fmt.Printf("%s timer completed \n", name)
}

func Time_Parse(time string) (min int, sec int) {

	t := strings.Split(time, ":")

	m, _ := strconv.Atoi(t[0])
	s, _ := strconv.Atoi(t[1])

	return m, s
}

type model struct {
	timer tmr.Model
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	t := m.timer.Timeout

	s := "Time remaining " + time.Duration.String(t)
	return s
}
