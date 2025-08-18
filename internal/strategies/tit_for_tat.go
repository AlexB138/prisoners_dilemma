package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type TitForTat struct {
	history        round.History
	name           string
	participantNum round.Participant
}

func init() { Register(NewTitForTat) }

func NewTitForTat() Strategy {
	return &TitForTat{name: "TitForTat"}
}

func (t *TitForTat) Description() string {
	return "Cooperates on the first round and imitates its opponent's previous move thereafter."
}

func (t *TitForTat) Name() string {
	return t.name
}

func (t *TitForTat) ParticipantNumber() round.Participant {
	return t.participantNum
}

func (t *TitForTat) MakeChoice(roundNum int) action.Action {
	opPreviousAction, ok := t.getOpponentsPreviousMove(roundNum)
	if !ok {
		// If no previous opponent action is present, cooperate.
		// This always happens the first round.
		return action.Cooperate
	}

	// If they previously defected, defect. Vice versa.
	return opPreviousAction
}

func (t *TitForTat) ReceiveResult(roundNum int, participantNum round.Participant, r round.Round) {
	if t == nil {
		return
	}

	if t.history == nil {
		t.history = make(round.History)
	}

	if t.participantNum == round.ParticipantNone {
		t.participantNum = participantNum
	}

	t.history[roundNum] = &r
}

func (t *TitForTat) Reset() {
	t.history = make(round.History)
	t.participantNum = round.ParticipantNone
}

func (t *TitForTat) getOpponentsPreviousMove(roundNum int) (action.Action, bool) {
	if t == nil || t.history == nil {
		return action.Cooperate, false
	}

	if r, ok := t.history[roundNum-1]; ok {
		opponentData, ok := r.GetOpponentData(t.participantNum)
		if !ok {
			return action.Cooperate, false
		}
		return opponentData.Action, true
	} else {
		return action.Cooperate, false
	}
}
