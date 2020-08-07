package config

const (
	PathCustomMap = "../maps/custom.txt"

	MessageCustomMap = "You can use a map that already exists by specifying it as a program argument\nor\nEnter map data here\nPress Ctrl+D when finished"
	MessageElapsed   = "\nElapsed: %vs\n"
	MessageTurns     = "\nTurns:"
	MessageLines     = "\nLines:"

	ErrorManyArgs = "Too many arguments"

	ErrorBase      = "Invalid input: %v\n"
	ErrorBaseExact = "Invalid input at line #%v: %v\n"

	ErrorAnts        = "ivalid ants amount"
	ErrorRoom        = "ivalid room params"
	ErrorLink        = "invalid link params"
	ErrorSore        = "invalid %s params" //sore == start or end
	ErrorAnotherSore = "another %s declaration"

	ErrorNoPaths      = "no valid paths were found"
	ErrorNoData       = "no data provided"
	ErrorNoStart      = "no start room"
	ErrorNoEnd        = "no end room"
	ErrorNoSoreCoords = "no %s coords"
)
