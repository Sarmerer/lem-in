package main

import (
	"lem-in/parser"
	"lem-in/solver"
)

func main() {
	solver.Solve(parser.ParseFile("./maps/example00.txt"))
}
