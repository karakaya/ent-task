package user

import (
	"context"
	"entain-golang-task/cmd/api-service/user/internal"
	"entain-golang-task/pkg/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/rs/zerolog"
)

type UserAccountBalanceService struct {
	logger                    zerolog.Logger
	userAccountBalanceHandler *internal.UserAccountBalanceHandler
}

func NewUserAccountBalanceService(logger zerolog.Logger) httprouter.Handle {
	service := &UserAccountBalanceService{
		logger:                    logger,
		userAccountBalanceHandler: internal.NewUserAccountBalanceHandler(logger),
	}

	return service.Handle
}

func (s *UserAccountBalanceService) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userIdStr := ps.ByName("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONError(s.logger, w, http.StatusBadRequest, utils.ErrInvalidUserId)
		return
	}

	accountBalanceOutput, err := s.userAccountBalanceHandler.GetAccountBalance(context.TODO(), userId)
	if err != nil {
		utils.WriteJSONError(s.logger, w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSONResponse(s.logger, w, http.StatusOK, accountBalanceOutput)
}
