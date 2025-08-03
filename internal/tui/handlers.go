package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
)

type handlerFunc func(*App, tea.KeyMsg) (tea.Model, tea.Cmd)

var stateToHandler = map[appState]handlerFunc{
	stateStrategy1:  (*App).handleStrategySelection,
	stateStrategy2:  (*App).handleStrategySelection,
	stateRounds:     (*App).handleRoundsInput,
	stateSimType:    (*App).handleSimTypeSelection,
	stateIterations: (*App).handleIterationsInput,
	stateRunning:    (*App).handleResultsView,
	stateResults:    (*App).handleResultsView,
}

func (a *App) handleStrategySelection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var s strategies.Strategy

	switch msg.String() {
	case "1":
		s = strategies.NewCooperator()
	case "2":
		s = strategies.NewDefector()
	case "3":
		s = strategies.NewRandom()
	case "4":
		s = strategies.NewTitForTat()
	case "b":
		a.previousState()
	case "q", "ctrl+c":
		return a, tea.Quit
	}

	if a.state == stateStrategy1 {
		a.settings.Strategy1 = s
	} else if a.state == stateStrategy2 {
		a.settings.Strategy2 = s
	}

	if s != nil {
		a.nextState()
	}
	
	return a, nil
}

func (a *App) handleRoundsInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if a.settings.Rounds == 1 {
			a.settings.Rounds = 5
		} else if a.settings.Rounds < 100 {
			a.settings.Rounds += 5
		}
	case "down":
		if a.settings.Rounds > 5 {
			a.settings.Rounds -= 5
		} else if a.settings.Rounds == 5 {
			a.settings.Rounds = 1
		}
	case "enter":
		a.nextState()
	case "q", "ctrl+c":
		return a, tea.Quit
	case "b":
		a.previousState()
	}

	return a, nil
}

func (a *App) handleSimTypeSelection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "1":
		a.settings.Type = simulation.SingleEvent
		a.settings.Iterations = 1
		a.transitionTo(stateRunning)
		return a, a.runSimulation()
	case "2":
		a.settings.Type = simulation.BestOfN
		a.transitionTo(stateIterations)
	case "b":
		a.previousState()
	case "q", "ctrl+c":
		return a, tea.Quit
	}
	return a, nil
}

func (a *App) handleIterationsInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if a.settings.Iterations < 99 {
			a.settings.Iterations += 2
		}
	case "down":
		if a.settings.Iterations > 1 {
			a.settings.Iterations -= 2
		}
	case "enter":
		a.nextState()
	case "q", "ctrl+c":
		return a, tea.Quit
	case "b":
		a.previousState()
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

func (a *App) handlePanic(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	panic("Invalid transition specified")
	return a, nil
}
