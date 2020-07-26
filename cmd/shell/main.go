package main

import (
	"fmt"
	"lem-in/parser"
)

func main() {
	fmt.Println(parser.ParseFile("../maps/example00.txt"))
}
