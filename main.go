package main

import (
	"fmt"
	"github.com/AlexB138/prisoners_dilemma/simulation"
	"github.com/AlexB138/prisoners_dilemma/strategies"
)

/*
TODO: Add CLI interface for selecting strategies
TODO: Create environment for multiple strategies to compete
*/

func main() {
	//s1 := strategies.NewDefector()
	//s2 := strategies.NewCooperator()
	s1 := strategies.NewRandom()
	s2 := strategies.NewTitForTat()

	sim := simulation.NewSimulation(15, s1, s2)
	sim.Run()

	fmt.Println(sim)
}
