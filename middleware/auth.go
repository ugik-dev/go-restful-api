package middleware

import (
	"net/http"

	"github.com/ugik-dev/go-restful-api.git/helper"
	"github.com/ugik-dev/go-restful-api.git/models/web"
)

type Auth struct {
	http.Handler
}

func NewAuth(handler http.Handler) *Auth {
	return &Auth{Handler: handler}
}
func (auth *Auth) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// fmt.Println("Before LogMidleware")
	// fmt.Println("---- Execute : ", req.URL.Path)
	// auth.Handler.ServeHTTP(res, req)
	// fmt.Println("After LogMidleware")
	if req.Header.Get("X-API-KEY") == "SECRET" {
		auth.Handler.ServeHTTP(res, req)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZE",
		}
		helper.WriteResponse(res, webResponse)
	}
}
