package main

import "math"

type Light struct {
	Intensity Color
	Position  Tuple
}

func NewLight(position, intensity [3]float64) Light {
	return Light{
		Intensity: NewColor(intensity[0], intensity[1], intensity[2]),
		Position:  Point(position[0], position[1], position[2]),
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

func EffectiveLighting(mat Material, shape Shape, light Light, point Tuple, eyeVec Tuple, normalVec Tuple, inShadow bool) Color {

	var color Color

	if mat.Pattern == nil {
		color = mat.Color
	} else {
		color = mat.Pattern.PatternAtShape(shape, point)
	}

	effectiveColor := color.Multiply(light.Intensity)

	lightVec := Normalize(light.Position.Subtract(point))

	ambient := effectiveColor.SMultiply(mat.Ambient)

	lightDotNormal := Dot(lightVec, normalVec)

	var specular Color
	var diffuse Color

	if lightDotNormal < 0 {

		specular = NewColor(0, 0, 0)
		diffuse = NewColor(0, 0, 0)

	} else {
		diffuse = effectiveColor.SMultiply(mat.Diffuse)
		diffuse = diffuse.SMultiply(lightDotNormal)

		reflectVec := lightVec.SMultiply(-1)
		reflectVec = reflectVec.ReflectBy(normalVec)

		reflectDotEye := Dot(reflectVec, eyeVec)

		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, mat.Shininess)
			specular = light.Intensity.SMultiply(mat.Specular * factor)
		}

	}

	if inShadow {
		return ambient
	}

	return (ambient.Add(diffuse.Add(specular)))
}
