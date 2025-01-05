package shapes

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Cylinder struct {
	*PrimeShape
	Minimum float64
	Maximum float64
	Closed  bool
	Bounds  *BoundingBox
}

func NewCylinder() *Cylinder {
	return &Cylinder{

		PrimeShape: CreateDefaultPrimeShape(),
		Minimum:    math.Inf(-1),
		Maximum:    math.Inf(1),
		Closed:     false,
		Bounds:     nil,
	}
}

func checkCap(ray Ray, t float64) bool {
	x := ray.origin.X + t*ray.Direction.X
	z := ray.origin.Z + t*ray.Direction.Z

	return ((x * x) + (z * z)) <= 1
}

func intersectCaps(cylinder *Cylinder, ray Ray, xs *Intersections) {
	if !cylinder.Closed || pm.AreFloatsEqual(ray.Direction.Y, 0.0) {
		return
	}

	t := (cylinder.Minimum - ray.origin.Y) / ray.Direction.Y
	if checkCap(ray, t) {
		xs.Add(NewIntersection(t, cylinder))
	}

	t = (cylinder.Maximum - ray.origin.Y) / ray.Direction.Y
	if checkCap(ray, t) {
		xs.Add(NewIntersection(t, cylinder))
	}
}

func (cylinder *Cylinder) LocalIntersect(ray Ray) *Intersections {

	intersections := Intersections{}
	a := (ray.Direction.X * ray.Direction.X) + (ray.Direction.Z * ray.Direction.Z)

	intersectCaps(cylinder, ray, &intersections)
	if pm.AreFloatsEqual(a, 0.0) {
		return &intersections
	}

	b := 2 * ((ray.origin.X * ray.Direction.X) + (ray.origin.Z * ray.Direction.Z))

	c := (ray.origin.X * ray.origin.X) + (ray.origin.Z * ray.origin.Z) - 1

	disc := (b * b) - (4 * a * c)

	if disc < 0.0 {
		return &intersections
	}

	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)

	if t0 > t1 {
		tempt0 := t0
		t0 = t1
		t1 = tempt0
	}

	y0 := ray.origin.Y + t0*ray.Direction.Y
	if cylinder.Minimum < y0 && y0 < cylinder.Maximum {
		intersections.Add(NewIntersection(t0, cylinder))
	}

	y1 := ray.origin.Y + t1*ray.Direction.Y
	if cylinder.Minimum < y1 && y1 < cylinder.Maximum {
		intersections.Add(NewIntersection(t1, cylinder))
	}

	return &intersections

}

func (cylinder *Cylinder) Intersect(ray *Ray) *Intersections {
	tray := ray.Transform(cylinder.GetInverseTransforms())
	return cylinder.LocalIntersect(tray)
}

func (cylinder *Cylinder) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	dist := (localPoint.X * localPoint.X) + (localPoint.Z * localPoint.Z)
	if dist < 1 && localPoint.Y >= cylinder.Maximum-pm.Epsilon {
		return pm.Vector(0, 1, 0)
	} else if dist < 1 && localPoint.Y <= cylinder.Minimum+pm.Epsilon {
		return pm.Vector(0, -1, 0)
	} else {

		return pm.Vector(localPoint.X, 0, localPoint.Z)
	}
}

func (cylinder *Cylinder) NormalAt(worldPoint pm.Tuple) pm.Tuple {
	return cylinder.LocalNormalAt(worldPoint, nil, nil)
}

func (cylinder *Cylinder) BoundingBox() *BoundingBox {
	if cylinder.Bounds == nil {
		cylinder.Bounds = &BoundingBox{
			Minimum: pm.Point(-1, cylinder.Minimum, -1),
			Maximum: pm.Point(1, cylinder.Maximum, 1),
		}
	}

	return cylinder.Bounds
}
