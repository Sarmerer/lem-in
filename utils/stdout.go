package utils

import (
	"fmt"
	"lem-in/config"
)

// PrintResult prints file contents and ants' moves
func PrintResult(lines *[]string, antMoves *[][]string) {
	for _, line := range *lines {
		fmt.Println(line)
	}
	fmt.Println("\nSolution:")
	var turns int
	var newLines int
	for _, move := range *antMoves {
		newLines++
		for _, m := range move {
			fmt.Print(m)
			turns++
		}
		fmt.Println("")
	}
	fmt.Println(config.MessageTurns, turns, config.MessageLines, newLines)
}
