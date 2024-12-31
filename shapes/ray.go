package shapes

import (
	"math"
	"strings"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Ray struct {
	origin    pm.Tuple
	Direction pm.Tuple
}

func NewRay(origin, Direction [3]float64) Ray {

	return Ray{origin: pm.Point(origin[0], origin[1], origin[2]), Direction: pm.Vector(Direction[0], Direction[1], Direction[2])}

}

/*
Multiply Direction by t then add the origin
*/
func (ray Ray) Position(t float64) pm.Tuple {
	dt := ray.Direction.SMultiply(t)
	return ray.origin.Add(dt)
}

func (ray Ray) Equal(other Ray) bool {
	return ray.origin.Equal(other.origin) && ray.Direction.Equal(other.Direction)
}

func (ray Ray) Print() string {
	var sb strings.Builder

	sb.WriteString("\nOrigin: " + ray.origin.Print())
	sb.WriteString("\nDirection: " + ray.Direction.Print() + "\n")

	return sb.String()

}

func (ray Ray) Translate(x, y, z float64) Ray {

	translationMatrix := pm.NewMatrix4x4([16]float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})

	newPoint := translationMatrix.TupleMultiply(ray.origin)

	rayDirection := ray.Direction

	return NewRay([3]float64{newPoint.X, newPoint.Y, newPoint.Z}, [3]float64{rayDirection.X, rayDirection.Y, rayDirection.Z})

}

func (ray Ray) Scale(x, y, z float64) Ray {
	scaleMatrix := pm.NewMatrix4x4([16]float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})

	newOrigin := scaleMatrix.TupleMultiply(ray.origin)
	newDirection := scaleMatrix.TupleMultiply(ray.Direction)

	return NewRay([3]float64{newOrigin.X, newOrigin.Y, newOrigin.Z}, [3]float64{newDirection.X, newDirection.Y, newDirection.Z})
}

func (ray Ray) Transform(m44 pm.Matrix4x4) Ray {
	return Ray{
		origin:    m44.TupleMultiply(ray.origin),
		Direction: m44.TupleMultiply(ray.Direction),
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
