package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

type Cooperator struct {
	name string
}

func NewCooperator() Strategy {
	return &Cooperator{name: "Cooperator"}
}

func (c *Cooperator) GetName() string {
	return c.name
}

func (c *Cooperator) MakeChoice(_ int) action.Action {
	return action.Cooperate
}

func (c *Cooperator) ReceiveResult(_, _ int, _ *round.Round) {
	return
}
