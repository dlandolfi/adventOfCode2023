package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// make valid-game map
	m := make(map[string]int)
	m["red"] = 12
	m["green"] = 13
	m["blue"] = 14

	var sum int
	filePath := "2-input"

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) // Trim spaces
		parts := strings.Split(line, ":")
		gameId, subsets := parts[0], strings.Split(parts[1], ";")
		gameNum, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSpace(gameId), "Game "))
		subsetList := make([][]string, len(subsets))

		for i, v := range subsets {
			subsetList[i] = strings.Split(strings.TrimSpace(v), ", ") // Trim spaces
		}

		isValid := compareSets(m, subsetList)

		if isValid {
			sum = sum + gameNum
		}
	}
	fmt.Println(sum)
}

func compareSets(counts map[string]int, subsets [][]string) bool {
	for _, subset := range subsets {
		cubeCounts := make(map[string]int)
		for _, cube := range subset {
			parts := strings.Split(cube, " ")
			color := parts[1]
			count, _ := strconv.Atoi(parts[0])
			cubeCounts[color] += count
		}

		for color, count := range cubeCounts {
			if count > counts[color] {
				return false
			}
		}
	}
	return true
}
