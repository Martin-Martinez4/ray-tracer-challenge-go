package main

import "testing"

func TestCheckerPatternAt(t *testing.T) {
	tests := []struct {
		name              string
		pattern           Pattern
		shape             Shape
		shapeTransforms   []*Matrix4x4
		patternTransforms []*Matrix4x4
		point             Tuple
		want              Color
	}{
		{
			name:              "Point(0,0,0) should be white",
			pattern:           NewChecker(WHITE, BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{},
			patternTransforms: []*Matrix4x4{},
			point:             Point(0, 0, 0),
			want:              WHITE,
		},
		{
			name:              "Point(0.99,0,0) should be a bit darker than white",
			pattern:           NewChecker(WHITE, BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{},
			patternTransforms: []*Matrix4x4{},
			point:             Point(0.99, 0, 0),
			want:              WHITE,
		},
		{
			name:              "Point(1.01,0,0) should be gray",
			pattern:           NewChecker(WHITE, BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{},
			patternTransforms: []*Matrix4x4{},
			point:             Point(1.01, 0, 0),
			want:              BLACK,
		},
		{
			name:              "Point(0,0.99,0) should be a bit darker than white",
			pattern:           NewChecker(WHITE, BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{},
			patternTransforms: []*Matrix4x4{},
			point:             Point(0, 0.99, 0),
			want:              WHITE,
		},
		{
			name:              "Point(0,1.01,0) should be gray",
			pattern:           NewChecker(WHITE, BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{},
			patternTransforms: []*Matrix4x4{},
			point:             Point(0, 1.01, 0),
			want:              BLACK,
		},
		{
			name:              "Point(0,0,0.99) should be a bit darker than white",
			pattern:           NewChecker(WHITE, BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{},
			patternTransforms: []*Matrix4x4{},
			point:             Point(0, 0, 0.99),
			want:              WHITE,
		},
		{
			name:              "Point(0,0,1.01) should be gray",
			pattern:           NewChecker(WHITE, BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{},
			patternTransforms: []*Matrix4x4{},
			point:             Point(0, 0, 1.01),
			want:              BLACK,
		},
	}

	for i, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if len(tt.shapeTransforms) > 0 {

				tt.shape.SetTransforms(tt.shapeTransforms)
			}

			if len(tt.patternTransforms) > 0 {
				tt.pattern.SetTransforms(tt.patternTransforms)
			}

			got := tt.pattern.PatternAtShape(tt.shape, tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("\n%d %s failed:\nwanted: %s\ngot: %s\n", i, tt.name, tt.want.Print(), got.Print())
			}
		})

	}
}
