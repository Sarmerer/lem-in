package types

type Data struct {
	AntsAmount int
	Start,
	End Room
}

// Room structure with name, (x, y) coordinates, and HasAnt bool to check if the room is empty.
type Room struct {
	Name   string
	X, Y   int
	HasAnt bool
}

// Ant structure with Ant ID, path to follow and position
type Ant struct {
	ID       int
	Path     []Room
	Position Room
}
