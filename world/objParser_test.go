package world

import (
	"testing"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
)

func TestParserObj(t *testing.T) {

	group1 := shapes.NewGroup()
	group1.AddChild(shapes.NewTriangle(pm.Point(-1, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0)))

	group2 := shapes.NewGroup()
	group2.AddChild(shapes.NewTriangle(pm.Point(-1, 1, 0), pm.Point(1, 0, 0), pm.Point(1, 1, 0)))

	tests := []struct {
		name  string
		input string
		want  *ParserOBJ
	}{
		{
			name:  "an empty string is parsed",
			input: "../obj-files/empty.obj",
			want:  &ParserOBJ{Vertices: []pm.Tuple{}, Triangles: []shapes.Triangle{}},
		},
		{
			name:  "gibberish string is parsed",
			input: "../obj-files/gibberish.obj",
			want:  &ParserOBJ{Vertices: []pm.Tuple{}, Triangles: []shapes.Triangle{}},
		},
		{
			name:  "parsing vertices",
			input: "../obj-files/verts.obj",
			want: &ParserOBJ{
				Vertices:  []pm.Tuple{pm.Point(1.5, 2, 1.3), pm.Point(1.4, -1.2, 0.12), pm.Point(-0.1, 0, -1.3)},
				Triangles: []shapes.Triangle{}},
		},
		{
			name:  "parsing vertices and a face (triangle)",
			input: "../obj-files/tris.obj",
			want:  &ParserOBJ{Vertices: []pm.Tuple{pm.Point(-1, -1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0), pm.Point(1, 1, 0)}, Triangles: []shapes.Triangle{*shapes.NewTriangle(pm.Point(-1, -1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0)), *shapes.NewTriangle(pm.Point(-1, -1, 0), pm.Point(1, 0, 0), pm.Point(1, 1, 0))}},
		},
		{
			name:  "parsing vertices and a fan face (triangle)",
			input: "../obj-files/fanTris.obj",
			want: &ParserOBJ{
				Vertices: []pm.Tuple{pm.Point(-1, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0), pm.Point(1, 1, 0), pm.Point(0, 2, 0)},
				Triangles: []shapes.Triangle{
					*shapes.NewTriangle(pm.Point(-1, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0)),
					*shapes.NewTriangle(pm.Point(-1, 1, 0), pm.Point(1, 0, 0), pm.Point(1, 1, 0)),
					*shapes.NewTriangle(pm.Point(-1, 1, 0), pm.Point(1, 1, 0), pm.Point(0, 2, 0)),
				},
			},
		},
		{
			name:  "parsing obj files with sub-groups",
			input: "../obj-files/groups.obj",
			want: &ParserOBJ{
				Vertices:  []pm.Tuple{pm.Point(-1, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0), pm.Point(1, 1, 0)},
				Triangles: []shapes.Triangle{},
				Groups:    []shapes.Group{*group1, *group2}},
		},
		{
			name:  "parsing obj files with normal vertices",
			input: "../obj-files/normalVerts.obj",
			want: &ParserOBJ{
				Vertices:  []pm.Tuple{pm.Point(1.5, 2, 1.3), pm.Point(1.4, -1.2, 0.12), pm.Point(-0.1, 0, -1.3)},
				Triangles: []shapes.Triangle{},
				Groups:    []shapes.Group{},
				Normals:   []pm.Tuple{pm.Point(1.5, 2, 1.3), pm.Point(1.4, -1.2, 0.12), pm.Point(-0.1, 0, -1.3)},
			},
		},
		{
			name:  "parsing obj files with normal vertices and smooth tris",
			input: "../obj-files/smoothTris.obj",
			want: &ParserOBJ{
				Vertices:  []pm.Tuple{pm.Point(0, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0)},
				Triangles: []shapes.Triangle{},
				STriangles: []shapes.SmoothTriangle{
					*shapes.NewSmoothTriangle(pm.Point(0, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0), pm.Vector(-1, 0, 0), pm.Vector(1, 0, 0), pm.Vector(0, 1, 0)),
					*shapes.NewSmoothTriangle(pm.Point(0, 1, 0), pm.Point(-1, 0, 0), pm.Point(1, 0, 0), pm.Vector(-1, 0, 0), pm.Vector(1, 0, 0), pm.Vector(0, 1, 0)),
				},
				Groups:  []shapes.Group{},
				Normals: []pm.Tuple{pm.Point(-1, 0, 0), pm.Point(1, 0, 0), pm.Point(0, 1, 0)},
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
