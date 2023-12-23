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

	// opens a file
	f, err := os.Open("1-input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close() // defers closing file until main() is done

	scanner1 := bufio.NewScanner(f)
	scanner2 := bufio.NewScanner(f)

	partOneValue := partOne(scanner1)
	fmt.Println("Part One:", partOneValue)

	// reset file pointer back to beginning
	f.Seek(0, 0)

	partTwoValue := partTwo(scanner2)
	fmt.Println("Part Two:", partTwoValue)
}

func partOne(scanner *bufio.Scanner) int {
	var sum int

	// reads line by line
	for scanner.Scan() {
		value := getLineCalValue(scanner.Text())
		sum = sum + value
	}
	return sum
}

func partTwo(scanner *bufio.Scanner) int {
	var sum int
	var newLine string

	for scanner.Scan() {
		line := scanner.Text()

		// replace all instances of number words with the number
		// adding the first and last letter to the replacement value
		// in case adjacent words share a letter. i.e. eightwo
		newLine = strings.Replace(line, "zero", "z0o", -1)
		newLine = strings.Replace(newLine, "one", "o1e", -1)
		newLine = strings.Replace(newLine, "two", "t2o", -1)
		newLine = strings.Replace(newLine, "three", "t3e", -1)
		newLine = strings.Replace(newLine, "four", "f4r", -1)
		newLine = strings.Replace(newLine, "five", "f5e", -1)
		newLine = strings.Replace(newLine, "six", "s6x", -1)
		newLine = strings.Replace(newLine, "seven", "s7n", -1)
		newLine = strings.Replace(newLine, "eight", "e8t", -1)
		newLine = strings.Replace(newLine, "nine", "n9e", -1)

		// parse newly created list
		value := getLineCalValue(newLine)
		sum = sum + value
	}
	return sum
}

func getLineCalValue(str string) int {
	var value1, value2 int
	// find the very first digit, then break
	for i := 0; i < len(str); i++ {
		if str[i] >= 48 && str[i] <= 57 {
			// Go returns type byte when indexing strings
			// ascii value gets converted to string, then converted to int
			value1, _ = strconv.Atoi((string(str[i])))
			break
		}
	}
	// find the very last digit, then break
	for j := len(str) - 1; j >= 0; j-- {
		if str[j] >= 48 && str[j] <= 57 {
			value2, _ = strconv.Atoi((string(str[j])))
			break
		}
	}
	return value1*10 + value2
}
