package tui

import (
	"fmt"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
)

type renderFunc func(*App) string

var stateToRender = map[appState]renderFunc{
	stateStrategy1:     (*App).renderStrategy1Selection,
	stateStrategy2:     (*App).renderStrategy2Selection,
	stateRounds:        (*App).renderRoundsInput,
	stateSimType:       (*App).renderSimTypeSelection,
	stateIterativeType: (*App).renderIterativeTypeInput,
	stateIterations:    (*App).renderIterationsInput,
	stateRunning:       (*App).renderRunning,
	stateResults:       (*App).renderResults,
}

func (a *App) renderStrategy1Selection() string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                     Prisoner's Dilemma                       ║
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
║                     Prisoner's Dilemma                       ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Strategy 1: %-15s 							       ║
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
║                     Prisoner's Dilemma                       ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Strategy 1: %-45s 										   ║
║  Strategy 2: %-45s 										   ║
║                                                              ║
║  Number of Rounds: %-4d                                      ║
║                                                              ║
║  Use ↑/↓ to adjust, Enter to continue, b to go back          ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`, a.settings.Strategy1.GetName(), a.settings.Strategy2.GetName(), a.settings.Rounds)
}

func (a *App) renderSimTypeSelection() string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                     Prisoner's Dilemma                       ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Strategy 1: %-45s 										   ║
║  Strategy 2: %-45s 										   ║
║  Rounds: %-4d                                                ║
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

func (a *App) renderIterativeTypeInput() string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                     Prisoner's Dilemma                       ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Strategy 1: %-45s 										   ║
║  Strategy 2: %-45s 										   ║
║  Rounds: %-3d                                                ║
║  Sim Type: %-3s                                              ║
║                                                              ║
║  Select Iterative Scoring Method:                            ║
║                                                              ║
║  1. %-20s                                                    ║
║  2. %-45s 										           ║
║  3. %-45s 										           ║
║  4. %-45s 										           ║
║                                                              ║
║  Press 1-4 to select, b to go back, q to quit                ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`,
		a.settings.Strategy1.GetName(),
		a.settings.Strategy2.GetName(),
		a.settings.Rounds,
		a.settings.Type,
		simulation.IterativeGameTypeMostWins,
		simulation.IterativeGameTypeHighestTotal,
		simulation.IterativeGameTypeHighestSingleEvent,
		simulation.IterativeGameTypeBestAverageScore,
	)
}

func (a *App) renderIterationsInput() string {
	return fmt.Sprintf(`
╔══════════════════════════════════════════════════════════════╗
║                     Prisoner's Dilemma                       ║
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
║                     Prisoner's Dilemma                       ║
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
║                     Prisoner's Dilemma                       ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  Settings:                                                   ║
║      Rounds: %-3d                                             ║
║      Type:   %-15s                                 ║
║                                                              ║
║                                                              ║
║  Results:                                                    ║
║      %-12s: %-4d points                               ║
║      %-12s: %-4d points                               ║
║                                                              ║
║      Winner: %-15s                                 ║
║                                                              ║
║  Press 'r' to run again, 'q' to quit                         ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`,
		a.settings.Rounds,
		a.settings.Type,
		name1,
		score1,
		name2,
		score2,
		w,
	)

	return result
}
