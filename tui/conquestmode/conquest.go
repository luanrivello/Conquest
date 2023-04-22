package conquestmode

import (
	"conquest/board"
	"conquest/tui/colors"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

// * Get Model * //
func GetConquestModel(board *board.Board) conquestModel {
	return initConquest(board)
}

// * Model * //
type conquestModel struct {
	board    *board.Board
	choices  []string
	cursor   int
	loading  bool
	typing   bool
	previous tea.Model
	err      error
}

// * Constructor * //
func initConquest(board *board.Board) conquestModel {
	return conquestModel{
		board: board,
	}
}

// * Startup * //
func (m conquestModel) Init() tea.Cmd {
	go m.board.Run()
	return ticktack()
}

// * Actions * //
func (m conquestModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		//* Keypress
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ", "l":

		}

	case TickMsg:
		return m, ticktack()

	}

	return m, nil
}

// * Render View * //
func (m conquestModel) View() string {
	//* Header
	galaxy := m.board.GetGalaxy()
	system := galaxy.GetSystem()
	sun := system.GetSun()
	planet := system.GetPlanet()
	result := defaultColor
	result += colors.BLUE + "✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦\n"
	result += colors.YELLOW + "Date:   " + strconv.Itoa(m.board.Calendar) + "\n"
	result += colors.PURPLE + "System: " + system.GetName() + "\n"
	result += colors.RED + "Sun:    " + sun.GetName() + "\n"
	result += colors.GREEN + "Planet: " + planet.GetName() + "\n"
	result += colors.BLUE + "✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦✦\n"
	result += colors.Reset

	result += colors.CYAN + ">\n"
	result += colors.RED + "↧\n"
	result += colors.GREEN + "O\n\n"
	result += colors.Reset

	// * Body
	tiles := planet.GetTiles()
	for _, tileLine := range tiles {

		for _, tile := range tileLine {

			var aux = ""
			aux += "["
			if len(tile.String()) != 0 {
				result += colors.GREEN
				aux += tile.String()
			} else {
				result += colors.Reset
				aux += "   "
			}

			//stack := tile.String()
			//aux += stack
			//stack := ""

			// Fill aux string untill it has lenght 9
			//for len(aux) < 1+len(stack)/3 {
			//	aux += " "
			//}

			aux += "]"
			aux += colors.Reset
			result += aux
		}
		result += "\n"
	}

	result += "\n"
	result += colors.Reset

	//* Footer
	result += "---------------------------------------------------------------------------------\n"
	result += colors.RED
	log := m.board.AkashicRecord.Actions
	limit := 0
	for i := len(log) - 1; i >= 0; i-- {
		if limit == 5 {
			break
		}
		result += log[i].String() + "\n"
		limit++
	}

	return result
}
