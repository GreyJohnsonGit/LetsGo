package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	var filenames = os.Args[1:]
	var fileContents = make(map[string]string)
	for _, filename := range filenames {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			continue
		}
		fileContents[filename] = string(data)
	}

	for filename, file := range fileContents {
		fmt.Printf("%s contains:\n%s\n", filename, file)
	}
}
