package testdata

import stdio "io"

func stdioEOFEqual(e1 error) bool {
	return e1 == stdio.EOF
}

func stdioEOFNotEqual(e1 error) bool {
	return e1 != stdio.EOF
}
