package materials

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Stripes struct {
	Color1     Color
	Color2     Color
	Transforms pm.Matrix4x4
}

func NewStripe(color1, color2 Color) *Stripes {

	return &Stripes{Color1: color1, Color2: color2, Transforms: pm.IdentitiyMatrix4x4()}
}

func (stripe *Stripes) GetColor1() Color {
	return stripe.Color1
}

func (stripe *Stripes) GetColor2() Color {
	return stripe.Color2
}

func (stripe *Stripes) SetTransform(mat44 *pm.Matrix4x4) pm.Matrix4x4 {

	stripe.Transforms = mat44.Multiply(stripe.Transforms)
	return stripe.Transforms
}

func (stripe *Stripes) SetTransforms(mat44 []*pm.Matrix4x4) {

	for _, transform := range mat44 {

		stripe.SetTransform(transform)
	}
}

func (stripe *Stripes) GetTransforms() pm.Matrix4x4 {
	return stripe.Transforms
}

func (stripe *Stripes) PatternAt(point pm.Tuple) Color {

	if int(math.Floor(point.X))%2 == 0 {
		return stripe.Color2
	}
	return stripe.Color1
}
