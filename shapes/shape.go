package shapes

import (
	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/google/uuid"
)

/*
	Shape will have all the functions needed
	The concrete implementations will have the actual struct
*/
/*
	Color:     NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
*/
type Shape interface {
	GetId() uuid.UUID
	GetTransforms() pm.Matrix4x4
	SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4
	SetTransforms(transform []*pm.Matrix4x4)

	GetMaterial() *mat.Material
	SetMaterial(material mat.Material)

	Intersect(ray *Ray) Intersections

	// NormalAt(point pm.Tuple) pm.Tuple
	LocalNormalAt(point pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple

	GetSavedRay() Ray
	SetSavedRay(ray Ray)

	GetParent() Shape
	SetParent(shape Shape)

	BoundingBox() *BoundingBox
	PatternAtShape(pattern mat.Pattern, worldPoint pm.Tuple) mat.Color
}

// short for primordial
type PrimeShape struct {
	id         uuid.UUID
	Material   mat.Material
	Transforms pm.Matrix4x4
	SavedRay   Ray
}

func (ps *PrimeShape) GetId() uuid.UUID {
	return ps.id
}
func (ps *PrimeShape) GetTransforms() pm.Matrix4x4 {
	return ps.Transforms

}
func (ps *PrimeShape) SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4 {
	ps.Transforms = transform.Multiply(ps.Transforms)
	return ps.Transforms
}
func (ps *PrimeShape) SetTransforms(transform []*pm.Matrix4x4) {
	// rework later to when memoizing the inverse transform
	// should keep a running multiplication then at the end memoize instead of calling setTransform for each transform
	for i := 0; i < len(transform); i++ {
		ps.SetTransform(transform[i])
	}
}

func (ps *PrimeShape) GetMaterial() *mat.Material {
	return &ps.Material
}

// copies material to avoid changing the material being passed in
// mat change later
func (ps *PrimeShape) SetMaterial(material mat.Material) {
	ps.Material = material
}

func (ps *PrimeShape) GetSavedRay() Ray {
	return ps.SavedRay
}
func (ps *PrimeShape) SetSavedRay(ray Ray) {
	ps.SavedRay = ray
}

func (ps *PrimeShape) PatternAtShape(pattern mat.Pattern, worldPoint pm.Tuple) mat.Color {
	inversObjTransform := ps.GetTransforms().Inverse()

	objectPoint := inversObjTransform.TupleMultiply(worldPoint)

	inversPatTransform := pattern.GetTransforms().Inverse()
	patternPoint := inversPatTransform.TupleMultiply(objectPoint)

	return pattern.PatternAt(patternPoint)
}

func CreateDefaultPrimeShape() *PrimeShape {
	id, err := uuid.NewUUID()
	identityMatix := pm.NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for cube")
	}

	return &PrimeShape{

		id:         id,
		Transforms: identityMatix,
		Material:   mat.DefaultMaterial(),
	}
}
