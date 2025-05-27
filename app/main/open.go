package main
//Just joining read and search together

import (
	b "github.com/Tukankamon/vible/app/backend"
	
	"fmt"
	//"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/charmbracelet/bubbles/viewport"

)

func OpenView(m model) string {      //After the first search, for some reason the bar moves down and it is driving me crazy
    if m.state == open {
        bar := searchBarStyle.Render(m.input.View())
        if m.err != nil {
            centeredBar := lipgloss.Place(
                m.width, m.height,
                lipgloss.Center, lipgloss.Center,
                bar + "\n\n\n" + m.err.Error(),     //Print the error
            )
            return centeredBar
        } else {
            centeredBar := lipgloss.Place(
                m.width, m.height,
                lipgloss.Center, lipgloss.Center,
                bar + "\n\n\n" + "Open a chapter",
            )  
            return centeredBar
        }
    }

    // m.state == opened

    if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())

}

func OpenUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
    if m.state == open{      //stil havent searched anything
        ti := textinput.New()
        ti.Placeholder = "Exodus"
        ti.Focus()
        ti.CharLimit = 156
        ti.Width = 20
        var cmd tea.Cmd

        switch msg := msg.(type) {
        case tea.KeyMsg:
            switch msg.Type {
            case tea.KeyCtrlC:
                return m, tea.Quit
            case tea.KeyEsc:    //Go back to previous menu
                m.input = ti    //clear the text
                m.state = home
            case tea.KeyEnter:
                m.content, m.err = b.Read(m.input.Value())
                if m.err != nil {
                    fmt.Printf("ERROR: %v", m.err.Error())
                    return m, tea.Quit
                }
                m.state = opened
            }

        // We handle errors just like any other message
        case errMsg:
            m.err = msg
            return m, nil
        }

        m.input, cmd = m.input.Update(msg)
        return m, cmd
    }

    
    // m.sate == opened <--


	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:

		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			wrapped := lipgloss.NewStyle().Width(m.viewport.Width).Render(m.content)
			m.viewport.SetContent(wrapped)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}