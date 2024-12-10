package fakeuseragent

import "errors"

var (
	ErrDatabaseIsEmpty = errors.New("database is empty")
	ErrNewOptions      = errors.New("failed to create new options")
)
