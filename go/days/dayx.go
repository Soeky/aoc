package days

import (
	"bufio"
	"fmt"
	"os"
)

func DayX() {
	fmt.Println("Solution for Day X, Part 1:", solveDayXPart1())
	fmt.Println("Solution for Day X, Part 2:", solveDayXPart2())
}

func solveDayXPart1() int {
	file, err := os.Open("days/Xinput.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return sum
}

func solveDayXPart2() int {
	file, err := os.Open("days/Xinput.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return sum
}
