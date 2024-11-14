package routes

import (
	"entain-golang-task/cmd/api/user"
	"entain-golang-task/cmd/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

func DefineRoutes(logger zerolog.Logger, r *httprouter.Router) {
	r.POST("/user/:userId/transaction", middleware.SourceTypeCheckMiddleware(logger)(user.NewUserTransactionService(logger, r)))

	// r.Handle("GET /user/{userId}/balance", nil)
}
