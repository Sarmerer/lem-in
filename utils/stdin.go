package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ProcessInput(args []string) string {
	fileName := "../maps/custom.txt"
	if len(args) == 1 {
		return args[0]
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		os.Exit(1)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var text string
		fmt.Println("Enter map data\nPress Ctrl+D when finished")
		for scanner.Scan() {
			text += scanner.Text() + "\n"
		}
		check(scanner.Err())
		f, err := os.Create(fileName)
		check(err)
		defer f.Close()
		defer f.Sync()
		_, err = f.WriteString(text)
		check(err)
	}
	return fileName
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
