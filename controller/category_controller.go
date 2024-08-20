package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Crate(w http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, request *http.Request, params httprouter.Params)
	Filter(w http.ResponseWriter, request *http.Request, params httprouter.Params)
}
