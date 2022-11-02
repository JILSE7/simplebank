package api

import (
	db "github.com/JILSE7/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(s db.Store) *Server {
	server := &Server{store: s}
	router := gin.Default()

	router.POST("/accounts", server.createAccountFactory)
	router.GET("/accounts/:id", server.getAccountById)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
