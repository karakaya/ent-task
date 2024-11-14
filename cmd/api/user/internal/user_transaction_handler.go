package internal

import (
	"entain-golang-task/pkg"

	"github.com/rs/zerolog"
)

type UserTransactionHandler struct {
	logger         zerolog.Logger
	userRepository pkg.UserRepository
}

func NewUserHandler(logger zerolog.Logger) *UserTransactionHandler {
	return &UserTransactionHandler{
		logger:         logger,
		userRepository: pkg.NewUserRepository(logger),
	}
}

func (h *UserTransactionHandler) Handle(userId uint64) error {

	return nil
}
