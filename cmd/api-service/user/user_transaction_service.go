package user

import (
	"context"
	"encoding/json"
	"entain-golang-task/cmd/api-service/user/internal"
	"entain-golang-task/pkg"
	"entain-golang-task/pkg/core"
	"entain-golang-task/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/rs/zerolog"
)

type AccountTransactionService struct {
	logger                 zerolog.Logger
	userTransactionHandler *internal.UserTransactionHandler
}

func NewUserTransactionService(logger zerolog.Logger) httprouter.Handle {
	service := &AccountTransactionService{
		logger:                 logger,
		userTransactionHandler: internal.NewUserTransactionHandler(logger),
	}

	return service.Handle
}

func (s *AccountTransactionService) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	//will not check for if amount 0
	validatedAmount, err := core.ValidateTransactionAmount(input.Amount)
	if err != nil {
		utils.WriteJSONError(s.logger, w, http.StatusBadRequest, utils.ErrInvalidAmount)
		return
	}

	input.Amount = validatedAmount

	if input.State == "" || (input.State != pkg.StateLose && input.State != pkg.StateWin) {
		utils.WriteJSONError(s.logger, w, http.StatusBadRequest, utils.ErrInvalidState)
		return
	}

	accountTransactionOutput, err := s.userTransactionHandler.SaveUserTransaction(context.TODO(), userId, input)
	//might not be sustainable
	if err != nil && (!errors.Is(err, utils.ErrTransactionExists) && !errors.Is(err, utils.ErrAccountBalanceCannotBeNegative)) {
		utils.WriteJSONError(s.logger, w, http.StatusInternalServerError, utils.ErrInternalServerErr)
		return
	}

	if err != nil {
		utils.WriteJSONError(s.logger, w, http.StatusBadRequest, err)
		return
	}

	//w.WriteHeader(http.StatusCreated) it's 200[OK] in the document, so...
	utils.WriteJSONResponse(s.logger, w, http.StatusOK, accountTransactionOutput)
}
