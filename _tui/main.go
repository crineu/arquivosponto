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
	// Key bindings are defined in keys.go
)

// item represents a stow package with its installation status
type item struct {
	name   string
	status Status
}

func (i item) FilterValue() string { return i.name }
func (i item) String() string      { return i.name }

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

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	// Determine style based on status
	var statusStyle lipgloss.Style
	switch i.status {
	case Installed:
		statusStyle = installedStyle
	case Outdated:
		statusStyle = outdatedStyle
	case NotInstalled:
		statusStyle = notInstalledStyle
	default:
		statusStyle = itemStyle
	}

	str := fmt.Sprintf("%s %s", i.status.Emoji(), i.name)
	rendered := statusStyle.Render(str)

	fn := func(s ...string) string { return rendered }
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("▶ " + strings.Join(s, "-"))
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
	statuses map[string]Status // package name -> stow status
}

func (m *model) refreshStatus(idx int) {
	if idx < 0 || idx >= len(m.list.Items()) {
		return
	}
	if it, ok := m.list.Items()[idx].(item); ok {
		st := CalculateStatus(it.name)
		it.status = st
		m.list.SetItem(idx, it)
	}
}

func (m model) Init() tea.Cmd {
	// Calculate all statuses synchronously at startup
	for i := range m.list.Items() {
		m.refreshStatus(i)
	}
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
				m.choice = i.name
				m.status = StowRestowDry(m.choice)
				m.refreshStatus(m.list.Index())
			}
			return m, nil

		case "d":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.name
				m.status = StowRemDry(m.choice)
				m.refreshStatus(m.list.Index())
			}
			return m, nil

		case "i", "enter", "p":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.name
				m.status = StowAddDry(m.choice)
				m.refreshStatus(m.list.Index())
			}
			return m, nil

		case "esc":
			m.choice = ""
			m.status = nil
			return m, nil

		case "a":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.name
				m.status = StowAdd(m.choice)
				m.refreshStatus(m.list.Index())
			}
			return m, nil

		case "x":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.name
				m.status = StowRem(m.choice)
				m.refreshStatus(m.list.Index())
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func helpView() string {
	return keyHelpStyle.Render("i/⏎/p preview  •  a apply  •  d preview rm  •  x remove  •  r restow  •  esc close  •  q quit")
}

func (m model) View() string {
	if m.quitting {
		return quitTextStyle.Render("👋 Até mais")
	}
	if m.choice != "" {
		titleMsg := statusTextStyle.Render(fmt.Sprintf("%s stow output: ", m.choice))
		outputMsg := stowTextStyle.Render(strings.Join(m.status[:], "\n"))

		helpLine := helpView()
		return lipgloss.JoinVertical(
			lipgloss.Left,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				borderStyle.Render(m.list.View()),
				borderStyle.Render(titleMsg+outputMsg)),
			helpLine)
	}
	helpLine := helpView()
	return lipgloss.JoinVertical(lipgloss.Left, borderStyle.Render(m.list.View()), helpLine)
}

func main() {
	if !stowAvailable {
		fmt.Println("Error: GNU Stow is not installed or not in PATH.")
		fmt.Println("Please install stow: sudo apt install stow  (or equivalent for your distro)")
		os.Exit(1)
	}

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
		items = append(items, item{name: tool, status: NotInstalled})
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

	m := model{list: l, keys: newListKeyMap()}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func RaiseErrorAndExit(msg string, err error) {
	fmt.Fprintf(os.Stderr, msg+":: error: %v\n", err)
	os.Exit(1)
}
