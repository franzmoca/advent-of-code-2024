package main

import (
	"flag"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/franzmoca/aoc24/utils"
)

func main() {
	now := time.Now()
	var part int
	var input string
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.StringVar(&input, "input", "input.txt", "input file")
	flag.Parse()
	fmt.Println("Running part", part)
	fmt.Println("Using input file:", input)

	lines, err := utils.ReadLines(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	column1, column2, shouldReturn := extractCols(lines)
	if shouldReturn {
		return
	}

	fmt.Println("Column 1:", column1)
	fmt.Println("Column 2:", column2)

	if part == 1 {
		part1(column1, column2)
	} else if part == 2 {
		part2(column2, column1)
	}
	//print execution time
	elapsed := time.Since(now)
	fmt.Println("Execution time:", elapsed)
}

func part1(column1 []int64, column2 []int64) {
	slices.Sort(column1)
	slices.Sort(column2)

	fmt.Println("Sort Column 1:", column1)
	fmt.Println("Sort Column 2:", column2)

	var distances []int64
	for i := 0; i < len(column1); i++ {
		diff := math.Abs(float64(column1[i] - column2[i]))
		distances = append(distances, int64(diff))
	}

	fmt.Println("Distances:", distances)

	sum := utils.SumInt(distances)

	fmt.Println("Sum of distances:", sum)
}

func part2(column2 []int64, column1 []int64) {
	counts := make(map[int64]int)
	for _, value := range column2 {
		counts[value]++
	}
	var similiarity []int64

	for _, value := range column1 {
		similiarity = append(similiarity, value*int64(counts[value]))
	}
	fmt.Println("Similiarity:", similiarity)

	sum := utils.SumInt(similiarity)

	fmt.Println("Sum of similiarity:", sum)
}

func extractCols(lines []string) ([]int64, []int64, bool) {
	var column1, column2 []int64

	for _, line := range lines {
		columns := strings.Split(line, "   ")
		if len(columns) >= 2 {
			col1Int, err := strconv.ParseInt(columns[0], 10, 64)
			if err != nil {
				fmt.Println("Error parsing column 1:", err)
				return nil, nil, true
			}
			col2Int, err := strconv.ParseInt(columns[1], 10, 64)
			if err != nil {
				fmt.Println("Error parsing column 2:", err)
				return nil, nil, true
			}
			column1 = append(column1, col1Int)
			column2 = append(column2, col2Int)
		}
	}
	return column1, column2, false
}
