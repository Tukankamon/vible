package main

//Manages the front-end when looking up a verse

import (
	b "github.com/Tukankamon/vible/app/backend"

    //"fmt"
    //"os"    //only used to exit

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/textinput"
)


func LookupView(m model) string {

    var text string
    switch m.state {
    case lookup, lookupQuote:
        text = "Search for quotes"
    case open:
        text = "Open a chapter"
    }
    bar := searchBarStyle.Render(m.input.View())
    centeredBar := lipgloss.Place(
        m.width, m.height,
        lipgloss.Center, lipgloss.Center,
        bar+ "\n\n\n" + text,
    )

    return centeredBar + "\n"
}
//Could join the two together
func LookupQuoteView(m model) string {      //After the first search, for some reason the bar moves down and it is driving me crazy
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
            bar + "\n\n\n" + m.quote,
        )  
        return centeredBar
    }
}

func LookupUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
    //To reset the input, probably are better ways to do this. Maybe set global var?
    ti := textinput.New()
	ti.Placeholder = "1 Kings 2:3"
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

            switch m.state {
            case lookupQuote, lookup:
                m.quote, m.err = b.Search(m.input.Value())
                m.state = lookupQuote
            case open:
                m.content, _ = b.Read(m.input.Value())
                m.state = opened    //Having trouble detecting errors here
            }
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
    }

    m.input, cmd = m.input.Update(msg)
    return m, cmd
}