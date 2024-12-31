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

type Plane struct {
	*PrimeShape
	SavedRay Ray
	Parent   Shape
	Bounds   *BoundingBox
}

func NewPlane() *Plane {

	return &Plane{
		PrimeShape: CreateDefaultPrimeShape(),
		Parent:     nil,
		Bounds:     nil,
	}
}

func (plane *Plane) GetTransforms() pm.Matrix4x4 {
	return plane.Transforms
}
func (plane *Plane) SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4 {
	plane.Transforms = transform.Multiply(plane.Transforms)
	return plane.Transforms
}
func (plane *Plane) SetTransforms(transform []*pm.Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		plane.SetTransform(transform[i])
	}
}

func (plane *Plane) GetMaterial() *mat.Material {
	return &plane.Material
}
func (plane *Plane) SetMaterial(material mat.Material) {
	plane.Material = material
}
func (plane *Plane) LocalIntersect(ray Ray) Intersections {
	if math.Abs(ray.Direction.Y) < pm.Epsilon {
		return Intersections{Intersections: []Intersection{}}
	}

	t := (-ray.origin.Y / ray.Direction.Y)

	return Intersections{Intersections: []Intersection{{T: t, S: plane}}}
}

func (plane *Plane) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(plane.Transforms.Inverse())
	return plane.LocalIntersect(tray)
}

func (plane *Plane) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	return pm.Vector(0, 1, 0)
}

func (plane *Plane) NormalAt(worldPoint pm.Tuple) pm.Tuple {
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
			Minimum: pm.Point(math.Inf(-1), 0, math.Inf(-1)),
			Maximum: pm.Point(math.Inf(1), 0, math.Inf(1)),
		}
	}

	return plane.Bounds
}
