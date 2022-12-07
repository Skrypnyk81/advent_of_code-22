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
		"B": 2,
		"C": 3,
		"A": 1,
	}
	myScoreNormal := map[string]int{
		"Y": 2,
		"Z": 3,
		"X": 1,
	}

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if words[1] == "Y" {
			total += myScore[words[0]]
			total += 3
		} else if words[1] == "X" {
			if words[0] == "A" {
				total += myScoreNormal["Z"]
			} else if words[0] == "B" {
				total += myScoreNormal["X"]
			} else if words[0] == "C" {
				total += myScoreNormal["Y"]
			}

		} else if words[1] == "Z" {
			if words[0] == "A" {
				total += myScoreNormal["Y"]
			} else if words[0] == "B" {
				total += myScoreNormal["Z"]
			} else if words[0] == "C" {
				total += myScoreNormal["X"]
			}
			total += 6
		}

	}

	fmt.Println(total)
	err = readFile.Close()
	if err != nil {
		fmt.Println("Errore nella chiusura del file", err)
	}
}
