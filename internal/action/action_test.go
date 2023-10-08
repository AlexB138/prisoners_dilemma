package action

import "testing"

func TestScoreActions(t *testing.T) {
	data := []struct {
		name           string
		action1        Action
		action2        Action
		expectedScore1 Score
		expectedScore2 Score
	}{
		{"Both Cooperate", Cooperate, Cooperate, Good, Good},
		{"Both Defect", Defect, Defect, Bad, Bad},
		{"Cooperate Defect", Cooperate, Defect, Minimum, Maximum},
		{"Defect Cooperate", Defect, Cooperate, Maximum, Minimum},
	}

	t.Parallel()

	for _, d := range data {
		d := d

		t.Run(d.name, func(t *testing.T) {
			score1, score2 := ScoreActions(d.action1, d.action2)
			if score1 != d.expectedScore1 || score2 != d.expectedScore2 {
				t.Errorf("Expected score %d and %d, got %d and %d", d.expectedScore1, d.expectedScore2, score1, score2)
			}
		})
	}
}
