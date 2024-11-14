package routes

import (
	"entain-golang-task/cmd/api/user"
	"net/http"

	"github.com/rs/zerolog"
)

func DefineRoutes(logger zerolog.Logger, r *http.ServeMux) {
	r.Handle("POST /user/{userId}/transaction", user.NewUserTransactionService(logger, r))
	// r.Handle("GET /user/{userId}/balance", nil)
}
