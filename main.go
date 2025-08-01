package main

import (
	"fmt"

	"github.com/AlexB138/prisoners_dilemma/internal/event"
	strategies "github.com/AlexB138/prisoners_dilemma/internal/strategies"
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

	e := event.NewEvent(15, s1, s2)
	e.Run()

	fmt.Println(e)
}
