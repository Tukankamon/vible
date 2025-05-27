package main

import (
    b "github.com/Tukankamon/vible/app/backend"

    "os"
    "fmt"
    //s "strings"
    //e "errors"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/charmbracelet/bubbles/viewport"       //needed for the struct, maybe I can define it in read.go
)

type (  //Dont really know what this is
    errMsg error
)

var (
    checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))   //unused

    //Box "css"     optionally use Align(lipgloss.Center)
    MenuBoxStyle = lipgloss.NewStyle().
        Padding(1, 2).
        Border(lipgloss.NormalBorder())

   searchBarStyle = lipgloss.NewStyle().
        Padding(0, 1).
        Border(lipgloss.RoundedBorder()).
        Width(50)

   quoteStyle = lipgloss.NewStyle().
        Padding(0, 1).
        Width(50)
)

func (m model) Init() tea.Cmd {
    return tea.Batch(
        textinput.Blink,
        tea.SetWindowTitle("Vible"), // window title

    )
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
    switch m.state {
    case home:
        return homeUpdate(m, msg)
    case lookup:
        return LookupUpdate(m, msg)
    case lookupQuote:
        return LookupUpdate(m, msg)
    case read:
        return ReadUpdate(m, msg)
    case open:
        return OpenUpdate(m, msg)
    case opened:
        return OpenUpdate(m, msg)
    default:
        homeUpdate(m, msg)
        }

	switch msg := msg.(type) {  //Needed for read.go
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
    return m, nil
}


func (m model) View() string {
    switch m.state {
    case home:
        return homeView(m)
    case lookup:
        return LookupView(m)
    case lookupQuote:
        return LookupQuoteView(m)
    case read:
        return ReadView(m)
    case open:
        return OpenView(m)
    case opened:
        return OpenView(m)
    default:
        homeView(m)
    }
    return ""
}

func homeView(m model) string {     //Default selection screen
    s := "Select prefered mode \n\n"

    for i, choice := range m.choices {

        cursor := " "   //no cursor
        if m.cursor == i {
            cursor = ">"
        }

        // This does the rendering
        s+= fmt.Sprintf("%s [ ] %s \n\n", cursor, choice)

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

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "1 Kings 2:3"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

    return model{
        choices: []string{  "Lookup", "Open chapter", "Continue reading", "Start Genesis"},
        selected: make(map[int]struct{}),
        state: home,
        input: ti,
        err: nil,
    }
}

type viewState int

const (
    home viewState = iota

    lookup
    lookupQuote

    read

    open
    opened  //Chapter open but with the change of opening another
)

type model struct {
    err error

    choices []string
    cursor int
    selected map[int]struct{}
    state viewState

    width int       //For centering
    height int

    input textinput.Model    //For search.go
    quote string    //The result of looking up a verse

    content  string     //for read.go
	ready    bool
	viewport viewport.Model
}

/*
func checkBoxView(m model) string {     //Currently not used but saved for the future
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
*/