package shapes

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Triangle struct {
	*PrimeShape
	P1     pm.Tuple
	P2     pm.Tuple
	P3     pm.Tuple
	Normal pm.Tuple

	E1 pm.Tuple
	E2 pm.Tuple

	Bounds *BoundingBox
}

/*
	type Shape interface {
	GetTransforms() pm.Matrix4x4
	SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4
	SetTransforms(transform []*pm.Matrix4x4)

	GetMaterial() *Material
	SetMaterial(material Material)

	Intersect(ray *Ray) Intersections

	NormalAt(point pm.Tuple) pm.Tuple

	GetSavedRay() Ray
	SetSavedRay(ray Ray)
}
*/

func NewTriangle(p1, p2, p3 pm.Tuple) *Triangle {

	// e1 = edge 1
	// p1 = point1
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)

	normal := pm.Normalize(pm.Cross(e2, e1))

	return &Triangle{
		PrimeShape: CreateDefaultPrimeShape(),
		P1:         p1,
		P2:         p2,
		P3:         p3,

		E1: e1,
		E2: e2,

		Normal: normal,
		Bounds: nil,
	}
}

func (triangle *Triangle) Equal(other *Triangle) bool {
	return triangle.P1.Equal(other.P1) && triangle.P2.Equal(other.P2) && triangle.P3.Equal(other.P3)
}

func (triangle *Triangle) LocalNormalAt(point pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	return triangle.Normal
}
func (triangle *Triangle) NormalAt(worldPoint pm.Tuple) pm.Tuple {
	return triangle.Normal
}

func (triangle *Triangle) LocalIntersect(ray Ray) Intersections {
	dirCrossE2 := pm.Cross(ray.Direction, triangle.E2)
	det := pm.Dot(triangle.E1, dirCrossE2)

	if math.Abs(det) < pm.Epsilon {
		return Intersections{Intersections: []Intersection{}}
	}

	f := 1.0 / det

	p1ToOrigin := ray.origin.Subtract(triangle.P1)
	u := f * pm.Dot(p1ToOrigin, dirCrossE2)

	if u < 0 || u > 1 {
		return Intersections{Intersections: []Intersection{}}
	}

	originCrossE1 := pm.Cross(p1ToOrigin, triangle.E1)
	v := f * pm.Dot(ray.Direction, originCrossE1)

	if v < 0 || (u+v) > 1 {
		return Intersections{Intersections: []Intersection{}}
	}

	t := f * pm.Dot(triangle.E2, originCrossE1)
	// will change later
	return Intersections{Intersections: []Intersection{NewIntersection(t, triangle)}}
}

func (triangle *Triangle) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(triangle.GetInverseTransforms())
	return triangle.LocalIntersect(tray)
}

func (triangle *Triangle) BoundingBox() *BoundingBox {
	if triangle.Bounds == nil {
		triangle.Bounds = &BoundingBox{
			Minimum: pm.Point(-1, -1, -1),
			Maximum: pm.Point(1, 1, 1),
		}
	}

	return triangle.Bounds
}
