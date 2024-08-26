package main

import (
	"testing"
)

func TestParserObj(t *testing.T) {

	group1 := NewGroup()
	group1.AddChild(NewTriangle(Point(-1, 1, 0), Point(-1, 0, 0), Point(1, 0, 0)))

	group2 := NewGroup()
	group2.AddChild(NewTriangle(Point(-1, 1, 0), Point(1, 0, 0), Point(1, 1, 0)))

	tests := []struct {
		name  string
		input string
		want  *ParserOBJ
	}{
		{
			name:  "an empty string is parsed",
			input: "obj-files/empty.obj",
			want:  &ParserOBJ{Vertices: []Tuple{}, Triangles: []Triangle{}},
		},
		{
			name:  "gibberish string is parsed",
			input: "obj-files/gibberish.obj",
			want:  &ParserOBJ{Vertices: []Tuple{}, Triangles: []Triangle{}},
		},
		{
			name:  "parsing vertices",
			input: "obj-files/verts.obj",
			want: &ParserOBJ{
				Vertices:  []Tuple{Point(1.5, 2, 1.3), Point(1.4, -1.2, 0.12), Point(-0.1, 0, -1.3)},
				Triangles: []Triangle{}},
		},
		{
			name:  "parsing vertices and a face (triangle)",
			input: "obj-files/tris.obj",
			want:  &ParserOBJ{Vertices: []Tuple{Point(-1, -1, 0), Point(-1, 0, 0), Point(1, 0, 0), Point(1, 1, 0)}, Triangles: []Triangle{*NewTriangle(Point(-1, -1, 0), Point(-1, 0, 0), Point(1, 0, 0)), *NewTriangle(Point(-1, -1, 0), Point(1, 0, 0), Point(1, 1, 0))}},
		},
		{
			name:  "parsing vertices and a fan face (triangle)",
			input: "obj-files/fanTris.obj",
			want: &ParserOBJ{
				Vertices: []Tuple{Point(-1, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Point(1, 1, 0), Point(0, 2, 0)},
				Triangles: []Triangle{
					*NewTriangle(Point(-1, 1, 0), Point(-1, 0, 0), Point(1, 0, 0)),
					*NewTriangle(Point(-1, 1, 0), Point(1, 0, 0), Point(1, 1, 0)),
					*NewTriangle(Point(-1, 1, 0), Point(1, 1, 0), Point(0, 2, 0)),
				},
			},
		},
		{
			name:  "parsing obj files with sub-groups",
			input: "obj-files/groups.obj",
			want: &ParserOBJ{
				Vertices:  []Tuple{Point(-1, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Point(1, 1, 0)},
				Triangles: []Triangle{},
				Groups:    []Group{*group1, *group2}},
		},
		{
			name:  "parsing obj files with normal vertices",
			input: "obj-files/normalVerts.obj",
			want: &ParserOBJ{
				Vertices:  []Tuple{Point(1.5, 2, 1.3), Point(1.4, -1.2, 0.12), Point(-0.1, 0, -1.3)},
				Triangles: []Triangle{},
				Groups:    []Group{},
				Normals:   []Tuple{Point(1.5, 2, 1.3), Point(1.4, -1.2, 0.12), Point(-0.1, 0, -1.3)},
			},
		},
		{
			name:  "parsing obj files with normal vertices and smooth tris",
			input: "obj-files/smoothTris.obj",
			want: &ParserOBJ{
				Vertices:  []Tuple{Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0)},
				Triangles: []Triangle{},
				STriangles: []SmoothTriangle{
					*NewSmoothTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Vector(-1, 0, 0), Vector(1, 0, 0), Vector(0, 1, 0)),
					*NewSmoothTriangle(Point(0, 1, 0), Point(-1, 0, 0), Point(1, 0, 0), Vector(-1, 0, 0), Vector(1, 0, 0), Vector(0, 1, 0)),
				},
				Groups:  []Group{},
				Normals: []Tuple{Point(-1, 0, 0), Point(1, 0, 0), Point(0, 1, 0)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ParseObjFile(tt.input)
			if !got.Equal(tt.want) {
				t.Errorf("%s did not pass", tt.name)
			}

		})
	}
}
