package world

import (
	"fmt"
	"math"
	"testing"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"
	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
)

func TestIntersectWorld(T *testing.T) {

	theWorld := NewDefaultWorld()

	tests := []struct {
		name string
		ray  shapes.Ray
		want shapes.Intersections
	}{
		{
			name: "a ray intersecting a default world should return an intersection struct with four members",
			ray:  shapes.NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			want: shapes.Intersections{
				{S: theWorld.Shapes[0], T: 4},
				{S: theWorld.Shapes[0], T: 4.5},
				{S: theWorld.Shapes[1], T: 5.5},
				{S: theWorld.Shapes[1], T: 6},
			},
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			got := RayWorldIntersect(tt.ray, theWorld)

			if len(tt.want) != len(got) {
				t.Errorf("Lengths did not match for test %d: %s", i, tt.name)
			}

			for k := 0; k < len(got); k++ {
				if !pm.AreFloatsEqual(tt.want[k].T, got[k].T) {
					t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", k, tt.want, got)
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
		ray          shapes.Ray
		sphere       shapes.Shape
		world        World
		intersection shapes.Intersection
		want         mat.Color
	}{
		{
			name:         "shading an intersection",
			ray:          shapes.NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			sphere:       theWorld.Shapes[0],
			world:        theWorld,
			intersection: shapes.NewIntersection(4, theWorld.Shapes[0]),
			want:         mat.NewColor(0.38066, 0.47583, 0.2855),
		},
		{
			name:         "shading an intersection from the inside",
			ray:          shapes.NewRay([3]float64{0, 0, 0}, [3]float64{0, 0, 1}),
			sphere:       theOtherWorld.Shapes[1],
			world:        theOtherWorld,
			intersection: shapes.NewIntersection(0.5, theOtherWorld.Shapes[1]),
			want:         mat.NewColor(0.90498, 0.90498, 0.90498),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := shapes.PrepareComputations(tt.ray, tt.sphere, tt.intersection)

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
		ray   shapes.Ray
		world World
		want  mat.Color
	}{
		{
			name:  "the color when a ray misses should be black (0,0,0)",
			ray:   shapes.NewRay([3]float64{0, 0, -5}, [3]float64{0, 1, 0}),
			world: theWorld,
			want:  mat.NewColor(0, 0, 0),
		},
		{
			name:  "the color when a ray hits",
			ray:   shapes.NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			world: theWorld,
			want:  mat.NewColor(0.38066, 0.47583, 0.2855),
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
		ray   shapes.Ray
		world World
		want  mat.Color
	}{
		{
			name:  "the color with an intersection behind the ray",
			ray:   shapes.NewRay([3]float64{0, 0, 0.75}, [3]float64{0, 0, -1}),
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
		point pm.Tuple
		want  bool
	}{
		{
			name:  "there is no shadow when nothing is collinear with point and light",
			world: NewDefaultWorld(),
			point: pm.Point(0, 10, 0),
			want:  false,
		},
		{
			name:  "there is no shadow when an object is behind the light",
			world: NewDefaultWorld(),
			point: pm.Point(-20, 20, -20),
			want:  false,
		},
		{
			name:  "the shadow exists when an object is between the point and the light",
			world: NewDefaultWorld(),
			point: pm.Point(10, -10, 10),
			want:  true,
		},
		{
			name:  "there is no shadow when an object is behind the point",
			world: NewDefaultWorld(),
			point: pm.Point(-22, 2, -2),
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
		from        pm.Tuple
		to          pm.Tuple
		up          pm.Tuple
		transform   pm.Matrix4x4
		pixelCoords [2]float64
		want        mat.Color
	}{
		{
			name:        "rendering a world with a camera",
			world:       NewDefaultWorld(),
			camera:      NewCamera(11, 11, math.Pi/2),
			from:        pm.Point(0, 0, -5),
			to:          pm.Point(0, 0, 0),
			up:          pm.Vector(0, 1, 0),
			pixelCoords: [2]float64{5, 5},
			want:        mat.NewColor(0.38066, 0.47583, 0.2855),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			tt.camera.Transform = pm.ViewTransformation(tt.from, tt.to, tt.up)

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

	s2 := shapes.NewSphere()
	s2.Transforms = s2.Transforms.Translate(0, 0, 10)

	theWorld.Shapes = []shapes.Shape{shapes.NewSphere(), s2}

	tests := []struct {
		name         string
		ray          shapes.Ray
		world        World
		intersection shapes.Intersection
		want         mat.Color
	}{
		{
			name:         "shading an intersection",
			ray:          shapes.NewRay([3]float64{0, 0, 5}, [3]float64{0, 0, 1}),
			world:        theWorld,
			intersection: shapes.NewIntersection(4, s2),
			want:         mat.NewColor(0.1, 0.1, 0.1),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := shapes.PrepareComputations(tt.ray, s2, tt.intersection)

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
		ray          shapes.Ray
		world        World
		intersection []shapes.Intersection
		want         mat.Color
	}{
		{
			name:         "shading an intersection",
			world:        theWorld,
			ray:          shapes.NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			intersection: []shapes.Intersection{shapes.NewIntersection(4, shape), shapes.NewIntersection(6, shape)},
			want:         mat.BLACK,
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := shapes.PrepareComputationsWithHit(tt.intersection[0], tt.ray, tt.intersection)

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
		ray          shapes.Ray
		world        World
		intersection []shapes.Intersection
		want         mat.Color
	}{
		{
			name:         "shading an intersection",
			world:        theWorld,
			ray:          shapes.NewRay([3]float64{0, 0, -5}, [3]float64{0, 0, 1}),
			intersection: []shapes.Intersection{shapes.NewIntersection(4, shape), shapes.NewIntersection(6, shape)},
			want:         mat.BLACK,
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := shapes.PrepareComputationsWithHit(tt.intersection[0], tt.ray, tt.intersection)

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
		ray          shapes.Ray
		world        World
		intersection []shapes.Intersection
		want         mat.Color
	}{
		{
			name:         "the refracted color under total internal relfection",
			world:        theWorld,
			ray:          shapes.NewRay([3]float64{0, 0, math.Sqrt(2) / 2}, [3]float64{0, 1, 0}),
			intersection: []shapes.Intersection{shapes.NewIntersection(-math.Sqrt(2)/2, shape), shapes.NewIntersection(math.Sqrt(2)/2, shape)},
			want:         mat.BLACK,
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {

			comps := shapes.PrepareComputationsWithHit(tt.intersection[1], tt.ray, tt.intersection)

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
	shapeA.GetMaterial().Pattern = mat.NewTestPattern(mat.BLACK, mat.WHITE)

	shapeB := theWorld.Shapes[1]
	shapeB.GetMaterial().Transparency = 1.0
	shapeB.GetMaterial().RefractiveIndex = 1.5

	tests := []struct {
		name         string
		ray          shapes.Ray
		world        World
		intersection []shapes.Intersection
		want         mat.Color
	}{
		{
			name:         "the refracted color under total internal relfection",
			world:        theWorld,
			ray:          shapes.NewRay([3]float64{0, 0, 0.1}, [3]float64{0, 1, 0}),
			intersection: []shapes.Intersection{shapes.NewIntersection(-0.9899, shapeA), shapes.NewIntersection(-0.4899, shapeB), shapes.NewIntersection(0.4899, shapeB), shapes.NewIntersection(0.9899, shapeA)},
			want:         mat.NewColor(0, 0.99888, 0.04725),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			comps := shapes.PrepareComputationsWithHit(tt.intersection[2], tt.ray, tt.intersection)

			got := RefreactedColor(tt.world, *comps, 5)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestRefractorShadeHit(T *testing.T) {

	theWorld := NewDefaultWorld()

	floor := shapes.NewPlane()
	floor.GetMaterial().Transparency = 0.5
	floor.GetMaterial().RefractiveIndex = 1.5
	floor.SetTransform(pm.Translate(0, -1, 0))

	ball := shapes.NewSphere()
	ball.GetMaterial().Color = mat.NewColor(1, 0, 0)
	ball.GetMaterial().Ambient = 0.5
	ball.SetTransform(pm.Translate(0, -3.5, -0.5))

	theWorld.Shapes = append(theWorld.Shapes, floor, ball)

	tests := []struct {
		name         string
		ray          shapes.Ray
		world        World
		intersection []shapes.Intersection
		want         mat.Color
	}{
		{
			name:         "ShadeHit with a transparent material",
			world:        theWorld,
			ray:          shapes.NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2}),
			intersection: []shapes.Intersection{shapes.NewIntersection(math.Sqrt(2), floor)},
			want:         mat.NewColor(0.93642, 0.68642, 0.68642),
		},
	}

	for i, tt := range tests {
		T.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			comps := shapes.PrepareComputationsWithHit(tt.intersection[0], tt.ray, tt.intersection)

			got := ShadeHit(theWorld, *comps, 5)

			if !got.Equal(tt.want) {
				t.Errorf("%d: \nwant: %v \ngot: %v \ndo not match", i, tt.want, got)
			}

		})
	}
}

func TestSchlickInternalReflection(T *testing.T) {

	shape := shapes.NewGlassSphere()

	name := "ShadeHit with a transparent material"
	ray := shapes.NewRay([3]float64{0, 0, math.Sqrt(2) / 2}, [3]float64{0, 1, 0})
	intersection := []shapes.Intersection{shapes.NewIntersection(-math.Sqrt(2)/2, shape), shapes.NewIntersection(math.Sqrt(2)/2, shape)}
	want := 1.0

	T.Run(name, func(t *testing.T) {
		comps := shapes.PrepareComputationsWithHit(intersection[1], ray, intersection)

		got := Schlick(comps)

		if !pm.AreFloatsEqual(got, want) {
			t.Errorf("\nwant: %v \ngot: %v \ndo not match", want, got)
		}

	})
}

func TestSchlickPerpendicular(T *testing.T) {

	shape := shapes.NewGlassSphere()

	name := "schilick with perpendicular viewing angle"
	ray := shapes.NewRay([3]float64{0, 0, 0}, [3]float64{0, 1, 0})
	intersection := []shapes.Intersection{shapes.NewIntersection(-1, shape), shapes.NewIntersection(1, shape)}
	want := 0.04

	T.Run(name, func(t *testing.T) {
		comps := shapes.PrepareComputationsWithHit(intersection[1], ray, intersection)

		got := Schlick(comps)

		if !pm.AreFloatsEqual(got, want) {
			t.Errorf("\nwant: %v \ngot: %v \ndo not match", want, got)
		}

	})
}

func TestSchlickSmallAngle(T *testing.T) {

	shape := shapes.NewGlassSphere()

	name := "schilick with small viewing angle, where n2 > n1"
	ray := shapes.NewRay([3]float64{0, 0.99, -2}, [3]float64{0, 0, 1})
	intersection := []shapes.Intersection{shapes.NewIntersection(1.8589, shape)}
	want := 0.48873

	T.Run(name, func(t *testing.T) {
		comps := shapes.PrepareComputationsWithHit(intersection[0], ray, intersection)

		got := Schlick(comps)

		if !pm.AreFloatsEqual(got, want) {
			t.Errorf("\nwant: %v \ngot: %v \ndo not match", want, got)
		}

	})
}

func TestSchlickReflectionAndRefraction(T *testing.T) {

	name := "schilick with reflection and refraction"
	ray := shapes.NewRay([3]float64{0, 0, -3}, [3]float64{0, -math.Sqrt(2) / 2, math.Sqrt(2) / 2})
	world := NewDefaultWorld()

	floor := shapes.NewPlane()
	floor.SetTransform(pm.Translate(0, -1, 0))
	floor.GetMaterial().Reflective = 0.5
	floor.GetMaterial().Transparency = 0.5
	floor.GetMaterial().RefractiveIndex = 1.5

	ball := shapes.NewSphere()
	ball.GetMaterial().Color = mat.NewColor(1, 0, 0)
	ball.GetMaterial().Ambient = 0.5
	ball.SetTransform(pm.Translate(0, -3.5, -0.5))

	world.Shapes = append(world.Shapes, floor, ball)

	intersection := []shapes.Intersection{shapes.NewIntersection(math.Sqrt(2), floor)}
	want := mat.NewColor(0.93391, 0.69643, 0.69243)

	T.Run(name, func(t *testing.T) {
		comps := shapes.PrepareComputationsWithHit(intersection[0], ray, intersection)

		got := ShadeHit(world, *comps, 5)

		if !got.Equal(want) {
			t.Errorf("\nwant: %v \ngot: %v \ndo not match", want, got)
		}

	})
}
