package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]
	statistics := getStatistics(words)

	showStatistics(statistics)
}

func getStatistics(words []string) map[string]int {
	statistics := make(map[string]int)

	for _, word := range words {
		initial := strings.ToUpper(string(word[0]))
		counter, found := statistics[initial]
		if found {
			statistics[initial] = counter + 1
		} else {
			statistics[initial] = 1
		}
	}

	return statistics
}

func showStatistics(statistics map[string]int) {
	fmt.Printf("Words count initialized in each letter: \n")
	for inital, counter := range statistics {
		fmt.Printf("%s = %d\n", inital, counter)
	}
}
