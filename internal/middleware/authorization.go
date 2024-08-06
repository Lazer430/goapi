package middleware

import (
	"errors"
	"fmt"
	"net/http"
	// "github.com/Lazer430/goapi/internal/tools"
	// "github.com/Lazer430/goapi/api"
	// log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Unauthorized username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the username from the request
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		// check if the username and token are valid
		if username == "" || token == "" {
			fmt.Println("Username or token is empty")
		}

	})
}
