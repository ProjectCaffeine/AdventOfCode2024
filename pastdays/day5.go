package day5

import (
	"bytes"
	"fmt"
	"strings"
)

func parseTextForDay4(data []byte) {
	splitLines := bytes.Split(data, []byte("\n"))
	cnt := 0

	for rowIdx, row := range splitLines {
		for colIdx, char := range row {
			if char != 'A' {
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
		//"up-down",
		//"left-right",
		"upleft-downright",
		"upright-downleft",
	}

	for _, dir := range directions {
		if searchDirection(data, currRow, currCol, dir) {
			matches++
		}
	}

	if matches >= 2 {
		return 1
	}

	return 0
}

func searchDirection(data [][]byte, currRow int, currCol int, direction string) bool {
	chars := []rune {
		'M',
		'S',
	}

	for _, char := range chars {
		hasChar := false
		for _, subDirction := range strings.Split(direction, "-") {
			if directionHasChar(data, currRow, currCol, subDirction, char) {
				hasChar = true
				break
			}
		}

		if !hasChar {
			return false
		}
	}
	
	return true
}

func directionHasChar(data [][]byte, rowToSearch int, colToSearch int, direction string, char rune) bool {

		if strings.Contains(direction, "up") {
			rowToSearch--
		}

		if strings.Contains(direction, "down") {
			rowToSearch++ 
		}

		if strings.Contains(direction, "left") {
			colToSearch-- 
		}

		if strings.Contains(direction, "right") {
			colToSearch++
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

	return true
}
