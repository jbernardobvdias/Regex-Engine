package main

import (
	"fmt"
	"regexengine/engine"
)

func main() {
	fmt.Println("What is the path to the file you want to process?")

	var path string
	fmt.Scanln(&path)

	fmt.Println("What is the pattern you are going to use?")

	var pattern string
	fmt.Scanln(&pattern)

	data := engine.Data{Input: path, Pattern: pattern}
	fmt.Println(data)

}
