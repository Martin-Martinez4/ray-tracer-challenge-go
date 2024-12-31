package materials

import (
	"fmt"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Color struct {
	R float64
	G float64
	B float64
}

var BLACK Color = NewColor(0, 0, 0)
var WHITE Color = NewColor(1, 1, 1)

func NewColor(r float64, g float64, b float64) Color {
	return Color{R: r, G: g, B: b}
}

func (c Color) Add(color Color) Color {
	return NewColor(c.R+color.R, c.G+color.G, c.B+color.B)
}

func (c Color) Subtract(color Color) Color {
	return NewColor(c.R-color.R, c.G-color.G, c.B-color.B)
}

func (c Color) SMultiply(scalar float64) Color {
	return NewColor(c.R*scalar, c.G*scalar, c.B*scalar)
}

func (c Color) Multiply(c2 Color) Color {
	return NewColor(c.R*c2.R, c.G*c2.G, c.B*c2.B)
}

func (c Color) Equal(c2 Color) bool {

	return pm.AreFloatsEqual(c.R, c2.R) && pm.AreFloatsEqual(c.G, c2.G) && pm.AreFloatsEqual(c.B, c2.B)
}

func (color Color) Print() string {
	return fmt.Sprintf("r: %f, g: %f, b: %f", color.R, color.G, color.B)
}

func Hadamard_product(c1 Color, c2 Color) Color {
	return NewColor(c1.R*c2.R, c1.G*c2.G, c1.B*c2.B)
}
