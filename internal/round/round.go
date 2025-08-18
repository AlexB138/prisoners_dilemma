package round

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
)

type Round struct {
	RoundNum         int
	Participant1Data *Data
	Participant2Data *Data
}

// getDataForParticipant is a helper function that returns the data for a given participant
func (r Round) getDataForParticipant(participantNum Participant) (*Data, bool) {
	switch participantNum {
	case Participant1:
		return r.Participant1Data, true
	case Participant2:
		return r.Participant2Data, true
	default:
		return nil, false
	}
}

// OpponentAction returns the action taken by the opponent in this Round. Note that participantNum is the requesting strategy number, not the opponent, for ease of use.
func (r Round) OpponentAction(participantNum Participant) (action.Action, bool) {
	if data, ok := r.getDataForParticipant(participantNum.Opponent()); ok {
		return data.Action, true
	}
	return action.Defect, false
}

// GetParticipantData returns the data for the specified participant
func (r Round) GetParticipantData(participantNum Participant) (*Data, bool) {
	return r.getDataForParticipant(participantNum)
}

// GetOpponentData returns the data for the opponent of the specified participant
func (r Round) GetOpponentData(participantNum Participant) (*Data, bool) {
	return r.getDataForParticipant(participantNum.Opponent())
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
