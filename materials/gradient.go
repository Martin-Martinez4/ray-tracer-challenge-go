package materials

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Gradient struct {
	Color1     Color
	Color2     Color
	Transforms pm.Matrix4x4
}

func linearInterpolation(point pm.Tuple, color1 Color, color2 Color) Color {

	// t := 2*point.x - math.Floor(2*point.x)
	// if math.Abs(point.x-math.Floor(point.x)) < 0.5 {
	// 	c2Minusc1 := color2.Subtract(color1)

	// 	return color1.Add(c2Minusc1.SMultiply(t))
	// } else {
	// 	c1Minusc2 := color1.Subtract(color2)
	// 	return color2.Add(c1Minusc2.SMultiply(t))
	// }

	// Same as below
	// r := color1.r + fraction*(color2.r-color1.r)
	// g := color1.g + fraction*(color2.g-color1.g)
	// b := color1.b + fraction*(color2.b-color1.b)
	distance := color2.Subtract(color1)
	fraction := (point.X) - math.Floor(point.X)
	dstTimeFrac := distance.SMultiply(fraction)

	return color1.Add(dstTimeFrac)
}

func NewGradient(color1, color2 Color) *Gradient {

	return &Gradient{Color1: color1, Color2: color2, Transforms: pm.IdentitiyMatrix4x4()}
}

func (gradient *Gradient) GetColor1() Color {
	return gradient.Color1
}

func (gradient *Gradient) GetColor2() Color {
	return gradient.Color2
}

func (gradient *Gradient) SetTransform(mat44 *pm.Matrix4x4) pm.Matrix4x4 {

	gradient.Transforms = mat44.Multiply(gradient.Transforms)
	return gradient.Transforms
}

func (gradient *Gradient) SetTransforms(mat44 []*pm.Matrix4x4) {

	for _, transform := range mat44 {

		gradient.SetTransform(transform)
	}
}

func (gradient *Gradient) GetTransforms() pm.Matrix4x4 {
	return gradient.Transforms
}

func (gradient *Gradient) PatternAt(point pm.Tuple) Color {

	return linearInterpolation(point, gradient.Color1, gradient.Color2)
}
