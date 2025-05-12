package fault

import (
	"fmt"
	"net/http"
)

type Fault struct {
	HttpStatusCode int    `json:"-"`
	Err            error  `json:"-"`
	Message        string `json:"message"`
}

func (f *Fault) Error() string {
	if f.Err != nil {
		return fmt.Sprintf("%s (caused by: %v)", f.Message, f.Err)
	}

	return fmt.Sprintf("%s", f.Message)
}

func NewFault(message string, err error, httpStatusCode int) *Fault {
	fault := Fault{
		HttpStatusCode: httpStatusCode,
		Err:            err,
		Message:        message,
	}

	return &fault
}

func NewNotFoundFault(message string) *Fault {
	return NewFault(
		message,
		nil,
		http.StatusNotFound,
	)
}

func NewUnprocessableEntityFault(message string) *Fault {
	return NewFault(
		message,
		nil,
		http.StatusUnprocessableEntity,
	)
}
