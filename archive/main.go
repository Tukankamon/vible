package main

    /*choices, err := loadChoicesFromJSON("./../bible/kjv.json")    #Use this later
    if err != nil {
        // fallback or exit
        fmt.Println("Error loading choices:", err)
        os.Exit(1)
    }*/

import (
    "os"
    "encoding/json"
    "fmt"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

var (
    checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))   //styling
)

// Reads the top-level keys from a JSON file and returns them as a slice of strings.
func loadChoicesFromJSON(filename string) ([]string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var obj map[string]interface{}
    if err := json.Unmarshal(data, &obj); err != nil {
        return nil, err
    }

    choices := make([]string, 0, len(obj))
    for k := range obj {
        choices = append(choices, k)
    }
    return choices, nil
}

func (m model) Init() tea.Cmd {
    return tea.SetWindowTitle("Vible") //Window title
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
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
                }
            }
    }

    return m, nil
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

func (m model) View() string {
    s := "This is a generic string \n"

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
        s+= fmt.Sprintf("%s [%s] %s \n", cursor, checked, choice)
    }
    s += "\n q to Quit"
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
    }
}

type model struct {
    choices []string
    cursor int
    selected map[int]struct{}
}