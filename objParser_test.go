package main

import "testing"

func TestParserObj(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *ParserOBJ
	}{
		{
			name:  "an empty string is parsed",
			input: "",
			want:  &ParserOBJ{Vertices: []Tuple{}, Triangles: []Triangle{}},
		},
		{
			name:  "gibberish string is parsed",
			input: "there was a young lady named Bright\nwho traveled much faster than light\nShe set out one day\nin a relative way\n",
			want:  &ParserOBJ{Vertices: []Tuple{}, Triangles: []Triangle{}},
		},
		{
			name:  "parsing vertices",
			input: "v 1.5 2 1.3\nv 1.4 -1.2 0.12\nv -0.1 0 -1.3",
			want:  &ParserOBJ{Vertices: []Tuple{Point(1.5, 2, 1.3), Point(1.4, -1.2, 0.12), Point(-0.1, 0, -1.3)}, Triangles: []Triangle{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := ParseObj(tt.input)
			if !got.Equal(tt.want) {
				t.Errorf("%s did not pass", tt.name)
			}

		})
	}
}
