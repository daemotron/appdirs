package appdirs

import "errors"

var (
	// ErrAllUsersProfileNotDefined %ALLUSERSPROFILE% environment not defined (Windows only)
	ErrAllUsersProfileNotDefined = errors.New("ALLUSERSPROFILE environment not defined")
)
