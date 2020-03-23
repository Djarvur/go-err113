package testdata

import eee "errors"

func DefineBad9() error {
	return eee.New("bad defined 29") // want `do not define dynamic errors, use wrapped static errors instead: `
}
