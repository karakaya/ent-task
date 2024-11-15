package user

import (
	"context"
	"encoding/json"
	"entain-golang-task/cmd/api/user/internal"
	"entain-golang-task/pkg"
	"entain-golang-task/pkg/utils"
	"errors"
	"net/http"
	"strconv"

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
	userIdStr := ps.ByName("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONError(s.logger, w, http.StatusBadRequest, utils.ErrInvalidUserId)
		return
	}

	var input internal.UserTransactionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJSONError(s.logger, w, http.StatusBadRequest, utils.ErrInvalidJsonBody)
		return
	}

	defer r.Body.Close()

	if input.State == "" || (input.State != pkg.StateLose && input.State != pkg.StateWin) {
		utils.WriteJSONError(s.logger, w, http.StatusBadRequest, utils.ErrIncorrectState)
		return
	}

	err = s.userHandler.Handle(context.TODO(), userId, input)
	if err != nil && !errors.Is(err, utils.ErrTransactionExists) {
		utils.WriteJSONError(s.logger, w, http.StatusInternalServerError, utils.ErrInternalServerErr)
		return
	}

	if err != nil {
		utils.WriteJSONError(s.logger, w, http.StatusBadRequest, err)
		return
	}

	//w.WriteHeader(http.StatusCreated) it's 200[OK] in the document, so...
	utils.WriteJSONMessage(w, http.StatusOK, "transaction processed successfully")
}
