package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day5() {
	fmt.Println("Solution for Day 5, Part 1:")
	solveDay5Part1()
	fmt.Println("Solution for Day 5, Part 2:")
	solveDay5Part2()
}

func solveDay5Part1() {
	file, err := os.Open("days/5input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	inputLines := false
	rules := make(map[int][]int)
	var tests [][]int
	index := 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if len(line) == 0 {
			inputLines = true
			continue
		}
		if !inputLines {
			parts := strings.Split(line, "|")
			firstNum, _ := strconv.Atoi(parts[0])
			secondNum, _ := strconv.Atoi(parts[1])
			rules[secondNum] = append(rules[secondNum], firstNum)
		} else {
			parts := strings.Split(line, ",")
			tests = append(tests, make([]int, len(parts)))
			for i, val := range parts {
				num, _ := strconv.Atoi(val)
				tests[index][i] = num
			}
			index++
		}
	}
	for i := 0; i < len(tests); i++ {
		isSafe := true
		for j := 0; j < len(tests[i]); j++ {
			currentNumber := tests[i][j]
			ruleForNumber := rules[currentNumber]

			for k := j + 1; k < len(tests[i]); k++ {
				nextNumber := tests[i][k]
				for _, ruleValue := range ruleForNumber {
					if nextNumber == ruleValue {
						isSafe = false
						break
					}
				}
			}
		}

		if isSafe {
			sum += tests[i][len(tests[i])/2]
		}
	}
	fmt.Println(sum)
}

func solveDay5Part2() {
	file, err := os.Open("days/5input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	inputLines := false
	rules := make(map[int][]int)
	var tests [][]int
	index := 0
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if len(line) == 0 {
			inputLines = true
			continue
		}
		if !inputLines {
			parts := strings.Split(line, "|")
			firstNum, _ := strconv.Atoi(parts[0])
			secondNum, _ := strconv.Atoi(parts[1])
			rules[secondNum] = append(rules[secondNum], firstNum)
		} else {
			parts := strings.Split(line, ",")
			tests = append(tests, make([]int, len(parts)))
			for i, val := range parts {
				num, _ := strconv.Atoi(val)
				tests[index][i] = num
			}
			index++
		}
	}
	var invalidTests [][]int
	for i := 0; i < len(tests); i++ {
	A:
		for j := 0; j < len(tests[i]); j++ {
			currentNumber := tests[i][j]
			ruleForNumber := rules[currentNumber]

			for k := j + 1; k < len(tests[i]); k++ {
				nextNumber := tests[i][k]
				for _, ruleValue := range ruleForNumber {
					if nextNumber == ruleValue {
						invalidTests = append(invalidTests, tests[i])
						break A
					}
				}
			}

		}
	}

	sortCorrectlyTest := func(test []int, rules map[int][]int) []int {

		updatedTest := append([]int{}, test...)

		for i := 0; i < len(updatedTest); i++ {
			currentNumber := updatedTest[i]
			ruleForNumber := rules[currentNumber]

			for k := i + 1; k < len(updatedTest); k++ {
				nextNumber := updatedTest[k]
				for _, ruleValue := range ruleForNumber {
					if nextNumber == ruleValue {
						newTest := append([]int{}, updatedTest[:i]...)   // Alles bis i-1
						newTest = append(newTest, updatedTest[i+1:k]...) // i+1 bis k-1
						newTest = append(newTest, updatedTest[k])        // Element k
						newTest = append(newTest, currentNumber)         // i selbst
						newTest = append(newTest, updatedTest[k+1:]...)  // k+1 bis Ende

						updatedTest = newTest
						i = -1
						break
					}
				}
				if i == -1 {
					break
				}
			}
		}

		return updatedTest
	}

	for _, invalidTest := range invalidTests {
		correctedTest := sortCorrectlyTest(invalidTest, rules)
		sum += correctedTest[len(correctedTest)/2]
	}

	fmt.Println(sum)
}
