package conquestmode

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luanrivello/conquest/board"
)

// * Model * //
type loadingModel struct {
	state    int
	start    bool
	loading  bool
	board    board.Board
	previous tea.Model
	err      error
}

// * Constructor * //
func initLoading(prev tea.Model) loadingModel {
	return loadingModel{
		start:    true,
		loading:  true,
		previous: prev,
	}
}

// * Get Model * //
func GetLoadingModel(prev tea.Model) loadingModel {
	return initLoading(prev)
}

// * Startup * //
func (m loadingModel) Init() tea.Cmd {
	createBoard(&m)
	return ticktack()
}

// * Actions * //
func (m loadingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		//* Keypress
		switch msg.String() {
		case "ctrl+c", "q":
			return m.previous, nil
		}

	case TickMsg:
		if m.start {
			m.start = false
			createBoard(&m)
		} else if !m.loading {
			nextModel := GetConquestModel(&m.board)
			return nextModel, nextModel.Init()
		}

		m.updateDots()
		return m, ticktack()

	}

	return m, nil
}

// * Render View * //
func (m loadingModel) View() string {
	//* Header
	result := defaultColor
	result += "Conquest Mode\n\n"

	//* Loading
	if m.loading {
		result += "    Loading"
		switch m.state {
		case 0:
			break
		case 1:
			result += "."
		case 2:
			result += ".."
		case 3:
			result += "..."
		}
		result += "\n"

	} else {
		result += "    Completed\n"
	}

	//* Footer
	result += "\nfooter\n"

	return result
}

func (m *loadingModel) updateDots() {
	if m.state == 3 {
		m.state = 0
	} else {
		m.state++
	}
}

type TickMsg time.Time

func ticktack() tea.Cmd {
	return tea.Tick(1*time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func createBoard(m *loadingModel) tea.Msg {
	m.board = *board.NewBoard()
	m.loading = false
	return nil
}
