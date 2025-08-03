package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
)

// App represents the main TUI application
type App struct {
	state         appState
	previousState appState
	strategy1     strategies.Strategy
	strategy2     strategies.Strategy
	rounds        int
	iterations    int
	simType       simulation.Type
	sim           *simulation.Simulation
	settings      *simulation.Settings
}

// simulationCompleteMsg is sent when simulation completes
type simulationCompleteMsg struct {
	sim *simulation.Simulation
}

// NewApp creates a new TUI application
func NewApp() *App {
	return &App{
		state:   stateStrategy1,
		rounds:  10, // default
		simType: simulation.SingleEvent,
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

func (a *App) runSimulation() tea.Cmd {
	return func() tea.Msg {
		settings := simulation.Settings{
			IterativeGameType: "",
			Iterations:        a.iterations,
			Rounds:            a.rounds,
			SettingType:       a.simType,
			Strategy1:         a.strategy1,
			Strategy2:         a.strategy2,
		}

		sim := simulation.NewSimulation(settings)
		sim.Run()

		return simulationCompleteMsg{sim: sim}
	}
}

// Run starts the TUI application
func Run() error {
	p := tea.NewProgram(NewApp())
	_, err := p.Run()
	return err
}
