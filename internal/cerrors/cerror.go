package cerrors

import (
	"encoding/json"
	"net/http"
)

type CError struct {
	Code    int              `json:"code"`
	Message map[string][]any `json:"message"`
}

func New(c int, msg map[string][]any) *CError {
	return &CError{
		Code:    c,
		Message: msg,
	}
}

func (e *CError) Error() string {
	for k, v := range e.Message {
		return k + ": " + v[0].(string)
	}
	return ""
}

func (e *CError) WithCode(c int) *CError {
	e.Code = c
	return e
}

func (e *CError) Render(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	return json.NewEncoder(w).Encode(e)
}

func HandleError(e error, w http.ResponseWriter) {
	if err, ok := e.(*CError); ok {
		err.Render(w)
		return
	}
	ErrInternal.Render(w)
}

func WrapError(e *CError, err error) *CError {
	newMessage := make(map[string][]any)
	for k, v := range e.Message {
		newMessage[k] = append([]any{}, v...)
	}
	for k, v := range newMessage {
		for i := range v {
			if str, ok := v[i].(string); ok {
				newMessage[k][i] = map[string]string{str: err.Error()}
			} else {
				newMessage[k] = append(newMessage[k], map[string]string{"error": err.Error()})
			}
		}
	}

	return &CError{
		Code:    e.Code,
		Message: newMessage,
	}
}
