package main

import (
	"math"
)

func (m44 Matrix4x4) Translate(x, y, z float64) Matrix4x4 {
	translationMatrix := NewMatrix4x4([16]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})

	return translationMatrix.Multiply(m44)
}

func (m44 Matrix4x4) TranslateInverse(x, y, z float64) Matrix4x4 {

	x *= -1
	y *= -1
	z *= -1

	translationMatrix := NewMatrix4x4([16]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})
	return translationMatrix.Multiply(m44)
}

func (m44 Matrix4x4) Scale(x, y, z float64) Matrix4x4 {
	scaleMatrix := NewMatrix4x4([16]float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})

	return scaleMatrix.Multiply(m44)

}

func (m44 Matrix4x4) ScaleInverse(x, y, z float64) Matrix4x4 {
	x = 1 / x
	y = 1 / y
	z = 1 / z

	scaleMatrix := NewMatrix4x4([16]float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})

	return scaleMatrix.Multiply(m44)

}

func (m44 Matrix4x4) ReflectX() Matrix4x4 {
	refmatrix := NewMatrix4x4([16]float64{-1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	return refmatrix.Multiply(m44)
}

func (m44 Matrix4x4) ReflectY() Matrix4x4 {
	refmatrix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, -1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	return refmatrix.Multiply(m44)
}

func (m44 Matrix4x4) ReflectZ() Matrix4x4 {
	refmatrix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -1, 0, 0, 0, 0, 1})

	return refmatrix.Multiply(m44)
}

func (m44 Matrix4x4) RotationAlongX(radians float64) Matrix4x4 {
	/*
		1, 0, 0 , 0,
		0, math.cos(radians), -math.sin(radians), 0,
		0, math.sin(radians), math.cos(radians), 0,
		0, 0, 0, 1
	*/
	rotmatrix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, math.Cos(radians), -math.Sin(radians), 0, 0, math.Sin(radians), math.Cos(radians), 0, 0, 0, 0, 1})
	return rotmatrix.Multiply(m44)
}

func (m44 Matrix4x4) RotationAlongY(radians float64) Matrix4x4 {
	/*
		math.cos(radians), 	0,	math.sin(radians) , 0,
		0, 					1, 	0, 					0,
		-math.sin(radians), 0, 	math.cos(radians), 	0,
		0, 					0, 	0, 					1
	*/
	rotmatrix := NewMatrix4x4([16]float64{math.Cos(radians), 0, math.Sin(radians), 0, 0, 1, 0, 0, -math.Sin(radians), 0, math.Cos(radians), 0, 0, 0, 0, 1})
	return rotmatrix.Multiply(m44)
}

func (m44 Matrix4x4) RotationAlongZ(radians float64) Matrix4x4 {
	/*
		math.Cos(radians), 		-math.Sin(radians),	0,	0,
		math.Sin(radians), 		math.Cos(radians),	0, 	0,
		0, 	0, 1, 0,
		0, 0, 0, 1
	*/
	rotmatrix := NewMatrix4x4([16]float64{math.Cos(radians), -math.Sin(radians), 0, 0, math.Sin(radians), math.Cos(radians), 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	return rotmatrix.Multiply(m44)
}

/*
Shearing is also called skewing.  Transformation has the effect of making straight lines slanted.  The x component changes in proportion to the other two components.  The x component will change in proportion to y and z.
*/
func (m44 Matrix4x4) Shear(xy, xz, yx, yz, zx, zy float64) Matrix4x4 {
	shearMatrix := NewMatrix4x4([16]float64{1, xy, xz, 0, yx, 1, yz, 0, zx, zy, 1, 0, 0, 0, 0, 1})

	return shearMatrix.Multiply(m44)
}
