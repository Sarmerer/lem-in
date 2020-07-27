package parser

import (
	"fmt"
	"lem-in/config"
	"os"
	"strconv"
	"strings"
)

//This function validates room parameters.
//It returns false if:
//coordinates are incorrect,
//room name is empty, or contains # or L.
func validRoom(line string, roomName *string) (int, int, bool) {
	spl := strings.Split(line, " ")
	expectedSplLen := 3

	if len(spl) != expectedSplLen {
		return 0, 0, false
	}
	xCoord, xErr := strconv.Atoi(spl[1])
	yCoord, yErr := strconv.Atoi(spl[2])
	if xErr != nil || yErr != nil || xCoord < 0 || yCoord < 0 {
		return 0, 0, false
	}
	if len(spl[0]) > 0 {
		if spl[0][0] == '#' || spl[0][0] == 'L' {
			return 0, 0, false
		}
	} else {
		return 0, 0, false
	}
	*roomName = spl[0]
	return xCoord, yCoord, true
}

//This function validates link parameters.
//It returns false if link declaration is not valid.
func validLink(line string, linkPointer *[]string) bool {
	spl := strings.Split(line, "-")

	if len(spl) != 2 {
		return false
	}
	linkFrom := spl[0]
	linkTo := spl[1]

	*linkPointer = append(*linkPointer, linkFrom)
	*linkPointer = append(*linkPointer, linkTo)
	return true
}

func invalidInput(line int, msg string) {
	if line >= 0 {
		fmt.Printf(config.ErrorBaseExact, line+1, msg)
	} else {
		fmt.Printf(config.ErrorBase, msg)
	}
	os.Exit(1)
}
