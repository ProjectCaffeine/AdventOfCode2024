package main

import (
	"bytes"
	"fmt"
	"maps"
	"math"
	"sort"
	"strconv"
)

func day1Main() {
	data := getFileData("day1input.txt")
	firstList, secondList := breakDataIntoLists(data)

	simMap := getSimilarityMap(firstList, secondList)
	totalSimilarity := sumSimilarityMap(simMap)

	fmt.Printf("Total similarity: %d\n", totalSimilarity)
}

func calculateTotalDifference(firstList []int, secondList []int) int {
	sort.Ints(firstList)
	sort.Ints(secondList)

	totalDifference := getTotalDifference(firstList, secondList)

	fmt.Printf("Total Difference: %d\n", totalDifference)

	return totalDifference
}

func getTotalDifference(firstList []int, secondList []int) int {
	totalDifference := 0

	for idx, v := range firstList {
		difference := int(math.Abs(float64(v) - float64(secondList[idx])))

		totalDifference += difference
	}

	return totalDifference
}

func printLists(firstList []int, secondList []int) {
	fmt.Printf("First List:\n")

	for _, v := range firstList {
		fmt.Printf("%d\n", v)
	}

	fmt.Printf("\n\nSecond List:\n")

	for _, v := range secondList {
		fmt.Printf("%d\n", v)
	}
}

func breakDataIntoLists(data []byte) (firstList []int, secondList []int) {
	dataByLines := bytes.Split(data, []byte("\n"))
	firstList = make([]int, len(dataByLines) - 1)
	secondList = make([]int, len(dataByLines) - 1)

	for idx, v := range dataByLines {
		if string(v) == "" {
			break
		}

		lineValues := bytes.Split(v, []byte("   "))
		firstNum, err := strconv.Atoi(string(lineValues[0]))

		handlePossibleError(err)

		secondNum, err := strconv.Atoi(string(lineValues[1]))

		handlePossibleError(err)

		firstList[idx] = firstNum
		secondList[idx] = secondNum
	}

	return firstList, secondList
}

func getSimilarityMap(firstList []int, secondList []int) map[int]int {
	simMap := make(map[int]int)

	for _, v := range firstList {
		simMap[v] = 0
	}

	for _, v := range secondList {
		fmt.Printf("Num: %d\n", v)
		occurrences, hasKey := simMap[v]

		if hasKey {
			fmt.Printf("Has Key: %d\n", v)
			simMap[v] = occurrences + 1
		}
	}

	return simMap
}

func sumSimilarityMap(simMap map[int]int) int{
	totalSimilarity := 0

	for k := range maps.Keys(simMap) {
		totalSimilarity += k * simMap[k]
	}

	return totalSimilarity
}
