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

type Plane struct {
	id         uuid.UUID
	Material   Material
	Transforms Matrix4x4
	SavedRay   Ray
	Parent     Shape
	Bounds     *BoundingBox
}

func NewPlane() *Plane {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for plane")
	}
	return &Plane{
		id:         id,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Parent:     nil,
		Bounds:     nil,
	}
}

func (plane *Plane) GetTransforms() Matrix4x4 {
	return plane.Transforms
}
func (plane *Plane) SetTransform(transform *Matrix4x4) Matrix4x4 {
	plane.Transforms = transform.Multiply(plane.Transforms)
	return plane.Transforms
}
func (plane *Plane) SetTransforms(transform []*Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		plane.SetTransform(transform[i])
	}
}

func (plane *Plane) GetMaterial() *Material {
	return &plane.Material
}
func (plane *Plane) SetMaterial(material Material) {
	plane.Material = material
}
func (plane *Plane) LocalIntersect(ray Ray) Intersections {
	if math.Abs(ray.direction.y) < Epsilon {
		return Intersections{intersections: []Intersection{}}
	}

	t := (-ray.origin.y / ray.direction.y)

	return Intersections{intersections: []Intersection{{T: t, S: plane}}}
}

func (plane *Plane) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(plane.Transforms.Inverse())
	return plane.LocalIntersect(tray)
}

func (plane *Plane) LocalNormalAt(localPoint Tuple, hitPoint *Tuple, intersection *Intersection) Tuple {
	return Vector(0, 1, 0)
}

func (plane *Plane) NormalAt(worldPoint Tuple) Tuple {
	return plane.LocalNormalAt(worldPoint, nil, nil)
}

func (plane *Plane) GetSavedRay() Ray {
	return plane.SavedRay
}
func (plane *Plane) SetSavedRay(ray Ray) {
	plane.SavedRay = ray
}

func (plane *Plane) GetParent() Shape {
	return plane.Parent
}

func (plane *Plane) SetParent(shape Shape) {
	plane.Parent = shape
}

func (plane *Plane) GetId() uuid.UUID {
	return plane.id
}

func (plane *Plane) BoundingBox() *BoundingBox {
	if plane.Bounds == nil {
		plane.Bounds = &BoundingBox{
			Minimum: Point(math.Inf(-1), 0, math.Inf(-1)),
			Maximum: Point(math.Inf(1), 0, math.Inf(1)),
		}
	}

	return plane.Bounds
}
