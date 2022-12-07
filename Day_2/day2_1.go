package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	readFile, err := os.Open("scissors.txt")

	if err != nil {
		fmt.Println("Errore in lettura del file", err)
	}
	scanner := bufio.NewScanner(readFile)
	total := 0
	myScore := map[string]int{
		"Y": 2,
		"Z": 3,
		"X": 1,
	}
	myPair := map[string]int{
		"A X": 3,
		"A Y": 6,
		"B Y": 3,
		"B Z": 6,
		"C Z": 3,
		"C X": 6,
	}

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		total += myScore[words[1]]
		listItem := strings.Join(words, " ")
		total += myPair[listItem]
	}

	fmt.Println(total)
	err = readFile.Close()
	if err != nil {
		fmt.Println("Errore nella chiusura del file", err)
	}
}

// A, X = Rock
// B, Y = paper
// C, Z = Scissors
