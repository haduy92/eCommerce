package errs

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type ErrResponse struct {
	Error ServiceError `json:"error"`
}

type ServiceError struct {
	Kind    string `json:"kind,omitempty"`
	Code    string `json:"code,omitempty"`
	Param   string `json:"param,omitempty"`
	Message string `json:"message,omitempty"`
}

// HTTPErrorResponse takes a writer, error and a logger, performs a
// type switch to determine if the type is an Error (which meets
// the Error interface as defined in this package), then sends the
// Error as a response to the client. If the type does not meet the
// Error interface as defined in this package, then a proper error
// is still formed and sent to the client, however, the Kind and
// Code will be Unanticipated.
func HTTPErrorResponse(w http.ResponseWriter, err error) {
	if err == nil {
		nilErrorResponse(w)
		return
	}

	var e *Error
	if errors.As(err, &e) {
		typicalErrorResponse(w, e)
		return
	}

	unknownErrorResponse(w, err)
}

// typicalErrorResponse replies to the request with the specified error
// message and HTTP code. It does not otherwise end the request; the
// caller should ensure no further writes are done to w.
func typicalErrorResponse(w http.ResponseWriter, e *Error) {

	httpStatusCode := httpErrorStatusCode(e.Kind)

	// We can retrieve the status here and write out a specific
	// HTTP status code. If the error is empty, just send the HTTP
	// Status Code as response. Error should not be empty, but it's
	// theoretically possible, so this is just in case...
	if e.isZero() {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// log the error with stacktrace
	log.WithFields(log.Fields{
		"http_status_code": httpStatusCode,
		"kind":             e.Kind.String(),
		"parameter":        string(e.Param),
		"code":             string(e.Code),
	}).Error(e.Err)

	// get ErrResponse
	er := newErrResponse(e)

	// Marshal errResponse struct to JSON for the response body
	errJSON, _ := json.Marshal(er)
	ej := string(errJSON)

	// Write HTTP Status Code
	w.WriteHeader(httpStatusCode)

	// Write response body (json)
	fmt.Fprintln(w, ej)
}

func newErrResponse(err *Error) ErrResponse {
	const msg string = "Internal server error - please contact support"

	switch err.Kind {
	case Internal, Database:
		return ErrResponse{
			Error: ServiceError{
				Kind:    Internal.String(),
				Message: msg,
			},
		}
	default:
		return ErrResponse{
			Error: ServiceError{
				Kind:    err.Kind.String(),
				Code:    string(err.Code),
				Param:   string(err.Param),
				Message: err.Error(),
			},
		}
	}
}

// nilErrorResponse responds with http status code 500 (Internal Server Error)
// and an empty response body. nil error should never be sent, but in case it is...
func nilErrorResponse(w http.ResponseWriter) {
	log.WithFields(log.Fields{
		"http_status_code": http.StatusInternalServerError,
	}).Error("nil error - no response body sent")
	w.WriteHeader(http.StatusInternalServerError)
}

// unknownErrorResponse responds with http status code 500 (Internal Server Error)
// and a json response body with unanticipated_error kind
func unknownErrorResponse(w http.ResponseWriter, err error) {
	er := ErrResponse{
		Error: ServiceError{
			Kind:    Unanticipated.String(),
			Code:    "Unanticipated",
			Message: "Unexpected error - contact support",
		},
	}

	log.WithFields(log.Fields{
		"http_status_code": http.StatusInternalServerError,
		"message":          err.Error(),
	}).Error("Unknown Error")

	// Marshal errResponse struct to JSON for the response body
	errJSON, _ := json.Marshal(er)
	ej := string(errJSON)

	// Write HTTP Status Code
	w.WriteHeader(http.StatusInternalServerError)

	// Write response body (json)
	fmt.Fprintln(w, ej)
}

// httpErrorStatusCode maps an error Kind to an HTTP Status Code
func httpErrorStatusCode(k Kind) int {
	switch k {
	case Invalid, Exist, NotExist, Private, BrokenLink, Validation, InvalidRequest:
		return http.StatusBadRequest
	// the zero value of Kind is Other, so if no Kind is present
	// in the error, Other is used. Errors should always have a
	// Kind set, otherwise, a 500 will be returned and no
	// error message will be sent to the caller
	case Other, Internal, Database:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
