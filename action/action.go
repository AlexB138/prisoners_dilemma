package action

type Action string

const (
	Defect    Action = "Defect"
	Cooperate Action = "Cooperate"
)

func ScoreActions(action1, action2 Action) (int, int) {
	if action1 == Cooperate {
		if action2 == Cooperate {
			// C, C
			return 3, 3
		} else {
			// C, D
			return 0, 5
		}
	} else {
		if action2 == Cooperate {
			// D, C
			return 5, 0
		} else {
			// D, D
			return 1, 1
		}
	}
}
