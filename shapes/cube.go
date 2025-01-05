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

type Cube struct {
	*PrimeShape
	Bounds *BoundingBox
}

func NewCube() *Cube {

	return &Cube{
		PrimeShape: CreateDefaultPrimeShape(),
	}
}

func CheckAxis(origin, Direction float64) (float64, float64) {
	tminNumerator := (-1 - origin)
	tmaxNumerator := (1 - origin)

	var tmin float64
	var tmax float64

	if math.Abs(Direction) >= pm.Epsilon {
		tmin = tminNumerator / Direction
		tmax = tmaxNumerator / Direction
	} else {
		tmin = tminNumerator * math.Inf(1)
		tmax = tmaxNumerator * math.Inf(1)
	}

	if tmin > tmax {
		minTemp := tmin
		tmin = tmax
		tmax = minTemp
	}

	return tmin, tmax
}

func (cube *Cube) LocalIntersect(ray Ray) *Intersections {
	xtmin, xtmax := CheckAxis(ray.origin.X, ray.Direction.X)
	ytmin, ytmax := CheckAxis(ray.origin.Y, ray.Direction.Y)
	ztmin, ztmax := CheckAxis(ray.origin.Z, ray.Direction.Z)

	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	if tmin > tmax {
		return &Intersections{}

	}

	return &Intersections{NewIntersection(tmin, cube), NewIntersection(tmax, cube)}
}

func (cube *Cube) Intersect(ray *Ray) *Intersections {
	tray := ray.Transform(cube.GetInverseTransforms())
	return cube.LocalIntersect(tray)
}

func (cube *Cube) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {

	absX := math.Abs(localPoint.X)
	absY := math.Abs(localPoint.Y)
	absZ := math.Abs(localPoint.Z)

	maxc := math.Max(math.Max(absX, absY), absZ)

	if maxc == absX {
		return pm.Vector(localPoint.X, 0, 0)
	} else if maxc == absY {
		return pm.Vector(0, localPoint.Y, 0)
	} else {
		return pm.Vector(0, 0, localPoint.Z)
	}
}

func (cube *Cube) NormalAt(worldPoint pm.Tuple) pm.Tuple {
	return cube.LocalNormalAt(worldPoint, nil, nil)
}

func (cube *Cube) BoundingBox() *BoundingBox {
	if cube.Bounds == nil {
		cube.Bounds = &BoundingBox{
			Minimum: pm.Point(-1, -1, -1),
			Maximum: pm.Point(1, 1, 1),
		}
	}

	return cube.Bounds
}
