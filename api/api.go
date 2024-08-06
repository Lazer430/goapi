package api

// import (
// 	"encoding/json"
// 	"net/http"
// )

// coin balance parameters struct
type CoinBalanceParams struct {
	username string
}

// coin balance response struct
type CoinBalanceResponse struct {
	balance   int
	errorCode int64
}

// error struct
type Error struct {
	code    int64
	message string
}

// time: 1:03:20
