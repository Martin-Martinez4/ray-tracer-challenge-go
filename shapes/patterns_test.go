package shapes

import (
	"testing"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

// It seems sphere is only being used as a dummy, maybe I could create a dummy shape and move the tests back to the materials module.

func TestCheckerPatternAt(t *testing.T) {
	tests := []struct {
		name              string
		pattern           mat.Pattern
		shape             Shape
		shapeTransforms   []*pm.Matrix4x4
		patternTransforms []*pm.Matrix4x4
		point             pm.Tuple
		want              mat.Color
	}{
		{
			name:              "pm.Point(0,0,0) should be white",
			pattern:           mat.NewChecker(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0, 0, 0),
			want:              mat.WHITE,
		},
		{
			name:              "pm.Point(0.99,0,0) should be a bit darker than white",
			pattern:           mat.NewChecker(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0.99, 0, 0),
			want:              mat.WHITE,
		},
		{
			name:              "pm.Point(1.01,0,0) should be gray",
			pattern:           mat.NewChecker(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(1.01, 0, 0),
			want:              mat.BLACK,
		},
		{
			name:              "pm.Point(0,0.99,0) should be a bit darker than white",
			pattern:           mat.NewChecker(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0, 0.99, 0),
			want:              mat.WHITE,
		},
		{
			name:              "pm.Point(0,1.01,0) should be gray",
			pattern:           mat.NewChecker(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0, 1.01, 0),
			want:              mat.BLACK,
		},
		{
			name:              "pm.Point(0,0,0.99) should be a bit darker than white",
			pattern:           mat.NewChecker(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0, 0, 0.99),
			want:              mat.WHITE,
		},
		{
			name:              "pm.Point(0,0,1.01) should be gray",
			pattern:           mat.NewChecker(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0, 0, 1.01),
			want:              mat.BLACK,
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

			got := tt.shape.PatternAtShape(tt.pattern, tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("\n%d %s failed:\nwanted: %s\ngot: %s\n", i, tt.name, tt.want.Print(), got.Print())
			}
		})

	}
}

func TestGradientPatternAt(t *testing.T) {
	tests := []struct {
		name              string
		pattern           mat.Pattern
		shape             Shape
		shapeTransforms   []*pm.Matrix4x4
		patternTransforms []*pm.Matrix4x4
		point             pm.Tuple
		want              mat.Color
	}{
		{
			name:              "pm.Point(0,0,0) should be white",
			pattern:           mat.NewGradient(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0, 0, 0),
			want:              mat.WHITE,
		},
		{
			name:              "pm.Point(0.25,0,0) should be a bit darker than white",
			pattern:           mat.NewGradient(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0.25, 0, 0),
			want:              mat.NewColor(0.75, 0.75, 0.75),
		},
		{
			name:              "pm.Point(0.25,0,0) should be gray",
			pattern:           mat.NewGradient(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0.5, 0, 0),
			want:              mat.NewColor(0.5, 0.5, 0.5),
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

			got := tt.shape.PatternAtShape(tt.pattern, tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("\n%d %s failed:\nwanted: %s\ngot: %s\n", i, tt.name, tt.want.Print(), got.Print())
			}
		})

	}
}

type testPattern struct {
	Color1     mat.Color
	Color2     mat.Color
	Transforms pm.Matrix4x4
}

func NewTestPattern(color1, color2 mat.Color) *testPattern {

	return &testPattern{Color1: color1, Color2: color2, Transforms: pm.IdentitiyMatrix4x4()}
}

func (testPattern *testPattern) GetColor1() mat.Color {
	return testPattern.Color1
}

func (testPattern *testPattern) GetColor2() mat.Color {
	return testPattern.Color2
}

func (testPattern *testPattern) SetTransform(mat44 *pm.Matrix4x4) pm.Matrix4x4 {

	testPattern.Transforms = mat44.Multiply(testPattern.Transforms)
	return testPattern.Transforms
}

func (testPattern *testPattern) SetTransforms(mat44 []*pm.Matrix4x4) {

	for _, transform := range mat44 {

		testPattern.SetTransform(transform)
	}
}

func (testPattern *testPattern) GetTransforms() pm.Matrix4x4 {
	return testPattern.Transforms
}

func (testPattern *testPattern) PatternAt(point pm.Tuple) mat.Color {

	return mat.NewColor(point.X, point.Y, point.Z)
}

