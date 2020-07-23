package parser

type data struct {
	Rooms      []roomStruct
	AntsAmount int
	Start      string
	End        string
}

type roomStruct struct {
	Name        string
	CoordX      int
	CoordY      int
	Connections []string
}

type antStruct struct {
	ID       int
	Position string
}
