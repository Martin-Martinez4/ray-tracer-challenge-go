package shapes

import (
	"math"

	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type Cone struct {
	*PrimeShape
	Minimum float64
	Maximum float64
	Closed  bool
	Bounds  *BoundingBox
}

func NewCone() *Cone {

	return &Cone{
		PrimeShape: CreateDefaultPrimeShape(),
		Minimum:    math.Inf(-1),
		Maximum:    math.Inf(1),
		Closed:     false,
		Bounds:     nil,
	}
}

func checkConeCap(ray Ray, t float64, y float64) bool {
	x := ray.origin.X + t*ray.Direction.X
	z := ray.origin.Z + t*ray.Direction.Z

	return x*x+z*z <= y*y
}

func intersectConeCaps(cone *Cone, ray Ray, xs *Intersections) {
	if !cone.Closed || pm.AreFloatsEqual(ray.Direction.Y, 0.0) {
		return
	}

	t := (cone.Minimum - ray.origin.Y) / ray.Direction.Y
	if checkConeCap(ray, t, cone.Minimum) {
		xs.Add(NewIntersection(t, cone))
	}

	t = (cone.Maximum - ray.origin.Y) / ray.Direction.Y
	if checkConeCap(ray, t, cone.Maximum) {
		xs.Add(NewIntersection(t, cone))
	}
}

func (cone *Cone) LocalIntersect(ray Ray) *Intersections {
	ray = ray.Transform(cone.Transforms.Inverse())

	intersections := Intersections{}
	a := (ray.Direction.X * ray.Direction.X) - (ray.Direction.Y * ray.Direction.Y) + (ray.Direction.Z * ray.Direction.Z)

	b := 2 * ((ray.origin.X * ray.Direction.X) - (ray.origin.Y * ray.Direction.Y) + (ray.origin.Z * ray.Direction.Z))

	c := (ray.origin.X * ray.origin.X) - (ray.origin.Y * ray.origin.Y) + (ray.origin.Z * ray.origin.Z)
	if pm.AreFloatsEqual(a, 0.0) {
		if pm.AreFloatsEqual(b, 0.0) {
			return &intersections
		}

		intersections.Add(Intersection{T: -c / (2 * b), S: cone})
	}
	intersectConeCaps(cone, ray, &intersections)

	disc := ((b * b) - (4 * a * c))

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
	if cone.Minimum < y0 && y0 < cone.Maximum {
		intersections.Add(NewIntersection(t0, cone))
	}

	y1 := ray.origin.Y + t1*ray.Direction.Y
	if cone.Minimum < y1 && y1 < cone.Maximum {
		intersections.Add(NewIntersection(t1, cone))
	}

	return &intersections

}

func (cone *Cone) Intersect(ray *Ray) *Intersections {
	tray := ray.Transform(cone.Transforms.Inverse())
	return cone.LocalIntersect(tray)
}

func (cone *Cone) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	dist := (localPoint.X * localPoint.X) + (localPoint.Z * localPoint.Z)
	if dist < 1 && localPoint.Y >= cone.Maximum-pm.Epsilon {
		return pm.Vector(0, 1, 0)
	} else if dist < 1 && localPoint.Y <= cone.Minimum+pm.Epsilon {
		return pm.Vector(0, -1, 0)
	} else {
		y := math.Sqrt((localPoint.X * localPoint.X) + (localPoint.Z * localPoint.Z))
		if localPoint.Y > 0 {
			y = -y
		}

		return pm.Vector(localPoint.X, y, localPoint.Z)
	}
}

func (cone *Cone) BoundingBox() *BoundingBox {
	if cone.Bounds == nil {
		a := math.Abs(cone.Minimum)
		b := math.Abs(cone.Maximum)
		limit := max(a, b)
		cone.Bounds = &BoundingBox{
			Minimum: pm.Point(-limit, cone.Minimum, -limit),
			Maximum: pm.Point(limit, cone.Maximum, limit),
		}
	}

	return cone.Bounds
}
