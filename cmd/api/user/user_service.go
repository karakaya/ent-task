package user

import (
	"entain-golang-task/cmd/api/user/internal"
	"net/http"

	"github.com/rs/zerolog"
)

type UserService struct {
	logger      zerolog.Logger
	userHandler *internal.UserHandler
}

func NewUserService(logger zerolog.Logger, r *http.ServeMux) http.HandlerFunc {
	service := &UserService{
		logger:      logger,
		userHandler: internal.NewUserHandler(logger),
	}

	return service.Handle
}

func (s *UserService) Handle(w http.ResponseWriter, r *http.Request) {
	s.logger.Info().Msg("HIT!")
	s.userHandler.Handle(1)
	w.WriteHeader(http.StatusNotFound)
}
