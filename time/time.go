package time

import (
	"fmt"
	"os"
	tmr "pomodoro-timer/timer"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// type TickMsg time.Time

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
	m := model{
		timer: tmr.New(time.Duration(s)),
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
			quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
		},
	}

	fmt.Printf("Running %s timer for %v seconds \n", name, s)
	m.keymap.start.SetEnabled(false)
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Uh oh, we encountered an error:", err)
		os.Exit(1)
	}
	// <-timer.C
	// fmt.Printf("%s timer completed \n", name)
}

func Time_Parse(time string) (min int, sec int) {

	t := strings.Split(time, ":")

	m, _ := strconv.Atoi(t[0])
	s, _ := strconv.Atoi(t[1])

	return m, s
}

const timeout = time.Second * 5

type model struct {
	timer    tmr.Model
	keymap   keymap
	quitting bool
}

type keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tmr.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case tmr.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.keymap.stop.SetEnabled(m.timer.Running())
		m.keymap.start.SetEnabled(!m.timer.Running())
		return m, cmd

	case tmr.TimeoutMsg:
		m.quitting = true
		return m, tea.Quit
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keymap.reset):
			m.timer.Timeout = timeout
		case key.Matches(msg, m.keymap.start, m.keymap.stop):
			return m, m.timer.Toggle()
		}

	}

	return m, nil
}

func (m model) View() string {
	var s string
	t := m.timer.Timeout

	// s := m.timer.View()

	if m.timer.Timedout() {
		s = "All done!"
	}
	s += "\n"
	if !m.quitting {
		s = "Time remaining " + time.Duration.String(t)
		s += "\n"
		s += "s - start/stop; r - reset; q - quit"
		// s = "Exiting in " + s
		// s += m.helpView()
	}
	return s
}

func Launch_Timer() {
	m := model{
		timer: tmr.New(timeout),
		keymap: keymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
			quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
		},
	}
	m.keymap.start.SetEnabled(false)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Uh oh, we encountered an error:", err)
		os.Exit(1)
	}
}
