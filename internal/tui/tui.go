package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
)

// App represents the main TUI application
type App struct {
	state    appState
	sim      *simulation.Simulation
	settings *simulation.Settings
	// helpOpen indicates whether contextual help is currently displayed
	helpOpen bool
	// helpIndex is the 1-based index of the currently highlighted item for help
	helpIndex int
}

// simulationCompleteMsg is sent when simulation completes
type simulationCompleteMsg struct {
	sim *simulation.Simulation
}

// NewApp creates a new TUI application
func NewApp() *App {
	return &App{
		state: stateStrategy1,
		settings: &simulation.Settings{
			Rounds:     10,
			Iterations: 1,
		},
		helpOpen:  false,
		helpIndex: 1,
	}
}

// Init initializes the TUI application
func (a *App) Init() tea.Cmd {
	return nil
}

// Update handles user input and state transitions
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m := msg.(type) {
	case tea.KeyMsg:
		if handler, ok := stateToHandler[a.state]; ok {
			return handler(a, m)
		}
	case simulationCompleteMsg:
		a.sim = m.sim
		a.state = stateResults
		return a, nil
	}

	return a, nil
}

// View renders the current state of the TUI
func (a *App) View() string {
	if render, ok := stateToRender[a.state]; ok {
		return render(a)
	} else {
		return "Unknown state"
	}
}

// Run starts the TUI application
func Run() error {
	p := tea.NewProgram(NewApp())
	_, err := p.Run()
	return err
}

func (a *App) runSimulation() tea.Cmd {
	return func() tea.Msg {
		sim := simulation.NewSimulation(*a.settings)
		sim.Run()

		return simulationCompleteMsg{sim: sim}
	}
}
