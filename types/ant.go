package types

// Ant structure with Ant ID, path to follow and position
type Ant struct {
	ID       int
	Path     []Room
	Position Room
}