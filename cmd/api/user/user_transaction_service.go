package user

import (
	"entain-golang-task/cmd/api/user/internal"
	"net/http"

	"github.com/rs/zerolog"
)

type UserTransactionService struct {
	logger      zerolog.Logger
	userHandler *internal.UserTransactionHandler
}

func NewUserTransactionService(logger zerolog.Logger, r *http.ServeMux) http.HandlerFunc {
	service := &UserTransactionService{
		logger:      logger,
		userHandler: internal.NewUserHandler(logger),
	}

	return service.Handle
}

func (s *UserTransactionService) Handle(w http.ResponseWriter, r *http.Request) {
	s.logger.Info().Msg("HIT!")
	s.userHandler.Handle(1)
}
