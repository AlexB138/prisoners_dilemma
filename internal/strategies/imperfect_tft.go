package strategies

import (
	"math/rand"

	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type ImperfectTitForTat struct {
	history        round.History
	name           string
	participantNum round.Participant
}

func init() {
	Register(NewImperfectTitForTat)
}

func NewImperfectTitForTat() Strategy {
	return &ImperfectTitForTat{name: "ImperfectTitForTat"}
}

func (s *ImperfectTitForTat) Description() string {
	return "Imitates opponent's last move with high (but less than certain) probability."
}

func (s *ImperfectTitForTat) Name() string {
	return s.name
}

func (s *ImperfectTitForTat) ParticipantNumber() round.Participant {
	return s.participantNum
}

// MakeChoice returns the opponent's last action.Action with 70% probability, and action.Cooperate or action.Defect with 15% probability each.'
func (s *ImperfectTitForTat) MakeChoice(roundNum int) action.Action {
	opPreviousAction, ok := s.getOpponentsPreviousMove(roundNum)
	if !ok {
		// If no previous opponent action is present, cooperate.
		// This always happens the first round.
		return action.Cooperate
	}

	probability := rand.Float64()
	if probability < 0.7 {
		return opPreviousAction
	} else if probability < 0.85 {
		return action.Cooperate
	}
	return action.Defect
}

func (s *ImperfectTitForTat) ReceiveResult(roundNum int, participantNum round.Participant, r round.Round) {
	if s == nil {
		return
	}

	if s.history == nil {
		s.history = make(round.History)
	}

	if s.participantNum == round.ParticipantNone {
		s.participantNum = participantNum
	}

	s.history[roundNum] = &r
}

func (s *ImperfectTitForTat) Reset() {
	s.history = make(round.History)
	s.participantNum = round.ParticipantNone
}

func (s *ImperfectTitForTat) getOpponentsPreviousMove(roundNum int) (action.Action, bool) {
	if s == nil || s.history == nil {
		return action.Cooperate, false
	}

	if r, ok := s.history[roundNum-1]; ok {
		opponentData, ok := r.GetOpponentData(s.participantNum)
		if !ok {
			return action.Cooperate, false
		}
		return opponentData.Action, true
	} else {
		return action.Cooperate, false
	}
}
