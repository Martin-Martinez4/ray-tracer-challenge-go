package materials

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Checker struct {
	Color1     Color
	Color2     Color
	Transforms pm.Matrix4x4
}

func NewChecker(color1, color2 Color) *Checker {

	return &Checker{Color1: color1, Color2: color2, Transforms: pm.IdentitiyMatrix4x4()}
}

func (checker *Checker) GetColor1() Color {
	return checker.Color1
}

func (checker *Checker) GetColor2() Color {
	return checker.Color2
}

func (checker *Checker) SetTransform(mat44 *pm.Matrix4x4) pm.Matrix4x4 {

	checker.Transforms = mat44.Multiply(checker.Transforms)
	return checker.Transforms
}

func (checker *Checker) SetTransforms(mat44 []*pm.Matrix4x4) {

	for _, transform := range mat44 {

		checker.SetTransform(transform)
	}
}

func (checker *Checker) GetTransforms() pm.Matrix4x4 {
	return checker.Transforms
}

func (checker *Checker) PatternAt(point pm.Tuple) Color {

	added := math.Floor(point.X) + math.Floor(point.Y) + math.Floor(point.Z)

	if math.Mod(added, 2) == 0 {
		return checker.Color1
	} else {
		return checker.Color2
	}
}
