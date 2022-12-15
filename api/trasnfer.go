package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	db "github.com/JILSE7/simplebank/db/sqlc"
	"github.com/JILSE7/simplebank/token"
	"github.com/gin-gonic/gin"
)

type tranferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency"  binding:"required,currency"`
}

func (s *Server) createTranferFactory(ctx *gin.Context) {
	var req tranferRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccount, isValid := s.isValidAccount(ctx, req.FromAccountID, req.Currency)

	if !isValid {
		return
	}

	// Get payload from the middleware
	authorizationPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authorizationPayload.Username != fromAccount.Owner {
		err := errors.New("from account doesn´t belong to the authenticated user")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, isValid = s.isValidAccount(ctx, req.ToAccountID, req.Currency)

	if !isValid {
		return
	}

	arg := db.TranferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	transfer, err := s.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, transfer)
}

func (s *Server) isValidAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := s.store.GetAccount(ctx, accountID)

	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true

}
