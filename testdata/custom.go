package testdata

type CustomError struct {
	msg string
}

func (e CustomError) Error() string {
	return e.msg
}

var _ error = (*CustomError)(nil)

func CustomEqual(e1, e2 CustomError) bool {
	return e1 == e2
}

func CustomNotEqual(e1, e2 CustomError) bool {
	return e1 != e2
}

func CustomCastedEqual(ce1, ce2 CustomError) bool {
	var e1, e2 error = ce1, ce2
	return e1 == e2 // want `do not compare errors directly \"e1 == e2\", use \"errors\.Is\(e1, e2\)\" instead`
}

func CustomCastedNotEqual(ce1, ce2 CustomError) bool {
	var e1, e2 error = ce1, ce2
	return e1 != e2 // want `do not compare errors directly \"e1 != e2\", use \"!errors\.Is\(e1, e2\)\" instead`
}

func CustomPCastedEqual(ce1, ce2 *CustomError) bool {
	var e1, e2 error = ce1, ce2
	return e1 == e2 // want `do not compare errors directly \"e1 == e2\", use \"errors\.Is\(e1, e2\)\" instead`
}

func CustomPCastedNotEqual(ce1, ce2 *CustomError) bool {
	var e1, e2 error = ce1, ce2
	return e1 != e2 // want `do not compare errors directly \"e1 != e2\", use \"!errors\.Is\(e1, e2\)\" instead`
}

func CustomHalfCastedEqual(e1 CustomError, e2 error) bool {
	return e1 == e2 // want `do not compare errors directly \"e1 == e2\", use \"errors\.Is\(e1, e2\)\" instead`
}
