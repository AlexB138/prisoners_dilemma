package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type Strategy interface {
	// GetName returns the Strategy name
	GetName() string
	// MakeChoice returns the action.Action for the round
	MakeChoice(roundNum int) action.Action
	// ReceiveResult sends the round number, an indicator of which participant the strategy is, and
	// simulation.Round containing results back to the Strategy,
	ReceiveResult(roundNum, participantNum int, round *round.Round)
}
