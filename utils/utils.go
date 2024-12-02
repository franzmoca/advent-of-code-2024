package utils

import (
	"os"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	dataStr := strings.TrimSpace(string(data))
	lines := strings.Split(dataStr, "\n")

	return lines, nil
}

func SumInt(ints []int64) int64 {
	var sum int64
	for _, val := range ints {
		sum += val
	}
	return sum
}

func Distance(a, b int) (int, bool) {
	diff := a - b
	positive := diff >= 0
	if diff < 0 {
		diff = -diff
	}
	return diff, positive
}

func DeleteAtIndex(slice []int, index int) []int {
	var newSlice []int
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)
	return newSlice
}
