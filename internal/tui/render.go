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

	winner := a.sim.Winner()

	header := fmt.Sprintf(`
Settings:
  Rounds    : %d
  Type      : %s`, a.settings.Rounds, a.settings.Type)

	var iterativeDetails string
	if a.settings.Type == simulation.BestOfN {
		iterativeDetails = fmt.Sprintf(`
  Iterations: %d
  Scoring   : %s`, a.settings.Iterations, a.settings.IterativeGameType)
	}

	results := a.renderResultsBlock()

	var winnerText string
	if winner == nil {
		winnerText = "Tie!"
	} else {
		winnerText = winner.GetName()
	}
	footer := fmt.Sprintf(`

  Winner: %s

Press 'r' to run again, 'q' to quit`, winnerText)

	// Combine all sections
	content := header + iterativeDetails + results + footer

	return boxStyle.Render(content)
}

func (a *App) renderResultsBlock() string {
	name1, name2 := a.sim.ParticipantNames()
	var s1, s2 string

	if a.settings.Type == simulation.SingleEvent {
		s1, s2 = a.renderSingleEventScore()
	} else {
		switch a.settings.IterativeGameType {
		case simulation.IterativeGameTypeHighestSingleEvent:
			s1, s2 = a.renderHighestSingleEventScore()

		case simulation.IterativeGameTypeHighestTotal:
			s1, s2 = a.renderHighestTotalScore()

		case simulation.IterativeGameTypeBestAverageScore:
			s1, s2 = a.renderBestAverageScore()

		case simulation.IterativeGameTypeMostWins:
			s1, s2 = a.renderMostWinsScore()
		}
	}

	l := len(name1)
	if len(name2) > len(name1) {
		l = len(name2)
	}

	return fmt.Sprintf(`

Results:
  %-*s: %s
  %-*s: %s`, l, name1, s1, l, name2, s2)
}

func (a *App) renderSingleEventScore() (string, string) {
	score1, score2 := a.sim.SingleEventScore()
	return fmt.Sprintf("%d points", score1), fmt.Sprintf("%d points", score2)
}

func (a *App) renderHighestSingleEventScore() (string, string) {
	highScore1, highScore2 := a.sim.HighestSingleEventScore()
	return fmt.Sprintf("%d high score", highScore1), fmt.Sprintf("%d high score", highScore2)
}

func (a *App) renderHighestTotalScore() (string, string) {
	total1, total2 := a.sim.HighestTotalScore()
	return fmt.Sprintf("%d total score", total1), fmt.Sprintf("%d total score", total2)
}

func (a *App) renderBestAverageScore() (string, string) {
	avg1, avg2 := a.sim.BestAverageScore()
	return fmt.Sprintf("%.2f average score", avg1), fmt.Sprintf("%.2f average score", avg2)
}

func (a *App) renderMostWinsScore() (string, string) {
	wins1, wins2 := a.sim.MostWinsScore()
	return fmt.Sprintf("%d wins", wins1), fmt.Sprintf("%d wins", wins2)
}
