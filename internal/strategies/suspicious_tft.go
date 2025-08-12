package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type SuspiciousTitForTat struct {
	history        round.History
	name           string
	participantNum int
}

func init() { Register(NewSuspiciousTitForTat) }

func NewSuspiciousTitForTat() Strategy {
	return &SuspiciousTitForTat{name: "SuspiciousTitForTat"}
}

func (t *SuspiciousTitForTat) Description() string {
	return "Defect on the first round and imitates its opponent's previous move thereafter."
}

func (t *SuspiciousTitForTat) Name() string {
	return t.name
}

func (t *SuspiciousTitForTat) MakeChoice(roundNum int) action.Action {
	opPreviousAction, ok := t.getOpponentsPreviousMove(roundNum)
	if !ok {
		// If no previous opponent action is present, defect.
		// This always happens in the first round.
		return action.Defect
	}

	// If they previously defected, defect. Vice versa.
	return opPreviousAction
}

func (t *SuspiciousTitForTat) ReceiveResult(roundNum, participantNum int, r round.Round) {
	if t == nil {
		return
	}

	if t.history == nil {
		t.history = make(round.History)
	}

	if t.participantNum == 0 {
		t.participantNum = participantNum
	}

	t.history[roundNum] = &r
}

func (t *SuspiciousTitForTat) Reset() {
	t.history = make(round.History)
}

func (t *SuspiciousTitForTat) getOpponentsPreviousMove(roundNum int) (action.Action, bool) {
	if t == nil || t.history == nil {
		return action.Defect, false
	}

	if r, ok := t.history[roundNum-1]; ok {
		opponentData := r.Participant1Data
		if t.participantNum == 1 {
			opponentData = r.Participant2Data
		}

		return opponentData.Action, true
	} else {
		return action.Defect, false
	}
}
