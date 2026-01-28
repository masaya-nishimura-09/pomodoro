package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor int
	timers []string
	page   int
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Select Timer")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.timers)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.page = 1
		}
	}

	return m, nil
}

func (m model) View() string {
	switch m.page {
	case 0:
		s := "Select timer type\n\n"

		for i, timer := range m.timers {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}

			s += fmt.Sprintf("%s %s\n", cursor, timer)
		}

		s += "\nPress q to quit.\n"

		return s
	case 1:
		s := fmt.Sprintf("You selected [%s].\n\n", m.timers[m.cursor])
		return s
	}

	return ""
}

func initialModel() model {
	return model{
		timers: []string{"Focus", "Short Break", "Long Break"},
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
