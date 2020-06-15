package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var parsedText []string

	file, err := os.Open("example00.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedText = append(parsedText, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for index, txt := range parsedText {
		fmt.Println("Index:", index, "Text:", txt, "Length:", len(txt))
	}
}
