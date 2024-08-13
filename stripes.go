package main

type Stripes struct {
	Color1     Color
	Color2     Color
	Transforms Matrix4x4
}

func Stripe(color1, color2 Color) *Stripes {

	return &Stripes{Color1: color1, Color2: color2, Transforms: IdentitiyMatrix4x4()}
}

func (stripe *Stripes) GetColor1() Color {
	return stripe.Color1
}

func (stripe *Stripes) GetColor2() Color {
	return stripe.Color2
}

func (stripe *Stripes) SetTransform(mat44 *Matrix4x4) Matrix4x4 {

	stripe.Transforms = mat44.Multiply(stripe.Transforms)
	return stripe.Transforms
}

func (stripe *Stripes) SetTransforms(mat44 []Matrix4x4) {

	for _, transform := range mat44 {

		stripe.SetTransform(&transform)
	}
}

func (stripe *Stripes) GetTransforms() Matrix4x4 {
	return stripe.Transforms
}

func (stripe *Stripes) StripeAt(point Tuple) Color {

	var color Color

	if int(point.x)%2 == 0 {
		color = stripe.Color2
	} else {
		color = stripe.Color1
	}

	if point.x-float64((int(point.x))) < 0 {
		if color.Equal(stripe.Color1) {
			color = stripe.Color2

		} else {
			color = stripe.Color1
		}
	}

	return color
}

func (stripe *Stripes) PatternAtShape(object Shape, worldPoint Tuple) Color {
	inversObjTransform := object.GetTransforms().Inverse()

	objectPoint := inversObjTransform.TupleMultiply(worldPoint)

	inversPatTransform := stripe.GetTransforms().Inverse()
	patternPoint := inversPatTransform.TupleMultiply(objectPoint)

	return stripe.StripeAt(patternPoint)

}
