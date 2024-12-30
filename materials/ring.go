package materials

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Ring struct {
	Color1     Color
	Color2     Color
	Transforms pm.Matrix4x4
}

func NewRing(color1, color2 Color) *Ring {

	return &Ring{Color1: color1, Color2: color2, Transforms: pm.IdentitiyMatrix4x4()}
}

func (ring *Ring) GetColor1() Color {
	return ring.Color1
}

func (ring *Ring) GetColor2() Color {
	return ring.Color2
}

func (ring *Ring) SetTransform(mat44 *pm.Matrix4x4) pm.Matrix4x4 {

	ring.Transforms = mat44.Multiply(ring.Transforms)
	return ring.Transforms
}

func (ring *Ring) SetTransforms(mat44 []*pm.Matrix4x4) {

	for _, transform := range mat44 {

		ring.SetTransform(transform)
	}
}

func (ring *Ring) GetTransforms() pm.Matrix4x4 {
	return ring.Transforms
}

func (ring *Ring) PatternAt(point pm.Tuple) Color {

	floored := math.Floor(math.Sqrt((point.X * point.X) + (point.Z * point.Z)))
	if int(floored)%2 == 0 {
		return ring.Color1
	} else {
		return ring.Color2
	}
}
