package solver

import (
	"fmt"
	"lem-in/config"
	"lem-in/types"
	"lem-in/utils"
)

// InitAntsAndAssignPaths initialazes Ants and assigns an optimal path to each Ant
// Use the following logic to calculate an optimal path:
// If Rooms in Path1 + Ants in Path1 > Rooms in Path2 send Ant to Path2
// Otherwise, send Ant to Path1
func InitAntsAndAssignPaths(data *types.Data, graph *types.Graph) ([]types.Ant, *types.Room) {
	// Find all valid paths with Edmonds-Karp and BFS
	paths := EdmondsKarp(graph, data.Start, data.End)
	// If var paths is empty - no valid paths were found
	// Exit the program
	if len(paths) == 0 {
		utils.InvalidInput(-1, config.ErrorNoPaths)
	}
	// Initialize ants array and assign each ant ID and start(source) position
	ants := make([]types.Ant, data.AntsAmount)
	for i := 0; i < data.AntsAmount; i++ {
		ants[i].ID = i
		ants[i].Position = data.Start
	}
	// Initialize antsInPath to know the current an=mount of ants in each path
	antsInPath := make([]int, len(paths))
	// Send ant1 to path1 because path1 is the shortest
	ants[0].Path = paths[0]
	// Now there is one ant in path1
	antsInPath[0]++
	// Iterate through the rest of ants to assign each a path
	for i, currPath := 1, 0; i < len(ants); i++ {
		// Make sure that next path exists in Paths Array
		// If Rooms in CurrentPath + Ants in CurrentPath > Rooms in NextPath
		// Assign Ant to NextPath
		if (currPath+1) < len(paths) &&
			(len(paths[currPath])+antsInPath[currPath]) > len(paths[currPath+1]) {
			currPath++
			antsInPath[currPath]++
			ants[i].Path = paths[currPath]
		} else {
			// If we get here: next path doesn't exist or
			// Rooms in CurrentPath + Ants in CurrentPath <= Rooms in NextPath
			// Set Current Path to Path1 and assign Ant to Current Path
			currPath = 0
			ants[i].Path = paths[currPath]
			antsInPath[currPath]++
		}
	}
	// All ants have paths assigned now so return
	return ants, &data.End
}

// MoveAnts moves all ants from source to sink in a correct order
func MoveAnts(ants []types.Ant, sink *types.Room) *[][]string {
	// resultToPriint to hold moves of ants
	var resultToPrint [][]string
	// Set next room to 1 because Room0 is the source(start)
	nextRoom := 1
	// Iterate until all ants reach the sink (end)
	for !allAntsIn(ants, *sink) {
		var iterationToPrint []string
		// Iterate through each ant and move it if next room is free
		for i := range ants {
			// If ant has already reached the sink - skip it
			if ants[i].Position != *sink {
				// Make sure sink is always empty
				sink.HasAnt = false
				// If next room in Ant's path is the sink:
				// Move Ant to sink, free up the previous room and continue
				if ants[i].Path[nextRoom] == *sink {
					iterationToPrint = append(iterationToPrint, fmt.Sprintf("L%v-%v ", i, ants[i].Path[nextRoom].Name))
					ants[i].Position = *sink
					ants[i].Path[nextRoom-1].HasAnt = false
					continue
				}
				// Otherwise, check the next room availabilty
				// If next room is empty: move Ant and free up the previous room and
				// cut the Ant's Path so that it always starts from the current room
				// If next room is not empty: skip the ant this time
				if !ants[i].Path[nextRoom].HasAnt {
					iterationToPrint = append(iterationToPrint, fmt.Sprintf("L%v-%v ", i, ants[i].Path[nextRoom].Name))
					ants[i].Position = ants[i].Path[nextRoom]
					ants[i].Path[nextRoom].HasAnt = true
					ants[i].Path[nextRoom-1].HasAnt = false
					ants[i].Path = ants[i].Path[nextRoom:]
				}
			}
		}
		resultToPrint = append(resultToPrint, iterationToPrint)
	}
	return &resultToPrint
}

// allAntsIn checks if all ants have reached the sink
func allAntsIn(ants []types.Ant, sink types.Room) bool {
	for _, ant := range ants {
		if ant.Position != sink {
			return false
		}
	}
	return true
}
