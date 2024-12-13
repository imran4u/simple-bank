package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/imran4u/simple-bank/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := Server{
		store: store,
	}
	router := gin.Default()

	// Add handler
	router.POST("/account", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return &server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {

	return gin.H{"Error": err.Error()}
}
