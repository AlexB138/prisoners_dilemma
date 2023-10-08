package main

import (
	"fmt"
	"github.com/AlexB138/prisoners_dilemma/internal/simulation"
	strategies2 "github.com/AlexB138/prisoners_dilemma/internal/strategies"
)

/*
TODO: Add CLI interface for selecting strategies
TODO: Create environment for multiple strategies to compete
*/

func main() {
	//s1 := strategies.NewDefector()
	//s2 := strategies.NewCooperator()
	s1 := strategies2.NewRandom()
	s2 := strategies2.NewTitForTat()

	sim := simulation.NewSimulation(15, s1, s2)
	sim.Run()

	fmt.Println(sim)
}
