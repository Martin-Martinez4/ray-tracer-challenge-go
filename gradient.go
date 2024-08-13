package main

import "math"

type Gradient struct {
	Color1     Color
	Color2     Color
	Transforms Matrix4x4
}

func linearInterpolation(point Tuple, color1 Color, color2 Color) Color {
	distance := color2.Subtract(color1)
	fraction := point.x - math.Floor(point.x)

	return color1.Add(distance.SMultiply(fraction))

}

func NewGradient(color1, color2 Color) *Gradient {

	return &Gradient{Color1: color1, Color2: color2, Transforms: IdentitiyMatrix4x4()}
}

func (gradient *Gradient) GetColor1() Color {
	return gradient.Color1
}

func (gradient *Gradient) GetColor2() Color {
	return gradient.Color2
}

func (gradient *Gradient) SetTransform(mat44 *Matrix4x4) Matrix4x4 {

	gradient.Transforms = mat44.Multiply(gradient.Transforms)
	return gradient.Transforms
}

func (gradient *Gradient) SetTransforms(mat44 []Matrix4x4) {

	for _, transform := range mat44 {

		gradient.SetTransform(&transform)
	}
}

func (gradient *Gradient) GetTransforms() Matrix4x4 {
	return gradient.Transforms
}

func (gradient *Gradient) PatternAt(point Tuple) Color {

	return linearInterpolation(point, gradient.Color1, gradient.Color2)
}

func (gradient *Gradient) PatternAtShape(object Shape, worldPoint Tuple) Color {
	inversObjTransform := object.GetTransforms().Inverse()

	objectPoint := inversObjTransform.TupleMultiply(worldPoint)

	inversPatTransform := gradient.GetTransforms().Inverse()
	patternPoint := inversPatTransform.TupleMultiply(objectPoint)

	return gradient.PatternAt(patternPoint)

}
