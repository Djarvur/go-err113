package testdata

func SimpleEqual(e1, e2 error) bool {
	_ = 1 + 2       //just to increase the coverage
	return e1 == e2 // want `do not compare errors directly \"e1 == e2\", use \"errors\.Is\(e1, e2\)\" instead`
}

func SimpleNotEqual(e1, e2 error) bool {
	return e1 != e2 // want `do not compare errors directly \"e1 != e2\", use \"!errors\.Is\(e1, e2\)\" instead`
}
