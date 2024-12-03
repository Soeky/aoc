package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Day3() {
	fmt.Println("Solution for Day 2, Part 1:")
	solveDay3Part1()
	fmt.Println("Solution for Day 2, Part 2:")
	solveDay3Part2()
}

func solveDay3Part1() {
	file, err := os.Open("days/3input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			first, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("error with match[1]:", match[1])
			}
			second, _ := strconv.Atoi(match[2])
			sum += first * second
		}
	}
	fmt.Println("sum = ", sum)
}

func solveDay3Part2() {
	file, err := os.Open("days/3input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllString(line, -1)
		do := true
		for _, match := range matches {
			if match == "do()" {
				do = true
			} else if match == "don't()" {
				do = false
			} else if do {
				sum += giveMul(match)
			}
		}
	}
	fmt.Println("sum = ", sum)
}

func giveMul(s string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(s, -1)

	if len(matches) == 0 {
		fmt.Println("No match found")
		return 0 // Rückgabewert 0, wenn kein Match gefunden wird
	}

	// Nimm das erste Match
	match := matches[0]

	// Konvertiere die Gruppen zu Integers
	first, err1 := strconv.Atoi(match[1])
	second, err2 := strconv.Atoi(match[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Error converting numbers")
		return 0
	}

	// Gib das Produkt zurück
	return first * second
}
