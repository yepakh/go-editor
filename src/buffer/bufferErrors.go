package buffer

import "errors"

var ChangesNotSaved = errors.New("cannot close file, changes not saved")
