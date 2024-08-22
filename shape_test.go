package main

import (
	"math"
	"testing"

	"github.com/google/uuid"
)

type TestShape struct {
	id         uuid.UUID
	Material   Material
	Transforms Matrix4x4
	SavedRay   Ray
	Parent     Shape
}

func (shape *TestShape) GetTransform() Matrix4x4 {
	return shape.Transforms
}

func NewTestShape() *TestShape {
	id, err := uuid.NewUUID()
	identityMatix := NewMatrix4x4([16]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})

	if err != nil {
		panic("not able to cerate unique id for shape")
	}
	return &TestShape{
		id:         id,
		Transforms: identityMatix,
		Material:   DefaultMaterial(),
		Parent:     nil,
	}
}

func (shape *TestShape) GetTransforms() Matrix4x4 {
	return shape.Transforms
}

func (shape *TestShape) SetTransform(mat44 *Matrix4x4) Matrix4x4 {
	shape.Transforms = mat44.Multiply(shape.Transforms)
	return shape.Transforms
}

func (shape *TestShape) SetTransforms(mat44 []*Matrix4x4) {

	for _, transform := range mat44 {

		shape.SetTransform(transform)
	}

}

func (shape *TestShape) GetMaterial() *Material {

	return &shape.Material
}

func (shape *TestShape) SetMaterial(material Material) {

	shape.Material = material
}

func (shape *TestShape) GetSavedRay() Ray {
	return shape.SavedRay
}
func (shape *TestShape) SetSavedRay(ray Ray) {
	shape.SavedRay = ray
}

func (shape *TestShape) LocalIntersect(ray Ray) Intersections {

	inters := Intersections{}

	sphereToRay := ray.origin.Subtract(Point(0, 0, 0))

	a := Dot(ray.direction, ray.direction)
	b := 2 * Dot(ray.direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1

	discriminant := (b * b) - (4*a)*c

	if discriminant < 0 {

	} else {
		d1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		d2 := (-b + math.Sqrt(discriminant)) / (2 * a)

		if !AreFloatsEqual(d1, d2) {

			inters.Add(Intersection{d1, shape})
			inters.Add(Intersection{d2, shape})

		} else {

			inters.Add(Intersection{d1, shape})

		}
	}

	return inters
}

func (shape *TestShape) Intersect(ray *Ray) Intersections {

	shape.SetSavedRay(ray.Transform(shape.Transforms.Inverse()))
	return shape.LocalIntersect(shape.GetSavedRay())

}

func (shape *TestShape) LocalNormalAt(localPoint Tuple) Tuple {
	return Vector(localPoint.x, localPoint.y, localPoint.z)
}

func (shape *TestShape) NormalAt(worldPoint Tuple) Tuple {
	invTransf := shape.GetTransforms().Inverse()
	objectPoint := invTransf.TupleMultiply(worldPoint)

	objectNormal := shape.LocalNormalAt(objectPoint)

	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.TupleMultiply(objectNormal)
	worldNormal.w = 0
	return Normalize(worldNormal)
}

func (shape *TestShape) GetParent() Shape {
	return shape.Parent
}

func (shape *TestShape) GetId() uuid.UUID {
	return shape.id
}

func TestShapeGetTransform(t *testing.T) {
	translate := Translate(2, 3, 4)
	tests := []struct {
		name       string
		shape      Shape
		transforms []*Matrix4x4
		want       Matrix4x4
	}{
		{
			name:       "the default transformation",
			shape:      NewTestShape(),
			transforms: []*Matrix4x4{},
			want:       IdentitiyMatrix4x4(),
		},
		{
			name:       "assiging a translate",
			shape:      NewTestShape(),
			transforms: []*Matrix4x4{Translate(2, 3, 4)},
			want:       *translate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for i := 0; i < len(tt.transforms); i++ {
				tt.shape.SetTransform(tt.transforms[i])
			}

			got := tt.shape.GetTransforms()

			if !got.Equal(tt.want) {
				t.Errorf("%s did not pass: \nGot: %s \nWanted: %s", tt.name, got.Print(), tt.want.Print())
			}

		})
	}
}

func TestShapeGetMaterial(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		material Material
		want     Material
	}{
		{
			name:     "the default material",
			shape:    NewTestShape(),
			material: DefaultMaterial(),
			want:     DefaultMaterial(),
		},
		{
			name:     "assigning a material",
			shape:    NewTestShape(),
			material: Material{Color: NewColor(1, 1, 1), Ambient: 0.1, Diffuse: 0.9, Specular: 0.9, Shininess: 200.0},
			want:     Material{Color: NewColor(1, 1, 1), Ambient: 0.1, Diffuse: 0.9, Specular: 0.9, Shininess: 200.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.shape.SetMaterial(tt.material)

			got := tt.shape.GetMaterial()

			if !got.Equal(tt.want) {
				t.Errorf("%s did not pass: \nGot: %s \nWanted: %s", tt.name, got.Print(), tt.want.Print())
			}

		})
	}
}

func TestShapeSavedRay(t *testing.T) {
	tests := []struct {
		name       string
		ray        Ray
		shape      Shape
		transforms []*Matrix4x4
		want       Ray
	}{
		{
			name:       "intersecting a scaled shape with a ray",
			ray:        NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			shape:      NewTestShape(),
			transforms: []*Matrix4x4{Scale(2, 2, 2)},
			want:       NewRay([3]float64{0, 0, -2.5}, [3]float64{0, 0, 0.5}),
		},
		{
			name:       "intersecting a translated shape with a ray",
			ray:        NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			shape:      NewTestShape(),
			transforms: []*Matrix4x4{Translate(5, 0, 0)},
			want:       NewRay([3]float64{-5, 0, -5}, [3]float64{0, 0, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.shape.Intersect(&tt.ray)

			got := tt.shape.GetSavedRay()

			if got.Equal(tt.want) {
				t.Errorf("%s did not pass: \nGot: %s \nWanted: %s", tt.name, got.Print(), tt.want.Print())
			}

		})
	}
}

func TestShapeNormalAt(t *testing.T) {
	tests := []struct {
		name       string
		shape      Shape
		point      Tuple
		transforms []*Matrix4x4
		want       Tuple
	}{
		{
			name:       "computing the normal on a translated shape",
			shape:      NewTestShape(),
			point:      Point(0, 1.70711, -0.70711),
			transforms: []*Matrix4x4{Translate(0, 1, 0)},
			want:       Vector(0, 0.70711, -0.70711),
		},
		{
			name:       "computing the normal on a transformed shape",
			shape:      NewTestShape(),
			point:      Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			transforms: []*Matrix4x4{RotationAlongZ(math.Pi / 5), Scale(1, 0.5, 1)},
			want:       Vector(0, 0.97014, -0.24254),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.shape.SetTransforms(tt.transforms)

			got := NormalAt(tt.shape, tt.point)

			if !got.Equal(tt.want) {
				t.Errorf("%s did not pass: \nGot: %s \nWanted: %s", tt.name, got.Print(), tt.want.Print())
			}

		})
	}
}
