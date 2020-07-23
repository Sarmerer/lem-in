package types

type Data struct {
	Rooms      []RoomStruct
	AntsAmount int
	Start      string
	End        string
}

type RoomStruct struct {
	Name        string
	CoordX      int
	CoordY      int
	Connections []string
}

type antStruct struct {
	ID       int
	Position string
}
