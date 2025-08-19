package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type GrimTrigger struct {
	name           string
	participantNum round.Participant
	triggered      bool
}

func init() { Register(NewGrimTrigger) }

func NewGrimTrigger() Strategy {
	return &GrimTrigger{name: "GrimTrigger"}
}

func (s *GrimTrigger) Description() string {
	return "Cooperates until its opponent defects, then defects forever."
}

func (s *GrimTrigger) Name() string {
	return s.name
}

func (s *GrimTrigger) ParticipantNumber() round.Participant {
	return s.participantNum
}

// MakeChoice returns action.Cooperate if the strategy has not been triggered, and action.Defect otherwise.
func (s *GrimTrigger) MakeChoice(roundNum int) action.Action {
	if s.triggered {
		return action.Defect
	}

	return action.Cooperate
}

func (s *GrimTrigger) ReceiveResult(_ int, participantNum round.Participant, r round.Round) {
	if s == nil {
		return
	}

	if a, ok := r.OpponentAction(s.ParticipantNumber()); ok && a == action.Defect {
		s.triggered = true
	}

	if s.participantNum == round.ParticipantNone {
		s.participantNum = participantNum
	}
}

func (s *GrimTrigger) Reset() {
	s.participantNum = round.ParticipantNone
	s.triggered = false
}
