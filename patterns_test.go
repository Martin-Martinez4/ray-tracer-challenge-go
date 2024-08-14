package main

import (
	"testing"
)

type testPattern struct {
	Color1     Color
	Color2     Color
	Transforms Matrix4x4
}

func NewTestPattern(color1, color2 Color) *testPattern {

	return &testPattern{Color1: color1, Color2: color2, Transforms: IdentitiyMatrix4x4()}
}

func (testPattern *testPattern) GetColor1() Color {
	return testPattern.Color1
}

func (testPattern *testPattern) GetColor2() Color {
	return testPattern.Color2
}

func (testPattern *testPattern) SetTransform(mat44 *Matrix4x4) Matrix4x4 {

	testPattern.Transforms = mat44.Multiply(testPattern.Transforms)
	return testPattern.Transforms
}

func (testPattern *testPattern) SetTransforms(mat44 []*Matrix4x4) {

	for _, transform := range mat44 {

		testPattern.SetTransform(transform)
	}
}

func (testPattern *testPattern) GetTransforms() Matrix4x4 {
	return testPattern.Transforms
}

func (testPattern *testPattern) PatternAt(point Tuple) Color {

	return NewColor(point.x, point.y, point.z)
}

func (testPattern *testPattern) PatternAtShape(object Shape, worldPoint Tuple) Color {
	inversObjTransform := object.GetTransforms().Inverse()

	objectPoint := inversObjTransform.TupleMultiply(worldPoint)

	inversPatTransform := testPattern.GetTransforms().Inverse()
	patternPoint := inversPatTransform.TupleMultiply(objectPoint)

	return testPattern.PatternAt(patternPoint)

}

func TestStripesWithTransformedObject(t *testing.T) {
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
			name:              "a pattern with an object transformation",
			pattern:           NewTestPattern(BLACK, WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{Scale(2, 2, 2)},
			patternTransforms: []*Matrix4x4{},
			point:             Point(2, 3, 4),
			want:              NewColor(1, 1.5, 2),
		},
		{
			name:              "a pattern with pattern transformation",
			pattern:           NewTestPattern(BLACK, WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{},
			patternTransforms: []*Matrix4x4{Scale(2, 2, 2)},
			point:             Point(2, 3, 4),
			want:              NewColor(1, 1.5, 2),
		},
		{
			name:              "a pattern with both object and pattern transformation",
			pattern:           NewTestPattern(BLACK, WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{Scale(2, 2, 2)},
			patternTransforms: []*Matrix4x4{Translate(.5, 1, 1.5)},
			point:             Point(2.5, 3, 3.5),
			want:              NewColor(0.75, 0.5, 0.25),
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

func TestPatternAt(t *testing.T) {
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
			name:              "stripes with an object transformation",
			pattern:           NewStripe(BLACK, WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{Scale(2, 2, 2)},
			patternTransforms: []*Matrix4x4{},
			point:             Point(1.5, 0, 0),
			want:              WHITE,
		},
		{
			name:              "stripes with pattern transformation",
			pattern:           NewStripe(BLACK, WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{},
			patternTransforms: []*Matrix4x4{Scale(2, 2, 2)},
			point:             Point(1.5, 0, 0),
			want:              WHITE,
		},
		{
			name:              "stripes with both object and pattern transformation",
			pattern:           NewStripe(BLACK, WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*Matrix4x4{Translate(0.5, 0, 0)},
			patternTransforms: []*Matrix4x4{Scale(2, 2, 2)},
			point:             Point(1.5, 0, 0),
			want:              WHITE,
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
