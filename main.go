package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data := getFileData("day3input.txt")

	parseText(data)
}

type command struct {
	Do bool
	Index []int
}

func getDosAndDonts(data []byte) []command {
	doRegex, err := regexp.Compile(`do\(\)`)

	handlePossibleError(err)
	dontRegex, err := regexp.Compile(`don't\(\)`)

	handlePossibleError(err)

	results := make([]command, 0, 256)

	doRegexIdxs := doRegex.FindAllIndex(data, -1)

	results = append(results, command {
		Do: true,
		Index: []int {0, 0},
	})

	if doRegexIdxs != nil {
		for _, val := range doRegexIdxs {
			results = append(results, command {
				Do: true,
				Index: val,
			})
		}
	}

	dontRegexIdxs := dontRegex.FindAllIndex(data, -1)

	if dontRegexIdxs != nil {
		for _, val := range dontRegexIdxs {
			results = append(results, command {
				Do: false,
				Index: val,
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[j].Index[0] > results[i].Index[0]
	})

	return results
}

func printDosAndDonts(dosAndDonts []command) {
	for _, val := range dosAndDonts {
		fmt.Printf("Do: %t\nIndex:[%d, %d]\n\n", val.Do, val.Index[0], val.Index[1])
	}
}

func parseText(data []byte) {
	dosAndDonts := getDosAndDonts(data)
	printDosAndDonts(dosAndDonts)

	regex, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	handlePossibleError(err)

	results := regex.FindAllIndex(data, -1)

	processEnabledCommands(dosAndDonts, results, data)
}

func processEnabledCommands(dosAndDonts []command, mulCommands [][]int, data []byte) {
	sum := 0
	enabled := true
	enabledIdx := 0

	for _, val := range dosAndDonts[1:] {
		if enabled && !val.Do {
			for _, mulCommand := range mulCommands {
				if mulCommand[0] < enabledIdx {
					continue
				} else if mulCommand[0] > val.Index[0] {
					break
				}

				sum += parseMul(string(data[mulCommand[0]:mulCommand[1]]))
			}
		}

		if !enabled && val.Do {
			enabledIdx = val.Index[0]
		}

		enabled = val.Do
	}

	fmt.Printf("The total sum is: %d\n", sum)
}

func parseMul(funcCall string) int {
	cleanedStr := strings.Replace(funcCall, "mul(", "", -1)
	cleanedStr = strings.Replace(cleanedStr, ")", "", -1)
	splitStr := strings.Split(cleanedStr, ",")
	firstNum, err := strconv.Atoi(splitStr[0])

	handlePossibleError(err)
	secondNum, err := strconv.Atoi(splitStr[1])

	handlePossibleError(err)

	return firstNum * secondNum
}
