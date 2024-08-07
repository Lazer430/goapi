package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Lazer430/goapi/api"
	"github.com/Lazer430/goapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

var ErrUnauthorized = errors.New("unauthorized username or token")

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		api.RequestErrorHandler(w, ErrUnauthorized)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance:   int((*tokenDetails).Coins),
		ErrorCode: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
