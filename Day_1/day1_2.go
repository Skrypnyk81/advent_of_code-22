package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	readFile, err := os.Open("calories.txt")

	if err != nil {
		fmt.Println("Errore in lettura del file", err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	totalCal := 0
	var allCalories []int

	for fileScanner.Scan() {
		number, _ := strconv.Atoi(fileScanner.Text())
		if number == 0 {
			totalCal += number
			allCalories = append(allCalories, totalCal)
			totalCal = 0
		} else {
			totalCal += number
		}

	}
	sort.Slice(allCalories, func(i, j int) bool {
		return allCalories[i] > allCalories[j]
	})
	result := 0
	for _, v := range allCalories[:3] {
		result += v
	}
	fmt.Println(result)
	readFile.Close()
}
