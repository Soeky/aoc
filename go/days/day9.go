package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day9() {
	fmt.Println("Solution for Day 9, Part 1:", solveDay9Part1())
	fmt.Println("Solution for Day 9, Part 2:", solveDay9Part2())
	solveDay9Part2()
}

func solveDay9Part1() int {
	file, err := os.Open("days/9input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var res []int

	for scanner.Scan() {
		line := scanner.Text()
		num := 0
		for i, char := range line {
			if i%2 == 0 {
				numBlocks := char - '0'
				for i := numBlocks; i > 0; {
					res = append(res, num)
					i--
				}
				num++
			} else {
				numBlocks := char - '0'
				for i := numBlocks; i > 0; {
					res = append(res, -1)
					i--
				}
			}
		}
	}
	l := 0
	r := len(res) - 1
	for l < r {
		for l < r && res[l] != -1 {
			l++
		}
		for l < r && res[r] == -1 {
			r--
		}
		if l < r {
			tmp := res[l]
			res[l] = res[r]
			res[r] = tmp
			l++
			r--
		}
	}
	sum := 0
	for i, val := range res {
		if val == -1 {
			break
		}
		sum += i * val
	}
	return sum
}

func solveDay9Part2() int {
	file, err := os.Open("days/9input.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var res []int

	num := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i, char := range line {
			if i%2 == 0 {
				numBlocks := char - '0'
				for i := numBlocks; i > 0; {
					res = append(res, num)
					i--
				}
				num++
			} else {
				numBlocks := char - '0'
				for i := numBlocks; i > 0; {
					res = append(res, -1)
					i--
				}
			}
		}
	}
	// max num = num-1
	num--

	for id := num; id >= 0; id-- {
		start, end := findFileRange(res, id)
		if start == -1 {
			continue
		}
		size := end - start + 1
		freeStart := findFreeStart(res, size, 0, end-1)
		if freeStart != -1 {
			for i := 0; i < size; i++ {
				res[freeStart+i] = res[start+i]
			}

			for i := start; i <= end; i++ {
				res[i] = -1
			}
		}
	}

	sum := 0
	for i, val := range res {
		if val == -1 {
			continue
		}
		sum += i * val
	}
	return sum
}

func findFileRange(res []int, id int) (int, int) {
	start := -1
	end := -1
	for i, v := range res {
		if v == id {
			if start == -1 {
				start = i
			}
			end = i
		}
	}
	return start, end
}

func findFreeStart(res []int, size int, left int, right int) int {
	if right < left || size <= 0 {
		return -1
	}
	length := 0
	startIndex := -1
	for i := left; i <= right; i++ {
		if res[i] == -1 {
			if startIndex == -1 {
				startIndex = i
			}
			length++
			if length == size {
				return startIndex
			}
		} else {
			length = 0
			startIndex = -1
		}
	}
	return -1
}
