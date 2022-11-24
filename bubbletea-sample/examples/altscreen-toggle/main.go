package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

var (
	color   = termenv.EnvColorProfile().Color
	keyword = termenv.Style{}.Foreground(color("204")).Background(color("235")).Styled
	help    = termenv.Style{}.Foreground(color("241")).Styled
)

type model struct {
	altscreen bool
	quitting  bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case " ":
			var cmd tea.Cmd
			if m.altscreen {
				cmd = tea.ExitAltScreen
			} else {
				cmd = tea.EnterAltScreen
			}
			m.altscreen = !m.altscreen
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "終了します\n"
	}

	const (
		altscreenMode = " 別画面モード "
		inlineMode    = " インラインモード "
	)

	var mode string
	if m.altscreen {
		mode = altscreenMode
	} else {
		mode = inlineMode
	}

	return fmt.Sprintf("\n\n  あなたは今%sです.\n\n\n", keyword(mode)) +
		help("    スペース: モード切替, q: 終了\n")
}

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
