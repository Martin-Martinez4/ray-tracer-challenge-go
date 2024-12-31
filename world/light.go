package world

import (
	"math"

	mat "github.com/Martin-Martinez4/ray-tracer-challenge-go/materials"
	pm "github.com/Martin-Martinez4/ray-tracer-challenge-go/primitive_math"

	"github.com/Martin-Martinez4/ray-tracer-challenge-go/shapes"
)

type Light struct {
	Intensity mat.Color
	Position  pm.Tuple
}

func NewLight(position, intensity [3]float64) Light {
	return Light{
		Position:  pm.Point(position[0], position[1], position[2]),
		Intensity: mat.NewColor(intensity[0], intensity[1], intensity[2]),
	}
}

/*
	Light() expects five arguments:
		the material,
		the point being illuminated,
		the light source, and
		the eye and
		the normal vectors from the Phong reflection model
*/

func EffectiveLighting(material mat.Material, shape shapes.Shape, light Light, point pm.Tuple, eyeVec pm.Tuple, normalVec pm.Tuple, inShadow bool) mat.Color {

	color := material.Color

	if material.Pattern != nil {

		// color = material.Pattern.PatternAtShape(shape, point)
		color = shape.PatternAtShape(material.Pattern, point)
	}

	effectiveColor := color.Multiply(light.Intensity)

	lightVec := pm.Normalize(light.Position.Subtract(point))

	ambient := effectiveColor.SMultiply(material.Ambient)

	lightDotNormal := pm.Dot(lightVec, normalVec)

	var specular mat.Color
	var diffuse mat.Color

	if lightDotNormal < 0 {

		specular = mat.NewColor(0, 0, 0)
		diffuse = mat.NewColor(0, 0, 0)

	} else {
		diffuse = effectiveColor.SMultiply(material.Diffuse)
		diffuse = diffuse.SMultiply(lightDotNormal)

		reflectVec := lightVec.SMultiply(-1)
		reflectVec = reflectVec.ReflectBy(normalVec)

		reflectDotEye := pm.Dot(reflectVec, eyeVec)

		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, material.Shininess)
			specular = light.Intensity.SMultiply(material.Specular * factor)
		}

	}

	if inShadow {
		return ambient
	}

	return (ambient.Add(diffuse.Add(specular)))
}
