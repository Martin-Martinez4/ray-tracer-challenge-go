package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCamera(T *testing.T) {

	theWorld := NewDefaultWorld()

	outer := theWorld.Shapes[0]
	outer.GetMaterial().Ambient = 1
	theWorld.Shapes[0] = outer

	inner := theWorld.Shapes[1]
	inner.GetMaterial().Ambient = 1
	theWorld.Shapes[0] = inner

	tests := []struct {
		name  string
		ray   Ray
		world World
		want  Color
	}{
		{
			name:  "the color with an intersection behind the ray",
			ray:   NewRay([3]float64{0, 0, 0.75}, [3]float64{0, 0, -1}),
			world: theWorld,
			want:  inner.GetMaterial().Color,
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := ColorAt(&tt.ray, &tt.world)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRayForPixel(T *testing.T) {

	type Transformation struct {
		name string
		args []float64
	}

	tests := []struct {
		name   string
		camera Camera
		px     float64
		py     float64
		args   []Transformation
		want   Ray
	}{
		{
			name:   "Constructing a ray through the center of the canvas",
			camera: NewCamera(201, 101, math.Pi/2),
			px:     100,
			py:     50,
			args:   []Transformation{},
			want:   NewRay([3]float64{0, 0, 0}, [3]float64{0, 0, -1}),
		},
		{
			name:   "Constructing a ray through a corner of the canvas",
			camera: NewCamera(201, 101, math.Pi/2),
			px:     0,
			py:     0,
			args:   []Transformation{},
			want:   NewRay([3]float64{0, 0, 0}, [3]float64{0.66519, 0.33259, -0.66851}),
		},
		{
			name:   "Constructing a ray when the camera is transformed",
			camera: NewCamera(201, 101, math.Pi/2),
			px:     100,
			py:     50,
			args:   []Transformation{{name: "rotation_y", args: []float64{math.Pi / 4}}, {name: "translate", args: []float64{0, -2, 5}}},
			want:   NewRay([3]float64{0, 2, -5}, [3]float64{math.Sqrt(2) / 2, 0, -math.Sqrt(2) / 2}),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			transformMatrix := IdentitiyMatrix4x4()

			// transformations applied last first, first last
			for index := len(tt.args) - 1; index >= 0; index-- {

				transform := tt.args[index]

				switch transform.name {
				case "rotation_y":
					if len(transform.args) != 1 {
						t.Errorf("Test %d Failed, transform %d: %s has invalid number of arguments expected: %d, got: %d", i, index, transform.name, 1, len(transform.args))
					} else {
						transformMatrix = transformMatrix.RotationAlongY(transform.args[0])
					}
					break
				case "translate":
					if len(transform.args) != 3 {
						t.Errorf("Test %d Failed, transform %d: %s has invalid number of arguments expected: %d, got: %d", i, index, transform.name, 3, len(transform.args))
					} else {
						transformMatrix = transformMatrix.Translate(transform.args[0], transform.args[1], transform.args[2])
					}
					break
				default:
					t.Errorf("Test %d Failed, transform %d: %s is invalid", i, index, transform.name)
				}
			}

			tt.camera.Transform = tt.camera.Transform.Multiply(transformMatrix)

			got := RayForPixel(tt.camera, tt.px, tt.py)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}
