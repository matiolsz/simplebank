package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/matiolsz/simplebank/db/sqlc"
	"github.com/matiolsz/simplebank/token"
)

type Server struct {
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(store db.Store) *Server {
	tokenMaker, err := token.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		return nil
		// , fmt.Errorf("cannot create token amker %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/entries", server.createEntry)
	router.GET("/entries/:id", server.getEntry)
	router.GET("/entries", server.listEntry)

	router.POST("/transfers", server.createTransfer)
	router.GET("/transfers/:id", server.getTransfer)
	router.GET("/transfers", server.listTransfer)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
