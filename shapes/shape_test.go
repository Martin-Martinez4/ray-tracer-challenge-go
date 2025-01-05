package shapes

import (
	"math"
	"testing"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
)

type TestShape struct {
	*PrimeShape
}

func (shape *TestShape) GetTransform() pm.Matrix4x4 {
	return shape.Transforms
}

func NewTestShape() *TestShape {

	return &TestShape{
		PrimeShape: CreateDefaultPrimeShape(),
	}
}

func (shape *TestShape) LocalIntersect(ray Ray) *Intersections {

	inters := Intersections{}

	sphereToRay := ray.origin.Subtract(pm.Point(0, 0, 0))

	a := pm.Dot(ray.Direction, ray.Direction)
	b := 2 * pm.Dot(ray.Direction, sphereToRay)
	c := pm.Dot(sphereToRay, sphereToRay) - 1

	discriminant := (b * b) - (4*a)*c

	if discriminant < 0 {

	} else {
		d1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		d2 := (-b + math.Sqrt(discriminant)) / (2 * a)

		if !pm.AreFloatsEqual(d1, d2) {

			inters.Add(NewIntersection(d1, shape))
			inters.Add(NewIntersection(d2, shape))

		} else {

			inters.Add(NewIntersection(d1, shape))

		}
	}

	return &inters
}

func (shape *TestShape) Intersect(ray *Ray) *Intersections {

	shape.SetSavedRay(ray.Transform(shape.Transforms.Inverse()))
	return shape.LocalIntersect(shape.GetSavedRay())

}

func (shape *TestShape) LocalNormalAt(localPoint pm.Tuple, hitPoint *pm.Tuple, intersection *Intersection) pm.Tuple {
	return pm.Vector(localPoint.X, localPoint.Y, localPoint.Z)
}

func (shape *TestShape) NormalAt(worldPoint pm.Tuple) pm.Tuple {
	invTransf := shape.GetTransforms().Inverse()
	objectPoint := invTransf.TupleMultiply(worldPoint)

	objectNormal := shape.LocalNormalAt(objectPoint, nil, nil)

	invTransfTransposed := invTransf.Transpose()
	worldNormal := invTransfTransposed.TupleMultiply(objectNormal)
	worldNormal.W = 0
	return pm.Normalize(worldNormal)
}

func (shape *TestShape) BoundingBox() *BoundingBox {
	return nil
}

func TestShapeGetTransform(t *testing.T) {
	translate := pm.Translate(2, 3, 4)
	tests := []struct {
		name       string
		shape      Shape
		transforms []*pm.Matrix4x4
		want       pm.Matrix4x4
	}{
		{
			name:       "the default transformation",
			shape:      NewTestShape(),
			transforms: []*pm.Matrix4x4{},
			want:       pm.IdentitiyMatrix4x4(),
		},
		{
			name:       "assiging a translate",
			shape:      NewTestShape(),
			transforms: []*pm.Matrix4x4{pm.Translate(2, 3, 4)},
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
		material mat.Material
		want     mat.Material
	}{
		{
			name:     "the default material",
			shape:    NewTestShape(),
			material: mat.DefaultMaterial(),
			want:     mat.DefaultMaterial(),
		},
		{
			name:     "assigning a material",
			shape:    NewTestShape(),
			material: mat.Material{Color: mat.NewColor(1, 1, 1), Ambient: 0.1, Diffuse: 0.9, Specular: 0.9, Shininess: 200.0},
			want:     mat.Material{Color: mat.NewColor(1, 1, 1), Ambient: 0.1, Diffuse: 0.9, Specular: 0.9, Shininess: 200.0},
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
		transforms []*pm.Matrix4x4
		want       Ray
	}{
		{
			name:       "intersecting a scaled shape with a ray",
			ray:        NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			shape:      NewTestShape(),
			transforms: []*pm.Matrix4x4{pm.Scale(2, 2, 2)},
			want:       NewRay([3]float64{0, 0, -2.5}, [3]float64{0, 0, 0.5}),
		},
		{
			name:       "intersecting a translated shape with a ray",
			ray:        NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			shape:      NewTestShape(),
			transforms: []*pm.Matrix4x4{pm.Translate(5, 0, 0)},
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
		point      pm.Tuple
		transforms []*pm.Matrix4x4
		want       pm.Tuple
	}{
		{
			name:       "computing the normal on a translated shape",
			shape:      NewTestShape(),
			point:      pm.Point(0, 1.70711, -0.70711),
			transforms: []*pm.Matrix4x4{pm.Translate(0, 1, 0)},
			want:       pm.Vector(0, 0.70711, -0.70711),
		},
		{
			name:       "computing the normal on a transformed shape",
			shape:      NewTestShape(),
			point:      pm.Point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			transforms: []*pm.Matrix4x4{pm.RotationAlongZ(math.Pi / 5), pm.Scale(1, 0.5, 1)},
			want:       pm.Vector(0, 0.97014, -0.24254),
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
