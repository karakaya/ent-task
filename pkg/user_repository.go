package pkg

import (
	"github.com/rs/zerolog"
)

type UserRepository interface{}
type userRepository struct {
	logger zerolog.Logger
}

type User struct {
	ID uint64 `json:"id"`
}

func NewUserRepository(logger zerolog.Logger) UserRepository {
	return &userRepository{
		logger: logger,
	}
}
