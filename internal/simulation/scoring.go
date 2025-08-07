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

func (s *Simulation) HighestSingleEventScore() (int, int) {
	var highScore1 action.Score = 0
	var highScore2 action.Score = 0

	for _, e := range s.events {
		score1, score2 := e.GetScore()
		if score1 > highScore1 {
			highScore1 = score1
		}

		if score2 > highScore2 {
			highScore2 = score2
		}
	}

	return int(highScore1), int(highScore2)
}

func (s *Simulation) highestSingleEventWinner() strategies.Strategy {
	highScore1, highScore2 := s.HighestSingleEventScore()

	if highScore1 == highScore2 {
		return nil
	} else if highScore1 > highScore2 {
		return s.settings.Strategy1
	} else {
		return s.settings.Strategy2
	}
}

func (s *Simulation) HighestTotalScore() (int, int) {
	var total1, total2 action.Score

	for _, e := range s.events {
		score1, score2 := e.GetScore()
		total1 += score1
		total2 += score2
	}

	return int(total1), int(total2)
}

func (s *Simulation) highestTotalWinner() strategies.Strategy {
	total1, total2 := s.HighestTotalScore()

	if total1 == total2 {
		return nil
	} else if total1 > total2 {
		return s.settings.Strategy1
	} else {
		return s.settings.Strategy2
	}

}

func (s *Simulation) BestAverageScore() (float64, float64) {
	var scores1, scores2 []int
	for _, e := range s.events {
		score1, score2 := e.GetScore()
		scores1 = append(scores1, int(score1))
		scores2 = append(scores2, int(score2))
	}

	return average(scores1), average(scores2)
}

func (s *Simulation) bestAverageScoreWinner() strategies.Strategy {
	avg1, avg2 := s.BestAverageScore()

	if avg1 == avg2 {
		return nil
	} else if avg1 > avg2 {
		return s.settings.Strategy1
	} else {
		return s.settings.Strategy2
	}
}

func (s *Simulation) MostWinsScore() (int, int) {
	var wins1, wins2 int
	for _, e := range s.events {
		if e.Winner() == s.settings.Strategy1 {
			wins1++
		} else if e.Winner() == s.settings.Strategy2 {
			wins2++
		}
	}

	return wins1, wins2
}

func (s *Simulation) mostWinsWinner() strategies.Strategy {
	wins1, wins2 := s.MostWinsScore()

	if wins1 == wins2 {
		return nil
	} else if wins1 > wins2 {
		return s.settings.Strategy1
	} else {
		return s.settings.Strategy2
	}
}

func (s *Simulation) SingleEventScore() (action.Score, action.Score) {
	return s.events[len(s.events)-1].GetScore()
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
