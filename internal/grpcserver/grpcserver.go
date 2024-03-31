package grpcserver

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/maximprokopchuk/address_service/internal/api"
	"github.com/maximprokopchuk/address_service/internal/sqlc"
	"github.com/maximprokopchuk/address_service/internal/store"
)

type GRPCServer struct {
	Store *store.Store
}

func New(st *store.Store) *GRPCServer {
	return &GRPCServer{Store: st}
}

func (server *GRPCServer) CreateAddress(ctx context.Context, req *api.CreateAddressRequest) (*api.CreateAddressResponse, error) {
	queries := sqlc.New(server.Store.Connection)
	fmt.Println(req.Parent)
	params := sqlc.CreateAddressParams{
		Name: req.Name,
		Type: req.Type,
	}
	if req.Parent != 0 {
		params.Parent = pgtype.Int4{Int32: req.Parent, Valid: true}
	}
	rec, err := queries.CreateAddress(ctx, params)

	if err != nil {
		return nil, err
	}
	return &api.CreateAddressResponse{Id: int32(rec.ID), Name: rec.Name, Type: rec.Type, Parent: req.Parent}, nil
}
