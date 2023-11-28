// Package datafile reads and returns data from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats returns float64 values from files
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, x)
	}
	// hello..
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}
