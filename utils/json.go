package utils

import (
	"encoding/json"
	"net/http"

	"github.com/nicksrandall/restful-starter-kit/errors"
)

func Write(res http.ResponseWriter, model interface{}, err error) {
	if err != nil {
		if httpError, ok := err.(errors.HTTPError); ok {
			http.Error(res, httpError.Error(), httpError.StatusCode())
		} else {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	} else {
		res.Header().Set("Content-Type", "application/json; charset=utf8")
		json.NewEncoder(res).Encode(model)
	}
}
