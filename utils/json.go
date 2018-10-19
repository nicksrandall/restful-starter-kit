package utils

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/nicksrandall/restful-starter-kit/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Write(res http.ResponseWriter, model interface{}, err error) {
	if err != nil {
		HandleError(res, err)
	} else {
		res.Header().Set("Content-Type", "application/json; charset=utf8")
		json.NewEncoder(res).Encode(model)
	}
}

func HandleError(res http.ResponseWriter, err error) {
	if httpError, ok := err.(errors.HTTPError); ok {
		http.Error(res, httpError.Error(), httpError.StatusCode())
	} else {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
