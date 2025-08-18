package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type Strategy interface {
	// Description returns a description of the strategy. Used for Help text.
	Description() string
	// Name returns the Strategy name
	Name() string
	// MakeChoice returns the action.Action for the round
	MakeChoice(roundNum int) action.Action
	// ParticipantNumber returns the participant number of an active or finished strategy. Used for helpers.
	ParticipantNumber() round.Participant
	// ReceiveResult sends the round number, an indicator of which participant the Strategy is, and
	// round.Round containing results back to the Strategy
	ReceiveResult(roundNum int, participantNum round.Participant, round round.Round)
	// Reset reinitializes a strategy. This allows it to participate in multiple events.
	Reset()
}
