package strategies

import (
	"math/rand"

	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type Random struct {
	name           string
	participantNum round.Participant
}

func init() { Register(NewRandom) }

func NewRandom() Strategy {
	return &Random{name: "Random"}
}

func (s *Random) Description() string {
	return "Randomly cooperates or defects with equal likelihood."
}

func (s *Random) Name() string {
	return s.name
}

func (s *Random) MakeChoice(_ int) action.Action {
	if rand.Int()%2 == 0 {
		return action.Cooperate
	} else {
		return action.Defect
	}
}

func (s *Random) ParticipantNumber() round.Participant {
	return s.participantNum
}

func (s *Random) ReceiveResult(roundNum int, participantNum round.Participant, _ round.Round) {
	if s.participantNum == round.ParticipantNone {
		s.participantNum = participantNum
	}
}

func (s *Random) Reset() {
	s.participantNum = round.ParticipantNone
}
