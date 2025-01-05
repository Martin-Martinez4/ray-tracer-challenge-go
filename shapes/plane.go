package shapes

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

// type Shape interface {
// 	GetTransforms() pm.Matrix4x4
// 	SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4
// 	SetTransforms(transform []pm.Matrix4x4)

// 	GetMaterial() *Material
// 	SetMaterial(material Material)

// 	Intersect(ray *Ray) Intersections

// 	NormalAt(point pm.Tuple) pm.Tuple

// 	GetSavedRay() Ray
// 	SetSavedRay(ray Ray)
// }

type Plane struct {
	*PrimeShape
	SavedRay Ray
	Bounds   *BoundingBox
}

func NewPlane() *Plane {

	return &Plane{
		PrimeShape: CreateDefaultPrimeShape(),
		Bounds:     nil,
	}
}

func (plane *Plane) LocalIntersect(ray Ray) *Intersections {
	if math.Abs(ray.Direction.Y) < pm.Epsilon {
		return &Intersections{}
	}

	t := (-ray.origin.Y / ray.Direction.Y)

	return &Intersections{{T: t, S: plane}}
}

func (plane *Plane) Intersect(ray *Ray) *Intersections {
	tray := ray.Transform(plane.GetInverseTransforms())
	return plane.LocalIntersect(tray)
}

func (plane *Plane) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	return pm.Vector(0, 1, 0)
}

func (plane *Plane) NormalAt(worldPoint pm.Tuple) pm.Tuple {
	return plane.LocalNormalAt(worldPoint, nil, nil)
}

func (plane *Plane) BoundingBox() *BoundingBox {
	if plane.Bounds == nil {
		plane.Bounds = &BoundingBox{
			Minimum: pm.Point(math.Inf(-1), 0, math.Inf(-1)),
			Maximum: pm.Point(math.Inf(1), 0, math.Inf(1)),
		}
	}

	return plane.Bounds
}
