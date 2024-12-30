package primitive_math

func ViewTransformation(from Tuple, to Tuple, up Tuple) Matrix4x4 {

	// Compute the forward vector by subtracting from from to.  Normalize the result
	// Compute the left vector by taking the cross product of forward and the normalized up vector
	// compute the trueUp vector by taking the cross product of the left and forward.
	/*
		Construct the orientation matrix
		[
			leftVec.X,		leftVec.Y, 	leftVec.Z, 	0
			trueUp.X,	trueUp.Y, 	trueUp.Z, 	0
			-forwardVec.X,	-forwardVec.Y,	-forwardVec.Z,	0
			0,			0,			0,			1
		]
	*/
	// Multiply orientation by translation(-from.X, -from.Y, -from.Z)

	forwardVec := Normalize(to.Subtract(from))
	leftVec := Cross(forwardVec, Normalize(up))
	trueUp := Cross(leftVec, forwardVec)

	orientationMat := NewMatrix4x4([16]float64{leftVec.X, leftVec.Y, leftVec.Z, 0, trueUp.X, trueUp.Y, trueUp.Z, 0, -forwardVec.X, -forwardVec.Y, -forwardVec.Z, 0, 0, 0, 0, 1})

	id := Translate(-from.X, -from.Y, -from.Z)

	return orientationMat.Multiply(*id)

}
