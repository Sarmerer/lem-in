package utils

import (
	"encoding/json"
	"fmt"
	"lem-in/types"
	"os"
)

type response struct {
	Nodes []node       `json:"nodes,omitempty"`
	Edges []edge       `json:"edges,omitempty"`
	Paths []pathStruct `json:"paths,omitempty"`
	Ants  int          `json:"ants,omitempty"`

	PathsCount int `json:"paths_count,omitempty"`
}

type node struct {
	ID    string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Type  string `json:"type,omitempty"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}

type edge struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
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
		n := node{ID: parent.Name, Label: parent.Name, X: parent.X, Y: parent.Y}
		if parent.Name == graph.Start.Name {
			n.Label = "start"
		} else if parent.Name == graph.End.Name {
			n.Label = "end"
		}
		res.Nodes = append(res.Nodes, n)
		for _, room := range rooms {
			if !edgeUsed(parent.Name, room.NeighbourRoom.Name, res.Edges) {
				res.Edges = append(res.Edges, edge{From: parent.Name, To: room.NeighbourRoom.Name})
			}
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

func edgeUsed(from, to string, edges []edge) bool {
	for _, edge := range edges {
		if (edge.From == from || edge.From == to) && (edge.To == from || edge.To == to) {
			return true
		}
	}
	return false
}
