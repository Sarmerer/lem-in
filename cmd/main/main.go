package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		res, _ := strconv.Atoi(txt)
		fmt.Println("Index:", index, "Text:", res, "Length:", len(txt))
	}
}
