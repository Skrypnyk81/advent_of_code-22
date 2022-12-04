package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	readFile, err := os.Open("calories.txt")

	if err != nil {
		fmt.Println("Errore in lettura del file", err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	maxCal := 0
	totalCal := 0

	for fileScanner.Scan() {
		number, _ := strconv.Atoi(fileScanner.Text())
		if number == 0 {
			if maxCal < totalCal {
				maxCal = totalCal
				totalCal = 0
			} else {
				totalCal = 0
			}
		} else {
			totalCal += number
		}

	}
	fmt.Println(maxCal)
	readFile.Close()
}
