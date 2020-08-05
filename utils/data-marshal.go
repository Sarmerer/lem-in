package utils

import (
	"encoding/json"
	"fmt"
	"lem-in/types"
	"os"
)

type response struct {
	Nodes []node       `json:"nodes"`
	Edges []edge       `json:"edges"`
	Paths []pathStruct `json:"paths"`
	Ants  int          `json:"ants"`

	PathsCount int `json:"paths_count"`
}

type node struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type,omitempty"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
	Size  int    `json:"size"`
}

type edge struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type pathStruct struct {
	ID   int      `json:"id"`
	Ants int      `json:"ants"`
	Path []string `json:"nodes"`
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
		n := node{ID: parent.Name, Label: parent.Name, X: parent.X, Y: parent.Y, Size: 7}
		if parent.Name == graph.Start.Name {
			n.Label += "(start)"
			n.Size = 10
		} else if parent.Name == graph.End.Name {
			n.Size = 10
			n.Label += "(end)"
		}
		res.Nodes = append(res.Nodes, n)
		for _, room := range rooms {
			if !edgeUsed(parent.Name, room.NeighbourRoom.Name, res.Edges) {
				res.Edges = append(res.Edges, edge{
					ID:     fmt.Sprint("e", counter),
					Source: parent.Name,
					Target: room.NeighbourRoom.Name,
				})
				counter++
			}
		}
	}
	counter = 0
	for index, path := range graph.Paths {
		var p []string
		for _, room := range path {
			p = append(p, room.Name)
		}
		res.Paths = append(res.Paths, pathStruct{ID: counter, Ants: graph.AntsInPaths[index], Path: p})
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

func edgeUsed(from, to string, edges []edge) bool {
	for _, edge := range edges {
		if (edge.Source == from || edge.Source == to) && (edge.Target == from || edge.Target == to) {
			return true
		}
	}
	return false
}
