package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	db "github.com/JILSE7/simplebank/db/sqlc"
	"github.com/JILSE7/simplebank/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/phihdn/simplebank/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password"  binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UserResponse struct {
	Username        string    `json:"username" binding:"required,alphanum"`
	FullName        string    `json:"full_name" binding:"required"`
	Email           string    `json:"email" binding:"required,email"`
	PasswordChanged time.Time `json:"password_changed"`
	CreatedAt       time.Time `json:"created_at"`
}

func newUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username:        user.Username,
		FullName:        user.FullName,
		Email:           user.Email,
		PasswordChanged: user.PasswordChanged,
		CreatedAt:       user.CreatedAt,
	}
}

func (s *Server) createUserFactory(ctx *gin.Context) {
	var req createUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := s.store.CreateUser(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code.Name())

			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return

			}
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	userResponse := newUserResponse(user)

	ctx.JSON(http.StatusOK, userResponse)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password"  binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	res := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, res)
}
