package utils

import (
	"encoding/json"
	"fmt"
	"lem-in/types"
	"os"
)

type responseStruct struct {
	Nodes []nodesStruct `json:"nodes"`
	Edges []edgesStruct `json:"edges"`
	Paths []pathStruct  `json:"paths"`
	Ants  int           `json:"ants"`
}
type nodesStruct struct {
	Data nodeStruct `json:"data"`
}

type nodeStruct struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type,omitempty"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}

type edgesStruct struct {
	Data edgeStruct `json:"data"`
}

type edgeStruct struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type pathStruct struct {
	ID    int      `json:"id"`
	Ants  int      `json:"ants"`
	Color string   `json:"color"`
	Nodes []string `json:"nodes"`
	Edges []string `json:"edges"`
}

func Marshal(graph *types.Graph) {
	filePath := "../visualizer/static/data.json"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Unable to create file.", err)
		os.Exit(1)
	}
	res := responseStruct{
		Ants: graph.AntsAmount,
	}
	var counter int
	for parent, rooms := range graph.Roommap {
		node := nodesStruct{nodeStruct{ID: parent.Name, Label: parent.Name, X: parent.X, Y: parent.Y}}
		if parent.Name == graph.Start.Name {
			node.Data.Label += "(start)"
			node.Data.Type = "start"
		} else if parent.Name == graph.End.Name {
			node.Data.Label += "(end)"
			node.Data.Type = "end"
		}
		res.Nodes = append(res.Nodes, node)
		for _, room := range rooms {
			if !edgeUsed(parent.Name, room.NeighbourRoom.Name, res.Edges) {
				res.Edges = append(res.Edges, edgesStruct{
					edgeStruct{
						ID:     fmt.Sprint("e", counter),
						Source: parent.Name,
						Target: room.NeighbourRoom.Name,
					},
				})
				counter++
			}
		}
	}
	counter = 0
	for index, path := range graph.Paths {
		var p []string
		var e []string
		for _, room := range path {
			p = append(p, room.Name)
			if len(p) >= 2 {
				e = append(e, getEdge(p[len(p)-2], p[len(p)-1], res.Edges))
			}
		}
		res.Paths = append(res.Paths, pathStruct{
			ID:    counter,
			Ants:  graph.AntsInPaths[index],
			Nodes: p,
			Edges: e,
		})
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

func getEdge(source, target string, edges []edgesStruct) string {
	var res string
	for _, edge := range edges {
		if edge.Data.Source == source &&
			edge.Data.Target == target ||
			edge.Data.Source == target &&
				edge.Data.Target == source {
			res = edge.Data.ID
			break
		}
	}
	return res
}

func edgeUsed(source, target string, edges []edgesStruct) bool {
	for _, edge := range edges {
		if (edge.Data.Source == source || edge.Data.Source == target) && (edge.Data.Target == source || edge.Data.Target == target) {
			return true
		}
	}
	return false
}
