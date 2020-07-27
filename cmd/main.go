package main

import (
	"fmt"
	"lem-in/parser"
	"lem-in/solver"
	"os"
	"time"
)

func main() {
	tStart := time.Now()
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Specify the file")
		return
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	data, graph := parser.ParseFile(args[0])
	ants := solver.InitAntsAndAssignPaths(data, graph)
	solver.MoveAnts(ants, data.End)
	fmt.Println("Elapsed:", time.Since(tStart).Seconds())
}