func (testPattern *testPattern) PatternAtShape(object Shape, worldPoint pm.Tuple) mat.Color {
	inversObjTransform := object.GetTransforms().Inverse()

	objectPoint := inversObjTransform.TupleMultiply(worldPoint)

	inversPatTransform := testPattern.GetTransforms().Inverse()
	patternPoint := inversPatTransform.TupleMultiply(objectPoint)

	return testPattern.PatternAt(patternPoint)

}

func TestStripesWithTransformedObject(t *testing.T) {
	tests := []struct {
		name              string
		pattern           mat.Pattern
		shape             Shape
		shapeTransforms   []*pm.Matrix4x4
		patternTransforms []*pm.Matrix4x4
		point             pm.Tuple
		want              mat.Color
	}{
		{
			name:              "a pattern with an object transformation",
			pattern:           NewTestPattern(mat.BLACK, mat.WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{pm.Scale(2, 2, 2)},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(2, 3, 4),
			want:              mat.NewColor(1, 1.5, 2),
		},
		{
			name:              "a pattern with pattern transformation",
			pattern:           NewTestPattern(mat.BLACK, mat.WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{pm.Scale(2, 2, 2)},
			point:             pm.Point(2, 3, 4),
			want:              mat.NewColor(1, 1.5, 2),
		},
		{
			name:              "a pattern with both object and pattern transformation",
			pattern:           NewTestPattern(mat.BLACK, mat.WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{pm.Scale(2, 2, 2)},
			patternTransforms: []*pm.Matrix4x4{pm.Translate(.5, 1, 1.5)},
			point:             pm.Point(2.5, 3, 3.5),
			want:              mat.NewColor(0.75, 0.5, 0.25),
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

			got := tt.shape.PatternAtShape(tt.pattern, tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("\n%d %s failed:\nwanted: %s\ngot: %s\n", i, tt.name, tt.want.Print(), got.Print())
			}
		})

	}
}

func TestPatternAt(t *testing.T) {
	tests := []struct {
		name              string
		pattern           mat.Pattern
		shape             Shape
		shapeTransforms   []*pm.Matrix4x4
		patternTransforms []*pm.Matrix4x4
		point             pm.Tuple
		want              mat.Color
	}{
		{
			name:              "stripes with an object transformation",
			pattern:           mat.NewStripe(mat.BLACK, mat.WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{pm.Scale(2, 2, 2)},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(1.5, 0, 0),
			want:              mat.WHITE,
		},
		{
			name:              "stripes with pattern transformation",
			pattern:           mat.NewStripe(mat.BLACK, mat.WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{pm.Scale(2, 2, 2)},
			point:             pm.Point(1.5, 0, 0),
			want:              mat.WHITE,
		},
		{
			name:              "stripes with both object and pattern transformation",
			pattern:           mat.NewStripe(mat.BLACK, mat.WHITE),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{pm.Translate(0.5, 0, 0)},
			patternTransforms: []*pm.Matrix4x4{pm.Scale(2, 2, 2)},
			point:             pm.Point(1.5, 0, 0),
			want:              mat.WHITE,
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

			got := tt.shape.PatternAtShape(tt.pattern, tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("\n%d %s failed:\nwanted: %s\ngot: %s\n", i, tt.name, tt.want.Print(), got.Print())
			}
		})

	}
}

func TestRingPatternAt(t *testing.T) {
	tests := []struct {
		name              string
		pattern           mat.Pattern
		shape             Shape
		shapeTransforms   []*pm.Matrix4x4
		patternTransforms []*pm.Matrix4x4
		point             pm.Tuple
		want              mat.Color
	}{
		{
			name:              "pm.Point(0,0,0) should be white",
			pattern:           mat.NewRing(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0, 0, 0),
			want:              mat.WHITE,
		},
		{
			name:              "pm.Point(0.25,0,0) should be a bit darker than white",
			pattern:           mat.NewRing(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(1, 0, 0),
			want:              mat.BLACK,
		},
		{
			name:              "pm.Point(0.25,0,0) should be gray",
			pattern:           mat.NewRing(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0, 0, 1),
			want:              mat.BLACK,
		},
		{
			name:              "pm.Point(0.25,0,0) should be gray",
			pattern:           mat.NewRing(mat.WHITE, mat.BLACK),
			shape:             NewSphere(),
			shapeTransforms:   []*pm.Matrix4x4{},
			patternTransforms: []*pm.Matrix4x4{},
			point:             pm.Point(0.708, 0, 0.708),
			want:              mat.BLACK,
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

			got := tt.shape.PatternAtShape(tt.pattern, tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("\n%d %s failed:\nwanted: %s\ngot: %s\n", i, tt.name, tt.want.Print(), got.Print())
			}
		})

	}
}
