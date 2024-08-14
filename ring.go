package main

import (
	"math"
)

type Ring struct {
	Color1     Color
	Color2     Color
	Transforms Matrix4x4
}

func NewRing(color1, color2 Color) *Ring {

	return &Ring{Color1: color1, Color2: color2, Transforms: IdentitiyMatrix4x4()}
}

func (ring *Ring) GetColor1() Color {
	return ring.Color1
}

func (ring *Ring) GetColor2() Color {
	return ring.Color2
}

func (ring *Ring) SetTransform(mat44 *Matrix4x4) Matrix4x4 {

	ring.Transforms = mat44.Multiply(ring.Transforms)
	return ring.Transforms
}

func (ring *Ring) SetTransforms(mat44 []*Matrix4x4) {

	for _, transform := range mat44 {

		ring.SetTransform(transform)
	}
}

func (ring *Ring) GetTransforms() Matrix4x4 {
	return ring.Transforms
}

func (ring *Ring) PatternAt(point Tuple) Color {

	floored := math.Floor(math.Sqrt((point.x * point.x) + (point.z * point.z)))
	if int(floored)%2 == 0 {
		return ring.Color1
	} else {
		return ring.Color2
	}
}

func (ring *Ring) PatternAtShape(object Shape, worldPoint Tuple) Color {
	inversObjTransform := object.GetTransforms().Inverse()

	objectPoint := inversObjTransform.TupleMultiply(worldPoint)

	inversPatTransform := ring.GetTransforms().Inverse()
	patternPoint := inversPatTransform.TupleMultiply(objectPoint)

	return ring.PatternAt(patternPoint)

}
