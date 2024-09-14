package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(filename string) ([]float64, error) {
	var storage []float64
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		storage = append(storage, temp)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return storage, nil
}
