package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	readFile, err := os.Open("rucksack.txt")

	if err != nil {
		fmt.Println("Errore in lettura del file", err)
	}
	scanner := bufio.NewScanner(readFile)
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	total := 0
	for scanner.Scan() {
		rucksack := scanner.Text()
		midString := len(rucksack) / 2
		sub1 := rucksack[:midString]
		sub2 := rucksack[midString:]
		check := ""
		for word := range sub1 {
			res1 := strings.Contains(sub2, string(sub1[word]))
			res2 := strings.Contains(check, string(sub1[word]))
			if res1 {
				if res2 == false {
					total += strings.Index(alphabet, string(sub1[word])) + 1
					check += string(sub1[word])
				}
			}
		}
		check = ""
	}
	fmt.Println(total)
	err = readFile.Close()
	if err != nil {
		fmt.Println("Errore nella chiusura del file", err)
	}
}
