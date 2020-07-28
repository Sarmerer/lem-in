package utils

import (
	"fmt"
	"lem-in/config"
	"os"
)

func InvalidInput(line int, msg string) {
	if line >= 0 {
		fmt.Printf(config.ErrorBaseExact, line+1, msg)
	} else {
		fmt.Printf(config.ErrorBase, msg)
	}
	os.Exit(1)
}
