package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
)

type appState int

const (
	stateStrategy1 appState = iota
	stateStrategy2
	stateRounds
	stateIterations
	stateSimType
	stateRunning
	stateResults
)

type handlerFunc func(*App, tea.KeyMsg) (tea.Model, tea.Cmd)

var stateToHandler = map[appState]handlerFunc{
	stateStrategy1:  (*App).handleStrategy1Selection,
	stateStrategy2:  (*App).handleStrategy2Selection,
	stateRounds:     (*App).handleRoundsInput,
	stateSimType:    (*App).handleSimTypeSelection,
	stateIterations: (*App).handleIterationsInput,
	stateRunning:    (*App).handleResultsView,
	stateResults:    (*App).handleResultsView,
}

func (a *App) handleStrategy1Selection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "1":
		a.strategy1 = strategies.NewCooperator()
		a.state = stateStrategy2
	case "2":
		a.strategy1 = strategies.NewDefector()
		a.state = stateStrategy2
	case "3":
		a.strategy1 = strategies.NewRandom()
		a.state = stateStrategy2
	case "4":
		a.strategy1 = strategies.NewTitForTat()
		a.state = stateStrategy2
	case "q", "ctrl+c":
		return a, tea.Quit
	}
	return a, nil
}

func (a *App) handleStrategy2Selection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "1":
		a.strategy2 = strategies.NewCooperator()
		a.state = stateRounds
	case "2":
		a.strategy2 = strategies.NewDefector()
		a.state = stateRounds
	case "3":
		a.strategy2 = strategies.NewRandom()
		a.state = stateRounds
	case "4":
		a.strategy2 = strategies.NewTitForTat()
		a.state = stateRounds
	case "b":
		a.state = stateStrategy1
	case "q", "ctrl+c":
		return a, tea.Quit
	}
	return a, nil
}

func (a *App) handleRoundsInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if a.rounds == 1 {
			a.rounds = 5
		} else if a.rounds < 100 {
			a.rounds += 5
		}
	case "down":
		if a.rounds > 5 {
			a.rounds -= 5
		} else if a.rounds == 5 {
			a.rounds = 1
		}
	case "enter":
		a.state = stateSimType
	case "q", "ctrl+c":
		return a, tea.Quit
	case "b":
		a.state = stateStrategy2
	}

	return a, nil
}

func (a *App) handleSimTypeSelection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "1":
		a.simType = simulation.SingleEvent
		a.iterations = 1
		a.state = stateRunning
		return a, a.runSimulation()
	case "2":
		a.simType = simulation.BestOfN
		a.state = stateIterations
	case "b":
		a.state = stateRounds
	case "q", "ctrl+c":
		return a, tea.Quit
	}
	return a, nil
}

func (a *App) handleIterationsInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if a.iterations < 99 {
			a.iterations += 2
		}
	case "down":
		if a.iterations > 1 {
			a.iterations -= 2
		}
	case "enter":
		a.state = stateSimType
	case "q", "ctrl+c":
		return a, tea.Quit
	case "b":
		a.state = stateSimType
	}

	return a, nil
}

func (a *App) handleResultsView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "r":
		a.sim.Reset()
		a.state = stateRunning
		return a, a.runSimulation()
	case "q", "ctrl+c":
		return a, tea.Quit
	}
	return a, nil
}
