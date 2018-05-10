package errstack

import (
	"runtime"
)

type Error struct {
	msg     string
	err     error
	callers []uintptr
}

func (e *Error) Error() string { return e.msg } // dummy

// Callers returns the error's calling PCs.
func (e *Error) Callers() []uintptr {
	if e.callers != nil {
		return e.callers
	}
	if e2, ok := e.err.(*Error); ok {
		return e2.Callers()
	}
	return nil
}

func wrap1(err error, msg string) error {
	e := &Error{msg: msg, err: err}
	if _, ok := err.(*Error); !ok {
		e.callers = callers()
	}
	return e
}

func wrap2(err error, msg string) error {
	e := &Error{msg: msg, err: err, callers: callers()}

	e2, ok := err.(*Error)
	if !ok {
		return e
	}

	// Move distinct callers from inner error to outer error and throw the
	// common callers away so that we only print the stack trace once.
	i := 0
	shared := false
	for ; i < len(e.callers) && i < len(e2.callers); i++ {
		if e.callers[len(e.callers)-1-i] != e2.callers[len(e2.callers)-1-i] {
			break
		}
		shared = true
	}
	if shared {
		// The stacks have a PC suffix in common: e2.callers[len(e2.callers)-i:]
		head := e2.callers[:len(e2.callers)-i]
		tail := e.callers
		e.callers = make([]uintptr, len(head)+len(tail))
		copy(e.callers, head)
		copy(e.callers[len(head):], tail)
		e2.callers = nil
	}

	return e
}

// Wrapper for runtime.Callers that allocates a slice.
func callers() []uintptr {
	var stack [64]uintptr
	const skip = 4 // skip callers(), wrap(), and New/Wrap/etc.
	n := runtime.Callers(skip, stack[:])
	return stack[:n]
}

func getCallers(err error) []uintptr {
	if e2, ok := err.(*Error); ok {
		return e2.Callers()
	}
	return nil
}
