package main

import (
	"math"
)

func (tuple Tuple) Translate(x, y, z float64) Tuple {
	translationMatrix := NewMatrix4x4([16]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})

	return translationMatrix.TupleMultiply(tuple)
}

func (tuple Tuple) TranslateInverse(x, y, z float64) Tuple {

	x *= -1
	y *= -1
	z *= -1

	translationMatrix := NewMatrix4x4([16]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})
	return translationMatrix.TupleMultiply(tuple)
}

func (tuple Tuple) Scale(x, y, z float64) Tuple {
	scaleMatrix := NewMatrix4x4([16]float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})

	return scaleMatrix.TupleMultiply(tuple)

}

func (tuple Tuple) ScaleInverse(x, y, z float64) Tuple {
	x = 1 / x
	y = 1 / y
	z = 1 / z

	scaleMatrix := NewMatrix4x4([16]float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})

	return scaleMatrix.TupleMultiply(tuple)

}

func (tuple Tuple) ReflectX() Tuple {
	refmatrix := NewMatrix4x4([16]float64{-1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	return refmatrix.TupleMultiply(tuple)
}

func (tuple Tuple) ReflectY() Tuple {
	refmatrix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, -1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	return refmatrix.TupleMultiply(tuple)
}

func (tuple Tuple) ReflectZ() Tuple {
	refmatrix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, -1, 0, 0, 0, 0, 1})

	return refmatrix.TupleMultiply(tuple)
}

func (tuple Tuple) RotationAlongX(radians float64) Tuple {
	/*
		1, 0, 0 , 0,
		0, math.cos(radians), -math.sin(radians), 0,
		0, math.sin(radians), math.cos(radians), 0,
		0, 0, 0, 1
	*/
	rotmatrix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, math.Cos(radians), -math.Sin(radians), 0, 0, math.Sin(radians), math.Cos(radians), 0, 0, 0, 0, 1})
	return rotmatrix.TupleMultiply(tuple)
}

func (tuple Tuple) RotationAlongY(radians float64) Tuple {
	/*
		math.cos(radians), 	0,	math.sin(radians) , 0,
		0, 					1, 	0, 					0,
		-math.sin(radians), 0, 	math.cos(radians), 	0,
		0, 					0, 	0, 					1
	*/
	rotmatrix := NewMatrix4x4([16]float64{math.Cos(radians), 0, math.Sin(radians), 0, 0, 1, 0, 0, -math.Sin(radians), 0, math.Cos(radians), 0, 0, 0, 0, 1})
	return rotmatrix.TupleMultiply(tuple)
}

func (tuple Tuple) RotationAlongZ(radians float64) Tuple {
	/*
		math.Cos(radians), 		-math.Sin(radians),	0,	0,
		math.Sin(radians), 		math.Cos(radians),	0, 	0,
		0, 	0, 1, 0,
		0, 0, 0, 1
	*/
	rotmatrix := NewMatrix4x4([16]float64{math.Cos(radians), -math.Sin(radians), 0, 0, math.Sin(radians), math.Cos(radians), 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
	return rotmatrix.TupleMultiply(tuple)
}

/*
Shearing is also called skewing.  Transformation has the effect of making straight lines slanted.  The x component changes in proportion to the other two components.  The x component will change in proportion to y and z.
*/
func (tuple Tuple) Shear(xy, xz, yx, yz, zx, zy float64) Tuple {
	shearMatrix := NewMatrix4x4([16]float64{1, xy, xz, 0, yx, 1, yz, 0, zx, zy, 1, 0, 0, 0, 0, 1})

	return shearMatrix.TupleMultiply(tuple)
}

func (tuple Tuple) ReflectBy(normalVector Tuple) Tuple {

	sMult := (normalVector.SMultiply(2))
	dot := Dot(tuple, normalVector)

	return tuple.Subtract(sMult.SMultiply(dot))
}
