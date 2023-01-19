package errors

import "net/http"

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, message string) *ResponseError {
	return &ResponseError{
		Code:    code,
		Message: message,
	}
}

func (e *ResponseError) Render(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	_, err := w.Write([]byte(e.Message))
	return err
}
