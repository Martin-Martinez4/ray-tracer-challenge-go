package main

// Pattern will need to be a struct
// type Pattern struct {
// 	Color1     Color
// 	Color2     Color
// 	Transforms Matrix4x4
// }

type Pattern interface {
	GetColor1() Color
	GetColor2() Color

	GetTransforms() Matrix4x4
	SetTransform(transform *Matrix4x4) Matrix4x4
	SetTransforms(mats []Matrix4x4)

	PatternAtShape(object Shape, worldPoint Tuple) Color
}
