package strategies

import (
	"math"
	"math/rand/v2"

	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type GenerousTitForTat struct {
	history        round.History
	name           string
	participantNum int
}

func init() { Register(NewGenerousTitForTat) }

func NewGenerousTitForTat() Strategy {
	return &GenerousTitForTat{name: "GenerousTitForTat"}
}

func (g *GenerousTitForTat) Description() string {
	return "Cooperates on the first round and after its opponent cooperates. Following a defection, it cooperates with probability g(R,P,T,S)=min{1 - (T - R)/(R - S), (R - P)/(T - P)}, where R, P, T and S are the reward, punishment, temptation and sucker payoffs."
}

func (g *GenerousTitForTat) Name() string {
	return g.name
}

func (g *GenerousTitForTat) MakeChoice(roundNum int) action.Action {
	opPreviousAction, ok := g.getOpponentsPreviousMove(roundNum)
	if !ok {
		// If no previous opponent action is present, cooperate.
		// This always happens in the first round.
		return action.Cooperate
	}

	// If they previously cooperated, cooperate.
	if opPreviousAction == action.Cooperate {
		return action.Cooperate
	}

	return g.chooseDefectResponse()
}

func (g *GenerousTitForTat) ReceiveResult(roundNum, participantNum int, r round.Round) {
	if g == nil {
		return
	}

	if g.history == nil {
		g.history = make(round.History)
	}

	if g.participantNum == 0 {
		g.participantNum = participantNum
	}

	g.history[roundNum] = &r
}

func (g *GenerousTitForTat) Reset() {
	g.history = make(round.History)
}

func (g *GenerousTitForTat) getOpponentsPreviousMove(roundNum int) (action.Action, bool) {
	if g == nil || g.history == nil {
		return action.Cooperate, false
	}

	if r, ok := g.history[roundNum-1]; ok {
		opponentData := r.Participant1Data
		if g.participantNum == 1 {
			opponentData = r.Participant2Data
		}

		return opponentData.Action, true
	} else {
		return action.Cooperate, false
	}
}

func (g *GenerousTitForTat) chooseDefectResponse() action.Action {
	/*
		T = Temptation (you defect, opponent cooperates — best personal payoff)
		R = Reward (mutual cooperation — good but not maximal payoff)
		P = Punishment (mutual defection — worse than cooperation)
		S = Sucker’s payoff (you cooperate, opponent defects — worst for you)
	*/
	t := float64(action.Maximum)
	r := float64(action.Good)
	p := float64(action.Bad)
	s := float64(action.Minimum)

	/*
		temptation: T - R / R - S. Compares the temptation to defect to the loss from being exploited.
		T - R: The extra gain you get from exploiting someone instead of mutually cooperating.
		R - S: The “cost” to you of being exploited (cooperating while the other defects).

		Subtracting temptation from 1 gives a forgiveness factor: if temptation is high relative to sucker loss, you
			forgive less.
	*/
	temptation := 1.0 - (t-r)/(r-s)

	/*
		benevolence: R - P / T - P. This ratio reflects how much you value cooperation over mutual defection compared to
			the lure of defection — higher values mean more incentive to forgive.
		R - P: The benefit of mutual cooperation over mutual defection.
		T - P: The benefit of exploiting someone over mutual defection.
	*/
	benevolence := (r - p) / (t - p)

	// The strategy takes the more conservative (lower) probability from the two formulas, ensuring it doesn’t become
	// too forgiving in scenarios where that would be exploited.
	prob := math.Min(temptation, benevolence)
	if prob < 0 {
		prob = 0
	}
	if prob > 1 {
		prob = 1
	}

	if rand.Float64() <= prob {
		return action.Cooperate
	}

	return action.Defect
}
