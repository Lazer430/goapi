package api

import (
	"encoding/json"
	"net/http"
)

// coin balance parameters struct
type CoinBalanceParams struct {
	Username string
}

// coin balance response struct
type CoinBalanceResponse struct {
	Balance   int
	ErrorCode int64
}

// error struct
type Error struct {
	Code    int64
	Message string
}

// time: 1:03:20
func writeError(w http.ResponseWriter, message string, code int64) {
	// create an error struct
	err := Error{
		Code:    code,
		Message: message,
	}

	// write the error to the response
	w.Header().Set("Content-Type", "application/json") // set the content type to json
	w.WriteHeader(int(code))                           // set the status code

	json.NewEncoder(w).Encode(err) // encode the error struct and write it to the response

}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal server error", http.StatusInternalServerError)
	}
)
