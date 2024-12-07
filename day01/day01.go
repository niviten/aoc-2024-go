package day01

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func RunSampleInputPartOne() {
	runPartOne("day01/sample_input")
}

func RunPartOne() {
	runPartOne("day01/input")
}

func runPartOne(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error while opening file: %s\n", filePath)
		fmt.Println("error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		a, _ := strconv.ParseInt(parts[0], 10, 32)
		b, _ := strconv.ParseInt(parts[len(parts)-1], 10, 32)
		leftList = append(leftList, int(a))
		rightList = append(rightList, int(b))
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	sum := 0

	for idx, a := range leftList {
		sum = sum + diff(a, rightList[idx])
	}

	fmt.Printf("Answer: %d\n", sum)
}

func RunSampleInputPartTwo() {
	runPartTwo("day01/sample_input")
}

func RunPartTwo() {
	runPartTwo("day01/input")
}

func runPartTwo(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error while opening file: %s\n", filePath)
		fmt.Println("error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	leftList := []int{}
	rightList := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		a, _ := strconv.ParseInt(parts[0], 10, 32)
		b, _ := strconv.ParseInt(parts[len(parts)-1], 10, 32)
		leftList = append(leftList, int(a))
		rightList = append(rightList, int(b))
	}

	m := make(map[int]int)

	for _, a := range rightList {
		m[a] = m[a] + 1
	}

	sum := 0

	for _, a := range leftList {
		sum = sum + (a * m[a])
	}

	fmt.Printf("Answer: %d\n", sum)
}

func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
