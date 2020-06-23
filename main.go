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

var roomsMap map[string]string

func parseFile(fileName string) {
	var start string
	var end string
	var rooms []string
	var links [][]string

	usedIndexes := []int{0}

	var antsAmount int

	startFound := false
	endFound := false
	roomsMap = make(map[string]string)

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
			if err != nil || a < 1 {
				invalidInput("invalid ants amount")
			}
			antsAmount = a
		} else {
			if line == "##start" {
				soreCheck(&arr, &usedIndexes, &startFound, &start, index, "start")
			} else if line == "##end" {
				soreCheck(&arr, &usedIndexes, &endFound, &end, index, "end")
			}
		}
	}
	if !startFound {
		invalidInput("no start room")
	} else if !endFound {
		invalidInput("no end room")
	}
	parseComments(&arr, &usedIndexes)
	parseRooms(&arr, &rooms, &usedIndexes)
	parseLinks(&arr, &links, &usedIndexes)
	fmt.Printf("Ants amount: %v\nStart: %v\nEnd: %v\nRooms: %v\nLinks: %v\n", antsAmount, start, end, rooms, links)
}

func parseComments(arr *[]string, usedIndexes *[]int) {
	for index, line := range *arr {
		if len(line) > 0 {
			if line[0] == '#' {
				spl := strings.Split(line, " ")
				if len(spl) == 3 {
					_, xErr := strconv.Atoi(spl[1])
					_, yErr := strconv.Atoi(spl[2])
					if xErr != nil || yErr != nil {
						invalidInput("invalid room params")
					}
				}
				*usedIndexes = append(*usedIndexes, index)
			}
		}
	}
}

func parseRooms(arr, rooms *[]string, usedIndexes *[]int) {
	for index, line := range *arr {
		var room string
		var extra []string
		if indexIsFree(index, usedIndexes) {
			if validRoom(line, &room) {
				if _, ok := roomsMap[room]; !ok {
					roomsMap[room] = room
				} else {
					invalidInput("invalid room params")
				}
				*rooms = append(*rooms, room)
				*usedIndexes = append(*usedIndexes, index)
			} else if !validLink(line, &extra) {
				invalidInput("invalid room params")
			}
		}
	}
}

func parseLinks(arr *[]string, links *[][]string, usedIndexes *[]int) {
	for index, line := range *arr {
		var link []string
		if indexIsFree(index, usedIndexes) {
			if validLink(line, &link) {
				if _, ok := roomsMap[link[0]]; !ok {
					invalidInput("invalid link params")
				} else if _, ok := roomsMap[link[1]]; !ok {
					invalidInput("invalid link params")
				} else if link[0] == link[1] {
					invalidInput("invalid link params")
				}
				*links = append(*links, link)
				link = nil
			}
		}
	}
}

func roomExists() {

}

func indexIsFree(index int, usedIndexes *[]int) bool {
	for _, idx := range *usedIndexes {
		if idx == index {
			return false
		}
	}
	return true
}

func soreCheck(arr *[]string, usedIndexes *[]int, found *bool, sorePointer *string, index int, sore string) { //sore == start or end
	if !*found {
		if index < len(*arr)-1 {
			if !validRoom((*arr)[index+1], sorePointer) {
				invalidInput("invalid " + sore + " room params")
			} else {
				*usedIndexes = append(*usedIndexes, index)
			}
			*found = true
		} else {
			invalidInput("no " + sore + " room coords")
		}
	} else {
		invalidInput("another " + sore + " declaration")
	}
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

func invalidInput(msg string) {
	fmt.Printf("Invalid input: %v\n", msg)
	os.Exit(1)
}

func main() {
	parseFile("example00.txt")
}
