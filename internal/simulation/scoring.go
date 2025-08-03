package simulation

import (
	"github.com/AlexB138/prisoners_dilemma/internal/action"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
)

func (s *Simulation) bestOfNWinner() strategies.Strategy {
	switch s.settings.IterativeGameType {
	case IterativeGameTypeHighestSingleEvent:
		return s.highestSingleEventWinner()

	case IterativeGameTypeHighestTotal:
		return s.highestTotalWinner()

	case IterativeGameTypeBestAverageScore:
		return s.bestAverageScoreWinner()

	case IterativeGameTypeMostWins:
		return s.mostWinsWinner()

	default:
		return nil
	}
}

func (s *Simulation) highestSingleEventWinner() strategies.Strategy {
	var highScore action.Score = 0
	var winner strategies.Strategy

	for _, e := range s.events {
		score1, score2 := e.GetScore()
		if score1 > highScore {
			highScore = score1
			winner = s.settings.Strategy1

		} else if score2 > highScore {
			highScore = score2
			winner = s.settings.Strategy2
		}
	}

	return winner
}

func (s *Simulation) highestTotalWinner() strategies.Strategy {
	var total1, total2 action.Score

	for _, e := range s.events {
		score1, score2 := e.GetScore()
		total1 += score1
		total2 += score2
	}

	if total1 > total2 {
		return s.settings.Strategy1
	} else {
		return s.settings.Strategy2
	}

}

func (s *Simulation) bestAverageScoreWinner() strategies.Strategy {
	var scores1, scores2 []int
	for _, e := range s.events {
		score1, score2 := e.GetScore()
		scores1 = append(scores1, int(score1))
		scores2 = append(scores2, int(score2))
	}

	avg1 := average(scores1)
	avg2 := average(scores2)

	if avg1 > avg2 {
		return s.settings.Strategy1
	} else {
		return s.settings.Strategy2
	}
}

func (s *Simulation) mostWinsWinner() strategies.Strategy {
	var wins1, wins2 int
	for _, e := range s.events {
		if e.Winner() == s.settings.Strategy1 {
			wins1++
		} else {
			wins2++
		}
	}

	if wins1 > wins2 {
		return s.settings.Strategy1
	} else {
		return s.settings.Strategy2
	}
}

func average(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	var sum int
	for _, n := range nums {
		sum += n
	}

	return float64(sum) / float64(len(nums))
}
