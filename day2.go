package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func processFileForDay2(data []byte) {
	lines := bytes.Split(data, []byte("\n"))
	reports := make([]string, len(lines))
	totalValid := 0

	for idx, val := range lines {
		strVal := string(val)
		if strVal == "" {
			continue
		}

		reports[idx] = strVal

		if isReportValid(reports[idx]) {
			totalValid++
		}
	}

	fmt.Printf("Total Valid: %d\n", totalValid)
}

func isReportValid(report string) bool {
	indivNumStrs := strings.Split(report, " ")
	indivNums := convertStrArrToIntArr(indivNumStrs)
	num := indivNums[0]
	skipped := false

	if len(indivNums) == 1 {
		return true
	}

	isAsc := false

	if indivNums[0] == indivNums[1] {
		isAsc = indivNums[1] > indivNums[0]
	} else {
		isAsc = indivNums[2] > indivNums[0]
	}

	for _, val := range indivNums[1:] {
		dif := math.Abs(float64(val) - float64(num))

		if dif < 1 || dif > 3 {
			if !skipped {
				skipped = true
			} else {
				return false
			}
		}

		if (isAsc && val < num) ||(!isAsc && val > num) {
			if !skipped {
				skipped = true
			} else {
				return false
			}
		}

		num = val
	}

	return true
}

func convertStrArrToIntArr(arr []string) []int {
	results := make([]int, len(arr))

	for idx, val := range arr {
		num, err := strconv.Atoi(val)

		handlePossibleError(err)

		results[idx] = num
	}

	return results
}
