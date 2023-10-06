package strategies

import "github.com/AlexB138/prisoners_dilemma/action"

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

func (c *Cooperator) ReceiveResult(_ int, _ action.Score, _ action.Action) {
	return
}
