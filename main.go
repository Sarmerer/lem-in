package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rooms struct {
	Rooms []room
	Start string
	End   string
}

type room struct {
	Name        string
	Ants        int
	Connections []string
}

type ant struct {
	ID       int
	Position string
}

func parseFile(fileName string) {
	var start string
	var antsAmount int
	var end string
	startFound := false
	endFound := false

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []string
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	for index, line := range arr {
		if index == 0 {
			a, err := strconv.Atoi(line)
			if err != nil {
				invalidInput("invalid ants amount")
			}
			antsAmount = a
		} else {
			if line == "##start" {
				soreCheck(&arr, &startFound, &start, index, "start")
			} else if line == "##end" {
				soreCheck(&arr, &endFound, &end, index, "end")
			} else {

			}
		}
	}
	if !startFound {
		invalidInput("no start room")
	} else if !endFound {
		invalidInput("no end room")
	}
	fmt.Printf("Ants amount: %v\nStart: %v\nEnd: %v\n", antsAmount, start, end)
}

func soreCheck(arr *[]string, found *bool, sorePointer *string, index int, sore string) { //sore == start or end
	if !*found {
		if index < len(*arr)-1 {
			if !validRoom((*arr)[index+1], "room", sorePointer) {
				invalidInput("invalid " + sore + " room params")
			}
			*found = true
		} else {
			invalidInput("no " + sore + " room coords")
		}
	} else {
		invalidInput("another " + sore + " declaration")
	}
}

func invalidInput(msg string) {
	fmt.Printf("Invalid input: %v\n", msg)
	os.Exit(1)
}

func validRoom(line, lineType string, roomPointer *string) bool {
	spl := strings.Split(line, " ")
	expectedSplLen := -1
	switch lineType {
	case "room":
		expectedSplLen = 3
	case "link":
		expectedSplLen = 2
	}
	if len(spl) != expectedSplLen {
		return false
	}
	xCoord, xErr := strconv.Atoi(spl[1])
	if xErr != nil {
		return false
	}
	_ = xCoord
	yCoord, yErr := strconv.Atoi(spl[2])
	if yErr != nil {
		return false
	}
	_ = yCoord
	if len(spl[0]) > 0 {
		if spl[0][0] == '#' || spl[0][0] == 'L' {
			return false
		}
	}
	*roomPointer = spl[0]
	return true
}

func main() {
	parseFile("example00.txt")
}
