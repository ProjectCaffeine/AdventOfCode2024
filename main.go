package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data := getFileData("day5input.txt")

	fmt.Printf("%q", data)
}

func parseData(data []byte) {
	splitData := bytes.Split(data, []byte("\n\n"))
	firstPageRules, secondPageRules := getRules(splitData[0])

	processUpdates(splitData[1], firstPageRules, secondPageRules)
}

func processUpdates(updates []byte, firstPageRules map[int][]int, secondPageRules map[int][]int) int {
	sum := 0

	for _, update := range bytes.Split(updates, []byte("\n")) {
		if !isUpdateValid(update, firstPageRules, secondPageRules) {
			continue
		}

		//Add sum here
	}

	return sum
}

func isUpdateValid(update []byte, firstPageRules map[int][]int, secondPageRules map[int][]int) bool {
	pageNumbers := bytes.Split(update, []byte(","))
	updatedPages := make(map[int]bool, len(pageNumbers))

	for _, pageNum := range pageNumbers {
		num, err := strconv.Atoi(string(pageNum))

		handlePossibleError(err)

		laterPages, hasFirstPageRules := firstPageRules[num]

		if hasFirstPageRules {
			for _, page := range laterPages {
				_, hasPage := updatedPages[page]

				if hasPage {
					return false
				}
			}
		}

		earlierPages, hasSecondPageRules := secondPageRules[num]

		if hasSecondPageRules {
			for _, page := range earlierPages {
				_, hasPage := updatedPages[page]

				if !hasPage {
					return false 
				}
			}
		}

		updatedPages[num] = true
	}

	return true
}

func getRules(data []byte) (map[int][]int, map[int][]int){
	lines := bytes.Split(data, []byte("\n"))
	firstPageRules := make(map[int][]int, len(lines))
	secondPageRules := make(map[int][]int, len(lines))

	for _, line := range lines {
		strLine := string(line)
		if strLine == "" {
			break
		}

		splitLine := strings.Split(strLine, "|")
		firstPage, err := strconv.Atoi(splitLine[0])

		handlePossibleError(err)
		secondPage, err := strconv.Atoi(splitLine[1])

		handlePossibleError(err)

		_, ok := firstPageRules[firstPage]

		if ok && !slices.Contains(firstPageRules[firstPage], secondPage) {
			firstPageRules[firstPage] = append(firstPageRules[firstPage], secondPage)
		} else if !ok {
			firstPageRules[firstPage] = []int {secondPage}
		}

		_, secondPageOk := secondPageRules[secondPage]

		if secondPageOk && !slices.Contains(firstPageRules[firstPage], secondPage) {
			secondPageRules[secondPage] = append(secondPageRules[secondPage], firstPage)
		} else if !secondPageOk {
			secondPageRules[secondPage] = []int {firstPage}
		}
	}

	return firstPageRules, secondPageRules
}
