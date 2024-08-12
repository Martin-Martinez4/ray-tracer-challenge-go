package main

/*
returns a Color
	material default material
	position Point(0,0,0)

	* Will Change *
	eyev
	normalv
	light
*/

import (
	"math"
	"testing"
)

func TestEffectiveLighting(t *testing.T) {

	defaultMat := DefaultMaterial()
	point := Point(0, 0, 0)

	tests := []struct {
		name      string
		material  Material
		eyeVec    Tuple
		normalVec Tuple
		point     Tuple
		light     Light
		want      Color
	}{
		{

			name:      "lighting with the eye between the light and the surface",
			material:  defaultMat,
			eyeVec:    Vector(0, 0, -1),
			normalVec: Vector(0, 0, -1),
			point:     point,
			light:     NewLight([3]float64{0, 0, -10}, [3]float64{1, 1, 1}),
			want:      NewColor(1.9, 1.9, 1.9),
		},
		{
			name:      "lighting with the eye between the light and the surface, eye offset 45deg",
			material:  defaultMat,
			eyeVec:    Vector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			normalVec: Vector(0, 0, -1),
			point:     point,
			light:     NewLight([3]float64{0, 0, -10}, [3]float64{1, 1, 1}),
			want:      NewColor(1.0, 1.0, 1.0),
		},
		{
			name:      "lighting with the eye opposite surface, eye offset 45deg",
			material:  defaultMat,
			eyeVec:    Vector(0, 0, -1),
			normalVec: Vector(0, 0, -1),
			point:     point,
			light:     NewLight([3]float64{0, 10, -10}, [3]float64{1, 1, 1}),
			want:      NewColor(0.7364, 0.7364, 0.7364),
		},
		{
			name:      "lighting with the eye in the path of the reflection vector",
			material:  defaultMat,
			eyeVec:    Vector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2),
			normalVec: Vector(0, 0, -1),
			point:     point,
			light:     NewLight([3]float64{0, 10, -10}, [3]float64{1, 1, 1}),
			want:      NewColor(1.6364, 1.6364, 1.6364),
		},
		{
			name:      "lighting with the light behind the surface",
			material:  defaultMat,
			eyeVec:    Vector(0, 0, -1),
			normalVec: Vector(0, 0, -1),
			point:     point,
			light:     NewLight([3]float64{0, 0, 10}, [3]float64{1, 1, 1}),
			want:      NewColor(0.1, 0.1, 0.1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := EffectiveLighting(tt.material, tt.light, tt.point, tt.eyeVec, tt.normalVec, false)

			if !got.Equal(tt.want) {
				t.Errorf("%s failed\nWanted:\n%v\nGot:\n%v\n", tt.name, tt.want, got)
			}

		})
	}
}

func TestLightingWithShadows(t *testing.T) {

	defaultMat := DefaultMaterial()
	point := Point(0, 0, 0)

	tests := []struct {
		name      string
		material  Material
		eyeVec    Tuple
		normalVec Tuple
		light     Light
		point     Tuple
		inShadow  bool
		want      Color
	}{
		{

			name:      "lighting with the surface in shadow",
			material:  defaultMat,
			eyeVec:    Vector(0, 0, -1),
			normalVec: Vector(0, 0, -1),
			light:     NewLight([3]float64{0, 0, -10}, [3]float64{1, 1, 1}),
			inShadow:  true,
			point:     point,
			want:      NewColor(0.1, 0.1, 0.1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := EffectiveLighting(tt.material, tt.light, tt.point, tt.eyeVec, tt.normalVec, tt.inShadow)

			if !got.Equal(tt.want) {
				t.Errorf("%s failed\nWanted:\n%v\nGot:\n%v\n", tt.name, tt.want, got)
			}

		})
	}
}
