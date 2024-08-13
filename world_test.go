package main

import (
	"fmt"
	"math"
	"testing"
)

func TestIntersectWorld(T *testing.T) {

	theWorld := NewDefaultWorld()

	tests := []struct {
		name string
		ray  Ray
		want Intersections
	}{
		{
			name: "a ray intersecting a default world should return an intersection struct with four members",
			ray:  NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			want: Intersections{
				intersections: []Intersection{
					{S: theWorld.Shapes[0], T: 4},
					{S: theWorld.Shapes[0], T: 4.5},
					{S: theWorld.Shapes[1], T: 5.5},
					{S: theWorld.Shapes[1], T: 6},
				},
			},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := RayWorldIntersect(tt.ray, theWorld)

			if len(tt.want.intersections) != len(got.intersections) {
				t.Errorf("Lengths did not match for test %d: %s", i, tt.name)
			}

			for k := 0; k < len(got.intersections); k++ {
				if !AreFloatsEqual(tt.want.intersections[k].T, got.intersections[k].T) {
					t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", k, tt.want.intersections, got.intersections)
				}
			}

		})
	}
}

func TestShadeHit(T *testing.T) {

	theWorld := NewDefaultWorld()
	otherLight := (NewLight([3]float64{0, 0.25, 0}, [3]float64{1, 1, 1}))
	theOtherWorld := NewWorld(nil, &otherLight)

	tests := []struct {
		name         string
		ray          Ray
		sphere       Shape
		world        World
		intersection Intersection
		want         Color
	}{
		{
			name:         "shading an intersection",
			ray:          NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:       theWorld.Shapes[0],
			world:        theWorld,
			intersection: Intersection{4, theWorld.Shapes[0]},
			want:         NewColor(0.38066, 0.47583, 0.2855),
		},
		{
			name:         "shading an intersection from the inside",
			ray:          NewRay([3]float64{0, 0, 0}, [3]float64{0, 0, 1}),
			sphere:       theOtherWorld.Shapes[1],
			world:        theOtherWorld,
			intersection: Intersection{0.5, theOtherWorld.Shapes[1]},
			want:         NewColor(0.90498, 0.90498, 0.90498),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := PrepareComputations(tt.ray, tt.sphere, tt.intersection)

			got := ShadeHit(&tt.world, &comps)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestColorAt(T *testing.T) {

	theWorld := NewDefaultWorld()

	tests := []struct {
		name  string
		ray   Ray
		world World
		want  Color
	}{
		{
			name:  "the color when a ray misses should be black (0,0,0)",
			ray:   NewRay([3]float64{0, 0, -5}, [3]float64{0, 1, 0}),
			world: theWorld,
			want:  NewColor(0, 0, 0),
		},
		{
			name:  "the color when a ray hits",
			ray:   NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			world: theWorld,
			want:  NewColor(0.38066, 0.47583, 0.2855),
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

func TestColorAtInner(T *testing.T) {

	theWorld := NewDefaultWorld()

	outer := theWorld.Shapes[0]
	outer.GetMaterial().SetAmbient(1)
	theWorld.Shapes[0] = outer

	inner := theWorld.Shapes[1]
	inner.GetMaterial().SetAmbient(1)
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

func TestIsShadowed(T *testing.T) {

	tests := []struct {
		name  string
		world World
		point Tuple
		want  bool
	}{
		{
			name:  "there is no shadow when nothing is collinear with point and light",
			world: NewDefaultWorld(),
			point: Point(0, 10, 0),
			want:  false,
		},
		{
			name:  "there is no shadow when an object is behind the light",
			world: NewDefaultWorld(),
			point: Point(-20, 20, -20),
			want:  false,
		},
		{
			name:  "the shadow exists when an object is between the point and the light",
			world: NewDefaultWorld(),
			point: Point(10, -10, 10),
			want:  true,
		},
		{
			name:  "there is no shadow when an object is behind the point",
			world: NewDefaultWorld(),
			point: Point(-22, 2, -2),
			want:  false,
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			got := IsShadowed(tt.world, tt.point)

			if got != tt.want {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRender(T *testing.T) {
	tests := []struct {
		name        string
		world       World
		camera      Camera
		from        Tuple
		to          Tuple
		up          Tuple
		transform   Matrix4x4
		pixelCoords [2]float64
		want        Color
	}{
		{
			name:        "rendering a world with a camera",
			world:       NewDefaultWorld(),
			camera:      NewCamera(11, 11, math.Pi/2),
			from:        Point(0, 0, -5),
			to:          Point(0, 0, 0),
			up:          Vector(0, 1, 0),
			pixelCoords: [2]float64{5, 5},
			want:        NewColor(0.38066, 0.47583, 0.2855),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			tt.camera.Transform = ViewTransformation(tt.from, tt.to, tt.up)

			render := Render(tt.camera, tt.world)
			got := render.GetPixel(int32(tt.pixelCoords[0]), int32(tt.pixelCoords[1]))

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestShadeHitWithShadow(T *testing.T) {

	theWorld := NewDefaultWorld()
	theWorld.Light = NewLight([3]float64{0, 0, -10}, [3]float64{1, 1, 1})

	s2 := NewSphere()
	s2.Transforms = s2.Transforms.Translate(0, 0, 10)

	theWorld.Shapes = []Shape{NewSphere(), s2}

	tests := []struct {
		name         string
		ray          Ray
		world        World
		intersection Intersection
		want         Color
	}{
		{
			name:         "shading an intersection",
			ray:          NewRay([3]float64{0, 0, 5}, [3]float64{0, 0, 1}),
			world:        theWorld,
			intersection: Intersection{4, s2},
			want:         NewColor(0.1, 0.1, 0.1),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := PrepareComputations(tt.ray, s2, tt.intersection)

			got := ShadeHit(&tt.world, &comps)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}
