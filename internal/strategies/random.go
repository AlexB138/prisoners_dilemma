package strategies

import (
	"math/rand"

	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type Random struct {
	name string
}

func init() { Register(NewRandom) }

func NewRandom() Strategy {
	return &Random{name: "Random"}
}

func (r *Random) Description() string {
	return "Randomly cooperates or defects with equal likelihood."
}

func (r *Random) Name() string {
	return r.name
}

func (r *Random) MakeChoice(_ int) action.Action {
	if rand.Int()%2 == 0 {
		return action.Cooperate
	} else {
		return action.Defect
	}
}

func (r *Random) ReceiveResult(_, _ int, _ round.Round) {}

func (r *Random) Reset() {}
