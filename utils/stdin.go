package utils

import (
	"bufio"
	"fmt"
	"lem-in/config"
	"os"
)

func ProcessInput(args []string) string {
	fileName := config.PathCustomMap
	if len(args) == 1 {
		return args[0]
	} else if len(args) > 1 {
		fmt.Println(config.ErrorManyArgs)
		os.Exit(1)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var text string
		fmt.Println(config.MessageCustomMap)
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
