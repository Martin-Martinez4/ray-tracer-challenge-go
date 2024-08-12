package main

import "fmt"

type Color struct {
	r float64
	g float64
	b float64
}

func NewColor(r float64, g float64, b float64) Color {
	return Color{r, g, b}
}

func (c Color) Add(color Color) Color {
	return NewColor(c.r+color.r, c.g+color.g, c.b+color.b)
}

func (c Color) Subtract(color Color) Color {
	return NewColor(c.r-color.r, c.g-color.g, c.b-color.b)
}

func (c Color) SMultiply(scalar float64) Color {
	return NewColor(c.r*scalar, c.g*scalar, c.b*scalar)
}

func (c Color) Multiply(c2 Color) Color {
	return NewColor(c.r*c2.r, c.g*c2.g, c.b*c2.b)
}

func (c Color) Equal(c2 Color) bool {

	return AreFloatsEqual(c.r, c2.r) && AreFloatsEqual(c.g, c2.g) && AreFloatsEqual(c.b, c2.b)
}

func (color Color) Print() string {
	return fmt.Sprintf("r: %f, g: %f, b: %f", color.r, color.g, color.b)
}

func Hadamard_product(c1 Color, c2 Color) Color {
	return NewColor(c1.r*c2.r, c1.g*c2.g, c1.b*c2.b)
}
