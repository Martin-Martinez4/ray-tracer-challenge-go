package main

func ViewTransformation(from Tuple, to Tuple, up Tuple) Matrix4x4 {

	// Compute the forward vector by subtracting from from to.  Normalize the result
	// Compute the left vector by taking the cross product of forward and the normalized up vector
	// compute the trueUp vector by taking the cross product of the left and forward.
	/*
		Construct the orientation matrix
		[
			leftVec.x,		leftVec.y, 	leftVec.z, 	0
			trueUp.x,	trueUp.y, 	trueUp.z, 	0
			-forwardVec.x,	-forwardVec.y,	-forwardVec.z,	0
			0,			0,			0,			1
		]
	*/
	// Multiply orientation by translation(-from.x, -from.y, -from.z)

	forwardVec := Normalize(to.Subtract(from))
	leftVec := Cross(forwardVec, Normalize(up))
	trueUp := Cross(leftVec, forwardVec)

	orientationMat := NewMatrix4x4([16]float64{leftVec.x, leftVec.y, leftVec.z, 0, trueUp.x, trueUp.y, trueUp.z, 0, -forwardVec.x, -forwardVec.y, -forwardVec.z, 0, 0, 0, 0, 1})

	id := IdentitiyMatrix4x4().Translate(-from.x, -from.y, -from.z)

	return orientationMat.Multiply(id)

}
