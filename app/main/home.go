package main

import (
	b "github.com/Tukankamon/vible/app/backend"
	
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"fmt"
	
)

func HomeView(m model) string { //Default selection screen
	s := "Select prefered mode \n\n"

	for i, choice := range m.choices {

		cursor := " " //no cursor
		if m.cursor == i {
			cursor = ">"
		}

		// This does the rendering
		s += fmt.Sprintf("%s [ ] %s \n\n", cursor, choice)

	}
	tip := "\n q to Quit"

	box := MenuBoxStyle.Render(s)
	centeredBox := lipgloss.Place( //Manually place it so "tip" can go outside
		m.width, m.height, lipgloss.Center, lipgloss.Center, box,
	)

	return centeredBox + tip
}

func HomeUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg: //Check for key press

		switch msg.String() {

		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

            case "enter", " ", "l", "right":
                switch m.cursor {
                case 0:
                    m.input.Placeholder = "1 Kings 2:3"
                    m.state = lookup
                case 1:
                    m.input.Placeholder = "Exodus 1"      //Set it before state so it actually shows up
                    m.state = open
                case 3:
                    m.content, _ = b.Read("Genesis 1:1")
                    m.state = read
                default:
                    m.state = home
                }
            }
        case tea.WindowSizeMsg: //For centering and positioning
            m.width = msg.Width
            m.height = msg.Height
    }

	return m, nil
}
