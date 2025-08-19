package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type GradualTitForTat struct {
	apologizing        bool
	apologyStreak      int
	history            round.History
	name               string
	opponentDefections int
	participantNum     round.Participant
	punishing          bool
	punishmentStreak   int
}

func init() { Register(NewGradualTitForTat) }

func NewGradualTitForTat() Strategy {
	return &GradualTitForTat{name: "GradualTitForTat"}
}

func (s *GradualTitForTat) Description() string {
	return "TFT with two differences. First, it increases the string of punishing defection responses with each additional defection by its opponent. Second, it apologizes for each string of defections by cooperating in the subsequent two rounds."
}

func (s *GradualTitForTat) Name() string {
	return s.name
}

func (s *GradualTitForTat) ParticipantNumber() round.Participant {
	return s.participantNum
}

func (s *GradualTitForTat) MakeChoice(roundNum int) action.Action {
	opPreviousAction, ok := s.opponentsPreviousMove(roundNum)
	if !ok {
		// If no previous opponent action is present, cooperate.
		// This always happens the first round.
		return action.Cooperate
	}

	/*
		If opponent defected, I increment s.opponentDefections and initiate a punishment streak.
		If on a punishment streak, I defect until I have defected s.opponentDefections times.
		If I complete a punishment streak, I start an apology.
		If I am apologizing, I participate until s.apologyStreak is 2.
		If opponent participates and I am not on any streak, I participate.
	*/
	s.updateState(opPreviousAction)

	// Note that apologizing supersedes punishing
	if s.apologizing {
		s.apologyStreak++
		return action.Cooperate
	}

	if s.punishing {
		s.punishmentStreak++
		return action.Defect
	}

	return action.Cooperate
}

func (s *GradualTitForTat) ReceiveResult(roundNum int, participantNum round.Participant, r round.Round) {
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

func (s *GradualTitForTat) Reset() {
	s.history = make(round.History)
	s.apologyStreak = 0
	s.opponentDefections = 0
	s.punishmentStreak = 0
	s.apologizing = false
	s.punishing = false
	s.participantNum = round.ParticipantNone
}

// updateState handles the management of apology and punishment
func (s *GradualTitForTat) updateState(opPreviousAction action.Action) {
	if opPreviousAction == action.Defect {
		s.opponentDefections++
		s.punishing = true
	}

	// Complete a punishment streak, begin apology
	if s.punishing && s.punishmentStreak == s.opponentDefections {
		s.punishing = false
		s.punishmentStreak = 0
		s.apologizing = true
	}

	// Complete apology
	if s.apologizing && s.apologyStreak == 2 {
		s.apologizing = false
		s.apologyStreak = 0
	}
}

func (s *GradualTitForTat) opponentsPreviousMove(roundNum int) (action.Action, bool) {
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
