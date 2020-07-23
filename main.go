package main

import (
	"lem-in/parser"
	"lem-in/solver"
)

func main() {
	a := parser.ParseFile("./maps/example00.txt")
	solver.Solve(a)
}
