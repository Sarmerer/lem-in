package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func validRoom(line string, roomPointer *string) bool {
	spl := strings.Split(line, " ")
	expectedSplLen := 3

	if len(spl) != expectedSplLen {
		return false
	}
	xCoord, xErr := strconv.Atoi(spl[1])
	if xErr != nil {
		return false
	}
	yCoord, yErr := strconv.Atoi(spl[2])
	if yErr != nil {
		return false
	}
	_ = xCoord
	_ = yCoord
	if len(spl[0]) > 0 {
		if spl[0][0] == '#' || spl[0][0] == 'L' {
			return false
		}
	}
	*roomPointer = spl[0]
	return true
}

func invalidInput(msg string) {
	fmt.Printf("Invalid input: %v\n", msg)
	os.Exit(1)
}
