package globals

import "github.com/lczm/as/errors"

var ErrorList []errors.Error

func init() {
	ErrorList = make([]errors.Error, 0)
}
