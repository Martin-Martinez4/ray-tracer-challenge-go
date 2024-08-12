package main

import "fmt"

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func DefaultMaterial() Material {
	return Material{
		Color:     NewColor(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200.0,
	}
}

func (mat Material) Equal(other Material) bool {

	return mat.Color.Equal(other.Color) && mat.Ambient == other.Ambient && mat.Diffuse == other.Diffuse && mat.Specular == other.Specular && mat.Shininess == other.Shininess
}

func (mat *Material) Print() string {
	return fmt.Sprintf("color: %s, ambient: %f, diffuse: %f, specular: %f, shininess: %f", mat.Color.Print(), mat.Ambient, mat.Diffuse, mat.Specular, mat.Shininess)
}

func (mat *Material) SetColor(color Color) {
	mat.Color = color
}

func (mat *Material) SetAmbient(ambient float64) {
	mat.Ambient = ambient
}

func (mat *Material) SetDiffuse(diffuse float64) {
	mat.Diffuse = diffuse
}

func (mat *Material) SetSpecular(specular float64) {
	mat.Specular = specular
}

func (mat *Material) SetShininess(shininess float64) {
	mat.Shininess = shininess
}
