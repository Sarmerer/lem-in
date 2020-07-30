package parser

import (
	"bufio"
	"fmt"
	"lem-in/config"
	"lem-in/types"
	"lem-in/utils"
	"os"
	"strconv"
	"strings"
)

/*
ReadFile function takes a fileName,
which should be a path to a desired file,
and returns it's content as an array of strings.
*/
func ReadFile(fileName string) *[]string {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		utils.InvalidInput(-1, config.ErrorNoData)
	}
	return &lines
}

/*
ParseFile function takes an array of strings to find:
start room, end room, all other rooms and all links.
It also checks if there is: valid ants amount, no
rooms duplicates, no links to unexistent rooms,
no rooms with invalid names or coordinates.
*/
func ParseFile(lines *[]string) (*types.Data, *types.Graph) {
	graph := types.InitGraph()
	var data types.Data
	usedIndexes := []int{0}
	//usedCoordsMap := make(map[int]int)
	roomsMap := make(map[string]int)

	parseSoreAndAnts(lines, &usedIndexes, &data)
	parseComments(lines, &usedIndexes)
	parseRooms(lines, graph, &usedIndexes, roomsMap)
	parseLinks(lines, &usedIndexes, graph, roomsMap)
	return &data, graph
}

//sore == start or end.
//parseSoreAndAnts function finds the start and the  end room and the amount of ants.
//Program wii terminate if: there is no start/end room, start/end has duplicates,
//start/end have invalid name or coordinate, ants amount is lesser or equal 0.
func parseSoreAndAnts(lines *[]string, usedIndexes *[]int, data *types.Data) {
	start := &data.Start
	end := &data.End
	startFound := false
	endFound := false
	for index, line := range *lines {
		if index == 0 {
			a, err := strconv.Atoi(line)
			if err != nil || a < 1 {
				utils.InvalidInput(index, config.ErrorAnts)
			}
			data.AntsAmount = a
		} else {
			if line == "##start" {
				soreCheck(lines, usedIndexes, &startFound, start, index, "start")
			} else if line == "##end" {
				soreCheck(lines, usedIndexes, &endFound, end, index, "end")
			}
		}
	}
	if !startFound {
		utils.InvalidInput(-1, config.ErrorNoStart)
	} else if !endFound {
		utils.InvalidInput(-1, config.ErrorNoEnd)
	}
}

//parseComments function finds all the comments in the file.
//Note:
//Project instructions say: "A room will never start with the letter L or with # and must have no spaces"
//It is not clear if a program should consider a room, starting with L or #, as an invalid input,
//or it is a rule of how rooms should be named.
//
//That is why: "#comment 12 45", will be considered as wrong room parameters.
func parseComments(arr *[]string, usedIndexes *[]int) {
	for index, line := range *arr {
		if len(line) > 0 {
			if line[0] == '#' {
				spl := strings.Split(line, " ")
				if len(spl) == 3 {
					_, xErr := strconv.Atoi(spl[1])
					_, yErr := strconv.Atoi(spl[2])
					if xErr == nil && yErr == nil {
						utils.InvalidInput(index, config.ErrorRoom)
					}
				}
				*usedIndexes = append(*usedIndexes, index)
			}
		}
	}
}

//parseRooms function finds all the rooms.
//It checks if a room:
//is unique,
//has positive numeric coordinates,
//doesn't contain # or L in it's name.
func parseRooms(lines *[]string, graph *types.Graph, usedIndexes *[]int, roomsMap map[string]int) {
	for index, line := range *lines {
		var room string
		if indexIsFree(index, usedIndexes) {
			if x, y, valid := validRoom(line, &room); valid {
				if _, ok := roomsMap[room]; !ok {
					roomsMap[room] = len(graph.GetRoomList()) + 1
				} else {
					utils.InvalidInput(index, config.ErrorRoom)
				}
				graph.AddRoom(types.Room{Name: room, X: x, Y: y, HasAnt: false})
				*usedIndexes = append(*usedIndexes, index)
			} else if !validLink(line, &[]string{}) {
				utils.InvalidInput(index, config.ErrorRoom)
			}
		}
	}
}

//parseLinks function finds all the links.
//It checks if a link:
//doesn't connect unexistent rooms,
//doesn't connect room with itself.
func parseLinks(arr *[]string, usedIndexes *[]int, graph *types.Graph, roomsMap map[string]int) {
	for index, line := range *arr {
		var link []string
		if indexIsFree(index, usedIndexes) {
			if validLink(line, &link) {
				if _, ok := roomsMap[link[0]]; !ok {
					utils.InvalidInput(index, config.ErrorLink)
				} else if _, ok := roomsMap[link[1]]; !ok {
					utils.InvalidInput(index, config.ErrorLink)
				} else if link[0] == link[1] {
					utils.InvalidInput(index, config.ErrorLink)
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
