package main

import (
	"fmt"
	"log"
	"os"
)

/*
	To do:
		- test matrix transformations?
		-
*/

func printToFile(str string, filepath string) {

	f, err := os.Create(filepath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(str)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}

func main() {

	// fmt.Println("Should be creating Canvas")
	// for chapter 5
	// printch6(ch6())

	printToFile(ch8(), "chapter8.ppm")

}
