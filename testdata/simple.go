package testdata

func SimpleEqual(e1, e2 error) bool {
	_ = 1 + 2       //just to increase the coverage
	return e1 == e2 // want `do not compare errors directly, use errors\.Is\(\) instead: "e1 == e2"`
}

func SimpleNotEqual(e1, e2 error) bool {
	return e1 != e2 // want `do not compare errors directly, use errors\.Is\(\) instead: "e1 \!= e2"`
}
