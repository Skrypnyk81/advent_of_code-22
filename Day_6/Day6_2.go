package main

import (
	"bufio"
	"fmt"
	"os"
)

func unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}

func main() {

	readFile, err := os.Open("signals.txt")

	if err != nil {
		fmt.Println("Errore in lettura del file", err)
	}
	scanner := bufio.NewScanner(readFile)

	for scanner.Scan() {
		lineSector := scanner.Text()
		for i := 0; i < len(lineSector); i++ {
			subString := lineSector[i : i+14]
			check := unique(subString)
			if check == true {
				fmt.Println(i + 14)
				break
			}
		}
		err = readFile.Close()
		if err != nil {
			fmt.Println("Errore nella chiusura del file", err)
		}
	}
}
