package simulation

type Type string

const (
	SingleEvent Type = "Single Event"
	BestOfN     Type = "Best Of N"
)

type IterativeGameType string

const (
	IterativeGameTypeHighestSingleRound IterativeGameType = "Highest Single Round"
	IterativeGameTypeHighestTotal       IterativeGameType = "Highest Total"
	IterativeGameTypeBestAverageScore   IterativeGameType = "Best Average"
	IterativeGameTypeMostWins           IterativeGameType = "Most Wins"
)
