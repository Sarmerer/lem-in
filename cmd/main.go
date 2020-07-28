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
	result :=
		solver.MoveAnts(
			solver.InitAntsAndAssignPaths(
				parser.ParseFile(lines),
			),
		)
	elapsed := time.Since(tStart).Seconds()
	utils.PrintResult(lines, result)
	fmt.Printf(config.MessageElapsed, elapsed)
}
