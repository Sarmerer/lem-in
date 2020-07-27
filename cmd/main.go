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
	lines := parser.ReadFile(file)
	result :=
		solver.MoveAnts(
			solver.InitAntsAndAssignPaths(
				parser.ParseFile(lines),
			),
		)
	utils.PrintResult(lines, result)
	fmt.Printf("\nTime elapsed: %vs\n", time.Since(tStart).Seconds())
}
