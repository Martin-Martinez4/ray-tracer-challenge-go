package main

import (
	"math"
	"strings"
)

type Ray struct {
	origin    Tuple
	direction Tuple
}

func NewRay(origin, direction [3]float64) Ray {

	return Ray{origin: Point(origin[0], origin[1], origin[2]), direction: Vector(direction[0], direction[1], direction[2])}

}

/*
Multiply direction by t then add the origin
*/
func (ray Ray) Position(t float64) Tuple {
	dt := ray.direction.SMultiply(t)
	return ray.origin.Add(dt)
}

func (ray Ray) Equal(other Ray) bool {
	return ray.origin.Equal(other.origin) && ray.direction.Equal(other.direction)
}

func (ray Ray) Print() string {
	var sb strings.Builder

	sb.WriteString("\nOrigin: " + ray.origin.Print())
	sb.WriteString("\nDirection: " + ray.direction.Print() + "\n")

	return sb.String()

}

func (ray Ray) Translate(x, y, z float64) Ray {

	translationMatrix := NewMatrix4x4([16]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})

	newPoint := translationMatrix.TupleMultiply(ray.origin)

	rayDirection := ray.direction

	return NewRay([3]float64{newPoint.x, newPoint.y, newPoint.z}, [3]float64{rayDirection.x, rayDirection.y, rayDirection.z})

}

func (ray Ray) Scale(x, y, z float64) Ray {
	scaleMatrix := NewMatrix4x4([16]float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})

	newOrigin := scaleMatrix.TupleMultiply(ray.origin)
	newDirection := scaleMatrix.TupleMultiply(ray.direction)

	return NewRay([3]float64{newOrigin.x, newOrigin.y, newOrigin.z}, [3]float64{newDirection.x, newDirection.y, newDirection.z})
}

func (ray Ray) Transform(m44 Matrix4x4) Ray {
	return Ray{
		origin:    m44.TupleMultiply(ray.origin),
		direction: m44.TupleMultiply(ray.direction),
	}
}

// Hit finds the first intersection with a positive T (the passed intersections are assumed to have been sorted already)
func Hit(intersections []Intersection) (Intersection, bool) {

	lowestNonNegative := Intersection{T: math.MaxFloat64, S: nil}
	for _, intersection := range intersections {
		if intersection.T > 0 && intersection.T < lowestNonNegative.T {
			lowestNonNegative = intersection
		}
	}
	if lowestNonNegative.T < math.MaxFloat64 {
		return lowestNonNegative, true
	} else {
		return lowestNonNegative, false
	}
	// inters := intersections.intersections

	// Filter out all negatives
	// for i := 0; i < len(inters); i++ {

	// 	if inters[i].T > 0 {
	// 		return inters[i], true
	// 		//xs = append(xs, i)
	// 	}
	// }

	// return Intersection{}, false
}
