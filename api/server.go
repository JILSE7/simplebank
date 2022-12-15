package api

import (
	"fmt"

	db "github.com/JILSE7/simplebank/db/sqlc"
	"github.com/JILSE7/simplebank/token"
	"github.com/JILSE7/simplebank/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     utils.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config utils.Config, s db.Store) (*Server, error) {
	token, err := token.NewPasetoMaker(config.TokenSymmetrictKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      s,
		tokenMaker: token,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUserFactory)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccountFactory)
	authRoutes.GET("/accounts/:id", server.getAccountById)
	authRoutes.GET("/accounts", server.listAccounts)

	authRoutes.POST("/transfer", server.createTranferFactory)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
