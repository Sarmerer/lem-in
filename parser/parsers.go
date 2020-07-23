package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseFile(fileName string) {
	var links [][]string
	var Data data

	usedIndexes := []int{0}

	roomsMap = make(map[string]string)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		invalidInput("no data provided")
	}
	parseSoreAndAnts(&lines, &usedIndexes, &Data.Start, &Data.End, &Data.AntsAmount)
	parseComments(&lines, &usedIndexes)
	parseRooms(&lines, &Data.Rooms, &usedIndexes)
	parseLinks(&lines, &links, &usedIndexes)
	fmt.Printf("Ants amount: %v\nStart: %v\nEnd: %v\nRooms:", Data.AntsAmount, Data.Start, Data.End)
	for _, r := range Data.Rooms {
		fmt.Printf("\n Name: %v\n  x: %v\n  y: %v\n  Links: %v", r.Name, r.CoordX, r.CoordY, r.Connections)
	}
}

func parseSoreAndAnts(lines *[]string, usedIndexes *[]int, start, end *string, antsAmount *int) {
	startFound := false
	endFound := false
	for index, line := range *lines {
		if index == 0 {
			antsAmount, err := strconv.Atoi(line)
			if err != nil || antsAmount < 1 {
				invalidInput("invalid ants amount")
			}
		} else {
			if line == "##start" {
				soreCheck(lines, usedIndexes, &startFound, start, index, "start")
			} else if line == "##end" {
				soreCheck(lines, usedIndexes, &endFound, end, index, "end")
			}
		}
	}
	if !startFound {
		invalidInput("no start room")
	} else if !endFound {
		invalidInput("no end room")
	}
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

func parseRooms(lines *[]string, rooms *[]roomStruct, usedIndexes *[]int) {
	for index, line := range *lines {
		var room string
		var extra []string
		if indexIsFree(index, usedIndexes) {
			if x, y, valid := validRoom(line, &room); valid {
				if _, ok := roomsMap[room]; !ok {
					roomsMap[room] = room
				} else {
					invalidInput("invalid room params")
				}
				*rooms = append(*rooms, roomStruct{room, x, y, []string{}})
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
