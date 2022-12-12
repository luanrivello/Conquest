package conquestmode

import (
	tea "github.com/charmbracelet/bubbletea"
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

		case "enter", " ", "l":
			//Choices

		}

	}

	return m, nil

}

// * Render View * //
func (m conquestModel) View() string {
	//* Header
	system := m.galaxy.GetSystem()
	sun := system.GetSun()
	planet := system.GetPlanet()
	result := defaultColor
	result += colors.BLUE + "✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦\n"
	result += colors.YELLOW + "Date:   " + "0" + "\n"
	result += colors.PURPLE + "System: " + system.GetName() + "\n"
	result += colors.RED + "Sun:    " + sun.GetName() + "\n"
	result += colors.GREEN + "Planet: " + planet.GetName() + "\n"
	result += colors.BLUE + "✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦\n"
	result += colors.Reset

	result += colors.CYAN + ">\n"
	result += colors.RED + "↧\n"
	result += colors.GREEN + "O\n\n"
	result += colors.Reset

	tiles := planet.GetTiles()

	for _, tileLine := range tiles {

		for _, tile := range tileLine {

			var aux = ""
			if len(tile.String()) != 0 {
				result += colors.GREEN
			}

			aux += "["

			//stack := tile.String()
			//aux += stack
			stack := ""

			// Fill aux string untill it has lenght 9
			for len(aux) < 1+len(stack)/3 {
				aux += " "
			}

			aux += "]"
			result += aux + colors.Reset
		}
		result += "\n"
	}

	result += "\n"

	//* Footer
	result += "\n| footer |\n"
	return result
}
