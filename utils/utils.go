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

func DeleteAtIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
