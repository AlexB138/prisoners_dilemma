//go:build testonly

package round

import "github.com/AlexB138/prisoners_dilemma/internal/action"

// MakeTestRound creates a Round with the supplied actions. For testing purposes.
func MakeTestRound(p1action, p2action action.Action) *Round {
	return &Round{
		RoundNum: 0,
		Participant1Data: &Data{
			Action:       p1action,
			Score:        0,
			RunningScore: 0,
		},
		Participant2Data: &Data{
			Action:       p2action,
			Score:        0,
			RunningScore: 0,
		},
	}
}

// MakeTestHistory creates a History. For testing purposes.
func MakeTestHistory(rounds ...*Round) History {
	h := make(History)

	for i, r := range rounds {
		h[i+1] = r
	}

	return h
}
