package main

import (
    "fmt"
    "os"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type model struct {
    width  int
    height int
}

func main() {
    if err := tea.NewProgram(model{}).Start(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
    case tea.KeyMsg:
        if msg.String() == "q" {
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string {
    content := "🍵 Lip Gloss Centered App\nPress q to quit."

    // Style for the content box
    boxStyle := lipgloss.NewStyle().
        Padding(1, 4).
        Border(lipgloss.DoubleBorder(), true).
        BorderForeground(lipgloss.Color("63")).
        Align(lipgloss.Center)

    // Render the content inside the box
    box := boxStyle.Render(content)

    // Center the box using lipgloss.Place
    return lipgloss.Place(
        m.width, m.height,
        lipgloss.Center, lipgloss.Center,
        box,
        lipgloss.WithWhitespaceChars(" "),
        lipgloss.WithWhitespaceForeground(lipgloss.Color("238")),
    )
}