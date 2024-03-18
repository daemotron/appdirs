package appdirs

import "errors"

// Custom errors
var (
	ErrAllUsersProfileNotDefined = errors.New("ALLUSERSPROFILE environment not defined") // %ALLUSERSPROFILE% environment not defined (Windows only)
)
