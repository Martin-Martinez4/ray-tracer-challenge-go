package main

import (
	"math"

	"github.com/google/uuid"
)

type Cone struct {
	id         uuid.UUID
	Material   Material
	Transforms Matrix4x4
	SavedRay   Ray
	Minimum    float64
	Maximum    float64
	Closed     bool
	Parent     Shape
}

func NewCone() *Cone {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for cube")
	}
	return &Cone{
		id:         id,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Minimum:    math.Inf(-1),
		Maximum:    math.Inf(1),
		Closed:     false,
		Parent:     nil,
	}
}

func (cone *Cone) GetTransforms() Matrix4x4 {
	return cone.Transforms
}
func (cone *Cone) SetTransform(transform *Matrix4x4) Matrix4x4 {
	cone.Transforms = transform.Multiply(cone.Transforms)
	return cone.Transforms
}
func (cone *Cone) SetTransforms(transform []*Matrix4x4) {
	for i := 0; i < len(transform); i++ {
		cone.SetTransform(transform[i])
	}
}

func (cone *Cone) GetMaterial() *Material {
	return &cone.Material
}
func (cone *Cone) SetMaterial(material Material) {
	cone.Material = material
}

func checkConeCap(ray Ray, t float64, y float64) bool {
	x := ray.origin.x + t*ray.direction.x
	z := ray.origin.z + t*ray.direction.z

	return x*x+z*z <= y*y
}

func intersectConeCaps(cone *Cone, ray Ray, xs *Intersections) {
	if !cone.Closed || AreFloatsEqual(ray.direction.y, 0.0) {
		return
	}

	t := (cone.Minimum - ray.origin.y) / ray.direction.y
	if checkConeCap(ray, t, cone.Minimum) {
		xs.Add(Intersection{t, cone})
	}

	t = (cone.Maximum - ray.origin.y) / ray.direction.y
	if checkConeCap(ray, t, cone.Maximum) {
		xs.Add(Intersection{t, cone})
	}
}

func (cone *Cone) LocalIntersect(ray Ray) Intersections {

	intersections := Intersections{intersections: []Intersection{}}
	a := (ray.direction.x * ray.direction.x) - (ray.direction.y * ray.direction.y) + (ray.direction.z * ray.direction.z)

	b := 2 * ((ray.origin.x * ray.direction.x) - (ray.origin.y * ray.direction.y) + (ray.origin.z * ray.direction.z))

	c := (ray.origin.x * ray.origin.x) - (ray.origin.y * ray.origin.y) + (ray.origin.z * ray.origin.z)
	if AreFloatsEqual(a, 0.0) {
		if AreFloatsEqual(b, 0.0) {
			return intersections
		}

		intersections.Add(Intersection{T: -c / (2 * b), S: cone})
	}
	intersectConeCaps(cone, ray, &intersections)

	disc := ((b * b) - (4 * a * c))

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
	if cone.Minimum < y0 && y0 < cone.Maximum {
		intersections.Add(Intersection{t0, cone})
	}

	y1 := ray.origin.y + t1*ray.direction.y
	if cone.Minimum < y1 && y1 < cone.Maximum {
		intersections.Add(Intersection{t1, cone})
	}

	return intersections

}

func (cone *Cone) Intersect(ray *Ray) Intersections {
	tray := ray.Transform(cone.Transforms.Inverse())
	return cone.LocalIntersect(tray)
}

func (cone *Cone) LocalNormalAt(localPoint Tuple) Tuple {
	dist := (localPoint.x * localPoint.x) + (localPoint.z * localPoint.z)
	if dist < 1 && localPoint.y >= cone.Maximum-Epsilon {
		return Vector(0, 1, 0)
	} else if dist < 1 && localPoint.y <= cone.Minimum+Epsilon {
		return Vector(0, -1, 0)
	} else {
		y := math.Sqrt((localPoint.x * localPoint.x) + (localPoint.z * localPoint.z))
		if localPoint.y > 0 {
			y = -y
		}

		return Vector(localPoint.x, y, localPoint.z)
	}
}

func (cone *Cone) GetSavedRay() Ray {
	return cone.SavedRay
}
func (cone *Cone) SetSavedRay(ray Ray) {
	cone.SavedRay = ray
}

func (cone *Cone) GetParent() Shape {
	return cone.Parent
}

func (cone *Cone) GetId() uuid.UUID {
	return cone.id
}
