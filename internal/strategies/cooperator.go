package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type Cooperator struct {
	name           string
	participantNum round.Participant
}

func init() { Register(NewCooperator) }

func NewCooperator() Strategy {
	return &Cooperator{name: "Cooperator"}
}

func (s *Cooperator) Description() string {
	return "Cooperates unconditionally."
}

func (s *Cooperator) Name() string {
	return s.name
}

func (s *Cooperator) MakeChoice(_ int) action.Action {
	return action.Cooperate
}

func (s *Cooperator) ParticipantNumber() round.Participant {
	return s.participantNum
}

func (s *Cooperator) ReceiveResult(roundNum int, participantNum round.Participant, _ round.Round) {
	if s.participantNum == round.ParticipantNone {
		s.participantNum = participantNum
	}
}

func (s *Cooperator) Reset() {
	s.participantNum = round.ParticipantNone
}
