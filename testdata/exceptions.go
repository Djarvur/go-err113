package testdata

import "io"

func NilEqual(e1 error) bool {
	return e1 == nil
}

func NilNotEqual(e1 error) bool {
	return e1 != nil
}

func EOFEqual(e1 error) bool {
	return e1 == io.EOF
}

func EOFNotEqual(e1 error) bool {
	return e1 != io.EOF
}
