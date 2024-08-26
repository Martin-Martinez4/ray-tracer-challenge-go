package main

import (
	"math"

	"github.com/google/uuid"
)

type Cylinder struct {
	id         uuid.UUID
	Material   Material
	Transforms Matrix4x4
	SavedRay   Ray
	Minimum    float64
	Maximum    float64
	Closed     bool
	Parent     Shape
	Bounds     *BoundingBox
}

func NewCylinder() *Cylinder {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for cube")
	}
	return &Cylinder{
		id:         id,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Minimum:    math.Inf(-1),
		Maximum:    math.Inf(1),
		Closed:     false,
		Parent:     nil,
		Bounds:     nil,
	}
}

func (cylinder *Cylinder) GetTransforms() Matrix4x4 {
	return cylinder.Transforms
}
func (cylinder *Cylinder) SetTransform(transform *Matrix4x4) Matrix4x4 {
	cylinder.Transforms = transform.Multiply(cylinder.Transforms)
	return cylinder.Transforms
}
func (cylinder *Cylinder) SetTransforms(transform []*Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		cylinder.SetTransform(transform[i])
	}
}

func (cylinder *Cylinder) GetMaterial() *Material {
	return &cylinder.Material
}
func (cylinder *Cylinder) SetMaterial(material Material) {
	cylinder.Material = material
}

func checkCap(ray Ray, t float64) bool {
	x := ray.origin.x + t*ray.direction.x
	z := ray.origin.z + t*ray.direction.z

	return ((x * x) + (z * z)) <= 1
}

func intersectCaps(cylinder *Cylinder, ray Ray, xs *Intersections) {
	if !cylinder.Closed || AreFloatsEqual(ray.direction.y, 0.0) {
		return
	}

	t := (cylinder.Minimum - ray.origin.y) / ray.direction.y
	if checkCap(ray, t) {
		xs.Add(NewIntersection(t, cylinder))
	}

	t = (cylinder.Maximum - ray.origin.y) / ray.direction.y
	if checkCap(ray, t) {
		xs.Add(NewIntersection(t, cylinder))
	}
}

func (cylinder *Cylinder) LocalIntersect(ray Ray) Intersections {

	intersections := Intersections{intersections: []Intersection{}}
	a := (ray.direction.x * ray.direction.x) + (ray.direction.z * ray.direction.z)

	intersectCaps(cylinder, ray, &intersections)
	if AreFloatsEqual(a, 0.0) {
		return intersections
	}

	b := 2 * ((ray.origin.x * ray.direction.x) + (ray.origin.z * ray.direction.z))

	c := (ray.origin.x * ray.origin.x) + (ray.origin.z * ray.origin.z) - 1

	disc := (b * b) - (4 * a * c)

	if disc < 0.0 {
		return intersections
	}

	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)

	if t0 > t1 {
		tempt0 := t0
		t0 = t1
		t1 = tempt0
	}

	y0 := ray.origin.y + t0*ray.direction.y
	if cylinder.Minimum < y0 && y0 < cylinder.Maximum {
		intersections.Add(NewIntersection(t0, cylinder))
	}

	y1 := ray.origin.y + t1*ray.direction.y
	if cylinder.Minimum < y1 && y1 < cylinder.Maximum {
		intersections.Add(NewIntersection(t1, cylinder))
	}

	return intersections

}

func (cylinder *Cylinder) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(cylinder.Transforms.Inverse())
	return cylinder.LocalIntersect(tray)
}

func (cylinder *Cylinder) LocalNormalAt(localPoint Tuple, hitPoint *Tuple, intersection *Intersection) Tuple {
	dist := (localPoint.x * localPoint.x) + (localPoint.z * localPoint.z)
	if dist < 1 && localPoint.y >= cylinder.Maximum-Epsilon {
		return Vector(0, 1, 0)
	} else if dist < 1 && localPoint.y <= cylinder.Minimum+Epsilon {
		return Vector(0, -1, 0)
	} else {

		return Vector(localPoint.x, 0, localPoint.z)
	}
}

func (cylinder *Cylinder) NormalAt(worldPoint Tuple) Tuple {
	return cylinder.LocalNormalAt(worldPoint, nil, nil)
}

func (cylinder *Cylinder) GetSavedRay() Ray {
	return cylinder.SavedRay
}
func (cylinder *Cylinder) SetSavedRay(ray Ray) {
	cylinder.SavedRay = ray
}

func (cylinder *Cylinder) GetParent() Shape {
	return cylinder.Parent
}

func (cylinder *Cylinder) SetParent(shape Shape) {
	cylinder.Parent = shape
}

func (cylinder *Cylinder) GetId() uuid.UUID {
	return cylinder.id
}

func (cylinder *Cylinder) BoundingBox() *BoundingBox {
	if cylinder.Bounds == nil {
		cylinder.Bounds = &BoundingBox{
			Minimum: Point(-1, cylinder.Minimum, -1),
			Maximum: Point(1, cylinder.Maximum, 1),
		}
	}

	return cylinder.Bounds
}
