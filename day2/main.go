package main

import (
	"flag"
	"fmt"
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

	levels, shouldReturn := extractLevels(lines)
	if shouldReturn {
		return
	}

	if part == 1 {
		part1(levels)
	} else if part == 2 {
		part2(levels)
	}
	//print execution time
	elapsed := time.Since(now)
	fmt.Println("Execution time:", elapsed)
}

func part1(levels map[int][]int) {
	fmt.Println("Part 1")
	//So, a report only counts as safe if both of the following are true:
	//The levels are either all increasing or all decreasing.
	//Any two adjacent levels differ by at least one and at most three.
	safeCount := 0

	for i, level := range levels {
		isLevelSafe := checkLevelSafe(i, level)

		if isLevelSafe {
			safeCount++
		}
	}
	fmt.Println("Safe levels:", safeCount)
}

func checkLevelSafe(i int, level []int) bool {
	lastPositive := false
	isLevelSafe := true
	for j := 0; j < len(level)-1; j++ {
		distance, positive := utils.Distance(level[j+1], level[j])
		if j == 0 {
			lastPositive = positive
		}

		if lastPositive != positive {
			isLevelSafe = false
		}

		lastPositive = positive

		//fmt.Println("Distance between", level[j+1], "and", level[j], "is", distance)
		if distance < 1 || distance > 3 {
			isLevelSafe = false
		}

	}
	if isLevelSafe {
		fmt.Println("Level", i, "is safe")
	} else {
		fmt.Println("Level", i, "is not safe")
	}
	return isLevelSafe
}

func part2(levels map[int][]int) {
	fmt.Println("Part 2")
	//So, a report only counts as safe if both of the following are true:
	//The levels are either all increasing or all decreasing.
	//Any two adjacent levels differ by at least one and at most three.
	safeCount := 0

	for i, level := range levels {
		fmt.Println("Level:", i, "Level:", level)
		isLevelSafe := checkLevelSafe(i, level)

		if isLevelSafe {
			safeCount++
		} else {
			//remove 1 item and check if it is safe
			for j := 0; j < len(level); j++ {
				newLevel := utils.DeleteAtIndex(level, j)
				fmt.Println("New level:", newLevel)
				isLevelSafe := checkLevelSafe(i, newLevel)
				if isLevelSafe {
					safeCount++
					break
				}
			}
		}
	}
	fmt.Println("Safe levels:", safeCount)
}

func extractLevels(lines []string) (map[int][]int, bool) {
	var levels = make(map[int][]int)
	for i, line := range lines {
		levelStrings := strings.Split(line, " ")
		for _, levelString := range levelStrings {
			level, err := strconv.ParseInt(levelString, 10, 64)
			if err != nil {
				fmt.Println("Error parsing level:", err)
				return nil, true
			}
			levels[i] = append(levels[i], int(level))
		}
	}
	return levels, false
}
