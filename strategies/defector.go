package strategies

import (
	"github.com/AlexB138/prisoners_dilemma/action"
	"github.com/AlexB138/prisoners_dilemma/round"
)

type Defector struct {
	name string
}

func NewDefector() Strategy {
	return &Defector{name: "Defector"}
}

func (d *Defector) GetName() string {
	return d.name
}

func (d *Defector) MakeChoice(_ int) action.Action {
	return action.Defect
}

func (d *Defector) ReceiveResult(_, _ int, _ *round.Round) {
	return
}
