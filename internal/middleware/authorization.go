package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Lazer430/goapi/api"
	"github.com/Lazer430/goapi/internal/tools"
)

var ErrUnauthorized = errors.New("unauthorized username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the username from the request
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		// check if the username and token are valid
		if username == "" || token == "" {
			fmt.Println("Username or token is empty")
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		var database *tools.DatabaseInterface
		var err error
		database, err = tools.NewDatabase()

		if err != nil {
			fmt.Println("Error connecting to database")
			api.RequestErrorHandler(w, err)
			return
		}

		// get the acceptable usernames from the database
		loginDetails := (*database).GetLoginDetails(username)

		if loginDetails == nil || token != (*loginDetails).AuthToken {
			fmt.Println("Invalid username or token")
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}
