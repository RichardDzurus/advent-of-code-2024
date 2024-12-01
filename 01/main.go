package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func convertLineToNumbers(line string) (int, int, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return 0, 0, errors.New("expected exactly two numbers in the line")
	}

	num1, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		return 0, 0, err1
	}

	num2, err2 := strconv.Atoi(parts[1])
	if err2 != nil {
		return 0, 0, err2
	}

	return num1, num2, nil
}

func getInput(fileName string) ([]int, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1 []int
	var list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		num1, num2, err := convertLineToNumbers(line)
		if err != nil {
			return nil, nil, err
		}
		list1 = append(list1, num1)
		list2 = append(list2, num2)

	}

	return list1, list2, nil
}

func getTotalDistanceOfUnsortedLists(list1 []int, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0

	for i := range list1 {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff
		}
		totalDistance += diff
	}

	return totalDistance
}

func getSimilarityScore(list1 []int, list2 []int) int64 {
	similarityMap := make(map[int]int)

	for _, number := range list2 {
		similarityMap[number] += 1
	}

	var similarityScore int64 = 0

	for _, number := range list1 {
		similarityScore += int64(number * similarityMap[number])
	}

	return similarityScore
}

func main() {
	list1, list2, err := getInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	totalDistance := getTotalDistanceOfUnsortedLists(list1, list2)
	similarityScore := getSimilarityScore(list1, list2)

	fmt.Println(totalDistance)
	fmt.Println(similarityScore)
}
