package main

import (
	"fmt"
	"lem-in/config"
	"lem-in/parser"
	"lem-in/solver"
	"lem-in/utils"
	"lem-in/visualizer"
	"os"
	"time"
)

func main() {
	file := utils.ProcessInput(os.Args[1:])
	tStart := time.Now()
	lines := parser.ReadFile(file)

	graph := parser.ParseFile(lines)
	solver.InitAntsAndAssignPaths(graph)
	result := solver.MoveAnts(graph)

	elapsed := time.Since(tStart).Seconds()
	utils.PrintResult(lines, result)
	fmt.Printf(config.MessageElapsed, elapsed)
	utils.Marshal(graph)
	visualizer.StartServer()
}
