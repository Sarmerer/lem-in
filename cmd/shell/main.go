package main

import (
	"fmt"
	"lem-in/parser"
	"lem-in/solver"
)

func main() {
	//TO-DO: wtf is start and end in data????
	_, graph := parser.ParseFile("example00.txt")
	rooms := graph.GetRoomList()
	paths := solver.EdmondsKarp(graph, rooms[0], rooms[len(rooms)-1])
	fmt.Println("Paths: ", paths)
}
