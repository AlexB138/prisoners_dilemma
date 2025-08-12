package tui

type appState int

const (
	stateNone appState = iota
	stateStrategy1
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

// stateTransitions tracks the previous and next states of each state. states without values either don't transition
// or are handled in their handler logic
var stateTransitions = map[appState]transition{
	stateStrategy1:     {next: stateStrategy2},
	stateStrategy2:     {prev: stateStrategy1, next: stateRounds},
	stateRounds:        {prev: stateStrategy2, next: stateSimType},
	stateSimType:       {prev: stateRounds},
	stateIterativeType: {prev: stateSimType, next: stateIterations},
	stateIterations:    {prev: stateIterativeType},
	stateRunning:       {prev: stateSimType},
	stateResults:       {prev: stateSimType},
}

func (a *App) transitionTo(newState appState) {
	a.state = newState
	// Reset help UI on state change
	a.helpOpen = false
	a.helpIndex = 1
}

func (a *App) nextState() {
	if t, ok := stateTransitions[a.state]; ok {
		a.transitionTo(t.next)
	}
}

func (a *App) previousState() {
	if t, ok := stateTransitions[a.state]; ok {
		if t.prev == stateNone {
			return
		}
		a.transitionTo(t.prev)
	}
}
