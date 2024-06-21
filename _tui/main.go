package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Green  = "#25A065"
	// White  = "#FFFDF5"
	// Pink   = "#FF06B7"
	// Gray   = "#768676"
	// Yellow = "#FF9F1C"
	// Blue   = "#02A9EA"
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Align(lipgloss.Center).
			Padding(0, 1)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Background(lipgloss.Color("#25A065"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	statusTextStyle   = lipgloss.NewStyle().Margin(1, 0, 2, 4).Foreground(lipgloss.Color("#FF9F1C"))
	stowTextStyle     = lipgloss.NewStyle().Margin(1, 0, 1, 4).Foreground(lipgloss.Color("#767676"))
	borderStyle       = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#25A065")).
				PaddingRight(2)
)

type listKeyMap struct {
	stowAdd    key.Binding
	stowRemove key.Binding
	stowRestow key.Binding
	stowCheck  key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		stowAdd: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "stow tool"),
		),
		stowRemove: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "delete tool"),
		),
		stowRestow: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "restow tool"),
		),
		stowCheck: key.NewBinding(
			key.WithKeys("c"),
			key.WithHelp("c", "check tool status"),
		),
	}
}

type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("🌟" + strings.Join(s, "-"))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	keys     *listKeyMap
	choice   string
	status   []string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "r":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
				m.status = StowRestowDry(m.choice)
			}
			return m, nil

		case "d":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
				m.status = StowRemDry(m.choice)
			}
			return m, nil

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
				m.status = StowAddDry(m.choice)
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return quitTextStyle.Render("👋 Até mais")
	}
	if m.choice != "" {
		// titleMsg := statusTextStyle.Render(fmt.Sprintf("%s stow output: ", m.choice))
		titleMsg := ""
		outputMsg := stowTextStyle.Render(strings.Join(m.status[:], "\n"))

		return lipgloss.JoinHorizontal(
			lipgloss.Top,
			borderStyle.Render(m.list.View()),
			borderStyle.Render(titleMsg+outputMsg))
	}
	return borderStyle.Render(m.list.View())
}

func main() {
	if os.Getenv("HELP_DEBUG") != "" {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("Couldn't open a file for logging:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	var items []list.Item

	tools := ListArquivosPontoTools()
	for _, tool := range tools {
		items = append(items, item(tool))
	}

	const defaultWidth = 14
	const listHeight = 30

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "⭐ Arquivos Ponto ⭐\nFerramentas disponíveis"
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func RaiseErrorAndExit(msg string, err error) {
	fmt.Fprintf(os.Stderr, msg+":: error: %v\n", err)
	os.Exit(1)
}
