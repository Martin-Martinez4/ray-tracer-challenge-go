package main

import (
	"fmt"
	"math"
)

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

var Epsilon = 0.00003

func IsPoint(tuple Tuple) bool {
	return tuple.w == 1
}

func IsVector(tuple Tuple) bool {
	return tuple.w == 0
}

func Point(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func Vector(x float64, y float64, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

func AreFloatsEqual(first float64, second float64) bool {

	return math.Abs(float64(first-second)) < Epsilon
}

func (t *Tuple) Equal(compare Tuple) bool {
	return AreFloatsEqual(t.x, compare.x) &&
		AreFloatsEqual(t.y, compare.y) &&
		AreFloatsEqual(t.z, compare.z) &&
		AreFloatsEqual(t.w, compare.w)
}

func (t *Tuple) Add(addend Tuple) Tuple {
	return Tuple{t.x + addend.x, t.y + addend.y, t.z + addend.z, t.w + addend.w}
}

func (t *Tuple) Subtract(addend Tuple) Tuple {

	return Tuple{t.x - addend.x, t.y - addend.y, t.z - addend.z, t.w - addend.w}
}

func (t *Tuple) Negate() Tuple {
	return Tuple{
		x: -t.x,
		y: -t.y,
		z: -t.z,
		w: -t.w,
	}
}

func (t *Tuple) SMultiply(aFloat float64) Tuple {
	return Tuple{t.x * aFloat, t.y * aFloat, t.z * aFloat, t.w * aFloat}
}

func (t *Tuple) SDivide(aFloat float64) Tuple {
	return Tuple{t.x / aFloat, t.y / aFloat, t.z / aFloat, t.w / aFloat}
}

func (t *Tuple) Magnitude() float64 {
	return math.Sqrt((t.x * t.x) + (t.y * t.y) + (t.z * t.z) + (t.w * t.w))
}

func Normalize(t Tuple) Tuple {
	magnitude := t.Magnitude()

	return Tuple{
		t.x / magnitude,
		t.y / magnitude,
		t.z / magnitude,
		t.w / magnitude,
	}
}

func Dot(t1 Tuple, t2 Tuple) float64 {

	return t1.x*t2.x + t1.y*t2.y + t1.z*t2.z + t1.w*t2.w
}

func Cross(t1 Tuple, t2 Tuple) Tuple {
	return Vector(
		t1.y*t2.z-t1.z*t2.y,
		t1.z*t2.x-t1.x*t2.z,
		t1.x*t2.y-t1.y*t2.x,
	)
}

func (t Tuple) Print() string {

	if t.w == 0 {

		return fmt.Sprintf("Vector: %f, %f, %f", t.x, t.y, t.z)
	} else {
		return fmt.Sprintf("Point: %f, %f, %f", t.x, t.y, t.z)
	}

}

func (t Tuple) Reflect(normal Tuple) Tuple {
	return t.Subtract(normal.SMultiply(Dot(normal, t) * 2))
}
