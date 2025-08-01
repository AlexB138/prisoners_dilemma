package simulation

import (
	"github.com/AlexB138/prisoners_dilemma/internal/event"
)

// Simulation is a container for one or more event.Event.
type Simulation struct {
	events   []event.Event
	settings Settings
}

func NewSimulation(settings Settings) *Simulation {
	return &Simulation{
		events:   make([]event.Event, 0),
		settings: settings,
	}
}

func (s *Simulation) Run() {
	for i := 0; i < s.settings.Iterations(); i++ {

		e := event.NewEvent(15, s1, s2)
		e.Run()
	}
}
