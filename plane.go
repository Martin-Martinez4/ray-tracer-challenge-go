package main

import "github.com/google/uuid"

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

func (plane *Plane) GetTransforms() Matrix4x4 {

}
func (plane *Plane) SetTransform(transform *Matrix4x4) Matrix4x4
func (plane *Plane) SetTransforms(transform []Matrix4x4)

func (plane *Plane) GetMaterial() *Material
func (plane *Plane) SetMaterial(material Material)

func (plane *Plane) Intersect(ray *Ray) Intersections

func (plane *Plane) NormalAt(point Tuple) Tuple

func (plane *Plane) GetSavedRay() Ray
func (plane *Plane) SetSavedRay(ray Ray)
