package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/yamad07/go-modular-monolith/pkg/apperror"
)

var CodeStatuses = map[apperror.Code]int{
	apperror.CodeError:        http.StatusInternalServerError,
	apperror.CodeSuccess:      http.StatusOK,
	apperror.CodeNotFound:     http.StatusNotFound,
	apperror.CodeInvalid:      http.StatusBadRequest,
	apperror.CodeUnauthorized: http.StatusUnauthorized,
}

type Response struct {
	Message string `json:"message"`
}

func NewResponse(msg string) Response {
	return Response{
		msg,
	}
}

func ApplicationException(w http.ResponseWriter, aerr apperror.Error) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(CodeStatuses[aerr.Code()])

	body := NewResponse(aerr.Error())
	encode(w, body)
}

func InternalServerError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	body := NewResponse(err.Error())
	encode(w, body)
}

func BadRequestError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	body := NewResponse(err.Error())
	encode(w, body)
}

func Success(w http.ResponseWriter) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	body := NewResponse("success")
	encode(w, body)
}

func Encode(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	encode(w, body)
}

func encode(w http.ResponseWriter, body interface{}) {
	json.NewEncoder(w).Encode(body)
}
