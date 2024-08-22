package main

import (
	"sort"

	"github.com/google/uuid"
)

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

type ChildrenMap map[uuid.UUID]Shape

type Group struct {
	id         uuid.UUID
	Children   ChildrenMap
	Transforms Matrix4x4
	Material   Material
	Parent     Shape
	SavedRay   Ray
}

func NewGroup() *Group {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for cube")
	}
	return &Group{
		id:         id,
		Children:   ChildrenMap{},
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Parent:     nil,
	}
}

func (group *Group) GetId() uuid.UUID {
	return group.id
}
func (group *Group) GetParent() Shape {
	return group.Parent
}

func (group *Group) GetSavedRay() Ray {
	return group.SavedRay
}
func (group *Group) SetSavedRay(ray Ray) {
	group.SavedRay = ray
}

func (group *Group) AddChild(shape Shape) {
	group.Children[shape.GetId()] = shape
}

func (group *Group) GetTransforms() Matrix4x4 {
	return group.Transforms
}
func (group *Group) SetTransform(transform *Matrix4x4) Matrix4x4 {
	group.Transforms = transform.Multiply(group.Transforms)
	return group.Transforms
}
func (group *Group) SetTransforms(transform []*Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		group.SetTransform(transform[i])
	}
}

func (group *Group) GetMaterial() *Material {
	return &group.Material
}
func (group *Group) SetMaterial(material Material) {
	group.Material = material
}

// Could really be improved
func (group *Group) LocalIntersect(ray Ray) Intersections {
	// Def need to change this
	overAllIntersections := Intersections{intersections: []Intersection{}}
	for _, shape := range group.Children {
		inters := shape.Intersect(&ray)
		overAllIntersections.intersections = append(overAllIntersections.intersections, inters.intersections...)
	}
	sort.Slice(overAllIntersections.intersections, func(i, j int) bool {
		return overAllIntersections.intersections[i].T < overAllIntersections.intersections[j].T
	})

	return overAllIntersections
}

func (group *Group) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(group.Transforms.Inverse())
	return group.LocalIntersect(tray)
}

func WorldToObject(shape Shape, point Tuple) Tuple {
	if shape.GetParent() != nil {
		point = WorldToObject(shape.GetParent(), point)
	}

	return shape.GetTransforms().Inverse().TupleMultiply(point)
}

func NormalToWorld(shape Shape, normal Tuple) Tuple {
	shapTransform := shape.GetTransforms().Inverse().Transpose()
	normal = shapTransform.TupleMultiply(normal)
	normal.w = 0
	normal = Normalize(normal)

	if shape.GetParent() != nil {
		normal = NormalToWorld(shape.GetParent(), normal)
	}

	return normal
}

func (group *Group) LocalNormalAt(localPoint Tuple) Tuple {
	panic("local normal called on group")
	return Vector(0, 1, 0)
}

func NormalAt(shape Shape, point Tuple) Tuple {
	localPoint := WorldToObject(shape, point)
	localNormal := shape.LocalNormalAt(localPoint)
	return NormalToWorld(shape, localNormal)
}
