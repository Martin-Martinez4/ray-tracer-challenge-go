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

	GetInverseTransforms() pm.Matrix4x4
	setInverseTransforms(transforms *pm.Matrix4x4)

	GetMaterial() *mat.Material
	SetMaterial(material mat.Material)

	Intersect(ray *Ray) *Intersections

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
	id               uuid.UUID
	Material         mat.Material
	Transforms       pm.Matrix4x4
	inverseTransform pm.Matrix4x4
	SavedRay         Ray
	Parent           Shape
}

func CreateDefaultPrimeShape() *PrimeShape {
	id, err := uuid.NewUUID()
	identityMatix := pm.NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for cube")
	}

	return &PrimeShape{

		id:               id,
		Transforms:       identityMatix,
		inverseTransform: pm.NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}),
		Material:         mat.DefaultMaterial(),
		Parent:           nil,
	}
}

func (ps *PrimeShape) GetId() uuid.UUID {
	return ps.id
}
func (ps *PrimeShape) GetTransforms() pm.Matrix4x4 {
	return ps.Transforms

}
func (ps *PrimeShape) SetTransform(transform *pm.Matrix4x4) pm.Matrix4x4 {
	ps.Transforms = transform.Multiply(ps.Transforms)
	inverse := ps.Transforms.Inverse()
	ps.setInverseTransforms(&inverse)
	return ps.Transforms
}

func (ps *PrimeShape) SetTransforms(transform []*pm.Matrix4x4) {

	transformation := pm.NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	for i := 0; i < len(transform); i++ {
		// modifying the transformation instead of creating new ones each time might make it faster
		transformation = transform[i].Multiply(transformation)
	}

	ps.SetTransform(&transformation)

}

// Other version I do not know which is better right now they seem about equal
// func (ps *PrimeShape) SetTransforms(transform []*pm.Matrix4x4) {

// 	for i := 0; i < len(transform); i++ {
// 		ps.Transforms = transform[i].Multiply(ps.Transforms)
// 	}
// 	inverse := ps.Transforms.Inverse()
// 	ps.setInverseTransforms(&inverse)

// }

func (ps *PrimeShape) GetInverseTransforms() pm.Matrix4x4 {
	return ps.inverseTransform
}

func (ps *PrimeShape) setInverseTransforms(transforms *pm.Matrix4x4) {
	ps.inverseTransform = *transforms
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

func (ps *PrimeShape) GetParent() Shape {
	return ps.Parent
}

func (ps *PrimeShape) SetParent(shape Shape) {
	ps.Parent = shape
}

func (ps *PrimeShape) PatternAtShape(pattern mat.Pattern, worldPoint pm.Tuple) mat.Color {
	inversObjTransform := ps.GetTransforms().Inverse()

	objectPoint := inversObjTransform.TupleMultiply(worldPoint)

	inversPatTransform := pattern.GetTransforms().Inverse()
	patternPoint := inversPatTransform.TupleMultiply(objectPoint)

	return pattern.PatternAt(patternPoint)
}
