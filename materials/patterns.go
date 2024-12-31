package materials

import (
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

// Pattern will need to be a struct
// type Pattern struct {
// 	Color1     Color
// 	Color2     Color
// 	Transforms pm.Matrix4x4
// }

type Pattern interface {
	GetColor1() Color
	GetColor2() Color

	GetTransforms() pm.Matrix4x4
	SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4
	SetTransforms(mats []*pm.Matrix4x4)

	PatternAt(point pm.Tuple) Color
}

type testPattern struct {
	Color1     Color
	Color2     Color
	Transforms pm.Matrix4x4
}

func NewTestPattern(color1, color2 Color) *testPattern {

	return &testPattern{Color1: color1, Color2: color2, Transforms: pm.IdentitiyMatrix4x4()}
}

func (testPattern *testPattern) GetColor1() Color {
	return testPattern.Color1
}

func (testPattern *testPattern) GetColor2() Color {
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

func (testPattern *testPattern) PatternAt(point pm.Tuple) Color {

	return NewColor(point.X, point.Y, point.Z)
}
