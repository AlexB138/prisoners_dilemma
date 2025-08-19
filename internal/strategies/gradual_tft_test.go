package strategies

import (
	"testing"

	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

func TestMakeChoice_GradualTitForTat(t *testing.T) {
	opDefectRound := round.MakeTestRound(action.Cooperate, action.Defect)
	opDefectHistory := round.MakeTestHistory(opDefectRound)

	inputs := []struct {
		name        string
		roundNum    int
		expected    action.Action
		history     round.History
		apologizing bool
		punishing   bool
	}{
		{
			name:        "Always cooperate when apologizing",
			roundNum:    2,
			expected:    action.Cooperate,
			history:     opDefectHistory,
			apologizing: true,
			punishing:   false,
		},
		{
			name:        "Apology overrides punishment",
			roundNum:    2,
			expected:    action.Cooperate,
			history:     opDefectHistory,
			apologizing: true,
			punishing:   true,
		},
		{
			name:        "Always defect when punishing and not apologizing",
			roundNum:    2,
			expected:    action.Defect,
			history:     opDefectHistory,
			apologizing: false,
			punishing:   true,
		},
		{
			name:        "Cooperate by default",
			roundNum:    1,
			expected:    action.Cooperate,
			history:     nil,
			apologizing: false,
			punishing:   false,
		},
	}

	gtft := GradualTitForTat{
		apologizing:        false,
		apologyStreak:      0,
		history:            nil,
		name:               "",
		opponentDefections: 0,
		participantNum:     round.Participant1,
		punishing:          false,
		punishmentStreak:   0,
	}

	for _, input := range inputs {
		t.Run(input.name, func(t *testing.T) {
			gtft.Reset()
			gtft.participantNum = round.Participant1
			gtft.apologizing = input.apologizing
			gtft.punishing = input.punishing
			gtft.history = input.history
			received := gtft.MakeChoice(input.roundNum)

			if received != input.expected {
				t.Errorf("Received: %s, Expected: %s", received, input.expected)
			}
		})
	}
}

func Test_updateState_GradualTitForTat(t *testing.T) {
	inputs := []struct {
		name                       string
		opPreviousAction           action.Action
		apologyStreak              int
		opponentDefections         int
		punishmentStreak           int
		expectedApologyStreak      int
		expectedOpponentDefections int
		expectedPunishmentStreak   int
		apologizing                bool
		punishing                  bool
		expectedApologizing        bool
		expectedPunishing          bool
	}{
		{
			name:                       "Punish on defect",
			opPreviousAction:           action.Defect,
			opponentDefections:         0,
			expectedOpponentDefections: 1,
			expectedApologizing:        false,
			expectedPunishing:          true,
		},
		{
			name:                       "Complete punishment",
			opPreviousAction:           action.Cooperate,
			opponentDefections:         2,
			punishmentStreak:           2,
			expectedOpponentDefections: 2,
			expectedPunishmentStreak:   0,
			apologizing:                false,
			punishing:                  true,
			expectedApologizing:        true, // apologizing and punishment swap at completion
			expectedPunishing:          false,
		},
		{
			name:                  "Complete apology",
			opPreviousAction:      action.Cooperate,
			apologyStreak:         2,
			expectedApologyStreak: 0,
			apologizing:           true,
			expectedApologizing:   false,
		},
		{
			name:                       "Continue punishing on additional defect",
			opPreviousAction:           action.Defect,
			opponentDefections:         2,
			expectedOpponentDefections: 3, // An additional defect causes punishment to extend
			apologizing:                false,
			punishing:                  true,
			expectedApologizing:        false,
			expectedPunishing:          true,
		},
		{
			name:                       "Queue punishment while apologizing",
			opPreviousAction:           action.Defect,
			opponentDefections:         0,
			expectedOpponentDefections: 1,
			apologizing:                true,
			punishing:                  false,
			expectedApologizing:        true,
			expectedPunishing:          true, // Punishment is toggled on, even while apologizing
		},
	}

	gtft := GradualTitForTat{
		apologizing:        false,
		apologyStreak:      0,
		history:            nil,
		name:               "",
		opponentDefections: 0,
		participantNum:     0,
		punishing:          false,
		punishmentStreak:   0,
	}

	for _, input := range inputs {
		t.Run(input.name, func(t *testing.T) {
			gtft.Reset()
			gtft.apologyStreak = input.apologyStreak
			gtft.opponentDefections = input.opponentDefections
			gtft.punishmentStreak = input.punishmentStreak
			gtft.apologizing = input.apologizing
			gtft.punishing = input.punishing

			gtft.updateState(input.opPreviousAction)

			if gtft.apologyStreak != input.expectedApologyStreak {
				t.Errorf("Apology streak: %d, Expected: %d", gtft.apologyStreak, input.expectedApologyStreak)
			}

			if gtft.opponentDefections != input.expectedOpponentDefections {
				t.Errorf("Opponent defections: %d, Expected: %d", gtft.opponentDefections, input.expectedOpponentDefections)
			}

			if gtft.punishmentStreak != input.expectedPunishmentStreak {
				t.Errorf("Punishment streak: %d, Expected: %d", gtft.punishmentStreak, input.expectedPunishmentStreak)
			}

			if gtft.apologizing != input.expectedApologizing {
				t.Errorf("Apologizing: %t, Expected: %t", gtft.apologizing, input.expectedApologizing)
			}

			if gtft.punishing != input.expectedPunishing {
				t.Errorf("Punishing: %t, Expected: %t", gtft.punishing, input.expectedPunishing)
			}
		})
	}
}

func TestReset_GradualTitForTat(t *testing.T) {
	r := round.MakeTestRound(action.Cooperate, action.Defect)
	h := round.MakeTestHistory(r)

	gtft := GradualTitForTat{
		apologizing:        true,
		apologyStreak:      1,
		history:            h,
		opponentDefections: 4,
		participantNum:     round.Participant1,
		punishing:          true,
		punishmentStreak:   3,
	}

	gtft.Reset()

	if gtft.apologizing {
		t.Error("Expected apologizing to be false")
	}

	if gtft.apologyStreak != 0 {
		t.Error("Expected apologyStreak to be 0")
	}

	if len(gtft.history) != 0 {
		t.Error("Expected history to be empty")
	}

	if gtft.opponentDefections != 0 {
		t.Error("Expected opponentDefections to be 0")
	}

	if gtft.punishing {
		t.Error("Expected punishing to be false")
	}

	if gtft.punishmentStreak != 0 {
		t.Error("Expected punishmentStreak to be 0")
	}

	if gtft.participantNum != round.ParticipantNone {
		t.Errorf("Expected participantNum to be %d, got %d", round.ParticipantNone, gtft.participantNum)
	}
}
