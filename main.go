package main

import (
	"fmt"
	"github.com/AlexB138/prisoners_dilemma/simulation"
	"github.com/AlexB138/prisoners_dilemma/strategies"
)

func main() {
	s1 := strategies.NewDefector()
	s2 := strategies.NewCooperator()

	sim := simulation.NewSimulation(15, s1, s2)
	sim.Run()

	fmt.Println(sim)
}
