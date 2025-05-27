package main
//Just joining read and search together

import (
	//b "github.com/Tukankamon/vible/app/backend"
	
	"fmt"
	//"strings"
    //e "errors"
	tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/lipgloss"
    //"github.com/charmbracelet/bubbles/textinput"
    //"github.com/charmbracelet/bubbles/viewport"

)

func OpenView(m model) string {      //After the first search, for some reason the bar moves down and it is driving me crazy
    switch m.state {
    case open:
        return LookupView(m)
    case opened:
        if !m.ready {
		    return "\n  Initializing..."
        }
	    return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
    default:
        return LookupView(m)
    }

}

func OpenUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
    
    switch m.state {
    case open:
        return LookupUpdate(m, msg)
    case opened:
        return ReadUpdate(m, msg)
    default:
        return LookupUpdate(m, msg)
    }
}