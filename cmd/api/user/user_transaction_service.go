package user

import (
	"encoding/json"
	"entain-golang-task/cmd/api/user/internal"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/rs/zerolog"
)

type UserTransactionService struct {
	logger      zerolog.Logger
	userHandler *internal.UserTransactionHandler
}

func NewUserTransactionService(logger zerolog.Logger, r *httprouter.Router) httprouter.Handle {
	service := &UserTransactionService{
		logger:      logger,
		userHandler: internal.NewUserHandler(logger),
	}

	return service.Handle
}

func (s *UserTransactionService) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var input UserTransactionInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	s.userHandler.Handle(1)
}

type UserTransactionInput struct {
	State         string  `json:"state"`
	Amount        float64 `json:"amount"`
	TransactionId string  `json:"transactionId"`
}
