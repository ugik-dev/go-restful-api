package helper

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&result)
	PanicIfError(err)
}

func WriteResponse(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-type", "application/json")
	encode := json.NewEncoder(w)
	err := encode.Encode(response)
	PanicIfError(err)
}
