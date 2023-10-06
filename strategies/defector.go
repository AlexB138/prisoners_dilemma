package strategies

import "github.com/AlexB138/prisoners_dilemma/action"

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

func (d *Defector) ReceiveResult(_ int, _ action.Score, _ action.Action) {
	return
}
