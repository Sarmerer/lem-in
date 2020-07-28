package solver

import (
	"fmt"
	"lem-in/config"
	"os"
)

// Exit function exits the program if error occured
func Exit(msg string) {
	fmt.Printf(config.ErrorBase, msg)
	os.Exit(1)

}
