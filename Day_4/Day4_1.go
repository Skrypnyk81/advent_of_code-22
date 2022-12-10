package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func extractDigit(line string) []int {
	re := regexp.MustCompile("[0-9]+")
	result := re.FindAllString(line, -1)
	ints := make([]int, len(result))

	for i, s := range result {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func main() {

	readFile, err := os.Open("pairs.txt")

	if err != nil {
		fmt.Println("Errore in lettura del file", err)
	}
	scanner := bufio.NewScanner(readFile)
	total := 0

	for scanner.Scan() {
		lineSector := scanner.Text()
		numberSectors := extractDigit(lineSector)
		firstEqual := numberSectors[0] <= numberSectors[2] && numberSectors[1] >= numberSectors[3]
		secondEqual := numberSectors[0] >= numberSectors[2] && numberSectors[1] <= numberSectors[3]
		if firstEqual || secondEqual {
			total++
		}
	}
	fmt.Println(total)
	err = readFile.Close()
	if err != nil {
		fmt.Println("Errore nella chiusura del file", err)
	}
}
