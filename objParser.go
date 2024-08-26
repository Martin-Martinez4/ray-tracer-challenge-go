package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
	Regex for Vertex
	v ((-?)[0-9]+(\.[0-9]+\s|\s)){2}((-?)[0-9]+(\.[0-9]+)?)\n
*/

type ParserOBJ struct {
	Vertices   []Tuple
	Triangles  []Triangle
	STriangles []SmoothTriangle
	Normals    []Tuple
	Groups     []Group
}

func parserVertices(contents string, parser *ParserOBJ) {
	// re := regexp.MustCompile(`v (?:-?[0-9]+(?:\.[0-9]+)?\s){2}-?[0-9]+(?:\.[0-9]+)?`)
	re2 := regexp.MustCompile(`(?:-?[0-9]+(?:\.[0-9]+)?)`)

	// for _, match := range re.FindAllStringSubmatch(contents, -1) {
	numbs := re2.FindAllStringSubmatch(contents, -1)

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

	if contents[1] == 'n' {

		parser.Normals = append(parser.Normals, vertex)
	} else {

		parser.Vertices = append(parser.Vertices, vertex)
	}

}

func parseFanTris(contents string, parser *ParserOBJ, gIndex int) {
	/*
		f 1 2 3 4 5
		f = face
		numbers will be digits becuase they correspond to the vertices in the Vertex slice

		f 1 2 3 4 5 should yield three tris
	*/
	// re := regexp.MustCompile(`f (?:[0-9]+\s){2,}(?:[0-9]+\n)`)
	// re2 := regexp.MustCompile(`[0-9]+`)
	re2 := regexp.MustCompile(`[0-9]+(?:\/\/[0-9]+|\/[0-9]+\/[0-9]+)?`)
	// for _, match := range re.FindAllStringSubmatch(contents, -1) {
	numbs := re2.FindAllStringSubmatch(contents, -1)

	if (len(numbs)-3)%2 != 0 && len(numbs) != 0 {
		panic("v length is incorrect")
	}

	for i := 1; i < len(numbs)-1; i++ {

		var numbs00 int
		var norm00 int
		var numbs01 int
		var norm01 int
		var numbs02 int
		var norm02 int
		smooth := false
		var err error

		if strings.Contains(numbs[0][0], "//") || strings.Contains(numbs[0][0], "/") {
			var sr []string
			if strings.Contains(numbs[0][0], "//") {
				sr = strings.Split(numbs[0][0], "//")

			} else {

				sr = strings.Split(numbs[0][0], "/")
			}

			numbs00, err = strconv.Atoi(strings.TrimSpace(sr[0]))
			if err != nil {
				// maybe reutrn error instead
				panic("failed to convert to int")
			}

			// the last one becuase the normal is the last one in 1/2/3
			norm00, err = strconv.Atoi(strings.TrimSpace(sr[len(sr)-1]))
			if err != nil {
				// maybe reutrn error instead
				panic("failed to convert to int")
			}
			smooth = true

		} else {
			numbs00, err = strconv.Atoi(strings.TrimSpace(numbs[0][0]))
			if err != nil {
				// maybe reutrn error instead
				panic("failed to convert to int 3")
			}
		}

		if strings.Contains(numbs[i][0], "//") || strings.Contains(numbs[i][0], "/") {
			var sr []string
			if strings.Contains(numbs[i][0], "//") {
				sr = strings.Split(numbs[i][0], "//")

			} else {

				sr = strings.Split(numbs[i][0], "/")
			}
			numbs01, err = strconv.Atoi(strings.TrimSpace(sr[0]))
			if err != nil {
				// maybe reutrn error instead
				panic("failed to convert to int")
			}

			norm01, err = strconv.Atoi(strings.TrimSpace(sr[len(sr)-1]))
			if err != nil {
				// maybe reutrn error instead
				panic("failed to convert to int")
			}
			smooth = true

		} else {

			numbs01, err = strconv.Atoi(strings.TrimSpace(numbs[i][0]))
			if err != nil {
				// maybe reutrn error instead
				panic("failed to convert to int")
			}
		}

		if strings.Contains(numbs[i+1][0], "//") || strings.Contains(numbs[i+1][0], "/") {
			var sr []string
			if strings.Contains(numbs[i+1][0], "//") {
				sr = strings.Split(numbs[i+1][0], "//")

			} else {

				sr = strings.Split(numbs[i+1][0], "/")
			}

			numbs02, err = strconv.Atoi(strings.TrimSpace(sr[0]))
			if err != nil {
				// maybe reutrn error instead
				panic("failed to convert to int")
			}

			norm02, err = strconv.Atoi(strings.TrimSpace(sr[len(sr)-1]))
			if err != nil {
				// maybe reutrn error instead
				panic("failed to convert to int")
			}
			smooth = true

		} else {

			numbs02, err = strconv.Atoi(strings.TrimSpace(numbs[i+1][0]))
			if err != nil {
				// maybe reutrn error instead
				panic("failed to convert to int")
			}
		}

		numbs00--
		numbs01--
		numbs02--

		norm00--
		norm01--
		norm02--

		verts := parser.Vertices

		norms := parser.Normals

		if smooth {

			stri := NewSmoothTriangle(verts[numbs00], verts[numbs01], verts[numbs02], norms[norm00], norms[norm01], norms[norm02])
			if gIndex == -1 {

				parser.STriangles = append(parser.STriangles, *stri)
			} else {
				parser.Groups[gIndex].AddChild(stri)
			}
		} else {
			tri := NewTriangle(verts[numbs00], verts[numbs01], verts[numbs02])
			if gIndex == -1 {

				parser.Triangles = append(parser.Triangles, *tri)
			} else {
				parser.Groups[gIndex].AddChild(tri)
			}
		}

	}

	// }

}

