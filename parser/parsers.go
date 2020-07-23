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
	var start string
	var end string
	var rooms []string
	var links [][]string
	var test data

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
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		fmt.Println("Not enough data")
		os.Exit(1)
	}

	for index, line := range lines {
		if index == 0 {
			a, err := strconv.Atoi(line)
			if err != nil || a < 1 {
				invalidInput("invalid ants amount")
			}
			antsAmount = a
			test.AntsAmount = a
		} else {
			if line == "##start" {
				soreCheck(&lines, &usedIndexes, &startFound, &start, index, "start")
			} else if line == "##end" {
				soreCheck(&lines, &usedIndexes, &endFound, &end, index, "end")
			}
		}
	}
	if !startFound {
		invalidInput("no start room")
	} else if !endFound {
		invalidInput("no end room")
	}
	parseComments(&lines, &usedIndexes)
	parseRooms(&lines, &rooms, &usedIndexes)
	parseLinks(&lines, &links, &usedIndexes)
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
