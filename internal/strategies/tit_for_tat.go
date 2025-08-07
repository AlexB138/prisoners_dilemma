package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type TitForTat struct {
	history        round.History
	name           string
	participantNum int
}

func NewTitForTat() Strategy {
	return &TitForTat{name: "TitForTat"}
}

func (c *TitForTat) GetName() string {
	return c.name
}

func (c *TitForTat) MakeChoice(roundNum int) action.Action {
	opPreviousAction, ok := c.getOpponentsPreviousMove(roundNum)
	if !ok {
		// If no previous opponent action is present, cooperate.
		// This always happens the first round.
		return action.Cooperate
	}

	// If they previously defected, defect. Vice versa.
	return opPreviousAction
}

func (c *TitForTat) ReceiveResult(roundNum, participantNum int, r round.Round) {
	if c == nil {
		return
	}

	if c.history == nil {
		c.history = make(round.History)
	}

	if c.participantNum == 0 {
		c.participantNum = participantNum
	}

	c.history[roundNum] = &r
}

func (c *TitForTat) Reset() {
	c.history = make(round.History)
}

func (c *TitForTat) getOpponentsPreviousMove(roundNum int) (action.Action, bool) {
	if c == nil || c.history == nil {
		return action.Cooperate, false
	}

	if r, ok := c.history[roundNum-1]; ok {
		opponentData := r.Participant1Data
		if c.participantNum == 1 {
			opponentData = r.Participant2Data
		}

		return opponentData.Action, true
	} else {
		return action.Cooperate, false
	}
}
