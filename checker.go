package main

import (
	"math"
)

type Checker struct {
	Color1     Color
	Color2     Color
	Transforms Matrix4x4
}

func NewChecker(color1, color2 Color) *Checker {

	return &Checker{Color1: color1, Color2: color2, Transforms: IdentitiyMatrix4x4()}
}

func (checker *Checker) GetColor1() Color {
	return checker.Color1
}

func (checker *Checker) GetColor2() Color {
	return checker.Color2
}

func (checker *Checker) SetTransform(mat44 *Matrix4x4) Matrix4x4 {

	checker.Transforms = mat44.Multiply(checker.Transforms)
	return checker.Transforms
}

func (checker *Checker) SetTransforms(mat44 []*Matrix4x4) {

	for _, transform := range mat44 {

		checker.SetTransform(transform)
	}
}

func (checker *Checker) GetTransforms() Matrix4x4 {
	return checker.Transforms
}

func (checker *Checker) PatternAt(point Tuple) Color {

	added := math.Floor(point.x) + math.Floor(point.y) + math.Floor(point.z)

	if math.Mod(added, 2) == 0 {
		return checker.Color1
	} else {
		return checker.Color2
	}
}

func (checker *Checker) PatternAtShape(object Shape, worldPoint Tuple) Color {
	inversObjTransform := object.GetTransforms().Inverse()

	objectPoint := inversObjTransform.TupleMultiply(worldPoint)

	inversPatTransform := checker.GetTransforms().Inverse()
	patternPoint := inversPatTransform.TupleMultiply(objectPoint)

	return checker.PatternAt(patternPoint)

}
