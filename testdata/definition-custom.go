package testdata

import stderr "errors"

func DefineBad9() error {
	return stderr.New("bad defined 29") // want `do not define dynamic errors, use wrapped static errors instead: `
}
