package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
)

var boxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63")).
	Padding(1, 2).
	Width(60).
	Height(15)

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
	content := `
Select Strategy 1:

  1. Cooperator (Always Cooperate)
  2. Defector (Always Defect)
  3. Random (Random Choice)
  4. Tit for Tat (Copy Opponent's Last Move)

Press 1-4 to select, q to quit
`
	return boxStyle.Render(content)
}

func (a *App) renderStrategy2Selection() string {
	content := fmt.Sprintf(`
Strategy 1: %s

Select Strategy 2:

  1. Cooperator (Always Cooperate)
  2. Defector (Always Defect)
  3. Random (Random Choice)
  4. Tit for Tat (Copy Opponent's Last Move)

Press 1-4 to select, b to go back, q to quit
`, a.settings.Strategy1.GetName())
	return boxStyle.Render(content)
}

func (a *App) renderRoundsInput() string {
	content := fmt.Sprintf(`
Strategy 1: %s
Strategy 2: %s

Number of Rounds: %d

Use ↑/↓ to adjust, Enter to continue, b to go back
`, a.settings.Strategy1.GetName(), a.settings.Strategy2.GetName(), a.settings.Rounds)
	return boxStyle.Render(content)
}

func (a *App) renderSimTypeSelection() string {
	content := fmt.Sprintf(`
Strategy 1: %s
Strategy 2: %s
Rounds: %d

Simulation Type:

  1. Single Event (One match)
  2. Best of N (Multiple matches, best wins)

Press 1-2 to select, b to go back, q to quit
`, a.settings.Strategy1.GetName(), a.settings.Strategy2.GetName(), a.settings.Rounds)
	return boxStyle.Render(content)
}

func (a *App) renderIterativeTypeInput() string {
	content := fmt.Sprintf(`
Strategy 1: %s
Strategy 2: %s
Rounds: %d
Sim Type: %s

Select Iterative Scoring Method:

  1. %s
  2. %s
  3. %s
  4. %s

Press 1-4 to select, b to go back, q to quit
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
	return boxStyle.Render(content)
}

func (a *App) renderIterationsInput() string {
	content := fmt.Sprintf(`
Strategy 1: %s
Strategy 2: %s
Rounds: %d
Sim Type: %s

Number of Events: %d

Use ↑/↓ to adjust, Enter to continue, b to go back
`,
		a.settings.Strategy1.GetName(),
		a.settings.Strategy2.GetName(),
		a.settings.Rounds,
		a.settings.Type,
		a.settings.Iterations,
	)
	return boxStyle.Render(content)
}

func (a *App) renderRunning() string {
	content := `

                    Running Simulation...

                    Please wait...

`
	return boxStyle.Render(content)
}

func (a *App) renderResults() string {
	if a.sim == nil {
		return boxStyle.Render("No sim available")
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

	content := fmt.Sprintf(`
Settings:
  Rounds: %d
  Type:   %s

Results:
  %s: %d points
  %s: %d points

  Winner: %s

Press 'r' to run again, 'q' to quit
`,
		a.settings.Rounds,
		a.settings.Type,
		name1,
		score1,
		name2,
		score2,
		w,
	)

	return boxStyle.Render(content)
}
