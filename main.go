package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
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

}

func breakDataIntoLists(data []byte) (firstList []int, secondList []int) {
	dataByLines := bytes.Split(data, []byte("\n"))
	firstList = make([]int, len(dataByLines))
	secondList = make([]int, len(dataByLines))

	for _, v := range dataByLines {
		lineValues := bytes.Split(v, []byte("   "))
		firstNum, err := strconv.Atoi(string(lineValues[0]))

		handlePossibleError(err)

		secondNum, err := strconv.Atoi(string(lineValues[1]))

		handlePossibleError(err)

		firstList = append(firstList, firstNum)
		secondList = append(secondList, secondNum)
	}

	return firstList, secondList
}

func getFileData() []byte {
	data, err := os.ReadFile("./day1input.txt")

	handlePossibleError(err)

	return data
}
