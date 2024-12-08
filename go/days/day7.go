package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day7() {
	fmt.Println("Solution for Day 7, Part 1:")
	solveDay7Part1()
	fmt.Println("Solution for Day 7, Part 2:")
	solveDay7Part2()
}

func solveDay7Part1() {
	file, err := os.Open("days/7input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var tests [][]int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		var currentTest []int

		firstNum, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		currentTest = append(currentTest, firstNum)

		numStrs := strings.Fields(parts[1])
		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			currentTest = append(currentTest, num)
		}
		tests = append(tests, currentTest)
	}
	sum := 0
	for _, test := range tests {
		var prefix []int
		prefix = append(prefix, test[1])
		for i := 2; i < len(test); i++ {
			var current []int
			for _, tmp := range prefix {
				plus := tmp + test[i]
				mal := tmp * test[i]
				current = append(current, plus, mal)
				prefix = current
			}
		}
		for _, res := range prefix {
			if res == test[0] {
				sum += test[0]
				break
			}
		}
	}
	fmt.Println("Sum =", sum)
}

func solveDay7Part2() {
	file, err := os.Open("days/7input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var tests [][]int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		var currentTest []int

		firstNum, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		currentTest = append(currentTest, firstNum)

		numStrs := strings.Fields(parts[1])
		for _, numStr := range numStrs {
			num, _ := strconv.Atoi(numStr)
			currentTest = append(currentTest, num)
		}
		tests = append(tests, currentTest)
	}

	sum := 0
	for _, test := range tests {
		var prefix []int
		prefix = append(prefix, test[1])

		for i := 2; i < len(test); i++ {
			var current []int
			for _, tmp := range prefix {
				// Addition und Multiplikation
				plus := tmp + test[i]
				mal := tmp * test[i]
				current = append(current, plus, mal)

				// Konkatenation
				concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", tmp, test[i]))
				current = append(current, concat)
			}
			prefix = current
		}

		// Prüfen, ob eines der Ergebnisse gleich der ersten Zahl ist
		for _, res := range prefix {
			if res == test[0] {
				sum += test[0]
				break
			}
		}
	}

	fmt.Println("Sum (Part 2) =", sum)
}
