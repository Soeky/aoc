package days

import (
	"bufio"
	"fmt"
	"os"
)

func DayX() {
	fmt.Println("Solution for Day X, Part 1:")
	solveDay2Part1()
	fmt.Println("Solution for Day X, Part 2:")
	solveDay2Part2()
}

func solveDayXPart1() {
	file, err := os.Open("days/Xinput.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func solveDayXPart2() {
	file, err := os.Open("days/Xinput.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
