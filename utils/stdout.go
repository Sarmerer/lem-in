package utils

import "fmt"

func PrintResult(lines *[]string) {
	for _, line := range *lines {
		fmt.Println(line)
	}
}
