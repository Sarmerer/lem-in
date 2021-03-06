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
func InitAntsAndAssignPaths(graph *types.Graph) {
	// Find all valid paths with Edmonds-Karp and BFS
	graph.Paths = EdmondsKarp(graph, graph.Start, graph.End)
	// If var paths is empty - no valid paths were found
	// Exit the program
	if len(graph.Paths) == 0 {
		utils.InvalidInput(-1, config.ErrorNoPaths)
	}
	// Initialize ants array and assign each ant ID and start(source) position
	graph.Ants = make([]types.Ant, graph.AntsAmount)
	for i := 0; i < graph.AntsAmount; i++ {
		graph.Ants[i].ID = i
		graph.Ants[i].Position = graph.Start
	}
	// Initialize antsInPath to know the current an=mount of ants in each path
	antsInPath := make([]int, len(graph.Paths))
	// Send ant1 to path1 because path1 is the shortest
	graph.Ants[0].Path = graph.Paths[0]
	// Now there is one ant in path1
	antsInPath[0]++
	// Iterate through the rest of ants to assign each a path
	for i, currPath := 1, 0; i < len(graph.Ants); i++ {
		// If there is a direct path between source and sink
		// assign each ant path source-sink
		if len(graph.Paths[currPath]) == 2 {
			graph.Ants[i].Path = graph.Paths[currPath]
			continue
		}
		// Make sure that next path exists in Paths Array
		// If Rooms in CurrentPath + Ants in CurrentPath > Rooms in NextPath
		// Assign Ant to NextPath
		if (currPath+1) < len(graph.Paths) &&
			(len(graph.Paths[currPath])+antsInPath[currPath]) > len(graph.Paths[currPath+1]) {
			currPath++
			antsInPath[currPath]++
			graph.Ants[i].Path = graph.Paths[currPath]
		} else {
			// If we get here: next path doesn't exist or
			// Rooms in CurrentPath + Ants in CurrentPath <= Rooms in NextPath
			// Set Current Path to Path1 and assign Ant to Current Path
			currPath = 0
			graph.Ants[i].Path = graph.Paths[currPath]
			antsInPath[currPath]++
		}
	}
	graph.AntsInPaths = antsInPath
}

// MoveAnts moves all ants from source to sink in a correct order
func MoveAnts(graph *types.Graph) *[][]string {
	// resultToPriint to hold moves of ants
	var resultToPrint [][]string
	// Set next room to 1 because Room0 is the source(start)
	nextRoom := 1
	// Iterate until all ants reach the sink (end)
	for !allAntsIn(graph.Ants, graph.End) {
		var iterationToPrint []string
		// Iterate through each ant and move it if next room is free
		for i := range graph.Ants {
			// If ant has already reached the sink - skip it
			if graph.Ants[i].Position != graph.End {
				// Make sure sink is always empty
				graph.End.HasAnt = false
				// If next room in Ant's path is the sink:
				// Move Ant to sink, free up the previous room and continue
				if graph.Ants[i].Path[nextRoom] == graph.End {
					iterationToPrint = append(iterationToPrint, fmt.Sprintf("L%v-%v ", i, graph.Ants[i].Path[nextRoom].Name))
					graph.Ants[i].Position = graph.End
					graph.Ants[i].Path[nextRoom-1].HasAnt = false
					continue
				}
				// Otherwise, check the next room availabilty
				// If next room is empty: move Ant and free up the previous room and
				// cut the Ant's Path so that it always starts from the current room
				// If next room is not empty: skip the ant this time
				if !graph.Ants[i].Path[nextRoom].HasAnt {
					iterationToPrint = append(iterationToPrint, fmt.Sprintf("L%v-%v ", i, graph.Ants[i].Path[nextRoom].Name))
					graph.Ants[i].Position = graph.Ants[i].Path[nextRoom]
					graph.Ants[i].Path[nextRoom].HasAnt = true
					graph.Ants[i].Path[nextRoom-1].HasAnt = false
					graph.Ants[i].Path = graph.Ants[i].Path[nextRoom:]
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
