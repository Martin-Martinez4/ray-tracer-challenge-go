package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
	Regex for Vertex
	v ((-?)[0-9]+(\.[0-9]+\s|\s)){2}((-?)[0-9]+(\.[0-9]+)?)\n
*/

type ParserOBJ struct {
	Vertices  []Tuple
	Triangles []Triangle
}

func parserVertices(contents string, parser *ParserOBJ) {
	re := regexp.MustCompile(`v (?:-?[0-9]+(?:\.[0-9]+)?\s){2}-?[0-9]+(?:\.[0-9]+)?`)
	re2 := regexp.MustCompile(`(?:-?[0-9]+(?:\.[0-9]+)?)`)

	for _, thing := range re.FindAllStringSubmatch(contents, -1) {
		numbs := re2.FindAllStringSubmatch(thing[0], -1)

		fmt.Printf("%v\n", numbs)

		if len(numbs) != 3 {
			panic("v length is incorrect")

		}

		numbs00, err := strconv.ParseFloat(strings.TrimSpace(numbs[0][0]), 64)
		if err != nil {
			// maybe reutrn error instead
			panic("failed to convert to float64")
		}
		numbs01, err := strconv.ParseFloat(strings.TrimSpace(numbs[1][0]), 64)
		if err != nil {
			panic("failed to convert to float64")
		}
		numbs02, err := strconv.ParseFloat(strings.TrimSpace(numbs[2][0]), 64)
		if err != nil {
			panic("failed to convert to float64")
		}

		vertex := Point(numbs00, numbs01, numbs02)

		parser.Vertices = append(parser.Vertices, vertex)
	}
}

func parseTris(contents string, parser ParserOBJ) {
	/*
		f 1 2 3
		f = face
		numbers will be digits becuase they correspond to the vertices in the Vertex slice

		may need to make \n optional
	*/
	re := regexp.MustCompile(`f (?:[0-9]+\s){2}(?:[0-9]+\n)`)
	for _, thing := range re.FindAllStringSubmatch(contents, -1) {
		numbs := re2.FindAllStringSubmatch(thing[0], -1)

		fmt.Printf("%v\n", numbs)

		if len(numbs) != 3 {
			panic("v length is incorrect")

		}

		numbs00, err := strconv.ParseFloat(strings.TrimSpace(numbs[0][0]), 64)
		if err != nil {
			// maybe reutrn error instead
			panic("failed to convert to float64")
		}
		numbs01, err := strconv.ParseFloat(strings.TrimSpace(numbs[1][0]), 64)
		if err != nil {
			panic("failed to convert to float64")
		}
		numbs02, err := strconv.ParseFloat(strings.TrimSpace(numbs[2][0]), 64)
		if err != nil {
			panic("failed to convert to float64")
		}

		vertex := Point(numbs00, numbs01, numbs02)

		parser.Vertices = append(parser.Vertices, vertex)
	}

}

func ParseObj(contents string) *ParserOBJ {

	parser := ParserOBJ{Vertices: []Tuple{}, Triangles: []Triangle{}}

	parserVertices(contents, &parser)

	return &parser
}

func (pO *ParserOBJ) Equal(other *ParserOBJ) bool {

	if len(pO.Vertices) != len(other.Vertices) || len(pO.Triangles) != len(other.Triangles) {
		return false
	}

	for i, vertex := range pO.Vertices {
		if !vertex.Equal(other.Vertices[i]) {
			return false
		}
	}

	for i, triangle := range pO.Triangles {
		if !triangle.Equal(&other.Triangles[i]) {
			return false
		}
	}

	return true
}
