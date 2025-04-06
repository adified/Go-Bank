package gapi

import (
	"fmt"

	db "github.com/adified/Go-Bank/db/sqlc"
	"github.com/adified/Go-Bank/pb"
	"github.com/adified/Go-Bank/token"
	"github.com/adified/Go-Bank/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSelfBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
