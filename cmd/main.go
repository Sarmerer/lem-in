package main

import (
	"fmt"
	"lem-in/config"
	"lem-in/parser"
	"lem-in/solver"
	"lem-in/utils"
	"os"
	"time"
)

func main() {
	file := utils.ProcessInput(os.Args[1:])
	tStart := time.Now()
	lines := parser.ReadFile(file)
	data, graph := parser.ParseFile(lines)
	ants, paths := solver.InitAntsAndAssignPaths(data, graph)
	result := solver.MoveAnts(ants, &data.End)
	elapsed := time.Since(tStart).Seconds()
	utils.PrintResult(lines, result)
	fmt.Printf(config.MessageElapsed, elapsed)
	utils.Marshal(data, graph, paths)
}
