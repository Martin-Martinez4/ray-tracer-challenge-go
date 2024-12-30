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
