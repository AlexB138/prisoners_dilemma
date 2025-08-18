package event

import (
	"fmt"

	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
)

// An Event is a single contest between two strategies
type Event struct {
	totalRounds  int
	currentRound int
	result       round.History
	participant1 *participant
	participant2 *participant
}

type participant struct {
	strategy strategies.Strategy
	score    action.Score
}

func (p *participant) updateScore(data *round.Data) {
	p.score += data.Score
	data.RunningScore = p.score
}

func NewEvent(rounds int, strategy1, strategy2 strategies.Strategy) *Event {
	p1 := &participant{
		strategy: strategy1,
		score:    0,
	}

	p2 := &participant{
		strategy: strategy2,
		score:    0,
	}

	return &Event{
		totalRounds:  rounds,
		currentRound: 0,
		result:       make(round.History),
		participant1: p1,
		participant2: p2,
	}
}

func (e *Event) GetParticipantNames() (string, string) {
	return e.participant1.strategy.Name(), e.participant2.strategy.Name()
}

func (e *Event) GetScore() (action.Score, action.Score) {
	return e.participant1.score, e.participant2.score
}

func (e *Event) Run() {
	for e.currentRound < e.totalRounds {
		e.currentRound++

		r := e.executeRound(e.currentRound)

		e.result[e.currentRound] = r

		// Update scores
		e.participant1.updateScore(r.Participant1Data)
		e.participant2.updateScore(r.Participant2Data)

		// Send results to strategies
		e.participant1.strategy.ReceiveResult(e.currentRound, round.Participant1, *r)
		e.participant2.strategy.ReceiveResult(e.currentRound, round.Participant2, *r)
	}
}

func (e *Event) executeRound(roundNum int) *round.Round {
	d1 := &round.Data{}
	d2 := &round.Data{}

	d1.Action = e.participant1.strategy.MakeChoice(roundNum)
	d2.Action = e.participant2.strategy.MakeChoice(roundNum)

	d1.Score, d2.Score = action.ScoreActions(d1.Action, d2.Action)

	return &round.Round{
		RoundNum:         roundNum,
		Participant1Data: d1,
		Participant2Data: d2,
	}
}

func (e *Event) String() string {
	var n1, n2, output string

	if e.participant1 != nil && e.participant1.strategy != nil {
		n1 = e.participant1.strategy.Name()
	} else {
		return ""
	}

	if e.participant2 != nil && e.participant2.strategy != nil {
		n2 = e.participant2.strategy.Name()
	}

	output = fmt.Sprintf("\t%s\t\t\t\t%s\n", n1, n2)
	output += "Round\tAction\tScore\tTotal\t\tAction\tScore\tTotal\n"

	for i := 1; i <= e.totalRounds; i++ {
		r := e.result[i]

		p1d := r.Participant1Data
		p2d := r.Participant2Data

		output += fmt.Sprintf("%d\t", i)
		output += fmt.Sprintf("%s\t%d\t%d\t\t", p1d.Action, p1d.Score, p1d.RunningScore)
		output += fmt.Sprintf("%s\t%d\t%d\n", p2d.Action, p2d.Score, p2d.RunningScore)
	}

	output += fmt.Sprintf("\nFinal:\t%d\t\t\t\t%d\n", e.participant1.score, e.participant2.score)

	return output
}

func (e *Event) Winner() strategies.Strategy {
	if e.participant1.score == e.participant2.score {
		return nil
	} else if e.participant1.score > e.participant2.score {
		return e.participant1.strategy
	} else {
		return e.participant2.strategy
	}
}

// GetParticipant returns the participant for the specified participant number
func (e *Event) GetParticipant(participantNum round.Participant) (*participant, bool) {
	switch participantNum {
	case round.Participant1:
		return e.participant1, true
	case round.Participant2:
		return e.participant2, true
	default:
		return nil, false
	}
}

// GetParticipantStrategy returns the strategy for the specified participant number
func (e *Event) GetParticipantStrategy(participantNum round.Participant) (strategies.Strategy, bool) {
	if p, ok := e.GetParticipant(participantNum); ok {
		return p.strategy, true
	}
	return nil, false
}

// GetParticipantScore returns the score for the specified participant number
func (e *Event) GetParticipantScore(participantNum round.Participant) (action.Score, bool) {
	if p, ok := e.GetParticipant(participantNum); ok {
		return p.score, true
	}
	return 0, false
}
