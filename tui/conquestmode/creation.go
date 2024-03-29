package conquestmode

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"conquest/tui/colors"
)

var defaultColor = colors.BLUE

// * Get Model * //
func GetCreationModel(prev tea.Model) creationModel {
	return initCreation(prev)
}

// * Model * //
type creationModel struct {
	choices  []string
	cursor   int
	loading  bool
	typing   bool
	previous tea.Model
	err      error
}

// * Constructor * //
func initCreation(prev tea.Model) creationModel {
	return creationModel{
		choices:  []string{"Galaxy Name", "Planet Name", "Next", "Back"},
		previous: prev,
	}
}

// * Startup * //
func (m creationModel) Init() tea.Cmd {
	return ticktack()
}

// * Actions * //
func (m creationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Keypress
	case tea.KeyMsg:
		//Which Key
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
			case 0:
				//TODO: set name

			case 1:
				//TODO: set name

			case 2:
				nextModel := GetLoadingModel(m.previous)
				return nextModel, nextModel.Init()

			default:
				return m.previous, nil
			}
		}

	}

	return m, nil

}

// * Render View * //
func (m creationModel) View() string {
	//* Header
	result := defaultColor
	result += "Conquest Mode\n\n"

	//* Print Choices
	for i, choice := range m.choices {

		var cursor string
		if m.cursor == i {
			cursor = colors.GREEN + "    "
		} else {
			cursor = "   "
		}

		result += fmt.Sprintf("%s %s%s\n", cursor, choice, defaultColor)

	}

	//* Footer
	result += "\nfooter\n"

	return result
}
