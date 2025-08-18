package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type Defector struct {
	name           string
	participantNum round.Participant
}

func init() { Register(NewDefector) }

func NewDefector() Strategy {
	return &Defector{name: "Defector"}
}

func (s *Defector) Description() string {
	return "Defects unconditionally."
}

func (s *Defector) Name() string {
	return s.name
}

func (s *Defector) MakeChoice(_ int) action.Action {
	return action.Defect
}

func (s *Defector) ParticipantNumber() round.Participant {
	return s.participantNum
}

func (s *Defector) ReceiveResult(roundNum int, participantNum round.Participant, _ round.Round) {
	if s.participantNum == round.ParticipantNone {
		s.participantNum = participantNum
	}
}

func (s *Defector) Reset() {
	s.participantNum = round.ParticipantNone
}
