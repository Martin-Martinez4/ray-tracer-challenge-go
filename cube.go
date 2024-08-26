package main

import (
	"math"

	"github.com/google/uuid"
)

// type Shape interface {
// 	GetTransforms() Matrix4x4
// 	SetTransform(transform *Matrix4x4) Matrix4x4
// 	SetTransforms(transform []Matrix4x4)

// 	GetMaterial() *Material
// 	SetMaterial(material Material)

// 	Intersect(ray *Ray) Intersections

// 	NormalAt(point Tuple) Tuple

// 	GetSavedRay() Ray
// 	SetSavedRay(ray Ray)
// }

type Cube struct {
	id         uuid.UUID
	Material   Material
	Transforms Matrix4x4
	SavedRay   Ray
	Parent     Shape
	Bounds     *BoundingBox
}

func NewCube() *Cube {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for cube")
	}
	return &Cube{
		id:         id,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Parent:     nil,
	}
}

func (cube *Cube) GetTransforms() Matrix4x4 {
	return cube.Transforms
}
func (cube *Cube) SetTransform(transform *Matrix4x4) Matrix4x4 {
	cube.Transforms = transform.Multiply(cube.Transforms)
	return cube.Transforms
}
func (cube *Cube) SetTransforms(transform []*Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		cube.SetTransform(transform[i])
	}
}

func (cube *Cube) GetMaterial() *Material {
	return &cube.Material
}
func (cube *Cube) SetMaterial(material Material) {
	cube.Material = material
}

func CheckAxis(origin, direction float64) (float64, float64) {
	tminNumerator := (-1 - origin)
	tmaxNumerator := (1 - origin)

	var tmin float64
	var tmax float64

	if math.Abs(direction) >= Epsilon {
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
	xtmin, xtmax := CheckAxis(ray.origin.x, ray.direction.x)
	ytmin, ytmax := CheckAxis(ray.origin.y, ray.direction.y)
	ztmin, ztmax := CheckAxis(ray.origin.z, ray.direction.z)

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

func (cube *Cube) LocalNormalAt(localPoint Tuple, hitPoint *Tuple, intersection *Intersection) Tuple {

	absX := math.Abs(localPoint.x)
	absY := math.Abs(localPoint.y)
	absZ := math.Abs(localPoint.z)

	maxc := math.Max(math.Max(absX, absY), absZ)

	if maxc == absX {
		return Vector(localPoint.x, 0, 0)
	} else if maxc == absY {
		return Vector(0, localPoint.y, 0)
	} else {
		return Vector(0, 0, localPoint.z)
	}
}

func (cube *Cube) NormalAt(worldPoint Tuple) Tuple {
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
			Minimum: Point(-1, -1, -1),
			Maximum: Point(1, 1, 1),
		}
	}

	return cube.Bounds
}
