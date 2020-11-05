package ctgov

import "errors"

// ErrNotParseable is an error instance for cases where parser recognizes corrupted input.
var ErrNotParseable = errors.New("Not parsable")
