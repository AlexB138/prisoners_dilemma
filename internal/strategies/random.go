package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/round"
	"math/rand"

	"github.com/AlexB138/prisoners_dilemma/internal/action"
)

type Random struct {
	name string
}

func NewRandom() Strategy {
	return &Random{name: "Random"}
}

func (d *Random) GetName() string {
	return d.name
}

func (d *Random) MakeChoice(_ int) action.Action {
	if rand.Int()%2 == 0 {
		return action.Cooperate
	} else {
		return action.Defect
	}
}

func (d *Random) ReceiveResult(_, _ int, _ *round.Round) {
	return
}
