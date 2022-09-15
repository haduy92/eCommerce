package errs

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
)

type Error struct {
	// Kind is the class of error, such as permission failure,
	// or "Other" if its class is unknown or irrelevant.
	Kind Kind
	// Param represents the parameter related to the error.
	Param Parameter
	// Code is a human-readable, short representation of the error
	Code Code
	// The underlying error that triggered this one, if any.
	Err error
}

func (e *Error) isZero() bool {
	return e.Kind == 0 && e.Param == "" && e.Code == "" && e.Err == nil
}

// Unwrap method allows for unwrapping errors using errors.As
func (e Error) Unwrap() error {
	return e.Err
}

func (e *Error) Error() string {
	return e.Err.Error()
}

type Kind uint8

type Parameter string

type Code string

const (
	Other          Kind = iota // Unclassified error. This value is not printed in the error message.
	Invalid                    // Invalid operation for this type of item.
	Exist                      // Item already exists.
	NotExist                   // Item does not exist.
	Private                    // Information withheld.
	Internal                   // Internal error or inconsistency.
	BrokenLink                 // Link target does not exist.
	Database                   // Error from database.
	Validation                 // Input validation error.
	InvalidRequest             // Invalid Request
	Unanticipated              // Unanticipated error.
)

func (k Kind) String() string {
	switch k {
	case Other:
		return "other_error"
	case Invalid:
		return "invalid_operation"
	case Exist:
		return "item_already_exists"
	case NotExist:
		return "item_does_not_exist"
	case BrokenLink:
		return "link_target_does_not_exist"
	case Private:
		return "information_withheld"
	case Internal:
		return "internal_error"
	case Database:
		return "database_error"
	case Validation:
		return "input_validation_error"
	case InvalidRequest:
		return "invalid_request_error"
	case Unanticipated:
		return "unanticipated_error"
	}
	return "unknown_error_kind"
}

// E builds an error value from its arguments.
// There must be at least one argument or E panics.
// The type of each argument determines its meaning.
// If more than one argument of a given type is presented,
// only the last one is recorded.
//
// The types are:
//
//	UserName
//		The username of the user attempting the operation.
//	string
//		Treated as an error message and assigned to the
//		Err field after a call to errors.New.
//	errors.Kind
//		The class of error, such as permission failure.
//	error
//		The underlying error that triggered this one.
//
// If the error is printed, only those items that have been
// set to non-zero values will appear in the result.
//
// If Kind is not specified or Other, we set it to the Kind of
// the underlying error.
func E(args ...interface{}) error {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	if len(args) == 0 {
		panic("call to errors.E with no arguments")
	}
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case string:
			e.Err = errors.New(arg)
		case Kind:
			e.Kind = arg
		case *Error:
			e.Err = arg
		case error:
			// if the error implements stackTracer, then it is
			// a pkg/errors error type and does not need to have
			// the stack added
			_, ok := arg.(stackTracer)
			if ok {
				e.Err = arg
			} else {
				e.Err = errors.WithStack(arg)
			}
		case Code:
			e.Code = arg
		case Parameter:
			e.Param = arg
		default:
			_, file, line, _ := runtime.Caller(1)
			return fmt.Errorf("errors.E: bad call from %s:%d: %v, unknown type %T, value %v in error call", file, line, args, arg, arg)
		}
	}

	prev, ok := e.Err.(*Error)
	if !ok {
		return e
	}
	// If this error has Kind unset or Other, pull up the inner one.
	if e.Kind == Other {
		e.Kind = prev.Kind
		prev.Kind = Other
	}

	if prev.Code == e.Code {
		prev.Code = ""
	}
	// If this error has Code == "", pull up the inner one.
	if e.Code == "" {
		e.Code = prev.Code
		prev.Code = ""
	}

	if prev.Param == e.Param {
		prev.Param = ""
	}
	// If this error has Param == "", pull up the inner one.
	if e.Param == "" {
		e.Param = prev.Param
		prev.Param = ""
	}

	return e
}

// Match compares its two error arguments. It can be used to check
// for expected errors in tests. Both arguments must have underlying
// type *Error or Match will return false. Otherwise it returns true
// if every non-zero element of the first error is equal to the
// corresponding element of the second.
// If the Err field is a *Error, Match recurs on that field;
// otherwise it compares the strings returned by the Error methods.
// Elements that are in the second argument but not present in
// the first are ignored.
//
// For example,
//
//		Match(errors.E(upspin.UserName("joe@schmoe.com"), errors.Permission), err)
//	 tests whether err is an Error with Kind=Permission and User=joe@schmoe.com.
func Match(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}
	e1, ok := err1.(*Error)
	if !ok {
		return false
	}
	e2, ok := err2.(*Error)
	if !ok {
		return false
	}
	if e1.Kind != Other && e2.Kind != e1.Kind {
		return false
	}
	if e1.Param != "" && e2.Param != e1.Param {
		return false
	}
	if e1.Code != "" && e2.Code != e1.Code {
		return false
	}
	if e1.Err != nil {
		if _, ok := e1.Err.(*Error); ok {
			return Match(e1.Err, e2.Err)
		}
		if e2.Err == nil || e2.Err.Error() != e1.Err.Error() {
			return false
		}
	}
	return true
}

// KindIs reports whether err is an *Error of the given Kind.
// If err is nil then KindIs returns false.
func KindIs(kind Kind, err error) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	if e.Kind != Other {
		return e.Kind == kind
	}
	if e.Err != nil {
		return KindIs(kind, e.Err)
	}
	return false
}
