package utils

import (
	"fmt"
	"lem-in/config"
)

func PrintResult(lines *[]string, antMoves *[][]string) {
	for _, line := range *lines {
		fmt.Println(line)
	}
	fmt.Println("\nSolution:")
	var counter int
	for _, move := range *antMoves {
		for _, m := range move {
			fmt.Print(m)
			counter++
		}
		fmt.Println("")
	}
	fmt.Println(config.MessageTurns, counter)
}
