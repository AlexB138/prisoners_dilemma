package simulation

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/event"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
)

// Simulation is a container for one or more event.Event.
type Simulation struct {
	events   []event.Event
	settings Settings
	winner   strategies.Strategy
}

func NewSimulation(settings Settings) *Simulation {
	return &Simulation{
		events:   make([]event.Event, 0),
		settings: settings,
		winner:   nil,
	}
}

func (s *Simulation) Run() {
	for i := 0; i < s.settings.Iterations; i++ {
		s.resetStrategies()

		e := event.NewEvent(s.settings.Rounds, s.settings.Strategy1, s.settings.Strategy2)
		e.Run()
		s.events = append(s.events, *e)
	}

	var winner strategies.Strategy

	if s.settings.SettingType == SingleEvent {
		winner = s.events[0].Winner()
	} else if s.settings.SettingType == BestOfN {
		winner = s.bestOfNWinner()
	}

	s.winner = winner
}

func (s *Simulation) resetStrategies() {
	s.settings.Strategy1.Reset()
	s.settings.Strategy2.Reset()
}

// GetEvents returns all events in the simulation
func (s *Simulation) GetEvents() []event.Event {
	return s.events
}

func (s *Simulation) GetFinalScores() (action.Score, action.Score) {
	return s.events[len(s.events)-1].GetScore()
}

func (s *Simulation) GetParticipantNames() (string, string) {
	var n1, n2 string

	if s.settings.Strategy1 != nil {
		n1 = s.settings.Strategy1.GetName()
	}

	if s.settings.Strategy2 != nil {
		n2 = s.settings.Strategy2.GetName()
	}

	return n1, n2
}

// GetWinner returns the winner of the simulation, nil after running indicates a tie
func (s *Simulation) GetWinner() strategies.Strategy {
	return s.winner
}

func (s *Simulation) Reset() {
	s.events = make([]event.Event, 0)
	s.winner = nil
	s.resetStrategies()
}
