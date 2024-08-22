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

			got := ShadeHit(tt.world, comps, 1)

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

			got := ColorAt(tt.ray, tt.world, 1)

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

			got := ColorAt(tt.ray, tt.world, 1)

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

			got := ShadeHit(tt.world, comps, 1)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRefractedColorOpaque(T *testing.T) {

	theWorld := NewDefaultWorld()
	shape := theWorld.Shapes[0]

	tests := []struct {
		name         string
		ray          Ray
		world        World
		intersection []Intersection
		want         Color
	}{
		{
			name:         "shading an intersection",
			world:        theWorld,
			ray:          NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			intersection: []Intersection{{4, shape}, {6, shape}},
			want:         BLACK,
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := PrepareComputationsWithHit(tt.intersection[0], tt.ray, tt.intersection)

			got := RefreactedColor(tt.world, *comps, 5)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRefractedColorMax0(T *testing.T) {

	theWorld := NewDefaultWorld()
	shape := theWorld.Shapes[0]
	shape.GetMaterial().Transparency = 1.0
	shape.GetMaterial().RefractiveIndex = 1.5

	tests := []struct {
		name         string
		ray          Ray
		world        World
		intersection []Intersection
		want         Color
	}{
		{
			name:         "shading an intersection",
			world:        theWorld,
			ray:          NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			intersection: []Intersection{{4, shape}, {6, shape}},
			want:         BLACK,
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := PrepareComputationsWithHit(tt.intersection[0], tt.ray, tt.intersection)

			got := RefreactedColor(tt.world, *comps, 0)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRefractedColorSnellLaw(T *testing.T) {

	theWorld := NewDefaultWorld()
	shape := theWorld.Shapes[0]
	shape.GetMaterial().Transparency = 1.0
	shape.GetMaterial().RefractiveIndex = 1.5

	tests := []struct {
		name         string
		ray          Ray
		world        World
		intersection []Intersection
		want         Color
	}{
		{
			name:         "the refracted color under total internal relfection",
			world:        theWorld,
			ray:          NewRay([3]float64{0, 0, math.Sqrt(2) / 2}, [3]float64{0, 1, 0}),
			intersection: []Intersection{{-math.Sqrt(2) / 2, shape}, {math.Sqrt(2) / 2, shape}},
			want:         BLACK,
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := PrepareComputationsWithHit(tt.intersection[1], tt.ray, tt.intersection)

			got := RefreactedColor(tt.world, *comps, 5)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRefractedColorRefractedRay(T *testing.T) {

	theWorld := NewDefaultWorld()

	shapeA := theWorld.Shapes[0]
	shapeA.GetMaterial().Ambient = 1.0
	shapeA.GetMaterial().Pattern = NewTestPattern(BLACK, WHITE)

	shapeB := theWorld.Shapes[1]
	shapeB.GetMaterial().Transparency = 1.0
	shapeB.GetMaterial().RefractiveIndex = 1.5

	tests := []struct {
		name         string
		ray          Ray
		world        World
		intersection []Intersection
		want         Color
	}{
		{
			name:         "the refracted color under total internal relfection",
			world:        theWorld,
			ray:          NewRay([3]float64{0, 0, 0.1}, [3]float64{0, 1, 0}),
			intersection: []Intersection{{-0.9899, shapeA}, {-0.4899, shapeB}, {0.4899, shapeB}, {0.9899, shapeA}},
			want:         NewColor(0, 0.99888, 0.04725),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			comps := PrepareComputationsWithHit(tt.intersection[2], tt.ray, tt.intersection)

			got := RefreactedColor(tt.world, *comps, 5)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRefractorShadeHit(T *testing.T) {

	theWorld := NewDefaultWorld()

	floor := NewPlane()
	floor.GetMaterial().Transparency = 0.5
	floor.GetMaterial().RefractiveIndex = 1.5
	floor.SetTransform(Translate(0, -1, 0))

	ball := NewSphere()
	ball.GetMaterial().Color = NewColor(1, 0, 0)
	ball.GetMaterial().Ambient = 0.5
	ball.SetTransform(Translate(0, -3.5, -0.5))

	theWorld.Shapes = append(theWorld.Shapes, floor, ball)

	tests := []struct {
		name         string
		ray          Ray
		world        World
		intersection []Intersection
		want         Color
	}{
		{
			name:         "ShadeHit with a transparent material",
			world:        theWorld,
			ray:          NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: []Intersection{{math.Sqrt(2), floor}},
			want:         NewColor(0.93642, 0.68642, 0.68642),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			comps := PrepareComputationsWithHit(tt.intersection[0], tt.ray, tt.intersection)

			got := ShadeHit(theWorld, *comps, 5)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestSchlickInternalReflection(T *testing.T) {

	shape := NewGlassSphere()

	name := "ShadeHit with a transparent material"
	ray := NewRay([3]float64{0, 0, math.Sqrt(2) / 2}, [3]float64{0, 1, 0})
	intersection := []Intersection{{-math.Sqrt(2) / 2, shape}, {math.Sqrt(2) / 2, shape}}
	want := 1.0

	T.Run(name, func(t *testing.T) {
		comps := PrepareComputationsWithHit(intersection[1], ray, intersection)

		got := Schlick(comps)

		if !AreFloatsEqual(got, want) {
			t.Errorf("\nwant: %v \ngot: %v \ndo not match", want, got)
		}

	})
}

func TestSchlickPerpendicular(T *testing.T) {

	shape := NewGlassSphere()

	name := "schilick with perpendicular viewing angle"
	ray := NewRay([3]float64{0, 0, 0}, [3]float64{0, 1, 0})
	intersection := []Intersection{{-1, shape}, {1, shape}}
	want := 0.04

	T.Run(name, func(t *testing.T) {
		comps := PrepareComputationsWithHit(intersection[1], ray, intersection)

		got := Schlick(comps)

		if !AreFloatsEqual(got, want) {
			t.Errorf("\nwant: %v \ngot: %v \ndo not match", want, got)
		}

	})
}

func TestSchlickSmallAngle(T *testing.T) {

	shape := NewGlassSphere()

	name := "schilick with small viewing angle, where n2 > n1"
	ray := NewRay([3]float64{0, 0.99, -2}, [3]float64{0, 0, 1})
	intersection := []Intersection{{1.8589, shape}}
	want := 0.48873

	T.Run(name, func(t *testing.T) {
		comps := PrepareComputationsWithHit(intersection[0], ray, intersection)

		got := Schlick(comps)

		if !AreFloatsEqual(got, want) {
			t.Errorf("\nwant: %v \ngot: %v \ndo not match", want, got)
		}

	})
}

func TestSchlickReflectionAndRefraction(T *testing.T) {

	name := "schilick with reflection and refraction"
	ray := NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2})
	world := NewDefaultWorld()

	floor := NewPlane()
	floor.SetTransform(Translate(0, -1, 0))
	floor.GetMaterial().Reflective = 0.5
	floor.GetMaterial().Transparency = 0.5
	floor.GetMaterial().RefractiveIndex = 1.5

	ball := NewSphere()
	ball.GetMaterial().Color = NewColor(1, 0, 0)
	ball.GetMaterial().Ambient = 0.5
	ball.SetTransform(Translate(0, -3.5, -0.5))

	world.Shapes = append(world.Shapes, floor, ball)

	intersection := []Intersection{{math.Sqrt(2), floor}}
	want := NewColor(0.93391, 0.69643, 0.69243)

	T.Run(name, func(t *testing.T) {
		comps := PrepareComputationsWithHit(intersection[0], ray, intersection)

		got := ShadeHit(world, *comps, 5)

		if !got.Equal(want) {
			t.Errorf("\nwant: %v \ngot: %v \ndo not match", want, got)
		}

	})
}