// func ParseObj(contents string) *ParserOBJ {

// 	parser := ParserOBJ{Vertices: []Tuple{}, Triangles: []Triangle{}}

// 	parserVertices(contents, &parser)
// 	parseFanTris(contents, &parser)

// 	return &parser
// }

func ParseObjFile(filePath string) *ParserOBJ {

	parser := ParserOBJ{Vertices: []Tuple{}, Triangles: []Triangle{}, Groups: []Group{}}

	file, err := os.Open(filePath)
	if err != nil {
		panic("opening OBJ file: " + err.Error())
	}
	defer file.Close()

	// start g counter at -1
	// if line starts with g add one to the g counter push a new group to the parserObj
	// if line starts with v run changed parseVertices,
	// if line starts with f run changed parseTris,
	// if g not equal -1 add the triangle to parserObj.Groups[g]

	gIndex := -1

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == 'g' {
			parser.Groups = append(parser.Groups, *NewGroup())
			gIndex++

		} else if line[0] == 'v' {
			parserVertices(line, &parser)

		} else if line[0] == 'f' {

			parseFanTris(line, &parser, gIndex)
		}
	}

	// fmt.Println(parser.PrintGroups())
	// later create a function that turns parsed data into a group, where groups are added to the childern of the main group along with the tris
	return &parser
}

func (pO *ParserOBJ) parserToGroup() *Group {
	group := NewGroup()

	for _, tri := range pO.Triangles {
		group.AddChild(&tri)
	}

	for _, group := range pO.Groups {
		group.AddChild(&group)
	}

	return group
}

func (pO *ParserOBJ) PrintGroups() string {
	var sb strings.Builder

	for i, group := range pO.Groups {
		sb.WriteString(fmt.Sprintf("\nGroup: %d\n", i))

		for j, child := range group.Children {
			sb.WriteString(fmt.Sprintf("\tchild %d: %v\n", j, child))

		}

	}

	return sb.String()
}

func (pO *ParserOBJ) Equal(other *ParserOBJ) bool {

	if len(pO.Vertices) != len(other.Vertices) || len(pO.Triangles) != len(other.Triangles) || len(pO.Normals) != len(other.Normals) || len(pO.STriangles) != len(other.STriangles) {
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

	for i, normals := range pO.Normals {
		if !normals.Equal(other.Normals[i]) {
			return false
		}
	}

	for i, smooth := range pO.STriangles {
		if !smooth.Equal(&other.STriangles[i]) {
			return false
		}
	}

	return true
}
