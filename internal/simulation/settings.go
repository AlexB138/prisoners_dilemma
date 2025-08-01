package simulation

import "github.com/AlexB138/prisoners_dilemma/internal/strategies"

type Settings interface {
	Iterations() int
	Rounds() int
	Strategies() (strategies.Strategy, strategies.Strategy)
	Type() Type
}

type BaseSettings struct {
	settingType          Type
	strategy1, strategy2 strategies.Strategy
	rounds               int
}

type SingleEventSettings struct {
	BaseSettings
}

func (bs *BaseSettings) Type() Type {
	return bs.settingType
}

func (bs *BaseSettings) Rounds() int {
	return bs.rounds
}

func (bs *BaseSettings) Iterations() int {
	return 1
}

type BestOfNSettings struct {
	BaseSettings
	iterations int
}

func (ns *BestOfNSettings) Iterations() int {
	return ns.iterations
}
