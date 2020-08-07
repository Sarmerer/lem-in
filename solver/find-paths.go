// algorithm.go
// Edmonds-Karp Algorithm with BFS to find valid paths in the graph.

package solver

import (
	"lem-in/types"
)

// EdmondsKarp accepts a graph with a valid source and sink room,
// and returns the list of all valid paths in the graph
func EdmondsKarp(g *types.Graph, source types.Room, sink types.Room) [][]types.Room {
	// Initialize the list of paths
	var pathList [][]types.Room
	// Run BFS until no valid path is found
	for {
		pathcap, pathmap := BreadthFirstSearch(g, source, sink)
		if pathcap == 0 {
			break
		}
		// Backtrack search through the graph to record path and update weight.
		room := sink
		// Initialize path that will contain path from sink to source
		revPath := []types.Room{room}
		for room != source {
			// Grab the parent of the room.
			parent := pathmap[room]
			revPath = append(revPath, parent)
			// Update weight of the connection.
			g.UpdateWeight(parent, room, pathcap)
			// Set the room equal to the parent to continue backtrack.
			room = parent
		}
		// Order path in ascending order and add it to the pathList
		addPath(revPath, &pathList)
	}
	return pathList
}

// BreadthFirstSearch accepts a graph with a valid source
// and sink and returns a valid path
func BreadthFirstSearch(g *types.Graph, source types.Room, sink types.Room) (int, map[types.Room]types.Room) {
	// Create a map in which rooms have room keys corresponding to a
	// parent/source to child/destination relationship.  This will be the
	// path returned.
	roomlist := g.GetRoomList()
	length := len(roomlist)
	path := make(map[types.Room]types.Room, length)
	// Initialize a room which can be used to tell if a room has been
	// discovered yet or not, and give every room that key to begin.
	notvisited := types.Room{Name: "", X: -1, Y: -1, HasAnt: false}
	for _, room := range roomlist {
		path[room] = notvisited
	}
	// Give the source a different key to ensure it is not rediscovered.
	path[source] = types.Room{Name: "", X: -2, Y: -2, HasAnt: false}
	// Initialize a queue and enqueue the source room.
	q := GenQueue(0)
	q.Enqueue(source)
	// Loop until the queue is empty.
	for q.GetSize() > 0 {
		// Grab the first room in the queue and check all neighbours
		// until one is found where flow can be pushed.
		u := q.Dequeue()
		for _, v := range g.GetNeighbours(u) {
			// If there is available capacity and the neighbour
			// has not been visited yet...
			if path[v.NeighbourRoom] == notvisited && (v.Capacity-v.Weight > 0 || (v.Weight < 0 && v.Capacity < 0)) {
				// Path can proceed from u to v.
				// Set u to be the parent of v.
				path[v.NeighbourRoom] = u

				if v.NeighbourRoom != sink {
					// We have not reached the sink. We
					// enqueue v.NeighbourRoom and
					// continue.
					q.Enqueue(v.NeighbourRoom)
				} else {
					// We have reached the sink and we
					// return.
					return 1, path
				}
			}
		}
	}
	// No paths were found, so we return 0 and whatever path was built.
	return 0, path
}

//addPath orders revPath in ascending order and adds it to the pathList
func addPath(revPath []types.Room, pathList *[][]types.Room) {
	path := make([]types.Room, len(revPath))
	counter := 0
	for i := len(path) - 1; i >= 0; i-- {
		path[counter] = revPath[i]
		counter++
	}
	*pathList = append(*pathList, path)
}
