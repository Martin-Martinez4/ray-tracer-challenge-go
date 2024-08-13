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
	}
}

func (plane *Plane) GetTransforms() Matrix4x4 {
	return plane.Transforms
}
func (plane *Plane) SetTransform(transform *Matrix4x4) Matrix4x4 {
	return IdentitiyMatrix4x4()
}
func (plane *Plane) SetTransforms(transform []Matrix4x4) {
	return
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

	t := -(ray.origin.y / ray.direction.y)

	return Intersections{intersections: []Intersection{{T: t, S: plane}}}
}

func (plane *Plane) Intersect(ray *Ray) Intersections {
	return plane.LocalIntersect(*ray)
}

func (plane *Plane) LocalNormalAt(localPoint Tuple) Tuple {
	return Vector(0, 1, 0)
}

func (plane *Plane) NormalAt(worldPoint Tuple) Tuple {
	invTransf := plane.GetTransforms().Inverse()
	objectPoint := invTransf.TupleMultiply(worldPoint)

	objectNormal := plane.LocalNormalAt(objectPoint)

	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.TupleMultiply(objectNormal)
	worldNormal.w = 0
	return Normalize(worldNormal)
}

func (plane *Plane) GetSavedRay() Ray {
	return plane.SavedRay
}
func (plane *Plane) SetSavedRay(ray Ray) {
	plane.SavedRay = ray
}
