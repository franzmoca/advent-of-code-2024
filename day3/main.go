package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
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

	if part == 1 {
		part1(lines)
	} else if part == 2 {
		part2(lines)
	}
	//print execution time
	elapsed := time.Since(now)
	fmt.Println("Execution time:", elapsed)
}

func part1(lines []string) {
	fmt.Println("Part 1")
	mults, shouldReturn := extractMultiplications1(lines)
	if shouldReturn {
		return
	}

	var total int = 0
	for _, mult := range mults {
		total += mult.a * mult.b
	}
	fmt.Println("Total:", total)

}

func part2(lines []string) {
	fmt.Println("Part 2")
	mults, shouldReturn := extractMultiplications2(lines)
	if shouldReturn {
		return
	}

	var total int = 0
	for _, mult := range mults {
		total += mult.a * mult.b
	}
	fmt.Println("Total:", total)
}

type Mult struct {
	a int
	b int
}

func extractMultiplications1(lines []string) ([]Mult, bool) {
	//regex to match the following pattern: mul(X,Y), where X and Y are each 1-3 digit numbers.
	//The numbers are separated by a comma and the whole string is enclosed in parentheses.
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	var mults []Mult //slice to store the extracted numbers
	for _, line := range lines {
		matches := mulRegex.FindAllStringSubmatch(line, -1)
		if len(matches) > 0 {
			for _, match := range matches {
				a, err := strconv.Atoi(match[1])
				if err != nil {
					fmt.Println("Error parsing a:", err)
					return nil, true
				}
				b, err := strconv.Atoi(match[2])
				if err != nil {
					fmt.Println("Error parsing b:", err)
					return nil, true
				}
				mult := Mult{a, b}
				mults = append(mults, mult)
			}
		}

	}
	return mults, false
}

func extractMultiplications2(lines []string) ([]Mult, bool) {

	//merge lines into one string
	linesStr := ""
	for _, line := range lines {
		linesStr += line
	}

	r, _ := regexp.Compile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)

	var mults []Mult //slice to store the extracted numbers

	enabled := true
	for _, match := range r.FindAllString(linesStr, -1) {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			nums := regexp.MustCompile(`\d+`)
			numsStr := nums.FindAllString(match, -1)
			a, err := strconv.Atoi(numsStr[0])
			if err != nil {
				fmt.Println("Error parsing a:", err)
				return nil, true
			}
			b, err := strconv.Atoi(numsStr[1])
			if err != nil {
				fmt.Println("Error parsing b:", err)
				return nil, true
			}
			mult := Mult{a, b}
			mults = append(mults, mult)
		}
	}
	return mults, false
}
