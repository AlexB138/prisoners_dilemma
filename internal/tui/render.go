package tui

import "fmt"

type renderFunc func(*App) string

var stateToRender = map[appState]renderFunc{
	stateStrategy1:  (*App).renderStrategy1Selection,
	stateStrategy2:  (*App).renderStrategy2Selection,
	stateRounds:     (*App).renderRoundsInput,
	stateSimType:    (*App).renderSimTypeSelection,
	stateIterations: (*App).renderIterationsInput,
	stateRunning:    (*App).renderRunning,
	stateResults:    (*App).renderResults,
}

func (a *App) renderStrategy1Selection() string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                    Prisoner's Dilemma TUI                    ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Select Strategy 1:                                          ║
║                                                              ║
║  1. Cooperator (Always Cooperate)                            ║
║  2. Defector (Always Defect)                                 ║
║  3. Random (Random Choice)                                   ║
║  4. Tit for Tat (Copy Opponent's Last Move)                  ║
║                                                              ║
║  Press 1-4 to select, q to quit                              ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`)
}

func (a *App) renderStrategy2Selection() string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                    Prisoner's Dilemma TUI                    ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Strategy 1: %-45s 										   ║
║                                                              ║
║  Select Strategy 2:                                          ║
║                                                              ║
║  1. Cooperator (Always Cooperate)                            ║
║  2. Defector (Always Defect)                                 ║
║  3. Random (Random Choice)                                   ║
║  4. Tit for Tat (Copy Opponent's Last Move)                  ║
║                                                              ║
║  Press 1-4 to select, b to go back, q to quit                ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`, a.settings.Strategy1.GetName())
}

func (a *App) renderRoundsInput() string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                    Prisoner's Dilemma TUI                    ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Strategy 1: %-45s 										   ║
║  Strategy 2: %-45s 										   ║
║                                                              ║
║  Number of Rounds: %-3d                                      ║
║                                                              ║
║  Use ↑/↓ to adjust, Enter to continue, b to go back          ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`, a.settings.Strategy1.GetName(), a.settings.Strategy2.GetName(), a.settings.Rounds)
}

func (a *App) renderSimTypeSelection() string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                    Prisoner's Dilemma TUI                    ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Strategy 1: %-45s 										   ║
║  Strategy 2: %-45s 										   ║
║  Rounds: %-3d                                                ║
║                                                              ║
║  Simulation Type:                                            ║
║                                                              ║
║  1. Single Event (One match)                                 ║
║  2. Best of N (Multiple matches, best wins)                  ║
║                                                              ║
║  Press 1-2 to select, b to go back, q to quit                ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`, a.settings.Strategy1.GetName(), a.settings.Strategy2.GetName(), a.settings.Rounds)
}

func (a *App) renderIterationTypeInput() string {
	// TODO: Continue here
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                    Prisoner's Dilemma TUI                    ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Strategy 1: %-45s 										   ║
║  Strategy 2: %-45s 										   ║
║  Rounds: %-3d                                                ║
║  Sim Type: %-3s                                                ║
║                                                              ║
║  Number of Events: %-3d                                      ║
║                                                              ║
║  Use ↑/↓ to adjust, Enter to continue, b to go back          ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`,
		a.settings.Strategy1.GetName(),
		a.settings.Strategy2.GetName(),
		a.settings.Rounds,
		a.settings.Type,
		a.settings.Iterations,
	)
}

func (a *App) renderIterationsInput() string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                    Prisoner's Dilemma TUI                    ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Strategy 1: %-45s 										   ║
║  Strategy 2: %-45s 										   ║
║  Rounds: %-3d                                                ║
║  Sim Type: %-3s                                                ║
║                                                              ║
║  Number of Events: %-3d                                      ║
║                                                              ║
║  Use ↑/↓ to adjust, Enter to continue, b to go back          ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`,
		a.settings.Strategy1.GetName(),
		a.settings.Strategy2.GetName(),
		a.settings.Rounds,
		a.settings.Type,
		a.settings.Iterations,
	)
}

func (a *App) renderRunning() string {
	return `
╔══════════════════════════════════════════════════════════════╗
║                    Prisoner's Dilemma TUI                    ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║                        Running Simulation...                 ║
║                                                              ║
║                    Please wait...                            ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`
}

func (a *App) renderResults() string {
	if a.sim == nil {
		return "No sim available"
	}

	name1, name2 := a.sim.GetParticipantNames()
	score1, score2 := a.sim.GetFinalScores()
	winner := a.sim.GetWinner()

	var w string
	if winner == nil {
		w = "Tie!"
	} else {
		w = winner.GetName()
	}

	result := fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                    Prisoner's Dilemma TUI                    ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Results:                                                    ║
║                                                              ║
║  %s: %d points                                               ║
║  %s: %d points                                               ║
║                                                              ║
║  Winner: %s                                                  ║
║                                                              ║
║  Press 'r' to run again, 'q' to quit                         ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`, name1, score1, name2, score2, w)

	return result
}
