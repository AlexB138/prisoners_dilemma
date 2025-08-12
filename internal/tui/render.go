package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
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
	content := a.renderStrategySelectionContent("Select Strategy 1:", false)
	return boxStyle.Render(content)
}

func (a *App) renderStrategy2Selection() string {
	header := fmt.Sprintf("Strategy 1: %s\n\n", a.settings.Strategy1.Name())
	body := a.renderStrategySelectionContent("Select Strategy 2:", true)
	return boxStyle.Render(header + body)
}

func (a *App) renderRoundsInput() string {
	content := fmt.Sprintf(`
Strategy 1: %s
Strategy 2: %s

Number of Rounds: %d

Use ↑/↓ to adjust, Enter to continue, b to go back
`, a.settings.Strategy1.Name(), a.settings.Strategy2.Name(), a.settings.Rounds)
	return boxStyle.Render(content)
}

func (a *App) renderSimTypeSelection() string {
	header := fmt.Sprintf("Strategy 1: %s\nStrategy 2: %s\nRounds: %d\n\nSimulation Type:\n\n", a.settings.Strategy1.Name(), a.settings.Strategy2.Name(), a.settings.Rounds)
	options := []string{
		fmt.Sprintf("1. %s", simulation.SingleEvent),
		fmt.Sprintf("2. %s", simulation.BestOfN),
	}
	content := header + renderHelpList(options, a.helpOpen, a.helpIndex) + "\n" + renderSimTypeHelp(a)
	content += "\nPress 1-2 to select, h for help, b to go back, q to quit"
	return boxStyle.Render(content)
}

func (a *App) renderIterativeTypeInput() string {
	header := fmt.Sprintf("Strategy 1: %s\nStrategy 2: %s\nRounds: %d\nSim Type: %s\n\nSelect Iterative Scoring Method:\n\n",
		a.settings.Strategy1.Name(),
		a.settings.Strategy2.Name(),
		a.settings.Rounds,
		a.settings.Type,
	)
	options := []string{
		fmt.Sprintf("1. %s", simulation.IterativeGameTypeMostWins),
		fmt.Sprintf("2. %s", simulation.IterativeGameTypeHighestTotal),
		fmt.Sprintf("3. %s", simulation.IterativeGameTypeHighestSingleEvent),
		fmt.Sprintf("4. %s", simulation.IterativeGameTypeBestAverageScore),
	}
	content := header + renderHelpList(options, a.helpOpen, a.helpIndex) + "\n" + renderIterativeTypeHelp(a)
	content += "\nPress 1-4 to select, h for help, b to go back, q to quit"
	return boxStyle.Render(content)
}

func (a *App) renderIterationsInput() string {
	content := fmt.Sprintf(`
Strategy 1: %s
Strategy 2: %s
Rounds: %d
Sim Type: %s
Scoring Method: %s

Number of Events: %d

Use ↑/↓ to adjust, Enter to continue, b to go back
`,
		a.settings.Strategy1.Name(),
		a.settings.Strategy2.Name(),
		a.settings.Rounds,
		a.settings.Type,
		a.settings.IterativeGameType,
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
		winnerText = winner.Name()
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

// --- Shared help rendering utilities ---

func renderHelpList(options []string, helpOpen bool, helpIndex int) string {
	// Render a simple list with a caret for the help-highlighted item if help is open
	result := ""
	for idx, opt := range options {
		prefix := "  "
		if helpOpen && (idx+1) == helpIndex {
			prefix = "> "
		}
		result += fmt.Sprintf("%s%s\n", prefix, opt)
	}
	return result
}

func renderSimTypeHelp(a *App) string {
	if !a.helpOpen {
		return ""
	}
	selected := a.helpIndex
	var key simulation.Type
	if selected == 1 {
		key = simulation.SingleEvent
	} else {
		key = simulation.BestOfN
	}
	return fmt.Sprintf("\n[Help: %s]\n%s\n", key, simulation.HelpForType(key))
}

func renderIterativeTypeHelp(a *App) string {
	if !a.helpOpen {
		return ""
	}
	var key simulation.IterativeGameType
	switch a.helpIndex {
	case 1:
		key = simulation.IterativeGameTypeMostWins
	case 2:
		key = simulation.IterativeGameTypeHighestTotal
	case 3:
		key = simulation.IterativeGameTypeHighestSingleEvent
	case 4:
		key = simulation.IterativeGameTypeBestAverageScore
	}
	return fmt.Sprintf("\n[Help: %s]\n%s\n", key, simulation.HelpForIterativeType(key))
}

func (a *App) renderStrategySelectionContent(title string, includeBack bool) string {
	options := strategies.Discover()

	list := title + "\n\n"
	for i, s := range options {
		line := fmt.Sprintf("%d. %s", i+1, s.Name())
		if a.helpOpen && a.helpIndex == i+1 {
			list += "> " + line + "\n"
		} else {
			list += "  " + line + "\n"
		}
	}

	// Instructions reflect dynamic option count; we support 1-9 via number keys or Enter on highlighted when help is open
	if includeBack {
		list += "\nPress 1-9 (or Enter when highlighted) to select, h for help, b to go back, q to quit\n"
	} else {
		list += "\nPress 1-9 (or Enter when highlighted) to select, h for help, q to quit\n"
	}

	if a.helpOpen {
		// Show help for the highlighted strategy
		idx := a.helpIndex - 1
		if idx < 0 {
			idx = 0
		}
		if idx > len(options)-1 {
			idx = len(options) - 1
		}
		selected := options[idx]
		list += fmt.Sprintf("\n[Help: %s]\n%s\n", selected.Name(), selected.Description())
	}

	return list
}
