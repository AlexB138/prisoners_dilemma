package simulation

type Type string

const (
	SingleEvent Type = "Single Event"
	BestOfN     Type = "Best Of N"
)

var typeToHelp = map[Type]string{
	SingleEvent: "Plays a single event, scored on total points.",
	BestOfN:     "Plays N events with configurable scoring criteria.",
}

type IterativeGameType string

const (
	IterativeGameTypeHighestSingleEvent IterativeGameType = "Highest Single Event"
	IterativeGameTypeHighestTotal       IterativeGameType = "Highest Total"
	IterativeGameTypeBestAverageScore   IterativeGameType = "Best Average"
	IterativeGameTypeMostWins           IterativeGameType = "Most Wins"
)

var iterativeGameTypeToHelp = map[IterativeGameType]string{
	IterativeGameTypeHighestSingleEvent: "The strategy with the highest score in a single event wins.",
	IterativeGameTypeHighestTotal:       "The strategy with the highest aggregate score across all events wins.",
	IterativeGameTypeBestAverageScore:   "The strategy with the highest average score across all events wins. Average is calculated as total score / N. N is the number of events played. E.g. 100/5 = 20.00.",
	IterativeGameTypeMostWins:           "The strategy that wins the most events wins."}
