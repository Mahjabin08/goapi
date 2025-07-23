package api

import (
	"encoding/json"
	"net/http"
)

//coin balance params
type CoinBalanceParams struct {
	Username string
}
type CoinBalanceResponse struct {
	//status code: success code, ususally 200
	Code int

	//account balance
	Balance int64
}
type Error struct{
	//error code
	Code int
	//error message
	Message string
}

func writeError (w http.ResponseWriter, message string,code int){
	resp := Error{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(),http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected error occurred",http.StatusInternalServerError)
	}
)