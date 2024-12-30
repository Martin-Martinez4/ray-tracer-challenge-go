package shapes

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/google/uuid"
)

type Cylinder struct {
	*PrimeShape
	Minimum float64
	Maximum float64
	Closed  bool
	Parent  Shape
	Bounds  *BoundingBox
}

func NewCylinder() *Cylinder {
	return &Cylinder{

		PrimeShape: CreateDefaultPrimeShape(),
		Minimum:    math.Inf(-1),
		Maximum:    math.Inf(1),
		Closed:     false,
		Parent:     nil,
		Bounds:     nil,
	}
}

func checkCap(ray Ray, t float64) bool {
	x := ray.origin.X + t*ray.direction.X
	z := ray.origin.Z + t*ray.direction.Z

	return ((x * x) + (z * z)) <= 1
}

func intersectCaps(cylinder *Cylinder, ray Ray, xs *Intersections) {
	if !cylinder.Closed || pm.AreFloatsEqual(ray.direction.Y, 0.0) {
		return
	}

	t := (cylinder.Minimum - ray.origin.Y) / ray.direction.Y
	if checkCap(ray, t) {
		xs.Add(NewIntersection(t, cylinder))
	}

	t = (cylinder.Maximum - ray.origin.Y) / ray.direction.Y
	if checkCap(ray, t) {
		xs.Add(NewIntersection(t, cylinder))
	}
}

func (cylinder *Cylinder) LocalIntersect(ray Ray) Intersections {

	intersections := Intersections{intersections: []Intersection{}}
	a := (ray.direction.X * ray.direction.X) + (ray.direction.Z * ray.direction.Z)

	intersectCaps(cylinder, ray, &intersections)
	if pm.AreFloatsEqual(a, 0.0) {
		return intersections
	}

	b := 2 * ((ray.origin.X * ray.direction.X) + (ray.origin.Z * ray.direction.Z))

	c := (ray.origin.X * ray.origin.X) + (ray.origin.Z * ray.origin.Z) - 1

	disc := (b * b) - (4 * a * c)

	if disc < 0.0 {
		return intersections
	}

	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)

	if t0 > t1 {
		tempt0 := t0
		t0 = t1
		t1 = tempt0
	}

	y0 := ray.origin.Y + t0*ray.direction.Y
	if cylinder.Minimum < y0 && y0 < cylinder.Maximum {
		intersections.Add(NewIntersection(t0, cylinder))
	}

	y1 := ray.origin.Y + t1*ray.direction.Y
	if cylinder.Minimum < y1 && y1 < cylinder.Maximum {
		intersections.Add(NewIntersection(t1, cylinder))
	}

	return intersections

}

func (cylinder *Cylinder) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(cylinder.Transforms.Inverse())
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

func (cylinder *Cylinder) GetSavedRay() Ray {
	return cylinder.SavedRay
}
func (cylinder *Cylinder) SetSavedRay(ray Ray) {
	cylinder.SavedRay = ray
}

func (cylinder *Cylinder) GetParent() Shape {
	return cylinder.Parent
}

func (cylinder *Cylinder) SetParent(shape Shape) {
	cylinder.Parent = shape
}

func (cylinder *Cylinder) GetId() uuid.UUID {
	return cylinder.id
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
