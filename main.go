package main

import (
	"log"
	"os"

	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
	"github.com/AlexB138/prisoners_dilemma/internal/strategies"
	"github.com/AlexB138/prisoners_dilemma/internal/tui"
)

/*
TODO:
- Create random ecosystem encounters with global "win"
- Add more strategies
- Update renderer to auto discover strategies
- Update tui with help entries
- Add detailed result view
- Make TUI full screen using alternate screen buffer
*/

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
		score1, score2 := sim.SingleEventScore()
		w := sim.Winner()

		if w == nil {
			log.Println("Tie! Score:", score1, " - ", score2)
			return
		}

		log.Println(sim.Winner().Name(), "won! Score:", score1, " - ", score2)
	}

}
