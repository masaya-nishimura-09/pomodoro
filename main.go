package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/masaya-nishimura-09/pomodoro/model"
)

func initialModel() model.Model {
	return model.Model{
		Timers:   []string{"Focus", "Short Break", "Long Break"},
		Selected: make(map[int]struct{}),
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
