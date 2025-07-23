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