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

// NotFound returns a pointer to Err with status
// pb.NOT_FOUND and given message.
func NotFound(msg string) error {
	return &Err{pb.Status_NOT_FOUND, msg}
}

// Unimplemented returns a pointer to Err with status
// pb.UNIMPLEMENTED and given message.
func Unimplemented(msg string) error {
	return &Err{pb.Status_UNIMPLEMENTED, msg}
}

// Error returns a pointer to Err with given status code and message.
func Error(code pb.Status, msg string) error {
	return &Err{code, msg}
}
