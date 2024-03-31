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

func (server *GRPCServer) CreateAddress(ctx context.Context, req *api.CreateAddressRequest) (*api.AddressResponse, error) {
	params := sqlc.CreateAddressParams{
		Name: req.Name,
		Type: req.Type,
	}
	if req.Parent != 0 {
		params.Parent = pgtype.Int4{Int32: req.Parent, Valid: true}
	}
	rec, err := server.Store.Queries.CreateAddress(ctx, params)

	if err != nil {
		return nil, err
	}
	return &api.AddressResponse{Id: int32(rec.ID), Name: rec.Name, Type: rec.Type, Parent: rec.Parent.Int32}, nil
}

func (server *GRPCServer) GetAddressesByParent(ctx context.Context, req *api.GetAddressesByParentRequest) (*api.GetAddressesByParentResponse, error) {
	rec, err := server.Store.Queries.ListAddressesForParent(ctx, pgtype.Int4{Int32: req.Parent, Valid: true})
	if err != nil {
		return nil, err
	}

	fmt.Println(len(rec))
	var result []*api.AddressResponse
	for _, element := range rec {
		newRec := api.AddressResponse{Id: int32(element.ID), Name: element.Name, Type: element.Type, Parent: element.Parent.Int32}
		result = append(result, &newRec)
	}
	return &api.GetAddressesByParentResponse{Items: result}, nil

}

func (server *GRPCServer) GetAddressById(ctx context.Context, req *api.GetAddressByIdRequest) (*api.AddressResponse, error) {
	rec, err := server.Store.Queries.GetAddress(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &api.AddressResponse{Id: int32(rec.ID), Name: rec.Name, Type: rec.Type, Parent: rec.Parent.Int32}, nil
}

func (server *GRPCServer) DeleteAddress(ctx context.Context, req *api.DeleteAddressRequest) (*api.Empty, error) {
	err := server.Store.Queries.DeleteAddress(ctx, int64(req.Id))
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}
