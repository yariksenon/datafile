package datafile

import (
	"bufio"
	"fmt"
	"os"
)

func GetFloats(filename string) ([]string, error) {
	var numbers []string
	file, err := os.Open(filename)
	if err != nil {
		return numbers, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number := scanner.Text()
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

func SayHello() {
	fmt.Println("Hello")
}
