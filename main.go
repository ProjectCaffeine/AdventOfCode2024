package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func handlePossibleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data := getFileData()
	firstList, secondList := breakDataIntoLists(data)

	sort.Ints(firstList)
	sort.Ints(secondList)

	totalDifference := getTotalDifference(firstList, secondList)

	fmt.Printf("Total Difference: %d\n", totalDifference)
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

func getFileData() []byte {
	data, err := os.ReadFile("./day1input.txt")

	handlePossibleError(err)

	return data
}
