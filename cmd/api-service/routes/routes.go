package routes

import (
	"entain-golang-task/cmd/api-service/user"
	"entain-golang-task/cmd/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

func DefineRoutes(logger zerolog.Logger, r *httprouter.Router) {
	r.POST("/user/:userId/transaction", middleware.Chain(user.NewUserTransactionService(logger), middleware.SourceTypeCheckMiddleware(logger), middleware.ContentTypeCheckMiddleware(logger)))

	r.GET("/user/:userId/balance", user.NewUserAccountBalanceService(logger))
}
