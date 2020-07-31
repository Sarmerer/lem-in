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

type nodeStruct struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

type edgeStruct struct {
	ID     string `json:"id,omitempty"`
	Source string `json:"spurce,omitempty"`
	Target string `json:"target,omitempty"`
}
type pathStruct struct {
	ID    string   `json:"id,omitempty"`
	Ants  int      `json:"ants,omitempty"`
	Nodes []string `json:"nodes,omitempty"`
}

type nodesStruct struct {
	Data nodeStruct `json:"data,omitempty"`
}

type edgesStruct struct {
	Data edgeStruct `json:"data,omitempty"`
}

func Marshal(data *types.Data, graph *types.Graph, paths *[][]types.Room) {
	path := "../visualizer/static/test.json"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Unable to create file.", err)
		os.Exit(1)
	}
	res := response{
		Ants:       data.AntsAmount,
		PathsCount: len(*paths),
	}
	var counter int
	for parent, rooms := range graph.Roommap {
		node := nodesStruct{}
		edge := edgesStruct{}
		if parent.Name == data.Start.Name {
			node.Data = nodeStruct{ID: parent.Name, Type: "start"}
		} else if parent.Name == data.End.Name {
			node.Data = nodeStruct{ID: parent.Name, Type: "end"}
		} else {
			node.Data = nodeStruct{ID: parent.Name}
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
	for _, path := range *paths {
		p := pathStruct{}
		p.ID = fmt.Sprint("path", counter)
		for _, i := range path {
			p.Nodes = append(p.Nodes, i.Name)
		}
		res.Paths = append(res.Paths, p)
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
