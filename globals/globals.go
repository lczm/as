package globals

import "github.com/lczm/as/errors"

// Errors to be updated here
var ErrorList []errors.Error

// Warnings to be updated here
var WarningList []errors.Error

func init() {
	ErrorList = make([]errors.Error, 0)
	WarningList = make([]errors.Error, 0)
}
