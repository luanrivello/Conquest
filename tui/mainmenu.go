package tui

import (
	"fmt"
	"log"
	"os"

	"conquest/loop"
	"conquest/tui/colors"
	"conquest/tui/conquestmode"

	tea "github.com/charmbracelet/bubbletea"
)

var defaultColor = colors.BLUE
var selectedColor = colors.GREEN

// * Main Menu * //
func MainMenu() {
	//* Clear Terminal
	fmt.Print("\033[H\033[2J")

	//* Initialize Bubbletea TUI
	tui := tea.NewProgram(initialModel())
	if err := tui.Start(); err != nil {
		fmt.Fprint(os.Stderr, err)
		log.Fatal(err)
		os.Exit(1)
	}

	//* Clear Terminal
	fmt.Print("\033[H\033[2J")
}

// * Model * //
type mainModel struct {
	choices []string
	cursor  int
}

// * Constructor * //
func initialModel() mainModel {
	return mainModel{
		choices: []string{"Conquest", "Explore", "Encyclopedia Galactica", "Dev", "Exit"},
	}
}

// * Startup * //
func (m mainModel) Init() tea.Cmd {
	return nil
}

// * Actions * //
func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		//* Keypress
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

		case "enter", " ", "l":
			//* Choices
			switch m.cursor {
			//* Conquest
			case 0:
				return conquestmode.GetCreationModel(m), nil

			//* Explore
			case 1:
				return conquestmode.GetCreationModel(m), nil

			//* Encyclopedia
			case 2:
				return conquestmode.GetCreationModel(m), nil

			//* Dev
			case 3:
				loop.Gameloop()

			//* Exit
			default:
				fmt.Println("Exiting...")
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

// * Render View * //
func (m mainModel) View() string {
	//* HEADER
	result := defaultColor
	result += "Galaxy Conquest\n\n"

	//* Print Choices
	for i, choice := range m.choices {

		var cursor string
		if m.cursor == i {
			cursor = selectedColor + "    "
		} else {
			cursor = "   "
		}

		result += fmt.Sprintf("%s %s%s\n", cursor, choice, defaultColor)

	}

	//* FOOTER
	result += "\nfooter\n"

	return result
}
