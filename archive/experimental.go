package main

    /*choices, err := loadChoicesFromJSON("./../bible/kjv.json")    #Use this later
    if err != nil {
        // fallback or exit
        fmt.Println("Error loading choices:", err)
        os.Exit(1)
    }*/

import (
    "os"
    "fmt"
    //s "strings"
    //e "errors"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

var (
    checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))   //styling

    //Box "css"     optionally use Align(lipgloss.Center)
    MenuBoxStyle = lipgloss.NewStyle().
        Padding(1, 2).
        Border(lipgloss.NormalBorder())
)

func (m model) Init() tea.Cmd {
    return tea.SetWindowTitle("Vible") //Window title
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
    return homeUpdate(m, msg)
    //Add a switch when there are more
}

func (m model) View() string {
    switch m.state {
    case home:
        return homeView(m)
    case lookup:
        return lookupView(m)
    default:
        return homeView(m)
    }
}

func homeView(m model) string {     //Default selection screen
    s := "Select prefered mode \n\n"

    for i, choice := range m.choices {

        cursor := " "   //no cursor
        if m.cursor == i {
            cursor = ">"
        }

        checked := " "
        if _, ok := m.selected[i]; ok {
            checked = "x"
        }

        // This does the rendering
        s+= fmt.Sprintf("%s [%s] %s \n\n", cursor, checked, choice)

    }
    tip := "\n q to Quit"

    box := MenuBoxStyle.Render(s)
    centeredBox := lipgloss.Place(      //Manually place it so "tip" can go outside
        m.width, m.height, lipgloss.Center, lipgloss.Center, box,
    )

    return centeredBox + tip
}

func homeUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type){
        case tea.KeyMsg:    //Check for key press

            switch msg.String() {

            case "ctrl+c", "q":
                return m, tea.Quit

            case "up", "k":
                if m.cursor > 0 {
                    m.cursor--
                }

            case "down", "j":
                if m.cursor < len(m.choices)-1 {
                    m.cursor++
                }

            case "enter", " ":
                _, ok := m.selected[m.cursor]

                if ok {
                    delete(m.selected, m.cursor)
                } else {
                    m.selected[m.cursor] = struct{}{}

                    switch m.cursor {
                    case 1:
                        m.state = lookup
                    default:
                        m.state = home
                    }
                }
            }
        case tea.WindowSizeMsg: //For centering and positioning
            m.width = msg.Width
            m.height = msg.Height
    }

    return m, nil
}

func lookupView(m model) string {
    s := "This is a test view"
    return s
}

func main() {
    p := tea.NewProgram(initialModel(), tea.WithAltScreen())    //tea.WithAltScreen makes it full screen
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

func initialModel() model {
    return model{
        choices: []string{ "Continue reading", "Lookup", "Start Genesis 1:1"},
        selected: make(map[int]struct{}),
        state: home,
    }
}

type viewState int

const (
    home viewState = iota
    lookup
)

type model struct {
    choices []string
    cursor int
    selected map[int]struct{}
    state viewState

    width int
    height int
}