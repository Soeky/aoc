package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	fmt.Println("Solution for Day 2, Part 1:")
	solveDay2Part1()
	fmt.Println("Solution for Day 2, Part 2:")
	solveDay2Part2()
}

func solveDay2Part1() {
	file, err := os.Open("days/2input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	safeLines := 0

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		lastnum := -1
		increasing := false
		isSafe := true
		for i, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("konnte Zahl nicht lesen")
				return
			}
			if lastnum == -1 {
				lastnum = number
				continue
			} else {
				diff := lastnum - number
				lastnum = number
				if i == 1 {
					if diff < 0 {
						increasing = true
					}
				}
				if increasing && (diff >= 0 || diff < -3) {
					isSafe = false
					break
				} else if !increasing && (diff <= 0 || diff > 3) {
					isSafe = false
					break
				}
			}
		}
		if isSafe {
			safeLines++
		}
	}
	fmt.Println(safeLines)
}

func solveDay2Part2() {
	file, err := os.Open("days/2input.txt")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeLines := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		nums := []int{}
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("konnte Zahl nicht lesen")
				return
			}
			nums = append(nums, number)
		}

		if isSafeSequence(nums) {
			safeLines++
			continue
		}

		for i := 0; i < len(nums); i++ {
			temp := append([]int{}, nums[:i]...)
			temp = append(temp, nums[i+1:]...)
			if isSafeSequence(temp) {
				safeLines++
				break
			}
		}
	}

	fmt.Println(safeLines)
}

func isSafeSequence(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	increasing := nums[1] > nums[0]
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		if increasing && diff < 0 {
			return false
		}
		if !increasing && diff > 0 {
			return false
		}
	}

	return true
}
