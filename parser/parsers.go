package parser

import (
	"bufio"
	"lem-in/types"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseFile(fileName string) (*types.Data, *types.Graph) {
	graph := types.InitGraph()

	var data types.Data
	usedIndexes := []int{0}
	roomsMap := make(map[string]int)

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
		invalidInput(-1, "no data provided")
	}
	parseSoreAndAnts(&lines, &usedIndexes, &data.Start, &data.End, &data.AntsAmount)
	parseComments(&lines, &usedIndexes)
	parseRooms(&lines, graph, &usedIndexes, roomsMap)
	parseLinks(&lines, &usedIndexes, graph, roomsMap)
	// fmt.Printf("Ants amount: %v\nStart: %v\nEnd: %v\nRooms:", Data.AntsAmount, Data.Start, Data.End)
	// for _, r := range Data.Rooms {
	// 	fmt.Printf("\n Name: %v\n  x: %v\n  y: %v\n  Links: %v", r.Name, r.CoordX, r.CoordY, r.Connections)
	// }
	return &data, graph
}

func parseSoreAndAnts(lines *[]string, usedIndexes *[]int, start, end *string, antsAmount *int) {
	startFound := false
	endFound := false
	for index, line := range *lines {
		if index == 0 {
			a, err := strconv.Atoi(line)
			if err != nil || a < 1 {
				invalidInput(index, "invalid ants amount")
			}
			*antsAmount = a
		} else {
			if line == "##start" {
				soreCheck(lines, usedIndexes, &startFound, start, index, "start")
			} else if line == "##end" {
				soreCheck(lines, usedIndexes, &endFound, end, index, "end")
			}
		}
	}
	if !startFound {
		invalidInput(-1, "no start room")
	} else if !endFound {
		invalidInput(-1, "no end room")
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
						invalidInput(index, "invalid room params")
					}
				}
				*usedIndexes = append(*usedIndexes, index)
			}
		}
	}
}

func parseRooms(lines *[]string, graph *types.Graph, usedIndexes *[]int, roomsMap map[string]int) {
	for index, line := range *lines {
		var room string
		if indexIsFree(index, usedIndexes) {
			if x, y, valid := validRoom(line, &room); valid {
				if _, ok := roomsMap[room]; !ok {
					roomsMap[room] = len(graph.GetRoomList()) + 1
				} else {
					invalidInput(index, "invalid room params")
				}
				graph.AddRoom(types.Room{room, x, y, false})
				*usedIndexes = append(*usedIndexes, index)
			} else if !validLink(line, &[]string{}) {
				invalidInput(index, "invalid room params")
			}
		}
	}
}

func parseLinks(arr *[]string, usedIndexes *[]int, graph *types.Graph, roomsMap map[string]int) {
	for index, line := range *arr {
		var link []string
		if indexIsFree(index, usedIndexes) {
			if validLink(line, &link) {
				if _, ok := roomsMap[link[0]]; !ok {
					invalidInput(index, "invalid link params")
				} else if _, ok := roomsMap[link[1]]; !ok {
					invalidInput(index, "invalid link params")
				} else if link[0] == link[1] {
					invalidInput(index, "invalid link params")
				}
				var sourceRoom types.Room
				var destRoom types.Room
				for _, room := range graph.GetRoomList() {
					if room.Name == link[0] {
						sourceRoom = room
					}
					if room.Name == link[1] {
						destRoom = room
					}
				}
				graph.AddNeighbour(sourceRoom, destRoom, 1)
			}
		}
	}
}
