package round

import (
	"fmt"
)

// Participant represents a participant number in a game
type Participant int

const (
	// Participant1 represents the first participant
	Participant1 Participant = 1
	// Participant2 represents the second participant
	Participant2 Participant = 2
	// ParticipantNone represents an unassigned participant
	ParticipantNone Participant = 0
)

// String returns the string representation of the participant
func (p Participant) String() string {
	switch p {
	case Participant1:
		return "1"
	case Participant2:
		return "2"
	case ParticipantNone:
		return "none"
	default:
		return fmt.Sprintf("invalid(%d)", int(p))
	}
}

// IsValid returns true if the participant number is valid
func (p Participant) IsValid() bool {
	return p == Participant1 || p == Participant2
}

// IsAssigned returns true if the participant has been assigned a number
func (p Participant) IsAssigned() bool {
	return p != ParticipantNone
}

// Opponent returns the opponent participant number
func (p Participant) Opponent() Participant {
	switch p {
	case Participant1:
		return Participant2
	case Participant2:
		return Participant1
	default:
		return ParticipantNone
	}
}

// ToInt returns the participant as an int for backward compatibility
func (p Participant) ToInt() int {
	return int(p)
}

// FromInt creates a Participant from an int, returning ParticipantNone if invalid
func FromInt(n int) Participant {
	switch n {
	case 1:
		return Participant1
	case 2:
		return Participant2
	default:
		return ParticipantNone
	}
}

// ValidateInt validates an int as a participant number
func ValidateInt(n int) bool {
	return n == 1 || n == 2
}

// MustFromInt creates a Participant from an int, panicking if invalid
// Use this only when you're certain the int is valid (1 or 2)
func MustFromInt(n int) Participant {
	if !ValidateInt(n) {
		panic(fmt.Sprintf("invalid participant number: %d", n))
	}
	return FromInt(n)
}

// ValidateParticipant validates a Participant value
func ValidateParticipant(p Participant) bool {
	return p.IsValid()
}

// IsOpponent returns true if the two participants are opponents
func IsOpponent(p1, p2 Participant) bool {
	return p1.IsValid() && p2.IsValid() && p1 != p2
}

// IsSame returns true if the two participants are the same
func IsSame(p1, p2 Participant) bool {
	return p1.IsValid() && p2.IsValid() && p1 == p2
}
