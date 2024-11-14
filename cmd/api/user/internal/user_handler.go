package internal

import (
	"entain-golang-task/pkg"

	"github.com/rs/zerolog"
)

type UserHandler struct {
	logger         zerolog.Logger
	userRepository pkg.UserRepository
}

func NewUserHandler(logger zerolog.Logger) *UserHandler {
	return &UserHandler{
		logger:         logger,
		userRepository: pkg.NewUserRepository(logger),
	}
}

func (h *UserHandler) Handle(userId uint64) error {

	return nil
}
