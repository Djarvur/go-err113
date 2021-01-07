package testdata

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrWellDefined1 = errors.New("well defined 1")
	ErrWellDefined2 = fmt.Errorf("well defined 2 %d", 2)
	irrelevant1     = time.Now()
	irrelevant2     = irrelevant1.Add(time.Second)
)

var ErrWellDefined3 = fmt.Errorf("well defined 3 %d: %w", 3, ErrWellDefined1)

var ErrWellDefined4, ErrWellDefined5 = errors.New("well defined 4"), errors.New("well defined 5")

func DefineWell1() error {
	_ = time.Now().Add(time.Second)
	err := fmt.Errorf("well defined 1 %d: %w", 11, ErrWellDefined1)
	return err
}

func DefineWell2() error {
	return fmt.Errorf("well defined 12 %d: %w", 12, ErrWellDefined1)
}

const errWellTemplate3 = "well defined 13 %d: %w"

func DefineWell3() error {
	return fmt.Errorf(errWellTemplate3, 13, ErrWellDefined1)
}

func DefineWell4() error {
	return fmt.Errorf("well defined 14 xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx %d: %w", 14, ErrWellDefined1)
}

func DefineBad1() error {
	err := fmt.Errorf("bad defined 21 %d", 21) // want `do not define dynamic errors, use wrapped static errors instead: `
	return err
}

func DefineBad2() error {
	return fmt.Errorf("bad defined 22 %d", 22) // want `do not define dynamic errors, use wrapped static errors instead: `
}

func DefineBad3() error {
	err := errors.New("bad defined 23") // want `do not define dynamic errors, use wrapped static errors instead: `
	return err
}

func DefineBad4() error {
	return errors.New("bad defined 24") // want `do not define dynamic errors, use wrapped static errors instead: `
}

func DefineBad5() error {
	makeErr := func() error {
		return errors.New("bad defined 25") // want `do not define dynamic errors, use wrapped static errors instead: `
	}

	return makeErr()
}

const errBadTemplate6 = "bad defined 26 %d"

func DefineBad6() error {
	return fmt.Errorf(errBadTemplate6, 26) // want `do not define dynamic errors, use wrapped static errors instead: `
}

var errBadTemplate7 = "bad defined 27 %d"

func DefineBad7() error {
	return fmt.Errorf(errBadTemplate7, 27) // want `do not define dynamic errors, use wrapped static errors instead: `
}
