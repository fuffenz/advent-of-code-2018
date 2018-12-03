package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	data := readInput()

	result := getResult(data)
	log.Printf("part 1: %d\n", result)

	firstDuplicate := getFirstDuplicate(data)

	log.Printf("part 2: %d\n", firstDuplicate)
}

func readInput() (data []int64) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		data = append(data, n)
	}

	return data
}

func getResult(data []int64) (result int64) {
	result = 0
	for n := range data {
		result += data[n]
	}
	return result
}

func getFirstDuplicate(data []int64) int64 {
	m := make(map[int64]bool)
	var result int64

	for true {
		for n := range data {
			result += data[n]
			if m[result] {
				return result
			}
			m[result] = true
		}
	}

	return 0
}
