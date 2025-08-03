package main

import (
	"log"
	"os"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
	"github.com/AlexB138/prisoners_dilemma/internal/tui"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "--tui" {
		if err := tui.Run(); err != nil {
			log.Fatal(err)
		}
	} else {
		s := simulation.Settings{
			IterativeGameType: "",
			Iterations:        1,
			Rounds:            9,
			Type:              simulation.SingleEvent,
			Strategy1:         strategies.NewRandom(),
			Strategy2:         strategies.NewTitForTat(),
		}

		sim := simulation.NewSimulation(s)
		sim.Run()
		score1, score2 := sim.GetFinalScores()
		w := sim.GetWinner()

		if w == nil {
			log.Println("Tie! Score:", score1, " - ", score2)
			return
		}

		log.Println(sim.GetWinner().GetName(), "won! Score:", score1, " - ", score2)
	}

}
