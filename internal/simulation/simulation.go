package simulation

import (
	"fmt"
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
)

type Simulation struct {
	totalRounds  int
	currentRound int
	Result       round.History
	participant1 *participant
	participant2 *participant
}

type participant struct {
	strategy strategies.Strategy
	score    action.Score
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
		Result:       make(round.History),
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
		s.participant1.score += r.Participant1Data.Score
		r.Participant1Data.RunningScore = s.participant1.score
		s.participant2.score += r.Participant2Data.Score
		r.Participant2Data.RunningScore = s.participant2.score

		// Send results to strategies
		s.participant1.strategy.ReceiveResult(s.currentRound, 1, r)
		s.participant2.strategy.ReceiveResult(s.currentRound, 2, r)
	}
}

func (s *Simulation) executeRound(roundNum int) *round.Round {
	var d1, d2 round.Data

	d1.Action = s.participant1.strategy.MakeChoice(roundNum)
	d2.Action = s.participant2.strategy.MakeChoice(roundNum)

	d1.Score, d2.Score = action.ScoreActions(d1.Action, d2.Action)

	return &round.Round{
		Participant1Data: d1,
		Participant2Data: d2,
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

	output = fmt.Sprintf("\t%s\t\t\t\t%s\n", n1, n2)
	output += "Round\tAction\tScore\tTotal\t\tAction\tScore\tTotal\n"

	for i := 1; i <= s.totalRounds; i++ {
		r := s.Result[i]

		p1d := r.Participant1Data
		p2d := r.Participant2Data

		output += fmt.Sprintf("%d\t", i)
		output += fmt.Sprintf("%s\t%d\t%d\t\t", p1d.Action, p1d.Score, p1d.RunningScore)
		output += fmt.Sprintf("%s\t%d\t%d\n", p2d.Action, p2d.Score, p2d.RunningScore)
	}

	output += fmt.Sprintf("\nFinal:\t%d\t\t\t\t%d\n", s.participant1.score, s.participant2.score)

	return output
}
