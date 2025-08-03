package round

import "github.com/AlexB138/prisoners_dilemma/internal/action"

type Round struct {
	RoundNum         int
	Participant1Data *Data
	Participant2Data *Data
}

// Data stores information about a given Round for a single participant
type Data struct {
	// Action is the action.Action taken by the participant in this Round
	Action action.Action
	// Score is the Score awarded for the Action taken in this Round
	Score action.Score
	// RunningScore is the total Score for the participant at the end of this Round
	RunningScore action.Score
}

type History map[int]*Round
