package routes

import (
	"entain-golang-task/cmd/api/user"
	"entain-golang-task/cmd/middleware"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

func DefineRoutes(logger zerolog.Logger, r *httprouter.Router) {
	r.POST("/user/:userId/transaction", middleware.Chain(user.NewUserTransactionService(logger, r), middleware.SourceTypeCheckMiddleware(logger), middleware.ContentTypeCheckMiddleware(logger)))

	//r.GET("/user/:userId/balance", middleware.Chain())
}
