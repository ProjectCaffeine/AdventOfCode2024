package main

import (
	"log"
	"os"
)

func handlePossibleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getFileData(fileName string) []byte {
	data, err := os.ReadFile("./" + fileName)

	handlePossibleError(err)

	return data
}
