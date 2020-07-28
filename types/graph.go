package types

// Neighbour structure contains connected room
// For this project, only 1 ant per vertex is aloud,
// So Capacity will always be equal to 1
type Neighbour struct {
	NeighbourRoom    Room
	Weight, Capacity int
}

// Graph structure contains a map of connected rooms to the key Room
type Graph struct {
	Roommap map[Room][]Neighbour
}

// InitGraph initializes an empty graph
func InitGraph() *Graph {
	return &Graph{
		Roommap: make(map[Room][]Neighbour),
	}
}

//GetNeighbours returns the neighbours of a given room
func (g *Graph) GetNeighbours(room Room) []Neighbour {
	return g.Roommap[room]
}

// GetRoomList returns a list of all rooms in graph
func (g *Graph) GetRoomList() []Room {
	roomlist := []Room{}
	for room := range g.Roommap {
		roomlist = append(roomlist, room)
	}
	return roomlist
}

// AddRoom adds a room with no neighbours into the map
// If the room already exists, returns graph unchanged
func (g *Graph) AddRoom(room Room) *Graph {
	if _, found := g.Roommap[room]; !found {
		null := []Neighbour{}
		g.Roommap[room] = null
	}
	return g
}

// AddNeighbour adds connection with no weight between a given source and destination rooms
// if the rooms are the same/don't exist, or if the connection already exists, returns graph unchanged
func (g *Graph) AddNeighbour(source Room, destination Room, capacity int) *Graph {
	// Check if the rooms are the same
	if source == destination {
		return g
	}
	// Check if rooms exist
	if _, found := g.Roommap[source]; !found {
		return g
	} else if _, found := g.Roommap[destination]; !found {
		return g
	}
	// Check if the connection already exists
	neighbourlist := g.GetNeighbours(source)
	for neighbour := range neighbourlist {
		if neighbourlist[neighbour].NeighbourRoom == destination {
			return g
		}
	}
	// At this point: both rooms exist in the graph, but
	// there is no connection between them.
	// Connect the rooms
	neighbour := Neighbour{destination, 0, capacity}
	g.Roommap[source] = append(g.Roommap[source], neighbour)
	neighbour = Neighbour{source, 0, (capacity)}
	g.Roommap[destination] = append(g.Roommap[destination], neighbour)
	return g
}

//UpdateWeight is a graph-associated method that updates the weight of the connection
//Returns the graph unchanged if the new weight is greater than the connection's capacity.
func (g *Graph) UpdateWeight(source Room, destination Room, weight int) *Graph {
	neighbourlistsource := g.GetNeighbours(source)
	for i := 0; i < len(neighbourlistsource); i++ {
		//Iterate until we find the destination...
		if neighbourlistsource[i].NeighbourRoom == destination {
			//Check to make sure the weight is not greater than
			//the capacity.
			if Abs(weight) <= Abs(neighbourlistsource[i].Capacity) {
				//Update the connection with the new weight.
				updatedconnect := Neighbour{destination, weight, neighbourlistsource[i].Capacity}
				g.Roommap[source][i] = updatedconnect
			}
		}
	}
	return g
}
