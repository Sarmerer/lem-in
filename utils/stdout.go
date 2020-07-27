package utils

import "fmt"

func PrintResult(lines *[]string, antMoves *[][]string) {
	for _, line := range *lines {
		fmt.Println(line)
	}
	fmt.Println("\nSolution:")
	for _, move := range *antMoves {
		for _, m := range move {
			fmt.Print(m)
		}
		fmt.Println("")
	}
}
