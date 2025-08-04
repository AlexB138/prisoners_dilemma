package tui

type appState int

const (
	stateStrategy1 appState = iota
	stateStrategy2
	stateRounds
	stateIterativeType
	stateIterations
	stateSimType
	stateRunning
	stateResults
)

type transition struct {
	next appState
	prev appState
}

var stateTransitions = map[appState]transition{
	stateStrategy1:     {next: stateStrategy2},
	stateStrategy2:     {prev: stateStrategy1, next: stateRounds},
	stateRounds:        {prev: stateStrategy2, next: stateSimType},
	stateSimType:       {prev: stateRounds}, // intentionally no next, controlled by handler
	stateIterativeType: {prev: stateSimType, next: stateIterations},
	stateIterations:    {prev: stateIterativeType, next: stateRunning},
	stateRunning:       {prev: stateSimType}, // intentionally no next
	stateResults:       {prev: stateSimType}, // intentionally no next
}

func (a *App) transitionTo(newState appState) {
	a.state = newState
}

func (a *App) nextState() {
	if t, ok := stateTransitions[a.state]; ok {
		a.transitionTo(t.next)
	}
}

func (a *App) previousState() {
	if t, ok := stateTransitions[a.state]; ok {
		a.transitionTo(t.prev)
	}
}
