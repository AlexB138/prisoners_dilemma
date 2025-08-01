package action

type (
	Action string
	Score  int
)

const (
	Defect    Action = "Defect"
	Cooperate Action = "Coop"
	Maximum   Score  = 5
	Good      Score  = 3
	Bad       Score  = 1
	Minimum   Score  = 0
)

// ScoreActions returns the scores for the input actions. Score1, Score2.
func ScoreActions(action1, action2 Action) (Score, Score) {
	if action1 == Cooperate {
		if action2 == Cooperate {
			// C, C
			return Good, Good
		} else {
			// C, D
			return Minimum, Maximum
		}
	} else {
		if action2 == Cooperate {
			// D, C
			return Maximum, Minimum
		} else {
			// D, D
			return Bad, Bad
		}
	}
}
