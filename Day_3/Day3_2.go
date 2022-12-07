package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countCommon(bag []string) string {
	commmon := ""
	sub := bag[0]
	sub1 := bag[1]
	sub2 := bag[2]
	for value := range sub {
		res1 := strings.Contains(sub1, string(sub[value]))
		res2 := strings.Contains(sub2, string(sub[value]))
		if res1 && res2 {
			checkDouble := strings.Contains(commmon, string(sub[value]))
			if checkDouble == false {
				commmon += string(sub[value])
			}
		}
	}
	return commmon
}

func main() {

	readFile, err := os.Open("rucksack.txt")

	if err != nil {
		fmt.Println("Errore in lettura del file", err)
	}
	scanner := bufio.NewScanner(readFile)
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	total := 0
	var listThree []string
	for scanner.Scan() {
		rucksack := scanner.Text()
		listThree = append(listThree, rucksack)
		if len(listThree) == 3 {
			char := countCommon(listThree)
			total += strings.Index(alphabet, char) + 1
			listThree = []string{}
		}
	}
	fmt.Println(total)
	err = readFile.Close()
	if err != nil {
		fmt.Println("Errore nella chiusura del file", err)
	}
}
