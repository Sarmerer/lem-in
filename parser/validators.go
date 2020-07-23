package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validRoom(line string, roomPointer *string) (int, int, bool) {
	spl := strings.Split(line, " ")
	expectedSplLen := 3

	if len(spl) != expectedSplLen {
		return 0, 0, false
	}
	xCoord, xErr := strconv.Atoi(spl[1])
	if xErr != nil {
		return 0, 0, false
	}
	yCoord, yErr := strconv.Atoi(spl[2])
	if yErr != nil {
		return 0, 0, false
	}
	if len(spl[0]) > 0 {
		if spl[0][0] == '#' || spl[0][0] == 'L' {
			return 0, 0, false
		}
	}
	*roomPointer = spl[0]
	return xCoord, yCoord, true
}

func validLink(line string, linkPointer *[]string) bool {
	spl := strings.Split(line, "-")
	expectedSplLen := 2

	if len(spl) != expectedSplLen {
		return false
	}
	linkFrom := spl[0]
	linkTo := spl[1]

	*linkPointer = append(*linkPointer, linkFrom)
	*linkPointer = append(*linkPointer, linkTo)
	return true
}

func invalidInput(msg string) {
	fmt.Printf("Invalid input: %v\n", msg)
	os.Exit(1)
}
