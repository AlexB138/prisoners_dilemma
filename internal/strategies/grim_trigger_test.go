package strategies

import (
	"testing"

	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/round"
)

func TestMakeChoice_GrimTrigger(t *testing.T) {
	inputs := []struct {
		name      string
		roundNum  int
		expected  action.Action
		triggered bool
	}{
		{
			name:      "Always cooperate when not triggered",
			roundNum:  2,
			expected:  action.Cooperate,
			triggered: false,
		},
		{
			name:      "Always defect when triggered",
			roundNum:  2,
			expected:  action.Defect,
			triggered: true,
		},
	}

	gt := GrimTrigger{
		name:           "",
		participantNum: round.Participant1,
		triggered:      false,
	}

	for _, input := range inputs {
		t.Run(input.name, func(t *testing.T) {
			gt.Reset()
			gt.participantNum = round.Participant1
			gt.triggered = input.triggered

			received := gt.MakeChoice(input.roundNum)

			if received != input.expected {
				t.Errorf("Received: %s, Expected: %s", received, input.expected)
			}
		})
	}
}

func TestReset_GrimTrigger(t *testing.T) {
	gt := GrimTrigger{
		participantNum: round.Participant1,
		triggered:      true,
	}

	gt.Reset()

	if gt.triggered != false {
		t.Errorf("Expected triggered to be false, got %v", gt.triggered)
	}

	if gt.participantNum != round.ParticipantNone {
		t.Errorf("Expected participantNum to be %d, got %d", round.ParticipantNone, gt.participantNum)
	}
}
