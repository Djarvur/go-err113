package testdata

func NilEqual(e1 error) bool {
	return e1 == nil
}

func NilNotEqual(e1 error) bool {
	return e1 != nil
}
