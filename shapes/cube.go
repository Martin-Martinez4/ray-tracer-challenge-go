package shapes

import (
	"math"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/google/uuid"
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
	Parent Shape
	Bounds *BoundingBox
}

func NewCube() *Cube {

	return &Cube{
		PrimeShape: CreateDefaultPrimeShape(),
		Parent:     nil,
	}
}

func (cube *Cube) GetTransforms() pm.Matrix4x4 {
	return cube.Transforms
}
func (cube *Cube) SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4 {
	cube.Transforms = transform.Multiply(cube.Transforms)
	return cube.Transforms
}
func (cube *Cube) SetTransforms(transform []*pm.Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		cube.SetTransform(transform[i])
	}
}

func (cube *Cube) GetMaterial() *mat.Material {
	return &cube.Material
}
func (cube *Cube) SetMaterial(material mat.Material) {
	cube.Material = material
}

func CheckAxis(origin, direction float64) (float64, float64) {
	tminNumerator := (-1 - origin)
	tmaxNumerator := (1 - origin)

	var tmin float64
	var tmax float64

	if math.Abs(direction) >= pm.Epsilon {
		tmin = tminNumerator / direction
		tmax = tmaxNumerator / direction
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

func (cube *Cube) LocalIntersect(ray Ray) Intersections {
	xtmin, xtmax := CheckAxis(ray.origin.X, ray.direction.X)
	ytmin, ytmax := CheckAxis(ray.origin.Y, ray.direction.Y)
	ztmin, ztmax := CheckAxis(ray.origin.Z, ray.direction.Z)

	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	if tmin > tmax {
		return Intersections{intersections: []Intersection{}}

	}

	return Intersections{intersections: []Intersection{NewIntersection(tmin, cube), NewIntersection(tmax, cube)}}
}

func (cube *Cube) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(cube.Transforms.Inverse())
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

func (cube *Cube) GetSavedRay() Ray {
	return cube.SavedRay
}
func (cube *Cube) SetSavedRay(ray Ray) {
	cube.SavedRay = ray
}

func (cube *Cube) GetParent() Shape {
	return cube.Parent
}

func (cube *Cube) SetParent(shape Shape) {
	cube.Parent = shape
}

func (cube *Cube) GetId() uuid.UUID {
	return cube.id
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
