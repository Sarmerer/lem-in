package utils

import (
	"encoding/json"
	"fmt"
	"lem-in/types"
	"os"
)

type response struct {
	Nodes []nodesStruct `json:"nodes,omitempty"`
	Edges []edgesStruct `json:"edges,omitempty"`
	Paths []pathStruct  `json:"paths,omitempty"`
	Ants  int           `json:"ants,omitempty"`

	PathsCount int `json:"paths_count,omitempty"`
}

type dataStruct struct {
	//for nodes
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`

	//for edges
	Source string `json:"source,omitempty"`
	Target string `json:"target,omitempty"`

	//for paths
	Ants  int      `json:"ants,omitempty"`
	Nodes []string `json:"nodes,omitempty"`
}

type nodesStruct struct {
	Data dataStruct `json:"data,omitempty"`
}

type edgesStruct struct {
	Data dataStruct `json:"data,omitempty"`
}

type pathStruct struct {
	ID   int      `json:"id"`
	Ants int      `json:"ants,omitempty"`
	Path []string `json:"nodes,omitempty"`
}

func Marshal(graph *types.Graph) {
	path := "../visualizer/static/data.json"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Unable to create file.", err)
		os.Exit(1)
	}
	res := response{
		Ants:       graph.AntsAmount,
		PathsCount: len(graph.Paths),
	}
	var counter int
	for parent, rooms := range graph.Roommap {
		var node nodesStruct
		var edge edgesStruct
		if parent.Name == graph.Start.Name {
			node.Data = dataStruct{ID: parent.Name, Type: "start"}
		} else if parent.Name == graph.End.Name {
			node.Data = dataStruct{ID: parent.Name, Type: "end"}
		} else {
			node.Data = dataStruct{ID: parent.Name}
		}
		res.Nodes = append(res.Nodes, node)
		for _, room := range rooms {
			edge.Data.ID = fmt.Sprint("edge", counter)
			edge.Data.Source = parent.Name
			edge.Data.Target = room.NeighbourRoom.Name
			res.Edges = append(res.Edges, edge)
			counter++
		}
	}
	counter = 0
	for _, ant := range graph.Ants {
		var path []string
		for _, i := range ant.Path {
			path = append(path, i.Name)
		}
		if index, found := pathAlreadyFound(&res.Paths, path); found {
			res.Paths[index].Ants++
		} else {
			res.Paths = append(res.Paths, pathStruct{ID: counter, Ants: 1, Path: path})
			counter++
		}
	}
	r, err := json.Marshal(res)
	if err != nil {
		fmt.Println("Unable to marshal JSON.", err)
		os.Exit(1)
	}
	_, err = file.Write(r)
	if err != nil {
		fmt.Println("Unable to write file.", err)
		os.Exit(1)
	}
}

func pathAlreadyFound(paths *[]pathStruct, path []string) (int, bool) {
	for index, p := range *paths {
		counter := 0
		if len(p.Path) != len(path) {
			return 0, false
		}
		for i, node := range p.Path {
			if node == path[i] {
				counter++
			}
		}
		if counter == len(path) {
			return index, true
		}
	}
	return 0, false
}
