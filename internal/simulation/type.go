package simulation

type Type string

const (
	SingleEvent Type = "Single Event"
	BestOfN     Type = "Best Of N"
)

type IterativeGameType string

const (
	IterativeGameTypeHighestSingleEvent IterativeGameType = "Highest Single Event"
	IterativeGameTypeHighestTotal       IterativeGameType = "Highest Total"
	IterativeGameTypeBestAverageScore   IterativeGameType = "Best Average"
	IterativeGameTypeMostWins           IterativeGameType = "Most Wins"
)
