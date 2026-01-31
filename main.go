package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	screen     int
	homeModel  homeModel
	timerModel timerModel
}

type homeModel struct {
	cursor int
	timers []string
}

type timerModel struct {
	timerType int
	timers    []string
	time      int
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
			if m.homeModel.cursor > 0 {
				m.homeModel.cursor--
			}
		case "down", "j":
			if m.homeModel.cursor < len(m.homeModel.timers)-1 {
				m.homeModel.cursor++
			}
		case "enter", " ":
			m.screen = 1
			m.timerModel.timerType = m.homeModel.cursor
		}
	}

	return m, nil
}

func (m timerModel) View() string {
	s := fmt.Sprintf("%s\n\n", m.timers[m.timerType])

	switch m.timerType {
	case 0:
		s += fmt.Sprintf("%s %s\n", cursor, timer)
	}
	return s
}

func (m homeModel) View() string {
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
}

func (m model) View() string {
	switch m.screen {
	case 0:
		return m.homeModel.View()
	case 1:
		return m.timerModel.View()
	}

	return ""
}

func initialModel() model {
	return model{
		screen: 0,
		homeModel: homeModel{
			timers: []string{"Focus", "Short Break", "Long Break"},
		},
		timerModel: timerModel{
			timers: []string{"Focus", "Short Break", "Long Break"},
		},
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
