package primitive_math

import (
	"fmt"
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

var Epsilon = 0.00003

func IsPoint(tuple Tuple) bool {
	return tuple.W == 1
}

func IsVector(tuple Tuple) bool {
	return tuple.W == 0
}

func Point(X float64, Y float64, Z float64) Tuple {
	return Tuple{X, Y, Z, 1}
}

func Vector(X float64, Y float64, Z float64) Tuple {
	return Tuple{X, Y, Z, 0}
}

func AreFloatsEqual(first float64, second float64) bool {

	return math.Abs(float64(first-second)) < Epsilon
}

func (t *Tuple) Equal(compare Tuple) bool {
	return AreFloatsEqual(t.X, compare.X) &&
		AreFloatsEqual(t.Y, compare.Y) &&
		AreFloatsEqual(t.Z, compare.Z) &&
		AreFloatsEqual(t.W, compare.W)
}

func (t *Tuple) Add(addend Tuple) Tuple {
	return Tuple{t.X + addend.X, t.Y + addend.Y, t.Z + addend.Z, t.W + addend.W}
}

func (t *Tuple) Subtract(addend Tuple) Tuple {

	return Tuple{t.X - addend.X, t.Y - addend.Y, t.Z - addend.Z, t.W - addend.W}
}

func (t *Tuple) Negate() Tuple {
	return Tuple{
		X: -t.X,
		Y: -t.Y,
		Z: -t.Z,
		W: -t.W,
	}
}

func (t *Tuple) SMultiply(aFloat float64) Tuple {
	return Tuple{t.X * aFloat, t.Y * aFloat, t.Z * aFloat, t.W * aFloat}
}

func (t *Tuple) SDivide(aFloat float64) Tuple {
	return Tuple{t.X / aFloat, t.Y / aFloat, t.Z / aFloat, t.W / aFloat}
}

func (t *Tuple) Magnitude() float64 {
	return math.Sqrt((t.X * t.X) + (t.Y * t.Y) + (t.Z * t.Z) + (t.W * t.W))
}

func Normalize(t Tuple) Tuple {
	magnitude := t.Magnitude()

	return Tuple{
		t.X / magnitude,
		t.Y / magnitude,
		t.Z / magnitude,
		t.W / magnitude,
	}
}

func Dot(t1 Tuple, t2 Tuple) float64 {

	return t1.X*t2.X + t1.Y*t2.Y + t1.Z*t2.Z + t1.W*t2.W
}

func Cross(t1 Tuple, t2 Tuple) Tuple {
	return Vector(
		t1.Y*t2.Z-t1.Z*t2.Y,
		t1.Z*t2.X-t1.X*t2.Z,
		t1.X*t2.Y-t1.Y*t2.X,
	)
}

func (t Tuple) Print() string {

	if t.W == 0 {

		return fmt.Sprintf("Vector: %f, %f, %f", t.X, t.Y, t.Z)
	} else {
		return fmt.Sprintf("Point: %f, %f, %f", t.X, t.Y, t.Z)
	}

}

func (t Tuple) Reflect(normal Tuple) Tuple {
	return t.Subtract(normal.SMultiply(Dot(normal, t) * 2))
}
