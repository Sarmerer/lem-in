package types

// Ant structure with Ant ID, path to follow and position
type Ant struct {
	ID       int
	Path     []Room
	Position Room
}

// InitAnts initializes an array of ants of size antsAmount
// and assigns each ant ID and a start(source) position
func InitAnts(antsAmount int, source Room) []Ant {
	ants := make([]Ant, antsAmount)
	for i := 0; i < antsAmount; i++ {
		ants[i].ID = i
		ants[i].Position = source
	}
	return ants
}

// AssignPathToAnts assigns an optimal path to each Ant
// Use the following logic to calculate an optimal path:
// If Rooms in Path1 + Ants in Path1 > Rooms in Path2 send Ant to Path2
// Otherwise, send Ant to Path1
func AssignPathToAnts() {

}

// MoveAnts moves all ants from source to sink in a correct order
func MoveAnts() {

}
