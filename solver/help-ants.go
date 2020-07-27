package solver

import (
	"fmt"
	"lem-in/types"
)

// InitAntsAndAssignPaths initialazes Ants and assigns an optimal path to each Ant
// Use the following logic to calculate an optimal path:
// If Rooms in Path1 + Ants in Path1 > Rooms in Path2 send Ant to Path2
// Otherwise, send Ant to Path1
func InitAntsAndAssignPaths(data *types.Data, graph *types.Graph) []types.Ant {
	paths := EdmondsKarp(graph, data.Start, data.End)
	ants := make([]types.Ant, data.AntsAmount)
	for i := 0; i < data.AntsAmount; i++ {
		ants[i].ID = i
		ants[i].Position = data.Start
	}
	//ASSIGN PATHS TO ANTS
	antsInPath := make([]int, len(paths))
	ants[0].Path = paths[0]
	antsInPath[0]++
	for i, currPath := 1, 0; i < len(ants); i++ {
		if (currPath+1) < len(paths) &&
			(len(paths[currPath])+antsInPath[currPath]) > len(paths[currPath+1]) {
			currPath++
			antsInPath[currPath]++
			ants[i].Path = paths[currPath]
		} else {
			currPath = 0
			ants[i].Path = paths[currPath]
			antsInPath[currPath]++
		}
	}

	return ants
}

// MoveAnts moves all ants from source to sink in a correct order
func MoveAnts(ants []types.Ant, sink types.Room) {
	currPath := 1
	//MOVE ANTS
	for !allAntsIn(ants, sink) {
		for i := range ants {
			if ants[i].Position != sink {
				sink.HasAnt = false
				if ants[i].Path[currPath] == sink {
					fmt.Printf("L%v-%v\t", i, ants[i].Path[currPath].Name)
					ants[i].Position = sink
					ants[i].Path[currPath-1].HasAnt = false
					continue
				}
				if !ants[i].Path[currPath].HasAnt {
					fmt.Printf("L%v-%v\t", i, ants[i].Path[currPath].Name)
					ants[i].Position = ants[i].Path[currPath]
					ants[i].Path[currPath].HasAnt = true
					ants[i].Path[currPath-1].HasAnt = false
					ants[i].Path = ants[i].Path[currPath:]
				}
			}
		}
		fmt.Println("")
	}
}

func allAntsIn(ants []types.Ant, sink types.Room) bool {
	for _, ant := range ants {
		if ant.Position != sink {
			return false
		}
	}
	return true
}
