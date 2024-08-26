package main

import (
	"math"

	"github.com/google/uuid"
)

type Triangle struct {
	id uuid.UUID

	P1     Tuple
	P2     Tuple
	P3     Tuple
	Normal Tuple

	E1 Tuple
	E2 Tuple

	Transforms Matrix4x4
	Material   Material
	Parent     Shape
	Bounds     *BoundingBox
	SavedRay   Ray
}

/*
	type Shape interface {
	GetTransforms() Matrix4x4
	SetTransform(transform *Matrix4x4) Matrix4x4
	SetTransforms(transform []*Matrix4x4)

	GetMaterial() *Material
	SetMaterial(material Material)

	Intersect(ray *Ray) Intersections

	NormalAt(point Tuple) Tuple

	GetSavedRay() Ray
	SetSavedRay(ray Ray)
}
*/

func NewTriangle(p1, p2, p3 Tuple) *Triangle {

	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for sphere")
	}
	// e1 = edge 1
	// p1 = point1
	e1 := p2.Subtract(p1)
	e2 := p3.Subtract(p1)

	normal := Normalize(Cross(e2, e1))

	return &Triangle{
		id: id,

		P1: p1,
		P2: p2,
		P3: p3,

		E1: e1,
		E2: e2,

		Normal:     normal,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Parent:     nil,
		Bounds:     nil,
	}
}

func (triangle *Triangle) Equal(other *Triangle) bool {
	return triangle.P1.Equal(other.P1) && triangle.P2.Equal(other.P2) && triangle.P3.Equal(other.P3)
}

func (triangle *Triangle) GetId() uuid.UUID {
	return triangle.id
}

func (triangle *Triangle) GetParent() Shape {
	return triangle.Parent
}
func (triangle *Triangle) SetParent(shape Shape) {
	triangle.Parent = shape
}

func (triangle *Triangle) GetTransforms() Matrix4x4 {
	return triangle.Transforms
}
func (triangle *Triangle) SetTransform(transform *Matrix4x4) Matrix4x4 {
	triangle.Transforms = transform.Multiply(triangle.Transforms)
	return triangle.Transforms
}
func (triangle *Triangle) SetTransforms(transform []*Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		triangle.SetTransform(transform[i])
	}
}

func (triangle *Triangle) GetMaterial() *Material {
	return &triangle.Material
}
func (triangle *Triangle) SetMaterial(material Material) {
	triangle.Material = material
}
func (triangle *Triangle) LocalNormalAt(point Tuple, hitPoint *Tuple, intersection *Intersection) Tuple {
	return triangle.Normal
}
func (triangle *Triangle) NormalAt(worldPoint Tuple) Tuple {
	return triangle.Normal
}

func (triangle *Triangle) LocalIntersect(ray Ray) Intersections {
	dirCrossE2 := Cross(ray.direction, triangle.E2)
	det := Dot(triangle.E1, dirCrossE2)

	if math.Abs(det) < Epsilon {
		return Intersections{intersections: []Intersection{}}
	}

	f := 1.0 / det

	p1ToOrigin := ray.origin.Subtract(triangle.P1)
	u := f * Dot(p1ToOrigin, dirCrossE2)

	if u < 0 || u > 1 {
		return Intersections{intersections: []Intersection{}}
	}

	originCrossE1 := Cross(p1ToOrigin, triangle.E1)
	v := f * Dot(ray.direction, originCrossE1)

	if v < 0 || (u+v) > 1 {
		return Intersections{intersections: []Intersection{}}
	}

	t := f * Dot(triangle.E2, originCrossE1)
	// will change later
	return Intersections{intersections: []Intersection{NewIntersection(t, triangle)}}
}

func (triangle *Triangle) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(triangle.Transforms.Inverse())
	return triangle.LocalIntersect(tray)
}

func (triangle *Triangle) GetSavedRay() Ray {
	return triangle.SavedRay
}
func (triangle *Triangle) SetSavedRay(ray Ray) {
	triangle.SavedRay = ray
}

func (triangle *Triangle) BoundingBox() *BoundingBox {
	if triangle.Bounds == nil {
		triangle.Bounds = &BoundingBox{
			Minimum: Point(-1, -1, -1),
			Maximum: Point(1, 1, 1),
		}
	}

	return triangle.Bounds
}
