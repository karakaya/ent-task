package routes

import (
	"entain-golang-task/cmd/api/user"
	"net/http"

	"github.com/rs/zerolog"
)

func DefineRoutes(logger zerolog.Logger, r *http.ServeMux) {
	r.Handle("/", user.NewUserService(logger, r))
}
