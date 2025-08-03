package simulation

import "github.com/AlexB138/prisoners_dilemma/internal/strategies"

type Settings struct {
	IterativeGameType    IterativeGameType
	Iterations           int
	Rounds               int
	SettingType          Type
	Strategy1, Strategy2 strategies.Strategy
}
