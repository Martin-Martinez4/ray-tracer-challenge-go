package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
	Current State:
		Mostly finished with the main Ray Tracer.  I would like to move on to a new project, but I can come back later to fix, refactor, and add features
	To do:
		- Store the inverse Transform to increase performance
		- Add a function to check for a child shape to Shape
		- Finish chapter 16 tests
		- Intersections need to be refactored
		- Put the obj parser into a sub package

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

	// printToFile(ch11A(), "chapter11A.ppm")
	// printToFile(ch15B(), "chapter15B.ppm")
	// ParseObjFile("test.obj")

	sr := strings.Split("1//3", "/")
	sr2 := strings.Split("1/2/3", "/")

	fmt.Println(sr[1] == "")
	fmt.Println(sr2)

}
