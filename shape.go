package main

import (
	"math"

	"github.com/google/uuid"
)

/*
	Shape will have all the functions needed
	The concrete implementations will have the actual struct
*/
/*
	Color:     NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
*/
type Shape interface {
	GetId() uuid.UUID
	GetTransforms() Matrix4x4
	SetTransform(transform *Matrix4x4) Matrix4x4
	SetTransforms(transform []*Matrix4x4)

	GetMaterial() *Material
	SetMaterial(material Material)

	Intersect(ray *Ray) Intersections

	// NormalAt(point Tuple) Tuple
	LocalNormalAt(point Tuple) Tuple

	GetSavedRay() Ray
	SetSavedRay(ray Ray)

	GetParent() Shape
}

// Created functions that return transforms as a Matrix4x4, aids with using the SetTransforms function

func Translate(x, y, z float64) *Matrix4x4 {
	translationMatrix := NewMatrix4x4([16]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})

	return &translationMatrix
}

func TranslateInverse(x, y, z float64) Matrix4x4 {

	x *= -1
	y *= -1
	z *= -1

	translationMatrix := NewMatrix4x4([16]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})

	return translationMatrix
}

func Scale(x, y, z float64) *Matrix4x4 {
	scaleMatrix := NewMatrix4x4([16]float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})

	return &scaleMatrix

}

func ScaleInverse(x, y, z float64) Matrix4x4 {
	x = 1 / x
	y = 1 / y
	z = 1 / z

	scaleMatrix := NewMatrix4x4([16]float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})

	return scaleMatrix

}

func ReflectX() Matrix4x4 {
	refmatrix := NewMatrix4x4([16]float64{-1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	return refmatrix
}

func ReflectY() Matrix4x4 {
	refmatrix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, -1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	return refmatrix
}

func ReflectZ() Matrix4x4 {
	refmatrix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -1, 0, 0, 0, 0, 1})

	return refmatrix
}

func RotationAlongX(radians float64) *Matrix4x4 {
	/*
		1, 0, 0 , 0,
		0, math.cos(radians), -math.sin(radians), 0,
		0, math.sin(radians), math.cos(radians), 0,
		0, 0, 0, 1
	*/
	rotmatrix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, math.Cos(radians), -math.Sin(radians), 0, 0, math.Sin(radians), math.Cos(radians), 0, 0, 0, 0, 1})
	return &rotmatrix
}

func RotationAlongY(radians float64) *Matrix4x4 {
	/*
		math.cos(radians), 	0,	math.sin(radians) , 0,
		0, 					1, 	0, 					0,
		-math.sin(radians), 0, 	math.cos(radians), 	0,
		0, 					0, 	0, 					1
	*/
	rotmatrix := NewMatrix4x4([16]float64{math.Cos(radians), 0, math.Sin(radians), 0, 0, 1, 0, 0, -math.Sin(radians), 0, math.Cos(radians), 0, 0, 0, 0, 1})
	return &rotmatrix
}

func RotationAlongZ(radians float64) *Matrix4x4 {
	/*
		math.Cos(radians), 		-math.Sin(radians),	0,	0,
		math.Sin(radians), 		math.Cos(radians),	0, 	0,
		0, 	0, 1, 0,
		0, 0, 0, 1
	*/
	rotmatrix := NewMatrix4x4([16]float64{math.Cos(radians), -math.Sin(radians), 0, 0, math.Sin(radians), math.Cos(radians), 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	return &rotmatrix
}

/*
Shearing is also called skewing.  Transformation has the effect of making straight lines slanted.  The x component changes in proportion to the other two components.  The x component will change in proportion to y and z.
*/
func Shear(xy, xz, yx, yz, zx, zy float64) Matrix4x4 {
	shearMatrix := NewMatrix4x4([16]float64{1, xy, xz, 0, yx, 1, yz, 0, zx, zy, 1, 0, 0, 0, 0, 1})

	return shearMatrix
}
