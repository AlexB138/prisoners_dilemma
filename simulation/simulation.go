package simulation

import (
	"fmt"
	"github.com/AlexB138/prisoners_dilemma/action"
	"github.com/AlexB138/prisoners_dilemma/strategies"
)

type Simulation struct {
	totalRounds  int
	currentRound int
	Result       roundHistory
	participant1 *participant
	participant2 *participant
}

type participant struct {
	strategy strategies.Strategy
	score    action.Score
}

type roundHistory map[int]*round

type round struct {
	participant1Data roundData
	participant2Data roundData
}

// roundData stores information about given round for a single participant
type roundData struct {
	// action is the action.Action taken by the participant in this round
	action action.Action
	// score is the score awarded for the action taken in this round
	score action.Score
	// runningScore is the total score for the participant at the end of this round
	runningScore action.Score
}

func NewSimulation(rounds int, strategy1, strategy2 strategies.Strategy) *Simulation {
	p1 := &participant{
		strategy: strategy1,
		score:    0,
	}

	p2 := &participant{
		strategy: strategy2,
		score:    0,
	}

	return &Simulation{
		totalRounds:  rounds,
		currentRound: 0,
		Result:       make(roundHistory),
		participant1: p1,
		participant2: p2,
	}
}

func (s *Simulation) Run() {
	for s.currentRound < s.totalRounds {
		s.currentRound++

		r := s.executeRound(s.currentRound)

		s.Result[s.currentRound] = r

		// Update scores
		s.participant1.score += r.participant1Data.score
		r.participant1Data.runningScore = s.participant1.score
		s.participant2.score += r.participant2Data.score
		r.participant2Data.runningScore = s.participant2.score

		// Send results to strategies
		s.participant1.strategy.ReceiveResult(s.currentRound, r.participant1Data.score, r.participant2Data.action)
		s.participant2.strategy.ReceiveResult(s.currentRound, r.participant2Data.score, r.participant1Data.action)
	}
}

func (s *Simulation) executeRound(roundNum int) *round {
	var d1, d2 roundData

	d1.action = s.participant1.strategy.MakeChoice(roundNum)
	d2.action = s.participant2.strategy.MakeChoice(roundNum)

	d1.score, d2.score = action.ScoreActions(d1.action, d2.action)

	return &round{
		participant1Data: d1,
		participant2Data: d2,
	}
}

func (s *Simulation) String() string {
	var n1, n2, output string

	if s.participant1 != nil && s.participant1.strategy != nil {
		n1 = s.participant1.strategy.GetName()
	} else {
		return ""
	}

	if s.participant2 != nil && s.participant2.strategy != nil {
		n2 = s.participant2.strategy.GetName()
	}

	output = fmt.Sprintf("\t%s\t\t\t%s\n", n1, n2)
	output += "Round\tAction\tScore\tTotal\t\tAction\tScore\tTotal\n"

	for i := 1; i <= s.totalRounds; i++ {
		r := s.Result[i]

		p1d := r.participant1Data
		p2d := r.participant2Data

		output += fmt.Sprintf("%d\t", i)
		output += fmt.Sprintf("%s\t%d\t%d\t\t", p1d.action, p1d.score, p1d.runningScore)
		output += fmt.Sprintf("%s\t%d\t%d\n", p2d.action, p2d.score, p2d.runningScore)
	}

	output += fmt.Sprintf("\nFinal:\t%d\t\t\t\t%d\n", s.participant1.score, s.participant2.score)

	return output
}
