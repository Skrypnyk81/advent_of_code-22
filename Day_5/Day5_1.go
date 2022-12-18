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
func createList(list []string) []string {
	count := 0
	var listTest []string
	for _, letter := range list {
		if string(letter) != " " {
			listTest = append(listTest, letter)
			count = 0
		} else if count == 4 {
			listTest = append(listTest, letter)
			count = 0
		}
		count += 1
	}
	return listTest
}

func addToMaps(map1 map[int][]string, list []string) map[int][]string {

	for index, letter := range list {
		if string(letter) != " " {
			map1[index+1] = append([]string{letter}, map1[index+1]...)
		}
	}
	return map1
}

func main() {
	crates := map[int][]string{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
		9: {},
	}
	readFile, err := os.Open("crates.txt")

	if err != nil {
		fmt.Println("Errore in lettura del file", err)
	}
	scanner := bufio.NewScanner(readFile)
	//total := 0
	var mapsCrates map[int][]string
	var numberMoves []int
	for scanner.Scan() {
		lineSector := scanner.Text()
		re := regexp.MustCompile(`[A-Z\s]`)
		result := re.FindAllString(lineSector, -1)
		listRow := createList(result)
		mapsCrates = addToMaps(crates, listRow)

		if lineSector != "" {
			if string(lineSector[:4]) == "move" {
				numberMoves = extractDigit(lineSector)
				step := numberMoves[0]
				from := mapsCrates[numberMoves[1]]
				to := mapsCrates[numberMoves[2]]
				for i := 0; i < step; i++ {
					to, from = append(to, from[len(from)-1]), from[:len(from)-1]
				}
				mapsCrates[numberMoves[1]] = from
				mapsCrates[numberMoves[2]] = to
			}
		}
	}
	var result string
	for i := 1; i < 10; i++ {
		value := mapsCrates[i]
		result += value[len(value)-1]
	}
	fmt.Println(result)
	err = readFile.Close()
	if err != nil {
		fmt.Println("Errore nella chiusura del file", err)
	}
}
