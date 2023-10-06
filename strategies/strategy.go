package strategies

import "github.com/AlexB138/prisoners_dilemma/action"

type Strategy interface {
	// GetName returns the Strategy name
	GetName() string
	// MakeChoice returns the action.Action for the round
	MakeChoice(round int) action.Action
	// ReceiveResult sends the result of the round to the Strategy, including its score for that round and the
	// action.Action taken by the opponent
	ReceiveResult(round int, score action.Score, opponentAction action.Action)
}
