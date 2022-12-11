package conquestmode

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luanrivello/conquest/loop"
	"github.com/luanrivello/conquest/spacetime"
	"github.com/luanrivello/conquest/tui/colors"
)

// * Get Model * //
func GetConquestModel(prev tea.Model, gal *spacetime.Galaxy) conquestModel {
	return initConquest(prev, gal)
}

// * Model * //
type conquestModel struct {
	galaxy   *spacetime.Galaxy
	choices  []string
	cursor   int
	loading  bool
	typing   bool
	previous tea.Model
}

// * Constructor * //
func initConquest(prev tea.Model, gal *spacetime.Galaxy) conquestModel {
	return conquestModel{
		galaxy:   gal,
		choices:  []string{"Galaxy Name", "Planet Name", "Next", "Back"},
		previous: prev,
	}
}

// * Startup * //
func (m conquestModel) Init() tea.Cmd {
	return nil
}

// * Actions * //
func (m conquestModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			//Choises
			switch m.cursor {
			case 0:
				loop.Gameloop()
			case 1:
				loop.Gameloop()
			case 2:
				loop.Gameloop()
			default:
				return m.previous, nil
			}

		}

	}

	return m, nil

}

// * Render View * //
func (m conquestModel) View() string {
	//* Header
	system := m.galaxy.GetSystem()
	result := defaultColor
	result += "Conquest Mode\n"
	result += colors.BLUE + "✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦\n"
	result += colors.YELLOW + "Date:   " + "0" + "\n"
	result += colors.PURPLE + "System: " + system.GetName() + "\n"
	result += colors.RED + "Sun:    " + system.GetSun().GetName() + "\n"
	result += colors.GREEN + "Planet: " + system.GetPlanet().GetName() + "\n"

	//* Print Choises
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
