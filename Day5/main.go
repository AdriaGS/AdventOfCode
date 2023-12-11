package main

import (
	"flag"
	"fmt"
	"os"
)


func main() {

	filename := flag.String("file", "", "filename to read")

	flag.Parse()

	if *filename == "" {
		flag.PrintDefaults()
		return
	}

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	inputParser := NewInputParser(file)
	almanac := inputParser.ParseFile()

	// Find lowest location value
	lowestLocationValue := almanac.findLowestLocationNumber()
	fmt.Println("Lowest Location Value: ", lowestLocationValue)
}
