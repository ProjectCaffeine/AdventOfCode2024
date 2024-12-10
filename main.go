package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	data := getFileData("day4input.txt")

	parseTextForDay4(data)
}

func parseTextForDay4(data []byte) {
	splitLines := bytes.Split(data, []byte("\n"))
	cnt := 0

	for rowIdx, row := range splitLines {
		for colIdx, char := range row {
			if char != 'X' {
				continue
			}

			cnt += searchForWord(splitLines, rowIdx, colIdx)
		}
	}

	fmt.Printf("Total found: %d\n", cnt)
}

func searchForWord(data [][]byte, currRow int, currCol int) int {
	matches := 0
	directions := []string {
		"up",
		"down",
		"left",
		"right",
		"upleft",
		"upright",
		"downleft",
		"downright",
	}

	for _, dir := range directions {
		//if !canSearch(data, currRow, currCol, dir) {
		//	continue
		//}

		if searchDirection(data, currRow, currCol, dir) {
			matches++
		}
	}

	return matches
}

func searchDirection(data [][]byte, currRow int, currCol int, direction string) bool {
	chars := []rune {
		'X',
		'M',
		'A',
		'S',
	}
	rowToSearch := currRow
	colToSearch := currCol
	
	for idx, char := range chars {
		if strings.Contains(direction, "up") {
			rowToSearch -= idx
		}

		if strings.Contains(direction, "down") {
			rowToSearch += idx
		}

		if strings.Contains(direction, "left") {
			colToSearch -= idx
		}

		if strings.Contains(direction, "right") {
			colToSearch += idx
		}

		if rowToSearch >= len(data) || rowToSearch < 0 {
			return false
		}

		if colToSearch < 0 || colToSearch >= len(data[rowToSearch]) {
			return false
		}
		
		
		if data[rowToSearch][colToSearch] != byte(char) {
			return false
		}

		rowToSearch = currRow
		colToSearch = currCol
	}

	return true
}

func canSearch(data [][]byte, currRow int, currCol int, direction string) bool {
	if strings.Contains(direction, "up") && currRow < 3 {
		return false
	}

	if strings.Contains(direction, "down") && currRow > len(data) - 4 {
		return false
	}

	if strings.Contains(direction, "left") && currCol < 3 {
		return false
	}

	if strings.Contains(direction, "right") && currCol > len(data[currRow]) - 4 {
		return false
	}

	return true
}
