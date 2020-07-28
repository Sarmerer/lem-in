package main

import (
	"fmt"
	"lem-in/config"
	"lem-in/parser"
	"lem-in/solver"
	"lem-in/utils"
	"time"
)

func main() {
	// file := utils.ProcessInput(os.Args[1:])
	tStart := time.Now()
	lines := parser.ReadFile("../maps/alem-audit/e0.txt")
	result :=
		solver.MoveAnts(
			solver.InitAntsAndAssignPaths(
				parser.ParseFile(lines),
			),
		)
	utils.PrintResult(lines, result)
	fmt.Printf(config.MessageElapsed, time.Since(tStart).Seconds())
}
