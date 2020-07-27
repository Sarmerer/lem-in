package main

import (
	"lem-in/parser"
	"lem-in/solver"
)

func main() {
	//TO-DO: wtf is start and end in data????
	data, graph := parser.ParseFile("../maps/example00.txt")
	ants, sink := solver.InitAntsAndAssignPaths(data.AntsAmount, graph)
	solver.MoveAnts(ants, sink)
}
