package market

import (
	"github.com/Hunsin/compass/trade/pb"
)

// An Err implements the error interface.
type Err struct {
	s pb.Status
	e string
}

func (e *Err) Error() string {
	return "market: " + e.e
}

// Status returns the pb's status code.
func (e *Err) Status() pb.Status {
	return e.s
}

// Cancelled returns a pointer to Err with status CANCELLED
// and given message.
func Cancelled(msg string) error {
	return &Err{pb.Status_CANCELLED, msg}
}

// NotFound returns a pointer to Err with status NOT_FOUND
// and given message.
func NotFound(msg string) error {
	return &Err{pb.Status_NOT_FOUND, msg}
}

// Unimplemented returns a pointer to Err with status UNIMPLEMENTED
// and given message.
func Unimplemented(msg string) error {
	return &Err{pb.Status_UNIMPLEMENTED, msg}
}

// Unlisted returns a pointer to Err with status UNLISTED
// and given message.
func Unlisted(msg string) error {
	return &Err{pb.Status_UNLISTED, msg}
}

// Error returns a pointer to Err with given status code and message.
func Error(code pb.Status, msg string) error {
	return &Err{code, msg}
}
