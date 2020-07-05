package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileNames := os.Args[1:]
	fileContents := make(map[string]string)
	for _, fileName := range fileNames {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			continue
		}
		fileContents[fileName] = string(data)
	}

	for filename, file := range fileContents {
		fmt.Printf("%s contains:\n%s\n", filename, file)
	}

	for fileName, file := range fileContents {
		byteString := []byte(file)
		for i, character := range byteString {
			byteString[i] = character + 1
		}
		fileContents[fileName] = string(byteString)
	}

	fmt.Printf("\nAnd Scramble!\n\n");
	for filename, file := range fileContents {
		fmt.Printf("%s contains:\n%s\n", filename, file)
	}
}
