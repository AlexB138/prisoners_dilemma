package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
)

type handlerFunc func(*App, tea.KeyMsg) (tea.Model, tea.Cmd)

var stateToHandler = map[appState]handlerFunc{
	stateStrategy1:     (*App).handleStrategySelection,
	stateStrategy2:     (*App).handleStrategySelection,
	stateRounds:        (*App).handleRoundsInput,
	stateSimType:       (*App).handleSimTypeSelection,
	stateIterativeType: (*App).handleIterativeTypeSelection,
	stateIterations:    (*App).handleIterationsInput,
	stateRunning:       (*App).handleRunningView,
	stateResults:       (*App).handleResultsView,
}

func (a *App) handleStrategySelection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var s strategies.Strategy
	optionCount := strategies.Count()

	switch msg.String() {
	case "h", "?":
		a.helpOpen = !a.helpOpen
	case "up":
		if a.helpOpen {
			if a.helpIndex > 1 {
				a.helpIndex--
			}
		}
	case "down":
		if a.helpOpen {
			if a.helpIndex < optionCount {
				a.helpIndex++
			}
		}
	case "1", "2", "3", "4", "5", "6", "7", "8", "9":
		idx := int(msg.Runes[0] - '1')
		if idx >= 0 && idx < optionCount {
			s = strategies.NewByIndex(idx)
		}
	case "enter":
		if a.helpOpen {
			idx := a.helpIndex - 1
			if idx >= 0 && idx < optionCount {
				s = strategies.NewByIndex(idx)
			}
		}
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
	case "h", "?":
		a.helpOpen = !a.helpOpen
	case "up":
		if a.helpOpen && a.settings.Type == "" {
			if a.helpIndex > 1 {
				a.helpIndex--
			}
		}
	case "down":
		if a.helpOpen && a.settings.Type == "" {
			if a.helpIndex < 2 {
				a.helpIndex++
			}
		}
	case "1":
		a.settings.Type = simulation.SingleEvent
		a.settings.Iterations = 1
		a.transitionTo(stateRunning)
		return a, a.runSimulation()
	case "2":
		a.settings.Type = simulation.BestOfN
		a.transitionTo(stateIterativeType)
	case "b":
		a.previousState()
	case "q", "ctrl+c":
		return a, tea.Quit
	}
	return a, nil
}

func (a *App) handleIterativeTypeSelection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	a.settings.IterativeGameType = simulation.IterativeGameTypeNone

	switch msg.String() {
	case "h", "?":
		a.helpOpen = !a.helpOpen
	case "up":
		if a.helpOpen {
			if a.helpIndex > 1 {
				a.helpIndex--
			}
		}
	case "down":
		if a.helpOpen {
			if a.helpIndex < 4 {
				a.helpIndex++
			}
		}
	case "1":
		a.settings.IterativeGameType = simulation.IterativeGameTypeMostWins
	case "2":
		a.settings.IterativeGameType = simulation.IterativeGameTypeHighestTotal
	case "3":
		a.settings.IterativeGameType = simulation.IterativeGameTypeHighestSingleEvent
	case "4":
		a.settings.IterativeGameType = simulation.IterativeGameTypeBestAverageScore
	case "b":
		a.previousState()
		return a, nil
	case "q", "ctrl+c":
		return a, tea.Quit
	}

	if a.settings.IterativeGameType != simulation.IterativeGameTypeNone {
		a.nextState()
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
		a.transitionTo(stateRunning)
		return a, a.runSimulation()
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

func (a *App) handleRunningView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return a, tea.Quit
	}
	return a, nil
}

func (a *App) handlePanic(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	panic("Invalid transition specified")
}
