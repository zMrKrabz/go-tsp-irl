package main

import (
	"log"

	"github.com/zMrKrabz/go-tsp-irl/solver"
)

func main() {
	// initiralize solver
	var s solver.Solver

	// initialize and add coordiantes
	first := solver.Coordinate{
		X: 13.388860,
		Y: 52.517037,
	}

	second := solver.Coordinate{
		X: 13.397634,
		Y: 52.529407,
	}

	third := solver.Coordinate{
		X: 13.428555,
		Y: 52.523219,
	}

	s.AddCoordinate(first)
	s.AddCoordinate(second)
	s.AddCoordinate(third)

	// Solve!
	log.Println(s.String()) // initial order of array
	s.Solve()
	log.Println(s.String()) // new order of array
}
