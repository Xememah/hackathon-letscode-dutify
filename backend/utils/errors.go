package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type ErrorResponse struct {
	Errors      []string `json:"errors"`
	DebugErrors []string `json:"debug_errors,omitempty"`
	Code        int      `json:"-"`
}

func NewErrorResponse(code int, errors ...error) *ErrorResponse {
	if len(errors) == 1 {
		if err, ok := errors[0].(*ErrorResponse); ok {
			return err
		}
	}

	resp := &ErrorResponse{Code: code}
	for _, err := range errors {
		if erp, ok := err.(*ErrorResponse); ok {
			resp.Errors = append(resp.Errors, erp.Errors...)
			resp.DebugErrors = append(resp.DebugErrors, erp.Errors...)
		} else {
			resp.Errors = append(resp.Errors, err.Error())
		}
	}
	return resp
}

func (re *ErrorResponse) Append(errors ...error) *ErrorResponse {
	for _, err := range errors {
		re.DebugErrors = append(re.Errors, err.Error())
	}
	return re
}

func (re *ErrorResponse) AppendDebug(errors ...error) *ErrorResponse {
	for _, err := range errors {
		re.DebugErrors = append(re.DebugErrors, err.Error())
	}
	return re
}

func (re *ErrorResponse) Error() string {
	return strings.Join(re.DebugErrors, ", ")
}

func (re *ErrorResponse) String() string {
	return re.Error()
}

func (re *ErrorResponse) Write(rw http.ResponseWriter) {
	rw.WriteHeader(re.Code)
	if os.Getenv("DEBUG") != "TRUE" {
		re.DebugErrors = []string{}
	}
	body, err := json.Marshal(re)
	if err != nil {
		panic(err)
	}
	if _, err := rw.Write(body); err != nil {
		panic(err)
	}
}
