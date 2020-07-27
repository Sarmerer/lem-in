package main

import (
	"fmt"
	"lem-in/parser"
	"lem-in/solver"
	"lem-in/utils"
	"os"
	"time"
)

func main() {
	file := utils.ProcessInput(os.Args[1:])
	tStart := time.Now()
	data, graph, lines := parser.ParseFile(file)
	ants := solver.InitAntsAndAssignPaths(data, graph)
	utils.PrintResult(lines)
	solver.MoveAnts(ants, data.End)
	fmt.Println("Elapsed:", time.Since(tStart).Seconds())
}
