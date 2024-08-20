package exception

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/ugik-dev/go-restful-api.git/helper"
	"github.com/ugik-dev/go-restful-api.git/models/web"
)

func PanicHandler(res http.ResponseWriter, req *http.Request, err interface{}) {
	if notFoundError(res, req, err) {
		return
	}

	if validationError(res, req, err) {
		return
	}

	internalServerError(res, req, err)
}
func validationError(res http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		helper.WriteResponse(res, webResponse)
		return true
	} else {
		return false
	}
}
func notFoundError(res http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "DATA NOT FOUND",
			Data:   exception.Error,
		}
		helper.WriteResponse(res, webResponse)
		return true
	} else {
		return false
	}
}
func internalServerError(res http.ResponseWriter, req *http.Request, err interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}
	helper.WriteResponse(res, webResponse)
}

type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
